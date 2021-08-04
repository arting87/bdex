<template>
  <div class="newsdetail">
    <span>
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/news' }">公告首页</el-breadcrumb-item>
        <el-breadcrumb-item>公告详情</el-breadcrumb-item>
      </el-breadcrumb>
    </span>
    <h1 class="title">{{ newsHtml === null ? this.$router.push({name: 'NotFound'}) : newsHtml.title}}</h1>
    <div v-html="newsHtml.content"></div>
  </div>
</template>

<script>
import store from '../../store/index'
export default {
  name: 'NewsDetail',
  data () {
    return {
      newstitle: '上币申请公告'
      // rawHtml: `<div class="write"><h5> <a name="header-n41" class="md-header-anchor"></a>亲爱的用户： </h5> <p class="indent"> <span></span>BDEX将在2019年7月1日18:00正式开放交易，首期上线ETH/EOS跨链交易，随后陆续开放ETH与EOS链内交易。 </p> <ul> <p>BDEX用户操作手册如下：</p> <ul> <li> ETH交易： <a href="https://www.bdex.global/cn/nd.jsp?id=32#_jcp=4_9" target="_blank" class="url" >https://www.bdex.global/cn/nd.jsp?id=32#_jcp=4_9</a> </li> <li> EOS交易： <a href="https://www.bdex.global/cn/nd.jsp?id=31#_jcp=4_9" target="_blank" class="url" >https://www.bdex.global/cn/nd.jsp?id=31#_jcp=4_9</a> </li> <li> 跨链交易： <a href="https://www.bdex.global/cn/nd.jsp?id=33#_jcp=4_9" target="_blank" class="url" >https://www.bdex.global/cn/nd.jsp?id=33#_jcp=4_9</a> <span></span> </li> </ul> </ul> <p class="indent">如果您在使用的过程中，遇到任何问题，可通过官方邮箱联系我们。</p> <blockquote> <p> BDEX官方邮箱： <a href="mailto:business@bdex.me">business@bdex.me</a> </p> </blockquote> <p class="indent">您的宝贵建议将促使我们不断完善BDEX，力争发展成为全球一流的去中心化跨链交易所。</p></div>`,
      // raw1Html: `<div class="write"><h5>亲爱的用户：</h5> <p> 为了在bdex拥有更好的交易体验，我们开放自助上币通道，以下是上币和币种下线说明。 </p> <p>&nbsp;</p> <h5> <a name="header-n5" class="md-header-anchor"></a>上线币种说明 </h5> <h5>亲爱的用户：</h5> <p> 为了保证用户交易体验、保护投资者利益，我们会严格评估项目的投资潜力和风险，考察内容如下： </p> <ul> <li>团队实力</li> <li>社区维护</li> <li>技术实力</li> <li>项目发展</li> <li>上线合作</li> </ul> <p> 请在线填写相关信息发送至社区邮箱： <a href="business@bdex.me" target="_blank" class="url">business@bdex.me</a> </p> <p>&nbsp;</p> <h5> <a name="header-n9" class="md-header-anchor"></a>币种下线说明 </h5> <h5>亲爱的用户：</h5> <p> bdex作为优质数字资产交易平台，拥有严选优质项目上线交易的权利，为了保护投资者利益，也天然保留无偿下线项目的权利。以下情形，将会提高项目下线的可能性： </p> <ul> <li>交易量持续低迷</li> <li>项目停滞不前</li> <li>团队投入度不够</li> <li>申请或沟通的信息造假</li> <li>重大负面新闻</li> <li>项目方主动要求下线</li> <li>其他不支持继续交易的事项</li> </ul> <p> 对于确定下线的项目，我们会提前5天发出公告，用户将有30天时间从交易所提现该资产。 </p> </div>`
    }
  },
  computed: {
    newsHtml () {
      let id = this.$route.params.id
      let newsContent = this.$store.state.news.newsContent
      if (newsContent !== null) {
        if (newsContent.hasOwnProperty(id.toString())) {
          return newsContent[`${id}`]
        }
      }
      return null // 处理获取新闻详情页面数据失败的情况, 跳转404页面
    }
  },
  beforeRouteEnter (to, from, next) {
    let newsContent = store.state.news.newsContent
    let id = to.params.id
    if (newsContent) {
      if (newsContent.hasOwnProperty(`${id}`)) {
        next()
      }
    }
    let timeoutId = ''
    let intervalId = setInterval(() => {
      if (newsContent) {
        if (newsContent.hasOwnProperty(`${id}`)) {
          clearInterval(intervalId)
          clearTimeout(timeoutId)
          next()
        }
      }
    }, 300)
    timeoutId = setTimeout(() => {
      clearInterval(intervalId)
      clearTimeout(timeoutId)
      next()
    }, 3000)
  }
}
</script>

<style>
.newsdetail {
  padding: 40px;
  background-color: #ffffff;
}
.write {
  bottom: 80px;
  max-width: 1000px;
  height: 100px;
  margin: 0 auto;
  padding: 0px 40px;
  border: 1px solid #ffffff;
  background-color: #ffffff;
  color: #000000;
}
.indent {
  text-indent: 2.5em;
}
</style>
<style scoped>
.title {
  text-align: center;
}
.el-button {
  margin-top: 20px;
  font-size: 16px;
}
.el-breadcrumb {
  margin-left: 15%;
}
</style>
