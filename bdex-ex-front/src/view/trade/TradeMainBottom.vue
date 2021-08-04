<template>
  <div class="trademainbottom">
    <el-tabs type="card">

      <!-- 限价交易 -->
      <el-tab-pane :label="$t('trade.limitorder')">
        <el-row>

          <!-- 买入 -->
          <el-col :span="11" :offset="0" class="buy">
              <el-row class="titleheader">
                <el-col :span="8" :offset="2" class="color-green">
                  <div class="grid-content bg-purple">{{$t('trade.buy') }} {{ currentToken }}</div>
                </el-col>
                <el-col :span="14" :offset="0" class="balancecolor">
                  <div class="grid-content bg-purple">{{$t('trade.balance') }} {{ buyBalance }} {{ baseToken }}</div>
                </el-col>
              </el-row>

              <el-form
                :label-position="labelPosition"
                size="small"
                label-width="21px"
                :model="limitBuyOrder"
              >
                <el-form-item>
                  <el-input v-model="limitBuyOrder.price" clearable type="number" min="0.000001" @blur="handlePrice" class="bottominput">
                    <template slot="append">{{ baseToken }}</template>
                    <!-- <i slot="suffix">EOS</i> -->
                  </el-input>
                </el-form-item>

                <el-form-item  class="sliderbottom">
                  <el-input v-model="limitBuyOrder.amount" clearable type="number" min="0.0001" @blur="handleAmount" class="bottominput">
                    <!-- <i slot="suffix">BTC</i> -->
                    <template slot="append">{{ currentToken }}</template>
                  </el-input>
                </el-form-item>

                <el-col :span="16" :offset="3" class="slider">
                  <el-slider v-model="buyslider" :show-tooltip="false" :step="25" show-stops :marks="marks"></el-slider>
                </el-col>

                <el-form-item>
                  <el-input v-model="limitBuyOrder.total" clearable type="number" min="0.0001" @blur="handleTotal" class="bottominput">
                    <!-- <i slot="suffix">EOS</i> -->
                    <template slot="append">{{ baseToken }}</template>
                  </el-input>
                </el-form-item>

                <!-- @click="onSubmit" -->
                <el-form-item>
                  <el-button v-if="showApproveButBde" class="buy-sell-btn accredit-btn" @click="ethAprroveBde()">{{ $t('trade.accreditbtn') }}</el-button>
                  <el-button class="buy-sell-btn but-btn" @click="buyOrder()">{{$t('trade.buy') }}</el-button>
                  <el-dialog
                    :modal="modal"
                    :visible.sync="accreditDialogBde"
                    width="480px"
                    top="32vh"
                    class="newaccountdialog"
                  >
                    <div class="newaccountbody">
                      <el-image :src="exportaccounturl" fit="fill" class="exportaccountimg"></el-image>
                      <div class="newaccounttitle">{{ $t('trade.accredit') }}</div>
                          <el-input
                            v-model="accreditamountBde"
                            type="number"
                            autocomplete="off"
                            min="0"
                            class="accredit-input"
                            :placeholder="$t('trade.accreditPlaceholder')"
                          ></el-input>
                      <div slot="footer" class="dialog-footer">
                        <div>{{ $t('trade.beauthorized') }}{{ haveAccreditAmountBde }} {{ baseToken }}</div>
                        <el-button
                          type="primary"
                          @click="approveTokenBde()"
                          :disabled="showApproveBde"
                        >{{ $t('trade.accredit') }}</el-button>
                      </div>
                    </div>
                  </el-dialog>
                </el-form-item>

              </el-form>
          </el-col>

          <!-- 卖出 -->
          <el-col :span="11" :offset="0" class="sell">
            <div>
              <el-row class="titleheader">
                <el-col :span="8" :offset="2" class="color-red">
                  <div class="grid-content bg-purple">{{ $t('trade.sell') }} {{ currentToken }}</div>
                </el-col>
                <el-col :span="14" :offset="0" class="balancecolor">
                  <div class="grid-content bg-purple">{{$t('trade.balance') }} {{ sellBalance }} {{ currentToken }}</div>
                </el-col>
              </el-row>

              <el-form
                :label-position="labelPosition"
                size="small"
                label-width="21px"
                :model="limitSellOrder"
              >
                <el-form-item>
                  <el-input v-model="limitSellOrder.price" clearable type="number" min="0.000001" @blur="handlePrice" class="bottominput">
                    <!-- <i slot="suffix">EOS</i> -->
                    <template slot="append">{{ baseToken }}</template>
                  </el-input>
                </el-form-item>
                <el-form-item class="sliderbottom">
                  <el-input v-model="limitSellOrder.amount" clearable type="number" min="0.0001" @blur="handleAmount" class="bottominput">
                    <!-- <i slot="suffix">BTC</i> -->
                    <template slot="append">{{ currentToken }}</template>
                  </el-input>
                </el-form-item>
                <el-col :span="16" :offset="3" class="slider">
                  <el-slider v-model="sellslider" :step="25" :show-tooltip="false" show-stops :marks="marks"></el-slider>
                </el-col>
                <el-form-item>
                  <el-input v-model="limitSellOrder.total" clearable type="number" min="0.0001" @blur="handleTotal" class="bottominput">
                    <!-- <i slot="suffix">EOS</i> -->
                    <template slot="append">{{ baseToken }}</template>
                  </el-input>
                </el-form-item>

                <!-- @click="onSubmit"  v-if="showApproveBut"-->
                <el-form-item>
                  <el-button v-if="showApproveBut" class="buy-sell-btn accredit-btn" @click="ethAprrove()">{{ $t('trade.accreditbtn') }}</el-button>
                  <el-button class="buy-sell-btn sell-btn" @click="sellOrder()">{{ $t('trade.sell') }}</el-button>
                  <el-dialog
                    :modal="modal"
                    :visible.sync="accreditDialog"
                    width="480px"
                    top="32vh"
                    class="newaccountdialog"
                  >
                    <div class="newaccountbody">
                      <el-image :src="exportaccounturl" fit="fill" class="exportaccountimg"></el-image>
                      <div class="newaccounttitle">{{ $t('trade.accredit') }}</div>
                          <el-input
                            v-model="accreditamount"
                            type="number"
                            autocomplete="off"
                            min="0"
                            class="accredit-input"
                            :placeholder="$t('trade.accreditPlaceholder')"
                          ></el-input>
                      <div slot="footer" class="dialog-footer">
                        <div>{{ $t('trade.beauthorized') }}{{ haveAccreditAmount }} {{ currentToken }}</div>
                        <el-button
                          type="primary"
                          @click="approveToken()"
                          :disabled="showApprove"
                        >{{ $t('trade.accredit') }}</el-button>
                      </div>
                    </div>
                  </el-dialog>
                </el-form-item>

              </el-form>
            </div>
          </el-col>
        </el-row>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { getCurrencyBalance, getEosTransferSign } from '../../plugins/scatter/scatter'
import web3 from '../../lib/eth/web3'
import abiContract from '../../lib/eth/eth'
import exchange from '../../lib/eth/exchange'
import Tx from 'ethereumjs-tx'
import BigNumber from 'bignumber.js'
import {handleEthPrice, numFloat, handBig} from '../../utils/handleFloat'
import {myEosContract} from '../../lib/eos/config'

export default {
  name: 'TradeMainBottom',
  props: { // 组件间通信
    routeParam: {
      type: String,
      required: true
    }
  },
  data () {
    return {
      exportaccounturl: require('../../assets/img/exportaccount.png'),
      showApproveBut: false, // 是否显示授权按钮
      showApproveButBde: false, // todo bde交易区是否显示授权按钮
      haveAccreditAmount: 0,
      haveAccreditAmountBde: 0, // todo 已授权额度
      // 买入授权的dialog
      accreditDialogBde: false, // todo bde 买入的dialog
      // 卖出授权的dialog
      accreditDialog: false,
      modal: false,
      accreditamount: 0, // eth erc20授权额度
      accreditamountBde: 0, // todo bde授权额度
      showApprove: true, // 是否显示 授权合约
      showApproveBde: true, // 是否显示 授权合约
      buyBalance: 0.0000, // 使用监听方案测试
      sellBalance: 0.0000,
      labelPosition: 'right', // 控制label位置
      isReceive: false,
      limitBuyOrder: {
        price: 0.000000,
        amount: 0.0000,
        total: 0.0000,
        adress: '',
        remind: '请输入EOS的接收账户'
      },
      limitSellOrder: {
        price: 0.000000,
        amount: 0.0000,
        total: 0.0000,
        adress: '',
        remind: '请输入ETH的接收账户'
      },
      buyslider: 0,
      sellslider: 0,
      marks: {
        0: '0%',
        25: '25%',
        50: '50%',
        75: '75%',
        100: '100%'
      },
      whIdentity: [0, 0, 0, 0, 0, 0], // 循环监听标识，防止死循环监听 length === 6
      buyIntervalId: '', // 5/29更改，使用定时器不停的查询余额
      sellIntervalId: ''
    }
  },
  computed: {
    orderAction () { // 监听使用，更新余额
      let eosOrder = this.$store.state.user.eosOrder
      let ethOrder = this.$store.state.user.ethOrder
      let bdeOrder = this.$store.state.user.bdeOrder // todo 8.14新增
      let eosWithdraw = this.$store.state.user.eosWithdraw
      let ethWithdraw = this.$store.state.user.ethWithdraw
      let bdeWithdraw = this.$store.state.user.bdeWithdraw // todo 8.14新增
      return {
        eosOrder,
        ethOrder,
        bdeOrder,
        eosWithdraw,
        ethWithdraw,
        bdeWithdraw
      }
    },
    baseToken () {
      let param = this.routeParam
      let code = param.split('-')[1]
      if (code === '2') {
        return 'ETH'
      } else if (code === '3') { // todo 8.13新增
        return 'BDE'
      } else {
        return 'EOS'
      }
    },
    currentToken () {
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      return tokenName
    },
    account () {
      let myEos = this.$store.state.account.myEos
      let myEth = this.$store.state.account.myEth
      return {
        myEos,
        myEth
      }
    },
    limitBalanceBuy () {
      let param = this.routeParam
      let code = param.split('-')[1]
      if (code === '2' || code === '3') { // eth
        let buy = handBig.minus(parseFloat(this.buyBalance), 0.03) // 保留0.03个eth
        if (buy >= 0) {
          return buy
        } else {
          return 0
        }
      } else { // eos
        let buy = handBig.minus(parseFloat(this.buyBalance), 0.01) // 保留0.01个eos
        if (buy >= 0) {
          return buy
        } else {
          return 0
        }
      }
    },
    limitBalanceSell () {
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      let code = param.split('-')[1]
      if (tokenName === 'ETH' && code === '4') {
        let sell = handBig.minus(parseFloat(this.sellBalance), 0.03) // 跨链, 卖eth, 保留 0.03个
        if (sell >= 0) {
          return sell
        } else {
          return 0
        }
      } else {
        return parseFloat(this.sellBalance)
      }
    },
    tablePrice () {
      let price = this.$store.state.orderApi.price
      return price
    }
  },
  watch: {
    routeParam: {
      handler: async function (newVal) {
        // 清空 limitBuyOrder/limitSellOrder
        this.limitBuyOrder.price = 0
        this.limitBuyOrder.amount = 0
        this.limitBuyOrder.total = 0
        this.limitSellOrder.price = 0
        this.limitSellOrder.amount = 0
        this.limitSellOrder.total = 0
        // 清空 buyslider、sellslider
        this.buyslider = 0
        this.sellslider = 0
        // 清空Order  price
        this.$store.commit('orderApi/setPrice', 0.000000)
        let param = newVal
        let tokenName = param.split('-')[0]
        let code = param.split('-')[1]
        let myEos = this.$store.state.account.myEos
        let myEth = this.$store.state.account.myEth
        if (code === '1') {
          this.showApproveBut = false
          this.showApproveButBde = false
          if (myEos) {
            let eosAddress = this.$store.state.api.eos && this.$store.state.api.eos[tokenName].contractAddress.address
            if (eosAddress) {
              let buyBalance = await getCurrencyBalance('eosio.token', 'EOS') // 测试使用，应该是EOS
              this.buyBalance = parseFloat(buyBalance)
              let sellBalance = await getCurrencyBalance(eosAddress, tokenName)
              this.sellBalance = parseFloat(sellBalance)

              // 清空定时器
              clearInterval(this.buyIntervalId)
              clearInterval(this.sellIntervalId)
              // 5/29 修改，使用定时器1s获取余额；开启定时器
              this.buyIntervalId = setInterval(async () => {
                let buyBalance = await getCurrencyBalance('eosio.token', 'EOS') // 测试使用，应该是EOS
                this.buyBalance = parseFloat(buyBalance)
              }, 2000)
              this.sellIntervalId = setInterval(async () => {
                let sellBalance = await getCurrencyBalance(eosAddress, tokenName)
                this.sellBalance = parseFloat(sellBalance)
              }, 2000)
            }
          } else {
            // 清空定时器
            clearInterval(this.buyIntervalId)
            clearInterval(this.sellIntervalId)
            this.buyBalance = 0.0000
            this.sellBalance = 0.0000
          }
        } else if (code === '2') { // eth处理
          this.showApproveBut = true
          if (myEth) {
            let ethAddress = this.$store.state.api.eth && this.$store.state.api.eth[tokenName].contractAddress.address
            if (ethAddress) {
              let buyBalance = await web3.eth.getBalance(this.$store.state.account.myEth.address) // eth address
              buyBalance = (web3.utils.fromWei(buyBalance, 'ether'))
              this.buyBalance = numFloat(buyBalance, 4)
              const coin = new web3.eth.Contract(abiContract, ethAddress)
              let sellBalance = await coin.methods.balanceOf(this.$store.state.account.myEth.address).call()
              let ret = new BigNumber(sellBalance)
              const decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
              const de = 10 ** decimal
              const d = new BigNumber(de)
              sellBalance = parseFloat(ret.dividedBy(d)).toFixed(4)
              this.sellBalance = numFloat(sellBalance, 4)

              // 清空定时器
              clearInterval(this.buyIntervalId)
              clearInterval(this.sellIntervalId)
              // 5/29 使用定时器2s获取余额；开启定时器
              this.buyIntervalId = setInterval(async () => {
                let buyBalance = await web3.eth.getBalance(this.$store.state.account.myEth.address) // eth address
                buyBalance = (web3.utils.fromWei(buyBalance, 'ether'))
                this.buyBalance = numFloat(buyBalance, 4)
              }, 2000)
              this.sellIntervalId = setInterval(async () => {
                const coin = new web3.eth.Contract(abiContract, ethAddress)
                let sellBalance = await coin.methods.balanceOf(this.$store.state.account.myEth.address).call()
                let ret = new BigNumber(sellBalance)
                const decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
                const de = 10 ** decimal
                const d = new BigNumber(de)
                sellBalance = parseFloat(ret.dividedBy(d)).toFixed(4)
                this.sellBalance = numFloat(sellBalance, 4)
              }, 2000)
            }
          } else {
            // 清空定时器
            clearInterval(this.buyIntervalId)
            clearInterval(this.sellIntervalId)
            this.buyBalance = 0.0000
            this.sellBalance = 0.0000
          }
        } else if (code === '3') { // 侧链处理 todo 待做，8.13 新需求
          this.showApproveBut = true // 显示授权按钮操作
          this.showApproveButBde = true // 显示授权按钮操作
          if (myEth) {
            // erc20 代币bde的合约地址
            const bdeAddress = this.$store.state.api.eth && this.$store.state.api.eth['BDE'].contractAddress.address
            let ethAddress = this.$store.state.api.eth && this.$store.state.api.eth[tokenName].contractAddress.address
            if (bdeAddress && ethAddress) {
              // bde的余额获取和更新
              const bdeCoin = new web3.eth.Contract(abiContract, bdeAddress)
              let buyBalance = await bdeCoin.methods.balanceOf(this.$store.state.account.myEth.address).call()
              let buyBig = new BigNumber(buyBalance)
              const bdeDecimal = this.$store.state.api.eth['BDE'].contractAddress.decimal
              const bdeDe = 10 ** bdeDecimal
              const bdeD = new BigNumber(bdeDe)
              buyBalance = parseFloat(buyBig.dividedBy(bdeD)).toFixed(6) // todo eth数据会全部转化为整数进行处理，所以获取的bde数量需要divid对应的精度
              this.buyBalance = numFloat(buyBalance, 4)
              // bde交易区，其它erc20代币的余额
              const coin = new web3.eth.Contract(abiContract, ethAddress)
              let sellBalance = await coin.methods.balanceOf(this.$store.state.account.myEth.address).call()
              let ret = new BigNumber(sellBalance)
              const decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
              const de = 10 ** decimal
              const d = new BigNumber(de)
              sellBalance = parseFloat(ret.dividedBy(d)).toFixed(4)
              this.sellBalance = numFloat(sellBalance, 4)

              // 清空定时器
              clearInterval(this.buyIntervalId)
              clearInterval(this.sellIntervalId)

              // 5/29 使用定时器2s获取余额；开启定时器
              this.buyIntervalId = setInterval(async () => {
                const bdeCoin = new web3.eth.Contract(abiContract, bdeAddress)
                let buyBalance = await bdeCoin.methods.balanceOf(this.$store.state.account.myEth.address).call()
                let buyBig = new BigNumber(buyBalance)
                const bdeDecimal = this.$store.state.api.eth['BDE'].contractAddress.decimal
                const bdeDe = 10 ** bdeDecimal
                const bdeD = new BigNumber(bdeDe)
                buyBalance = parseFloat(buyBig.dividedBy(bdeD)).toFixed(6) // todo eth数据会全部转化为整数进行处理，所以获取的bde数量需要divid对应的精度
                this.buyBalance = numFloat(buyBalance, 4)
              }, 2000)
              this.sellIntervalId = setInterval(async () => {
                const coin = new web3.eth.Contract(abiContract, ethAddress)
                let sellBalance = await coin.methods.balanceOf(this.$store.state.account.myEth.address).call()
                let ret = new BigNumber(sellBalance)
                const decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
                const de = 10 ** decimal
                const d = new BigNumber(de)
                sellBalance = parseFloat(ret.dividedBy(d)).toFixed(4)
                this.sellBalance = numFloat(sellBalance, 4)
              }, 2000)
            }
          } else {
            // 清空定时器
            clearInterval(this.buyIntervalId)
            clearInterval(this.sellIntervalId)
            this.buyBalance = 0.0000
            this.sellBalance = 0.0000
          }
        } else if (code === '4') {
          this.showApproveBut = false
          this.showApproveButBde = false
          if (myEos) {
            let buyBalance = await getCurrencyBalance('eosio.token', 'EOS') // 测试使用，应该是EOS
            this.buyBalance = parseFloat(buyBalance)

            // 清空定时器
            clearInterval(this.buyIntervalId)
            // 开启定时器
            this.buyIntervalId = setInterval(async () => {
              let buyBalance = await getCurrencyBalance('eosio.token', 'EOS') // 测试使用，应该是EOS
              this.buyBalance = parseFloat(buyBalance)
            }, 2000)
          } else {
            // 清空定时器
            clearInterval(this.buyIntervalId)
            this.buyBalance = 0.0000
          }
          if (myEth) { // 获取eth余额
            let sellBalance = await web3.eth.getBalance(this.$store.state.account.myEth.address) // eth address
            sellBalance = (web3.utils.fromWei(sellBalance, 'ether'))
            this.sellBalance = numFloat(sellBalance, 4)

            // 清空定时器
            clearInterval(this.sellIntervalId)
            // 开启定时器
            this.sellIntervalId = setInterval(async () => {
              let sellBalance = await web3.eth.getBalance(this.$store.state.account.myEth.address) // eth address
              sellBalance = (web3.utils.fromWei(sellBalance, 'ether'))
              this.sellBalance = numFloat(sellBalance, 4)
            }, 2000)
          } else {
            // 清空定时器
            clearInterval(this.sellIntervalId)
            this.sellBalance = 0.0000
          }
        }
      },
      immediate: true
    },
    'account.myEos': {
      handler: async function (newVal) { // 用户登录和退出时，都会监听
        if (newVal !== null) {
          let param = this.routeParam
          let tokenName = param.split('-')[0]
          let code = param.split('-')[1]
          if (code === '1') {
            let eosAddress = this.$store.state.api.eos && this.$store.state.api.eos[tokenName].contractAddress.address
            if (eosAddress) {
              let buyBalance = await getCurrencyBalance('eosio.token', 'EOS') // 测试使用，应该是EOS
              this.buyBalance = parseFloat(buyBalance)
              let sellBalance = await getCurrencyBalance(eosAddress, tokenName)
              this.sellBalance = parseFloat(sellBalance)

              // 清空定时器
              clearInterval(this.buyIntervalId)
              clearInterval(this.sellIntervalId)
              // 开启定时器
              this.buyIntervalId = setInterval(async () => {
                let buyBalance = await getCurrencyBalance('eosio.token', 'EOS') // 测试使用，应该是EOS
                this.buyBalance = parseFloat(buyBalance)
              }, 2000)
              this.sellIntervalId = setInterval(async () => {
                let sellBalance = await getCurrencyBalance(eosAddress, tokenName)
                this.sellBalance = parseFloat(sellBalance)
              }, 2000)
            }
          } else if (code === '3') { // 侧链处理 todo 是否需要处理，待定 8.13新需求；不需要处理，eth erc20代币与eos账号状态无关

          } else if (code === '4') {
            let buyBalance = await getCurrencyBalance('eosio.token', 'EOS') // 测试使用，应该是EOS
            this.buyBalance = parseFloat(buyBalance)

            // 清空定时器
            clearInterval(this.buyIntervalId)
            // 开启定时器
            this.buyIntervalId = setInterval(async () => {
              let buyBalance = await getCurrencyBalance('eosio.token', 'EOS') // 测试使用，应该是EOS
              this.buyBalance = parseFloat(buyBalance)
            }, 2000)
          }
        } else {
          let param = this.routeParam
          let code = param.split('-')[1]
          if (code === '1') {
            // 清空定时器
            clearInterval(this.buyIntervalId)
            clearInterval(this.sellIntervalId)
            this.buyBalance = 0.0000
            this.sellBalance = 0.0000
          } else if (code === '3') { // todo 8.13 改动，code === '3'时，是eth erc20代币的bde交易区；eos账号的状态改动，于此无关
            // this.buyBalance = 0.0000
          } else if (code === '4') {
            // 清空定时器
            clearInterval(this.buyIntervalId)
            this.buyBalance = 0.0000
          }
        }
      }
      // immediate: true
    },
    'account.myEth': { // eth登录退出处理
      handler: async function (newVal) {
        if (newVal !== null) {
          let param = this.routeParam
          let tokenName = param.split('-')[0]
          let code = param.split('-')[1]
          if (code === '2') {
            let ethAddress = this.$store.state.api.eth && this.$store.state.api.eth[tokenName].contractAddress.address
            if (ethAddress) {
              let buyBalance = await web3.eth.getBalance(this.$store.state.account.myEth.address) // eth address
              buyBalance = (web3.utils.fromWei(buyBalance, 'ether'))
              this.buyBalance = numFloat(buyBalance, 4)
              const coin = new web3.eth.Contract(abiContract, this.$store.state.api.eth[tokenName].contractAddress.address)
              let sellBalance = await coin.methods.balanceOf(this.$store.state.account.myEth.address).call()
              let ret = new BigNumber(sellBalance)
              const dec = this.$store.state.api.eth[tokenName].contractAddress.decimal
              sellBalance = parseFloat(ret / Math.pow(10, dec))
              this.sellBalance = numFloat(sellBalance, 4)

              // 清空定时器
              clearInterval(this.buyIntervalId)
              clearInterval(this.sellIntervalId)
              // 开启定时器，每2s获取余额
              this.buyIntervalId = setInterval(async () => {
                let buyBalance = await web3.eth.getBalance(this.$store.state.account.myEth.address) // eth address
                buyBalance = (web3.utils.fromWei(buyBalance, 'ether'))
                this.buyBalance = numFloat(buyBalance, 4)
              }, 2000)
              this.sellIntervalId = setInterval(async () => {
                const coin = new web3.eth.Contract(abiContract, ethAddress)
                let sellBalance = await coin.methods.balanceOf(this.$store.state.account.myEth.address).call()
                let ret = new BigNumber(sellBalance)
                const decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
                const de = 10 ** decimal
                const d = new BigNumber(de)
                sellBalance = parseFloat(ret.dividedBy(d)).toFixed(4)
                this.sellBalance = numFloat(sellBalance, 4)
              }, 2000)
            }
          } else if (code === '3') { // todo 需要处理，获取余额，更新数据
            // erc20 代币bde的合约地址
            const bdeAddress = this.$store.state.api.eth && this.$store.state.api.eth['BDE'].contractAddress.address
            let ethAddress = this.$store.state.api.eth && this.$store.state.api.eth[tokenName].contractAddress.address
            if (bdeAddress && ethAddress) {
              // bde的余额获取和更新
              const bdeCoin = new web3.eth.Contract(abiContract, bdeAddress)
              let buyBalance = await bdeCoin.methods.balanceOf(this.$store.state.account.myEth.address).call()
              let buyBig = new BigNumber(buyBalance)
              const bdeDecimal = this.$store.state.api.eth['BDE'].contractAddress.decimal
              const bdeDe = 10 ** bdeDecimal
              const bdeD = new BigNumber(bdeDe)
              buyBalance = parseFloat(buyBig.dividedBy(bdeD)).toFixed(6) // todo eth数据会全部转化为整数进行处理，所以获取的bde数量需要divid对应的精度
              this.buyBalance = numFloat(buyBalance, 4)
              // bde交易区，其它erc20代币的余额
              const coin = new web3.eth.Contract(abiContract, ethAddress)
              let sellBalance = await coin.methods.balanceOf(this.$store.state.account.myEth.address).call()
              let ret = new BigNumber(sellBalance)
              const decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
              const de = 10 ** decimal
              const d = new BigNumber(de)
              sellBalance = parseFloat(ret.dividedBy(d)).toFixed(4)
              this.sellBalance = numFloat(sellBalance, 4)

              // 清空定时器
              clearInterval(this.buyIntervalId)
              clearInterval(this.sellIntervalId)

              // 5/29 使用定时器2s获取余额；开启定时器
              this.buyIntervalId = setInterval(async () => {
                const bdeCoin = new web3.eth.Contract(abiContract, bdeAddress)
                let buyBalance = await bdeCoin.methods.balanceOf(this.$store.state.account.myEth.address).call()
                let buyBig = new BigNumber(buyBalance)
                const bdeDecimal = this.$store.state.api.eth['BDE'].contractAddress.decimal
                const bdeDe = 10 ** bdeDecimal
                const bdeD = new BigNumber(bdeDe)
                buyBalance = parseFloat(buyBig.dividedBy(bdeD)).toFixed(6) // todo eth数据会全部转化为整数进行处理，所以获取的bde数量需要divid对应的精度
                this.buyBalance = numFloat(buyBalance, 4)
              }, 2000)
              this.sellIntervalId = setInterval(async () => {
                const coin = new web3.eth.Contract(abiContract, ethAddress)
                let sellBalance = await coin.methods.balanceOf(this.$store.state.account.myEth.address).call()
                let ret = new BigNumber(sellBalance)
                const decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
                const de = 10 ** decimal
                const d = new BigNumber(de)
                sellBalance = parseFloat(ret.dividedBy(d)).toFixed(4)
                this.sellBalance = numFloat(sellBalance, 4)
              }, 2000)
            }
          } else if (code === '4') {
            let sellBalance = await web3.eth.getBalance(this.$store.state.account.myEth.address) // eth address
            sellBalance = (web3.utils.fromWei(sellBalance, 'ether'))
            this.sellBalance = numFloat(sellBalance, 4)

            // 清空定时器
            clearInterval(this.sellIntervalId)
            // 开启定时器
            this.sellIntervalId = setInterval(async () => {
              let sellBalance = await web3.eth.getBalance(this.$store.state.account.myEth.address) // eth address
              sellBalance = (web3.utils.fromWei(sellBalance, 'ether'))
              this.sellBalance = numFloat(sellBalance, 4)
            }, 2000)
          }
        } else {
          let param = this.routeParam
          let code = param.split('-')[1]
          if (code === '2') {
            // 清空定时器
            clearInterval(this.buyIntervalId)
            clearInterval(this.sellIntervalId)
            this.buyBalance = 0.0000
            this.sellBalance = 0.0000
          } else if (code === '3') { // todo 需要处理，更新数据
            // 清空定时器
            clearInterval(this.buyIntervalId)
            clearInterval(this.sellIntervalId)
            this.buyBalance = 0.0000
            this.sellBalance = 0.0000
          } else if (code === '4') {
            clearInterval(this.sellIntervalId)
            this.sellBalance = 0.0000
          }
        }
      }
      // immediate: true
    },
    'limitBuyOrder.price': function (newVal) { // 赋值相同不会再次监听，比较的是值；引用类型除外
      this.whIdentity[0] = 1
      let filterVal = newVal
      filterVal = numFloat(filterVal, 6)
      this.limitBuyOrder.total = numFloat(handBig.times(filterVal, this.limitBuyOrder.amount), 4)
    },
    'limitBuyOrder.amount': function (newVal) {
      if (this.whIdentity[2] === 0) { // 自身变动
        this.whIdentity[1] = 1
        let filterVal = newVal
        filterVal = numFloat(filterVal, 4)
        let total = handBig.times(this.limitBuyOrder.price, filterVal)
        this.limitBuyOrder.total = numFloat(total, 4)
      } else { // total变动引起amount变动
        this.whIdentity[2] = 0
      }
    },
    'limitBuyOrder.total': function (newVal) {
      if (this.whIdentity[0] === 1) { // price变动引起total变动
        this.whIdentity[0] = 0
      } else {
        if (this.whIdentity[1] === 0) {
          this.whIdentity[2] = 1
          if (this.limitBuyOrder.price === 0) {
            this.limitBuyOrder.amount = 0.0000
          } else {
            let filterVal = newVal
            filterVal = numFloat(filterVal, 4)
            let amount = handBig.divide(filterVal, this.limitBuyOrder.price)
            this.limitBuyOrder.amount = numFloat(amount, 4) // 保留四位小数
          }
        } else { // amount变动引起total变动
          this.whIdentity[1] = 0
        }
      }
    },
    'limitSellOrder.price': function (newVal) {
      this.whIdentity[3] = 1
      let filterVal = newVal
      filterVal = numFloat(filterVal, 6)
      this.limitSellOrder.total = numFloat(handBig.times(filterVal, this.limitSellOrder.amount), 4)
    },
    'limitSellOrder.amount': function (newVal) {
      if (this.whIdentity[5] === 0) { // 自身变动
        this.whIdentity[4] = 1
        let filterVal = newVal
        filterVal = numFloat(filterVal, 4)
        let total = handBig.times(this.limitSellOrder.price, filterVal)
        this.limitSellOrder.total = numFloat(total, 4)
      } else { // total变动引起amount变动
        this.whIdentity[5] = 0
      }
    },
    'limitSellOrder.total': function (newVal) {
      if (this.whIdentity[3] === 1) { // price变动引起total变动
        this.whIdentity[3] = 0
      } else {
        if (this.whIdentity[4] === 0) {
          this.whIdentity[5] = 1
          if (this.limitSellOrder.price === 0) {
            this.limitSellOrder.amount = 0.0000
          } else {
            let filterVal = newVal
            filterVal = numFloat(filterVal, 4)
            let amount = handBig.divide(filterVal, this.limitSellOrder.price)
            this.limitSellOrder.amount = numFloat(amount, 4) // 保留四位小数
          }
        } else { // amount变动引起total变动
          this.whIdentity[4] = 0
        }
      }
    },
    buyslider: function (newVal) {
      if (newVal !== 0) {
        if (this.limitBuyOrder.price === 0) {
          this.buyslider = 0
        } else if (this.limitBalanceBuy !== 0) {
          let price = this.limitBuyOrder.price
          newVal = newVal / 100
          let total = handBig.times(this.limitBalanceBuy, newVal)
          this.limitBuyOrder.total = numFloat(total, 4)
          let amount = this.limitBuyOrder.total / price
          this.limitBuyOrder.amount = numFloat(amount, 4)
        }
      }
    },
    sellslider: function (newVal) {
      if (newVal !== 0) {
        if (this.limitSellOrder.price === 0) {
          this.sellslider = 0
        } else if (this.limitBalanceSell !== 0) {
          newVal = newVal / 100
          let total = handBig.times(this.limitBalanceSell, newVal)
          this.limitSellOrder.total = numFloat(total, 4)
          let amount = this.limitSellOrder.total
          this.limitSellOrder.amount = numFloat(amount, 4)
        }
      }
    },
    tablePrice: function (newVal) {
      if (newVal !== 0.000000) {
        this.limitBuyOrder.price = newVal
        this.limitSellOrder.price = newVal
        this.$store.commit('orderApi/setPrice', 0.000000)
      }
    },
    // 监听并更新余额 todo 需要新增 bde交易区 的数据监听
    'orderAction.eosOrder': async function (newVal) { // eos下单成功, 更新eos账号余额
      if (newVal !== null && newVal.status) {
        this.$message('您的eos账号下单成功，订单id:' + newVal.orderId)
        let param = this.routeParam
        let tokenName = param.split('-')[0]
        let code = param.split('-')[1]
        let myEos = this.$store.state.account.myEos
        if (myEos) {
          if (code === '1') {
            let eosAddress = this.$store.state.api.eos && this.$store.state.api.eos[tokenName].contractAddress.address
            if (eosAddress) {
              let buyBalance = await getCurrencyBalance('eosio.token', 'EOS') // 测试使用，应该是EOS
              this.buyBalance = parseFloat(buyBalance)
              let sellBalance = await getCurrencyBalance(eosAddress, tokenName)
              this.sellBalance = parseFloat(sellBalance)
            } // 缺少侧链处理
          } else if (code === '4') {
            let buyBalance = await getCurrencyBalance('eosio.token', 'EOS') // 测试使用，应该是EOS
            this.buyBalance = parseFloat(buyBalance)
          }
        }
        this.$store.commit('user/setEosOrder', null)
      } else if (newVal !== null && !newVal.status) {
        this.$message('您的eos账号下单失败，请稍后重试！订单id:' + newVal.orderId)
      }
    },
    'orderAction.ethOrder': async function (newVal) { // eth下单成功, 更新eth账号余额
      if (newVal !== null && newVal.status) {
        this.$message('您的eth账号下单成功，订单id:' + newVal.orderId)
        let param = this.routeParam
        let tokenName = param.split('-')[0]
        let code = param.split('-')[1]
        let myEth = this.$store.state.account.myEth
        if (myEth) {
          if (code === '2') {
            let ethAddress = this.$store.state.api.eth && this.$store.state.api.eth[tokenName].contractAddress.address
            if (ethAddress) {
              let buyBalance = await web3.eth.getBalance(this.$store.state.account.myEth.address) // eth address
              buyBalance = (web3.utils.fromWei(buyBalance, 'ether'))
              this.buyBalance = numFloat(buyBalance, 4)
              const coin = new web3.eth.Contract(abiContract, ethAddress)
              let sellBalance = await coin.methods.balanceOf(this.$store.state.account.myEth.address).call()
              let ret = new BigNumber(sellBalance)
              const decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
              const de = 10 ** decimal
              const d = new BigNumber(de)
              sellBalance = parseFloat(ret.dividedBy(d)).toFixed(4)
              this.sellBalance = numFloat(sellBalance, 4)
            }
          } else if (code === '4') {
            let sellBalance = await web3.eth.getBalance(this.$store.state.account.myEth.address) // eth address
            sellBalance = (web3.utils.fromWei(sellBalance, 'ether'))
            this.sellBalance = numFloat(sellBalance, 4)
          }
        }
        this.$store.commit('user/setEthOrder', null)
      } else if (newVal !== null && !newVal.status) {
        this.$message('您的eth账号下单失败，请等待上一笔交易确认之后在次尝试！订单id:' + newVal.orderId)
      }
    },
    'orderAction.bdeOrder': async function (newVal) { // bde下单成功, 更新eth账号余额
      if (newVal !== null && newVal.status) {
        this.$message('您的bde相关交易下单成功, 订单id:' + newVal.orderId)
        let param = this.routeParam
        let tokenName = param.split('-')[0]
        let code = param.split('-')[1]
        let myEth = this.$store.state.account.myEth
        if (myEth) {
          if (code === '3') {
            // erc20 代币bde的合约地址
            const bdeAddress = this.$store.state.api.eth && this.$store.state.api.eth['BDE'].contractAddress.address
            let ethAddress = this.$store.state.api.eth && this.$store.state.api.eth[tokenName].contractAddress.address
            if (bdeAddress && ethAddress) {
              // bde的余额获取和更新
              const bdeCoin = new web3.eth.Contract(abiContract, bdeAddress)
              let buyBalance = await bdeCoin.methods.balanceOf(this.$store.state.account.myEth.address).call()
              let buyBig = new BigNumber(buyBalance)
              const bdeDecimal = this.$store.state.api.eth['BDE'].contractAddress.decimal
              const bdeDe = 10 ** bdeDecimal
              const bdeD = new BigNumber(bdeDe)
              buyBalance = parseFloat(buyBig.dividedBy(bdeD)).toFixed(6) // todo eth数据会全部转化为整数进行处理，所以获取的bde数量需要divid对应的精度
              this.buyBalance = numFloat(buyBalance, 4)
              // bde交易区，其它erc20代币的余额
              const coin = new web3.eth.Contract(abiContract, ethAddress)
              let sellBalance = await coin.methods.balanceOf(this.$store.state.account.myEth.address).call()
              let ret = new BigNumber(sellBalance)
              const decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
              const de = 10 ** decimal
              const d = new BigNumber(de)
              sellBalance = parseFloat(ret.dividedBy(d)).toFixed(4)
              this.sellBalance = numFloat(sellBalance, 4)
            }
          }
        }
        this.$store.commit('user/setBdeOrder', null)
      } else if (newVal !== null && !newVal.status) {
        this.$message('您的bde相关交易下单失败，请等待上一笔交易确认之后再次尝试！订单id:' + newVal.orderId)
      }
    },
    'orderAction.eosWithdraw': async function (newVal) { // eos撤单成功, 更新eos账号余额
      if (newVal !== null && newVal.status) {
        this.$message('您的eos账号撤单成功，订单id:' + newVal.orderId)
        let param = this.routeParam
        let tokenName = param.split('-')[0]
        let code = param.split('-')[1]
        let myEos = this.$store.state.account.myEos
        if (myEos) {
          if (code === '1') {
            let eosAddress = this.$store.state.api.eos && this.$store.state.api.eos[tokenName].contractAddress.address
            if (eosAddress) {
              let buyBalance = await getCurrencyBalance('eosio.token', 'EOS') // 测试使用，应该是EOS
              this.buyBalance = parseFloat(buyBalance)
              let sellBalance = await getCurrencyBalance(eosAddress, tokenName)
              this.sellBalance = parseFloat(sellBalance)
            }
          } else if (code === '4') {
            let buyBalance = await getCurrencyBalance('eosio.token', 'EOS') // 测试使用，应该是EOS
            this.buyBalance = parseFloat(buyBalance)
          }
        }
        this.$store.commit('user/setEosWithdraw', null)
      } else if (newVal !== null && !newVal.status) {
        this.$message('您的eos账号撤单失败，请稍后重试！订单id:' + newVal.orderId)
      }
    },
    'orderAction.ethWithdraw': async function (newVal) { // eth撤单成功, 更新eth账号余额
      if (newVal !== null && newVal.status) {
        this.$message('您的eth账号撤单成功，订单id:' + newVal.orderId)
        let param = this.routeParam
        let tokenName = param.split('-')[0]
        let code = param.split('-')[1]
        let myEth = this.$store.state.account.myEth
        if (myEth) {
          if (code === '2') {
            let ethAddress = this.$store.state.api.eth && this.$store.state.api.eth[tokenName].contractAddress.address
            if (ethAddress) {
              let buyBalance = await web3.eth.getBalance(this.$store.state.account.myEth.address) // eth address
              buyBalance = (web3.utils.fromWei(buyBalance, 'ether'))
              this.buyBalance = numFloat(buyBalance, 4)
              const coin = new web3.eth.Contract(abiContract, ethAddress)
              let sellBalance = await coin.methods.balanceOf(this.$store.state.account.myEth.address).call()
              let ret = new BigNumber(sellBalance)
              const decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
              const de = 10 ** decimal
              const d = new BigNumber(de)
              sellBalance = parseFloat(ret.dividedBy(d)).toFixed(4)
              this.sellBalance = numFloat(sellBalance, 4)
            }
          } else if (code === '4') {
            let sellBalance = await web3.eth.getBalance(this.$store.state.account.myEth.address) // eth address
            sellBalance = (web3.utils.fromWei(sellBalance, 'ether'))
            this.sellBalance = numFloat(sellBalance, 4)
          }
        }
        this.$store.commit('user/setEthWithdraw', null)
      } else if (newVal !== null && !newVal.status) {
        this.$message('您的eth账号撤单失败，请稍后重试！订单id:' + newVal.orderId)
      }
    },
    'orderAction.bdeWithdraw': async function (newVal) { // eth撤单成功, 更新eth账号余额
      if (newVal !== null && newVal.status) {
        this.$message('您的bde相关交易撤单成功，订单id:' + newVal.orderId)
        let param = this.routeParam
        let tokenName = param.split('-')[0]
        let code = param.split('-')[1]
        let myEth = this.$store.state.account.myEth
        if (myEth) {
          if (code === '3') {
            // erc20 代币bde的合约地址
            const bdeAddress = this.$store.state.api.eth && this.$store.state.api.eth['BDE'].contractAddress.address
            let ethAddress = this.$store.state.api.eth && this.$store.state.api.eth[tokenName].contractAddress.address
            if (bdeAddress && ethAddress) {
              // bde的余额获取和更新
              const bdeCoin = new web3.eth.Contract(abiContract, bdeAddress)
              let buyBalance = await bdeCoin.methods.balanceOf(this.$store.state.account.myEth.address).call()
              let buyBig = new BigNumber(buyBalance)
              const bdeDecimal = this.$store.state.api.eth['BDE'].contractAddress.decimal
              const bdeDe = 10 ** bdeDecimal
              const bdeD = new BigNumber(bdeDe)
              buyBalance = parseFloat(buyBig.dividedBy(bdeD)).toFixed(6) // todo eth数据会全部转化为整数进行处理，所以获取的bde数量需要divid对应的精度
              this.buyBalance = numFloat(buyBalance, 4)
              // bde交易区，其它erc20代币的余额
              const coin = new web3.eth.Contract(abiContract, ethAddress)
              let sellBalance = await coin.methods.balanceOf(this.$store.state.account.myEth.address).call()
              let ret = new BigNumber(sellBalance)
              const decimal = this.$store.state.api.eth[tokenName].contractAddress.decimal
              const de = 10 ** decimal
              const d = new BigNumber(de)
              sellBalance = parseFloat(ret.dividedBy(d)).toFixed(4)
              this.sellBalance = numFloat(sellBalance, 4)
            }
          }
        }
        this.$store.commit('user/setBdeWithdraw', null)
      } else if (newVal !== null && !newVal.status) {
        this.$message('您的bde相关交易撤单失败，请稍后重试！订单id:' + newVal.orderId)
      }
    },
    accreditamount: function (newVal) {
      if (newVal > 0 && newVal <= parseFloat(this.sellBalance)) { // 授权额度是否符合要求
        // 精度判断，是否满足交易手续费规则
        let param = this.routeParam
        let tokenName = param.split('-')[0]
        let eth = this.$store.state.api.eth
        let filDecimal = eth[tokenName].contractAddress.decimal
        let filQuantity = numFloat(newVal, 4)
        if (filDecimal > 8) {
          filDecimal = 8
        }

        let len = newVal.toString().split('.').length
        if (len === 2) {
          if (newVal.toString().split('.')[1].length > filDecimal) {
            return
          }
        }
        if ((filQuantity * (10 ** filDecimal)) < 2) {
          return
        }
        this.showApprove = false
      } else {
        this.showApprove = true
      }
    },
    accreditamountBde: function (newVal) {
      if (newVal > 0 && newVal <= parseFloat(this.buyBalance)) { // 授权额度是否符合要求
        // 精度判断，是否满足交易手续费规则
        let tokenName = 'BDE'
        let eth = this.$store.state.api.eth
        let filDecimal = eth[tokenName].contractAddress.decimal
        let filQuantity = numFloat(newVal, 4)
        if (filDecimal < 4) {
          filQuantity = numFloat(newVal, parseInt(filDecimal))
        }
        if (filDecimal > 8) {
          filDecimal = 8
        }

        let len = newVal.toString().split('.').length
        if (len === 2) {
          if (newVal.toString().split('.')[1].length > filDecimal) {
            return
          }
        }
        if ((filQuantity * (10 ** filDecimal)) < 2) {
          return
        }
        this.showApproveBde = false
      } else {
        this.showApproveBde = true
      }
    }
  },
  methods: {
    async buyOrder () { // 下单 买单
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      let code = param.split('-')[1]
      if (code === '1') { // eos  买单
        let myEos = this.$store.state.account.myEos
        let isConnected = this.$store.state.isConnected
        let idArr = this.$store.state.user.idArr
        let eos = this.$store.state.api.eos
        if (myEos === null || !isConnected || idArr === null || eos === null) {
          this.$message('未登录或者网络故障, 请重新加载')
        } else {
          if (this.limitBuyOrder.total <= this.limitBalanceBuy && this.limitBuyOrder.total > 0) {
            const contractAccount = 'eosio.token'
            const myContract = myEosContract
            // 精度判断，是否满足交易手续费规则
            // let filQuantity = (Math.floor(this.limitBuyOrder.total * 10000) / 10000)
            let filQuantity = numFloat(this.limitBuyOrder.total, 4)
            let filDecimal = 4
            if ((filQuantity * (10 ** filDecimal)) < 2) {
              this.$message('交易额过低，无法交易！')
              return
            }

            // 新增 5/30 shj
            let currentFil = eos[tokenName].contractAddress.decimal
            const len1 = this.limitBuyOrder.amount.toString().split('.').length
            const len2 = this.limitBuyOrder.total.toString().split('.').length
            if (len1 === 2) {
              if (currentFil >= 4) {
                if (this.limitBuyOrder.amount.toString().split('.')[1].length > currentFil) {
                  this.$message('交易精度过低，无法交易！')
                  return
                }
              }
            }
            if (len2 === 2) {
              if (this.limitBuyOrder.total.toString().split('.')[1].length > filDecimal) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }

            // const quantity = (Math.floor(this.limitBuyOrder.total * 10000) / 10000).toFixed(4) + ' EOS'
            const quantity = numFloat(this.limitBuyOrder.total, 4) + ' EOS'
            const id = idArr[0] // string
            // 更新订单id
            if (idArr.length === 1) {
              this.$socket.sendObj({
                namespace: 'user',
                action: 'getid',
                data: ''
              })
            }
            this.$store.commit('user/delId', 0)
            const a = '0'
            // let price1 = Math.floor(this.limitBuyOrder.price * 1000000) / 1000000
            let price = numFloat(this.limitBuyOrder.price, 6)
            const p = handBig.times(price, 10 ** 8) + ''
            let decimal = eos[tokenName].contractAddress.decimal + ''
            const bs = tokenName + ',' + decimal
            const bc = eos[tokenName].contractAddress.address + ''
            const ta = this.$store.state.account.myEos.identity.accounts[0].name
            const memo = `{"i": "${id}", "a": "${a}", "p": "${p}", "bs": "${bs}", "bc": "${bc}", "ta": "${ta}"}` // 先不加"", 若是不行再改
            let trxSign = await getEosTransferSign(contractAccount, myContract, quantity, memo)
            if (trxSign) {
              this.$socket.sendObj({
                namespace: 'user',
                action: 'sendeos',
                data: {
                  order: JSON.stringify(trxSign),
                  marketCode: '1',
                  chainCode: '1',
                  targetCode: '1'
                }
              })
              this.$message('正在下单，请稍等....................')
            } else {
              // this.$message('放弃下单')
            }
          } else {
            this.$message('请输入正确的订单数据===========')
          }
        }
      } else if (code === '2') { // eth 买单
        let myEth = this.$store.state.account.myEth // 删除.address,shj618
        let isConnected = this.$store.state.isConnected
        let idArr = this.$store.state.user.idArr
        let eth = this.$store.state.api.eth
        if (myEth === null || !isConnected || idArr === null || eth === null) {
          this.$message('未登录或者网络故障, 请重新加载')
        } else if (myEth === null && sessionStorage.getItem('myEth')) {
          this.$message('请解锁钱包')
        } else {
          if (this.limitBuyOrder.total <= this.limitBalanceBuy && this.limitBuyOrder.total > 0) {
            // let filQuantity = (Math.floor(this.limitBuyOrder.total * 10000) / 10000)
            let filQuantity = numFloat(this.limitBuyOrder.total, 4)
            let filDecimal = 8

            // 新增5/30，shj
            let currentDe = eth[tokenName].contractAddress.decimal
            const len1 = this.limitBuyOrder.total.toString().split('.').length
            const len2 = this.limitBuyOrder.amount.toString().split('.').length
            if (len1 === 2) {
              if (this.limitBuyOrder.total.toString().split('.')[1].length > filDecimal) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }
            if (len2 === 2 && currentDe >= 4) { // todo 新增，精度低于4的处理
              if (this.limitBuyOrder.amount.toString().split('.')[1].length > currentDe) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }

            if ((filQuantity * (10 ** filDecimal)) < 2) {
              this.$message('交易额过低，无法交易！')
              return
            }
            // 更新订单id
            if (idArr.length === 1) {
              this.$socket.sendObj({
                namespace: 'user',
                action: 'getid',
                data: ''
              })
            }
            const orderid = idArr[0]
            this.$store.commit('user/delId', 0)
            const useraddress = this.$store.state.account.myEth.address
            const receiver = this.$store.state.account.myEth.address
            const tradetype = false // 以太坊链内交易买单为false，跨链为true
            let sellamount = await web3.utils.toWei(this.limitBuyOrder.total.toString(), 'ether')
            const price = handleEthPrice(this.limitBuyOrder.price) // 6.18shj新增
            if (!price) {
              this.$message('请输入正确的订单数据===========')
              return
            }
            // const buyaddress = eth[tokenName].contractAddress.address // err
            // const dateTime = Date.now()
            // const createTime = Math.floor(dateTime / 1000)
            const sellTokenAddress = '0x0000000000000000000000000000000000000000'
            const buyTokenAddress = eth[tokenName].contractAddress.address
            const contractdata = await exchange.methods.depositsOrder(web3.utils.toHex(orderid), receiver, tradetype, web3.utils.toHex(sellamount), sellTokenAddress, price, buyTokenAddress, 2)
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
              value: web3.utils.toHex(sellamount),
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
              action: 'sendeth',
              data: {
                approveTx: '0x' + serializedTx.toString('hex'),
                marketType: 2,
                tradetype: false
              }
            })
            this.$message('正在下单，请稍等....................')
          } else {
            this.$message('请输入正确的订单数据===========')
          }
        }
      } else if (code === '3') { // todo 需要处理，bde买单
        // bde的精度是2位，price需要截取
        let myEth = this.$store.state.account.myEth
        let isConnected = this.$store.state.isConnected
        let idArr = this.$store.state.user.idArr
        let eth = this.$store.state.api.eth
        if (myEth === null || !isConnected || eth === null) {
          this.$message('未登录或者网络故障, 请重新加载')
        } else {
          // 授权额度，判断
          const contractAccountBde = eth['BDE'].contractAddress.address // bde的合约地址
          // const contractAccount = eth[tokenName].contractAddress.address
          const myContract = exchange.address
          const useraddress = myEth.address

          const coin = new web3.eth.Contract(abiContract, contractAccountBde)
          let quota = await coin.methods.allowance(useraddress, myContract).call() // 授权额度
          let approveQuota = new BigNumber(quota) // todo 改动
          console.log('查询bde的授权额度-------返回值', approveQuota)
          const dec = eth['BDE'].contractAddress.decimal
          console.log('bde的精度====================', dec)
          this.haveAccreditAmountBde = parseFloat(approveQuota / Math.pow(10, dec)) // 不会报错吗 todo 这里需要修改，使用专门的变量存储bde授权额度
          console.log('bde买单，获取已经授权的额度----------------', this.haveAccreditAmountBde)
          if (this.limitBuyOrder.total > this.haveAccreditAmountBde) {
            this.$message('交易额度超过授权额度，请确定授权额度大于交易额度')
            return
          }
          if (this.limitBuyOrder.total <= this.limitBalanceBuy && this.limitBuyOrder.total > 0) {
            const oriPrice = this.limitBuyOrder.price
            const oriAmount = this.limitBuyOrder.amount
            const oriTotal = this.limitBuyOrder.total
            // todo 用于处理精度低于4的erc20代币
            // 精度判断，是否满足交易手续费规则, bde精度判断
            let bdeDecimal = eth[tokenName].contractAddress.decimal
            if (bdeDecimal > 8) {
              bdeDecimal = 8
            }
            let filQuantityBde = numFloat(oriTotal, 4) // todo 下单时强制使用两位精度数据，bde精度2位
            if (bdeDecimal < 4) {
              filQuantityBde = numFloat(oriTotal, bdeDecimal) // todo 下单时强制使用两位精度数据，bde精度2位
            }
            // erc20代币 精度判断
            let filDecimal = eth[tokenName].contractAddress.decimal // bde交易区erc20代币精度
            // let filQuantity = (Math.floor(this.limitSellOrder.amount * 10000) / 10000)
            if (filDecimal > 8) {
              filDecimal = 8
            }
            // todo 新增处理，代币精度小于4的情况处理
            // 5/30新增，shj
            const len1 = oriAmount.toString().split('.').length
            const len2 = oriTotal.toString().split('.').length
            if (len1 === 2 && filDecimal >= 4) {
              if (oriAmount.toString().split('.')[1].length > filDecimal) { // 判断erc20代币精度
                this.$message('交易精度过低，无法交易！')
                return
              }
            }
            if (len2 === 2 && bdeDecimal >= 4) { // 精度低于4的代币不检查
              if (oriTotal.toString().split('.')[1].length > bdeDecimal) { // 判断bde精度
                this.$message('交易精度过低，无法交易！')
                return
              }
            }

            if ((filQuantityBde * (10 ** bdeDecimal)) < 2) {
              this.$message('交易额过低，无法交易！')
              return
            }

            const orderId = idArr[0]
            // 更新订单id
            if (idArr.length === 1) {
              this.$socket.sendObj({
                namespace: 'user',
                action: 'getid',
                data: ''
              })
            }
            let total = 0
            if (bdeDecimal < 4) {
              let lowFloatTotal = numFloat(oriTotal, bdeDecimal)
              total = new BigNumber(lowFloatTotal * (10 ** eth['BDE'].contractAddress.decimal))
            } else {
              total = new BigNumber(this.limitBuyOrder.total * (10 ** eth['BDE'].contractAddress.decimal))
            }
            const price = handleEthPrice(oriPrice) // shj新增6.18;todo price放大8倍，暂使用之前的处理
            if (!price) {
              this.$message('请输入正确的订单数据===========')
              return
            }
            const tradeType = false // 1 || true : 卖单 0 || false: 买单
            const useraddress = this.$store.state.account.myEth.address
            // const datetime = Date.now()
            // const createtime = Math.floor(datetime / 1000)
            const gas = await web3.eth.getGasPrice()
            const nonce = await web3.eth.getTransactionCount(useraddress)
            const gasPriceHex = web3.utils.numberToHex(gas)
            const nonceHex = web3.utils.numberToHex(nonce)
            const gasLimitHex = web3.utils.numberToHex(400000)
            // todo 暂时使用，后面会改接口
            const sellTokenAddress = eth['BDE'].contractAddress.address
            const buyTokenAddress = eth[tokenName].contractAddress.address
            const contractdata = await exchange.methods.depositsOrder(web3.utils.toHex(orderId), useraddress, tradeType, web3.utils.toHex(total), sellTokenAddress, price, buyTokenAddress, 3)
            const payload = contractdata.encodeABI()

            const rawTx = {
              nonce: nonceHex,
              gasPrice: gasPriceHex,
              gasLimit: gasLimitHex,
              value: '0x00',
              from: useraddress,
              to: myContract,
              data: payload
            }
            const tx = new Tx(rawTx)
            const privateKey = Buffer.from(this.$store.state.account.myEth.privateKey.slice(2), 'hex')
            tx.sign(privateKey)
            const serializedTx = tx.serialize()
            this.$socket.sendObj({
              namespace: 'user',
              action: 'sendbde',
              data: {
                tradeType: false, // 1 || true : 卖单 0 || false: 买单
                marketType: 3, // 对应于marketCode
                approveTx: '0x' + serializedTx.toString('hex')
              }
            })
            this.$message('正在下单, 请稍后...')
          } else {
            this.$message('请输入正确的订单数据')
          }
        }
      } else if (code === '4') { // 跨链 买单 eos
        // eos
        let myEos = this.$store.state.account.myEos
        let myEth = this.$store.state.account.myEth
        let isConnected = this.$store.state.isConnected
        let idArr = this.$store.state.user.idArr
        let eos = this.$store.state.api.eos
        if (myEos === null || !isConnected || idArr === null || eos === null) {
          this.$message('未登录或者网络故障, 请重新加载')
        } else if (myEth === null) {
          this.$message('跨链交易，请登录Eth账号')
        } else {
          if (this.limitBuyOrder.total <= this.limitBalanceBuy && this.limitBuyOrder.total > 0) {
            const contractAccount = 'eosio.token'
            const myContract = myEosContract
            // 精度判断，是否满足交易手续费规则
            // let filQuantity = (Math.floor(this.limitBuyOrder.total * 10000) / 10000)
            let filQuantity = numFloat(this.limitBuyOrder.total, 4)
            let filDecimal = 4
            if ((filQuantity * (10 ** filDecimal)) < 2) {
              this.$message('交易额过低，无法交易！')
              return
            }

            // 新增 5/30 shj
            let ethDecimal = 8
            const len1 = this.limitBuyOrder.amount.toString().split('.').length
            const len2 = this.limitBuyOrder.total.toString().split('.').length
            if (len1 === 2) {
              if (this.limitBuyOrder.amount.toString().split('.')[1].length > ethDecimal) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }
            if (len2 === 2) {
              if (this.limitBuyOrder.total.toString().split('.')[1].length > filDecimal) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }

            // const quantity = (Math.floor(this.limitBuyOrder.total * 10000) / 10000).toFixed(4) + ' EOS'
            const quantity = (numFloat(this.limitBuyOrder.total, 4)) + ' EOS'
            const id = idArr[0] // string
            // 更新订单id
            if (idArr.length === 1) {
              this.$socket.sendObj({
                namespace: 'user',
                action: 'getid',
                data: ''
              })
            }
            this.$store.commit('user/delId', 0)
            const a = '0'
            // let price = Math.floor(this.limitBuyOrder.price * 1000000) / 1000000
            let price = numFloat(this.limitBuyOrder.price, 6)
            const p = handBig.times(price, 10 ** 8) + ''
            // let decimal = eos[tokenName].contractAddress.decimal + ''
            const bs = 'ETH' + ',' + '8'
            const bc = '0x0000000000000000000000000000000000000000'
            const ta = myEth.address // 这里需要加一个审核判断，eth是否已登录
            const memo = `{"i": "${id}", "a": "${a}", "p": "${p}", "bs": "${bs}", "bc": "${bc}", "ta": "${ta}"}` // 先不加"", 若是不行再改
            let trxSign = await getEosTransferSign(contractAccount, myContract, quantity, memo)
            if (trxSign) {
              this.$socket.sendObj({
                namespace: 'user',
                action: 'sendeos',
                data: {
                  order: JSON.stringify(trxSign),
                  marketCode: '4',
                  chainCode: '1',
                  targetCode: '2'
                }
              })
              this.$message('正在下单，请稍等....................')
            } else {

            }
          } else {
            this.$message('请输入正确的订单数据===========')
          }
        }
      }
    },
    async sellOrder () { // 下单 卖单
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      let code = param.split('-')[1]
      if (code === '1') {
        let myEos = this.$store.state.account.myEos
        let isConnected = this.$store.state.isConnected
        let idArr = this.$store.state.user.idArr
        let eos = this.$store.state.api.eos
        if (myEos === null || !isConnected || idArr === null || eos === null) {
          this.$message('未登录或者网络故障, 请重新加载')
        } else {
          if (this.limitSellOrder.amount <= this.limitBalanceSell && this.limitSellOrder.total > 0) {
            const contractAccount = eos[tokenName].contractAddress.address
            const myContract = myEosContract
            // 精度判断，是否满足交易手续费规则
            let filDecimal = eos[tokenName].contractAddress.decimal
            // let filQuantity = (Math.floor(this.limitSellOrder.amount * 10000) / 10000)
            let filQuantity = numFloat(this.limitSellOrder.amount, 4)
            if (filDecimal > 8) {
              filDecimal = 8
            }
            if ((filQuantity * (10 ** filDecimal)) < 2) {
              this.$message('交易额过低，无法交易！')
              return
            }

            // 新增 5/30 shj
            let eosDecimal = 4
            const len1 = this.limitSellOrder.amount.toString().split('.').length
            const len2 = this.limitSellOrder.total.toString().split('.').length
            if (len1 === 2) {
              if (this.limitSellOrder.amount.toString().split('.')[1].length > filDecimal) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }
            if (len2 === 2) {
              if (this.limitSellOrder.total.toString().split('.')[1].length > eosDecimal) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }

            const decimal = eos[tokenName].contractAddress.decimal
            // const quantity = (Math.floor(this.limitSellOrder.amount * 10000) / 10000).toFixed(decimal) + ' ' + tokenName
            const quantity = (numFloat(this.limitSellOrder.amount, decimal)) + ' ' + tokenName
            const id = idArr[0]
            // 更新订单id
            if (idArr.length === 1) {
              this.$socket.sendObj({
                namespace: 'user',
                action: 'getid',
                data: ''
              })
            }
            this.$store.commit('user/delId', 0)
            const a = '1' // '0' buy '1' sell
            // let price = Math.floor(this.limitSellOrder.price * 1000000) / 1000000
            let price = numFloat(this.limitSellOrder.price, 6)
            const p = handBig.times(price, 10 ** 8) + ''
            const bs = 'EOS,4' // test,
            const bc = 'eosio.token'
            const ta = this.$store.state.account.myEos.identity.accounts[0].name
            const memo = `{"i": "${id}", "a": "${a}", "p": "${p}", "bs": "${bs}", "bc": "${bc}", "ta": "${ta}"}` // 先不加"", 若是不行再改
            let trxSign = await getEosTransferSign(contractAccount, myContract, quantity, memo)
            if (trxSign) {
              this.$socket.sendObj({
                namespace: 'user',
                action: 'sendeos',
                data: {
                  order: JSON.stringify(trxSign),
                  marketCode: '1',
                  chainCode: '1',
                  targetCode: '1'
                }
              })
              this.$message('正在下单, 请稍后...')
            } else {
              // this.$message('取消下单')
            }
          } else {
            this.$message('请输入正确的订单数据')
          }
        }
      } else if (code === '2') { // eth 卖单
        let myEth = this.$store.state.account.myEth
        let isConnected = this.$store.state.isConnected
        let idArr = this.$store.state.user.idArr
        let eth = this.$store.state.api.eth
        if (myEth === null || !isConnected || eth === null) {
          this.$message('未登录或者网络故障, 请重新加载')
        } else {
          const contractAccount = eth[tokenName].contractAddress.address
          const myContract = exchange.address
          const useraddress = myEth.address

          const coin = new web3.eth.Contract(abiContract, contractAccount)
          let quota = await coin.methods.allowance(useraddress, myContract).call()
          let approveQuota = new BigNumber(quota)
          const dec = eth[tokenName].contractAddress.decimal
          this.haveAccreditAmount = parseFloat(approveQuota / Math.pow(10, dec))

          if (this.limitSellOrder.amount > this.haveAccreditAmount) {
            this.$message('交易额度超过授权额度，请确定授权额度大于交易额度')
            return
          }
          if (this.limitSellOrder.amount <= this.limitBalanceSell && this.limitSellOrder.total > 0) {
            // 精度判断，是否满足交易手续费规则
            let filDecimal = eth[tokenName].contractAddress.decimal
            // let filQuantity = (Math.floor(this.limitSellOrder.amount * 10000) / 10000)
            let filQuantity = numFloat(this.limitSellOrder.amount, 4)
            if (filDecimal < 4) {
              filQuantity = numFloat(this.limitSellOrder.amount, filDecimal)
            }
            if (filDecimal > 8) {
              filDecimal = 8
            }

            // 5/30新增，shj
            let ethDecimal = 8
            const len1 = this.limitSellOrder.amount.toString().split('.').length
            const len2 = this.limitSellOrder.total.toString().split('.').length
            if (len1 === 2 && filDecimal >= 4) {
              if (this.limitSellOrder.amount.toString().split('.')[1].length > filDecimal) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }
            if (len2 === 2) {
              if (this.limitSellOrder.total.toString().split('.')[1].length > ethDecimal) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }

            if ((filQuantity * (10 ** filDecimal)) < 2) {
              this.$message('交易额过低，无法交易！')
              return
            }

            const orderId = idArr[0]
            // 更新订单id
            if (idArr.length === 1) {
              this.$socket.sendObj({
                namespace: 'user',
                action: 'getid',
                data: ''
              })
            }
            this.$store.commit('user/delId', 0)
            let amount = new BigNumber(this.limitSellOrder.amount * (10 ** eth[tokenName].contractAddress.decimal))
            if (filDecimal < 4) { // todo 8.14
              let lowFloatAmount = numFloat(this.limitSellOrder.amount, filDecimal)
              amount = new BigNumber(lowFloatAmount * (10 ** eth[tokenName].contractAddress.decimal))
            }
            const price = handleEthPrice(this.limitSellOrder.price) // shj新增6.18
            if (!price) {
              this.$message('请输入正确的订单数据===========')
              return
            }
            const tradeType = true
            const useraddress = this.$store.state.account.myEth.address
            // const datetime = Date.now()
            // const createtime = Math.floor(datetime / 1000)
            const gas = await web3.eth.getGasPrice()
            const nonce = await web3.eth.getTransactionCount(useraddress)
            const gasPriceHex = web3.utils.numberToHex(gas)
            const nonceHex = web3.utils.numberToHex(nonce)
            const gasLimitHex = web3.utils.numberToHex(400000)
            const sellTokenAddress = eth[tokenName].contractAddress.address
            const buyTokenAddress = '0x0000000000000000000000000000000000000000'
            const contractdata = await exchange.methods.depositsOrder(web3.utils.toHex(orderId), useraddress, tradeType, web3.utils.toHex(amount), sellTokenAddress, price, buyTokenAddress, 2)
            const payload = contractdata.encodeABI()

            const rawTx = {
              nonce: nonceHex,
              gasPrice: gasPriceHex,
              gasLimit: gasLimitHex,
              value: '0x00',
              from: useraddress,
              to: myContract,
              data: payload
            }
            const tx = new Tx(rawTx)
            const privateKey = Buffer.from(this.$store.state.account.myEth.privateKey.slice(2), 'hex')
            tx.sign(privateKey)
            const serializedTx = tx.serialize()
            this.$socket.sendObj({
              namespace: 'user',
              action: 'sendeth',
              data: {
                tradeType: true,
                marketType: 2,
                approveTx: '0x' + serializedTx.toString('hex')
              }
            })
            this.$message('正在下单, 请稍后...')
          } else {
            this.$message('请输入正确的订单数据')
          }
        }
      } else if (code === '3') { // 同构跨链 todo 需要处理bde卖单，8.13
        let myEth = this.$store.state.account.myEth
        let isConnected = this.$store.state.isConnected
        let idArr = this.$store.state.user.idArr
        let eth = this.$store.state.api.eth
        if (myEth === null || !isConnected || eth === null) {
          this.$message('未登录或者网络故障, 请重新加载')
        } else {
          const contractAccount = eth[tokenName].contractAddress.address
          const myContract = exchange.address
          const useraddress = myEth.address

          const coin = new web3.eth.Contract(abiContract, contractAccount)
          let quota = await coin.methods.allowance(useraddress, myContract).call()
          let approveQuota = new BigNumber(quota)
          const dec = eth[tokenName].contractAddress.decimal
          this.haveAccreditAmount = parseFloat(approveQuota / Math.pow(10, dec))

          if (this.limitSellOrder.amount > this.haveAccreditAmount) {
            this.$message('交易额度超过授权额度，请确定授权额度大于交易额度')
            return
          }
          if (this.limitSellOrder.amount <= this.limitBalanceSell && this.limitSellOrder.total > 0) {
            // 精度判断，是否满足交易手续费规则
            let filDecimal = eth[tokenName].contractAddress.decimal
            // let filQuantity = (Math.floor(this.limitSellOrder.amount * 10000) / 10000)
            let filQuantity = numFloat(this.limitSellOrder.amount, 4)
            if (filDecimal < 4) {
              filQuantity = numFloat(this.limitSellOrder.amount, filDecimal)
            }
            if (filDecimal > 8) {
              filDecimal = 8
            }

            // 5/30新增，shj
            let bdeDecimal = eth['BDE'].contractAddress.decimal
            const len1 = this.limitSellOrder.amount.toString().split('.').length
            const len2 = this.limitSellOrder.total.toString().split('.').length
            if (len1 === 2 && filDecimal >= 4) {
              if (this.limitSellOrder.amount.toString().split('.')[1].length > filDecimal) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }
            if (len2 === 2 && bdeDecimal >= 4) {
              if (this.limitSellOrder.total.toString().split('.')[1].length > bdeDecimal) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }

            if ((filQuantity * (10 ** filDecimal)) < 2) {
              this.$message('交易额过低，无法交易！')
              return
            }

            const orderId = idArr[0]
            // 更新订单id
            if (idArr.length === 1) {
              this.$socket.sendObj({
                namespace: 'user',
                action: 'getid',
                data: ''
              })
            }
            this.$store.commit('user/delId', 0)
            let amount = new BigNumber(this.limitSellOrder.amount * (10 ** eth[tokenName].contractAddress.decimal))
            if (filDecimal < 4) { // todo 8.14
              let lowFloatAmount = numFloat(this.limitSellOrder.amount, filDecimal)
              amount = new BigNumber(lowFloatAmount * (10 ** eth[tokenName].contractAddress.decimal))
            }
            const price = handleEthPrice(this.limitSellOrder.price) // shj新增6.18
            if (!price) {
              this.$message('请输入正确的订单数据===========')
              return
            }
            const tradeType = true
            const useraddress = this.$store.state.account.myEth.address
            // const datetime = Date.now()
            // const createtime = Math.floor(datetime / 1000)
            const gas = await web3.eth.getGasPrice()
            const nonce = await web3.eth.getTransactionCount(useraddress)
            const gasPriceHex = web3.utils.numberToHex(gas)
            const nonceHex = web3.utils.numberToHex(nonce)
            const gasLimitHex = web3.utils.numberToHex(400000)
            // todo 接口暂时没确定，待定
            const sellTokenAddress = eth[tokenName].contractAddress.address
            const buyTokenAddress = eth['BDE'].contractAddress.address
            const contractdata = await exchange.methods.depositsOrder(web3.utils.toHex(orderId), useraddress, tradeType, web3.utils.toHex(amount), sellTokenAddress, price, buyTokenAddress, 3)
            const payload = contractdata.encodeABI()

            const rawTx = {
              nonce: nonceHex,
              gasPrice: gasPriceHex,
              gasLimit: gasLimitHex,
              value: '0x00',
              from: useraddress,
              to: myContract,
              data: payload
            }
            const tx = new Tx(rawTx)
            const privateKey = Buffer.from(this.$store.state.account.myEth.privateKey.slice(2), 'hex')
            tx.sign(privateKey)
            const serializedTx = tx.serialize()
            this.$socket.sendObj({
              namespace: 'user',
              action: 'sendbde',
              data: {
                tradeType: true,
                marketType: 3,
                approveTx: '0x' + serializedTx.toString('hex')
              }
            })
            this.$message('正在下单, 请稍后...')
          } else {
            this.$message('请输入正确的订单数据')
          }
        }
      } else if (code === '4') { // 非同构跨链 eth
        let myEos = this.$store.state.account.myEos
        let myEth = this.$store.state.account.myEth
        let isConnected = this.$store.state.isConnected
        let idArr = this.$store.state.user.idArr
        let eth = this.$store.state.api.eth
        if (myEth === null || !isConnected || idArr === null || eth === null) {
          this.$message('未登录或者网络故障, 请重新加载')
        } else if (myEos === null) {
          this.$message('跨链交易，请登录Eos账号')
        } else {
          if (this.limitSellOrder.amount <= this.limitBalanceSell && this.limitSellOrder.total > 0) {
            // let filQuantity = (Math.floor(this.limitSellOrder.total * 10000) / 10000)
            let filQuantity = numFloat(this.limitSellOrder.total, 4)
            let filDecimal = 8

            // 5/30新增，shj
            let eosDecimal = 4
            const len1 = this.limitSellOrder.amount.toString().split('.').length
            const len2 = this.limitSellOrder.total.toString().split('.').length
            if (len1 === 2) {
              if (this.limitSellOrder.amount.toString().split('.')[1].length > filDecimal) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }
            if (len2 === 2) {
              if (this.limitSellOrder.total.toString().split('.')[1].length > eosDecimal) {
                this.$message('交易精度过低，无法交易！')
                return
              }
            }

            if ((filQuantity * (10 ** filDecimal)) < 2) {
              this.$message('交易额过低，无法交易！')
              return
            }
            const myContract = exchange.address
            const orderid = idArr[0]
            // 更新订单id
            if (idArr.length === 1) {
              this.$socket.sendObj({
                namespace: 'user',
                action: 'getid',
                data: ''
              })
            }
            this.$store.commit('user/delId', 0)
            const useraddress = this.$store.state.account.myEth.address
            const receiver = this.$store.state.account.myEos.identity.accounts[0].name
            const tradetype = true // 以太坊链内交易买单为false，跨链为true
            let sellamount = web3.utils.toWei(this.limitSellOrder.amount.toString(), 'ether')
            const price = handleEthPrice(this.limitSellOrder.price) // shj 6.18新增
            if (!price) {
              this.$message('请输入正确的订单数据===========')
              return
            }
            // const buyaddress = '0x0000000000000000000000000000000000000000'
            // const dateTime = Date.now()
            // const createTime = Math.floor(dateTime / 1000)
            const sellTokenAddress = '0x0000000000000000000000000000000000000000'
            const buyTokenAddress = '0x0000000000000000000000000000000000000000'
            const contractdata = await exchange.methods.depositsOrder(web3.utils.toHex(orderid), receiver, tradetype, web3.utils.toHex(sellamount), sellTokenAddress, price, buyTokenAddress, 4)
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
              value: web3.utils.toHex(sellamount),
              from: useraddress,
              to: myContract,
              data: payload
            }
            const tx = new Tx(rawTx)
            const privateKey = Buffer.from(this.$store.state.account.myEth.privateKey.slice(2), 'hex')
            tx.sign(privateKey)
            const serializedTx = tx.serialize()
            this.$socket.sendObj({
              namespace: 'user',
              action: 'sendeth',
              data: {
                approveTx: '0x' + serializedTx.toString('hex'),
                marketType: 4
              }
            })
            this.$message('正在下单，请稍等...')
          } else {
            this.$message('请输入正确的订单数据')
          }
        }
      }
    },
    async ethAprroveBde () { // todo bde交易区，买单，bde授权
      let param = this.routeParam
      let tokenName = 'BDE'
      let code = param.split('-')[1]
      let myEth = this.$store.state.account.myEth
      let isConnected = this.$store.state.isConnected
      let eth = this.$store.state.api.eth
      if (myEth === null || !isConnected || eth === null) {
        this.$message('未登录或者网络故障, 请登录或解锁钱包或重新加载')
      } else {
        if (code === '3') { // todo 授权操作应该分buy和sell; '3' 8.14新增
          const useraddress = myEth.address
          const contractAccount = eth[tokenName].contractAddress.address
          const myContract = exchange.address
          const coin = new web3.eth.Contract(abiContract, contractAccount)
          let quota = await coin.methods.allowance(useraddress, myContract).call()
          let approveQuota = new BigNumber(quota)
          const dec = eth[tokenName].contractAddress.decimal
          this.haveAccreditAmountBde = parseFloat(approveQuota / Math.pow(10, dec))
          this.accreditDialogBde = true
        }
      }
    },
    async approveTokenBde () { // todo bde 交易区，买单的授权操作
      let param = this.routeParam
      let tokenName = 'BDE'
      let code = param.split('-')[1]
      let myEth = this.$store.state.account.myEth
      let isConnected = this.$store.state.isConnected
      let eth = this.$store.state.api.eth
      if (myEth === null || !isConnected || eth === null) {
        this.$message('未登录或者网络故障, 请登录或解锁钱包或重新加载')
      } else {
        if (code === '3') {
          const contractAccount = eth[tokenName].contractAddress.address
          const myContract = exchange.address
          const useraddress = myEth.address
          const gas = await web3.eth.getGasPrice()
          const gasPriceHex = web3.utils.numberToHex(gas)
          const gasLimitHex = web3.utils.numberToHex(300000)
          const coin = new web3.eth.Contract(abiContract, contractAccount)
          let amount = new BigNumber(this.accreditamountBde * (10 ** eth[tokenName].contractAddress.decimal))
          try {
            coin.methods.approve(myContract, web3.utils.toHex(amount)).send({
              from: useraddress,
              gasPrice: gasPriceHex,
              gas: gasLimitHex
            })
            this.accreditDialogBde = false
            this.accreditamountBde = 0 // 清空授权额度
            this.$message('授权合约正在进行，请稍后查询授权额度!')
          } catch (e) {
            this.accreditDialogBde = false
            this.accreditamountBde = 0 // 清空授权额度
          }
        }
      }
    },
    async ethAprrove () { // todo eth的卖单授权操作
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      let code = param.split('-')[1]
      let myEth = this.$store.state.account.myEth
      let isConnected = this.$store.state.isConnected
      let eth = this.$store.state.api.eth
      if (myEth === null || !isConnected || eth === null) {
        this.$message('未登录或者网络故障, 请登录或解锁钱包或重新加载')
      } else {
        if (code === '2' || code === '3') { // todo 授权操作应该分buy和sell; '3' 8.14新增
          const useraddress = myEth.address
          const contractAccount = eth[tokenName].contractAddress.address
          const myContract = exchange.address
          const coin = new web3.eth.Contract(abiContract, contractAccount)
          let quota = await coin.methods.allowance(useraddress, myContract).call()
          let approveQuota = new BigNumber(quota)
          const dec = eth[tokenName].contractAddress.decimal
          this.haveAccreditAmount = parseFloat(approveQuota / Math.pow(10, dec))
          this.accreditDialog = true
        }
      }
    },
    async approveToken () {
      // 这里为什么没有私钥签名,难道是账户已经解锁。。。估计是使用了web3.eth.accounts管理账户，并且传入了私钥，故调用web3.eth时会自动签名 todo 问问亚松
      let param = this.routeParam
      let tokenName = param.split('-')[0]
      let code = param.split('-')[1]
      let myEth = this.$store.state.account.myEth
      let isConnected = this.$store.state.isConnected
      let eth = this.$store.state.api.eth
      if (myEth === null || !isConnected || eth === null) {
        this.$message('未登录或者网络故障, 请登录或解锁钱包或重新加载')
      } else {
        if (code === '2' || '3') { // todo bde交易区 新增
          const contractAccount = eth[tokenName].contractAddress.address
          const myContract = exchange.address
          const useraddress = myEth.address
          const gas = await web3.eth.getGasPrice()
          const gasPriceHex = web3.utils.numberToHex(gas)
          const gasLimitHex = web3.utils.numberToHex(300000)
          const coin = new web3.eth.Contract(abiContract, contractAccount)
          let amount = new BigNumber(this.accreditamount * (10 ** eth[tokenName].contractAddress.decimal))
          try {
            coin.methods.approve(myContract, web3.utils.toHex(amount)).send({
              from: useraddress,
              gasPrice: gasPriceHex,
              gas: gasLimitHex
            })
            this.accreditDialog = false
            this.accreditamount = 0 // 清空授权额度
            this.$message('授权合约正在进行，请稍后查询授权额度!')
          } catch (e) {
            this.accreditDialog = false
            this.accreditamount = 0 // 清空授权额度
          }
        }
      }
    },
    handlePrice () {
      let buyPrice = this.limitBuyOrder.price
      let sellPrice = this.limitSellOrder.price
      const buyLen = buyPrice.toString().split('.').length
      const sellLen = sellPrice.toString().split('.').length
      if (buyLen === 2) {
        const buyDecimal = buyPrice.toString().split('.')[1].length
        if (buyDecimal > 6) {
          this.limitBuyOrder.price = numFloat(buyPrice, 6)
        }
      }
      if (sellLen === 2) {
        const sellDecimal = sellPrice.toString().split('.')[1].length
        if (sellDecimal > 6) {
          this.limitSellOrder.price = numFloat(sellPrice, 6)
        }
      }
    },
    handleAmount () {
      let buyAmount = this.limitBuyOrder.amount
      let sellAmount = this.limitSellOrder.amount
      const buyLen = buyAmount.toString().split('.').length
      const sellLen = sellAmount.toString().split('.').length
      if (buyLen === 2) {
        const buyDecimal = buyAmount.toString().split('.')[1].length
        if (buyDecimal > 4) {
          this.limitBuyOrder.amount = numFloat(buyAmount, 4)
        }
      }
      if (sellLen === 2) {
        const sellDecimal = sellAmount.toString().split('.')[1].length
        if (sellDecimal > 4) {
          this.limitSellOrder.amount = numFloat(sellAmount, 4)
        }
      }
    },
    handleTotal () {
      let buyTotal = this.limitBuyOrder.total
      let sellTotal = this.limitSellOrder.total
      const buyLen = buyTotal.toString().split('.').length
      const sellLen = sellTotal.toString().split('.').length
      if (buyLen === 2) {
        const buyDecimal = buyTotal.toString().split('.')[1].length
        if (buyDecimal > 4) {
          this.limitBuyOrder.total = numFloat(buyTotal, 4)
        }
      }
      if (sellLen === 2) {
        const sellDecimal = sellTotal.toString().split('.')[1].length
        if (sellDecimal > 4) {
          this.limitSellOrder.total = numFloat(sellTotal, 4)
        }
      }
    }
  },
  beforeDestroy () {
    // 清空定时器
    clearInterval(this.buyIntervalId)
    clearInterval(this.sellIntervalId)
  }
}
</script>

<style scoped>
  @import '../../assets/css/TradeMainBottom.css';
</style>
