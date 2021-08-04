<template>
  <div class="create-wallet">
    <navbar text="创建钱包" />
    <input-item :value.sync="name" type="text" label="钱包名称" placeholder="设置钱包名称" />
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
      name: '',
      pass: '',
      repass: ''
    }
  },
  components: {
    Navbar,
    InputItem
  },
  methods: {
    createWallet () {
      console.log(this.name, this.pass, this.repass)
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
