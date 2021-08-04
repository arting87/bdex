<template>
  <div class="trade">
    <div class="side-token-list" v-show="showSideTokenList">
      <div class="title"><span>币币</span></div>
      <div>
        <token-list :sideMode="true" @clickItem="clickItem" />
      </div>
    </div>
    <div class="main-area" ref="main">
    <div class="top">
      <div class="menu" @click="showSide">{{token.tokenName + '/' + token.baseCoin}}</div>
      <div class="kline-more">
        <span @click="gotoDetail()"></span>
        <span @click="changeFav()">{{isFav ? '取消自选' : '添加自选'}}</span>
      </div>
    </div>
    <tab :tabList="tabList" :tabIndex.sync="tabIndex" />
    <div class="trade-area">
      <div class="left">
        <div class="trade-type">
          限价
        </div>
        <div class="price-area">
          <div class="price-action-area">
            <div @click="minus">-</div>
            <input class="price" type="number" v-model="price" />
            <div @click="add">+</div>
          </div>
          <!-- <span>≈¥0.068</span> -->
        </div>
        <div class="amount-area">
          <div class="amount-input-area">
            <input v-model="amount" type="number" placeholder="数量">
            <p>{{token.tokenName}}</p>
          </div>
          <span>可用-{{token.baseCoin}}</span>
        </div>
        <div class="slide-bar">
          <div @click="selectAmount(0)" :class="{'slide-active': this.slideBarIndex === 0}"></div>
          <div></div>
          <div @click="selectAmount(1)" :class="{'slide-active': this.slideBarIndex === 1}"></div>
          <div></div>
          <div @click="selectAmount(2)" :class="{'slide-active': this.slideBarIndex === 2}"></div>
          <div></div>
          <div @click="selectAmount(3)" :class="{'slide-active': this.slideBarIndex === 3}"></div>
          <div></div>
          <div @click="selectAmount(4)" :class="{'slide-active': this.slideBarIndex === 4}"></div>
        </div>
        <div class="total">
          <div class="token-amount">
            <span>{{availableAmount}}</span>
            <span>{{token.baseCoin}}</span>
          </div>
          <div class="total-amount">
            <span>交易额</span>
            <span>{{price * amount | fixed}}</span>
          </div>
        </div>
        <div>
          <button class="trade-btn" @click="aprroveShowClick()" v-if="isShowApproveBut">授权</button>
          <button class="trade-btn" @click="tradeOrder()">{{ accountStatus }}</button>
            <van-dialog
              v-model="aprroveShow"
              title="授权额度"
              show-cancel-button
              confirm-button-text="确认授权"
              :before-close="approveToken"
              style="text-align:center"
            >
          <van-cell-group>
            <van-field v-model="approveValue" style="text-align:center" placeholder="请输入您要授权的额度" />
          </van-cell-group>
          <div style="height: 40px;line-height: 40px;">您已授权：{{ havaAccreditAmount }}</div>
          </van-dialog>
        </div>
      </div>
      <div class="right">
        <div class="price-amount">
          <span>价格</span>
          <span>数量</span>
        </div>
        <div class="sell-area">
          <div class="buy-sell-item sell" @click="selectItem(item)" v-for="(item, i) in order.sell" :key="i">
            <span>{{item.price | fixed}}</span>
            <span>{{item.amt}}</span>
          </div>
        </div>
        <div class="last-price">
          <!-- <span>{{order.lastPrice}}</span> -->
          <!-- <span>≈0.068 CNY</span> -->
        </div>
        <div class="buy-area">
          <div class="buy-sell-item buy" @click="selectItem(item)" v-for="(item, i) in order.buy" :key="i">
            <span>{{item.price | fixed}}</span>
            <span>{{item.amt}}</span>
          </div>
        </div>
        <div class="depth-area">
          <div class="depth-selection">
            深度
          </div>
          <div class="switch">

          </div>
        </div>
      </div>
    </div>
    <div class="order-area">
      <div class="order-desc">
        <span>当前委托</span>
        <span @click="$router.push({name: 'history', params: {tokenInfo: routeParam}})">全部</span>
      </div>
      <div class="divide-1" v-if="currentOrder.length > 0"></div>
      <div class="order-list" v-if="currentOrder.length > 0">
        <order-item v-for="(item, i) in currentOrder" :data="item" :key="i" /><!-- todo 委托单数据没有注入 fuck !!! -->
      </div>
      <div class="no-data" v-else>
        暂无数据
      </div>
    </div>
    </div>
  </div>
</template>
<script>
import store from '../../store/index'
import Tab from '../../components/Tab'
import OrderItem from '../../components/OrderItem'
import TokenList from '../../components/TokenList'
import ethTrade from '../../lib/eth/ethTrade'
import eosTrade from '../../lib/eos/eosTrade'
import web3 from '../../lib/eth/web3'
// import abiContract from '../../lib/eth/eth'
// import Bignumber from 'bignumber.js'
import {numFloat, handBig} from '../../utils/handleFloat'
import {getErc20Balance, getCurrentBalance} from '../../lib/eos/helper'
import { Notify } from 'vant'

export default {
  name: 'trade',
  components: {
    Tab,
    OrderItem,
    TokenList
  },
  data () {
    return {
      approveValue: 0, // 需要授权的值
      aprroveShow: false,
      accreditAmountBuy: 0,
      accreditAmountSell: 0,
      tabIndex: 0,
      tabList: ['买入', '卖出'],
      showSideTokenList: false,
      amount: '',
      price: '',
      isFav: false,
      availableAmount: 0,
      slideBarIndex: 0,
      codeList: ['', 'eos', 'eth', 'eosSide', 'crossChain'],
      routeParam: this.$route.path === '/trade' ? (this.$store.state.api.crossChain ? Object.keys(this.$store.state.api.crossChain)[0] + '-4' : '') : this.$route.params.tokenInfo
    } // routeParam很重要，localStorage.tradeRouteParam 需要处理; 替换成vuex或者简单的审核
  },
  computed: {
    isShowApproveBut () {
      const myEth = this.$store.state.account.myEth
      const param = this.routeParam
      const code = param.split('-')[1]
      const tabIndex = this.tabIndex
      if (code === '2' && tabIndex === 1 && myEth) {
        return true
      } else if (code === '3' && myEth) {
        return true
      }
      return false
    },
    havaAccreditAmount () { // 已经授权的额度
      const param = this.routeParam
      const tokenName = param.split('-')[0]
      const code = param.split('-')[1]
      const tabIndex = this.tabIndex // 0: 买入 1: 卖出
      if (code === '2' && tabIndex === 1) {
        return this.accreditAmountSell + ` ${tokenName}`
      } else if (code === '3') {
        if (tabIndex === 0) {
          return this.accreditAmountBuy + ' BDE'
        } else if (tabIndex === 1) {
          return this.accreditAmountSell + ` ${tokenName}`
        }
      }
    },
    accountStatus () { // todo 账户状态，显示登陆或者购买
      const myEos = this.$store.state.account.myEos
      const myEth = this.$store.state.account.myEth
      const param = this.routeParam
      const code = param.split('-')[1]
      const tabIndex = this.tabIndex
      if (code === '1') {
        if (tabIndex === 0) {
          if (myEos) {
            return '买入'
          } else {
            return '登录'
          }
        } else {
          if (myEos) {
            return '卖出'
          } else {
            return '登录'
          }
        }
      } else if (code === '2' || code === '3') {
        if (tabIndex === 0) {
          if (myEth) {
            return '买入'
          } else {
            return '登录'
          }
        } else {
          if (myEth) {
            return '卖出'
          } else {
            return '登录'
          }
        }
      } else if (code === '4') {
        if (tabIndex === 0) {
          if (myEth && myEos) {
            return '买入'
          } else {
            return '登录'
          }
        } else {
          if (myEth && myEos) {
            return '卖出'
          } else {
            return '登录'
          }
        }
      }
    },
    token: function () {
      const param = this.routeParam
      const tokenName = param.split('-')[0]
      const code = param.split('-')[1]
      // todo 8.15增加bde交易区,改动
      let baseCoin = '--'
      if (code === '1') {
        baseCoin = 'EOS'
      } else if (code === '2') {
        baseCoin = 'ETH'
      } else if (code === '3') {
        baseCoin = 'BDE'
      } else if (code === '4') {
        baseCoin = 'EOS'
      }
      // const baseCoin = code === '2' ? 'ETH' : 'EOS'
      return {
        tokenName,
        code,
        baseCoin
      }
    },
    order: function () {
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      let code = param.split('-')[1]
      code = parseInt(code)
      const buy = this.$store.state.orderApi[this.codeList[code]] && this.$store.state.orderApi[this.codeList[code]][tokenName]['buyOrder'].slice(0, 5)
      const sell = this.$store.state.orderApi[this.codeList[code]] && this.$store.state.orderApi[this.codeList[code]][tokenName]['sellOrder'].slice(0, 5)
      const lastPrice = this.$store.state.api[this.codeList[code]] && this.$store.state.api[this.codeList[code]][tokenName]['lastPrice']
      return { buy, sell, lastPrice }
    },
    currentOrder () {
      let param = this.routeParam
      const currentOrder = this.$store.state.user.currentOrder
      if (currentOrder !== null && currentOrder.hasOwnProperty(param)) {
        return currentOrder[param]
      } else {
        return []
      }
    }
  },
  methods: {
    async aprroveShowClick () { // 点击授权操作，出现弹窗
      // 查寻已授权额度
      // todo 路由变化更新授权余额
      const routeParam = this.routeParam
      const tokenName = routeParam.split('-')[0]
      const code = routeParam.split('-')[1]
      if (code === '2') {
        const haveAccreditAmountSell = await ethTrade.ethApproveAmount(tokenName)
        this.accreditAmountSell = haveAccreditAmountSell
        this.aprroveShow = true
      } else if (code === '3') {
        if (this.tabIndex === 0) {
          const haveAccreditAmountBuy = await ethTrade.ethApproveAmount('BDE')
          this.accreditAmountBuy = haveAccreditAmountBuy
          this.aprroveShow = true
        } else {
          const haveAccreditAmountSell = await ethTrade.ethApproveAmount(tokenName)
          this.accreditAmountSell = haveAccreditAmountSell
          this.aprroveShow = true
        }
      }
    },
    async approveToken (action, done) { // 授权操作 todo 测试通过
      if (action === 'confirm') {
        let approveAmount = this.approveValue
        approveAmount = numFloat(approveAmount, 4)
        if (approveAmount > 0) {
          const param = this.routeParam
          let tokenName = param.split('-')[0]
          const code = param.split('-')[1]
          const tabIndex = this.tabIndex // 0: 买入 1: 卖出
          if (code === '3' && tabIndex === 0) {
            tokenName = 'BDE'
          }
          const erc20Decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
          if (erc20Decimal < 4) {
            approveAmount = numFloat(approveAmount, erc20Decimal)
          }
          const result = await ethTrade.approveToken(tokenName, approveAmount)
          if (result === 'Processing') {
            Notify('正在授权，等待交易确认，请稍后查看')
          }
          done()
        }
      } else {
        done()
      }
    },
    async tradeOrder () { // 交易下单
      if (this.accountStatus === '登录') {
        this.$router.push({
          path: '/my'
        })
      } else {
        const param = this.routeParam
        const tokenName = param.split('-')[0]
        const code = param.split('-')[1]
        const tabIndex = this.tabIndex // 0: 买入 1: 卖出
        let price = this.price
        let amount = this.amount
        if (code === '1') { // eos下单处理, 数据审核过滤，price 6位,amount 4位, total 4位
          const eosAddress = this.$store.state.account.myEos.address
          if (tabIndex === 0) { // 买单
            if (price > 0 && amount > 0) {
              const eosBalance = getErc20Balance('eosio.token', eosAddress, 'EOS')
              // price和amount过滤
              const priArr = price.toString().split('.')
              const amtArr = amount.toString().split('.')
              const priLen = priArr.length
              const amtLen = amtArr.length
              if (priLen === 2) {
                if (priArr[1].length > 6) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              price = numFloat(price, 6)
              if (amtLen === 2) {
                if (amtArr[1].length > 4) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              amount = numFloat(amount, 4)
              let total = handBig.times(price, amount)
              total = numFloat(total, 4)
              const eosDecimal = 4
              if ((total * (10 ** eosDecimal)) < 2) {
                Notify('交易额过低,无法交易')
                return
              }
              if (eosBalance === 'error' || total > parseFloat(eosBalance)) {
                Notify('余额不足，请重新输入')
                return
              }
              const result = await eosTrade.buy(price, total, tokenName)
              if (result === 'Processing') {
                Notify('您的订单正在处理中...')
              } else {
                Notify('下单失败---------')
              }
            }
          } else { // 卖单
            if (price > 0 && amount > 0) {
              const erc20Address = this.$store.state.api.eos[tokenName].contractAddress.address
              const erc20Decimal = this.$store.state.api.eos[tokenName].contractAddress.decimal
              const erc20Balance = getErc20Balance(erc20Address, eosAddress, tokenName)
              // price和amount过滤
              // price和amount过滤
              const priArr = price.toString().split('.')
              const amtArr = amount.toString().split('.')
              const priLen = priArr.length
              const amtLen = amtArr.length
              if (priLen === 2) {
                if (priArr[1].length > 6) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              price = numFloat(price, 6)
              if (amtLen === 2) {
                if (amtArr[1].length > 4) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              amount = numFloat(amount, 4)
              if (erc20Decimal < 4) {
                amount = numFloat(amount, erc20Decimal)
              }
              if ((amount * (10 ** erc20Decimal)) < 2) {
                Notify('交易额过低,无法交易')
                return
              }
              if (erc20Balance === 'error' || (amount > parseFloat(erc20Balance))) {
                Notify('余额不足，请重新输入')
                return
              }
              const result = await eosTrade.sell(price, amount, tokenName)
              if (result === 'Processing') {
                Notify('您的订单正在处理中...')
              } else {
                Notify('下单失败---------')
              }
            }
          }
        } else if (code === '2') { // eth 下单处理
          const ethAddress = this.$store.state.account.myEth.address
          if (tabIndex === 0) { // 买单
            if (price > 0 && amount > 0) {
              // price和amount过滤
              // price和amount过滤
              const priArr = price.toString().split('.')
              const amtArr = amount.toString().split('.')
              const priLen = priArr.length
              const amtLen = amtArr.length
              if (priLen === 2) {
                if (priArr[1].length > 6) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              price = numFloat(price, 6)
              if (amtLen === 2) {
                if (amtArr[1].length > 4) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              amount = numFloat(amount, 4)
              let total = handBig.times(price, amount)
              total = numFloat(total, 4)
              const ethDecimal = 8
              if ((total * (10 ** ethDecimal)) < 2) {
                Notify('交易额过低,无法交易')
                return
              }
              let ethBalance = await web3.eth.getBalance(ethAddress)
              ethBalance = (web3.utils.fromWei(ethBalance, 'ether'))
              ethBalance = numFloat(ethBalance, 4)
              if (total > parseFloat(ethBalance)) {
                Notify('余额不足，请重新输入')
                return
              }
              const result = await ethTrade.buy(price, total, tokenName)
              if (result === 'Processing') {
                Notify('您的订单正在处理中...')
              } else {
                Notify('下单失败---------')
              }
            }
          } else { // 卖单
            if (price > 0 && amount > 0) {
              // const erc20Address = this.$store.state.api.eth[tokenName].contractAddress.address
              const erc20Decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
              // price和amount过滤
              const priArr = price.toString().split('.')
              const amtArr = amount.toString().split('.')
              const priLen = priArr.length
              const amtLen = amtArr.length
              if (priLen === 2) {
                if (priArr[1].length > 6) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              price = numFloat(price, 6)
              if (amtLen === 2) {
                if (amtArr[1].length > 4) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              amount = numFloat(amount, 4)
              if (erc20Decimal < 4) {
                amount = numFloat(amount, erc20Decimal)
              }
              if ((amount * (10 ** erc20Decimal)) < 2) {
                Notify('交易额过低,无法交易')
                return
              }
              this.accreditAmountSell = await ethTrade.ethApproveAmount(tokenName)
              if (amount > this.accreditAmountSell) {
                Notify('授权额度不足，下单失败')
                return
              }
              // todo 比较授权额度，而不是账户余额；授权时，比较账户余额
              // const coin = new web3.eth.Contract(abiContract, erc20Address)
              // let erc20Balance = await coin.methods.balanceOf(ethAddress).call()
              // let ret = new Bignumber(erc20Balance)
              // const de = 10 ** erc20Decimal
              // const d = new Bignumber(de)
              // erc20Balance = parseFloat(ret.dividedBy(d)).toFixed(4)
              // erc20Balance = numFloat(erc20Balance, 4)
              // if (amount > parseFloat(erc20Balance)) {
              //   Notify('余额不足，请重新输入')
              //   return
              // }
              const result = await ethTrade.sell(price, amount, tokenName)
              if (result === 'Processing') {
                Notify('您的订单正在处理中...')
              } else {
                Notify('下单失败---------')
              }
            }
          }
        } else if (code === '3') { // bde交易处理
          // const ethAddress = this.$store.state.account.myEth.address
          if (tabIndex === 0) { // 买单
            // const bdeAddress = this.$store.state.api.eth[tokenName].contractAddress.address
            const bdeDecimal = this.$store.state.api.eth['BDE'].contractAddress.decimal
            if (price > 0 && amount > 0) {
              // price和amount过滤
              const priArr = price.toString().split('.')
              const amtArr = amount.toString().split('.')
              const priLen = priArr.length
              const amtLen = amtArr.length
              if (priLen === 2) {
                if (priArr[1].length > 6) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              price = numFloat(price, 6)
              if (amtLen === 2) {
                if (amtArr[1].length > 4) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              amount = numFloat(amount, 4)
              let total = handBig.times(price, amount)
              total = numFloat(total, 4)
              if (bdeDecimal < 4) {
                total = numFloat(total, bdeDecimal)
              }
              if ((total * (10 ** bdeDecimal)) < 2) {
                Notify('交易额过低,无法交易')
                return
              }
              this.accreditAmountBuy = await ethTrade.ethApproveAmount('BDE')
              if (total > this.accreditAmountBuy) {
                Notify('授权额度不足，下单失败')
                return
              }
              // const coin = new web3.eth.Contract(abiContract, bdeAddress)
              // let bdeBalance = await coin.methods.balanceOf(ethAddress).call()
              // let ret = new Bignumber(bdeBalance)
              // const de = 10 ** bdeDecimal
              // const d = new Bignumber(de)
              // bdeBalance = parseFloat(ret.dividedBy(d)).toFixed(4)
              // bdeBalance = numFloat(bdeBalance, 4)
              // if (total > parseFloat(bdeBalance)) {
              //   Notify('余额不足，请重新输入')
              //   return
              // }
              const result = await ethTrade.bdeBuy(price, total, tokenName)
              if (result === 'Processing') {
                Notify('您的订单正在处理中...')
              } else {
                Notify('下单失败---------')
              }
            }
          } else { // 卖单
            if (price > 0 && amount > 0) {
              // const erc20Address = this.$store.state.api.eth[tokenName].contractAddress.address
              const erc20Decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
              // price和amount过滤
              const priArr = price.toString().split('.')
              const amtArr = amount.toString().split('.')
              const priLen = priArr.length
              const amtLen = amtArr.length
              if (priLen === 2) {
                if (priArr[1].length > 6) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              price = numFloat(price, 6)
              if (amtLen === 2) {
                if (amtArr[1].length > 4) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              amount = numFloat(amount, 4)
              if (erc20Decimal < 4) {
                amount = numFloat(amount, erc20Decimal)
              }
              if ((amount * (10 ** erc20Decimal)) < 2) {
                Notify('交易额过低,无法交易')
                return
              }
              this.accreditAmountSell = await ethTrade.ethApproveAmount(tokenName)
              if (amount > this.accreditAmountSell) {
                Notify('授权额度不足，下单失败')
                return
              }
              // const coin = new web3.eth.Contract(abiContract, erc20Address)
              // let erc20Balance = await coin.methods.balanceOf(ethAddress).call()
              // let ret = new Bignumber(erc20Balance)
              // const de = 10 ** erc20Decimal
              // const d = new Bignumber(de)
              // erc20Balance = parseFloat(ret.dividedBy(d)).toFixed(4)
              // erc20Balance = numFloat(erc20Balance, 4)
              // if (amount > parseFloat(erc20Balance)) {
              //   Notify('余额不足，请重新输入')
              //   return
              // }
              const result = await ethTrade.bdeSell(price, amount, tokenName)
              if (result === 'Processing') {
                Notify('您的订单正在处理中...')
              } else {
                Notify('下单失败---------')
              }
            }
          }
        } else if (code === '4') { // 跨链处理
          const eosAddress = this.$store.state.account.myEos.address
          const ethAddress = this.$store.state.account.myEth.address
          if (tabIndex === 0) { // 买单, eos buy eth
            if (price > 0 && amount > 0) {
              const eosBalance = getCurrentBalance(eosAddress)
              // price和amount过滤
              const priArr = price.toString().split('.')
              const amtArr = amount.toString().split('.')
              const priLen = priArr.length
              const amtLen = amtArr.length
              if (priLen === 2) {
                if (priArr[1].length > 6) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              price = numFloat(price, 6)
              if (amtLen === 2) {
                if (amtArr[1].length > 4) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              amount = numFloat(amount, 4)
              let total = handBig.times(price, amount)
              total = numFloat(total, 4)
              const eosDecimal = 4
              if ((total * (10 ** eosDecimal)) < 2) {
                Notify('交易额过低,无法交易')
                return
              }
              if (total > parseFloat(eosBalance)) {
                Notify('余额不足，请重新输入')
                return
              }
              const result = await eosTrade.eosBuyEth(price, total)
              if (result === 'Processing') {
                Notify('您的订单正在处理中...')
              } else {
                Notify('下单失败---------')
              }
            }
          } else { // 卖单
            if (price > 0 && amount > 0) {
              // price和amount过滤
              const priArr = price.toString().split('.')
              const amtArr = amount.toString().split('.')
              const priLen = priArr.length
              const amtLen = amtArr.length
              if (priLen === 2) {
                if (priArr[1].length > 6) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              price = numFloat(price, 6)
              if (amtLen === 2) {
                if (amtArr[1].length > 4) {
                  Notify('数据精度过低，请重新输入')
                  return
                }
              }
              amount = numFloat(amount, 4)
              const ethDecimal = 8
              if ((amount * (10 ** ethDecimal)) < 2) {
                Notify('交易额过低,无法交易')
                return
              }
              let ethBalance = await web3.eth.getBalance(ethAddress)
              ethBalance = (web3.utils.fromWei(ethBalance, 'ether'))
              ethBalance = numFloat(ethBalance, 4)
              if (amount > parseFloat(ethBalance)) {
                Notify('余额不足，请重新输入')
                return
              }
              const result = await ethTrade.ethBuyEos(price, amount)
              if (result === 'Processing') {
                Notify('您的订单正在处理中...')
              } else {
                Notify('下单失败---------')
              }
            }
          }
        }
      }
    },
    minus () {
      this.price = this.price - 1 < 0 ? 0 : this.price - 1
    },
    add () {
      this.price = parseInt(this.price) + 1
    },
    showSide () {
      this.showSideTokenList = !this.showSideTokenList
    },
    clickItem () {
      this.showSideTokenList = false
    },
    selectAmount (value) {
      if (!this.price) {
        return 0
      }
      this.slideBarIndex = value
      this.amount = (this.availableAmount / this.price * 0.25 * value).toFixed(4)
    },
    selectItem (item) {
      this.price = item.price.toFixed(4)
      this.amount = item.amt
    },
    gotoDetail () {
      this.$router.push('/market/' + 'ETH-4')
    },
    isThisFav () {
      const tokenInfo = this.token.tokenName + '-' + this.token.code
      let localFavTableData = localStorage.getItem('localFavTableData')
      if (!localFavTableData) {
        return false
      } else {
        localFavTableData = JSON.parse(localFavTableData)
        return localFavTableData.some(item => {
          console.log('xixixi', item, tokenInfo, (item === tokenInfo))
          return item === tokenInfo
        })
      }
    },
    changeFav () {
      const tokenInfo = this.token.tokenName + '-' + this.token.code
      try {
        let localFavTableData = localStorage.getItem('localFavTableData')
        if (localFavTableData) {
          localFavTableData = JSON.parse(localFavTableData)
          for (let i = 0; i < localFavTableData.length; i++) {
            if (localFavTableData[i] === tokenInfo) {
              localFavTableData.splice(i, 1)
              localStorage.setItem('localFavTableData', JSON.stringify(localFavTableData))
              this.$toast('取消成功')
              this.isFav = false
              return
            }
          }
          localFavTableData.push(tokenInfo)
          localStorage.setItem('localFavTableData', JSON.stringify(localFavTableData))
        } else {
          localFavTableData = []
          localFavTableData.push(tokenInfo)
          localStorage.setItem('localFavTableData', JSON.stringify(localFavTableData))
        }
        this.isFav = true
        this.$toast('添加成功')
      } catch (e) {
      }
    }
  },
  created () {
    document.addEventListener('click', (e) => {
      if (e.clientX > 300) {
        this.showSideTokenList = false
      }
    })
    if (sessionStorage.getItem('EncryptEth') && sessionStorage.getItem('pwdtEth')) {
      this.$store.dispatch('account/unlockMyEth')
    }
    const param = this.routeParam
    const code = param.split('-')[1]
    const self = this
    if (code === '2') {
      this.getETHBalance().then(v => {
        self.availableAmount = v
      })
    } else {
      this.getEOSBalance(this.$store.state.account.myEos.address).then(v => {
        self.availableAmount = v[0].split(' ')[0]
      })
    }
    this.isFav = this.isThisFav()
  },
  mounted () {
  },
  watch: {
    '$route' (to, from) {
      if (to.path === '/trade') { // 路由/trade复用组件时
        let crossChain = this.$store.state.api.crossChain
        if (!crossChain) {
          const id = setTimeout(() => {
            let crossChain2 = this.$store.state.api.crossChain
            if (crossChain2) {
              this.routeParam = Object.keys(crossChain2)[0] + '-4'
            } else {
              this.$message.error('亲，请检查网络并刷新页面！')
            }
            clearTimeout(id)
          }, 1500)
        } else {
          this.routeParam = Object.keys(crossChain)[0] + '-4'
        }
      } else {
        this.routeParam = to.params.tokenInfo
      }
    },
    routeParam: {
      handler: function (newVal) {
        if (newVal === '') {
          this.$message.error('亲，请检查网络并刷新页面！')
        }
      },
      immediate: true
    }
  },
  beforeRouteEnter (to, from, next) {
    if (store.state.api.eos) {
      next()
    } else {
      let id2 = ''
      let id = setInterval(() => {
        if (store.state.api.eos) {
          clearInterval(id)
          clearTimeout(id2)
          next()
        }
      }, 200)
      id2 = setTimeout(() => {
        clearInterval(id)
        clearTimeout(id2)
        next()
      }, 3000)
    }
  }
}
</script>

<style scoped lang='scss'>
/* 控制交易页面 */
.trade {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 1.2rem;
  padding-bottom: 5rem;
  .side-token-list {
    height: 100%;
    width: 28rem;
    position: absolute;
    top: 0;
    left: 0;
    display: flex;
    flex-direction: column;
    align-items: center;
    z-index: 100;
    background-color: $bg;
    box-shadow: 0.2rem 0rem 0.2rem 0rem rgba(0, 0, 0, 0.5);
    .title {
      width: 90%;
      display: flex;
      align-items: center;
      height: 6.8rem;
      font-size: 1.6rem;
      color: $textColor;
    }
  }
  .main-area {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .top {
    width: 33.5rem;
    height: 4.4rem;
    font-size: 1.4rem;
    color: $textColor;
    justify-content: space-between;
    align-items: center;
    display: flex;
    .menu::before {
      content: '';
      background: url("../../../static/images/头部-菜单@2x.png") no-repeat;
      width: 1.5rem;
      height: 1.1rem;
      display: inline-block;
      background-size: contain;
      margin-right: .2rem;
    }
    .kline-more {
      display: flex;
      align-items: center;
      :first-child::before {
        content: '';
        background: url("../../../static/images/头部-K线@2x.png") no-repeat;
        width: 1.4rem;
        height: 1.1rem;
        display: inline-block;
        background-size: contain;
      }
      :last-child {
        display: inline-block;
      }
    }
  }
  .trade-area {
    width: 33.5rem;
    height: 32.3rem;
    display: flex;
    .left {
      flex: 1;
      .trade-type {
        height: 5.4rem;
        display: flex;
        align-items: center;
        color: $whiteTextColor;
      }
      .price-area {
        display: flex;
        flex-direction: column;
        height: 7.1rem;
        span {
          font-size: 1rem;
          color: $textColor;
          margin-top: .7rem;
        }
        .price-action-area {
          width: 14.7rem;
          height: 3.4rem;
          border-radius: 0.2rem;
          border: solid 0.1rem #8ca1ff;
          display: flex;
          color: $textColor;
          div {
            display: flex;
            justify-content: center;
            align-items: center;
            width: 3.5rem;
          }
          :nth-child(2) {
            font-size: 1.2rem;
            width: 7.4rem;
            text-align: center;
            border-left: solid 0.1rem #8ca1ff;
            border-right: solid 0.1rem #8ca1ff;
          }
        }
      }
      .amount-area {
        display: flex;
        flex-direction: column;
        color: $textColor;
        height: 7.4rem;
        span {
          font-size: 1rem;
          color: $textColor;
          margin-top: .7rem;
        }
        input::placeholder {
          color: $textColor;
        }
        .amount-input-area {
          width: 14.7rem;
          height: 3.4rem;
          border-radius: 0.2rem;
          border: solid 0.1rem #8ca1ff;
          opacity: 0.5;
          font-size: 1.2rem;
          display: flex;
          justify-content: space-between;
          align-items: center;
          input {
            width: 10.2rem;
            text-indent: 1rem;
          }
          p {
            margin-right: 1rem;
          }
        }
      }
      .slide-bar {
        height: 1.4rem;
        width: 14.7rem;
        display: flex;
        align-items: center;
        justify-content: space-between;
        :nth-child(2n+1) {
          height: .7rem;
          width: .6rem;
          background-color: $textColor;
        }
        :nth-child(2n) {
          height: .1rem;
          width: 2.2rem;
          background-color: $textColor;
        }
        .slide-active {
          height: 1.4rem;
          width: 1.1rem;
          background: url("../../../static/images/滑块@2x.png") no-repeat;
          background-size: contain;
        }
      }
      .total {
        height: 6.9rem;
        span {
          font-size: 1rem;
          color: $textColor;
        }
        div {
          display: flex;
          justify-content: space-between;
          width: 14.7rem;
        }
        .token-amount {
          margin-top: .9rem;
        }
        .total-amount {
          margin-top: 1.9rem;
        }
      }
    }
    .right {
      flex: 1;
      .price-amount {
        display: flex;
        height: 5.4rem;
        color: $whiteTextColor;
        align-items: center;
        justify-content: space-between;
      }
      .buy-sell-item {
        height: 1.8rem;
        display: flex;
        justify-content: space-between;
        align-items: center;
      }
      .buy {
        color: $riseColor;
      }
      .sell {
        color: $fallColor;
      }
      .buy-area, sell-area {
        height: 9rem;
        display: flex;
        flex-direction: column;
        width: 100%;
      }
      .last-price {
        height: 4.7rem;
        display: flex;
        flex-direction: column;
        :first-child {
          margin-top: 1.15rem;
          color: $riseColor;
          font-size: 1.4rem;
        }
        :last-child {
          color: $textColor;
          font-size: 1rem;
        }
      }
      .depth-area {
        margin-top: 1.2rem;
        display: flex;
        justify-content: space-between;
        .depth-selection {
          width: 11.5rem;
          height: 3rem;
          border-radius: 0.2rem;
          border: solid 0.1rem #347ce1;
          display: flex;
          justify-content: center;
          align-items: center;
          color: $revokeTextColor;
        }
        .switch {
          height: 3rem;
          width: 3.8rem;
          border-radius: 0.2rem;
          border: solid 0.1rem #347ce1;
          background: url("../../../static/images/买入卖出切换@2x.png") no-repeat;
          background-size: 1.8rem 1.9rem;
          background-position: center;
        }
      }
    }
  }
  .order-area {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    .order-desc {
      width: 33.5rem;
      display: flex;
      height: 6rem;
      justify-content: space-between;
      align-items: center;
      color: $textColor;
      :first-child {
        color: $whiteTextColor;
        font-size: 1.4rem;
      }
      :last-child {
        height: 1.5rem;
        display: flex;
        align-items: center;
      }
      :last-child::before {
        content: '';
        display: inline-block;
        height: 1.2rem;
        width: .9rem;
        background: url("../../../static/images/全部@2x.png") no-repeat;
        background-size: cover;
        margin-right: .2rem;
      }
    }
    .order-list {
      color: #fff;
      width: 100%;
    }
  }
}
</style>
