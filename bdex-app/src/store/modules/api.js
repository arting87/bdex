import Vue from 'vue'
import {orderSort, changeListSort, volListSort} from '../../utils/sort'
// 行情数据状态管理, api模块

// 以链为主干
// 必须包含contractAccount Address
const state = {
  eos: null, // 默认为空， 存储{}结构数据 ;eos主网
  eth: null, // eth 主网
  eosSide: null, // eos同构跨链   以eos作为基准参考
  crossChain: null, // 非同构跨链  以eos作为基准参考
  changeList: [],
  volList: []
}

const mutations = {
  set (state, {eosMarket, ethMarket, eosSideMarket, crossChainMarket}) {
    state.eos = eosMarket
    state.eosSide = eosSideMarket
    state.eth = ethMarket
    state.crossChain = crossChainMarket
  },
  setChangeList (state, changList) {
    state.changeList = changList
  },
  setVolList (state, volList) {
    state.volList = volList
  }
}

const actions = {
  get ({commit}, msg) { // 测试时可以使用async试试
    // todo 涨幅榜和成交量榜单
    const changeList = [] // {pair: bde/eth, change: change}
    const volList = [] // {tokenName: bde, volume: volume}
    let ethPrice = '' // eth 兑换eos的价格
    let bdePrice = '' // bde 兑换eth的价格
    // eos
    let eosMarket = {}
    let eosOrder = {}

    // eth
    let ethMarket = {}
    let ethOrder = {}
    // eosSide
    let eosSideMarket = {}
    let eosSideOrder = {}

    // crossChain
    let crossChainMarket = {}
    let crossChainOrder = {}

    const datas = msg.data

    for (let i = 0; i < datas.length; i++) {
      const data = datas[i] // 数组中的一条数据
      const market = data.marketApi
      const order = data.orderApi

      const tokenName = data.tokenName
      const pair = data.tokenName
      let lastPrice = 0
      let change = 0
      let highPrice = 0
      let lowPrice = 0
      let volume = 0
      let total = 0
      if (market.basePrice !== 0) {
        lastPrice = market.price
        change = ((lastPrice - market.basePrice) / market.basePrice * 100).toFixed(2)
        highPrice = market.highestPrice
        lowPrice = market.lowestPrice
        volume = market.volume.toFixed(4)
        total = (lastPrice * volume).toFixed(4)
      }
      const code = data.marketCode
      const contractAddress = market.contractAddress
      const buyOrder = orderSort(order.buyOrder)
      const sellOrder = orderSort(order.sellOrder)
      const completeOrder = order.completeOrder

      if (data.marketCode === '1') {
        Vue.set(eosMarket, tokenName, {pair, lastPrice, change, highPrice, lowPrice, volume, total, code, contractAddress})
        Vue.set(eosOrder, tokenName, {buyOrder, sellOrder, completeOrder})
        // todo 涨幅榜
        const changeListPair = tokenName + '/EOS'
        changeList.push({pair: changeListPair, change: parseFloat(change)})
        // todo 成交量榜
        volList.push({tokenName: tokenName, vol: parseFloat(total), code: code})
      } else if (data.marketCode === '2') {
        Vue.set(ethMarket, tokenName, {pair, lastPrice, change, highPrice, lowPrice, volume, total, code, contractAddress})
        Vue.set(ethOrder, tokenName, {buyOrder, sellOrder, completeOrder})
        // todo 涨幅榜
        const changeListPair = tokenName + '/ETH'
        changeList.push({pair: changeListPair, change: parseFloat(change)})
        if (tokenName === 'BDE') {
          bdePrice = lastPrice
        }
        // todo 成交量榜
        volList.push({tokenName: tokenName, vol: parseFloat(total), code: code})
      } else if (data.marketCode === '3') {
        Vue.set(eosSideMarket, tokenName, {pair, lastPrice, change, highPrice, lowPrice, volume, total, code, contractAddress})
        Vue.set(eosSideOrder, tokenName, {buyOrder, sellOrder, completeOrder})
        // todo 涨幅榜
        const changeListPair = tokenName + '/BDE'
        changeList.push({pair: changeListPair, change: parseFloat(change)})
        // todo 成交量榜
        volList.push({tokenName: tokenName, vol: parseFloat(total), code: code})
      } else {
        Vue.set(crossChainMarket, tokenName, {pair, lastPrice, change, highPrice, lowPrice, volume, total, code, contractAddress})
        Vue.set(crossChainOrder, tokenName, {buyOrder, sellOrder, completeOrder})
        // todo 涨幅榜
        const changeListPair = tokenName + '/EOS'
        changeList.push({pair: changeListPair, change: parseFloat(change)})
        // todo 成交量榜
        ethPrice = lastPrice
        volList.push({tokenName: tokenName, vol: parseFloat(total), code: code})
      }
    }
    const newChangList = changeListSort(changeList)
    commit('setChangeList', newChangList)
    const newVolList = volListSort(volList, ethPrice, bdePrice)
    commit('setVolList', newVolList)
    commit('set', {eosMarket, ethMarket, eosSideMarket, crossChainMarket}) // 提交至api/set, 即本模块
    commit('orderApi/set', {eosOrder, ethOrder, eosSideOrder, crossChainOrder}, {root: true}) // 提交至orderApi/set
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
