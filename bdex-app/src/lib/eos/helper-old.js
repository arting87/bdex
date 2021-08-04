import { Api, JsonRpc } from 'eosjs'
import { JsSignatureProvider } from 'eosjs/dist/eosjs-jssig'

// const rpc = new JsonRpc('https://eos.greymass.com') test:http://47.103.108.61:8888
const rpc = new JsonRpc('http://47.103.108.61:8888')

async function getKeyAccounts (pk) {
  return rpc.history_get_key_accounts(pk)
}

async function getCurrentBalance (name) {
  return rpc.get_currency_balance('eosio.token', name, 'EOS')
}

async function getErc20Balance (code, name, symbol) { // 8.16 查询指定eos代币的余额
  return rpc.get_currency_balance(code, name, symbol)
}

async function sendEos (from, to, pk, amount, memo) {
  const defaultPrivateKey = pk
  const signatureProvider = new JsSignatureProvider([defaultPrivateKey])
  const api = new Api({rpc, signatureProvider})
  const option = {
    actions: [{
      account: 'eosio.token',
      name: 'transfer',
      authorization: [{
        actor: from,
        permission: 'active'
      }],
      data: {
        from: from,
        to: to,
        quantity: parseFloat(amount).toFixed(4) + ' EOS',
        memo: memo
      }
    }]
  }
  const params = {
    blocksBehind: 3,
    expireSeconds: 30
  }

  return api.transact(option, params)
}

async function buyRam (from, pk) {
  const defaultPrivateKey = pk
  const signatureProvider = new JsSignatureProvider([defaultPrivateKey])
  const api = new Api({rpc, signatureProvider})
  return api.transact({
    actions: [{
      account: 'eosio',
      name: 'buyrambytes',
      authorization: [{
        actor: from,
        permission: 'active'
      }],
      data: {
        payer: from,
        receiver: from,
        bytes: 8192
      }
    }]
  }, {
    blocksBehind: 3,
    expireSeconds: 30
  })
}

export { rpc, getKeyAccounts, sendEos, buyRam, getCurrentBalance, getErc20Balance }
