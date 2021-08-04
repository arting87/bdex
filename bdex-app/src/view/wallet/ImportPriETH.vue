<template>
  <div class="import-pri">
    <navbar text="通过私钥导入" />
    <div class="import-pri-input-item">
      <div class="item-container">
        <span>输入私钥</span>
        <textarea placeholder="粘贴密钥" v-model="pri" />
      </div>
    </div>
    <input-item :value.sync="pass" type="password" label="钱包密码" placeholder="至少8位字符，包含大小写字母和数字" />
    <input-item :value.sync="repass" type="password" label="重复钱包密码" placeholder="再次输入钱包密码" />
    <div class="service-term">
      <span>我已仔细阅读并同意《服务与隐私条款》</span>
    </div>
    <div class="custom-button" @click="importAccountClick">导入钱包</div>
  </div>
</template>

<script>
import Navbar from '../../components/Navbar'
import InputItem from '../../components/InputItem'

export default {
  name: 'import-pri-eth',
  data () {
    return {
      pri: '',
      pass: '',
      repass: ''
    }
  },
  components: {
    Navbar,
    InputItem
  },
  methods: {
    importAccountClick () {
      const self = this
      this.loading = true
      this.$store.dispatch('account/importEth', {
        privateKey: this.pri,
        password: this.pass
      }).then(wallet => {
        if (wallet) {
          this.address = wallet.address
          this.importAccountDialog = false
          this.loading = false
          this.importAccountInnerDialog = true
          this.passErr = true
          // 导入钱包输入框 置空
          this.pri = ''
          this.pass = ''
          self.$router.push('/my')
        } else {
          this.importAccountDialog = false
          this.loading = false
          this.passErr = true
          // 导入钱包输入框 置空
          this.pri = ''
          this.pass = ''
          this.$toast('导入钱包失败')
        }
      })
    }
  }
}
</script>

<style lang='scss' scoped>
.import-pri {
  display: flex;
  flex-direction: column;
  align-items: center;
  .import-pri-input-item {
    width: 100%;
    border-bottom: solid .1rem $divideColor;
    display: flex;
    justify-content: center;
    .item-container {
      width: 33.5rem;
      display: flex;
      flex-direction: column;
      span {
        margin-top: 2.5rem;
        color: $textColor;
        font-size: 1.4rem;
      }
      textarea {
        margin-top: 1rem;
        width: 100%;
        height: 6rem;
        border: solid 0.1rem #21263c;
        font-size: 1.2rem;
        background-color: $bg;
        opacity: 0.5;
        color: $textColor;
        ::placeholder {
          color: rgba(130, 150, 238, 0.5);
        }
      }
    }
  }
  .service-term {
    width: 33.5rem;
    margin-top: 1.5rem;
    font-size: 1.2rem;
    color: $textColor;
  }
  .custom-button {
    margin-top: 6.8rem;
  }
}
</style>
