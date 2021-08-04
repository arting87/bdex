// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import App from './App'
import router from './router'
import store from './store/index'
import i18n from './utils/i18n'
import axios from 'axios'
import VueNativeSocket from 'vue-native-websocket'

Vue.use(ElementUI)

// VueNative 配置
Vue.use(VueNativeSocket, 'ws://47.52.192.27:8082/ws', { // 'wss://bdex.me/ws' push时需要修改回来；还有index.js需要修改, test:ws://47.103.108.61:8082/ws
  store: store,
  format: 'json',
  reconnection: false, // 是否自动重连
  reconnectionAttempts: 3, // 自动重连次数
  reconnectionDelay: 2000 // close后多长时间重新连接,默认为1000
}) // protocol暂时先不管，等理解在设置子协议

Vue.config.productionTip = false

Vue.prototype.$axios = axios

new Vue({
  router,
  store,
  i18n,
  render: h => h(App)
}).$mount('#app')
