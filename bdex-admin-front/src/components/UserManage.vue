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
        BDEX后台管理系统
        <!-- <small>Version 2.0</small> -->
      </h1>
      <ol class="breadcrumb">
        <li>
          <a href="#">
            <i class="glyphicon glyphicon-home"></i> 主页
          </a>
        </li>
        <!--<li class="active">Dashboard</li>-->
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
                <thead>
                  <tr>
                    <th>用户唯一id</th>
                    <th>用户名</th>
                    <th>密码</th>
                    <th>权限等级:1=最高级，2=普通</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody id="userList">
                  <tr v-for="(item,index) in list" :key="index">
                    <td>{{item.id}}</td>
                    <td>{{item.name}}</td>
                    <td>{{item.pass_word}}</td>
                    <td>{{item.permission}}</td>
                    <td>
                      <button>删除</button>
                    </td>
                  </tr>
                </tbody>
              </table>
              <table class="table table-bordered">
                <thead>
                  <tr>
                    <th>用户唯一id</th>
                    <th>用户名</th>
                    <th>密码</th>
                    <th>权限等级:1=最高级，2=普通</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody>
                  <tr id="updateOrAdd">
                    <td>
                      <input type="number" placeholder="仅更新时生效" />
                    </td>
                    <td>
                      <input type="text" />
                    </td>
                    <td>
                      <input type="text" />
                    </td>
                    <td>
                      <input type="number" placeholder="1为最高级，2为普通级" />
                    </td>
                    <td>
                      <button id="update">更新</button>
                      <button id="add">添加</button>
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
import $ from "jquery";
export default {
  data() {
    return {
      list: [
        {
          id: 0,
          name: "",
          pass_word: "",
          permission: 0
        }
      ]
    };
  },
  mounted() {
    sessionStorage.setItem("currentPage", 5);
    this.getUser();
    this.deleteUser();
    this.updateUser();
    this.addUser();
  },
  methods: {
    /**
     * 获取用户列表
     */
    getUser() {
      let that = this;
      $.ajax({
        type:"post",
        url:"http://47.103.108.61:9090/api/admin/user/getUser",
        data:{},
        dataType:'json',
        headers: {
          "Authorization": "Bearer " + that.$store.state.login.token
        },
        success:function(response){
          let res = response
          // console.log(res)
          if (res.error_code === 0) {
            that.list = res.data;
          } else {
            that.displayInfo("获取用户管理列表失败!");
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
     * 删除对应用户
     * 若删除的是当前用户则提示重新登录
     */
    deleteUser() {
      let list = document.querySelector("#userList");
      list.addEventListener("click", ev => {
        if (ev.target.nodeName === "BUTTON") {
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
              url:"http://47.103.108.61:9090/api/admin/user/delete",
              data:JSON.stringify(obj),
              dataType:'json',
              headers: {
                "Authorization": "Bearer " + that.$store.state.login.token
              },
              success:function(response){
                let res = response
                // console.log(res)
                if (res.error_code === 0 && res.data.success === true) {
                  that.getUser();
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
          })
          .catch(() => {
            this.$message({
              type: "info",
              message: "已取消删除"
            });
          });
        }
      });
    },
    /**
     * 更新用户
     */
    updateUser() {
      let update = document.querySelector("#update");
      update.addEventListener("click", () => {
        let updateOrAdd = document.querySelector("#updateOrAdd");
        let id = id === "" ? "" : parseInt(updateOrAdd.children[0].children[0].value);
        let name = updateOrAdd.children[1].children[0].value.trim();
        let pass_word = updateOrAdd.children[2].children[0].value.trim();
        let permission = parseInt(updateOrAdd.children[3].children[0].value);
        //先获取数据再验证
        let that = this;
        $.ajax({
          type:"post",
          url:"http://47.103.108.61:9090/api/admin/user/getUser",
          data:{},
          dataType:'json',
          headers: {
            "Authorization": "Bearer " + that.$store.state.login.token
          },
          success:function(response){
            let res = response
            // console.log(res)
            if (res.error_code === 0) {
              if (id === "" ||name === "" ||pass_word === "" ||permission === "") {
                that.displayInfo("输入不能为空!");
              } else {
                for (let i = 0; i < res.data.length; i++) {
                  if (permission !== 1 && permission !== 2) {
                    that.displayInfo("权限错误!");
                    break;
                  } else if (parseInt(res.data[i].id) === id) {
                    //更新用户列表
                    let obj = {
                      id,
                      name,
                      pass_word,
                      permission
                    };
                    $.ajax({
                      type:"post",
                      url:"http://47.103.108.61:9090/api/admin/user/update",
                      data:JSON.stringify(obj),
                      dataType:'json',
                      headers: {
                        "Authorization": "Bearer " + that.$store.state.login.token
                      },
                      success:function(response){
                        let res = response
                        // console.log(res)
                        if (res.error_code === 0) {
                          that.getUser();
                          that.displayInfo("更新成功!");
                        } else {
                          that.displayInfo("更新失败!");
                        }
                      },
                      error:function(xhr, status, error){
                        that.$message({
                          type: "error",
                          message: "网络错误！"
                        });
                      }
                    })
                    break;
                  } else {
                    if (i === res.data.length - 1) {
                      that.displayInfo("id所对应的用户不存在!");
                    }
                  }
                }
              }
            } else {
              //报错
              this.displayInfo("更新失败!");
            }
          },
          error:function(xhr, status, error){
            that.$message({
              type: "error",
              message: "网络错误！"
            });
          }
        })
      });
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
    },
    /**
     * 添加用户
     */
    addUser() {
      let add = document.querySelector("#add");
      add.addEventListener("click", () => {
        let updateOrAdd = document.querySelector("#updateOrAdd");
        let name = updateOrAdd.children[1].children[0].value.trim();
        let pass_word = updateOrAdd.children[2].children[0].value.trim();
        let permission = parseInt(updateOrAdd.children[3].children[0].value);

        //先获取数据再验证
        let that = this;
        $.ajax({
          type:"post",
          url:"http://47.103.108.61:9090/api/admin/user/getUser",
          data:{},
          dataType:'json',
          headers: {
            "Authorization": "Bearer " + that.$store.state.login.token
          },
          success:function(response){
            let res = response
            // console.log(res)
            if (res.error_code === 0) {
              if (name === "" || pass_word === "" || permission === "") {
                that.displayInfo("输入不能为空!");
              } else {
                for (let i = 0; i < res.data.length; i++) {
                  if (res.data[i].name === name) {
                    that.displayInfo("用户名已存在!");
                    break;
                  } else if (permission !== 1 && permission !== 2) {
                    that.displayInfo("权限错误!");
                    break;
                  } else if (i === res.data.length - 1) {
                    //更新用户列表
                    let obj = {
                      name,
                      pass_word,
                      permission
                    };
                    $.ajax({
                      type:"post",
                      url:"http://47.103.108.61:9090/api/admin/user/create",
                      data:JSON.stringify(obj),
                      dataType:'json',
                      headers: {
                        "Authorization": "Bearer " + that.$store.state.login.token
                      },
                      success:function(response){
                        let res = response
                        // console.log(res)
                        if (res.error_code === 0) {
                          that.getUser();
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
                    break;
                  }
                }
              }
            } else {
              //报错
              this.displayInfo("添加失败！");
            }
          },
          error:function(xhr, status, error){
            that.$message({
              type: "error",
              message: "网络错误！"
            });
          }
        })
      });
    }
  }
};
</script>
<style scoped>
</style>