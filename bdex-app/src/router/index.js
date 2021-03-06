import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/view/index'
import Market from '@/view/market/Market'
import MarketIndex from '@/view/market/MarketIndex'
import Trade from '@/view/trade/Trade'
import TradeIndex from '@/view/trade/TradeIndex'
import My from '@/view/my/My'
import MyIndex from '@/view/my/MyIndex'
// import store from '../store/index'

// import { Login } from '../plugins/scatter/scatter'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/',
      redirect: '/index'
    },
    {
      path: '/index',
      name: 'index',
      component: Index,
      children: [
        {
          name: 'home',
          path: 'home',
          meta: {
            needTabbar: true
          },
          component: () => import('@/view/home/Home'),
          alias: '/'
        },
        {
          name: 'news',
          path: 'news',
          component: () => import('@/view/news/News')
        },
        {
          path: 'newsdetails',
          name: 'newsdetails',
          component: () => import('@/view/news/NewsDetails')
        }
      ]
    },
    {
      path: '/market',
      component: MarketIndex,
      children: [
        {
          path: '',
          name: 'market',
          meta: {
            needTabbar: true
          },
          component: Market
        },
        {
          path: ':tokenInfo',
          name: 'detail',
          component: () => import('@/view/market/MarketDetail')
        }
      ]
    },
    {
      path: '/trade',
      component: TradeIndex,
      children: [
        {
          path: '',
          name: 'trade',
          meta: {
            needTabbar: true
          },
          component: Trade
        },
        {
          path: 'token/:tokenInfo',
          component: Trade,
          meta: {
            needTabbar: true
          }
        },
        {
          path: 'history',
          name: 'history',
          component: () => import('@/view/trade/TradeHistory')
        }
      ]
    },
    {
      path: '/my',
      component: MyIndex,
      children: [
        {
          path: '',
          name: 'my',
          meta: {
            needTabbar: true
          },
          component: My
        },
        {
          path: 'withdrawETH',
          name: 'withdrawETH',
          component: () => import('@/view/wallet/WithdrawETH')
        },
        {
          path: 'depositETH',
          name: 'depositETH',
          component: () => import('@/view/wallet/DepositETH')
        },
        {
          path: 'withdrawEOS',
          name: 'withdrawEOS',
          component: () => import('@/view/wallet/WithdrawEOS')
        },
        {
          path: 'depositEOS',
          name: 'depositEOS',
          component: () => import('@/view/wallet/DepositEOS')
        },
        {
          path: 'manageResource',
          name: 'manageResource',
          component: () => import('@/view/my/ManageResource')
        },
        {
          path: 'manageRam',
          name: 'manageRam',
          component: () => import('@/view/my/ManageRam')
        },
        {
          path: 'importETHWallet',
          name: 'importETHWallet',
          component: () => import('@/view/wallet/ImportETHWallet')
        },
        {
          path: 'importEOSWallet',
          name: 'importEOSWallet',
          component: () => import('@/view/wallet/ImportEOSWallet')
        },
        {
          path: 'createETHWallet',
          name: 'createETHWallet',
          component: () => import('@/view/wallet/CreateETHWallet')
        },
        {
          path: 'createEOSWallet',
          name: 'createEOSWallet',
          component: () => import('@/view/wallet/CreateEOSWallet')
        },
        {
          path: 'importPriETH',
          name: 'importPriETH',
          component: () => import('@/view/wallet/ImportPriETH')
        },
        {
          path: 'importPriEOS',
          name: 'importPriEOS',
          component: () => import('@/view/wallet/ImportPriEOS')
        }
      ]
    }
  ]
})

// ??????????????????
router.beforeEach((to, from, next) => {
  // ?????? vuex?????????????????????
  next()
})

// ??????????????????
router.afterEach(async (to, from) => {
  // eos ??????????????????
  // if (sessionStorage.getItem('myEos') === 'myEos' && store.state.account.myEos === null) {
  //   await Login() // ????????????; ????????????????????????????????????===????????? await ??????????????????????????????????????????
  // }
})

export default router
