<template>
  <div class="news">
    <div class="announcement-center">

      <el-row class="centertitle">
        <el-col :span="18" :offset="2">{{ $t('news.announcement') }}</el-col>
      </el-row>

      <el-row class="newstop">
        <!-- 重要公告 -->
        <el-col :span="10" :offset="2">
          <h3>{{ $t('news.importantAnnouncement') }}</h3>
          <div class="latestnews">
            <div
              class="newstitel"
              v-for="(newstitle,index) in news.importantnews"
              @click="newsContent(newstitle.id)"
              :key="index"
            ><el-link type="_blank">{{ newstitle.title }}</el-link></div>
          </div>
        </el-col>
        <!-- 活动公告 -->
        <el-col :span="10" :offset="1">
          <h3>{{ $t('news.activity') }}</h3>
          <div class="activitynews">
            <div
              class="newstitel"
              v-for="(newstitle,index) in news.activitynews"
              @click="newsContent(newstitle.id)"
              :key="index"
            ><el-link type="_blank">{{ newstitle.title }}</el-link></div>
          </div>
        </el-col>
      </el-row>
      <!-- 新币上线 -->
      <el-row class="newsbottom">
        <el-col :span="10" :offset="2">
          <h3>{{ $t('news.currencyAnnouncement') }}</h3>
          <div class="latestnews">
            <div
              class="newstitel"
              v-for="(newstitle,index) in news.currencynews"
              @click="newsContent(newstitle.id)"
              :key="index"
            ><el-link type="_blank">{{ newstitle.title }}</el-link></div>
          </div>
        </el-col>
        <!-- 其他公告 -->
        <el-col :span="10" :offset="1">
          <h3>{{ $t('news.others') }}</h3>
          <div class="activitynews">
            <div
              class="newstitel"
              v-for="(newstitle,index) in news.othersnews"
              @click="newsContent(newstitle.id)"
              :key="index"
            ><el-link type="_blank">{{ newstitle.title }}</el-link></div>
          </div>
        </el-col>
      </el-row>

    </div>
  </div>
</template>

<script>
import Vue from 'vue'
export default {
  name: 'News',
  data () {
    return {
      activeName: 'first',
      importantnews: [
        { title: 'BDEX正式上线' }
      ],
      activitynews: [
        { title: '暂无公告' }
      ],
      currencynews: [
        { title: '暂无公告' }
      ],
      othersnews: [
        { title: '上币申请公告' }
      ]
    }
  },
  computed: {
    news () {
      const importantnews = this.$store.state.news.importantnews
      const activitynews = this.$store.state.news.activitynews
      const currencyNews = this.$store.state.news.currencyNews
      const otherNews = this.$store.state.news.otherNews
      return {
        importantnews,
        activitynews,
        currencyNews,
        otherNews
      }
    }
  },
  methods: {
    newsContent (id) { // 进入新闻详情页面
      let content = this.$store.state.news.newsContent
      if (content) {
        if (content.hasOwnProperty(`${id}`)) {
          return this.$router.push({
            name: 'newsdetail',
            params: {
              id: id
            }
          })
        }
      }
      this.$socket.sendObj({
        namespace: 'news',
        action: 'onenews',
        data: {
          id: id
        }
      })
      this.$router.push({
        name: 'newsdetail',
        params: {
          id: id
        }
      })
    }
  },
  beforeRouteEnter (to, from, next) { // 组件前置守卫
    Vue.prototype.$socket.sendObj({
      namespace: 'news',
      action: 'newsList',
      data: ''
    })
    next()
  }
}
</script>

<style scoped>
  @import '../../assets/css/News.css';
</style>
