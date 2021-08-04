import Vue from 'vue'
import web3 from '../../lib/eth/web3'
import abiContract from '../../lib/eth/eth'
import exchange from '../../lib/eth/exchange'
import Tx from 'ethereumjs-tx'
import BigNumber from 'bignumber.js'
import {handleEthPrice, numFloat} from '../../utils/handleFloat'
import store from '../../store/index'

export default {
  bdeRevokeOrder: async function (orderId) {
    const userPrivateKey = store.state.account.myEth.privateKey
    const useraddress = store.state.account.myEth.address
    const contractdata = await exchange.methods.withdrawal(web3.utils.toHex(orderId))
    const payload = contractdata.encodeABI()
    const gas = await web3.eth.getGasPrice()
    const nonce = await web3.eth.getTransactionCount(useraddress)
    const gasPriceHex = web3.utils.numberToHex(gas)
    const nonceHex = web3.utils.numberToHex(nonce)
    const gasLimitHex = web3.utils.numberToHex(300000)

    const rawTx = {
      nonce: nonceHex,
      gasPrice: gasPriceHex,
      gasLimit: gasLimitHex,
      value: '0x00',
      from: useraddress,
      to: exchange.address,
      data: payload
    }
    const tx = new Tx(rawTx)
    const privateKey = Buffer.from(userPrivateKey.slice(2), 'hex')
    tx.sign(privateKey)
    const serializedTx = tx.serialize()
    Vue.prototype.$socket.sendObj({
      namespace: 'user',
      action: 'withdrawbde',
      data: {
        orderId: orderId,
        approveTx: '0x' + serializedTx.toString('hex')
      }
    })
  },
  revokeOrder: async function (orderId) {
    const userPrivateKey = store.state.account.myEth.privateKey
    const useraddress = store.state.account.myEth.address
    const contractdata = await exchange.methods.withdrawal(web3.utils.toHex(orderId))
    const payload = contractdata.encodeABI()
    const gas = await web3.eth.getGasPrice()
    const nonce = await web3.eth.getTransactionCount(useraddress)
    const gasPriceHex = web3.utils.numberToHex(gas)
    const nonceHex = web3.utils.numberToHex(nonce)
    const gasLimitHex = web3.utils.numberToHex(300000)

    const rawTx = {
      nonce: nonceHex,
      gasPrice: gasPriceHex,
      gasLimit: gasLimitHex,
      value: '0x00',
      from: useraddress,
      to: exchange.address,
      data: payload
    }
    const tx = new Tx(rawTx)
    const privateKey = Buffer.from(userPrivateKey.slice(2), 'hex')
    tx.sign(privateKey)
    const serializedTx = tx.serialize()
    Vue.prototype.$socket.sendObj({
      namespace: 'user',
      action: 'withdraweth',
      data: {
        orderId: orderId,
        approveTx: '0x' + serializedTx.toString('hex')
      }
    })
  },
  ethApproveAmount: async function (tokenName) { // eth查询erc20代币已授权的数量  测试通过
    const eth = store.state.api.eth
    const useraddress = store.state.account.myEth.address
    const myContract = exchange.address
    const contractAccount = eth[tokenName].contractAddress.address
    const coin = new web3.eth.Contract(abiContract, contractAccount)
    let quota = await coin.methods.allowance(useraddress, myContract).call()
    let approveQuota = new BigNumber(quota)
    const dec = eth[tokenName].contractAddress.decimal
    const accreditAmount = parseFloat(approveQuota / Math.pow(10, dec))
    return accreditAmount
  },
  approveToken: async function (tokenName, accreditAmount) { // 授权
    const eth = store.state.api.eth
    const useraddress = store.state.account.myEth.address
    const myContract = exchange.address
    const contractAccount = eth[tokenName].contractAddress.address
    const gas = await web3.eth.getGasPrice()
    const gasPriceHex = web3.utils.numberToHex(gas)
    const gasLimitHex = web3.utils.numberToHex(300000)
    const coin = new web3.eth.Contract(abiContract, contractAccount)
    let amount = new BigNumber(accreditAmount * (10 ** eth[tokenName].contractAddress.decimal))
    try {
      coin.methods.approve(myContract, web3.utils.toHex(amount)).send({
        from: useraddress,
        gasPrice: gasPriceHex,
        gas: gasLimitHex
      })
      return 'Processing'
    } catch (e) {
      return false
    }
  },
  buy: async function (price, total, tokenName) { // eth 买单
    const userPrivateKey = store.state.account.myEth.privateKey
    const useraddress = store.state.account.myEth.address
    const receiver = store.state.account.myEth.address
    const eth = store.state.api.eth
    const data = {}
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
    data.tradetype = false // 以太坊链内交易买单为false，卖单true
    data.amount = await web3.utils.toWei(total.toString(), 'ether')
    data.sellTokenAddress = '0x0000000000000000000000000000000000000000' // 卖出的token的地址
    data.price = handleEthPrice(price)
    if (!data.price) {
      return 'error'
    }
    data.buyTokenAddress = eth[tokenName].contractAddress.address // 买入的token的地址
    // todo 8.20 接口改动
    // const dateTime = Date.now()
    // data.createTime = Math.floor(dateTime / 1000)
    const contractdata = await exchange.methods.depositsOrder(web3.utils.toHex(data.id), receiver, data.tradetype, web3.utils.toHex(data.amount), data.sellTokenAddress, data.price, data.buyTokenAddress, 2)
    const payload = contractdata.encodeABI()
    const gas = await web3.eth.getGasPrice()
    const nonce = await web3.eth.getTransactionCount(useraddress)
    const gasPriceHex = web3.utils.numberToHex(gas)
    const nonceHex = web3.utils.numberToHex(nonce)
    const gasLimitHex = web3.utils.numberToHex(300000)

    const rawTx = {
      nonce: nonceHex,
      gasPrice: gasPriceHex,
      gasLimit: gasLimitHex,
      value: web3.utils.toHex(data.amount),
      from: useraddress,
      to: exchange.address,
      data: payload
    }
    const tx = new Tx(rawTx)
    const privateKey = Buffer.from(userPrivateKey.slice(2), 'hex')
    tx.sign(privateKey)
    const serializedTx = tx.serialize()
    Vue.prototype.$socket.sendObj({
      namespace: 'user',
      action: 'sendeth',
      data: {
        approveTx: '0x' + serializedTx.toString('hex'),
        marketType: 2,
        tradetype: false
      }
    })
    return 'Processing' // 订单正在发送处理中
  },
  sell: async function (price, amount, tokenName) { // eth 卖单
    // const haveAccreditAmount = this.ethApproveAmount(tokenName)
    // if (amount > haveAccreditAmount) {
    //   console.error('eth卖单，授权额度过低===============')
    //   return 'error'
    // }
    const userPrivateKey = store.state.account.myEth.privateKey
    const useraddress = store.state.account.myEth.address
    const receiver = store.state.account.myEth.address
    const myContract = exchange.address
    const eth = store.state.api.eth
    const data = {}
    data.contractAccount = eth[tokenName].contractAddress.address
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
    data.price = handleEthPrice(price)
    if (!data.price) {
      return 'error'
    }
    if (eth[tokenName].contractAddress.decimal < 4) {
      amount = numFloat(amount, eth[tokenName].contractAddress.decimal)
    }
    data.sellTokenAddress = eth[tokenName].contractAddress.address // 卖出的token的地址
    data.buyTokenAddress = '0x0000000000000000000000000000000000000000' // 买入的token的地址
    data.amount = new BigNumber(amount * (10 ** eth[tokenName].contractAddress.decimal))
    data.tradetype = true // 以太坊链内交易买单为false，跨链为true
    // const dateTime = Date.now()
    // data.createTime = Math.floor(dateTime / 1000)
    const gas = await web3.eth.getGasPrice()
    const nonce = await web3.eth.getTransactionCount(useraddress)
    const gasPriceHex = web3.utils.numberToHex(gas)
    const nonceHex = web3.utils.numberToHex(nonce)
    const gasLimitHex = web3.utils.numberToHex(400000)
    const contractdata = await exchange.methods.depositsOrder(web3.utils.toHex(data.id), receiver, data.tradetype, web3.utils.toHex(data.amount), data.sellTokenAddress, data.price, data.buyTokenAddress, 2)
    const payload = contractdata.encodeABI()
    const rawTx = {
      nonce: nonceHex,
      gasPrice: gasPriceHex,
      gasLimit: gasLimitHex,
      value: '0x00',
      from: useraddress,
      to: myContract,
      data: payload
    }
    const tx = new Tx(rawTx)
    const privateKey = Buffer.from(userPrivateKey.slice(2), 'hex')
    tx.sign(privateKey)
    const serializedTx = tx.serialize()
    Vue.prototype.$socket.sendObj({
      namespace: 'user',
      action: 'sendeth',
      data: {
        tradeType: true,
        marketType: 2,
        approveTx: '0x' + serializedTx.toString('hex')
      }
    })
    return 'Processing' // 订单正在发送处理中
  },
  bdeBuy: async function (price, total, tokenName) { // bde 交易区买单
    // const haveAccreditAmount = this.ethApproveAmount('BDE')
    // if (total > haveAccreditAmount) {
    //   console.error('bde买单，授权额度过低===============')
    //   return 'error'
    // }
    const userPrivateKey = store.state.account.myEth.privateKey
    const useraddress = store.state.account.myEth.address
    const receiver = store.state.account.myEth.address
    const myContract = exchange.address
    const eth = store.state.api.eth
    const contractAccountBde = eth['BDE'].contractAddress.address // bde的合约地址
    const data = {}
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
    const bdeDecimal = eth['BDE'].contractAddress.decimal
    data.total = 0
    if (bdeDecimal < 4) {
      let lowFloatTotal = numFloat(total, bdeDecimal)
      data.total = new BigNumber(lowFloatTotal * (10 ** eth['BDE'].contractAddress.decimal))
    } else {
      data.total = new BigNumber(total * (10 ** eth['BDE'].contractAddress.decimal))
    }
    data.tradetype = false // 以太坊链内交易买单为false，卖单true
    data.sellTokenAddress = contractAccountBde
    data.buyTokenAddress = eth[tokenName].contractAddress.address
    data.price = handleEthPrice(price)
    // const dateTime = Date.now()
    // data.createTime = Math.floor(dateTime / 1000)
    const gas = await web3.eth.getGasPrice()
    const nonce = await web3.eth.getTransactionCount(useraddress)
    const gasPriceHex = web3.utils.numberToHex(gas)
    const nonceHex = web3.utils.numberToHex(nonce)
    const gasLimitHex = web3.utils.numberToHex(400000)
    // todo 暂时使用，后面会改接口
    console.log('开始调用合约， 执行bde操下单操作')
    const contractdata = await exchange.methods.depositsOrder(web3.utils.toHex(data.id), receiver, data.tradetype, web3.utils.toHex(data.total), data.sellTokenAddress, data.price, data.buyTokenAddress, 3)
    const payload = contractdata.encodeABI()

    const rawTx = {
      nonce: nonceHex,
      gasPrice: gasPriceHex,
      gasLimit: gasLimitHex,
      value: '0x00',
      from: useraddress,
      to: myContract,
      data: payload
    }
    const tx = new Tx(rawTx)
    const privateKey = Buffer.from(userPrivateKey.slice(2), 'hex')
    tx.sign(privateKey)
    const serializedTx = tx.serialize()
    Vue.prototype.$socket.sendObj({
      namespace: 'user',
      action: 'sendbde',
      data: {
        tradeType: false, // 1 || true : 卖单 0 || false: 买单
        marketType: 3, // 对应于marketCode
        approveTx: '0x' + serializedTx.toString('hex')
      }
    })
    return 'Processing' // 订单正在发送处理中
  },
  bdeSell: async function (price, amount, tokenName) { // bde 交易区卖单
    // const haveAccreditAmount = this.ethApproveAmount(tokenName)
    // if (amount > haveAccreditAmount) {
    //   console.error('eth卖单，授权额度过低===============')
    //   return 'error'
    // }
    const userPrivateKey = store.state.account.myEth.privateKey
    const useraddress = store.state.account.myEth.address
    const receiver = store.state.account.myEth.address
    const myContract = exchange.address
    const eth = store.state.api.eth
    const data = {}
    data.contractAccount = eth[tokenName].contractAddress.address
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
    data.price = handleEthPrice(price) // todo 注意下，bde的价格处理方式可能有问题
    if (!data.price) {
      return 'error'
    }
    data.sellTokenAddress = eth[tokenName].contractAddress.address
    data.buyTokenAddress = eth['BDE'].contractAddress.address
    data.amount = new BigNumber(amount * (10 ** eth[tokenName].contractAddress.decimal))
    if (eth[tokenName].contractAddress.decimal < 4) {
      let newAmount = numFloat(amount, eth[tokenName].contractAddress.decimal)
      data.amount = new BigNumber(newAmount * (10 ** eth[tokenName].contractAddress.decimal))
    }
    data.tradetype = true // 以太坊链内交易买单为false，跨链为true
    const dateTime = Date.now()
    data.createTime = Math.floor(dateTime / 1000)
    const gas = await web3.eth.getGasPrice()
    const nonce = await web3.eth.getTransactionCount(useraddress)
    const gasPriceHex = web3.utils.numberToHex(gas)
    const nonceHex = web3.utils.numberToHex(nonce)
    const gasLimitHex = web3.utils.numberToHex(400000)
    // todo bde交易区接口待定，暂时先如此处理
    const contractdata = await exchange.methods.depositsOrder(web3.utils.toHex(data.id), receiver, data.tradetype, web3.utils.toHex(data.amount), data.sellTokenAddress, data.price, data.buyTokenAddress, 3)
    const payload = contractdata.encodeABI()
    const rawTx = {
      nonce: nonceHex,
      gasPrice: gasPriceHex,
      gasLimit: gasLimitHex,
      value: '0x00',
      from: useraddress,
      to: myContract,
      data: payload
    }
    const tx = new Tx(rawTx)
    const privateKey = Buffer.from(userPrivateKey.slice(2), 'hex')
    tx.sign(privateKey)
    const serializedTx = tx.serialize()
    Vue.prototype.$socket.sendObj({
      namespace: 'user',
      action: 'sendeth',
      data: {
        tradeType: true,
        marketType: 2,
        approveTx: '0x' + serializedTx.toString('hex')
      }
    })
    return 'Processing' // 订单正在发送处理中
  },
  ethBuyEos: async function (price, amount) {
    const userPrivateKey = store.state.account.myEth.privateKey
    const useraddress = store.state.account.myEth.address
    const receiver = store.state.account.myEos.address
    const data = {}
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
    data.tradetype = true // 以太坊链内交易买单为false，卖单true,跨链true
    data.amount = await web3.utils.toWei(amount.toString(), 'ether')
    data.price = handleEthPrice(price)
    if (!data.price) {
      return 'error'
    }
    data.sellTokenAddress = '0x0000000000000000000000000000000000000000'
    data.buyTokenAddress = '0x0000000000000000000000000000000000000000'
    // const dateTime = Date.now()
    // data.createTime = Math.floor(dateTime / 1000)
    const contractdata = await exchange.methods.depositsOrder(web3.utils.toHex(data.id), receiver, data.tradetype, web3.utils.toHex(data.amount), data.sellTokenAddress, data.price, data.buyTokenAddress, 4)
    const payload = contractdata.encodeABI()
    const gas = await web3.eth.getGasPrice()
    const nonce = await web3.eth.getTransactionCount(useraddress)
    const gasPriceHex = web3.utils.numberToHex(gas)
    const nonceHex = web3.utils.numberToHex(nonce)
    const gasLimitHex = web3.utils.numberToHex(300000)

    const rawTx = {
      nonce: nonceHex,
      gasPrice: gasPriceHex,
      gasLimit: gasLimitHex,
      value: web3.utils.toHex(data.amount),
      from: useraddress,
      to: exchange.address,
      data: payload
    }
    const tx = new Tx(rawTx)
    const privateKey = Buffer.from(userPrivateKey.slice(2), 'hex')
    tx.sign(privateKey)
    const serializedTx = tx.serialize()
    Vue.prototype.$socket.sendObj({
      namespace: 'user',
      action: 'sendeth',
      data: {
        approveTx: '0x' + serializedTx.toString('hex'),
        marketType: 4
      }
    })
    return 'Processing' // 订单正在发送处理中
  }
}
