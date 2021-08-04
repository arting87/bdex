<template>
  <div class="market">
    <token-list :sideMode="false" />
  </div>
</template>

// 自选区方案：采用localStorage本地存储：[tokenInfo, tokenInfo...]
<script>
import TokenList from '../../components/TokenList'

export default {
  name: 'market',
  components: {
    TokenList
  },
  data () {
    return {
      tabIndex: 3,
      tabList: [
        '自选', 'EOS', 'ETH', 'BDE', '跨链'
      ],
      search: ''
    }
  },
  methods: {
  },
  computed: {
  },
  beforeDestroy () {
    sessionStorage.setItem('homeActive', this.activeName)
    // localStorage.homeActive = this.activeName // 保存用户离开此路由时，选中的tab
  },
  beforeRouteEnter (to, from, next) {
    // sessionStorage 防止注入攻击
    let testStorage = ['first', 'second', 'third', 'fifth'] // eos侧链现在不支持，故删除了'fourth'
    if (!testStorage.includes(sessionStorage.getItem('homeActive'))) {
      sessionStorage.setItem('homeActive', 'second')
      // localStorage.homeActive = 'second'
    }
    next()
  }
}
</script>

<style lang="scss" scoped>
.market {
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
