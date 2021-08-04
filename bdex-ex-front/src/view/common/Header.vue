<template>
  <div class="header">
    <el-menu
      :router="true"
      class="el-menu-demo"
      mode="horizontal"
      background-color="#131624"
      text-color="#FFFFFF"
    >
      <el-menu-item index="0" :route="{path: '/home'}" class="logostyle">
        <img src="../../assets/img/bdexlogo.png" alt="logo" class="logo">
      </el-menu-item>
      <el-menu-item index="0" :route="{path: '/home'}">{{ $t('header.home')}}</el-menu-item>
      <el-menu-item index="1" :route="{path: '/trade'}">{{ $t('header.exchange')}}</el-menu-item>
      <el-menu-item index="2" :route="{path: '/news'}">{{ $t('header.news')}}</el-menu-item>

      <el-submenu index="login">
        <template slot="title">{{ $t('header.login')}}</template>
        <el-menu-item>
          {{ $t('header.EosAccount')}}
          <el-button @click="LoginEos" type="text" style="font-size:16px">{{eosAccountName}}</el-button>
          <el-button @click="LogoutEos" type="text" size="small" style="font-size:16px">{{eosExit}}</el-button>
        </el-menu-item>
      </el-submenu>

      <!-- hide-timeout=30000 -->
      <el-submenu index="ethwallet">
        <template slot="title">{{ $t('header.wallet')}}</template>
        <el-menu-item class="account">{{ myEthAdress }}</el-menu-item>
        <el-menu-item
          @click="newAccountOuterDialog = true"
          :disabled="ethShow.create"
        >{{ $t('eth.newaccount')}}</el-menu-item>
        <el-menu-item
          @click="importAccountDialog = true"
          :disabled="ethShow.import"
        >{{ $t('eth.importaccount')}}</el-menu-item>
        <el-menu-item
          @click="unlockAccountDialog = true"
          :disabled="ethShow.unlock"
        >{{ $t('eth.unlockwallet')}}</el-menu-item>
        <el-menu-item
          @click="exportAccountOuterDialog = true"
          :disabled="ethShow.export"
        >{{ $t('eth.importkey')}}</el-menu-item>
        <el-menu-item
          @click="deleteAccountOuterDialog = true"
          :disabled="ethShow.delete"
        >{{ $t('eth.deleaccount')}}</el-menu-item>
      </el-submenu>

      <el-submenu index="4">
        <template slot="title">{{ $t('header.lang')}}</template>
        <el-menu-item @click="switchLang('zh')">简体中文</el-menu-item>
        <el-menu-item @click="switchLang('en')">English</el-menu-item>
      </el-submenu>
    </el-menu>

    <!-- 新建账号的 外层dialog :title="$t('eth.newaccount')"-->
    <el-dialog
      :modal="modal"
      :visible.sync="newAccountOuterDialog"
      v-loading="loading"
      width="480px"
      top="32vh"
      class="newaccountdialog"
    >
      <div class="newaccountbody">
        <el-image :src="newaccounturl" fit="fill" class="newaccountimg"></el-image>
        <div class="newaccounttitle">{{ $t('eth.newaccount') }}</div>
        <el-form :model="createEth">
          <el-form-item>
            <el-input
              v-model="createEth.password"
              autocomplete="off"
              show-password
              :placeholder="$t('eth.enterpassword')"
            ></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button
            type="primary"
            @click="newAccountClick"
            :disabled="passErr"
          >{{ $t('eth.createaccount') }}</el-button>
        </div>
      </div>
    </el-dialog>

    <!--新建账号的 第二层dialog -->
    <el-dialog
      :modal="modal"
      :visible.sync="newAccountInnerDialog"
      width="480px"
      top="32vh"
      class="newaccountinfodialog"
    >
      <div class="newaccountinfobody">
        <el-image :src="newaccountinfourl" fit="fill" class="warningimg"></el-image>
        <h3>{{ $t('eth.part1')}}</h3>
        <p>{{address}}</p>
        <h3 class="warntitle">{{ $t('eth.part2')}}</h3>
        <p>{{privateKey}}</p>
        <span slot="footer" class="dialog-footer">
          <el-button type="primary" @click="newAccountInnerDialog = false">{{ $t('eth.yes')}}</el-button>
        </span>
      </div>
    </el-dialog>

    <!-- 导入钱包dialog -->
    <el-dialog
      :modal="modal"
      :visible.sync="importAccountDialog"
      top="32vh"
      class="newaccountdialog"
      width="480px"
    >
      <div class="newaccountbody">
        <el-image :src="importaccounturl" fit="fill" class="importaccountimg"></el-image>
        <div class="importaccounttitle">{{ $t('eth.importaccount')}}</div>
        <el-form :model="impEth">
          <el-form-item>
            <el-input
              v-model="impEth.privateKey"
              autocomplete="off"
              :placeholder="$t('eth.enterprivatekey')"
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-input
              v-model="impEth.password"
              autocomplete="off"
              show-password
              :placeholder="$t('eth.enterpassword')"
            ></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button
            type="primary"
            @click="importAccountClick"
            :disabled="passErr"
          >{{ $t('eth.importaccount')}}</el-button>
        </div>
      </div>
    </el-dialog>

    <!--导入账号的 第二层dialog -->
    <el-dialog
      :modal="modal"
      :visible.sync="importAccountInnerDialog"
      width="480px"
      top="32vh"
      class="importinnerdialog"
    >
      <div class="importinnerbody">
        <el-image :src="newaccounturl" fit="fill" class="newaccountimg"></el-image>
        <h3>{{ $t('eth.part3')}}</h3>
        <div>
          <h4>{{address}}</h4>
        </div>
        <span slot="footer" class="dialog-footer">
          <el-button
            type="primary"
            @click="importAccountInnerDialog = false"
            class="importinnerbtn"
          >{{ $t('eth.yes')}}</el-button>
        </span>
      </div>
    </el-dialog>

    <!-- 导出账号的外层dialog-->
    <el-dialog
      :modal="modal"
      :visible.sync="exportAccountOuterDialog"
      v-loading="loading"
      top="32vh"
      class="newaccountdialog"
      width="480px"
    >
      <div class="newaccountbody">
        <el-image :src="exportaccounturl" fit="fill" class="exportaccountimg"></el-image>
        <div class="newaccounttitle">{{ $t('eth.importkey')}}</div>
        <el-form :model="exportEth">
          <el-form-item>
            <el-input
              v-model="exportEth.password"
              autocomplete="off"
              show-password
              :placeholder="$t('eth.enterpassword')"
            ></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button
            type="primary"
            @click="exportAccountClick"
            :disabled="passErr"
          >{{ $t('eth.importkey')}}</el-button>
        </div>
      </div>
    </el-dialog>

    <!-- 导出账号的第二层dialog -->
    <el-dialog
      :modal="modal"
      :visible.sync="exportAccountInnerDialog"
      top="32vh"
      class="newaccountinfodialog"
      width="480px"
    >
      <div class="newaccountinfobody">
        <el-image :src="newaccounturl" fit="fill" class="newaccountimg"></el-image>
        <h3>{{ $t('eth.part4')}}</h3>
        <p>{{ address }}</p>
        <h3>{{ $t('eth.part5')}}</h3>
        <p>{{ privateKey }}</p>
        <span slot="footer" class="dialog-footer">
          <el-button
            type="primary"
            @click="exportAccountInnerDialog = false"
            class="exportaccountinnerbtn"
          >{{ $t('eth.yes')}}</el-button>
        </span>
      </div>
    </el-dialog>

    <!-- 删除账号的第一层dialog -->
    <el-dialog
      :modal="modal"
      :visible.sync="deleteAccountOuterDialog"
      v-loading="loading"
      top="32vh"
      class="newaccountdialog"
      width="480px"
    >
      <div class="newaccountbody">
        <el-image :src="deleteaccounturl" fit="fill" class="deleteaccountimg"></el-image>
        <div class="newaccounttitle">{{ $t('eth.deleaccount')}}</div>
        <el-form :model="deleteEth">
          <el-form-item>
            <el-input
              v-model="deleteEth.password"
              autocomplete="off"
              show-password
              :placeholder="$t('eth.enterpassword')"
            ></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button
            type="primary"
            @click="deleteAccountClick"
            :disabled="passErr"
          >{{ $t('eth.deleaccount')}}</el-button>
        </div>
      </div>
    </el-dialog>
    <!-- 删除账号的第二层dialog -->
    <el-dialog
      :modal="modal"
      :visible.sync="deleteAccountInnerDialog"
      top="32vh"
      class="delinnerdialog"
      width="480px"
    >
      <div class="delinnerbody">
        <el-image :src="newaccountinfourl" fit="fill" class="warningimg"></el-image>
        <h3 class="h3style">{{ $t('eth.part6')}}</h3>
        <div class="ethadress">
          <h4>{{myEthAdress}}</h4>
        </div>
        <h3>{{ $t('eth.part7')}}</h3>
        <span slot="footer" class="dialog-footer">
          <el-button @click="deleteAccountInnerDialog = false">{{ $t('eth.cancel')}}</el-button>
          <el-button type="primary" @click="confirmDeleteAccount">{{ $t('eth.confirmDelete')}}</el-button>
        </span>
      </div>
    </el-dialog>

    <!-- 提醒解锁ETH钱包dialog-->
    <el-dialog
      :modal="modal"
      :visible.sync="unlockAccountDialog"
      v-loading="loading"
      top="32vh"
      class="newaccountdialog"
      width="480px"
    >
      <div class="newaccountbody">
        <el-image :src="unlockaccounturl" fit="fill" class="unlocaccountimg"></el-image>
        <div class="newaccounttitle">{{ $t('eth.unlockwallet')}}</div>
        <el-form :model="unlockEth">
          <el-form-item>
            <el-input
              v-model="unlockEth.password"
              autocomplete="off"
              show-password
              :placeholder="$t('eth.enterpassword')"
            ></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button
            type="primary"
            @click="unlocakAccountClick"
            :disabled="passErr"
          >{{ $t('eth.unlockwallet')}}</el-button>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {Login, Logout} from '../../plugins/scatter/scatter'
import store from '../../store/index'

export default {
  name: 'Header',
  data () {
    return {
      unlockaccounturl: require('../../assets/img/unlockwallet.png'),
      deleteaccounturl: require('../../assets/img/deleteaccount.png'),
      exportaccounturl: require('../../assets/img/exportaccount.png'),
      newaccounturl: require('../../assets/img/newaccount.png'),
      newaccountinfourl: require('../../assets/img/warning.png'),
      importaccounturl: require('../../assets/img/importaccount.png'),
      homeActiveIndex: '0',
      newAccountOuterDialog: false, // 新建账号 第一层dialog
      newAccountInnerDialog: false, // 第二层 dialog
      importAccountDialog: false, // 导入 钱包 1
      importAccountInnerDialog: false, // 2
      exportAccountOuterDialog: false, // 导出 第一层1
      exportAccountInnerDialog: false, // 第二层2
      deleteAccountOuterDialog: false, // 删除账号 第一层
      deleteAccountInnerDialog: false, // 删除账号 2
      unlockAccountDialog: false, // 解锁eth钱包
      loading: false,
      address: '',
      modal: false,
      privateKey: '',
      createEth: { // 创建钱包
        password: ''
      },
      impEth: { // 导入钱包时使用
        privateKey: '',
        password: ''
      },
      exportEth: {
        password: ''
      }, // 导出钱包
      unlockEth: {
        password: ''
      },
      deleteEth: {
        password: ''
      },
      passErr: true // 用于控制 是否可以操作
    }
  },
  mounted () {
  },
  computed: {
    eosAccountName () {
      const lang = this.$i18n.locale
      const myEos = this.$store.state.account.myEos
      if (myEos !== null) {
        return this.$store.state.account.myEos.identity.accounts[0].name
      } else {
        if (lang === 'zh') {
          return '登录'
        }
        return 'login'
      }
    },
    eosExit () {
      const lang = this.$i18n.locale
      const myEos = this.$store.state.account.myEos
      if (myEos === null) {
        return ''
      } else {
        if (lang === 'zh') {
          return '退出'
        }
        return 'logout'
      }
    },
    myEthAdress () {
      let myEth = this.$store.state.account.myEth
      if (myEth) {
        return myEth.address
      } else {
        if (sessionStorage.getItem('myEth')) {
          return this.$t('eth.accoutStateInfo2') // 支持多语言
        } else {
          return this.$t('eth.accoutStateInfo') // 支持多语言
        }
      }
    },
    ethShow () {
      let myEth = this.$store.state.account.myEth
      if (myEth) {
        return {
          create: true,
          import: true,
          unlock: true,
          export: false,
          delete: false
        }
      } else {
        if (sessionStorage.getItem('myEth')) {
          return {
            create: true,
            import: true,
            unlock: false,
            export: false,
            delete: false
          }
        } else {
          return {
            create: false,
            import: false,
            unlock: true,
            export: true,
            delete: true
          }
        }
      }
    },
    isConnected () {
      let conn = this.$store.state.isConnected
      return conn
    }
  },
  methods: {
    switchLang (lang) {
      this.$store.commit('i18n/switchLang', lang)
    },
    newAccountClick () {
      this.loading = true
      this.$store.dispatch('account/createEth', {
        password: this.createEth.password
      }).then(wallet => {
        if (wallet) {
          this.address = wallet.address
          this.privateKey = wallet.privateKey
          this.newAccountOuterDialog = false
          this.loading = false // 后期修改
          this.newAccountInnerDialog = true
          this.passErr = true
          // 清空输入框
          this.createEth.password = ''
        } else {
          this.newAccountOuterDialog = false
          this.loading = false
          this.passErr = true
          // 清空输入框
          this.createEth.password = ''
          this.$message.error('创建钱包失败')
        }
      })
    },
    importAccountClick () {
      this.loading = true
      this.$store.dispatch('account/importEth', {
        privateKey: this.impEth.privateKey,
        password: this.impEth.password
      }).then(wallet => {
        if (wallet) {
          this.address = wallet.address
          this.importAccountDialog = false
          this.loading = false
          this.importAccountInnerDialog = true
          this.passErr = true
          // 导入钱包输入框 置空
          this.impEth.privateKey = ''
          this.impEth.password = ''
        } else {
          this.importAccountDialog = false
          this.loading = false
          this.passErr = true
          // 导入钱包输入框 置空
          this.impEth.privateKey = ''
          this.impEth.password = ''
          this.$message.error('导入钱包失败')
        }
      })
    },
    exportAccountClick () {
      this.loading = true
      this.$store.dispatch('account/exportMyEth', {
        password: this.exportEth.password
      }).then(wallet => {
        if (wallet) {
          this.address = wallet.address
          this.privateKey = wallet.privateKey
          this.exportAccountOuterDialog = false
          this.loading = false
          this.exportAccountInnerDialog = true
          this.passErr = true
          // 输入框置空
          this.exportEth.password = ''
        } else {
          this.exportAccountOuterDialog = false
          this.loading = false
          this.passErr = true
          // 输入框置空
          this.exportEth.password = ''
          this.$message.error('导出钱包失败')
        }
      })
    },
    deleteAccountClick () {
      this.loading = true
      this.$store.dispatch('account/unlockMyEth', {
        password: this.deleteEth.password
      }).then(wallet => {
        if (wallet) {
          this.deleteAccountOuterDialog = false
          this.loading = false
          this.passErr = true
          this.deleteEth.password = ''
          this.deleteAccountInnerDialog = true
        } else {
          this.deleteAccountOuterDialog = false
          this.loading = false
          this.passErr = true
          this.deleteEth.password = ''
          this.$message.error('密码错误')
        }
      })
    },
    confirmDeleteAccount () {
      this.loading = true
      let myEth = this.$store.state.account.myEth
      if (myEth) {
        this.$store.dispatch('account/deleteMyEth').then(() => {
          this.deleteAccountInnerDialog = false
          this.loading = false
          this.$message('删除成功')
        })
      } else {
        this.deleteAccountInnerDialog = false
        this.loading = false
        this.$message.error('删除失败')
      }
    },
    unlocakAccountClick () {
      this.loading = true
      this.$store.dispatch('account/unlockMyEth', {
        password: this.unlockEth.password
      }).then(wallet => {
        if (wallet) {
          this.unlockAccountDialog = false
          this.loading = false
          this.passErr = true
          // 置空输入框
          this.unlockEth.password = ''
          this.$message('解锁成功')
        } else {
          this.unlockAccountDialog = false
          this.loading = false
          this.passErr = true
          // 置空输入框
          this.unlockEth.password = ''
          // window.alert('解锁失败')
          this.$message.error('解锁失败')
        }
      })
    },
    // eos 登录，退出操作
    async LoginEos () {
      if (this.eosAccountName === '登录' || this.eosAccountName === 'login') {
        // 用户没有主动退出，而是直接关闭网页！然后进入交易所，登录时需要把之前的证书删除
        if (sessionStorage.getItem('myEos') !== 'myEos' && sessionStorage.getItem('myEos') !== 'logoutEos') {
          await Logout()
        }
        await Login()
      }
    },
    async LogoutEos () {
      if (this.eosAccountName !== '登录' && this.eosAccountName !== 'login') {
        await Logout()
      }
    }
  },
  watch: {
    'createEth.password': function (newVal) {
      if (newVal !== '') {
        this.passErr = false
      } else {
        this.passErr = true
      }
    },
    'impEth.password': function (newVal) {
      if (newVal !== '' && this.impEth.privateKey !== '') {
        this.passErr = false
      } else {
        this.passErr = true
      }
    },
    'impEth.privateKey': function (newVal) {
      if (newVal !== '' && this.impEth.password !== '') {
        this.passErr = false
      } else {
        this.passErr = true
      }
    },
    'exportEth.password': function (newVal) {
      if (newVal !== '') {
        this.passErr = false
      } else {
        this.passErr = true
      }
    },
    'unlockEth.password': function (newVal) {
      if (newVal !== '') {
        this.passErr = false
      } else {
        this.passErr = true
      }
    },
    'deleteEth.password': function (newVal) {
      if (newVal !== '') {
        this.passErr = false
      } else {
        this.passErr = true
      }
    }
  },
  beforeRouteEnter (to, from, next) {
    if (sessionStorage.getItem('myEth') && store.account.myEth === null) {
      next(vm => {
        vm.unlockAccountDialog = true
      })
    } else {
      next()
    }
  }
}
</script>

<style scoped>
  @import '../../assets/css/Header.css';
</style>
