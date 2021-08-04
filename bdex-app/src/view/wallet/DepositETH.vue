<template>
  <div class="deposit">
    <navbar text="收款" />
    <div class="address-area">
      <span>{{address}}</span>
      <div class="qrcode">
        <canvas id="canvas"></canvas>
      </div>
    </div>
    <button class="custom-button copy-button bottom" :data-clipboard-text="address">复制收款帐号</button>
  </div>
</template>

<script>
import Navbar from '../../components/Navbar'
import QRCode from 'qrcode'
import Clipboard from 'clipboard'

export default {
  name: 'deposit',
  data () {
    return {
    }
  },
  computed: {
    myEth: function () { return this.$store.state.account.myEth },
    address: function () { return this.myEth.address }
  },
  components: {
    Navbar
  },
  methods: {
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
    const self = this
    const clipboard = new Clipboard('.copy-button')
    clipboard.on('success', function (e) {
      self.$dialog.alert({
        message: '复制成功'
      })
    })
    clipboard.on('error', function (e) {
      console.log(e)
    })
  }
}
</script>

<style lang="scss" scoped>
.deposit {
  display: flex;
  flex-direction: column;
  align-items: center;
  button {
    border: none;
  }
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
