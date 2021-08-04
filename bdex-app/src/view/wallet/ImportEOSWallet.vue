<template>
  <div class="import-eos-wallet">
    <navbar text="导入钱包" />
    <div class="have-wallet-area wallet-area">
      <div class="text-area">
        <span>我有eos账户</span>
      </div>
      <div class="import-area" @click="goto('/my/importPriEOS')">
        <div class="left">
          <span>导入钱包</span>
          <span>通过私钥导入钱包</span>
        </div>
        <div class="right">
        </div>
      </div>
    </div>
    <div class="no-wallet-area wallet-area">
      <div class="text-area">
        <span>我没有eos账户</span>
        <span>1.可下载imtoken钱包</span>
        <span>2.根据imtoken相关流程创建钱包</span>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from '../../components/Navbar'

export default {
  name: 'import-eos-wallet',
  data () {
    return {
    }
  },
  components: {
    Navbar
  },
  methods: {
    goto (url) {
      this.$router.push(url)
    }
  }
}
</script>

<style lang='scss' scoped>
.import-eos-wallet {
  display: flex;
  flex-direction: column;
  .wallet-area {
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 14.2rem;
    .text-area {
      width: 33.5rem;
      height: 6rem;
      font-size: 1.4rem;
      color: $textColor;
      height: 6rem;
      display: flex;
      align-items: center;
      font-weight: 500;
    }
  }
  .import-area, .create-area {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-repeat: no-repeat;
    background-size: cover;
    height: 8.2rem;
    width: 100%;
    .left {
      margin-left: 2rem;
      display: flex;
      flex-direction: column;
      justify-content: center;
      color: $whiteTextColor;
      :first-child {
        font-size: 1.4rem;
        font-weight: 500;
      }
      :last-child {
        margin-top: 1rem;
        font-size: 1.2rem;
      }
    }
    .right {
      margin-right: 2rem;
      background-image: url('../../../static/images/历史记录-右箭头@2x.png');
      background-repeat: no-repeat;
      background-size: cover;
      width: .7rem;
      height: 1.2rem;
    }
  }
  .import-area {
    background-image: url('../../../static/images/钱包-EOS提醒-导入钱包背景@2x.png');
  }
  .no-wallet-area {
    .text-area {
      margin-top: 3rem;
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      :nth-child(2) {
        margin-top: 1rem;
      }
      :nth-child(2), :nth-child(3) {
        color: rgba(133, 171, 255, 0.5);
        font-size: 1.2rem;
      }
    }
  }
}
</style>
