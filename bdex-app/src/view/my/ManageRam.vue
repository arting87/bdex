<template>
  <div class="manage-ram">
    <navbar text="内存" />
    <div class="fixed-width eos-net">
      <div class="eos-net-item">
        <span>内存</span>
        <span>价格：0 EOS/KB</span>
      </div>
      <div class="status-bar-area">
        <div class="status-bar">
          <div class="active-area"></div>
        </div>
        <span>剩余：{{accountInfo.ram_usage | getKB}}KB/{{accountInfo.ram_quota | getKB}}KB</span>
      </div>
    </div>
    <div class="divide-1"></div>
    <div class="fixed-width remain">
      <span>帐号余额：{{accountInfo.core_liquid_balance}}</span>
    </div>
    <div class="divide-1 fixed-width"></div>
    <div class="fixed-width item">
      <span>购买内存</span>
      <input placeholder="输入EOS数量" v-model="ramAmount"/>
    </div>
    <div class="divide-1 fixed-width"></div>
    <div class="fixed-width receive-account">
      <span>接收账户</span>
      <input placeholder="账户名" v-model="receiveAccount"/>
    </div>
    <div class="custom-button next-step" @click="nextStep">买入</div>
  </div>
</template>

<script>
import Navbar from '../../components/Navbar'
import { buyRam, sellRam } from '../../lib/eos/helper'

export default {
  name: 'manage-resouce',
  data () {
    return {
      ramAmount: 0,
      receiveAccount: '',
      action: 'buy'
    }
  },
  computed: {
    myEos: function () { return this.$store.state.account.myEos },
    accountInfo: function () { return this.$store.state.account.accountInfo }
  },
  components: {
    Navbar
  },
  created () {
    if (this.myEos && this.myEos.address) {
      this.$store.dispatch('account/getAccountInfo', this.myEos.address)
    }
  },
  mounted () {
  },
  methods: {
    nextStep () {
      const self = this
      const myEos = this.$store.state.account.myEos
      const from = myEos.address
      const pk = myEos.privateKey
      if (this.action === 'sell') {
        sellRam(from, pk, this.receiveAccount, this.ramAmount)
          .then(data => {
            console.log(data)
            self.$dialog.alert({
              message: '交易已发送，等待确认'
            })
          })
          .catch(err => {
            console.log(err)
            self.$dialog.alert({
              message: '发送失败'
            })
          })
        return
      }
      buyRam(from, pk, this.receiveAccount, this.ramAmount)
        .then(data => {
          console.log(data)
          self.$dialog.alert({
            message: '交易已发送，等待确认'
          })
        })
        .catch(err => {
          console.log(err)
          self.$dialog.alert({
            message: '发送失败'
          })
        })
    }
  }
}
</script>

<style lang="scss" scoped>
.manage-ram {
  display: flex;
  flex-direction: column;
  align-items: center;
  .status-bar-area {
    height: 2rem;
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    span {
      margin-top: -1.8rem;
      font-size: 1.2rem;
      color: #fff;
      justify-self: center;
    }
    .status-bar {
      width: 100%;
      height: 100%;
      background-color: #4E628D;
      border-radius: 1rem;
      .active-area {
        height: 100%;
        background-color: #6A87C4;
        border-top-left-radius: 1rem;
        border-bottom-left-radius: 1rem;
        width: 50%;
      }
    }
  }
  .eos-net {
    height: 10rem;
    justify-content: space-around;
    display: flex;
    flex-direction: column;
    .eos-net-item {
      :first-child {
        color: $textColor;
        font-size: 1.4rem;
      }
      :last-child {
        font-size: 1.2rem;
        color: $lightTextColor;
      }
    }
    .eos-net-item-pledge {
      :first-child {
        color: $lightTextColor;
      }
    }
  }
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
      color: $lightTextColor;
      font-size: 1.2rem;
    }
  }
  .receive-account {
    height: 4rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    span {
      color: $textColor;
      font-size: 1.4rem;
    }
    input {
      padding: 1rem 0;
      color: $lightTextColor;
      font-size: 1.2rem;
    }
  }
  .remain {
    height: 4.8rem;
    display: flex;
    font-size: 1.2rem;
    align-items: center;
    color: $lightTextColor;
  }
  .next-step {
    margin-top: 2rem;
  }
}
</style>
