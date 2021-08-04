<template>
  <div class="content-wrapper">
    <!-- 头部 面包屑路径导航 -->
    <section class="content-header">
      <h1>
        BDEX后台管理系统
        <!-- <small>Version 2.0</small> -->
      </h1>
      <ol class="breadcrumb">
        <li>
          <a href="#">
            <i class="glyphicon glyphicon-edit"></i>新闻
          </a>
        </li>
        <li class="active">新闻数据处理>详情</li>
      </ol>
    </section>

    <section class="content">
      <div class="row">
        <div class="col-xs-12">
          <div class="box">
            <div class="box-header with-border" id="title"></div>
            <div class="box-body" id="content"></div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
export default {
  data() {
    return {};
  },
  mounted() {
    sessionStorage.setItem("currentPage", 6);
    let that = this;
    $.ajax({
      type:"post",
      url:"http://47.103.108.61:9090/api/admin/getNews",
      data:{},
      dataType:'json',
      headers: {
        "Authorization": "Bearer " + that.$store.state.login.token
      },
      success:function(response){
        let res = response
        // console.log(res)
        if (res.error_code === 0) {
          that.items = res.data;
          let id = that.$store.state.login.data;
          let contentData = "";
          for (let i = 0; i < that.items.length; i++) {
            if (that.items[i].id === id) {
              contentData = that.items[i];
              break;
            }
          }
          // console.log(contentData)
          let title = document.querySelector("#title");
          let content = document.querySelector("#content");
          title.innerHTML = contentData.title;
          content.innerHTML = contentData.content;
          // console.log(that .items)
        } else {
          that.displayInfo("信息获取失败！");
        }
      },
      error:function(xhr, status, error){
        that.$message({
          type: "error",
          message: "网络错误！"
        });
      }
    })
  },
  methods: {
    /**
     * 设置cookie
     * @param {String} name  cookie名称
     * @param {String} value cookie值
     * @param {Number} time 过期时间 分钟
     */
    setCookie(name, value, minute) {
      let va = escape(value);
      let data = new Date();
      if (minute) {
        data.setTime(data.getTime() + 1000 * 60 * minute);
        document.cookie = `${name}=${va};expires=${data.toUTCString()}`;
      } else {
        document.cookie = `${name}=${va}`; //当前会话关闭就超时
      }
    },
    /**
     * 获取所有cookie
     * return 包含当前cookie值的对象
     */
    getCookieAll() {
      let coo = document.cookie;
      if (coo == "") {
        return null;
      }
      //清除空格
      coo = coo.replace(/\s+/g, "");
      let cArr = coo.split(";");
      let cookies = {};
      cArr.forEach(coo => {
        let c = coo.split("=");
        cookies[c[0]] = unescape(c[1]);
      });
      return cookies;
    },
    /**
     * 获取指定名称的cookie
     * 参数name: cookie名称
     * return :如果有返回cookie值，否则返回为null
     */
    getCookieByName(name) {
      let cs = this.getCookieAll();
      if (cs !== null) {
        if (cs.hasOwnProperty(name)) {
          return cs[name];
        } else {
          return null;
        }
      }
    },
    /**
     * 提示
     */
    displayInfo(text) {
      let errData = this.$store.state.login.token;
      let code = JSON.parse(window.atob(errData));
      let timeOver = code.time_stamp + 3600; //3600
      let timestamp = Date.parse(new Date()) / 1000;
      let res = timeOver - timestamp;
      if (res < 0) {
        this.$message({
          type: "warn",
          message: "您的登录状态已过期，请重新登录!"
        });
      } else {
        if (text.slice(text.length - 3, text.length - 1) === "成功") {
          this.$message({
            type: "success",
            message: text
          });
        } else {
          this.$message({
            type: "warn",
            message: text
          });
        }
      }
    }
  }
};
</script>
<style scoped>
#title {
  text-align: center;
  font-size: 48px;
}
</style>