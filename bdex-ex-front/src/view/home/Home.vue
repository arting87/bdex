<template>
  <div class="home">
    <el-row class="grid-content">
      <!-- <div class="dataoverview">
        <HomeDataOverview/> -->
      <!-- </div> -->
    <div class="block">
      <el-carousel height="648px">
        <el-carousel-item v-for="item in imagesbox" :key="item.id">
          <div class="bannertitle">
            <h3>{{item.titlezh}}</h3>
            <h5>{{item.titleen}}</h5>
          </div>
          <el-image :src="item.imgurl" class="image" fit="cover"></el-image>
        </el-carousel-item>
      </el-carousel>
    </div>
      <!-- <img :src="imgurl" class="image"> -->
    </el-row>

    <!-- 交易对数据 -->
        <!-- 搜索 -->
      <div class="searchbox">
        <el-input class="home-search" v-model="search" clearable size="mini" :placeholder="$t('home.search')"/>
      </div>
    <el-tabs v-model="activeName" @tab-click="handleClick" type="border-card">
      <!-- 自选区:favorites 缩写为fav -->

      <el-tab-pane :label="$t('home.favorites')" name="first">
        <el-table
          :data="api.favTableData.filter(data => !search || data.pair.toLowerCase().includes(search.toLowerCase()))"
          @row-click="rowClick"
          style="width: 100%"
        >
          <el-table-column prop="pair" :label="$t('home.pairs')" sortable width="160">
            <template slot-scope="scope">
              <span v-if="isShow(scope)"><i class="el-icon-star-on" @click="changeFav(scope)"></i></span>
              <span v-else=""><i class="el-icon-star-off" @click="changeFav(scope)"></i></span>
              <span style="margin-left: 15px">{{ scope.row.code === '2' ? scope.row.pair + ' / ETH' : scope.row.pair + ' / EOS' }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="lastprice" :label="$t('home.lastPrice')" width="160">
            <template slot-scope="scope">
              <span>{{ scope.row.lastPrice === '' ? '--' : scope.row.code === '2' ? scope.row.lastPrice + ' ETH' : scope.row.lastPrice + ' EOS'}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="change" :label="$t('home.24hChange')" sortable width="200">
            <template slot-scope="scope">
              <span>{{ handleChange(scope) }} </span>
            </template>
          </el-table-column>
          <el-table-column prop="high" :label="$t('home.24hHigh')" width="140">
            <template slot-scope="scope">
              <span>{{ scope.row.highPrice === '' ? '--' : scope.row.code === '2' ? scope.row.highPrice + ' ETH' : scope.row.highPrice + ' EOS'}} </span>
            </template>
          </el-table-column>
          <el-table-column prop="low" :label="$t('home.24hLow')" width="140">
            <template slot-scope="scope">
              <span>{{ scope.row.lowPrice === '' ? '--' : scope.row.code === '2' ? scope.row.lowPrice + ' ETH' : scope.row.lowPrice + ' EOS'}} </span>
            </template>
          </el-table-column>
          <el-table-column prop="vol" :label="$t('home.24hVolume')" sortable width="185">
            <template slot-scope="scope">
              <span>{{ scope.row.volume === '' ? '--' : scope.row.volume + ' ' + scope.row.pair}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="total" :label="$t('home.24hTotal')" sortable width="185">
            <template slot-scope="scope">
              <span>{{ scope.row.total === '' ? '--' : scope.row.code === '2' ?  scope.row.total + ' ETH' : scope.row.total + ' EOS'}}</span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

      <!-- EOS链内 -->
      <el-tab-pane :label="$t('home.eosMain')" name="second">
        <!--data:显示的数据
        :default-sort="{prop: 'pair', order: 'ascending'}" 默认排序-->
        <el-table
          :data="api.eosTableData.filter(data => !search || data.pair.toLowerCase().includes(search.toLowerCase()))"
          @row-click="rowClick"
          style="width: 100%"
        >
          <el-table-column prop="pair" :label="$t('home.pairs')" sortable width="160">
            <template slot-scope="scope">
              <span v-if="isShow(scope)"><i class="el-icon-star-on" @click="changeFav(scope)"></i></span>
              <span v-else=""><i class="el-icon-star-off" @click="changeFav(scope)"></i></span>
              <span style="margin-left: 15px">{{ scope.row.pair }} / EOS</span>
            </template>
          </el-table-column>
          <el-table-column prop="lastprice" :label="$t('home.lastPrice')" width="160">
            <template slot-scope="scope">
              <span>{{ scope.row.lastPrice === '' ? '--' : scope.row.lastPrice + ' EOS'}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="change" :label="$t('home.24hChange')" sortable width="200">
            <template slot-scope="scope">
              <span>{{ handleChange(scope) }} </span>
            </template>
          </el-table-column>
          <el-table-column prop="high" :label="$t('home.24hHigh')" width="140">
            <template slot-scope="scope">
              <span>{{ scope.row.highPrice === '' ? '--' : scope.row.highPrice + ' EOS'}} </span>
            </template>
          </el-table-column>
          <el-table-column prop="low" :label="$t('home.24hLow')" width="140">
            <template slot-scope="scope">
              <span>{{ scope.row.lowPrice === '' ? '--' : scope.row.lowPrice + ' EOS'}} </span>
            </template>
          </el-table-column>
          <el-table-column prop="vol" :label="$t('home.24hVolume')" sortable width="185">
            <template slot-scope="scope">
              <span>{{ scope.row.volume === '' ? '--' : scope.row.volume + ' ' + scope.row.pair}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="total" :label="$t('home.24hTotal')" sortable width="185">
            <template slot-scope="scope">
              <span>{{ scope.row.total === '' ? '--' : scope.row.total + ' EOS'}}</span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

      <!-- ETH链内 -->
      <el-tab-pane :label="$t('home.ethSide')" name="third">
        <el-table
          :data="api.ethTableData.filter(data => !search || data.pair.toLowerCase().includes(search.toLowerCase()))"
          @row-click="rowClick"
          style="width: 100%"
        >
          <el-table-column prop="pair" :label="$t('home.pairs')" sortable width="160">
            <template slot-scope="scope">
              <span v-if="isShow(scope)"><i class="el-icon-star-on" @click="changeFav(scope)"></i></span>
              <span v-else=""><i class="el-icon-star-off" @click="changeFav(scope)"></i></span>
              <span style="margin-left: 15px">{{ scope.row.pair }} / ETH</span>
            </template>
          </el-table-column>
          <el-table-column prop="lastprice" :label="$t('home.lastPrice')" width="160">
            <template slot-scope="scope">
              <span>{{ scope.row.lastPrice === '' ? '--' : scope.row.lastPrice + ' ETH'}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="change" :label="$t('home.24hChange')" sortable width="200">
            <template slot-scope="scope">
              <span>{{ handleChange(scope) }} </span>
            </template>
          </el-table-column>
          <el-table-column prop="high" :label="$t('home.24hHigh')" width="140">
            <template slot-scope="scope">
              <span>{{ scope.row.highPrice === '' ? '--' : scope.row.highPrice + ' ETH'}} </span>
            </template>
          </el-table-column>
          <el-table-column prop="low" :label="$t('home.24hLow')" width="140">
            <template slot-scope="scope">
              <span>{{ scope.row.lowPrice === '' ? '--' : scope.row.lowPrice + ' ETH'}} </span>
            </template>
          </el-table-column>
          <el-table-column prop="vol" :label="$t('home.24hVolume')" sortable width="185">
            <template slot-scope="scope">
              <span>{{ scope.row.volume === '' ? '--' : scope.row.volume + ' ' + scope.row.pair}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="total" :label="$t('home.24hTotal')" sortable width="185">
            <template slot-scope="scope">
              <span>{{ scope.row.total === '' ? '--' : scope.row.total + ' ETH'}}</span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

      <!-- EOS侧链 --><!--todo 8.13 新需求，现作为eth erc20代币bdex的交易区-->
       <el-tab-pane :label="$t('home.eosSide')" name="fourth">
        <el-table
          :data="api.eosSideTableData.filter(data => !search || data.pair.toLowerCase().includes(search.toLowerCase()))"
          @row-click="rowClick"
          style="width: 100%"
        >
          <el-table-column prop="pair" :label="$t('home.pairs')" sortable width="160">
            <template slot-scope="scope">
              <span v-if="isShow(scope)"><i class="el-icon-star-on" @click="changeFav(scope)"></i></span>
              <span v-else=""><i class="el-icon-star-off" @click="changeFav(scope)"></i></span>
              <span style="margin-left: 15px">{{ scope.row.pair }} / BDE</span>
            </template>
          </el-table-column>
          <el-table-column prop="lastprice" :label="$t('home.lastPrice')" width="160">
            <template slot-scope="scope">
              <span>{{ scope.row.lastPrice === '' ? '--' : scope.row.lastPrice + ' BDE'}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="change" :label="$t('home.24hChange')" sortable width="200">
            <template slot-scope="scope">
              <span>{{ handleChange(scope) }} </span>
            </template>
          </el-table-column>
          <el-table-column prop="high" :label="$t('home.24hHigh')" width="140">
            <template slot-scope="scope">
              <span>{{ scope.row.highPrice === '' ? '--' : scope.row.highPrice + ' BDE'}} </span>
            </template>
          </el-table-column>
          <el-table-column prop="low" :label="$t('home.24hLow')" width="140">
            <template slot-scope="scope">
              <span>{{ scope.row.lowPrice === '' ? '--' : scope.row.lowPrice + ' BDE'}} </span>
            </template>
          </el-table-column>
          <el-table-column prop="vol" :label="$t('home.24hVolume')" sortable width="185">
            <template slot-scope="scope">
              <span>{{ scope.row.volume === '' ? '--' : scope.row.volume + ' ' + scope.row.pair}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="total" :label="$t('home.24hTotal')" sortable width="185">
            <template slot-scope="scope">
              <span>{{ scope.row.total === '' ? '--' : scope.row.total + ' BDE'}}</span>
            </template>
          </el-table-column>
        </el-table>
       </el-tab-pane>

      <!-- 跨链:Across the chain -->
      <el-tab-pane :label="$t('home.crossChain')" name="fifth">
        <el-table
          :data="api.crossChainTableData.filter(data => !search || data.pair.toLowerCase().includes(search.toLowerCase()))"
          @row-click="rowClick"
          style="width: 100%"
        >
          <el-table-column prop="pair" :label="$t('home.pairs')" sortable width="160">
            <template slot-scope="scope">
              <span v-if="isShow(scope)"><i class="el-icon-star-on" @click="changeFav(scope)"></i></span>
              <span v-else=""><i class="el-icon-star-off" @click="changeFav(scope)"></i></span>
              <span style="margin-left: 15px">{{ scope.row.pair }} / EOS</span>
            </template>
          </el-table-column>
          <el-table-column prop="lastprice" :label="$t('home.lastPrice')" width="160">
            <template slot-scope="scope">
              <span>{{ scope.row.lastPrice === '' ? '--' : scope.row.lastPrice.toFixed(4) + ' EOS'}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="change" :label="$t('home.24hChange')" sortable width="200">
            <template slot-scope="scope">
              <span>{{ handleChange(scope) }} </span>
            </template>
          </el-table-column>
          <el-table-column prop="high" :label="$t('home.24hHigh')" width="140">
            <template slot-scope="scope">
              <span>{{ scope.row.highPrice === '' ? '--' : scope.row.highPrice + ' EOS'}} </span>
            </template>
          </el-table-column>
          <el-table-column prop="low" :label="$t('home.24hLow')" width="140">
            <template slot-scope="scope">
              <span>{{ scope.row.lowPrice === '' ? '--' : scope.row.lowPrice + ' EOS'}} </span>
            </template>
          </el-table-column>
          <el-table-column prop="vol" :label="$t('home.24hVolume')" sortable width="185">
            <template slot-scope="scope">
              <span>{{ scope.row.volume === '' ? '--' : scope.row.volume + ' ' + scope.row.pair}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="total" :label="$t('home.24hTotal')" sortable width="185">
            <template slot-scope="scope">
              <span>{{ scope.row.total === '' ? '--' : scope.row.total + ' EOS'}}</span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

// 自选区方案：采用localStorage本地存储：[tokenInfo, tokenInfo...]
<script>
import HomeDataOverview from './HomeDataOverview'
export default {
  name: 'Home',
  components: {
    HomeDataOverview
  },
  data () {
    return {
      search: '',
      activeName: sessionStorage.getItem('homeActive') || 'second',
      icon1: 'el-icon-star-off',
      icon2: '',
      imagesbox: [
        { id: 0, imgurl: require('../../assets/img/home4.png'), titlezh: '全球首个去中心化多协议跨链交易平台', titleen: 'The worlds first decentralized multi-protocol cross-chain trading platform' },
        { id: 1, imgurl: require('../../assets/img/home1.png'), titlezh: '全球首个去中心化多协议跨链交易平台', titleen: `The world's first decentralized multi-protocol cross-chain trading platform` },
        { id: 2, imgurl: require('../../assets/img/home3.png'), titlezh: '全球首个去中心化多协议跨链交易平台', titleen: `The world's first decentralized multi-protocol cross-chain trading platform` }
      ]
    }
  },
  methods: {
    formatter (row, column) {
      return row.address
    },
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
          path: '/trade/' + row.pair + '-' + row.code // todo 8.13新增交易区，code新增为3的情况，以BDE作为参考币
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
  beforeDestroy () {
    sessionStorage.setItem('homeActive', this.activeName)
    // localStorage.homeActive = this.activeName // 保存用户离开此路由时，选中的tab
  },
  beforeRouteEnter (to, from, next) {
    // sessionStorage 防止注入攻击
    let testStorage = ['first', 'second', 'third', 'fourth', 'fifth'] // eos侧链现在不支持，故删除了'fourth';todo 8.13 新增
    if (!testStorage.includes(sessionStorage.getItem('homeActive'))) {
      sessionStorage.setItem('homeActive', 'second')
      // localStorage.homeActive = 'second'
    }
    next()
  }
}
</script>

<style scoped>
  @import '../../assets/css/Home.css';
</style>
