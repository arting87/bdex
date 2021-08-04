<template>
  <div class="traderight">
    <el-tabs type="border-card">
      <el-tab-pane>
        <span slot="label">
          {{$t('trade.trade') }}
        </span>
        <!-- sell and buy 缩写为sb -->
        <el-table
          :data="order.sell === null ? [] : order.sell.slice(0, 15)"
          style="width: 100%"
          @row-click="setPrice"
          height="395"
          class="sbtable"
        ><!--todo 8.13新需求，改动,bde交易区-->
          <el-table-column :label="$t('trade.price') + '(' + (token.code === '3' ? 'BDE' : (token.code === '2' ? 'ETH' : 'EOS')) + ')'" width="105">
            <template prop="price" slot-scope="scope">
              <span class="color-red">{{ handFloat(scope.row.price, 6) }}</span>
            </template>
          </el-table-column>
          <el-table-column :label="$t('trade.amount') + '(' + token.tokenName + ')' " width="85">
            <template prop="amount" slot-scope="scope">
              <span>{{ handFloat(scope.row.amt, 4) }}</span>
            </template>
          </el-table-column>
          <el-table-column :label="$t('trade.total') + '(' + (token.code === '3' ? 'BDE' : (token.code === '2' ? 'ETH' : 'EOS')) + ')' " width="95">
            <template prop="total" slot-scope="scope">
              <span>{{ handFloat(NH.times(scope.row.price, scope.row.amt), 4) }}</span>
            </template>
          </el-table-column>
        </el-table>

        <!-- {{$t('trade.lastprice') }}  -->
        <el-row>
          <el-col :span="24" class="middle">
            <span>{{$t('trade.lastprice') }}</span>
            <span class="lastprice">{{ order.lastPrice === null ? '-- --' : order.lastPrice }}</span>
          </el-col>
        </el-row>

        <el-table
          :data="order.buy === null ? [] : order.buy.slice(0, 15)"
          @row-click="setPrice"
          style="width: 100%"
          height="372"
          class="btable"
        >
          <el-table-column prop="price" width="105">
            <template slot-scope="scope">
              <span class="color-green">{{ handFloat(scope.row.price, 6) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="amount" width="85">
            <template slot-scope="scope">
              <span>{{ handFloat(scope.row.amt, 4) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="total" width="95">
            <template slot-scope="scope">
              <span>{{ handFloat(NH.times(scope.row.price, scope.row.amt), 4) }}</span>
            </template>
          </el-table-column>
        </el-table>

      </el-tab-pane>

      <el-tab-pane>
        <span slot="label">
          {{$t('trade.buy') }}
        </span>
        <el-table
          :data="order.buy === null ? [] : order.buy"
          @row-click="setPrice"
          style="width: 100%"
          height="816"
          class="sbtable"
        >
          <el-table-column prop="price" :label="$t('trade.price') + '(' + (token.code === '3' ? 'BDE' : (token.code === '2' ? 'ETH' : 'EOS')) + ')' " width="105">
            <template slot-scope="scope">
              <span class="color-green">{{ handFloat(scope.row.price, 6) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="amount" :label="$t('trade.amount') + '(' + token.tokenName + ')' " width="85">
            <template slot-scope="scope">
              <span>{{ handFloat(scope.row.amt, 4) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="total" :label="$t('trade.total') + '(' + (token.code === '3' ? 'BDE' : (token.code === '2' ? 'ETH' : 'EOS')) + ')' " width="95">
            <template slot-scope="scope">
              <span>{{ handFloat(NH.times(scope.row.price, scope.row.amt), 4) }}</span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

      <el-tab-pane>
        <span slot="label">
          {{$t('trade.sell') }}
        </span>
        <el-table
          :data="order.sell === null ? [] : order.sell"
          @row-click="setPrice"
          style="width: 100%"
          height="816"
          class="sbtable"
        >
          <el-table-column prop="price" :label="$t('trade.price') + '(' + (token.code === '3' ? 'BDE' : (token.code === '2' ? 'ETH' : 'EOS')) + ')' " width="105">
            <template slot-scope="scope">
              <span class="color-red">{{ handFloat(scope.row.price, 6) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="amount" :label="$t('trade.amount') + '(' + token.tokenName + ')' " width="85">
            <template slot-scope="scope">
              <span>{{ handFloat(scope.row.amt, 4) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="total" :label="$t('trade.total') + '(' + (token.code === '3' ? 'BDE' : (token.code === '2' ? 'ETH' : 'EOS')) + ')' " width="95">
            <template slot-scope="scope">
              <span>{{ handFloat((scope.row.price * scope.row.amt), 4) }}</span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import {numFloat, handBig} from '../../utils/handleFloat'
export default {
  name: 'TradeRight',
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
    token: function () {
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      let code = param.split('-')[1]
      return {
        tokenName,
        code
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
    setPrice (row) { // 点击挂单的，自动把价格填充到TradeMainBottom price输入框; 以vuex解决
      let price = row.price
      this.$store.commit('orderApi/setPrice', price)
    }
  }
}
</script>

<style scoped>
  @import '../../assets/css/TradeRight.css';
</style>
