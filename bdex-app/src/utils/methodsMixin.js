import web3 from '../lib/eth/web3'
import { getCurrentBalance } from '../lib/eos/helper'

const methodsMixin = {
  methods: {
    getCNYRate (token) {
      if (token === 'ETH') {
        return 1500
      } else if (token === 'EOS') {
        return 20
      }
      return 1
    },
    fixed (val) {
      if (typeof val === 'string') {
        if (val.indexOf('.') > -1) {
          const valList = val.split('.')
          valList[1] = valList[1].slice(0, 4)
          return valList.join('.')
        } else {
          return val.slice(0, 4)
        }
      }
      return val && val.toFixed(4)
    },
    async getETHBalance () {
      const myEth = this.$store.state.account.myEth
      console.log('查询余额获取的账户地址是------------------------', myEth.address)
      if (myEth && myEth.address) {
        const value = await web3.eth.getBalance(myEth.address)
        return this.fixed(web3.utils.fromWei(value))
      } else {
        return 0
      }
    },
    async getEOSBalance (name) {
      return getCurrentBalance(name)
    }
  }
}

export default methodsMixin
