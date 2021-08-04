<template>
  <div class="content-wrapper">
    <!-- 弹出框 -->
    <div
      class="modal fade bs-example-modal-sm"
      tabindex="-1"
      role="dialog"
      aria-labelledby="mySmallModalLabel"
      id="myModal"
    >
      <div class="modal-dialog modal-sm" role="document">
        <div class="modal-content">
          <div class="modal-header">提示</div>
          <div class="modal-body" id="text">用户名或密码不能为空！</div>
        </div>
      </div>
    </div>
    <!-- 登录界面 -->
    <div class="panel panel-primary">
      <div class="panel-heading">用户登录</div>
      <div class="panel-body">
        <form>
          <div>
            <label for="userName">账号:</label>
            <input id="userName" type="text" placeholder="请输入用户名" />
          </div>
          <div>
            <label for="password">密码:</label>
            <input id="password" type="password" placeholder="请输入密码" />
          </div>
          <button type="button" @click="login" class="btn btn-primary">登录</button>
        </form>
      </div>
    </div>
  </div>
</template>
<script>
import $ from "jquery";
//导入路由
import router from "../router.js";
import DelTransaction from "../components/DelTransaction.vue";
import ListOrders from "../components/ListOrders.vue";
import TransactionState from "../components/TransactionStation.vue";
import NewsData from "../components/NewsData.vue";
import NewsContent from "../components/NewsContent.vue";
import UserManage from "../components/UserManage.vue";

export default {
  data() {
    return {};
  },
  mounted() {
    // console.log('login')
  },
  methods: {
    login() {
      // console.log(this.axios)
      let userName = document.querySelector("#userName").value;
      let password = document.querySelector("#password").value;
      //检测输入是否为空
      if (userName === "" || password === "") {
        $("#text").text("用户名或密码不能为空!");
        $("#myModal").modal("show");
      } else {
        // 更改路由情况，cookie后面再加，修改权限等级
        // var xhr = null;
        // if (window.XMLHttpRequest) {
        //   xhr = new XMLHttpRequest();
        // } else {
        //   // 兼容 IE5/6
        //   xhr = new ActiveXObject("Microsoft.XMLHTTP");
        // }
        // xhr.withCredentials = true;
        // xhr.open("post", "http://47.103.108.61:9090/api/admin/login", true);
        // //设置请求头
        // xhr.setRequestHeader(
        //   "Content-type",
        //   "application/x-www-form-urlencoded"
        // );
        // //  提交给服务器，并且设置请求参数。
        // xhr.send({
        //   name: userName,
        //   pass_word: password
        // });
        // xhr.onreadystatechange = function() {
        //   if (4 == xhr.readyState && 200 == xhr.status) {
        //     console.log(xhr.responseText);
        //   } else {
        //     console.log(xhr.responseText);
        //   }
        // };

        // this.$ajax.post('http://47.103.108.61:9090/api/admin/login',{
        //     name: userName,
        //     pass_word: password
        //   },{
        //   headers: {'Access-Control-Allow-Origin': '*'}
        // })
        // .then(res => { 
        //   console.log(res); 
        // }) 
        // .catch(err => { 
        //   console.log(err); 
        // })
        let that = this
        $.ajax({
          type:"post",
          url:"http://47.103.108.61:9090/api/admin/login",
          data:JSON.stringify({ //JSON.stringify变成字符串之后才可以，直接用json对象后台接收不到参数
            name: userName,
            pass_word: password
          }),
          dataType:'json',
          crossDomain: true,  //跨域,
          success:function(response){
            let res = response;
            // console.log(res)
            if (res.data.success === true && res.error_code === 0) {
              //登录成功才能使用路由
              //判断权限
              if (res.data.permission === 1) {
                //超级管理员
                sessionStorage.setItem("level", "super");
                // 登录成功加入路由
                that.$router.addRoutes([
                  { path: "/delTransaction", component: DelTransaction },
                  { path: "/newsContent", component: NewsContent },
                  { path: "/listOrders", component: ListOrders },
                  { path: "/transactionState", component: TransactionState },
                  { path: "/newsData", component: NewsData },
                  { path: "/userManage", component: UserManage }
                ]);
                that.$store.commit('changeToken',res.data.token)
                that.super();
              } else {
                //普通管理员
                sessionStorage.setItem("level", "common");
                that.$router.addRoutes([
                  { path: "/delTransaction", component: DelTransaction },
                  { path: "/newsContent", component: NewsContent },
                  { path: "/listOrders", component: ListOrders },
                  { path: "/transactionState", component: TransactionState },
                  { path: "/newsData", component: NewsData }
                ]);
                that.common();
                that.$store.commit('changeToken',res.data.token)
              }
              sessionStorage.setItem("user", userName);
              sessionStorage.setItem("token", res.data.token);
              that.$router.replace({ path: "/transactionState" }); //跳转到首页
            } else {
              //登录失败提示
              that.logInFailed();
            }
          },
          error:function(error){
            // console.log(error);
            that.$message({
              type: "error",
              message: "网络错误！"
            });
          }
        });
        // this.axios
        //   .post("http://47.103.108.61:9090/api/admin/login", JSON.stringify({
        //     name: userName,
        //     pass_word: password
        //   }))
        //   .then(response => {
        //     console.log(response);
        //     let res = response.data;
        //     // console.log(response)
        //     if (res.data.success === true && res.error_code === 0) {
        //       //登录成功才能使用路由
        //       //判断权限
        //       if (res.data.permission === 1) {
        //         //超级管理员
        //         sessionStorage.setItem("level", "super");
        //         // 登录成功加入路由
        //         this.$router.addRoutes([
        //           { path: "/delTransaction", component: DelTransaction },
        //           { path: "/newsContent", component: NewsContent },
        //           { path: "/listOrders", component: ListOrders },
        //           { path: "/transactionState", component: TransactionState },
        //           { path: "/newsData", component: NewsData },
        //           { path: "/userManage", component: UserManage }
        //         ]);
        //         this.super();
        //       } else {
        //         //普通管理员
        //         sessionStorage.setItem("level", "common");
        //         this.$router.addRoutes([
        //           { path: "/delTransaction", component: DelTransaction },
        //           { path: "/newsContent", component: NewsContent },
        //           { path: "/listOrders", component: ListOrders },
        //           { path: "/transactionState", component: TransactionState },
        //           { path: "/newsData", component: NewsData }
        //         ]);
        //         this.common();
        //       }
        //       sessionStorage.setItem("user", userName);
        //       this.$router.replace({ path: "/transactionState" }); //跳转到首页
        //     } else {
        //       //登录失败提示
        //       this.logInFailed();
        //     }
        //   })
        //   .catch(error => {
        //     console.log(error);
        //     this.$message({
        //       type: "error",
        //       message: "网络错误！"
        //     });
        //   });
      }
    },
    /**
     * 得出用户等级为超级管理员后的操作
     */
    super() {
      this.$store.commit("changeLevel", "super"); //更改vuex中的等级
    },
    /**
     *
     */
    common() {
      this.$store.commit("changeLevel", "common");
    },
    /**
     * 登录成功提示
     */
    logInSuccess() {
      $("#text").text("登录成功!");
      $("#myModal").modal("show");
    },
    /**
     * 登录失败提示
     */
    logInFailed() {
      $("#text").text("用户名或密码错误，请检查输入!");
      $("#myModal").modal("show");
    },
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

      if (cs.hasOwnProperty(name)) {
        return cs[name];
      } else {
        return null;
      }
    }
  }
};
</script>
<style scoped>
.panel-body {
  height: 400px;
  display: flex;
  align-items: center;
}
.panel-body > form {
  font-size: 18px;
  width: auto;
  height: auto;
  margin: 0 auto;
  text-align: center;
  /* background-color: peachpuff; */
}
form button {
  margin: 30px auto;
}
</style>