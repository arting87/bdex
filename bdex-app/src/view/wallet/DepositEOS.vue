<template>
  <div class="deposit">
    <navbar text="收款" />
    <div class="address-area">
      <span>{{address}}</span>
      <div class="qrcode">
        <canvas id="canvas"></canvas>
      </div>
    </div>
    <div class="custom-button bottom" @click="copyAddress">复制收款帐号</div>
  </div>
</template>

<script>
import Navbar from '../../components/Navbar'
import QRCode from 'qrcode'

export default {
  name: 'deposit',
  data () {
    return {
    }
  },
  computed: {
    myEos: function () { return this.$store.state.account.myEos },
    address: function () { return this.myEos.address }
  },
  components: {
    Navbar
  },
  methods: {
    copyAddress () {
      this.$dialog.alert({
        message: '复制成功'
      })
    },
    useqrcode () {
      const canvas = document.getElementById('canvas')
      if (this.address) {
        QRCode.toCanvas(canvas, this.address, function (error) {
          if (error) console.error(error)
          console.log('QRCode success!')
        })
      }
    }
  },
  mounted () {
    this.useqrcode()
  }
}
</script>

<style lang="scss" scoped>
.deposit {
  display: flex;
  flex-direction: column;
  align-items: center;
  .address-area {
    margin-top: 5rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    .qrcode {
      margin-top: 4rem;
    }
    span {
      font-size: 1.2rem;
      color: $textColor;
    }
  }
  .bottom {
    margin-top: 6rem;
  }
}
</style>
