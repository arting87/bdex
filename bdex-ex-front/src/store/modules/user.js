import {numFloat} from '../../utils/handleFloat'
// 用户的订单处理,已完成订单即历史委托,当前委托，下单处理
// state存储当前委托，考虑到当前委托的数量通常比较少。。。使用tokenInfo作为key,即'tokenName-marketCode'
const state = {
  currentOrder: null, // 对象+数组， key: tokenInfo value:[]
  historyOrder: null,
  idArr: null, // 存储数组
  eosOrder: null, // {oderId, status}
  ethOrder: null,
  eosWithdraw: null,
  ethWithdraw: null,
  bdeOrder: null, // todo 8.13新增
  bdeWithdraw: null
}

const mutations = {
  setCurrentOrder (state, currentOrder) {
    state.currentOrder = currentOrder
  },
  setHistoryOrder (state, historyOrder) {
    state.historyOrder = historyOrder
  },
  deleteOrder (state, val) { // 用户退出登录，清空用户委托单数据
    state.currentOrder = val
    state.historyOrder = val
  },
  setIdArr (state, idArr) {
    state.idArr = idArr
  },
  delId (state, index) { // index：0；为固定值，每次取第一个元素使用
    let oldArr = state.idArr
    let newArr = []
    if (index === 0) {
      for (let i = 1; i < oldArr.length; i++) {
        newArr.push(oldArr[i])
      }
      state.idArr = newArr
    } else {
    }
  },
  setEosOrder (state, eosOrder) {
    state.eosOrder = eosOrder
  },
  setEthOrder (state, ethOrder) {
    state.ethOrder = ethOrder
  },
  setEosWithdraw (state, eosWithdraw) {
    state.eosWithdraw = eosWithdraw
  },
  setEthWithdraw (state, ethWithdraw) {
    state.ethWithdraw = ethWithdraw
  },
  // todo 8.13 新增
  setBdeOrder (state, bdeOrder) {
    state.bdeOrder = bdeOrder
  },
  setBdeWithdraw (state, bdeWithdraw) {
    state.bdeWithdraw = bdeWithdraw
  }
}

const actions = {
  // todo 8.13 新需求，增加
  withdrawbde ({commit}, msg) {
    console.log('bde撤单获取的数据-----------', msg.data)
    let orderId = msg.data.orderId
    let status = msg.data.status
    let bdeWithdraw = {
      orderId,
      status
    }
    commit('setBdeWithdraw', bdeWithdraw)
  },
  sendbde ({commit}, msg) {
    console.log('bde下单返回获取的数据-----------', msg.data)
    let orderId = msg.data.orderId
    let status = msg.data.status
    let bdeOrder = {
      orderId,
      status
    }
    commit('setBdeOrder', bdeOrder)
  },
  withdraweos ({commit}, msg) {
    let orderId = msg.data.orderId
    let status = msg.data.status
    let eosWithdraw = {
      orderId,
      status
    }
    commit('setEosWithdraw', eosWithdraw)
  },
  withdraweth ({commit}, msg) {
    let orderId = msg.data.orderId
    let status = msg.data.status
    let ethWithdraw = {
      orderId,
      status
    }
    commit('setEthWithdraw', ethWithdraw)
  },
  sendeos ({commit}, msg) {
    let orderId = msg.data.orderId
    let status = msg.data.status
    let eosOrder = {
      orderId,
      status
    }
    commit('setEosOrder', eosOrder)
  },
  sendeth ({commit}, msg) {
    let orderId = msg.data.orderId
    let status = msg.data.status
    let ethOrder = {
      orderId,
      status
    }
    commit('setEthOrder', ethOrder)
  },
  getid ({commit}, msg) {
    let idArr = msg.data
    commit('setIdArr', idArr)
  },
  get ({commit}, msg) {
    let chainsData = msg.data
    let currentOrder = {}
    let historyOrder = {}

    for (let k = 0; k < chainsData.length; k++) {
      let data = chainsData[k].datas
      for (let i = 0; i < data.length; i++) {
        let order = data[i] // 一个token的订单
        const cArr = order.currentOrder // cArr: currentArr
        const hArr = order.historyOrder // hArr: historyArr

        let ncArr = []
        let nhArr = []
        const tokenInfo = order.tokenName + '-' + order.marketCode
        // currentOrder
        for (let j = 0; j < cArr.length; j++) {
          let cObj = {}
          cObj.orderId = cArr[j].orderId
          cObj.date = cArr[j].date
          cObj.pair = order.tokenName
          cObj.direction = parseInt(cArr[j].type, 10) // 0:买  1:卖 注意:string类型
          cObj.orderPrice = numFloat(cArr[j].orderPrice, 6)
          cObj.orderVol = numFloat(cArr[j].orderVol, 4)
          cObj.dealVol = numFloat(cArr[j].dealtVol, 4)
          cObj.pendingVol = numFloat(cArr[j].pendingVol, 4)
          cObj.status = cArr[j].status
          cObj.code = order.marketCode
          if (currentOrder.hasOwnProperty(tokenInfo)) {
            currentOrder[tokenInfo].push(cObj)
          } else {
            ncArr.push(cObj)
          }
        }
        // historyOrder
        for (let j = 0; j < hArr.length; j++) {
          let hObj = {}
          hObj.orderId = hArr[j].orderId
          hObj.date = hArr[j].date
          hObj.pair = order.tokenName
          hObj.direction = parseInt(hArr[j].type, 10) // 0:买  1:卖 注意：string类型
          hObj.orderPrice = numFloat(hArr[j].orderPrice, 6)
          hObj.orderVol = numFloat(hArr[j].orderVol, 4)
          hObj.dealVol = numFloat(hArr[j].dealtVol, 4)
          hObj.pendingVol = numFloat(hArr[j].pendingVol, 4)
          hObj.status = hArr[j].status // 1：完成    2：未完成   3：撤销 4: 下单失败 5：下单成功
          hObj.code = order.marketCode
          if (historyOrder.hasOwnProperty(tokenInfo)) {
            historyOrder[tokenInfo].push(hObj)
          } else {
            nhArr.push(hObj)
          }
        }
        if (!currentOrder.hasOwnProperty(tokenInfo)) {
          currentOrder[tokenInfo] = ncArr
          historyOrder[tokenInfo] = nhArr
        }
      }
    }
    commit('setCurrentOrder', currentOrder)
    commit('setHistoryOrder', historyOrder)
  }
}

export default {
  namespaced: true, // userOrder
  state,
  mutations,
  actions
}
