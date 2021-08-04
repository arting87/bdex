// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vant from 'vant'
// import ElementUI from 'element-ui'
// import 'element-ui/lib/theme-chalk/index.css'
import 'vant/lib/index.css'
import App from './App'
import router from './router'
import store from './store/index'
import i18n from './utils/i18n'
import filters from './utils/filters'
import methodsMixin from './utils/methodsMixin.js'
import axios from 'axios'
import VueNativeSocket from 'vue-native-websocket'

Vue.use(Vant)
// Vue.use(ElementUI)

// add global filters
Object.keys(filters).forEach(k => Vue.filter(k, filters[k]))

// add global mixins
Vue.mixin(methodsMixin)

const EventBus = new Vue()
Vue.prototype.$bus = EventBus

// VueNative 配置
Vue.use(VueNativeSocket, 'ws://47.52.192.27:8082/ws', { // 生产：wss://bdex.me/ws测试：ws://47.103.108.61:8082/ws  'ws://47.52.192.27:8082/ws'
  store: store,
  format: 'json',
  reconnection: false, // 是否自动重连
  reconnectionAttempts: 3, // 自动重连次数
  reconnectionDelay: 2000 // close后多长时间重新连接,默认为1000
}) // protocol暂时先不管，等理解在设置子协议

Vue.config.productionTip = false

Vue.prototype.$axios = axios

if (process.argv[2] && process.argv[2] === 'cordova') {
  document.addEventListener('deviceready', function () {
    new Vue({
      router,
      store,
      i18n,
      render: h => h(App)
    }).$mount('#app')
    // window.navigator.splashscreen.hide()
  }, false)
} else {
  new Vue({
    router,
    store,
    i18n,
    render: h => h(App)
  }).$mount('#app')
}
