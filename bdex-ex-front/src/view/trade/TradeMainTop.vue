<template>
  <div class="trademaintop">
    <trade-view ref="trade"></trade-view>
  </div>
</template>

<script>
import TradeView from '../../components/TradeView'
export default {
  name: 'TradeMainTop',
  components: {
    TradeView
  },
  props: { // 组件间通信
    routeParam: {
      type: String,
      required: true
    }
  },
  computed: {
    lang () {
      let lang = this.$store.state.i18n.lang
      if (lang !== 'zh' && lang !== 'en') {
        lang = 'zh'
      }
      return lang
    }
  },
  watch: {
    routeParam: function (newVal) {
      let param = newVal
      let tokenName = param.split('-')[0]
      let code = param.split('-')[1]
      code = parseInt(code)
      this.$refs.trade.init(tokenName, 5, code) // tokenName, interval 区间, marketCode
    },
    lang: function (newVal) {
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      let code = param.split('-')[1]
      code = parseInt(code)
      this.$refs.trade.init(tokenName, 5, code, newVal, true)
    }
  },
  mounted () {
    let param = this.routeParam
    let tokenName = param.split('-')[0]
    let code = param.split('-')[1]
    code = parseInt(code)
    let lang = localStorage.lang
    if (lang !== 'zh' && lang !== 'en') {
      lang = 'zh'
    }
    this.$refs.trade.init(tokenName, 5, code, lang, false) // tokenName, interval 区间, marketCode, lang 语言, changeLang改变语言
  }
}
</script>

<style scoped>
.trademaintop {
  min-width: 596px;
  max-width: 596px;
  margin-bottom: 10px;
  max-height: 442px;
  border: 1px solid #171b2b;
}
</style>
