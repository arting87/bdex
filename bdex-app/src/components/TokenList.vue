<template>
  <div class="token-list">
    <div class="token-list-top" v-if="!sideMode">
      <tab id="token-list-top-tab" :tabList="tabList" :tabIndex.sync="tabIndex"/>
      <div class="search-area">
        <input class="search-input" v-model="search" placeholder="搜索"/>
      </div>
    </div>
    <div class="token-list-top-side" v-else>
      <tab id="token-list-top-tab" :tabList="tabList" :tabIndex.sync="tabIndex"/>
      <div class="search-area">
        <input class="search-input" v-model="search" placeholder="搜索"/>
      </div>
    </div>
    <div class="filter-header" v-if="!sideMode">
      <div class="pair-volume">
        <span>交易对/24H量</span>
      </div>
      <div class="last-price">
        <span>最新价</span>
      </div>
      <div class="rise-fall">
        <span>涨跌</span>
      </div>
    </div>
    <div class="token-info-area">
      <div v-if="filterData(api.favTableData).length > 0">
        <token-info-item :data="item" :simple="sideMode" @clickItem="clickItem" v-show="tabIndex === 0" v-for="(item,i) in filterData(api.favTableData)" :key="i" />
      </div>
      <div class="no-data" v-show="tabIndex === 0" v-else >暂无数据</div>
      <div v-if="filterData(api.eosTableData).length > 0">
        <token-info-item :data="item" :simple="sideMode" @clickItem="clickItem" v-show="tabIndex === 1" v-for="(item,i) in filterData(api.eosTableData)" :key="i" />
      </div>
      <div class="no-data" v-show="tabIndex === 1" v-else >暂无数据</div>
      <div v-if="filterData(api.ethTableData).length > 0">
        <token-info-item :data="item" :simple="sideMode" @clickItem="clickItem" v-show="tabIndex === 2" v-for="(item,i) in filterData(api.ethTableData)" :key="i" />
      </div>
      <div class="no-data" v-show="tabIndex === 2" v-else >暂无数据</div>
      <div v-if="filterData(api.eosSideTableData).length > 0">
        <token-info-item :data="item" :simple="sideMode" @clickItem="clickItem" v-show="tabIndex === 3" v-for="(item,i) in filterData(api.eosSideTableData)" :key="i" />
      </div>
      <div class="no-data" v-show="tabIndex === 3" v-else >暂无数据</div>
      <div v-if="filterData(api.crossChainTableData).length > 0">
        <token-info-item :data="item" :simple="sideMode" @clickItem="clickItem" v-show="tabIndex === 4" v-for="(item,i) in filterData(api.crossChainTableData)" :key="i" />
      </div>
      <div class="no-data" v-show="tabIndex === 4" v-else >暂无数据</div>
    </div>
  </div>
</template>

// 自选区方案：采用localStorage本地存储：[tokenInfo, tokenInfo...]
<script>
import Tab from './Tab'
import TokenInfoItem from './TokenInfoItem'

export default {
  name: 'token-list',
  components: {
    Tab,
    TokenInfoItem
  },
  props: ['sideMode'],
  data () {
    return {
      tabIndex: 4,
      tabList: [
        '自选', 'EOS', 'ETH', 'BDE', '跨链'
      ],
      search: ''
    }
  },
  methods: {
    clickItem () {
      this.$emit('clickItem')
    },
    filterData (oldData) {
      return oldData.filter(data => !this.search || data.pair.toLowerCase().includes(this.search.toLowerCase()))
    },
    changeFav (scope) {
      if (event.target.nodeName === 'I') {
        const tokenInfo = scope.row.pair + '-' + scope.row.code
        try {
          let localFavTableData = localStorage.getItem('localFavTableData')
          if (localFavTableData) {
            localFavTableData = JSON.parse(localFavTableData)
            for (let i = 0; i < localFavTableData.length; i++) {
              if (localFavTableData[i] === tokenInfo) {
                localFavTableData.splice(i, 1)
                localStorage.setItem('localFavTableData', JSON.stringify(localFavTableData))
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
        } catch (e) {
        }
      }
    }
  },
  computed: {
    api () {
      let eos = this.$store.state.api.eos
      let eth = this.$store.state.api.eth
      let eosSide = this.$store.state.api.eosSide
      let crossChain = this.$store.state.api.crossChain
      // 非自选区处理
      let eosTableData = []
      let ethTableData = []
      let eosSideTableData = []
      let crossChainTableData = []
      if (eos) {
        eosTableData = Object.values(eos)
      }
      if (eth) {
        ethTableData = Object.values(eth)
      }
      if (eosSide) {
        eosSideTableData = Object.values(eosSide)
      }
      if (crossChain) {
        crossChainTableData = Object.values(crossChain)
      }
      // 自选区处理, 暂时使用
      let favTableData = []
      let localFavTableData = []
      if (localStorage.getItem('localFavTableData')) {
        try {
          localFavTableData = localStorage.getItem('localFavTableData')
          if (localFavTableData) {
            localFavTableData = JSON.parse(localFavTableData) // JSON反序列化
            localFavTableData.forEach(e => {
              let tokenName = e.split('-')[0]
              let code = e.split('-')[1]
              if (code === '1' && eos) {
                if (eos.hasOwnProperty(tokenName)) {
                  favTableData.push(eos[tokenName])
                }
              } else if (code === '2' && eth) {
                if (eth.hasOwnProperty(tokenName)) {
                  favTableData.push(eth[tokenName])
                }
              } else if (code === '3' && eosSide) {
                if (eosSide.hasOwnProperty(tokenName)) {
                  favTableData.push(eosSide[tokenName])
                }
              } else if (code === '4' && crossChain) {
                if (crossChain.hasOwnProperty(tokenName)) {
                  favTableData.push(crossChain[tokenName])
                }
              }
            })
          }
        } catch (e) {
          localFavTableData = []
        }
      }
      return {
        eosTableData,
        ethTableData,
        eosSideTableData,
        crossChainTableData,
        favTableData
      }
    }
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
.token-list {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 37.5rem;
  .token-list-top-side {
    display: flex;
    flex-direction: column;
    align-items: center;
    .search-area {
      margin-top: .7rem;
      display: flex;
      align-items: center;
      font-size: 1.2rem;
      .search-input {
        width: 25.2rem;
        height: 2.8rem;
        background-color: #21263c;
        box-shadow: inset 0rem 0rem 0.1rem 0.2rem rgba(0, 0, 0, 0.5);
        border-radius: 0.4rem;
        color: $whiteTextColor;
        text-align: center;
      }
      .search-input::placeholder {
        text-align: center;
      }
    }
  }
  .token-list-top {
    height: 4rem;
    display: flex;
    #token-list-top-tab {
      width: 28rem;
    }
    .search-area {
      width: 9.5rem;
      height: 100%;
      display: flex;
      align-items: center;
      font-size: 1.2rem;
      .search-input {
        width: 100%;
        height: 100%;
        background-color: #131624;
        box-shadow: inset 0rem 0rem 1rem 0rem rgba(0, 0, 0, 0.5);
        border-radius: 0.4rem;
        color: $whiteTextColor;
        text-align: center;
      }
      .search-input::placeholder {
        text-align: center;
      }
    }
  }
  .filter-header {
    height: 4rem;
    width: 33.5rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 1.2rem;
    color: $whiteTextColor;
  }
}
</style>
