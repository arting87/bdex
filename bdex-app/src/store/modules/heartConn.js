// websocket 心跳机制处理
import Vue from 'vue'
const state = {
  timeoutId: null
}

const mutations = {
  setTimeoutId (state) {
    state.timeoutId = setTimeout(function () {
      Vue.prototype.$socket.sendObj({
        namespace: 'ping',
        action: 'send',
        data: ''
      })
    }, 50000)
  },
  clearTimeoutId (state) {
    if (state.timeoutId !== null) {
      clearTimeout(state.timeoutId)
    }
  }
}

const actions = {
  PONG ({commit}, msg) {
    commit('clearTimeoutId')
    commit('setTimeoutId')
  }
}

export default {
  namespaced: true, // heart
  state,
  mutations,
  actions
}
