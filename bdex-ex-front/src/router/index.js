import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/view/index'
import store from '../store/index'

import {Login} from '../plugins/scatter/scatter'

Vue.use(Router)

const router = new Router({
  // mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index,
      children: [
        {
          name: 'home',
          path: 'home',
          component: () => import('@/view/home/Home'),
          alias: '/'
        },
        {
          name: 'trade',
          path: 'trade',
          component: () => import('@/view/trade/Trade')
        },
        {
          path: 'trade/:tokenInfo',
          component: () => import('@/view/trade/Trade')
        },
        {
          name: 'news',
          path: 'news',
          component: () => import('@/view/news/News')
        },
        {
          name: 'newsdetail',
          path: 'news/newsdetail',
          component: () => import('@/view/news/NewsDetail')
        },
        {
          path: 'ethwallet',
          name: 'ethwallet',
          component: () => import('@/view/wallet/ETHWallet')
        }
      ]
    },
    {
      path: '/menu',
      name: 'Menu',
      component: () => import('@/components/common/Menu')
    },
    {
      path: '/404',
      name: 'NotFound',
      component: () => import('@/view/common/NotFound')
    },
    {// 路由最下方为404页面
      path: '*',
      redirect: '/404'
    }
  ]
})

// 全局前置守卫
router.beforeEach((to, from, next) => {
  // 判断 vuex中的值是否还在
  next()
})

// 全局后置守卫
router.afterEach(async (to, from) => {
  // eos 账户状态检查
  if (sessionStorage.getItem('myEos') === 'myEos' && store.state.account.myEos === null) {
    await Login() // 连接登录; 这里是异步操作；执行顺序===注意； await 我觉得可以去掉！后面优化在说
  }
})

export default router
