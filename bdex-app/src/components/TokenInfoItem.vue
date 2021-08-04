<template>
  <div class="token-info-item" :class="data.change > 0 ? 'rise' : 'fall'" @click="clickItem">
    <div class="container fixed-width" v-if="!simple">
      <div class="left">
        <span class="symbol">{{ data.code === '3' ? (data.pair + ' / BDE') : (data.code === '2' ? data.pair + ' / ETH' : data.pair + ' / EOS') }}</span>
        <span>24h量 {{data.volume}}</span>
      </div>
      <div class="middle">
        <span class="price">{{data.lastPrice | fixed}}</span>
        <!--首页 <span>≈¥0.068</span> -->
      </div>
      <div class="right">
        <div class="change-area">
          {{handleChange()}}
        </div>
      </div>
    </div>
    <div class="simple-container" v-else>
      <span class="symbol">{{ data.code === '3' ? (data.pair + ' / BDE') : (data.code === '2' ? data.pair + ' / ETH' : data.pair + ' / EOS') }}</span>
      <span class="price">{{data.lastPrice | fixed}}</span>
    </div>
  </div>
</template>

<script>
/* eslint-disable */

export default {
  name: 'token-info-item',
  props: ['data', 'simple'],
  data() {
    return {
      
    }
  },
  methods: {
    clickItem () { // 行情列表，row-click
      this.$emit('clickItem')
      this.$bus.$emit('updateIndex', 2)
      this.$router.push({
        path: '/trade/token/' + this.data.pair + '-' + this.data.code
      })
    },
    handleChange () { // 更新数据是否会重新渲染？待测试
      if (this.data === null) {
        return '--'
      } else {
        let change = this.data.change
        if (change >= 0) {
          return '+' + change + '%'
        } else {
          return change + '%'
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.token-info-item {
  width: 100%;
  height: 7.6rem;
  background-color: $bg;
  border-bottom: solid .1rem #21263C;
  font-size: 1.2rem;
  display: flex;
  justify-content: center;
  .simple-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 25.2rem;
    height: 2.9rem;
    :first-child {
      font-size: 1.4rem;
    }
    :last-child {
      font-size: 1.2rem;
    }
  }
  .container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    div {
      display: flex;
      flex-direction: column;
      justify-content: space-between;
    }
  }
  .left {
    :first-child {
      color: $textColor;
      font-size: 1rem;
    }
  }
  .left, .middle {
    :first-child {
      font-size: 1.4rem;
    }
    :last-child {
      color: $whiteTextColor;
    }
    height: 4rem;
  }
  .right {
    .change-area {
      width: 6.6rem;
      height: 3.4rem;
      border-radius: 0.2rem;
      display: flex;
      justify-content: center;
      align-items: center;
      color: $whiteTextColor;
    }
  }
}
.rise {
  .price {
    color: $riseColor;
  }
  .change-area {
    background-color: $riseColor;
  }
}
.fall {
  .price {
    color: $fallColor;
  }
  .change-area {
    background-color: $fallColor;
  }
}
.symbol {
  color: #8ac1ff;
}
</style>
