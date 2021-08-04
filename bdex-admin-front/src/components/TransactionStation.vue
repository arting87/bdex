<!-- DelOrder.vue -->
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
        交易所状态
        <!-- <small>Version 2.0</small> -->
      </h1>
      <ol class="breadcrumb">
        <li>
          <a href="#">
            <i class="glyphicon glyphicon-tower"></i>交易所
          </a>
        </li>
        <li class="active">交易所状态</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">
      <div class="row">
        <div class="col-xs-12">
          <div class="box">
            <div class="box-header with-border">
              <form action class="form-inline">
                <div class="input-group" style="width:220px;">
                  <div class="input-group-addon">
                    <i class="fa fa-calendar"></i>
                  </div>
                  <input type="date" id="date" class="form-control pull-right" />
                </div>
                <div class="form-group">
                  <button type="button" class="btn btn-primary" id="search">查询</button>
                </div>
              </form>
            </div>
            <!-- /.box-header -->
            <div class="box-body">
              <div id="myChart" :style="{width: '1000px', height: '300px'}"></div>
              <table class="table table-bordered">
                <thead>
                  <tr>
                    <th>错误订单数量</th>
                    <th>撤单数量</th>
                    <th>完成订单数量</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>{{items[0].err_order_count}}</td>
                    <td>{{items[0].withdraw_order_count}}</td>
                    <td>{{items[0].deal_order_count}}</td>
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
      items: [
        {
          err_order_count: "",
          withdraw_order_count: "",
          deal_order_count: ""
        }
      ],
      orderTable: {
        pageData: [], //当前页的数据
        numberOfItems: 10, //一页展示的数量
        totalNumberOfPages: "",
        currentPage: 1 //当前页数
      },
      text: {
        name1: "当天24小时内eos链内交易数量",
        name2: "当天24小时内eth链内交易数量",
        name3: "当天24小时内bde链内交易数量",
        name4: "当天24小时内eos跨链交易数量"
      }
    };
  },
  created() {
    // this.init()
  },
  /**
   * 获得时间-->查询数据-->饼状图显示
   */
  mounted() {
    sessionStorage.setItem("currentPage", 3);
    this.searchOrder(); //查询数据,后台交互是异步的
    setTimeout(() => {
      // this.init()
      this.generateChart(); //饼状图生成
    }, 200);

    //用户点击饼状图时处理逻辑
    let chart = document.querySelector("#myChart"); //获取对象
    if (chart !== null) {
      chart.onclick = () => {
        switch (
          chart.children[1].innerText
            .split("源")[1]
            .split(" ")[0]
            .trim()
        ) {
          case this.text.name1: //eos链内
            // console.log(this.text.name1)
            this.$store.commit("changeData", 1);
            this.$store.commit("changeFlag", 1);
            break;
          case this.text.name2: //eos链内
            // console.log(this.text.name2)
            this.$store.commit("changeData", 2);
            this.$store.commit("changeFlag", 1);
            break;
          case this.text.name3: //eos链内
            // console.log(this.text.name3)
            this.$store.commit("changeData", 3);
            this.$store.commit("changeFlag", 1);
            break;
          case this.text.name4: //eos链内
            // console.log(this.text.name4)
            this.$store.commit("changeData", 4);
            this.$store.commit("changeFlag", 1);
            break;
        }
        this.$router.replace({ path: "/listOrders" });
      };
    }
  },
  watch: {
    /**
     * 1. 检测总数据
     * 2. 数据处理，分页
     * 3. 分页数据展示
     * 4. 展示的数据对应的分页按钮加active类
     */
    items() {
      if (this.items !== null && this.items !== [])
        this.orderTable.pageData = this.items.length / 10;
    }
  },
  methods: {
    /**
     * 获取当日开始时间及现在时间
     */
    getTime() {
      let date = new Date();
      let year = date.getFullYear();
      let month = date.getMonth();
      let day = date.getDate();
      let hour = date.getHours();
      let minute = date.getMinutes();
      let seconds = date.getSeconds();

      month++;
      if (month < 10) month = "0" + month;
      if (day < 10) day = "0" + day
      let commonDay = year + "-" + month + "-" + day;
      let dateDay = year + "-" + month + "-" + day + "T" + "00:00:00+08:00";
      let timeStart = year + "-" + month + "-" + day + " " + "00:00:00";
      let timeNow =
        year +
        "-" +
        month +
        "-" +
        day +
        " " +
        hour +
        ":" +
        minute +
        ":" +
        seconds;

      let time = { timeStart, timeNow, dateDay, commonDay };
      return time;
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
      this.changeStyle(this.orderTable.currentPage); //分页按钮的样式
    },
    /**
     * 初始化页数及展示第一页数据
     */
    init() {
      // console.log(this.items.length)
      this.orderTable.pageData = this.items.slice(
        0,
        this.orderTable.numberOfItems
      ); //第一页的数据
      this.orderTable.currentPage = 1; //当前页数设为第一页
      this.orderTable.totalNumberOfPages = Math.ceil(
        this.items.length / this.orderTable.numberOfItems
      ); //总页数
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
    changeStyle(val) {
      //获取对象
      let number1 = document.querySelector("#number1");
      let number2 = document.querySelector("#number2");
      let number3 = document.querySelector("#number3");
      //清除所有类
      number1.removeAttribute("class");
      number2.removeAttribute("class");
      number3.removeAttribute("class");
      //给页数对应的按钮对象加类
      if (val <= 3) {
        number3.children[0].innerText = 3;
        number2.children[0].innerText = 2;
        number1.children[0].innerText = 1;
      } else {
        number3.children[0].innerText = this.orderTable.currentPage;
        number2.children[0].innerText = this.orderTable.currentPage - 1;
        number1.children[0].innerText = this.orderTable.currentPage - 2;
      }

      if (val === 1) {
        number1.className = "active";
      } else if (val === 2) {
        number2.className = "active";
      } else if (val >= 3) {
        number3.className = "active";
      }
    },
    generateChart() {
      let myChart = this.$echarts.init(document.getElementById("myChart"));
      // 绘制图表
      // console.log(this.items)
      let option = {
        title: {
          text: "24小时内下单数量",
          subtext: this.items[0].order_count,
          x: "center"
        },
        tooltip: {
          trigger: "item",
          formatter: "{a} <br/>{b} : {c} ({d}%)"
        },
        legend: {
          orient: "vertical",
          x: "left",
          data: [
            this.text.name1,
            this.text.name2,
            this.text.name3,
            this.text.name4
          ]
        },
        toolbox: {
          show: true,
          feature: {
            mark: { show: true },
            dataView: { show: true, readOnly: false },
            magicType: {
              show: true,
              type: ["pie", "funnel"],
              option: {
                funnel: {
                  x: "25%",
                  width: "50%",
                  funnelAlign: "left",
                  max: 1548
                }
              }
            },
            restore: { show: true },
            saveAsImage: { show: true }
          }
        },
        calculable: true,
        series: [
          {
            name: "访问来源",
            type: "pie",
            radius: "55%",
            center: ["50%", "60%"],
            data: [
              { value: this.items[0].eos_in_deal_count, name: this.text.name1 },
              { value: this.items[0].eth_in_deal_count, name: this.text.name2 },
              { value: this.items[0].bde_in_deal_count, name: this.text.name3 },
              {
                value: this.items[0].eos_cross_deal_count,
                name: this.text.name4
              }
            ]
          }
        ]
      };
      myChart.setOption(option);
    },
    /**
     * 查询订单
     * 参数：
     *  day
     */
    searchOrder() {
      let search = document.querySelector("#search");
      let date = document.querySelector("#date");
      let time = this.getTime().dateDay;
      let obj = { day: time };

      date.setAttribute("value", this.getTime().commonDay);

      let that = this
      $.ajax({
        type:"post",
        url:"http://47.103.108.61:9090/api/admin/getDataInDay",
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
            that.items = [res.data];
          } else {
            that.displayInfo("获取数据失败！");
          }
        },
        error:function(xhr, status, error){
          that.$message({
            type: "error",
            message: "网络错误！"
          });
        }
      })

      search.onclick = () => {
        let date = document.querySelector("#date");
        let time = date.value + "T00:00:00+08:00";
        let obj2 = { day: time };
        $.ajax({
          type:"post",
          url:"http://47.103.108.61:9090/api/admin/getDataInDay",
          data:JSON.stringify(obj2),
          dataType:'json',
          crossDomain: true,  //跨域,
          headers: {
            "Authorization": "Bearer " + that.$store.state.login.token
          },
          success:function(response){
            let res = response
            if (res.error_code === 0) {
              that.items = [res.data];
              that.generateChart(); //重新渲染饼状图
            } else {
              that.displayInfo("获取数据失败！");
            }
          },
          error:function(error){
            // console.log(error);
            that.$message({
              type: "error",
              message: "网络错误！"
            });
          }
        })
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
      let errData = this.getCookieByName("token");
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
.content-wrapper {
  height: 100%;
  min-height: 1000px;
}
.box {
  padding-bottom: 100px;
}
</style>