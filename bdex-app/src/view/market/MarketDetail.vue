<template>
  <div class="market-detail">
    <navbar :text="symbol" />
    <div class="detail-desc">
      <div class="left">
        <span class="price" :id="token.change > 0 ? 'buy-color' : 'sell-color'">{{token.lastPrice}}</span>
        <!-- <span>≈0.068 CNY</span> -->
          <span :id="token.change > 0 ? 'buy-color' : 'sell-color'">{{handleChange(token)}}</span>
      </div>
      <div class="right">
        <div>
          <span>高</span>
          <span>{{token.highPrice}}</span>
        </div>
        <div>
          <span>低</span>
          <span>{{token.lowPrice}}</span>
        </div>
        <div>
          <span>24h</span>
          <span>{{token.volume}}</span>
        </div>
      </div>
    </div>
    <trade-view ref="trade"></trade-view>
    <tab :tabList="tabList" :tabIndex.sync="tabIndex" />
    <div class="order-area" v-show="tabIndex === 0">
      <div class="title title-open">
        <div>
          <span>买盘</span>
          <span>数量({{tradeCoin}})</span>
        </div>
        <div>
          <span>价格({{baseCoin}})</span>
        </div>
        <div>
          <span>数量({{tradeCoin}})</span>
          <span>卖盘</span>
        </div>
      </div>
      <div class="divide-1"></div>
      <div class="order-open">
        <div class="buy-area">
          <div class="open-order-item" v-for="(item, i) in order.buy" :key="i">
            <span>{{i+1}}</span>
            <span>{{item.amt}}</span>
            <span id="buy-color">{{item.price | fixed}}</span>
          </div>
        </div>
        <div class="sell-area">
          <div class="open-order-item" v-for="(item, i) in order.sell" :key="i">
            <span id="sell-color">{{item.price | fixed}}</span>
            <span>{{item.amt}}</span>
            <span>{{i+1}}</span>
          </div>
        </div>
      </div>
    </div>
    <div class="order-area" v-show="tabIndex === 1">
      <div class="title title-finished">
        <span>时间</span>
        <span>方向</span>
        <span>价格({{baseCoin}})</span>
        <span>数量({{tradeCoin}})</span>
      </div>
      <div class="divide-1"></div>
      <div class="order-finished">
        <div class="finished-order-item" v-for="(item, i) in completeOrder" :key="i">
          <span>{{ item.date | time(true) }}</span>
          <span id="buy-color">买入</span>
          <span>{{ item.price }}</span>
          <span>{{ Math.floor((item.amt / item.price) * 10000) / 10000 }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from '../../components/Navbar'
import TradeView from '../../components/TradeView'

export default {
  name: 'market-detail',
  data () {
    return {
      tabList: ['委托', '成交'],
      tabIndex: 0,
      routeParam: this.$route.path === '/market' ? (this.$store.state.api.crossChain ? Object.keys(this.$store.state.api.crossChain)[0] + '-4' : '') : this.$route.params.tokenInfo
    }
  },
  computed: {
    lang () {
      let lang = this.$store.state.i18n.lang
      if (lang !== 'zh' && lang !== 'en') {
        lang = 'zh'
      }
      return lang
    },
    token () {
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      let code = param.split('-')[1]
      if (code === '1') {
        return this.$store.state.api.eos && this.$store.state.api.eos[tokenName]
      } else if (code === '2') {
        return this.$store.state.api.eth && this.$store.state.api.eth[tokenName]
      } else if (code === '3') {
        return this.$store.state.api.eosSide && this.$store.state.api.eosSide[tokenName]
      } else {
        return this.$store.state.api.crossChain && this.$store.state.api.crossChain[tokenName]
      }
    },
    baseCoin () {
      let param = this.routeParam
      let code = param.split('-')[1]
      if (code === '2') {
        return 'ETH'
      } else if (code === '3') {
        return 'BDE'
      } else {
        return 'EOS'
      }
    },
    tradeCoin () {
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      return tokenName
    },
    symbol () {
      return this.tradeCoin + '/' + this.baseCoin
    },
    completeOrder: function () {
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      let code = param.split('-')[1]
      // 默认completeOrder的排序是正序，即数组0号元素是最新的成交订单
      if (code === '1') {
        return this.$store.state.orderApi.eos && this.$store.state.orderApi.eos[tokenName]['completeOrder']
      } else if (code === '2') {
        return this.$store.state.orderApi.eth && this.$store.state.orderApi.eth[tokenName]['completeOrder']
      } else if (code === '3') {
        return this.$store.state.orderApi.eosSide && this.$store.state.orderApi.eosSide[tokenName]['completeOrder']
      } else {
        return this.$store.state.orderApi.crossChain && this.$store.state.orderApi.crossChain[tokenName]['completeOrder']
      }
    },
    order: function () {
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      let code = param.split('-')[1]
      // 默认completeOrder的排序是正序，即数组0号元素是最新的成交订单
      if (code === '1') {
        const buy = this.$store.state.orderApi.eos && this.$store.state.orderApi.eos[tokenName]['buyOrder']
        const sell = this.$store.state.orderApi.eos && this.$store.state.orderApi.eos[tokenName]['sellOrder']
        const lastPrice = this.$store.state.api.eos && this.$store.state.api.eos[tokenName]['lastPrice']
        return {
          buy,
          sell,
          lastPrice
        }
      } else if (code === '2') {
        const buy = this.$store.state.orderApi.eth && this.$store.state.orderApi.eth[tokenName]['buyOrder']
        const sell = this.$store.state.orderApi.eth && this.$store.state.orderApi.eth[tokenName]['sellOrder']
        const lastPrice = this.$store.state.api.eth && this.$store.state.api.eth[tokenName]['lastPrice']
        return {
          buy,
          sell,
          lastPrice
        }
      } else if (code === '3') {
        const buy = this.$store.state.orderApi.eosSide && this.$store.state.orderApi.eosSide[tokenName]['buyOrder']
        const sell = this.$store.state.orderApi.eosSide && this.$store.state.orderApi.eosSide[tokenName]['sellOrder']
        const lastPrice = this.$store.state.api.eosSide && this.$store.state.api.eosSide[tokenName]['lastPrice']
        return {
          buy,
          sell,
          lastPrice
        }
      } else {
        const buy = this.$store.state.orderApi.crossChain && this.$store.state.orderApi.crossChain[tokenName]['buyOrder']
        const sell = this.$store.state.orderApi.crossChain && this.$store.state.orderApi.crossChain[tokenName]['sellOrder']
        const lastPrice = this.$store.state.api.crossChain && this.$store.state.api.crossChain[tokenName]['lastPrice']
        return {
          buy,
          sell,
          lastPrice
        }
      }
    }
  },
  methods: {
    handleChange (token) { // 更新数据是否会重新渲染？待测试
      if (token === null) {
        return '--'
      } else {
        let change = token.change
        if (change >= 0) {
          return '+' + change + '%'
        } else {
          return change + '%'
        }
      }
    }
  },
  components: {
    Navbar,
    TradeView
  },
  mounted () {
    let param = 'ETH-4'
    let tokenName = param.split('-')[0]
    let code = param.split('-')[1]
    code = parseInt(code)
    let lang = localStorage.lang
    if (lang !== 'zh' && lang !== 'en') {
      lang = 'zh'
    }
    this.$refs.trade.init(tokenName, 5, code, lang, false) // tokenName, interval 区间, marketCode, lang 语言, changeLang改变语言

    console.log(this.order)
  }
}
</script>

<style lang='scss' scoped>
.market-detail {
  display: flex;
  flex-direction: column;
  align-items: center;
  span {
    font-size: 1rem;
    color: $textColor;
  }
  .detail-desc {
    height: 8rem;
    width: 33.4rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    .left, .right {
      display: flex;
      flex-direction: column;
      justify-content: space-between;
    }
    .left {
      height: 4.4rem;
      .price {
        font-size: 1.6rem;
      }
    }
    .right {
      height: 5rem;
      div {
        display: flex;
        align-items: center;
        :first-child {
          opacity: .5;
          margin-right: 1rem;
        }
      }
    }
  }
  .order-area {
    display: flex;
    flex-direction: column;
    justify-content: center;
    .title {
      display: flex;
      width: 33.5rem;
      justify-content: space-between;
      height: 4.5rem;
      align-items: center;
      span {
        font-size: 1.2rem;
        color: rgba(255, 255, 255, .5);
      }
    }
    .title-finished {
      span {
        flex: 1;
      }
    }
    .order-finished, .order-open {
      display: flex;
      width: 33.5rem;
    }
    .order-finished {
      flex-direction: column;
      .finished-order-item {
        display: flex;
        justify-content: space-between;
        height: 2.8rem;
        align-items: center;
        span {
          font-size: 1.2rem;
          color: $textColor;
          flex: 1;
        }
      }
    }
    .order-open {
      .buy-area, .sell-area {
        flex: 1;
      }
      .open-order-item {
        display: flex;
        align-items: center;
        justify-content: space-between;
        color: $textColor;
        height: 2.8rem;
        span {
          flex: 1;
          font-size: 1.2rem;
        }
      }
      .sell-area {
        .open-order-item {
          :nth-child(2), :nth-child(3) {
            text-align: right;
          }
        }
      }
    }
  }
}
</style>
