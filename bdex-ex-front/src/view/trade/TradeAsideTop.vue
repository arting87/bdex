<template>
  <div class="tradeasidetop">
    <!-- 搜索币种  type="card"  type="border-card" -->
    <div class="search">
      <el-input v-model="search" size="mini" clearable :placeholder="$t('trade.placeholder')"/>
    </div>

    <el-tabs v-model="activeName" @tab-click="handleClick" type="border-card">

      <!-- 自选 -->
      <el-tab-pane :label="$t('trade.favorite')" name="first">
        <el-table
          :data="api.favTableData.filter(data => !search || data.pair.toLowerCase().includes(search.toLowerCase()))"
          @row-click="rowClick"
          style="width: 100%"
          height="335"
        >
          <el-table-column prop="pair" :label="$t('trade.pair')" width="115">
            <template slot-scope="scope">
              <span v-if="isShow(scope)"><i class="el-icon-star-on" @click="changeFav(scope)"></i></span>
              <span v-else=""><i class="el-icon-star-off" @click="changeFav(scope)"></i></span>
              <span>{{ scope.row.code === '3' ? (scope.row.pair + ' / BDE') : (scope.row.code === '2' ? scope.row.pair + ' / ETH' : scope.row.pair + ' / EOS') }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="lastprice" :label="$t('trade.lastprice')" width="90">
            <template slot-scope="scope">
              <span>{{ scope.row.lastPrice === '' ? '--' : scope.row.lastPrice}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="change" :label="$t('trade.change')" sortable width="105">
            <template slot-scope="scope">
              <span>{{ handleChange(scope) }} </span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

      <!-- EOS @row-click="$router.push({name:'trade',params: {id: tableData.pair}})" ({ path: `/trade/${data.pair}` }) -->
      <el-tab-pane label="EOS" name="second">
        <el-table
          :data="api.eosTableData.filter(data => !search || data.pair.toLowerCase().includes(search.toLowerCase()))"
          @row-click="rowClick"
          style="width: 100%"
          height="335"
        >
          <el-table-column prop="pair" :label="$t('trade.pair')" width="115">
            <template slot-scope="scope">
              <span v-if="isShow(scope)"><i class="el-icon-star-on" @click="changeFav(scope)"></i></span>
              <span v-else=""><i class="el-icon-star-off" @click="changeFav(scope)"></i></span>
              <span>{{ scope.row.pair }} / EOS</span>
            </template>
          </el-table-column>
          <el-table-column prop="lastprice" :label="$t('trade.lastprice')" width="90">
            <template slot-scope="scope">
              <span>{{ scope.row.lastPrice === '' ? '--' : scope.row.lastPrice }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="change" :label="$t('trade.change')" sortable width="105">
            <template slot-scope="scope">
              <span>{{ handleChange(scope) }} </span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

      <!-- ETH -->
      <el-tab-pane label="ETH" name="third">
        <el-table
          :data="api.ethTableData.filter(data => !search || data.pair.toLowerCase().includes(search.toLowerCase()))"
          @row-click="rowClick"
          style="width: 100%"
          height="335"
        >
          <el-table-column prop="pair" :label="$t('trade.pair')" width="115">
            <template slot-scope="scope">
              <span v-if="isShow(scope)"><i class="el-icon-star-on" @click="changeFav(scope)"></i></span>
              <span v-else=""><i class="el-icon-star-off" @click="changeFav(scope)"></i></span>
              <span>{{ scope.row.pair }} / ETH</span>
            </template>
          </el-table-column>
          <el-table-column prop="lastprice" :label="$t('trade.lastprice')" width="90">
            <template slot-scope="scope">
              <span>{{ scope.row.lastPrice === '' ? '--' : scope.row.lastPrice }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="change" :label="$t('trade.change')" sortable width="105">
            <template slot-scope="scope">
              <span>{{ handleChange(scope) }} </span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

      <!-- EOS侧链 --><!--todo 8.13 新增BDE交易区，解决方案：使用之前预留的eos同构侧链eosSide字段解决；这里需要调整样式，待处理-->
      <el-tab-pane :label="$t('trade.sideChain')" name="fourth">
        <el-table
          :data="api.eosSideTableData.filter(data => !search || data.pair.toLowerCase().includes(search.toLowerCase()))"
          @row-click="rowClick"
          style="width: 100%"
          height="335"
        >
          <el-table-column prop="pair" :label="$t('trade.pair')" width="115">
            <template slot-scope="scope">
              <span v-if="isShow(scope)"><i class="el-icon-star-on" @click="changeFav(scope)"></i></span>
              <span v-else=""><i class="el-icon-star-off" @click="changeFav(scope)"></i></span>
              <span>{{ scope.row.pair }} / BDE</span>
            </template>
          </el-table-column>
          <el-table-column prop="lastprice" :label="$t('trade.lastprice')" width="90">
            <template slot-scope="scope">
              <span>{{ scope.row.lastPrice === '' ? '--' : scope.row.lastPrice }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="change" :label="$t('trade.change')" sortable width="105">
            <template slot-scope="scope">
              <span>{{ handleChange(scope) }} </span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

      <!-- 跨链 -->
      <el-tab-pane :label="$t('trade.cross')" name="fifth">
        <el-table
          :data="api.crossChainTableData.filter(data => !search || data.pair.toLowerCase().includes(search.toLowerCase()))"
          @row-click="rowClick"
          style="width: 100%"
          height="335"
        >
          <el-table-column prop="pair" :label="$t('trade.pair')" width="115">
            <template slot-scope="scope">
              <span v-if="isShow(scope)"><i class="el-icon-star-on" @click="changeFav(scope)"></i></span>
              <span v-else=""><i class="el-icon-star-off" @click="changeFav(scope)"></i></span>
              <span>{{ scope.row.pair }} / EOS</span>
            </template>
          </el-table-column>
          <el-table-column prop="lastprice" :label="$t('trade.lastprice')" width="90">
            <template slot-scope="scope">
              <span>{{ scope.row.lastPrice === '' ? '--' : scope.row.lastPrice }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="change" :label="$t('trade.change')" sortable width="105">
            <template slot-scope="scope">
              <span>{{ handleChange(scope) }} </span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

    </el-tabs>
  </div>
</template>

<script>
export default {
  name: 'TradeAsideTop',
  props: { // 组件间通信
    routeParam: {
      type: String,
      required: true
    }
  },
  data () {
    return {
      activeNameEosSide: 'one',
      search: '',
      activeName: sessionStorage.getItem('asideTopActive') || 'fifth',
      icon1: 'el-icon-star-off',
      icon2: '',
      favTableData: [],
      eosTableData: this.$store.state.api.eos === null ? [] : Object.values(this.$store.state.api.eos),
      ethTableData: this.$store.state.api.eth === null ? [] : Object.values(this.$store.state.api.eth),
      eosSideTableData: this.$store.state.api.eosSide === null ? [] : Object.values(this.$store.state.api.eosSide),
      crossChainTableData: this.$store.state.api.crossChain === null ? [] : Object.values(this.$store.state.api.crossChain)
    }
  },
  methods: {
    handleClick (tab, event) {
    },
    handleChange (scope) { // 处理涨跌幅数据
      if (scope.row.change === '') {
        return '--'
      } else if (scope.row.change >= 0) {
        return '+' + scope.row.change + '%'
      } else {
        return scope.row.change + '%'
      }
    },
    rowClick (row) { // 行情列表，row-click
      if (event.target.nodeName === 'DIV') {
        this.$router.push({
          path: '/trade/' + row.pair + '-' + row.code
        })
      }
    },
    isShow (scope) {
      const tokenInfo = scope.row.pair + '-' + scope.row.code
      try {
        let localFavTableData = localStorage.getItem('localFavTableData')
        if (localFavTableData) {
          localFavTableData = JSON.parse(localFavTableData)
          for (let i = 0; i < localFavTableData.length; i++) {
            if (localFavTableData[i] === tokenInfo) {
              return true
            }
          }
        }
        return false
      } catch (e) {
        return false
      }
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
  beforeDestroy () { // 动态路由时不会调用，非动态路由时可能会调用，，，，待测试
    sessionStorage.setItem('asideTopActive', this.activeName)
  },
  beforeRouteEnter (to, from, next) {
    // sessionStorage 防止注入攻击
    let testStorage = ['first', 'second', 'third', 'fourth', 'fifth'] // eos侧链现在不支持，故删除了'fourth' todo 8.13 新需求，改动：新增fourth
    if (!testStorage.includes(sessionStorage.getItem('asideTopActive'))) {
      sessionStorage.setItem('asideTopActive', 'fifth')
    }
    next()
  }
}
</script>

<style scoped>
  @import '../../assets/css/TradeAsideTop.css';
</style>
