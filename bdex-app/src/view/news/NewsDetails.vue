<template>
<div class="newsdetails">
    <navbar text="通知公告" type="bright" />
  <div class="notice">
      <h1 class="title">
        <a name="header-n0" class="md-header-anchor"></a>上币申请公告
      </h1>
      <h5>亲爱的用户：</h5>
      <p>
        为了在bdex拥有更好的交易体验，我们开放自助上币通道，以下是上币和币种下线说明。
      </p>
      <h5>
        <a name="header-n5" class="md-header-anchor"></a>上线币种说明
      </h5>
      <h5>亲爱的用户：</h5>
      <p>
        为了保证用户交易体验、保护投资者利益，我们会严格评估项目的投资潜力和风险，考察内容如下：
      </p>
      <ul>
        <li>团队实力</li>
        <li>社区维护</li>
        <li>技术实力</li>
        <li>项目发展</li>
        <li>上线合作</li>
      </ul>
      <p>
        请在线填写相关信息发送至社区邮箱：
        <a href="business@bdex.me" target="_blank" class="url">business@bdex.me</a>
      </p>
      <h5>
        <a name="header-n9" class="md-header-anchor"></a>币种下线说明
      </h5>
      <h5>亲爱的用户：</h5>
      <p>
        bdex作为优质数字资产交易平台，拥有严选优质项目上线交易的权利，为了保护投资者利益，也天然保留无偿下线项目的权利。以下情形，将会提高项目下线的可能性：
      </p>
      <ul>
        <li>交易量持续低迷</li>
        <li>项目停滞不前</li>
        <li>团队投入度不够</li>
        <li>申请或沟通的信息造假</li>
        <li>重大负面新闻</li>
        <li>项目方主动要求下线</li>
        <li>其他不支持继续交易的事项</li>
    </ul>
        <p>
        对于确定下线的项目，我们会提前5天发出公告，用户将有30天时间从交易所提现该资产。
        </p>
  </div>
  <div @click="$router.push({path: '/index/news'})">查看更多公告</div>
</div>
</template>

<script>
import Navbar from '../../components/Navbar'

export default {
  name: 'NewsDetails',
  data () {
    return {
    }
  },
  components: {
    Navbar
  },
  methods: {
  }
}
</script>

<style lang="scss" scoped>
.newsdetails {
  width: 100%;
  height: 110vh;
  background-color: #ffffff;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.notice {
  margin: 0 auto;
  padding: 0px 40px;
  border: 1px solid #ffffff;
  background-color: #ffffff;
  color: #000000;
}
.notice p {
  text-indent: 2.5em;
}
.title {
  text-align: center;
}
</style>
