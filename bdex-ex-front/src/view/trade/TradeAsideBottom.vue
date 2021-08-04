<template>
  <div class="tradeasidebottom">
    <el-row class="boxbg">
      <el-col :span="24">
        <div class="latestdeal">{{ $t('trade.latestDeal') }}</div>
      </el-col>
    </el-row>
    <div>
      <!--:data="completeOrder === null ? [] : completeOrder"  -->
    <el-table
      :data="completeOrder"
      style="width: 100%"
      height="376"
    >
      <el-table-column prop="time" :label="$t('trade.time')" width="100">
        <template slot-scope="scope">
          <span>{{ scope.row.date | time }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="price" :label="$t('trade.price') + '(' + baseCoin + ')' " width="90">
        <template slot-scope="scope">
          <span>{{ scope.row.price }}</span><!--todo 精度没有处理，是之前遗忘？？？待定8.13-->
        </template>
      </el-table-column>
      <el-table-column prop="amount" :label="$t('trade.amount') + '(' + tradeCoin + ')' " width="102">
        <template slot-scope="scope">
          <span>{{ handFloat(NH.divide(scope.row.amt, scope.row.price), 4) }}</span><!--todo 8.13 为什么是divide,除非amt是总的花费，如：eos数量，不是代币的数量-->
        </template>
      </el-table-column>
    </el-table>
    </div>
  </div>
</template>

<script>
import {numFloat, handBig} from '../../utils/handleFloat'
export default {
  name: 'TradeAsideBottom',
  props: { // 组件间通信
    routeParam: {
      type: String,
      required: true
    }
  },
  data () {
    return {
      handFloat: numFloat,
      NH: handBig
    }
  },
  computed: {
    baseCoin () {
      let param = this.routeParam
      let code = param.split('-')[1]
      if (code === '2') {
        return 'ETH'
      } else if (code === '3') { // todo 8.13 新需求改动
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
    }
  },
  filters: {
    time: function (txdate) {
      const d = new Date(txdate)
      const month = (d.getMonth() + 1) < 10 ? '0' + (d.getMonth() + 1) : '' + (d.getMonth() + 1)
      const day = d.getDate() < 10 ? '0' + d.getDate() : '' + d.getDate()
      const hour = d.getHours() < 10 ? '0' + d.getHours() : '' + d.getHours()
      const minutes = d.getMinutes() < 10 ? '0' + d.getMinutes() : '' + d.getMinutes()
      return month + '-' + day + ' ' + hour + ':' + minutes
    }
  }
}
</script>

<style scoped>
  @import '../../assets/css/TradeAsideBottom.css';
</style>
