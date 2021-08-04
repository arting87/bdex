import ScatterJS from 'scatterjs-core'
import ScatterEOS from 'scatterjs-plugin-eosjs'
import Eos from 'eosjs'
import store from '../../store/index'
import {myEosNet} from '../../lib/eos/config'

// 注入ScatterEOS插件至ScatterJS中
// 注入ScatterWeb3插件至ScatterJS中
ScatterJS.plugins(new ScatterEOS())

// 注意 bug:是否要在关闭窗口网页时默认forgetIdentity,目前这里没做
// 目前是单账号登录，操作;
// 有冗余逻辑判断，后期再优化吧！保险的点，先这样吧！

// const myEosNet = {
//   blockchain: 'eos',
//   protocol: 'http',
//   host: '47.102.110.128',
//   port: 8888,
//   chainId: '6cbecff836a9fa60da53bf97a0a180103b2e76041d4414693d11bf39e2341547'
// }

// const myEosNet = { // scatter中提供的eos主网配置
//   blockchain: 'eos',
//   protocol: 'https',
//   host: 'nodes.get-scatter.com',
//   port: 443,
//   chainId: 'aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906'
// }

// 登录
const Login = async function () { // 目前只使用eos, 故可以使用死数据；以后需要多条链时，在改为通用接口
  // 连接Scatter
  let connected = await ScatterJS.scatter.connect('BDEX_Exchange') // exchange作为scatter中的证书名
  if (!connected) {
    window.alert('请安装或者打开scatter')
    return false
  }
  let scatter = ScatterJS.scatter
  const requiredFields = {accounts: [myEosNet]}

  try {
    await scatter.getIdentity(requiredFields)
    if (!scatter.identity) { // 什么情况下identity会为空？不记得了, 可能不会出现这种情况吧！先不管了；保留这个判断
      scatter.forgetIdentity()
      return false // 临时使用============
    }
    // console.log('登录账号==============信息：', JSON.stringify(scatter.identity))
    // let currentAccount = scatter.identity.accounts.find(x => x.blockchain === 'eos') // 暂时保留，后面若是不需要则删除
    // console.log('登录的当前账号信息:', JSON.stringify(currentAccount))
  } catch (error) {
    return false
  }
  // 数据状态处理
  store.commit('account/setMyEos', scatter)
  window.ScatterJS = null // ScatterJS是全局属性，前端可通过window.ScatterJS直接拿到，安全需要必须置空
  return scatter // 不返回true, 登录成功返回scatter; 不要再次从vuex中取
}

// 登出eos
const Logout = async function () { // 目前只做eos,故，netNameArr = ['eosTopblock']
  const scatter = store.state.account.myEos
  if (!scatter || !scatter.identity) {
    // 特殊情况，退出使用
    // 连接Scatter
    let connected = await ScatterJS.scatter.connect('BDEX_Exchange') // exchange作为scatter中的证书名
    if (!connected) {
      return true // 不需要退出，也可以算是退出成功
    }
    let scatter = ScatterJS.scatter
    // 数据状态处理
    try {
      await scatter.forgetIdentity() // 退出登录
    } catch (err) {
      return false
    }
    window.ScatterJS = null // ScatterJS是全局属性，前端可通过window.ScatterJS直接拿到，安全需要必须置空
  } else {
    // 正常退出使用
    // 退出登录
    store.dispatch('account/deleteMyEos')
    return true
  }
}

// 基于eosio的区块链
// 转账签名, 以交易所账号为参考,合约账号tbcexchange1
// contractAccount是被操作的token的合约
// myContract是接收的账号, 暂时先保留，后面可以写死
// quantity 数量 精度四位小数
// memo 备注  作为记录和传递数据使用
const getEosTransferSign = async function (contractAccount, myContract, quantity, memo) { // myContract: 'tbcexchange1'
  const scatter = store.state.account.myEos
  if (!scatter || !scatter.identity) { // 这个后期可以去掉，毕竟只有登录后才会执行这个方法；登录是前提
    window.alert('请安装或者打开scatter')
    return false
  }
  const currentAccount = scatter.identity.accounts[0]
  // 获取eos API对象
  const config = {expireInSeconds: 60} // eos超时配置
  const eos = scatter.eos(myEosNet, Eos, config)
  // 交易配置
  const transactionOptions = {
    authorization: [`${currentAccount.name}@${currentAccount.authority}`],
    broadcast: false,
    sign: true
  }
  try {
    let contract = await eos.contract(contractAccount) // 转eos时，使用'eosio.token'
    let trxSign = await contract.transfer(currentAccount.name, myContract, quantity, memo, transactionOptions)
    delete trxSign.transaction.compression
    return trxSign.transaction
  } catch (error) {
    return null
  }
}

// 撤单签名
// myContract交易所的合约
const getEosWithdrawSign = async function (myContract, orderId) { // myContract: 'tbcexchange1'
  const scatter = store.state.account.myEos
  if (!scatter || !scatter.identity) { // 这个后期可以去掉，毕竟只有登录后才会执行这个方法；登录是前提
    window.alert('请安装或者打开scatter')
    return false
  }
  const currentAccount = scatter.identity.accounts[0]
  // 获取eos API对象
  const config = {expireInSeconds: 60} // eos超时配置
  const eos = scatter.eos(myEosNet, Eos, config)
  // 交易配置
  const transactionOptions = { // withdraw，新加transactionOptions，不知是否可以
    authorization: [`${currentAccount.name}@${currentAccount.authority}`],
    broadcast: false,
    sign: true
  }
  try {
    let contract = await eos.contract(myContract)
    let wdSign = await contract.withdraw(currentAccount.name, orderId, transactionOptions) // wdSign撤单签名后的数据
    delete wdSign.transaction.compression
    return wdSign.transaction
  } catch (error) {
    return null
  }
}

// 查询token余额
const getCurrencyBalance = async function (code, symbol) { // code合约地址， symbol tokenName
  const scatter = store.state.account.myEos
  if (!scatter || !scatter.identity) { // 这个后期可以去掉，毕竟只有登录后才会执行这个方法；登录是前提
    window.alert('请安装或者打开scatter')
    return false
  }
  const currentAccount = scatter.identity.accounts[0]
  // 获取eos API对象
  const config = {expireInSeconds: 60} // eos超时配置
  try {
    const eos = scatter.eos(myEosNet, Eos, config)
    const balance = await eos.getCurrencyBalance(code, currentAccount.name, symbol)
    let bal = balance[0]
    return bal.split(' ')[0]
  } catch (e) {
    return 0.0000
  }
}

export {Login, Logout, getEosTransferSign, getEosWithdrawSign, getCurrencyBalance}
