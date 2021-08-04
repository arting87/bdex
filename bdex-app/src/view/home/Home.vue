<template>
  <div class="home">
    <div class="swipe-area">
      <van-swipe :autoplay="3000">
        <van-swipe-item v-for="(item, index) in images" :key="index">
          <img :src="item.url">
          <h3>{{item.titlezh}}</h3>
          <h5>{{item.titleen}}</h5>
        </van-swipe-item>
      </van-swipe>
    </div>
    <div class="news-area">
      <div></div>
      <span @click="gotoNews">上币申请公告</span>
      <div @click="goto('/index/news')"></div>
    </div>
    <manage-center />

    <van-tabs
      v-model="active"
      type="card"
      border
      background="#131624"
      color="#21263C"
      title-active-color="#fff"
      title-inactive-color="#fff"
    >

      <van-tab title="涨幅榜">
          <van-cell
          v-for="item in changList"
          :key="item"
          label="24h量"
          class="van-cell"
          title-style="color: #ffffff"
          >
          <template slot="title">
            <span>{{ item.pair }}</span>
          </template>
          <template slot="default">
            <van-tag type="danger" style="font-size:14px">{{ handleChange(item.change) }}</van-tag>
          </template>
          </van-cell>
      </van-tab>

      <van-tab title="24h成交量榜" style="padding: 0px;">
          <van-cell
          v-for="item in volList"
          :key="item"
          :title="item"
          label="24h量"
          class="van-cell"
          title-style="color: #ffffff"
          >
            <template slot="title">
            <span>{{ item.tokenName }}/EOS</span>
            </template>
            <template slot="default">
            <van-tag type="danger" style="font-size:14px">{{item.vol}} EOS</van-tag>
            </template>
          </van-cell>
      </van-tab>
    </van-tabs>
  </div>
</template>

<script>
import ManageCenter from '../../components/ManageCenter'

export default {
  name: 'Home',
  components: {
    ManageCenter
  },
  data () {
    return {
      active: 0,
      search: '',
      images: [
        { id: 0, url: require('../../assets/img/home4.png'), titlezh: '全球首个去中心化多协议跨链交易平台', titleen: 'The worlds first decentralized multi-protocol cross-chain trading platform' },
        { id: 1, url: require('../../assets/img/home1.png'), titlezh: '全球首个去中心化多协议跨链交易平台', titleen: `The world's first decentralized multi-protocol cross-chain trading platform` },
        { id: 2, url: require('../../assets/img/home3.png'), titlezh: '全球首个去中心化多协议跨链交易平台', titleen: `The world's first decentralized multi-protocol cross-chain trading platform` }
      ]
    }
  },
  computed: {
    changList () {
      return this.$store.state.api.changeList
    },
    volList () {
      return this.$store.state.api.volList
    }
  },
  methods: {
    handleChange (change) { // 更新数据是否会重新渲染？待测试
      if (change >= 0) {
        return '+ ' + change + ' %'
      } else {
        return change + '%'
      }
    },
    goto (url) {
      this.$router.push(url)
    },
    gotoNews () {
      this.$router.push('/index/newsdetails')
    }
  }
}
</script>

<style lang="scss" scoped>
.home {
  display: flex;
  flex-direction: column;
  .swipe-area {
    height: 17.2rem;
    img {
      height: 17.2rem;
      width: 37.5rem;
    }
    h3, h5 {
      margin-top: -8rem;
      color: #B3B5BD;
    }
    .van-swipe {
      height: 100%;
    }
    .van-swipe-item {
      display: flex;
      flex-direction: column;
      align-items: center;
    }
  }
  .news-area {
    height: 3.6rem;
    display: flex;
    align-items: center;
    justify-content: space-between;
    background-color: #21263c;
    color: #e6e6e6;
    font-size: 1.4rem;
    span {
      margin-left: -7rem;
    }
    :first-child {
      margin-left: 2rem;
      width: 1.1rem;
      height: 1.2rem;
      background-image: url('../../../static/images/公告喇叭@2x.png');
      background-repeat: no-repeat;
      background-size: cover;
    }
    :last-child {
      margin-right: 2rem;
      height: .9rem;
      width: 1.1rem;
      background-image: url('../../../static/images/公告菜单@2x.png');
      background-repeat: no-repeat;
      background-size: cover;
    }
  }
  .van-cell {
    background-color: #131624;
    padding: 10px 16px;
  }
  .van-tabs--card {
    padding-top: 5px;
  }
}
</style>
