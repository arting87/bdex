<template>
  <div class="trade-history">
    <navbar/>
    <tab :tabList="tabList" :tabIndex.sync="tabIndex" />
    <div class="order-area" v-show="tabIndex === 0">
      <div  v-if="openOrderList.length > 0">
        <order-item :data="item" v-for="(item,i) in openOrderList" :key="i" />
      </div>
      <div class="no-data" v-else>暂无数据</div>
    </div>
    <div class="order-area" v-show="tabIndex === 1">
      <div  v-if="finishOrderList.length > 0">
        <order-item :data="item" v-for="(item,i) in finishOrderList" :key="i" />
      </div>
      <div class="no-data" v-else>暂无数据</div>
    </div>
  </div>
</template>

<script>
import Navbar from '../../components/Navbar'
import Tab from '../../components/Tab'
import OrderItem from '../../components/OrderItem'

export default {
  name: 'trade-history',
  data () {
    return {
      tabList: ['全部委托', '历史纪录'],
      tabIndex: 0,
      // openOrderList: [],
      // finishOrderList: [],
      routeParam: this.$route.path === '/market' ? (this.$store.state.api.crossChain ? Object.keys(this.$store.state.api.crossChain)[0] + '-4' : '') : this.$route.params.tokenInfo
    }
  },
  computed: {
    openOrderList () {
      let param = this.$route.params.tokenInfo
      const currentOrder = this.$store.state.user.currentOrder
      if (currentOrder !== null && currentOrder.hasOwnProperty(param)) {
        return currentOrder[param]
      } else {
        return []
      }
    },
    finishOrderList () {
      let param = this.$route.params.tokenInfo
      const historyOrder = this.$store.state.user.historyOrder
      if (historyOrder !== null && historyOrder.hasOwnProperty(param)) {
        return historyOrder[param]
      } else {
        return []
      }
    }
  },
  watch: {

  },
  methods: {
  },
  components: {
    Navbar,
    Tab,
    OrderItem
  },
  mounted () {
  }
}
</script>

<style lang='scss' scoped>
.trade-history {
  display: flex;
  flex-direction: column;
  align-items: center;
  .order-area {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
}
</style>
