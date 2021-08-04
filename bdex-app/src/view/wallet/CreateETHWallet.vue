<template>
  <div class="create-wallet">
    <navbar text="创建钱包" />
    <input-item :value.sync="pass" type="password" label="钱包密码" placeholder="至少8位字符，包含大小写字母和数字" />
    <input-item :value.sync="repass" type="password" label="重复钱包密码" placeholder="再次输入钱包密码" />
    <div class="service-term">
      <span>我已仔细阅读并同意《服务与隐私条款》</span>
    </div>
    <div class="custom-button" @click="createWallet">创建钱包</div>
  </div>
</template>

<script>
import Navbar from '../../components/Navbar'
import InputItem from '../../components/InputItem'

export default {
  name: 'create-wallet',
  data () {
    return {
      pass: '',
      repass: ''
    }
  },
  components: {
    Navbar,
    InputItem
  },
  methods: {
    check () {
      if (!this.pass || !this.repass) {
        this.$toast('请输入密码')
        return false
      }
      if (this.pass.length < 8 || this.repass.length < 8) {
        this.$toast('密码长度最少为8位')
        return false
      }
      if (this.pass !== this.repass) {
        this.$toast('两次密码不相同')
        return false
      }
      return true
    },
    createWallet () {
      console.log(this.pass, this.repass)
      if (!this.check()) {
        return
      }
      const self = this
      this.loading = true
      this.$store.dispatch('account/createEth', {
        password: this.pass
      }).then(wallet => {
        if (wallet) {
          this.address = wallet.address
          this.privateKey = wallet.privateKey
          console.log(this.privateKey)
          this.newAccountOuterDialog = false
          this.loading = false // 后期修改
          this.newAccountInnerDialog = true
          this.passErr = true
          // 清空输入框
          this.pass = ''
          this.repass = ''
          self.$router.push('/my')
        } else {
          this.newAccountOuterDialog = false
          this.loading = false
          this.passErr = true
          // 清空输入框
          this.pass = ''
          this.repass = ''
          this.$message.error('创建钱包失败')
        }
      })
    }
  }
}
</script>

<style lang='scss' scoped>
.create-wallet {
  display: flex;
  flex-direction: column;
  align-items: center;
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
