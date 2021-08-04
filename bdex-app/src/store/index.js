// Vuex状态管理
import Vue from 'vue'
import Vuex from 'vuex'
import createLogger from 'vuex/dist/logger'
import account from './modules/account'
import api from './modules/api'
import user from './modules/user'
import orderApi from './modules/orderApi'
import i18n from './modules/i18n'
import PONG from './modules/heartConn'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
  modules: {
    account,
    api,
    orderApi,
    user,
    PONG,
    i18n
  },
  strict: debug,
  plugins: debug ? [createLogger()] : [],

  // 非模块的vuex状态部分
  // onerror问题需要处理,即发生网络故障时，重点...!!发生网络故障前后端是否都会触发onclose事件????
  // 是否要使用心跳机制。。。!!!
  // 关闭网页时后端是否能触发onclose事件;测试后，确定会触发;同时后端关闭服务器时，前端会触发onclose事件
  // 前端关闭连接时，会触发前端onclose事件，也会触发后端onclose事件
  // import Vue from 'vue'
  state: {
    isConnected: false,
    message: '',
    reconnectError: false
  },

  mutations: {
    SOCKET_ONOPEN (state, event) {
      // Vue.prototype.$socket = event.currentTarget
      state.isConnected = true
      // 每次连接成功，都需要获取api数据
      Vue.prototype.$socket.sendObj({
        namespace: 'api',
        action: 'get',
        data: ''
      })
      // 检查用户的登陆状态; 可能出现用户处于登录状态，但是ws连接失败，之后又连接成功以及重连接成功的情况
      let data = []
      let myEos = this.state.account.myEos
      if (myEos !== null) { // 判断eos是否在线
        let account = myEos.accounts[0]
        data.push({
          account: account,
          chainCode: '1'
        })
      }
      let myEth = this.state.account.myEth
      if (myEth !== null) { // 判断eth是否在线
        let address = this.state.account.myEth.address
        if (address.indexOf('0x') === -1) {
          address = '0x' + address
        }
        // todo 为什么不解密可以拿到地址？？？虽然少一个'0x'
        data.push({
          account: address,
          chainCode: '2'
        }) // push eth账号 !!!
      }
      if (data.length !== 0) {
        Vue.prototype.$socket.sendObj({
          namespace: 'user',
          action: 'get',
          data: data
        })
      }
      // 获取用户id
      Vue.prototype.$socket.sendObj({
        namespace: 'user',
        action: 'getid',
        data: ''
      })
      // 发送ping启动心跳机制
      Vue.prototype.$socket.sendObj({
        namespace: 'ping',
        action: 'send',
        data: ''
      })
    },
    SOCKET_ONCLOSE (state, event) {
      state.isConnected = false
    },
    SOCKET_ONERROR (state, event) { // 发生异常时调用！这里需要完善，1.网络连接失败/断网等情况发生时处理，后期测试时完善
      console.error(state, event)
    },
    SOCKET_RECONNECT (state, reconnectionCount) { // reconnectionCount重连次数, 是否需要记录后期确定
      console.info(state, reconnectionCount)
    },
    SOCKET_RECONNECT_ERROR (state, isRe) { // isRe: reconnectError = true
      state.reconnectError = isRe
    },
    SOCKET_ONMESSAGE (state, message) { // 正常情况下这里不会调用，冗余处理，防止传输的data中没有mutation和action
      state.message = message
    }
  }
})
