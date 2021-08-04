<template>
  <div class="news">
    <navbar text="通知公告" type="bright" />
    <div class="lastest">
      <div class="news-title">
        <span>最新公告</span>
      </div>
      <news-item :title="item" v-for="(item,i) in lastestList" :key="i" />
    </div>
    <!-- <div class="old">
      <div class="news-title">
        <span>通告</span>
      </div>
      <news-item :title="item" v-for="(item,i) in lastestList" :key="i" />
    </div> -->
  </div>
</template>

<script>
import Navbar from '../../components/Navbar'
import NewsItem from '../../components/NewsItem'

export default {
  name: 'news',
  data () {
    return {
      lastestList: [
        '上币申请公告'
      ]
    }
  },
  components: {
    Navbar,
    NewsItem
  },
  methods: {
  }
}
</script>

<style lang="scss" scoped>
.news {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #ffffff;
  height: 100vh;
  .lastest, .old {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    .news-title {
      height: 6rem;
      width: 33.5rem;
      display: flex;
      align-items: center;
      font-size: 1.4rem;
      color: #333333;
      font-weight: 500;
    }
  }
}
</style>
