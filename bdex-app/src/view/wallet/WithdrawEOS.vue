<template>
  <div class="withdraw">
    <navbar text="提现" />
    <div class="fixed-width item">
      <span>收款帐号</span>
      <input placeholder="请输入收款帐号" v-model="address"/>
    </div>
    <div class="divide-1 fixed-width"></div>
    <div class="fixed-width item">
      <span>转账金额</span>
      <input placeholder="请输入需要转出的数量" v-model="amount"/>
    </div>
    <div class="divide-1 fixed-width"></div>
    <div class="fixed-width remain">
      <span>帐号余额</span>
      <span>{{remainAmount}} EOS</span>
    </div>
    <div class="fixed-width item">
      <span>Memo</span>
      <input placeholder="Memo" v-model="memo"/>
    </div>
    <div class="divide-1 fixed-width"></div>
    <div class="custom-button next-step" @click="nextStep">下一步</div>
  </div>
</template>

<script>
import Navbar from '../../components/Navbar'
import { sendEos } from '../../lib/eos/helper'

export default {
  name: 'withdraw',
  data () {
    return {
      address: '',
      memo: '',
      amount: '',
      remainAmount: 0
    }
  },
  components: {
    Navbar
  },
  created () {

  },
  mounted () {
  },
  methods: {
    nextStep () {
      const self = this
      const myEos = this.$store.state.account.myEos
      const from = myEos.address
      const pk = myEos.privateKey
      sendEos(from, this.address, pk, this.amount, this.memo)
        .then(data => {
          console.log(data)
          self.$dialog.alert({
            message: '交易已发送，等待确认'
          })
        })
        .catch(err => {
          console.log(err)
          self.$dialog.alert({
            message: '提币失败'
          })
        })
    }
  }
}
</script>

<style lang="scss" scoped>
.withdraw {
  display: flex;
  flex-direction: column;
  align-items: center;
  .item {
    height: 8.4rem;
    display: flex;
    flex-direction: column;
    span {
      margin-top: 2.5rem;
      color: $textColor;
      font-size: 1.4rem;
    }
    input {
      padding: 1rem 0;
    }
    input {
      color: $lightTextColor;
      font-size: 1.2rem;
    }
  }
  .remain {
    margin-top: 1rem;
    display: flex;
    justify-content: space-between;
    font-size: 1.2rem;
    :first-child {
      color: $textColor;
    }
    :last-child {
      color: $lightTextColor;
    }
  }
  .next-step {
    margin-top: 6rem;
  }
}
</style>
