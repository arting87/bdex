package sync

import (
	"encoding/json"

	"github.com/GyhzzZ/eos-go"
)

const (
	ContractCode  = "bdexcontract"
	ContractScope = "bdexcontract"
)

func (s *Syncer) getTokenTable(lowerBound string, limit uint32) ([]*TokenTable, error) {
	log := log.Context("getTokenTable")
	assetTable, err := s.eosClient.GetTableRows(eos.GetTableRowsRequest{
		JSON:       true,
		Code:       ContractCode,
		Scope:      ContractScope,
		Table:      "token",
		LowerBound: lowerBound,
		Limit:      limit,
	})

	if err != nil {
		log.Errorw("Get token table", "error", err)
		return nil, err
	}
	var tokens []*TokenTable
	err = json.Unmarshal(assetTable.Rows, &tokens)

	if err != nil {
		log.Errorw("JSON unmarshal token", "error", err)
		return nil, err
	}

	return tokens, nil
}
