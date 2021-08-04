<!-- ListOrders.vue -->
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
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <h1>
        分页查询所有订单
        <!-- <small>Version 2.0</small> -->
      </h1>
      <ol class="breadcrumb">
        <li>
          <a href="#">
            <i class="glyphicon glyphicon-list-alt"></i>订单
          </a>
        </li>
        <li class="active">分页查询所有订单</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">
      <div class="row">
        <div class="col-xs-12">
          <div class="box">
            <!-- 查询条件 -->
            <div class="box-header with-border">
              <table class="table table-bordered">
                <tr>
                  <th>时间范围</th>
                  <th>订单创建时间/订单更新时间</th>
                  <th>市场类型</th>
                  <th>交易对名称</th>
                  <th>价格区间的开始</th>
                  <th>价格区间的结束</th>
                  <th>可否撤销状态</th>
                  <th>是否完成状态</th>
                  <th>订单状态</th>
                  <th>订单发送者名称</th>
                  <th>操作</th>
                </tr>
                <tr>
                  <td>
                    <el-date-picker
                      v-model="value2"
                      type="datetimerange"
                      :picker-options="pickerOptions"
                      range-separator="至"
                      start-placeholder="开始日期"
                      end-placeholder="结束日期"
                      align="right"
                    ></el-date-picker>
                  </td>
                  <td>
                    <select name="boolt" id="boolt">
                      <option value="0">请选择</option>
                      <option value="true">true</option>
                      <option value="false">false</option>
                    </select>
                  </td>
                  <td>
                    <input type="number" v-model="search.market_type" />
                  </td>
                  <td>
                    <input type="text" v-model="search.match_symbol" />
                  </td>
                  <td>
                    <input type="number" v-model="search.price_start" />
                  </td>
                  <td>
                    <input type="number" v-model="search.price_end" />
                  </td>
                  <td>
                    <select name="boolw" id="boolw">
                      <option value="0">请选择</option>
                      <option value="true" selected>true</option>
                      <option value="false">false</option>
                    </select>
                  </td>
                  <td>
                    <select name="bools" id="bools">
                      <option value="0">请选择</option>
                      <option value="true" selected>true</option>
                      <option value="false">false</option>
                    </select>
                  </td>
                  <td>
                    <input type="number" placeholder="int" v-model="search.status" />
                  </td>
                  <td>
                    <input type="text" v-model="search.name" />
                  </td>
                  <td>
                    <button id="search" class="btn btn-primary btn-sm">查询</button>
                  </td>
                </tr>
              </table>
            </div>
            <!-- 条件输入值 -->

            <!-- 查询出来的数据列表 -->
            <div class="box-body">
              <table class="table table-bordered">
                <thead>
                  <tr>
                    <th>订单Id</th>
                    <th>发送者</th>
                    <th>市场类型</th>
                    <th>详情</th>
                    <th>
                      订单的创建时间
                      <span id="direction" class="btn btn-primary btn-sm">前到后排序</span>
                    </th>
                    <th>订单当前的状态</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody id="delList">
                  <tr v-for="(item,index) in orderTable.pageData" :key="index">
                    <td>{{item.order_id }}</td>
                    <td>{{item.name }}</td>
                    <td>{{item.market_type }}</td>
                    <td>
                      <div class="dropdown">
                        <button
                          id="dLabel"
                          type="button"
                          data-toggle="dropdown"
                          aria-haspopup="true"
                          aria-expanded="false"
                        >
                          查看
                          <span class="caret"></span>
                        </button>
                        <ul class="dropdown-menu" aria-labelledby="dLabel">
                          <li>发送者属于哪条链:{{item.name_chain_code }}</li>
                          <li>接受地址属于那一条链:{{item.target_chain_code }}</li>
                          <li>订单的更新时间:{{item.update_at }}</li>
                          <li>交易类型:{{item.trade_type }}</li>
                          <li>交易对名称{{item.match_symbol }}</li>
                          <li>下单时的总量:{{item.dele_vol }}</li>
                          <li>卖出的代币名称:{{item.sell_token_symbol }}</li>
                          <li>卖出的代币数量:{{item.sell_token_amount }}</li>
                          <li>买的代币名称:{{item.buy_token_symbol }}</li>
                          <li>买的代币数量:{{item.buy_token_amount }}</li>
                          <li>卖出的代币合约:{{item.sell_token_contract }}</li>
                          <li>买的代币合约:{{item.buy_token_contract }}</li>
                          <li>下单价格:{{item.price }}</li>
                          <li>目前扣除的手续费:{{item.fee }}</li>
                          <li>是否可撤单:{{item.withdraw_able }}</li>
                          <li>订单是否完成:{{item.ex_success }}</li>
                          <li>接受代币的账户地址:{{item.target_address }}</li>
                          <li>订单下单的交易Hash: {{item.tx_hash }}</li>
                        </ul>
                      </div>
                    </td>
                    <td>{{item.create_at }}</td>
                    <td>{{item.status }}</td>
                    <td>
                      <section class="btn btn-primary btn-sm">删除</section>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            <!-- 分页按钮 -->
            <el-pagination
              background
              layout="prev, pager, next"
              @prev-click="pageMinus"
              @next-click="pagePlus"
              @current-change="orderCurrentPage"
              :page-size="orderTable.numberOfItems"
              :hide-on-single-page="true"
              :total="items.length"
            ></el-pagination>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
import $ from "jquery";
export default {
  data() {
    return {
      pickerOptions: {
        shortcuts: [
          {
            text: "最近一周",
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
              picker.$emit("pick", [start, end]);
            }
          },
          {
            text: "最近一个月",
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
              picker.$emit("pick", [start, end]);
            }
          },
          {
            text: "最近三个月",
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
              picker.$emit("pick", [start, end]);
            }
          }
        ]
      },
      value2: [new Date(2000, 0, 1, 0, 0, 0), new Date(9999, 0, 1, 0, 0, 0)],
      search: {
        time_start: "",
        time_end: "",
        time_type: false,
        market_type: 0,
        match_symbol: "",
        price_start: 0,
        price_end: 0,
        withdraw_able: true,
        success: true,
        status: 0,
        name: ""
      },
      orderTable: {
        pageData: [], //当前页的数据
        numberOfItems: 10, //一页展示的数量
        totalNumberOfPages: "",
        currentPage: 1 //当前页数
      },
      items: []
    };
  },
  created() {
    // this.init()
  },
  mounted() {
    $(".dropdown-toggle").dropdown();
    sessionStorage.setItem("currentPage", 2);
    this.testLevel();
    this.deleteOrder();

    //从交易所跳转过来的，读取数据
    if (this.$store.state.login.flag === 1) {
      let data = this.$store.state.login.data;
      let obj = {
        time_start: "1800-01-01T00:00:00Z",
        time_end: "9999-01-01T00:00:00Z",
        time_type: false,
        market_type: data,
        match_symbol: "",
        price_start: 0,
        price_end: 0,
        withdraw_able: false,
        success: true,
        status: 0,
        name: ""
      };
      this.$store.commit("changeFlag", 0);
      this.ajaxGetData(obj)
    }
    //排序按钮逻辑
    let direction = document.querySelector("#direction");
    direction.onclick = () => {
      if (direction.innerText === "后到前排序") {
        direction.innerText = "前到后排序";
      } else {
        direction.innerText = "后到前排序";
      }
      this.items = this.items.reverse();
    };

    //进入页面后就加载所有数据
    let sh = {
      time_start: "1800-01-01T00:00:00Z",
      time_end: "9999-12-31T00:00:00Z",
      time_type: false,
      market_type: 0,
      match_symbol: "",
      price_start: 0,
      price_end: 0,
      withdraw_able: true,
      success: true,
      status: 0,
      name: ""
    };
    this.ajaxGetData(sh)    

    //点击按钮查询对应条件
    let search = document.querySelector("#search");
    search.onclick = () => {
      this.searchOrder();
    };
  },
  watch: {
    /**
     * 1. 检测总数据
     * 2. 数据处理，分页
     * 3. 分页数据展示
     * 4. 展示的数据对应的分页按钮加active类
     */
    items() {
      if (this.items !== null) {
        this.orderTable.pageData = this.items.length / 10;
        this.init();
      }
    }
  },
  methods: {
    ajaxGetData (obj) {
      let that = this;
      $.ajax({
        type:"post",
        url:"http://47.103.108.61:9090/api/admin/getAllDeleteOrder",
        data:JSON.stringify(obj),
        dataType:'json',
        crossDomain: true,  //跨域,
        headers: {
          "Authorization": "Bearer " + that.$store.state.login.token
        },
        success:function(response){
          let res = response
          // console.log(res)
          if (res.error_code === 0) {
            if (res.data !== null) that.items = res.data;
            else that.items = [];
            that.displayInfo("查询成功!");
          } else {
            //提示错误
            that.displayInfo("查询失败");
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
    date() {
      let str1 = this.value2[0].toString();
      let str2 = this.value2[1].toString();
      str1 = str1.slice(4, 24);
      str2 = str2.slice(4, 24);
      let month = [
        { text: "Jan", index: "01" },
        { text: "Feb", index: "02" },
        { text: "Mar", index: "03" },
        { text: "Apr", index: "04" },
        { text: "May", index: "05" },
        { text: "Jun", index: "06" },
        { text: "Jul", index: "07" },
        { text: "Aug", index: "08" },
        { text: "Sep", index: "09" },
        { text: "Oct", index: "10" },
        { text: "Nov", index: "11" },
        { text: "Dec", index: "12" }
      ];
      for (let i = 0; i < 12; i++) {
        if (str1.slice(0, 3) === month[i].text) {
          str1 =
            str1.slice(7, 11) +
            "-" +
            month[i].index +
            "-" +
            str1.slice(4, 6) +
            "T" +
            str1.slice(12, 20) +
            "Z";
          break;
        }
      }
      for (let i = 0; i < 12; i++) {
        if (str2.slice(0, 3) === month[i].text) {
          str2 =
            str2.slice(7, 11) +
            "-" +
            month[i].index +
            "-" +
            str2.slice(4, 6) +
            "T" +
            str2.slice(12, 20) +
            "Z";
          break;
        }
      }
      // console.log(str1);
      // console.log(str2);
      return { str1, str2 };
    },
    /**
     * 检测用户权限等级
     */
    testLevel() {
      //检测普通用户及其他情况下都禁用修改权限，只有为超级管理员才能修改
      let deleteBtn = document.querySelectorAll("#delList button");
      for (let i = 0; i < deleteBtn.length; i++) {
        deleteBtn[i].setAttribute("style", "pointer-events:none;");
      }

      //判断是否为超级管理员
      if (this.$store.state.login.userLevel === "super") {
        let deleteBtn = document.querySelectorAll("#delList button");
        for (let i = 0; i < deleteBtn.length; i++) {
          deleteBtn[i].setAttribute("style", "pointer-events:auto;");
        }
      }
    },
    deleteOrder() {
      let delList = document.querySelector("#delList");
      delList.onclick = e => {
        let target = e.target || event.target;
        if (target.nodeName === "SECTION") {
          this.$confirm("此操作将永久删除该项, 是否继续?", "提示", {
            confirmButtonText: "确定",
            cancelButtonText: "取消",
            type: "warning"
          }).then(() => {
            let param = target.parentNode.parentNode.children[0].innerText;
            let delObj = { order_id: param };
            let that = this;
            $.ajax({
              type:"post",
              url:"http://47.103.108.61:9090/api/admin/dbDeleteOrder",
              data:JSON.stringify(delObj),
              dataType:'json',
              crossDomain: true,  //跨域,
              headers: {
                "Authorization": "Bearer " + that.$store.state.login.token
              },
              success:function(response){
                let res = response
                // console.log(res)
                if (res.error_code === 0 && res.data.success === true) {
                  that.displayInfo("删除成功!");
                  // that.ajaxGetData(obj)
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
          }).catch(() => {
            this.$message({
              type: "info",
              message: "已取消删除"
            });
          });
        }
      };
    },
    /**
     * 查询订单
     * 参数：
     *  time_start: '',
        time_end: '',
        time_type: false,
        market_type: 0,
        match_symbol: '',
        price_start: 0,
        price_end: 0,
        withdraw_able: false,
        success: false,
        status: 0,
        name: ''
     */
    searchOrder() {
      let time = this.date();
      this.search.time_start = time.str1;
      this.search.time_end = time.str2;

      // console.log('change after start'+this.search.time_start)
      // console.log('change after end'+this.search.time_end)
      //三个布尔选项>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  start
      let selectType = document.querySelector("#boolt");
      let selectWith = document.querySelector("#boolw");
      let selectSucc = document.querySelector("#bools");
      //获取选中项索引值
      let indext = selectType.selectedIndex;
      let indexw = selectWith.selectedIndex;
      let indexs = selectSucc.selectedIndex;
      // 获取选择的值
      if (selectType.options[indext].value !== "0") {
        this.search.time_type = selectType.options[indext].text;
      }
      if (selectWith.options[indexw].value !== "0") {
        this.search.withdraw_able = selectWith.options[indexw].text;
      }
      if (selectSucc.options[indexs].value !== "0") {
        this.search.success = selectSucc.options[indexs].text;
      }
      //字符串 转 布尔类型
      if (this.search.time_type === "true") this.search.time_type = true;
      else if (this.search.time_type === "false") this.search.time_type = false;

      if (this.search.withdraw_able === "true")
        this.search.withdraw_able = true;
      else if (this.search.withdraw_able === "false")
        this.search.withdraw_able = false;

      if (this.search.success === "true") this.search.success = true;
      else if (this.search.success === "false") this.search.success = false;
      //三个布尔选项>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  end

      //四个数字选项>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  start
      this.search.price_start = parseInt(this.search.price_start);
      this.search.price_end = parseInt(this.search.price_end);
      this.search.status = parseInt(this.search.status);
      this.search.market_type = parseInt(this.search.market_type);
      //四个数字选项>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  end
      let search = this.search;
      this.ajaxGetData(search)      
    },
    /**
     * 传一个参数作为要展示的数据页数进行展示
     */
    orderCurrentPage(val) {
      this.orderTable.currentPage = val; // 修改当前页数

      //检测是否为最后一页
      if (this.orderTable.totalNumberOfPages === val) {
        //当前页为最后一页
        this.orderTable.pageData = this.items.slice(
          (val - 1) * this.orderTable.numberOfItems,
          this.items.length
        ); //展示的页数为当前页数所对应的数据
      } else {
        //非最后一页
        this.orderTable.pageData = this.items.slice(
          (val - 1) * this.orderTable.numberOfItems,
          val * this.orderTable.numberOfItems
        ); //展示的页数为当前页数所对应的数据
      }
    },
    /**
     * 初始化页数及展示第一页数据
     */
    init() {
      this.orderTable.pageData = this.items.slice(
        0,
        this.orderTable.numberOfItems
      ); //第一页的数据
      this.orderTable.currentPage = 1; //当前页数设为第一页
      this.orderTable.totalNumberOfPages = Math.ceil(
        this.items.length / this.orderTable.numberOfItems
      ); //总页数
      //marketcode
      for (let i = 0; i < this.orderTable.pageData.length; i++) {
        if (this.orderTable.pageData[i].market_type === 1) {
          this.orderTable.pageData[i].market_type = "eos";
        } else if (this.orderTable.pageData[i].market_type === 2) {
          this.orderTable.pageData[i].market_type = "eth";
        } else if (this.orderTable.pageData[i].market_type === 3) {
          this.orderTable.pageData[i].market_type = "bdex";
        } else if (this.orderTable.pageData[i].market_type === 4) {
          this.orderTable.pageData[i].market_type = "跨链";
        }
      }
      for (let i = 0; i < this.orderTable.pageData.length; i++) {
        if (this.orderTable.pageData[i].trade_type === "1") {
          this.orderTable.pageData[i].trade_type = "卖单";
        } else if (this.orderTable.pageData[i].trade_type === "0") {
          this.orderTable.pageData[i].trade_type = "买单";
        }
      }
      // 1:完全撮合成功 2:未完全撮合 3:已被撤单 4:下单失败 5:下单成功 6:错误订单
      for (let i = 0; i < this.orderTable.pageData.length; i++) {
        if (this.orderTable.pageData[i].status === 1) {
          this.orderTable.pageData[i].status = "完全撮合成功";
        } else if (this.orderTable.pageData[i].status === 2) {
          this.orderTable.pageData[i].status = "未完全撮合";
        } else if (this.orderTable.pageData[i].status === 3) {
          this.orderTable.pageData[i].status = "已被撤单";
        } else if (this.orderTable.pageData[i].status === 4) {
          this.orderTable.pageData[i].status = "下单失败";
        } else if (this.orderTable.pageData[i].status === 5) {
          this.orderTable.pageData[i].status = "下单成功";
        } else if (this.orderTable.pageData[i].status === 6) {
          this.orderTable.pageData[i].status = "错误订单";
        }
      }
    },
    /**
     * 页数减
     */
    pageMinus() {
      if (this.orderTable.currentPage > 1) {
        this.orderTable.currentPage = this.orderTable.currentPage - 1;
        this.orderCurrentPage(this.orderTable.currentPage);
      }
    },
    /**
     * 页数加
     */
    pagePlus() {
      if (this.orderTable.currentPage < this.orderTable.totalNumberOfPages) {
        this.orderTable.currentPage = this.orderTable.currentPage + 1;
        this.orderCurrentPage(this.orderTable.currentPage);
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
th,
td {
  font-size: 15px;
  text-align: center;
  border: 1px solid gray;
  /* max-width: 350px; */
  /* min-width: 40px; */
  word-break: break-all;
}
/* table:last-of-type th:last-of-type{
  width: 70px;
} */
input[type="number"] {
  width: 100px;
}
input[type="checkbox"] {
  width: auto;
}
td input {
  text-align: center;
  width: 170px;
}
td input[type="datetime-local"] {
  width: 210px;
}
.dropdown-menu {
  width: 500px;
}
#direction {
  margin-left: 30px;
}
.content-wrapper {
  height: 100%;
  min-height: 1000px;
}
.box {
  padding-bottom: 100px;
}
</style>