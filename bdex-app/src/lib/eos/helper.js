import Eos from 'eosjs'

// const isProduction = process.env.NODE_ENV === 'production' // true 为生产模式
// const chainId = isProduction ? 'aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906' : '6cbecff836a9fa60da53bf97a0a180103b2e76041d4414693d11bf39e2341547'
// const httpEndpoint = isProduction ? 'https://eos.greymass.com' : 'http://47.103.108.61:8888' // https://eos.greymass.com  http://47.103.108.61:8888

const chainId = 'aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906'
const httpEndpoint = 'https://eos.greymass.com'

const config = {
  chainId: chainId, // 32 byte (64 char) hex string
  // keyProvider: [privateKey], // WIF string or array of keys..
  httpEndpoint: httpEndpoint,
  expireInSeconds: 30,
  broadcast: true,
  sign: true
}

const eos = Eos(config)

async function getKeyAccounts (pk) { // 获取账号地址
  try {
    const result = await eos.getKeyAccounts(pk)
    return result
  } catch (e) {
    console.error(e)
  }
}

async function getAccountInfo (name) {
  const result = await eos.getAccount(name)
  return result
}

async function getCurrentBalance (name) {
  const result = await eos.getCurrencyBalance('eosio.token', name, 'EOS')
  return result
}

async function getErc20Balance (code, name, symbol) {
  try {
    const result = await eos.getCurrencyBalance(code, name, symbol)
    let balance = result[0]
    balance = balance.split(' ')[0]
    return balance
  } catch (e) {
    console.error(e)
    return 'error'
  }
}

async function sendEos (from, to, pk, amount, memo) {
  const keyProvider = pk
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
    keyProvider: keyProvider,
    blocksBehind: 3
  }
  const transaction = eos.transaction(option, params)
  return transaction
}

async function buyRam (from, pk, receiver, ramAmount) {
  const keyProvider = pk
  const option = {
    actions: [{
      account: 'eosio',
      name: 'buyrambytes',
      authorization: [{
        actor: from,
        permission: 'active'
      }],
      data: {
        payer: from,
        receiver: receiver,
        bytes: ramAmount
      }
    }]
  }
  const params = {
    keyProvider: keyProvider,
    blocksBehind: 3
  }
  const transacton = await eos.transaction(option, params)
  return transacton
}

async function sellRam (from, pk, ramAmount) {
  const keyProvider = pk
  const option = {
    actions: [{
      account: 'eosio',
      name: 'sellram',
      authorization: [{
        actor: from,
        permission: 'active'
      }],
      data: {
        account: from,
        bytes: 8192
      }
    }]
  }
  const params = {
    keyProvider: keyProvider,
    blocksBehind: 3
  }
  const transacton = await eos.transaction(option, params)
  return transacton
}

async function buyResource (from, pk, netAmount, cpuAmount, receiver) {
  const keyProvider = pk
  const option = {
    actions: [{
      account: 'eosio',
      name: 'delegatebw',
      authorization: [{
        actor: from,
        permission: 'active'
      }],
      data: {
        from: from,
        receiver: receiver,
        stake_net_quantity: parseFloat(netAmount).toFixed(4) + ' EOS',
        stake_cpu_quantity: parseFloat(cpuAmount).toFixed(4) + ' EOS',
        transfer: false
      }
    }]
  }
  const params = {
    keyProvider: keyProvider,
    blocksBehind: 3
  }
  const transacton = await eos.transaction(option, params)
  return transacton
}

export {getKeyAccounts, sendEos, buyRam, getCurrentBalance, getErc20Balance, getAccountInfo, buyResource, sellRam}
