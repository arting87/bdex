<template>
  <div class="manage-resource">
    <navbar text="CPU/NET" />
    <div class="fixed-width eos-net">
      <div class="eos-net-item">
        <span>EOS</span>
        <span>价格：0 EOS/ms/天</span>
      </div>
      <div class="status-bar-area">
        <div class="status-bar">
          <div class="active-area"></div>
        </div>
        <span>剩余：{{cpuLimit.available | getCpu}}ms/{{cpuLimit.max | getCpu}}ms</span>
      </div>
      <div class="eos-net-item-pledge">
        <span>总抵押</span>
        <span>抵押：{{cpuLimit.used | getCpu}} EOS</span>
      </div>
      <div class="eos-net-item">
        <span>{{cpuLimit.used | getCpu}} EOS</span>
      </div>
    </div>
    <div class="divide-1"></div>
    <div class="fixed-width eos-net">
      <div class="eos-net-item">
        <span>网络</span>
        <span>价格：0.0004 EOS/KB/天</span>
      </div>
      <div class="status-bar-area">
        <div class="status-bar">
          <div class="active-area"></div>
        </div>
        <span>剩余：{{netLimit.available | getKB }}KB/{{netLimit.max | getKB }}KB</span>
      </div>
      <div class="eos-net-item-pledge">
        <span>总抵押</span>
        <span>抵押：{{netLimit.used | getKB }} EOS</span>
      </div>
      <div class="eos-net-item">
        <span>{{netLimit.used | getKB }} EOS</span>
      </div>
    </div>
    <div class="divide-1"></div>
    <div class="fixed-width remain">
      <span>帐号余额：{{accountInfo.core_liquid_balance}}</span>
    </div>
    <div class="divide-1 fixed-width"></div>
    <div class="fixed-width item">
      <span>计算抵押</span>
      <input placeholder="输入EOS数量" v-model="cpuAmount"/>
    </div>
    <div class="divide-1 fixed-width"></div>
    <div class="fixed-width item">
      <span>网络抵押</span>
      <input placeholder="输入EOS数量" v-model="netAmount"/>
    </div>
    <div class="divide-1 fixed-width"></div>
    <div class="fixed-width receive-account">
      <span>接收账户</span>
      <input placeholder="账户名" v-model="receiveAccount"/>
    </div>
    <div class="custom-button next-step" @click="nextStep">抵押</div>
  </div>
</template>

<script>
import Navbar from '../../components/Navbar'
import { buyResource } from '../../lib/eos/helper'

export default {
  name: 'manage-resouce',
  data () {
    return {
      remainAmount: 0,
      cpuAmount: 0,
      netAmount: 0,
      receiveAccount: 0
    }
  },
  computed: {
    myEos: function () { return this.$store.state.account.myEos },
    accountInfo: function () { return this.$store.state.account.accountInfo },
    cpuLimit: function () { return this.accountInfo && this.accountInfo.cpu_limit },
    netLimit: function () { return this.accountInfo && this.accountInfo.net_limit }
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
      console.log('pledge', this.computeAmount, this.netAmount)
      const self = this
      const myEos = this.$store.state.account.myEos
      const from = myEos.address
      const pk = myEos.privateKey
      buyResource(from, pk, this.netAmount, this.cpuAmount, this.receiveAccount, this.receiveAccount)
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
.manage-resource {
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
    height: 17rem;
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
