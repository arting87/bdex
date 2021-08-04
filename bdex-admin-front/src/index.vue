<!-- index.vue -->
<template>
  <div class="wrapper">
    <header class="main-header">
      <!-- Logo -->
      <a href="javascript:;" class="logo">
        <!-- mini logo for sidebar mini 50x50 pixels -->
        <span class="logo-mini">A</span>
        <!-- logo for regular state and mobile devices -->
        <span class="logo-lg"><b>Admin</b>Lte</span>
      </a>
      
      <!-- Header Navbar: style can be found in header.less -->
      <nav class="navbar navbar-static-top">
        <!-- Sidebar toggle button-->
        <a href="#" class="sidebar-toggle" data-toggle="offcanvas" role="button">
          <span class="sr-only">Toggle navigation</span>
        </a>
        <!-- 导航栏右边部分 -->
        <div class="navbar-custom-menu">
          <ul class="nav navbar-nav">
            <li class="dropdown user user-menu">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                <img src="./dist/img/user2-160x160.jpg" class="user-image" alt="User Image">
                <span class="hidden-xs">{{userName}}</span>
              </a>
              <ul class="dropdown-menu">
                <!-- User image -->
                <li class="user-header">
                  <img src="./dist/img/user2-160x160.jpg" class="img-circle" alt="User Image">
                  <p>
                    {{userName}}
                  </p>
                </li>
                <li class="user-footer">
                  <div class="pull-left">
                    <a href="javascript:;" id="userManage" @click="userManage()" class="btn btn-default btn-flat">用户管理</a>
                  </div>
                  <div class="pull-right">
                    <a href="#" class="btn btn-default btn-flat" @click="logInOrLogOut()">{{loginStateRight}}</a>
                  </div>
                </li>
              </ul>
            </li>
          </ul>
        </div>
      </nav>
    </header>

    <!-- 左侧边栏 -->
    <aside class="main-sidebar">
      <!-- 左侧边栏: style can be found in sidebar.less -->
      <section class="sidebar">

        <!-- 用户板块 -->
        <div class="user-panel">
          <div class="pull-left image">
            <img src="./dist/img/user2-160x160.jpg" class="img-circle" alt="User Image">
          </div>
          <div class="pull-left info">
            <!-- 左上角板块管理员级别 -->
            <p>{{userLevel}}</p>
            <!-- 左上角板块用户登录状态 -->
            <a href="#"><i class="fa fa-circle text-success"></i>{{loginState}}</a>
          </div>
        </div>

       

        <!-- 侧边栏功能菜单 -->
        <ul class="sidebar-menu" id="menuList">
          <li class="header">功能导航</li>
          <li class="treeview">
            <a href="#">
              <i class="glyphicon glyphicon-tower"></i> <span>交易所</span>
              <span class="pull-right-container">
                <i class="fa fa-angle-left pull-right"></i>
                <!-- <small class="label pull-right bg-green">new</small> -->
              </span>
            </a>
            <ul class="treeview-menu">
              <li><router-link to="transactionState"><i class="fa fa-circle-o"></i>交易所状态</router-link></li>
            </ul>
          </li>

          <li class="treeview">
            <a href="#">
              <i class="fa fa-pie-chart"></i>
              <span>token交易</span>
              <span class="pull-right-container">
                <i class="fa fa-angle-left pull-right"></i>
                <!-- <small class="label pull-right bg-green">new</small> -->
              </span>
            </a>
            <ul class="treeview-menu">
              <li><router-link to="delTransaction"><i class="fa fa-circle-o"></i>交易查询与修改</router-link></li>
            </ul>
          </li>

          <li class="treeview">
            <a href="#">
              <i class="glyphicon glyphicon-list-alt"></i> <span>订单</span>
              <span class="pull-right-container">
                <i class="fa fa-angle-left pull-right"></i>
                <!-- <small class="label pull-right bg-green">new</small> -->
              </span>
            </a>
            <ul class="treeview-menu">
              <li><router-link to="listOrders"><i class="fa fa-circle-o"></i>分页查询所有订单</router-link></li>
            </ul>
          </li>

          <li class="treeview">
            <a href="#">
              <i class="glyphicon glyphicon-edit"></i> <span>新闻</span>
              <span class="pull-right-container">
                <i class="fa fa-angle-left pull-right"></i>
                <!-- <small class="label pull-right bg-green">new</small> -->
              </span>
            </a>
            <ul class="treeview-menu">
              <li><router-link to="newsData"><i class="fa fa-circle-o"></i>新闻数据处理</router-link></li>
            </ul>
          </li>
        </ul>
      </section>
      <!-- /侧边栏 -->
    </aside>
    <router-view></router-view>
    <div class="control-sidebar-bg"></div>
  </div>
</template>
<script>
import router from './router.js'
import DelTransaction from './components/DelTransaction.vue'
import ListOrders from './components/ListOrders.vue'
import TransactionState from './components/TransactionStation.vue'
import NewsData from './components/NewsData.vue'
import NewsContent from './components/NewsContent.vue'
import UserManage from './components/UserManage.vue'
export default {
  data() {
    return {
      userLevel: '游客',
      userName: '请登录',
      loginState: '请登录',
      loginStateRight: '登录'
    }
  },
  mounted(){
    this.multiDevice()
    this.refresh()
    this.level()
  },
  computed: {
    storeUserLevel () {
      return this.$store.state.login.userLevel
    }
  },
  watch: {
    storeUserLevel () {
      this.level()
    }
  },
  methods: {
    /**
     * 刷新处理
     */
    refresh () {
      //用户刷新后跳转到当前页面
      if(sessionStorage.getItem('level') === 'super') {
        this.$store.commit('changeLevel','super') //更改vuex中的等级
        this.$router.addRoutes([
            { path: '/delTransaction',component: DelTransaction},
            { path: '/newsContent',component: NewsContent},
            { path: '/listOrders',component: ListOrders},
            { path: '/transactionState',component: TransactionState},
            { path: '/newsData',component: NewsData},
            { path: '/userManage', component:UserManage}
        ])
        switch (sessionStorage.getItem('currentPage')) {
          case '1':
            this.$router.push({ path: '/delTransaction'})
            break
          case '2':
            this.$router.push({ path: '/listOrders'})
            break
          case '3':
            this.$router.push({ path: '/transactionState'})
            break
          case '4':
            this.$router.push({ path: '/newsData'})
            break
          case '5':
            this.$router.push({ path: '/userManage'})
          case '6':
            this.$router.push({ path: '/newsContent'})
            break
          default:
            this.$router.replace({ path: '/login'})
            break
        }
      } else if (sessionStorage.getItem('level') === 'common') {
        this.$store.commit('changeLevel','common') //更改vuex中的等级
        this.$router.addRoutes([
            { path: '/delTransaction',component: DelTransaction},
            { path: '/newsContent',component: NewsContent},
            { path: '/listOrders',component: ListOrders},
            { path: '/transactionState',component: TransactionState},
            { path: '/newsData',component: NewsData}
        ])
        switch (sessionStorage.getItem('currentPage')) {
          case '1':
            this.$router.replace({ path: '/delTransaction'})
            break
          case '2':
            this.$router.replace({ path: '/listOrders'})
            break
          case '3':
            this.$router.replace({ path: '/transactionState'})
            break
          case '4':
            this.$router.replace({ path: '/newsData'})
            break
          case '6':
            this.$router.push({ path: '/newsContent'})
            break
          default:
            this.$router.replace({ path: '/login'})
            break
        }
      }
      if(sessionStorage.getItem( "token") !== null) {
        this.$store.commit( "changeToken", sessionStorage.getItem( "token"))
      }
    },
    /**
     * 管理权限及点击
     * 刚进入时禁用导航栏点击功能，然后判断是否是管理员或超级管理员，是的话就开放点击功能
     */
    level () {
      //禁用点击
      let lis = document.querySelectorAll('#menuList>li')
      for(let i = 0; i<lis.length;i++){
        lis[i].setAttribute('style','pointer-events:none;')
      }

      let userManage = document.querySelector('#userManage')
      userManage.setAttribute('style','pointer-events:none;')

      //判断权限等级
      if( this.storeUserLevel === 'common' || this.storeUserLevel === 'super') { //登录成功
        //开启导航菜单点击功能
        for(let i = 0; i<lis.length;i++){
          lis[i].setAttribute('style','pointer-events:auto;')
        }

        //修改用户板块显示情况
        if(this.storeUserLevel === 'common') {
          this.userLevel = '管理员'
        }else if (this.storeUserLevel === 'super') {
          this.userLevel = '超级管理员'
          userManage.setAttribute('style','pointer-events:auto;')
        }
        this.loginState = '在线',
        this.loginStateRight = '退出登录'
        this.userName = sessionStorage.getItem('user')
      }
    },
    /**
     * 点击登录或退出登录
     */
    logInOrLogOut () {
      this.$store.commit('changeLevel','') //恢复等级为空，即游客
      this.userLevel = '游客',
      this.userName = '请登录',
      this.loginState = '未登录',
      this.loginStateRight = '登录'
      sessionStorage.removeItem('user')
      sessionStorage.removeItem('currentPage')
      sessionStorage.removeItem('level')
      sessionStorage.removeItem('token')
      this.$router.push('/login')
    },
    /**
     * 用户权限点击
     */
    userManage () {
      // if()
      this.$router.push({path:'/userManage'})
    },
    /**
     * 设备兼容
     */
    multiDevice() {
      var sUserAgent = navigator.userAgent.toLowerCase();
      var bIsIpad = sUserAgent.match(/ipad/i) == "ipad";
      var bIsIphoneOs = sUserAgent.match(/iphone os/i) == "iphone os";
      var bIsMidp = sUserAgent.match(/midp/i) == "midp";
      var bIsUc7 = sUserAgent.match(/rv:1.2.3.4/i) == "rv:1.2.3.4";
      var bIsUc = sUserAgent.match(/ucweb/i) == "ucweb";
      var bIsAndroid = sUserAgent.match(/android/i) == "android";
      var bIsCE = sUserAgent.match(/windows ce/i) == "windows ce";
      var bIsWM = sUserAgent.match(/windows mobile/i) == "windows mobile";
      if (bIsIpad || bIsIphoneOs || bIsMidp || bIsUc7 || bIsUc || bIsAndroid || bIsCE || bIsWM) {
        alert(1);
          $("a[href!='#']").attr("data-toggle","offcanvas");
      }else{
        // alert(2);
      }
    }
  }
}
</script>
<style scoped>
#menuList{
  color: gray;
}
.wrapper {
  height: 100%;
}
</style>