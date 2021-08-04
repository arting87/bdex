<!-- delTransaction.vue -->
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
    <!-- 头部 面包屑路径导航 -->
    <!-- <content-header v-bind:menuList="menuList"></content-header> -->
    <section class="content-header">
      <h1>
        交易修改与查询
        <!-- <small>Version 2.0</small> -->
      </h1>
      <ol class="breadcrumb">
        <li>
          <a href="#">
            <i class="fa fa-pie-chart"></i>token交易
          </a>
        </li>
        <li class="active">交易查询与修改</li>
      </ol>
    </section>
    <!-- Main content -->
    <section class="content">
      <div class="row">
        <div class="col-xs-12">
          <div class="box">
            <!-- 查询代币条件部分 -->
            <div class="box-header with-border">
              <h3>查询</h3>
            </div>
            <div class="box-body">
              <table class="table table-bordered">
                <thead>
                  <tr>
                    <th>基本交易对的类型</th>
                    <th>代币地址</th>
                    <th>代币名称</th>
                    <th>代币精度</th>
                    <th>偏移量</th>
                    <th>取的个数</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody>
                  <tr id="send_search">
                    <td>
                      <input type="number" placeholder="int" />
                    </td>
                    <td>
                      <input type="text" placeholder="string" />
                    </td>
                    <td>
                      <input type="text" placeholder="string" />
                    </td>
                    <td>
                      <input type="number" placeholder="int" />
                    </td>
                    <td>
                      <input type="number" placeholder="int" />
                    </td>
                    <td>
                      <input type="number" placeholder="int" />
                    </td>
                    <td>
                      <button
                        class="btn btn-primary btn-sm"
                        id="search"
                        @click="send_SearchToken()"
                      >查询</button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>

            <!-- 交易列表 -->
            <div class="box-header with-border">
              <h3>删除</h3>
            </div>
            <div class="box-body">
              <table class="table table-bordered">
                <thead>
                  <tr>
                    <th>在数据库中唯一id</th>
                    <th>所在链上id</th>
                    <th>名称</th>
                    <th>地址</th>
                    <th>类型</th>
                    <th>精度</th>
                    <th>数据库代币处理精度</th>
                    <th>true是取消了的代币，false是上了的代币</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody id="delList">
                  <tr v-for="(item,index) in transactionList.pageData" :key="index">
                    <td>{{item.id}}</td>
                    <td>{{item.token_id}}</td>
                    <td>{{item.symbol}}</td>
                    <td>{{item.token_contract}}</td>
                    <td>{{item.market_code}}</td>
                    <td>{{item.decimal}}</td>
                    <td>{{item.order_decimal}}</td>
                    <td>{{item.is_cancel}}</td>
                    <td>
                      <button class="btn btn-primary btn-sm">删除</button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            <!-- 列表分页按钮部分 -->
            <el-pagination
              background
              layout="prev, pager, next"
              @prev-click="pageMinus"
              @next-click="pagePlus"
              @current-change="orderCurrentPage"
              :page-size="transactionList.numberOfItems"
              :hide-on-single-page="true"
              :total="items.length"
            ></el-pagination>
            <!-- 添加交易部分 -->
            <div class="box-header with-border">
              <h3>添加</h3>
            </div>
            <div class="box-body">
              <table class="table table-bordered">
                <thead>
                  <tr>
                    <th>代币地址</th>
                    <th>基本交易对的类型</th>
                    <th>代币符号</th>
                    <th>代币精度</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody>
                  <tr id="addList">
                    <td>
                      <input type="text" required placeholder="代币地址由13位及以内长度的小写字母组成" />
                    </td>
                    <td>
                      <input type="number" required placeholder="只能输入1-5之间的整数" />
                    </td>
                    <td>
                      <input type="text" required placeholder="代币符号必须为全大写字母" />
                    </td>
                    <td>
                      <input type="number" required placeholder="0~20" />
                    </td>
                    <td>
                      <button
                        type="button"
                        id="add"
                        @click="addTransaction()"
                        class="btn btn-primary btn-sm"
                      >添加</button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  data() {
    return {
      items: [],
      transactionList: {
        pageData: [], //当前页的数据
        numberOfItems: 10, //一页展示的数量
        totalNumberOfPages: "",
        currentPage: 1 //当前页数
      }
    };
  },
  created() {
    //获取
    this.init();
  },
  mounted() {
    sessionStorage.setItem("currentPage", 1);
    this.testLevel(); //渲染后查看权限，进行界面设置
    this.send_SearchToken(); //查询代币
    this.deleteTransactionPair(); //删除按钮的事件逻辑，发送删除数据
    this.addTransaction(); //添加按钮的事件逻辑，发送添加数据
    let search = document.querySelector("#search");
    search.onclick = () => {
      this.send_SearchToken();
    };
  },
  watch: {
    /**
     * 1. 检测总数据
     * 2. 数据处理，分页
     * 3. 分页数据展示
     * 4. 展示的数据对应的分页按钮加active类
     * 提示：这里的检测items是为了处理 用户更换查询条件后再次查询 返回结果改变 产生的数据列表变化 列表要重新渲染 页面回到第一页
     */
    items() {
      this.init(); //展示第一页数据
      this.orderCurrentPage(1); //页面回到第一页
    }
  },
  updated() {
    this.testLevel();
  },
  methods: {
    /**
     * 检测用户权限等级
     */
    testLevel() {
      //检测普通用户及其他情况下都禁用修改权限，只有为超级管理员才能修改
      let deleteBtn = document.querySelectorAll("#delList button");
      let add = document.querySelector("#add");

      for (let i = 0; i < deleteBtn.length; i++) {
        deleteBtn[i].setAttribute("disabled", "disabled");
      }
      add.setAttribute("disabled", "disabled");

      //判断是否为超级管理员
      if (this.$store.state.login.userLevel === "super") {
        for (let i = 0; i < deleteBtn.length; i++) {
          deleteBtn[i].removeAttribute("disabled");
        }
        add.removeAttribute("disabled");
      }
    },
    /**
     * 功能：
     * 1.初始化页数
     * 2.展示第一页数据
     */
    init() {
      if (this.items !== [] && this.items !== null) {
        this.transactionList.pageData = this.items.slice(
          0,
          this.transactionList.numberOfItems
        ); //第一页的数据
        this.transactionList.currentPage = 1; //当前页数设为第一页
        this.transactionList.totalNumberOfPages = Math.ceil(
          this.items.length / this.transactionList.numberOfItems
        ); //总页数
        // console.log(this.transactionList.pageData)
      }
    },
    /**
     * 展示页面数据
     * 参数：页面序号
     */
    orderCurrentPage(val) {
      this.transactionList.currentPage = val; // 修改当前页数

      //检测是否为最后一页
      if (this.transactionList.totalNumberOfPages === val) {
        //当前页为最后一页
        this.transactionList.pageData = this.items.slice(
          (val - 1) * this.transactionList.numberOfItems,
          this.items.length
        ); //展示的页数为当前页数所对应的数据
      } else {
        //非最后一页
        this.transactionList.pageData = this.items.slice(
          (val - 1) * this.transactionList.numberOfItems,
          val * this.transactionList.numberOfItems
        ); //展示的页数为当前页数所对应的数据
      }
    },
    /**
     * 页数减
     */
    pageMinus() {
      if (this.transactionList.currentPage > 1) {
        this.transactionList.currentPage = this.transactionList.currentPage - 1;
        this.orderCurrentPage(this.transactionList.currentPage);
      }
    },
    /**
     * 页数加
     */
    pagePlus() {
      if (
        this.transactionList.currentPage <
        this.transactionList.totalNumberOfPages
      ) {
        this.transactionList.currentPage = this.transactionList.currentPage + 1;
        this.orderCurrentPage(this.transactionList.currentPage);
      }
    },
    /**
     * 功能
     * 发送查询代币请求
     */
    send_SearchToken() {
      //获取输入的值并加入到对象里
      let send_search_list = document.querySelectorAll("#send_search input");
      let list = {};
      if (send_search_list[0].value === "") list.market_code = 0;
      else list.market_code = parseInt(send_search_list[0].value);

      list.token_contract = send_search_list[1].value;
      list.symbol = send_search_list[2].value;

      if (send_search_list[3].value === "") list.decimal = 0;
      else list.decimal = parseInt(send_search_list[3].value);

      if (send_search_list[4].value === "") list.offset = 0;
      else list.offset = parseInt(send_search_list[4].value);

      if (send_search_list[5].value === "") list.limit = 0;
      else list.limit = parseInt(send_search_list[5].value);
      // console.log(list)
      //将数据传到后台并接收查询结果
      let that = this
      $.ajax({
        type:"post",
        url:"http://47.103.108.61:9090/api/admin/getToken",
        data:JSON.stringify(list),
        dataType:'json',
        crossDomain: true,  //跨域,
        headers: {
          "Authorization": "Bearer " + that.$store.state.login.token
        },
        success:function(response){
          let data = response;
          // console.log(data)
          if (data.error_code === 0) {
            that.displayInfo("查询成功!");
            that.items = data.data; //渲染列表数据变化
          } else {
            //报错
            that.displayInfo("查询失败，请重新输入!");
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
    /**
     * 功能
     * 删除指定交易对
     * market_code token_contract
     * 步骤：点击后先获取要删除的对象，获取对象的类型和地址，将数据参数发送给后台
     */
    deleteTransactionPair() {
      let delList = document.querySelector("#delList");
      delList.onclick = e => {
        let target = e.target || event.target;
        if (target.nodeName === "BUTTON") {
          this.$confirm("此操作将永久删除该项, 是否继续?", "提示", {
            confirmButtonText: "确定",
            cancelButtonText: "取消",
            type: "warning"
          }).then(() => {
            let market_code = parseInt(target.parentNode.parentNode.children[4].innerText);
            let token_contract = target.parentNode.parentNode.children[3].innerText;
            if (market_code !== 4 && token_contract !== "0x0000000000000000000000000000000000000000") {
              let delObj = {
                market_code,
                token_contract
              };
              let that = this;
              $.ajax({
                type:"post",
                url:"http://47.103.108.61:9090/api/admin/deleteToken",
                data:JSON.stringify(delObj), //字符串化
                dataType:'json',
                crossDomain: true,  //跨域,
                headers: { //token方式跨域验证
                  "Authorization": "Bearer " + that.$store.state.login.token
                },
                success:function(response){
                  let res = response;
                  // console.log(res)
                  if (res.error_code === 0 && res.data.success === true) {
                    that.displayInfo("删除成功!");
                  } else {
                    that.displayInfo("删除失败!");
                  }
                },
                error:function(xhr, status, error){
                  that.$message({
                    type: "error",
                    message: "网络错误！"
                  });
                }
              })
            } else {
              this.displayInfo("该笔交易不可删除!");
            }
          }).catch(() => {
            this.$message({
              type: "info",
              message: "已取消删除!"
            });
          });
        }
      };
    },
    /**
     * 功能：添加交易对
     * 参数必选
     */
    addTransaction() {
      let addList = document.querySelector("#addList");
      addList.onclick = e => {
        let target = e.target || event.target;
        if (target.nodeName === "BUTTON") {
          let token_contract =
            target.parentNode.parentNode.children[0].children[0].value;
          let market_code = parseInt(
            target.parentNode.parentNode.children[1].children[0].value
          );
          let symbol =
            target.parentNode.parentNode.children[2].children[0].value;
          let decimal = parseInt(
            target.parentNode.parentNode.children[3].children[0].value
          );
          //代币符号与合约地址要做限制判断
          let symbol_test = /[A-Z]/g;
          // console.log(symbol.match(symbol_test))
          //输入检测
          if (
            token_contract === "" ||
            market_code === "" ||
            symbol === "" ||
            decimal === ""
          ) {
            this.displayInfo("输入不能为空，请检查输入!");
          } else if (
            symbol.match(symbol_test) === null ||
            symbol.match(symbol_test).length !== symbol.length
          ) {
            this.displayInfo("添加失败，代币符号只能由大写字母组成");
          } else if (
            (token_contract.length !== 42 && market_code !== 1) ||
            (token_contract.length > 13 && market_code === 1)
          ) {
            this.displayInfo("添加失败,地址长度错误!");
          } else if (decimal < 0 || decimal > 20) {
            this.displayInfo("添加失败，精度错误!");
          } else if (
            market_code !== 1 &&
            market_code !== 2 &&
            market_code !== 3 &&
            market_code !== 4 &&
            market_code !== 5
          ) {
            this.displayInfo("添加失败,交易类型只能为1到5!");
          } else {
            let addObj = {
              token_contract,
              market_code,
              symbol,
              decimal
            };
            let that = this;
            $.ajax({
              type:"post",
              url:"http://47.103.108.61:9090/api/admin/addToken",
              data:JSON.stringify(addObj), //字符串化
              dataType:'json',
              crossDomain: true,  //跨域,
              headers: { //token方式跨域验证
                "Authorization": "Bearer " + that.$store.state.login.token
              },
              success:function(response){
                let res = response;
                // console.log(res)
                if (res.error_code === 0 && res.data.success === true) {
                  that.displayInfo("添加成功!");
                } else {
                  that.displayInfo("添加失败!");
                }
              },
              error:function(xhr, status, error){
                that.$message({
                  type: "error",
                  message: "网络错误！"
                });
              }
            })
          }
        }
      };
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
form#add {
  text-align: center;
  margin: 0 auto;
  padding: 10px;
}
input {
  display: block;
  width: 100%;
}
.content-wrapper {
  height: 100%;
  min-height: 1000px;
}
.box {
  padding-bottom: 100px;
}
</style>