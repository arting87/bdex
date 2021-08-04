<template>
  <div class="my">
    <div class="top-area">
      <div class="account-info">
        <div class="account-desc">
          <div class="img"></div>
          <div class="name">
            <p>账户：{{currentAddress | omitAddress}}</p>
          </div>
          <div class="copy" ref="copyButton" :data-clipboard-text="currentAddress">
          <van-icon name="description"/>
          </div>
        </div>
      </div>
      <div class="assets">
        <span>总资产</span>
        <span>{{currentAssets}} {{currentToken}}</span>
        <span>≈￥{{getCNY()}}</span>
      </div>
      <div class="action" v-if="tabIndex === 0 && myEos || tabIndex === 1 && myEth">
        <div class="my-button" @click="deposit">
          充币
        </div>
        <div class="my-button" @click="withdraw">
          提币
        </div>
      </div>
      <div class="action" v-else>
        <div class="my-button" @click="toImportWallet">
          登录
        </div>
      </div>
    </div>
    <tab :tabList="tabList" :tabIndex.sync="tabIndex" />
    <div class="assets-desc" @click="toImportWallet">
      <div class="desc">
        <div>
          <span>{{currentToken}}</span>
          <span>总资产折合({{currentToken}})</span>
        </div>
        <div>
          <span>{{currentAssets}}</span>
          <span>≈{{getCNY(currentAssets)}}￥</span>
        </div>
      </div>
      <div class="right-arrow"></div>
    </div>
    <div class="divide-1"></div>
    <div class="manage-area">
      <manage-center />
    </div>
    <div class="hide-search">
      <div class="left">
        <div class="hide-token" :class="{'active': hideToken}" @click="toggleHideToken"></div>
        <span>隐藏小额币种</span>
        <div class="tip" @click="tipHideToken"></div>
      </div>
    </div>
    <div class="token-list" v-if="tokenList.length > 0">
      <div class="token-item">
        <div class="top">
          <span>EOS</span>
          <div class="right-arrow"></div>
        </div>
        <div class="content">
          <div>
            <span>可用</span>
            <span>22.3</span>
          </div>
          <div>
            <span>冻结</span>
            <span>2.3</span>
          </div>
          <div>
            <span>折合（CNY）</span>
            <span>232424</span>
          </div>
        </div>
      </div>
      <div class="divide-1"></div>
    </div>
    <div class="no-data" v-else>暂无数据</div>
  </div>
</template>

<script>
import ManageCenter from '../../components/ManageCenter'
import Tab from '../../components/Tab'
import Clipboard from 'clipboard'

export default {
  name: 'My',
  data () {
    return {
      tabIndex: 0,
      tabList: [
        'EOS账户',
        'ETH账户'
      ],
      assetsETH: 0,
      assetsEOS: 0,
      hideToken: false,
      tokenList: []
    }
  },
  computed: {
    currentAddress: function () {
      return this.tabIndex === 0 ? this.$store.state.account.myEos && this.$store.state.account.myEos.address
        : this.$store.state.account.myEth && this.$store.state.account.myEth.address
    },
    currentToken: function () { return this.tabIndex === 0 ? 'EOS' : 'ETH' },
    currentAssets: function () { return this.tabIndex === 0 ? this.assetsEOS : this.assetsETH },
    myEth: function () { return this.$store.state.account.myEth },
    myEos: function () { return this.$store.state.account.myEos }
  },
  created () {
    const self = this
    this.getETHBalance()
      .then(value => {
        self.assetsETH = value
      })
    this.getEOSBalance(this.myEos.address)
      .then(value => {
        self.assetsEOS = value[0].split(' ')[0]
      })
  },
  mounted () {
    this.initClipboard()
  },
  components: {
    ManageCenter,
    Tab
  },
  filters: {
    omitAddress (value) {
      if (!value) return ''
      if (value.length > 13) {
        return value.slice(0, 12) + '...' + value.slice(38, 42)
      }
      return value
    }
  },
  methods: {
    toggleHideToken () {
      this.hideToken = !this.hideToken
    },
    tipHideToken () {
      this.$toast('选中将会隐藏小额数字货币')
    },
    toImportWallet () {
      if (this.tabIndex === 0) {
        this.goto('/my/importEOSWallet')
      } else if (this.tabIndex === 1) {
        this.goto('/my/importETHWallet')
      }
    },
    withdraw () {
      if (this.tabIndex === 0) {
        this.goto('/my/withdrawEOS')
      } else if (this.tabIndex === 1) {
        this.goto('/my/withdrawETH')
      }
    },
    deposit () {
      if (this.tabIndex === 0) {
        this.goto('/my/depositEOS')
      } else if (this.tabIndex === 1) {
        this.goto('/my/depositETH')
      }
    },
    getCNY () {
      if (this.tabIndex === 0) {
        return this.getEOSCNY(this.currentAssets)
      } else {
        return this.getETHCNY(this.currentAssets)
      }
    },
    getEOSCNY (value) {
      return value * this.getCNYRate('EOS')
    },
    getETHCNY (value) {
      return value * this.getCNYRate('ETH')
    },
    goto (url) {
      this.$router.push(url)
    },
    initClipboard () {
      const clipboard = new Clipboard(this.$refs.copyButton)
      clipboard.on('success', (e) => {
        console.log('succsss')
        this.$dialog.alert({
          message: '复制成功'
        })
      })
    }
  },
  watch: {
  }
}
</script>

<style lang="scss" scoped>
.my {
  display: flex;
  flex-direction: column;
  color: #fff;
  align-items: center;
  .top-area {
    width: 100%;
    height: 25rem;
    background-image: url('../../../static/images/我的-背景@2x.png');
    background-repeat: no-repeat;
    background-size: cover;
    display: flex;
    flex-direction: column;
    align-items: center;
    .account-info {
      margin-top: 4rem;
      width: 33.5rem;
      display: flex;
      align-items: center;
      justify-content: space-between;
      .account-desc {
        display: flex;
        font-size: 1.2rem;
        .copy {
          width: 3.8rem;
          height: 4rem;
          background-color: #347ce100;
          border-radius: 0.2rem;
          display: flex;
          align-items: center;
          font-size: 1.4rem;
          color: #fff;
        }
      }
      // .account-setting {
      //   width: 2.1rem;
      //   height: 2.1rem;
      //   background-image: url('../../../static/images/我的-设置@2x.png');
      //   background-repeat: no-repeat;
      //   background-size: contain;
      // }
    }
    .assets {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      width: 33.5rem;
      font-size: 1.2rem;
      :nth-child(2) {
        font-size: 2.8rem;
      }
    }
    .action {
      margin-top: 2.6rem;
      width: 33.5rem;
      display: flex;
      align-items: center;
      justify-content: space-around;
      .my-button {
        width: 14.7rem;
        height: 3.4rem;
        border-radius: 0.2rem;
        border: solid 0.1rem #ffffff;
        display: flex;
        justify-content: space-around;
        align-items: center;
        font-size: 1.4rem;
      }
    }
  }
  .assets-desc {
    width: 33.5rem;
    height: 8.2rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    .desc {
      display: flex;
      flex-direction: column;
      font-size: 1.4rem;
      color: #8CA1FF;
      justify-content: space-around;
      height: 80%;
      div {
        display: flex;
        align-items: center;
        :last-child {
          margin-left: .1rem;
        }
      }
      :first-child {
        :last-child {
          font-size: 1rem;
          color: #8CA1FF;
        }
      }
      :last-child {
        :last-child {
          color: #fff;
          font-size: 1.2rem;
        }
      }
    }
  }
  .manage-area {
    width: 100%;
  }
  .hide-search {
    width: 33.5rem;
    height: 3rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    .left {
      display: flex;
      align-items: center;
      span {
        margin-left: 1rem;
      }
    }
    .hide-token {
      width: 1.4rem;
      height: 1.4rem;
      border-radius: 0.1rem;
      border: solid 0.1rem $textColor;
      opacity: 0.5;
    }
    .active {
      background-color: $textColor;
    }
    .tip {
      margin-left: 1rem;
      width: 1rem;
      height: 1rem;
      background-image: url('../../../static/images/我的-问号@2x.png');
      background-repeat: no-repeat;
      background-size: contain;
    }
  }
  .token-list {
    display: flex;
    flex-direction: column;
    align-items: center;
    .token-item {
      width: 33.5rem;
      height: 10.8rem;
      display: flex;
      flex-direction: column;
      justify-content: space-around;
      .top {
        font-size: 1.4rem;
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex: 1;
      }
      .content {
        flex: 1.2;
        font-size: 1.2rem;
        color: #fff;
        justify-content: space-between;
        display: flex;
        div {
          display: flex;
          flex-direction: column;
          :last-child {
            color: #8CA1FF;
            margin-top: 1rem;
          }
        }
      }
    }
  }
  .right-arrow {
    width: .7rem;
    height: 1.2rem;
    background-image: url('../../../static/images/历史记录-右箭头@2x.png');
    background-repeat: no-repeat;
    background-size: contain;
  }
}
</style>
