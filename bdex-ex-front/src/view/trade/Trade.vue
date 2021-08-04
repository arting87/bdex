<template>
  <div class="trade">
    <el-container class="trade-content">

      <!-- 代币行情 -->
      <el-header height="108px" class="header">
        <TradeHeader :route-param="routeParam"></TradeHeader>
      </el-header>

      <el-container>
        <el-aside style="width: 292px;">
          <el-container>
            <div class="trade-aside">
              <!-- 交易对 -->
              <div class="tradeaside-top">
                <TradeAsideTop :route-param="routeParam"></TradeAsideTop>
              </div>
              <!-- 实时交易 -->
              <div>
                <TradeAsideBottom :route-param="routeParam"></TradeAsideBottom>
              </div>
            </div>
          </el-container>

        </el-aside>

        <el-container>
          <el-main class="trade-main">
            <!-- K线图 -->
            <TradeMainTop :route-param="routeParam"></TradeMainTop>
            <!-- 交易 -->
            <TradeMainBottom :route-param="routeParam"></TradeMainBottom>
          </el-main>
          <!-- 买卖盘口 -->
          <el-main class="trade-right"><TradeRight :route-param="routeParam"></TradeRight></el-main>
        </el-container>

      </el-container>
      <!-- 当前委托 -->
      <el-footer height="414px"><TradeFooterTop :route-param="routeParam"></TradeFooterTop></el-footer>
      <!-- 历史委托 -->
      <el-footer height="414px"><TradeFooterBottom :route-param="routeParam"></TradeFooterBottom></el-footer>
    </el-container>

  </div>
</template>
<script>
import TradeHeader from './TradeHeader'
import TradeAsideTop from './TradeAsideTop'
import TradeAsideBottom from './TradeAsideBottom'
import TradeMainTop from './TradeMainTop'
import TradeMainBottom from './TradeMainBottom'
import TradeRight from './TradeRight'
import TradeFooterTop from './TradeFooterTop'
import TradeFooterBottom from './TradeFooterBottom'
import store from '../../store/index'
export default {
  name: 'Trade',
  components: {
    TradeHeader,
    TradeAsideTop,
    TradeAsideBottom,
    TradeMainTop,
    TradeMainBottom,
    TradeRight,
    TradeFooterTop,
    TradeFooterBottom
  },
  data () {
    return {
      routeParam: this.$route.path === '/trade' ? (this.$store.state.api.crossChain ? Object.keys(this.$store.state.api.crossChain)[0] + '-4' : '') : this.$route.params.tokenInfo
    }
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

<style scoped>
/* 控制交易页面 */
.trade {
  margin-top: 130px;
  padding: 0;
  background-color: #1A1D2C;
}
/* 控制交易页面的居中 */
.trade-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0;
}
/* 交易页面header */
.header {
  margin: 0;
  height: 108px;
  /* line-height: 45px; */
  border-radius: 4px;
}
.trade-right {
  min-width: 292px;
  width: 292px;
  max-width: 292px;
  height: 866px;
  border: 1px solid #131624;
  padding: 0;
  margin-top: 10px;
}
.trade-main {
  margin: 10px 10px 0px 10px;
}
.el-footer {
  text-align: center;
  border: 1px solid #131624;
  border-radius: 4px;
  margin-top: 20px;
}
.el-header,.el-aside,.el-main,.el-footer {
  padding: 0;
}
.trade-aside {
  margin-top: 10px;
}
/* .tradeaside-top {
  margin-bottom: 10px;
} */

</style>
<style>
/* 控制el-tab-pane label的颜色 */
/* .el-tabs__item {
  color: #b0b8db;
}
.el-table {
  color: #b0b8db;
} */
.el-table >>> .el-table__body tr:hover>td{
    background-color: #131624;
}
</style>
