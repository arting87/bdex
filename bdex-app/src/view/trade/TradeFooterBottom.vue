<template>
  <div class="tradefooterbottom">
    <el-tabs type="border-card">
      <el-tab-pane :label="$t('trade.orderHistory')">
        <el-table
          :data="historyOrder.filter(data => !search || data.pair.toLowerCase().includes(search.toLowerCase()))"
          style="width: 100%"
          height="361"
        >
          <!--data:显示的数据
          :default-sort="{prop: 'pair', order: 'ascending'}" 默认排序--><!--todo 8.13新增 bde交易区-->
          <el-table-column prop="pair" :label="$t('trade.pair')" width="100">
            <template slot-scope="scope">
              <span>{{ scope.row.pair }} {{ scope.row.code === '3' ? '/BDE' : (scope.row.code === '2' ? '/ ETH' : '/ EOS') }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="date" :label="$t('trade.time')" width="147">
            <template slot-scope="scope">
              <span>{{ scope.row.date | time }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="direction" :label="$t('trade.direction')" width="100">
            <template slot-scope="scope">
              <span>{{ scope.row.direction === 0 ? $t('trade.buy') : $t('trade.sell') }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="orderprice" :label="$t('trade.orderPrice')" width="150">
            <template slot-scope="scope">
              <span>{{ scope.row.orderPrice }} {{ scope.row.code === '3' ? 'BDE' : (scope.row.code === '2' ? ' ETH' : ' EOS') }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="ordervol" :label="$t('trade.orderVol')" width="150">
            <template slot-scope="scope">
              <span>{{ scope.row.orderVol.toFixed(4) }} {{ scope.row.pair }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="ordertotal" :label="$t('trade.orderTotal')" width="150">
            <template slot-scope="scope">
              <span>{{(scope.row.orderPrice * scope.row.orderVol).toFixed(4) }} {{ scope.row.code === '3' ? 'BDE' : (scope.row.code === '2' ? ' ETH' : ' EOS') }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="dealtvol" :label="$t('trade.dealtVol')" width="150">
            <template slot-scope="scope">
              <span>{{ scope.row.dealVol.toFixed(4) }} {{ scope.row.pair }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="pendingVol" :label="$t('trade.pendingVol')" width="150">
            <template slot-scope="scope">
              <span>{{ scope.row.pendingVol.toFixed(4) }} {{ scope.row.pair }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="status" :label="$t('trade.status')" width="100">
            <template slot-scope="scope">
              <span>{{ status(scope) }}</span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
export default {
  name: 'TradeFooterBottom',
  props: { // 组件间通信
    routeParam: {
      type: String,
      required: true
    }
  },
  data () {
    return {
      search: ''
    }
  },
  computed: {
    historyOrder () {
      let param = this.routeParam
      const historyOrder = this.$store.state.user.historyOrder
      if (historyOrder !== null && historyOrder.hasOwnProperty(param)) {
        return historyOrder[param]
      } else {
        return []
      }
    }
  },
  methods: {
    status (scope) {
      let status = scope.row.status
      const lang = this.$i18n.locale
      if (lang === 'zh') {
        if (status === 1) {
          return '已成交'
        } else if (status === 2) {
          return '未成交'
        } else if (status === 3) {
          return '撤单'
        } else if (status === 4) {
          return '下单失败'
        } else if (status === 5) {
          return '下单成功'
        } else {
          return '未知状况'
        }
      } else {
        if (status === 1) {
          return 'Deal Done'
        } else if (status === 2) {
          return 'Unfilled'
        } else if (status === 3) {
          return 'Withdrawal'
        } else if (status === 4) {
          return 'Order failed'
        } else if (status === 5) {
          return 'Order success'
        } else {
          return 'Unknown'
        }
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
      // const second = d.getSeconds() < 10 ? '0' + d.getSeconds() : '' + d.getSeconds()
      return month + '-' + day + ' ' + hour + ':' + minutes
    }
  }
}
</script>

<style scoped>
  @import '../../assets/css/TradeFooter.css';
</style>
