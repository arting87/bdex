package core

import (
	"encoding/hex"
	"encoding/json"
	"time"

	"bdex.in/bdex/bdex-ex-backend/bdexerrors"
	eosContract "bdex.in/bdex/bdex-ex-backend/contract/eos"
	"bdex.in/bdex/bdex-ex-backend/core/matchpool"
	"bdex.in/bdex/bdex-ex-backend/utils"

	"github.com/GyhzzZ/eos-go"
	"github.com/GyhzzZ/eos-go/ecc"
)

//解包交易
func UnPackEosTransaction(tx string) (*ReqBody, []byte, error) {
	log := log.Context("UnPackEosTransaction")
	reqBody := &ReqBody{}
	err := json.Unmarshal([]byte(tx), reqBody)
	if err != nil {
		log.Errorw("JSON unmarshal transaction", "error", err)
		return nil, nil, err
	}
	if reqBody.Transaction == nil {
		return reqBody, nil, bdexerrors.ErrBadRequest
	}
	if len(reqBody.Transaction.Actions) < 1 {
		return reqBody, nil, bdexerrors.ErrBadRequest
	}
	b, ok := reqBody.Transaction.Actions[0].Data.(string)
	if !ok {
		return reqBody, nil, bdexerrors.ErrBadRequest
	}
	c, err := hex.DecodeString(b)
	if err != nil {
		return reqBody, nil, err
	}
	return reqBody, c, nil
}

//发送交易
func SendEosTransaction(transaction *eos.Transaction, signatures []ecc.Signature) (*eos.PushTransactionFullResp, error) {
	log := log.Context("SendEosTransaction")

	signedTx := &eos.SignedTransaction{
		Transaction: transaction,
		Signatures:  signatures,
	}
	packedTx, err := signedTx.Pack(eos.CompressionNone)
	if err != nil {
		log.Errorw("Pack transaction", "error", err)
		return nil, err
	}

	r, err := EXCore.EOS.PushTransaction(packedTx)
	if err != nil {
		log.Errorw("Push transaction", "error", err)
		return nil, err
	}

	return r, nil
}

//验证交易签名
func (s *ExService) verifySignature(account eos.AccountName, create *time.Time, body, signstr string) error {
	log := log.Context("verifySignature")
	if create.Before(time.Now().UTC()) {
		log.Errorw("request create time expired")
		return bdexerrors.ErrBadRequest
	}
	signature, err := ecc.NewSignature(signstr)
	if err != nil {
		log.Errorw("new signature fail", "error", err)
		return err
	}
	accountResp, err := s.EOS.GetAccount(account)
	if err != nil {
		log.Errorw("Can`t get account from chain")
		return bdexerrors.ErrInvalidName
	}
	var pubKey ecc.PublicKey
	for _, permiss := range accountResp.Permissions {
		if permiss.PermName == "active" && len(permiss.RequiredAuth.Keys) > 0 {
			pubKey = permiss.RequiredAuth.Keys[0].PublicKey
			break
		}
	}
	if pubKey.Content == nil {
		log.Errorw("Can`t get account pub key ")
		return bdexerrors.ErrInvalidName
	}

	signResult := signature.Verify(utils.SigDigest([]byte(body)), pubKey)
	if !signResult {
		log.Errorw("verify fail")
		return bdexerrors.ErrUnauthorized
	}
	return nil
}

func BackEosOrder(order *matchpool.CommonOrder, chainnum string) (string, error) {

	backOrder := eosContract.NewBackorder(order.Oid)

	txresp, err := eosContract.BuildTx(backOrder, chainnum)
	if err != nil {
		log.Errorw("Send Submit to "+chainnum+" Chain", "ErrorInMem", err)
		return "", err
	}

	return txresp.TransactionID, nil
}
