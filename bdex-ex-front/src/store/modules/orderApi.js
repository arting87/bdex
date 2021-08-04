// 订单行情数据，也即交易对行情数据，orderApi模块

const state = {
  eos: null,
  eth: null,
  eosSide: null,
  crossChain: null, // 8.13 todo 新增eth erc20代币交易区;占用此字段
  price: 0.000000
}

const mutations = {
  set (state, {eosOrder, ethOrder, eosSideOrder, crossChainOrder}) {
    state.eos = eosOrder
    state.eth = ethOrder
    state.eosSide = eosSideOrder
    state.crossChain = crossChainOrder
  },
  setPrice (state, price) {
    state.price = price
  }
}

const actions = {

}

export default {
  namespaced: true, // orderApi
  state,
  mutations,
  actions
}
