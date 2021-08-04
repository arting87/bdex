// router
import Vue from "vue";
import VueRouter from "vue-router";
import axios from "axios";
import VueAxios from "vue-axios";
import index from "./index";
import Login from "./login/Login.vue";
import store from "./store/index";
import echarts from "echarts";
//element UI
//全局引入
// import ElementUI from 'element-ui'
// Vue.use(ElementUI)
import "element-ui/lib/theme-chalk/index.css";
//按需引入
import {
  Pagination,
  MessageBox,
  Message,
  DatePicker,
  TimePicker
} from "element-ui";
Vue.component(Message.name, Message); //使用use方式会在进入页面时弹出空提示
Vue.component(MessageBox.name, MessageBox);
Vue.use(Pagination);
Vue.use(DatePicker);
Vue.use(TimePicker);
Vue.prototype.$msgbox = MessageBox;
Vue.prototype.$confirm = MessageBox.confirm;
Vue.prototype.$message = Message;
//Echarts
Vue.prototype.$echarts = echarts;
//axios
Vue.use(VueAxios, axios);
axios.defaults.withCredentials = true; //设置为打开向后台发送cookie
Vue.prototype.$ajax = axios;
// axios.defaults.headers.post['Content-Type'] = 'text/plain'
// Vue.prototype.$axios = axios

Vue.use(VueRouter);
const routes = [
  { path: "/", name: "index", component: index },
  { path: "/login", name: "login", component: Login },
  { path: "*", redirect: { name: "index" } }
];

const router = new VueRouter({
  mode: "history",
  routes
});

const app = new Vue({
  store,
  router,
  render: h => h(index)
}).$mount("#app");
