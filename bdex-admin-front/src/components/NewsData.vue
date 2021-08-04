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
    <section class="content-header">
      <h1>
        新闻数据处理
        <!-- <small>Version 2.0</small> -->
      </h1>
      <ol class="breadcrumb">
        <li>
          <a href="#">
            <i class="glyphicon glyphicon-edit"></i>新闻
          </a>
        </li>
        <li class="active">新闻数据处理</li>
      </ol>
    </section>
    <!-- Main content -->
    <section class="content">
      <div class="row">
        <div class="col-xs-12">
          <div class="box">
            <!-- 查询代币条件部分 -->
            <!-- 交易列表 -->
            <div class="box-header with-border">
              <h3>展示与删除</h3>
            </div>
            <div class="box-body">
              <table class="table table-bordered">
                <thead>
                  <tr>
                    <th>新闻唯一id</th>
                    <th>新闻标题</th>
                    <th>新闻类型</th>
                    <th>新闻内容</th>
                    <th>时间</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody id="newsList">
                  <tr v-for="(item,index) in newsList.pageData" :key="index">
                    <td>{{item.id}}</td>
                    <td>{{item.title}}</td>
                    <td>{{item.type_code}}</td>
                    <td>
                      <a href="javascript:;" class="details" @click="toDetailContent()">详情</a>
                    </td>
                    <td>{{item.time}}</td>
                    <td>
                      <button class="btn btn-primary btn-sm" @click="deleteNews()">删除</button>
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
              :page-size="newsList.numberOfItems"
              :hide-on-single-page="false"
              :total="items.length"
            ></el-pagination>
            <!-- 添加新闻部分 -->
            <div class="box-header with-border">
              <h3>新闻添加</h3>
            </div>
            <div class="box-body">
              <header>
                新闻标题
                <input type="text" id="title" />
              </header>
              <main>
                <article>
                  <!-- 富文本编辑器 -->
                  <!-- <textarea name="content" id="editor"></textarea>
                   -->
                  <div id="editor"></div>
                </article>
              </main>
              <footer>
                <section>
                  <p>
                    新闻类型:
                    <!-- <input type="number" id="newsType"> -->
                    <!-- newsType -->
                    <select name="newsType" id="newsType">
                      <option value="1" selected>重要活动</option>
                      <option value="2">活动公告</option>
                      <option value="3">新币上线</option>
                      <option value="4">其他公告</option>
                    </select>
                  </p>
                </section>
                <section>
                  <button id="add" class="btn btn-primary btn-sm">添加</button>
                </section>
              </footer>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
// import ClassicEditor from "@ckeditor/ckeditor5-build-classic";
var E = require('wangeditor')
export default {
  data() {
    return {
      items: [],
      newsList: {
        pageData: [], //当前页的数据
        numberOfItems: 10, //一页展示的数量
        totalNumberOfPages: "",
        currentPage: 1 //当前页数
      }
    };
  },
  created() {
    this.init();
  },
  mounted() {
    sessionStorage.setItem("currentPage", 4);
    this.testLevel(); //检测权限

    let getNews = document.querySelector("#getNews");
    this.getNews(); //获取新闻

    let newsList = document.querySelector("#newsList");
    newsList.onclick = ev => {
      this.deleteNews(ev); //删除新闻
    };
    this.toDetailContent(); //跳转到详情页面
    this.addNews(); //添加新闻
  },
  watch: {
    /**
     * 1. 检测总数据
     * 2. 数据处理，分页
     * 3. 分页数据展示
     * 4. 展示的数据对应的分页按钮加active类
     */
    items() {
      this.init(); //重新渲染
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
      let deleteBtn = document.querySelectorAll("#newsList button");
      for (let i = 0; i < deleteBtn.length; i++) {
        deleteBtn[i].setAttribute("disabled", "disabled");
      }
      let add = document.querySelector("#add");
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
     * 获取新闻
     */
    getNews() {
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
            that.testLevel();
          } else {
            that.displayInfo("查询失败！");
          }
        },
        error:function(xhr, status, error){
          // console.log(xhr)
          // console.log(status, error);
          that.$message({
            type: "error",
            message: "网络错误！"
          });
        }
      })
    },
    /**
     * 删除新闻
     */
    deleteNews(ev) {
      if (ev !== undefined && ev.target.nodeName === "BUTTON") {
        this.$confirm("此操作将永久删除该项, 是否继续?", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }).then(() => {
          let id = parseInt(
            ev.target.parentNode.parentNode.children[0].innerText
          );
          let obj = { id };
          let that = this;
          $.ajax({
            type:"post",
            url:"http://47.103.108.61:9090/api/admin/deleteNews",
            data:JSON.stringify(obj),
            dataType:'json',
            headers: {
              "Authorization": "Bearer " + that.$store.state.login.token
            },
            success:function(response){
              let res = response
              // console.log(res)
              if (res.error_code === 0 && res.data.success === true) {
                //删除成功后立即更新列表
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
                      that.displayInfo("删除成功！");
                    } else {
                      that.displayInfo("删除失败！");
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
                this.displayInfo("删除失败!");
              }
            },
            error:function(xhr, status, error){
              that.$message({
                type: "error",
                message: "网络错误！"
              });
            }
          })
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
      }
    },
    /**
     * 用户点击详情时跳转到详情界面显示具体的新闻详情
     */
    toDetailContent() {
      let toDetails = document.querySelectorAll(".details");
      let newsList = document.querySelector("#newsList");
      newsList.addEventListener("click", ev => {
        if (ev.target.nodeName === "A") {
          let id = parseInt(
            ev.target.parentNode.parentNode.children[0].innerText
          );
          this.$store.commit("changeData", id);
          this.$router.push({ path: "/newsContent" });
        }
      });
    },
    /**
     * 用户点击添加按钮时向后台发送新添加的新闻数据
     */
    addNews() {
      // 先获取富文本编辑器里输入后转化的内容---->content
      var editor = new E('#editor')
      editor.create()
      let add = document.querySelector("#add");
      add.onclick = () => {
        // console.log(editor.txt.html());
        let content = editor.txt.html(); //输入的新闻内容html格式
        let newsType = parseInt(document.querySelector("#newsType").value);
        let title = document.querySelector("#title");
        let obj = {
          title: title.value,
          type_Code: newsType,
          content
        };
        //向后台发送添加数据
        let that = this;
        $.ajax({
          type:"post",
          url:"http://47.103.108.61:9090/api/admin/addNews",
          data:JSON.stringify(obj),
          dataType:'json',
          headers: {
            "Authorization": "Bearer " + that.$store.state.login.token
          },
          success:function(response){
            let res = response
            // console.log(res)
            if (res.error_code === 0 && res.data.success === true) {
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
                    title.value = "";
                    // editor.setData("");
                    editor.txt.html('')
                    document.querySelector("#newsType").value = "1";
                    that.displayInfo("添加成功!");
                  }else {
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
            } else {
              that.displayInfo("添加失败!");
            }
          },
          error:function(xhr, status, error){
            // console.log(xhr)
            // console.log(status, error);
            that.$message({
              type: "error",
              message: "网络错误！"
            });
          }
        })
      };
    },
    /**
     * 传一个参数作为要展示的数据页数进行展示
     */
    orderCurrentPage(val) {
      this.newsList.currentPage = val; // 修改当前页数
      if (this.items !== [] && this.items !== null) {
        //检测是否为最后一页
        if (this.newsList.totalNumberOfPages === val) {
          //当前页为最后一页
          this.newsList.pageData = this.items.slice(
            (val - 1) * this.newsList.numberOfItems,
            this.items.length
          ); //展示的页数为当前页数所对应的数据
        } else {
          //非最后一页
          this.newsList.pageData = this.items.slice(
            (val - 1) * this.newsList.numberOfItems,
            val * this.newsList.numberOfItems
          ); //展示的页数为当前页数所对应的数据
        }
        for (let i = 0; i < this.newsList.pageData.length; i++) {
          if (this.newsList.pageData[i].type_code === 1) {
            this.newsList.pageData[i].type_code = "重要活动";
          } else if (this.newsList.pageData[i].type_code === 2) {
            this.newsList.pageData[i].type_code = "活动公告";
          } else if (this.newsList.pageData[i].type_code === 3) {
            this.newsList.pageData[i].type_code = "新币上线";
          } else if (this.newsList.pageData[i].type_code === 4) {
            this.newsList.pageData[i].type_code = "其他公告";
          }
        }
      }
      // this.changeStyle(this.newsList.currentPage) //分页按钮的样式
    },
    /**
     * 初始化页数及展示第一页数据
     */
    init() {
      if (this.items !== [] && this.items !== null) {
        if (this.items.length < 10) {
          this.newsList.pageData = this.items; //第一页的数据
        } else {
          this.newsList.pageData = this.items.slice(
            0,
            this.newsList.numberOfItems
          ); //第一页的数据
        }
        this.newsList.currentPage = 1; //当前页数设为第一页
        this.newsList.totalNumberOfPages = Math.ceil(
          this.items.length / this.newsList.numberOfItems
        ); //总页数
        // 1 重要公告，2 活动公告，3 新币上线，4 其它公告
        for (let i = 0; i < this.newsList.pageData.length; i++) {
          if (this.newsList.pageData[i].type_code === 1) {
            this.newsList.pageData[i].type_code = "重要活动";
          } else if (this.newsList.pageData[i].type_code === 2) {
            this.newsList.pageData[i].type_code = "活动公告";
          } else if (this.newsList.pageData[i].type_code === 3) {
            this.newsList.pageData[i].type_code = "新币上线";
          } else if (this.newsList.pageData[i].type_code === 4) {
            this.newsList.pageData[i].type_code = "其他公告";
          }
        }
      }
    },
    /**
     * 页数减
     */
    pageMinus() {
      if (this.newsList.currentPage > 1) {
        this.newsList.currentPage = this.newsList.currentPage - 1;
        this.orderCurrentPage(this.newsList.currentPage);
      }
    },
    /**
     * 页数加
     */
    pagePlus() {
      if (this.newsList.currentPage < this.newsList.totalNumberOfPages) {
        this.newsList.currentPage = this.newsList.currentPage + 1;
        this.orderCurrentPage(this.newsList.currentPage);
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
footer {
  margin-top: 20px;
}
form#add {
  text-align: center;
  margin: 0 auto;
  padding: 10px;
}
.box {
  padding-bottom: 100px;
}
.content-wrapper {
  height: 100%;
  min-height: 1000px;
}
input {
  display: block;
  width: 100%;
}
#getNews {
  margin-left: 30px;
}
footer input {
  width: auto;
}
</style>