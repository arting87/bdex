<template>
  <div class="tradefootertop">
    <el-tabs type="border-card" v-model="activeName" >
      <el-tab-pane :label="$t('trade.openOrder')" name="first">
        <el-table
          :data="currentOrder"
          style="width: 100%"
          v-if="activeName === 'first'"
          height="361"
        >
          <el-table-column prop="pair" :label="$t('trade.pair')" width="100">
            <template slot-scope="scope">
              <span>{{ scope.row.pair }} {{ scope.row.code === '3' ? '/BDE' : (scope.row.code === '2' ? '/ ETH' : '/ EOS') }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="date" :label="$t('trade.time')" width="147">
            <template slot-scope="scope">
              <span>{{ scope.row.date | time }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="direction" :label="$t('trade.direction')" width="100">
            <template slot-scope="scope">
              <span>{{ scope.row.direction === 0 ? $t('trade.buy') : $t('trade.sell') }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="orderprice" :label="$t('trade.orderPrice')" width="150">
            <template slot-scope="scope">
              <span>{{ scope.row.orderPrice }} {{ scope.row.code === '3' ? 'BDE' : (scope.row.code === '2' ? ' ETH' : ' EOS') }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="ordervol" :label="$t('trade.orderVol')" width="150">
            <template slot-scope="scope">
              <span>{{ scope.row.orderVol.toFixed(4) }} {{ scope.row.pair}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="ordertotal" :label="$t('trade.orderTotal')" width="150">
            <template slot-scope="scope">
              <span>{{ (scope.row.orderVol * scope.row.orderPrice).toFixed(4) }} {{ scope.row.code === '3' ? 'BDE' : (scope.row.code === '2' ? ' ETH' : ' EOS') }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="dealtvol" :label="$t('trade.dealtVol')" width="150">
            <template slot-scope="scope">
              <span>{{ scope.row.dealVol.toFixed(4) }} {{ scope.row.pair }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="pendingVol" :label="$t('trade.pendingVol')" width="150">
            <template slot-scope="scope">
              <span>{{ scope.row.pendingVol.toFixed(4) }} {{ scope.row.pair }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="action" :label="$t('trade.action')" width="100">
            <template slot-scope="scope">
              <el-button @click="hint(scope.row)" type="text" size="small">{{$t('trade.withdraw') }}</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { getEosWithdrawSign } from '../../plugins/scatter/scatter'
import exchange from '../../lib/eth/exchange'
import Tx from 'ethereumjs-tx'
import web3 from '../../lib/eth/web3'

export default {
  name: 'TradeFooterTop',
  props: { // 组件间通信
    routeParam: {
      type: String,
      required: true
    }
  },
  data () {
    return {
      search: '',
      activeName: 'first'
    }
  },
  computed: {
    currentOrder () {
      let param = this.routeParam
      const currentOrder = this.$store.state.user.currentOrder
      if (currentOrder !== null && currentOrder.hasOwnProperty(param)) {
        return currentOrder[param]
      } else {
        return []
      }
    }
  },
  filters: {
    time: function (txdate) {
      const d = new Date(txdate)
      const month = (d.getMonth() + 1) < 10 ? '0' + (d.getMonth() + 1) : '' + (d.getMonth() + 1)
      const day = d.getDate() < 10 ? '0' + d.getDate() : '' + d.getDate()
      const hour = d.getHours() < 10 ? '0' + d.getHours() : '' + d.getHours()
      const minutes = d.getMinutes() < 10 ? '0' + d.getMinutes() : '' + d.getMinutes()
      // const second = d.getSeconds() < 10 ? '0' + d.getSeconds() : '' + d.getSeconds()
      return month + '-' + day + ' ' + hour + ':' + minutes
    }
  },
  methods: {
    async hint (row) {
      this.$confirm('您确定撤销该订单吗？', '温馨提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        let code = row.code
        let direction = row.direction
        if (code === '1') { // eos撤单
          let orderId = row.orderId
          let myContract = 'bdexcontract'
          let myEos = this.$store.state.account.myEos
          let isConnected = this.$store.state.isConnected
          if (myEos !== null && isConnected) {
            let wdSign = await getEosWithdrawSign(myContract, orderId)
            if (wdSign) {
              this.$socket.sendObj({
                namespace: 'user',
                action: 'withdraweos',
                data: {
                  orderId: orderId,
                  approveTx: JSON.stringify(wdSign)
                }
              })
              // 提示
              this.$message({
                type: 'success',
                message: '撤单请求已发出，稍等......'
              })
            } else {
              this.$message('取消撤单')
            }
          } else {
            this.$message({
              type: 'success',
              message: '账号未登录或网络故障！'
            })
          }
        } else if (code === '2') {
          const txid = row.orderId
          let myEth = this.$store.state.account.myEth
          let isConnected = this.$store.state.isConnected
          if (myEth !== null && isConnected) {
            const useraddress = this.$store.state.account.myEth.address
            const contractdata = await exchange.methods.withdrawal(web3.utils.toHex(txid))
            const payload = contractdata.encodeABI()

            const gas = await web3.eth.getGasPrice()
            const nonce = await web3.eth.getTransactionCount(useraddress)
            const gasPriceHex = web3.utils.numberToHex(gas)
            const nonceHex = web3.utils.numberToHex(nonce)
            const gasLimitHex = web3.utils.numberToHex(300000)

            const rawTx = {
              nonce: nonceHex,
              gasPrice: gasPriceHex,
              gasLimit: gasLimitHex,
              value: '0x00',
              from: useraddress,
              to: exchange.address,
              data: payload
            }
            const tx = new Tx(rawTx)
            const privateKey = Buffer.from(this.$store.state.account.myEth.privateKey.slice(2), 'hex')
            tx.sign(privateKey)
            const serializedTx = tx.serialize()

            this.$socket.sendObj({
              namespace: 'user',
              action: 'withdraweth',
              data: {
                orderId: txid,
                approveTx: '0x' + serializedTx.toString('hex')
              }
            })
            this.$message({
              type: 'success',
              message: '撤单请求已发出，稍等......'
            })
          } else {
            this.$message({
              type: 'success',
              message: '账号未登录或网络故障！'
            })
          }
        } else if (code === '3') { // todo 待做，8.13 新增需求，bde交易区
          const txid = row.orderId
          let myEth = this.$store.state.account.myEth
          let isConnected = this.$store.state.isConnected
          if (myEth !== null && isConnected) {
            const useraddress = this.$store.state.account.myEth.address
            const contractdata = await exchange.methods.withdrawal(web3.utils.toHex(txid))
            const payload = contractdata.encodeABI()

            const gas = await web3.eth.getGasPrice()
            const nonce = await web3.eth.getTransactionCount(useraddress)
            const gasPriceHex = web3.utils.numberToHex(gas)
            const nonceHex = web3.utils.numberToHex(nonce)
            const gasLimitHex = web3.utils.numberToHex(300000)

            const rawTx = {
              nonce: nonceHex,
              gasPrice: gasPriceHex,
              gasLimit: gasLimitHex,
              value: '0x00',
              from: useraddress,
              to: exchange.address,
              data: payload
            }
            const tx = new Tx(rawTx)
            const privateKey = Buffer.from(this.$store.state.account.myEth.privateKey.slice(2), 'hex')
            tx.sign(privateKey)
            const serializedTx = tx.serialize()

            this.$socket.sendObj({
              namespace: 'user',
              action: 'withdraweth',
              data: {
                orderId: txid,
                approveTx: '0x' + serializedTx.toString('hex')
              }
            })
            this.$message({
              type: 'success',
              message: '撤单请求已发出，稍等......'
            })
          } else {
            this.$message({
              type: 'success',
              message: '账号未登录或网络故障！'
            })
          }
        } else if (code === '4') {
          if (direction === 0) { // 跨链的撤单 eos
            let orderId = row.orderId
            let myContract = 'bdexcontract'
            let myEos = this.$store.state.account.myEos
            let isConnected = this.$store.state.isConnected
            if (myEos !== null && isConnected) {
              let wdSign = await getEosWithdrawSign(myContract, orderId)
              if (wdSign) {
                this.$socket.sendObj({
                  namespace: 'user',
                  action: 'withdraweos',
                  data: {
                    orderId: orderId,
                    approveTx: JSON.stringify(wdSign)
                  }
                })
                // 提示
                this.$message({
                  type: 'success',
                  message: '撤单请求已发出，稍等......'
                })
              } else {
                this.$message('取消撤单')
              }
            } else {
              this.$message({
                type: 'success',
                message: '账号未登录或网络故障！'
              })
            }
          } else { // 跨链的撤单 eth
            const txid = row.orderId
            let myEth = this.$store.state.account.myEth
            let isConnected = this.$store.state.isConnected
            if (myEth !== null && isConnected) {
              const useraddress = this.$store.state.account.myEth.address
              const contractdata = await exchange.methods.withdrawal(web3.utils.toHex(txid))
              const payload = contractdata.encodeABI()

              const gas = await web3.eth.getGasPrice()
              const nonce = await web3.eth.getTransactionCount(useraddress)
              const gasPriceHex = web3.utils.numberToHex(gas)
              const nonceHex = web3.utils.numberToHex(nonce)
              const gasLimitHex = web3.utils.numberToHex(300000)
              const rawTx = {
                nonce: nonceHex,
                gasPrice: gasPriceHex,
                gasLimit: gasLimitHex,
                value: '0x00',
                from: useraddress,
                to: exchange.address,
                data: payload
              }
              const tx = new Tx(rawTx)
              const privateKey = Buffer.from(this.$store.state.account.myEth.privateKey.slice(2), 'hex')
              tx.sign(privateKey)
              const serializedTx = tx.serialize()

              this.$socket.sendObj({
                namespace: 'user',
                action: 'withdraweth',
                data: {
                  orderId: txid,
                  approveTx: '0x' + serializedTx.toString('hex')
                }
              })
            }
          }
        }
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消撤单'
        })
      })
    }
  }
}
</script>

<style scoped>
  @import '../../assets/css/TradeFooter.css';
</style>
