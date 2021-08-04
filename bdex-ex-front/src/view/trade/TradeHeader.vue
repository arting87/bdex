<template>
  <div class="tradeheader">
    <el-row>
      <el-col :span="4" :offset="0" class="pair"><!--todo 8.13 新增bde交易区-->
        <span class="frontCoin">{{ token === null ? '--' : token.pair }}</span>
        <span class="backCoin">/ {{ token === null ? '--' : (token.code === '3' ? 'BDE' : (token.code === '2' ? ' ETH' : ' EOS')) }}</span>
      </el-col>
      <el-col :span="3" :offset="0">
        <span class="backCoin">{{ $t('home.lastPrice') }}</span>
        <div>
          <span class="change">{{ token === null ? '--' : token.lastPrice }}</span>
          <span>{{ token === null ? '--' : (token.code === '3' ? 'BDE' : (token.code === '2' ? ' ETH' : ' EOS')) }}</span>
        </div>
      </el-col>
      <el-col :span="3">
        <span class="backCoin">{{ $t('home.24hChange') }}</span>
        <div class="change">{{ handleChange(token) }}</div>
      </el-col>
      <el-col :span="3">
        <span class="backCoin">{{ $t('home.24hHigh') }}</span>
        <div>{{ token === null ? '--' : token.highPrice }} {{ token === null ? '--' : (token.code === '3' ? 'BDE' : (token.code === '2' ? ' ETH' : ' EOS')) }}</div>
      </el-col>
      <el-col :span="3">
        <span class="backCoin">{{ $t('home.24hLow') }}</span>
        <div>{{ token === null ? '--' : token.lowPrice }} {{ token === null ? '--' : (token.code === '3' ? 'BDE' : (token.code === '2' ? ' ETH' : ' EOS')) }}</div>
      </el-col>
      <el-col :span="3">
        <span class="backCoin">{{ $t('home.24hVolume') }}</span>
        <div>{{ token === null ? '--' : token.volume }} {{ token === null ? '--' : token.pair }}</div>
      </el-col>
      <el-col :span="4">
        <span class="backCoin">{{ $t('home.24hTotal') }}</span>
        <div>{{ token === null ? '--' : token.total }} {{ token === null ? '--' : (token.code === '3' ? 'BDE' : (token.code === '2' ? ' ETH' : ' EOS')) }}</div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'TradeHeader',
  props: { // 组件间通信
    routeParam: {
      type: String,
      required: true
    }
  },
  computed: {
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
  }
}
</script>

<style scoped>
  @import '../../assets/css/TradeHeader.css';
</style>
