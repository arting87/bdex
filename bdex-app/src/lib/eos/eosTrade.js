import store from '../../store/index'
import {handBig, numFloat} from '../../utils/handleFloat'
import Vue from 'vue'
import Eos from 'eosjs'

// const isProduction = process.env.NODE_ENV === 'production' // true 为生产模式
// const chainId = isProduction ? 'aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906' : '6cbecff836a9fa60da53bf97a0a180103b2e76041d4414693d11bf39e2341547'
// const httpEndpoint = isProduction ? 'https://eos.greymass.com' : 'http://47.103.108.61:8888'
// const myEosContract = isProduction ? 'bdexcontract' : 'tbcexchange1'
// const eosContractAcc = 'eosio.token'

const chainId = 'aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906'
const httpEndpoint = 'https://eos.greymass.com'
const myEosContract = '' // TODO 替换部署账户名
const eosContractAcc = 'eosio.token'

export default {
  revokeOrder: async function (orderId) {
    const privateKey = store.state.account.myEos.privateKey
    const currentAccount = store.state.account.myEos.address
    const config = {
      chainId: chainId, // 32 byte (64 char) hex string
      keyProvider: [privateKey], // WIF string or array of keys..
      httpEndpoint: httpEndpoint,
      expireInSeconds: 60,
      broadcast: false, // 只签名，不发送
      sign: true
    }
    const eos = Eos(config)

    const options = {
      authorization: `${currentAccount}@active`,
      broadcast: false,
      sign: true
    }
    try {
      let contract = await eos.contract(myEosContract)
      let wdSign = await contract.withdraw(currentAccount, orderId, options) // wdSign撤单签名后的数据
      delete wdSign.transaction.compression
      wdSign = wdSign.transaction
      if (wdSign) {
        Vue.prototype.$socket.sendObj({
          namespace: 'user',
          action: 'withdraweos',
          data: {
            orderId: orderId,
            approveTx: JSON.stringify(wdSign)
          }
        })
      }
      return 'revoking'
    } catch (error) {
      return 'error'
    }
  },
  buy: async function (price, total, tokenName) { // 买单
    const privateKey = store.state.account.myEos.privateKey
    const currentAccount = store.state.account.myEos.address
    const data = {}
    const config = {
      chainId: chainId, // 32 byte (64 char) hex string
      keyProvider: [privateKey], // WIF string or array of keys..
      httpEndpoint: httpEndpoint,
      expireInSeconds: 60,
      broadcast: false, // 只签名，不发送
      sign: true
    }
    const eos = Eos(config)

    const apiEos = store.state.api.eos
    const idArr = store.state.user.idArr
    data.id = idArr[0] // string
    // 更新订单id
    if (idArr.length === 1) {
      Vue.prototype.$socket.sendObj({
        namespace: 'user',
        action: 'getid',
        data: ''
      })
    }
    store.commit('user/delId', 0)
    data.a = '0'
    data.p = handBig.times(price, 10 ** 8) + ''
    const decimal = apiEos[tokenName].contractAddress.decimal + ''
    data.bs = tokenName + ',' + decimal
    data.bc = apiEos[tokenName].contractAddress.address + ''
    data.ta = currentAccount
    const memo = `{"i": "${data.id}", "a": "${data.a}", "p": "${data.p}", "bs": "${data.bs}", "bc": "${data.bc}", "ta": "${data.ta}"}`
    const options = {
      authorization: `${currentAccount}@active`,
      broadcast: false,
      sign: true
    }
    const quantity = numFloat(total, 4) + ' EOS'
    try {
      const contract = await eos.contract(eosContractAcc)
      let trxSign = await contract.transfer(currentAccount, myEosContract, quantity, memo, options)
      delete trxSign.transaction.compression
      trxSign = trxSign.transaction
      if (trxSign) {
        Vue.prototype.$socket.sendObj({
          namespace: 'user',
          action: 'sendeos',
          data: {
            order: JSON.stringify(trxSign),
            marketCode: '1',
            chainCode: '1',
            targetCode: '1'
          }
        })
      }
      return 'Processing'
    } catch (e) {
      console.error(e)
      return 'error'
    }
  },
  sell: async function (price, amount, tokenName) { // 卖单
    const privateKey = store.state.account.myEos.privateKey
    const currentAccount = store.state.account.myEos.address
    const data = {}
    const config = {
      chainId: chainId, // 32 byte (64 char) hex string
      keyProvider: [privateKey], // WIF string or array of keys..
      httpEndpoint: httpEndpoint,
      expireInSeconds: 60,
      broadcast: false, // 只签名，不发送
      sign: true
    }
    const eos = Eos(config)

    const apiEos = store.state.api.eos
    const idArr = store.state.user.idArr
    data.id = idArr[0] // string
    // 更新订单id
    if (idArr.length === 1) {
      Vue.prototype.$socket.sendObj({
        namespace: 'user',
        action: 'getid',
        data: ''
      })
    }
    store.commit('user/delId', 0)
    data.a = '1'
    data.p = handBig.times(price, 10 ** 8) + ''
    data.bs = 'EOS,4' // buy token name and float
    data.bc = 'eosio.token'
    data.ta = currentAccount
    const memo = `{"i": "${data.id}", "a": "${data.a}", "p": "${data.p}", "bs": "${data.bs}", "bc": "${data.bc}", "ta": "${data.ta}"}`
    const decimal = apiEos[tokenName].contractAddress.decimal
    const quantity = (numFloat(amount, decimal)) + ' ' + tokenName
    const contractAccount = apiEos[tokenName].contractAddress.address
    const options = {
      authorization: `${currentAccount}@active`,
      broadcast: false,
      sign: true
    }
    try {
      const contract = await eos.contract(eosContractAcc)
      let trxSign = await contract.transfer(contractAccount, myEosContract, quantity, memo, options)
      delete trxSign.transaction.compression
      trxSign = trxSign.transaction
      if (trxSign) {
        Vue.prototype.$socket.sendObj({
          namespace: 'user',
          action: 'sendeos',
          data: {
            order: JSON.stringify(trxSign),
            marketCode: '1',
            chainCode: '1',
            targetCode: '1'
          }
        })
      }
      return 'Processing'
    } catch (e) {
      console.error(e)
      return 'error'
    }
  },
  eosBuyEth: async function (price, total) { // 跨链买单，eos买eth
    const privateKey = store.state.account.myEos.privateKey
    const currentAccountEos = store.state.account.myEos.address
    const currentAccountEth = store.state.account.myEth.address
    const data = {}
    const config = {
      chainId: chainId, // 32 byte (64 char) hex string
      keyProvider: [privateKey], // WIF string or array of keys..
      httpEndpoint: httpEndpoint,
      expireInSeconds: 60,
      broadcast: false, // 只签名，不发送
      sign: true
    }
    const eos = Eos(config)

    const idArr = store.state.user.idArr
    data.id = idArr[0] // string
    // 更新订单id
    if (idArr.length === 1) {
      Vue.prototype.$socket.sendObj({
        namespace: 'user',
        action: 'getid',
        data: ''
      })
    }
    store.commit('user/delId', 0)
    data.a = '0'
    data.p = handBig.times(price, 10 ** 8) + ''
    data.bs = 'ETH' + ',' + '8'
    data.bc = '0x0000000000000000000000000000000000000000'
    data.ta = currentAccountEth
    const memo = `{"i": "${data.id}", "a": "${data.a}", "p": "${data.p}", "bs": "${data.bs}", "bc": "${data.bc}", "ta": "${data.ta}"}`
    const options = {
      authorization: `${currentAccountEos}@active`,
      broadcast: false,
      sign: true
    }
    const quantity = numFloat(total, 4) + ' EOS'
    try {
      const contract = await eos.contract(eosContractAcc)
      let trxSign = await contract.transfer(currentAccountEos, myEosContract, quantity, memo, options)
      delete trxSign.transaction.compression
      trxSign = trxSign.transaction
      if (trxSign) {
        Vue.prototype.$socket.sendObj({
          namespace: 'user',
          action: 'sendeos',
          data: {
            order: JSON.stringify(trxSign),
            marketCode: '4',
            chainCode: '1',
            targetCode: '2'
          }
        })
      }
      return 'Processing'
    } catch (e) {
      console.error(e)
      return 'error'
    }
  }
}
