import Vue from 'vue'
import store from '../index'
import web3 from '../../lib/eth/web3'

const state = {
  myEos: null,
  myEth: null
}

const mutations = {
  setMyEos (state, scatter) {
    state.myEos = scatter
    sessionStorage.setItem('myEos', 'myEos')
    if (store.state.isConnected === true) {
      let data = []
      let account = scatter.identity.accounts[0].name
      data.push({
        account: account,
        chainCode: '1'
      })
      if (state.myEth !== null) {
        data.push({
          account: state.myEth.address,
          chainCode: '2'
        }) // push eth账号 !!!
      }
      Vue.prototype.$socket.sendObj({
        namespace: 'user',
        action: 'get',
        data: data
      })
    }
  },
  setMyEth (state, wallet) {
    state.myEth = wallet
    if (store.state.isConnected === true) {
      let data = []
      data.push({
        account: wallet.address,
        chainCode: '2'
      }) // push eth账号 !!!
      if (state.myEos !== null) {
        let account = state.myEos.identity.accounts[0].name
        data.push({
          account: account,
          chainCode: '1'
        })
      }
      Vue.prototype.$socket.sendObj({
        namespace: 'user',
        action: 'get',
        data: data
      })
    }
  },
  delMyEos (state) {
    state.myEos.forgetIdentity()
    state.myEos = null
    sessionStorage.setItem('myEos', 'logoutEos')
  },
  delMyEth (state) { // 删除以太坊钱包
    state.myEth = null
    sessionStorage.removeItem('myEth')
  }
}

const actions = {
  // eos 处理
  async deleteMyEos ({commit}) {
    commit('delMyEos')
    // 通知后台，eos用户退出
    let data = []
    if (store.state.isConnected === true) {
      if (state.myEth !== null) { // 判断eth是否在线
        data.push({
          account: state.myEth.address,
          chainCode: '2'
        }) // push eth账号 !!!
      }
      Vue.prototype.$socket.sendObj({
        namespace: 'user',
        action: 'get',
        data: data
      })
    }
    if (data.length === 0) {
      commit('user/deleteOrder', null, {root: true}) // 在mutations里不可以提交
    }
  },
  // eth 处理
  async createEth ({commit}, opts) { // eth创建钱包
    const wallet = web3.eth.accounts.create(opts.password)
    const pri = web3.eth.accounts.encrypt(wallet.privateKey, opts.password)
    // 保存加密后的私钥
    sessionStorage.setItem('myEth', JSON.stringify(pri))
    commit('setMyEth', wallet) // 保存钱包至myEth
    return wallet
  },
  async importEth ({commit}, opts) { // eth导入钱包
    try {
      let pri
      let wallet
      const priv = opts.privateKey.toLowerCase()
      if (priv.indexOf('0x') !== -1) {
        pri = web3.eth.accounts.encrypt(priv, opts.password)
        wallet = web3.eth.accounts.wallet.add(priv)
      } else {
        pri = web3.eth.accounts.encrypt('0x' + priv, opts.password)
        wallet = web3.eth.accounts.wallet.add('0x' + priv)
      }
      // 保存加密后的私钥
      sessionStorage.setItem('myEth', JSON.stringify(pri))
      commit('setMyEth', wallet) // 保存钱包至myEth
      return wallet
    } catch (err) {
      return false
    }
  },
  async exportMyEth ({commit}, opts) {
    let sessionMyEth = sessionStorage.getItem('myEth')
    if (sessionMyEth) {
      try {
        let encryptMyEth = JSON.parse(sessionStorage.getItem('myEth'))
        const wallet = web3.eth.accounts.decrypt(encryptMyEth, opts.password)
        return wallet
      } catch (err) {
        return false
      }
    } else {
      return false
    }
  },
  async unlockMyEth ({commit}, opts) { // 解锁eth wallet
    let sessionMyEth = sessionStorage.getItem('myEth')
    if (sessionMyEth) {
      try {
        let encryptMyEth = JSON.parse(sessionStorage.getItem('myEth'))
        const wallet = web3.eth.accounts.decrypt(encryptMyEth, opts.password)
        commit('setMyEth', wallet)
        return true
      } catch (err) {
        return false
      }
    } else {
      return false
    }
  },
  async deleteMyEth ({commit}) {
    commit('delMyEth')
    // 通知后台，eth用户退出
    let data = []
    if (store.state.isConnected === true) {
      if (state.myEos !== null) { // 判断eos是否在线
        let account = state.myEos.identity.accounts[0].name
        data.push({
          account: account,
          chainCode: '1'
        })
      }
      Vue.prototype.$socket.sendObj({
        namespace: 'user',
        action: 'get',
        data: data
      })
    }
    if (data.length === 0) {
      commit('user/deleteOrder', null, {root: true})
    }
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
