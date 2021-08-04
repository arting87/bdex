<template>
  <div class="order-item">
    <div class="order-item-area unfinished-order" v-if="(data.status === 2 || data.status === 5)">
      <div class="top order-item-row">
        <div class="action-desc">
          <span :class="{buy: data.direction === 0}">{{data.direction === 0 ? $t('trade.buy') : $t('trade.sell')}}</span>
          <span>{{ data.date | time }}</span>
        </div>
        <span class="action" @click="revoke(data.orderId, data.code, data.direction)">撤销</span>
      </div>
      <div class="bottom order-item-row">
        <div>
          <span>价格({{data.code === '3' ? 'BDE' : (data.code === '2' ? 'ETH' : 'EOS')}})</span>
          <span>{{data.orderPrice}}</span>
        </div>
        <div>
          <span>数量({{data.pair}})</span>
          <span>{{data.orderVol | fixed}}</span>
        </div>
        <div>
          <span>实际成交({{data.pair}})</span>
          <span>{{data.dealVol | fixed}}</span>
        </div>
      </div>
    </div>
    <div class="order-item-area finished-order" v-else>
      <div class="top finished-top order-item-row">
        <div class="action-desc">
          <span :class="{buy: data.direction === 0}">{{data.direction === 0 ? $t('trade.buy') : $t('trade.sell')}}</span>
          <span>{{ data.pair }} {{ data.code === '3' ? '/ BDE' : (data.code === '2' ? ' / ETH' : ' / EOS') }}</span>
        </div>
        <span class="action" @click="revoke">已成交</span>
      </div>
      <div class="bottom order-item-row">
        <div>
          <span>时间</span>
          <span>{{ data.date | time }}</span>
        </div>
        <div>
          <span>委托价({{data.code === '3' ? 'BDE' : (data.code === '2' ? 'ETH' : 'EOS')}})</span>
          <span>{{data.orderPrice}}</span>
        </div>
        <div>
          <span>委托量({{data.pair}})</span>
          <span>{{data.orderVol | fixed}}</span>
        </div>
      </div>
      <div class="bottom order-item-row">
        <div>
          <span>成交量总额({{data.code === '3' ? 'BDE' : (data.code === '2' ? 'ETH' : 'EOS')}})</span>
          <span>{{(data.orderPrice * data.orderVol) | fixed}}</span>
        </div>
        <div>
          <span>成交均价({{data.code === '3' ? 'BDE' : (data.code === '2' ? 'ETH' : 'EOS')}})</span>
          <span>{{data.orderPrice}}</span>
        </div>
        <div>
          <span>委托量({{data.pair}})</span>
          <span>{{data.orderVol | fixed}}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
/* eslint-disable */
import eosTrade from '../lib/eos/eosTrade'
import ethTrade from '../lib/eth/ethTrade'
export default {
  name: 'order-item',
  props: ['data'],
  data() {
    return {
      
    }
  },
  methods: {
    revoke (orderId, code, direction) {
      console.log('执行撤销操作-------------', orderId, code, direction)
      if (code === '1') { // eos撤单
        eosTrade.revokeOrder(orderId)
      } else if (code === '2') { // eth撤单
        ethTrade.revokeOrder(orderId)
      } else if (code === '3') { //
        ethTrade.bdeRevokeOrder(orderId)
      } else if (code === '4') {
        if (direction === 0) {
          eosTrade.revokeOrder(orderId)
        } else if (direction === 1) {
          ethTrade.revokeOrder(orderId)
        }
      }
    },

  }
}
</script>

<style lang="scss" scoped>
.order-item {
  font-size: 1.2rem;
  .order-item-area {
    width: 100%;
    border-bottom: solid .1rem $divideColor;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
    :last-child {
      margin-bottom: 2.2rem;
    }
    .order-item-row {
      width: 33.5rem;
      height: 3.8rem;
      display: flex;
      justify-content: space-between;
    }
    .top {
      margin-top: 2.2rem;
      .action-desc {
        :first-child {
          font-size: 1.4rem;
          color: $sellColor;
        }
        .buy {
          color: $buyColor;
        }
        :last-child {
          margin-left: .5rem;
          color: $textColor;
        }
      }
      .action {
        text-decoration: underline;
        color: $revokeTextColor;
      }
    }
    .finished-top {
      .action-desc {
        :last-child {
          font-size: 1.4rem;
        }
      }
      .action {
        text-decoration: none;
        color: $whiteTextColor;
      }
    }
    .bottom {
      display: flex;
      justify-content: space-between;
      :nth-child(2) {
        align-items: center;
      }
      :last-child {
        align-items: flex-end;
      }
      div {
        display: flex;
        flex-direction: column;
        flex: 1;
        :first-child {
          color: $whiteTextColor;
          opacity: .5;
        }
        :last-child {
          color: $textColor;
        }
      }
    }
  }
  .unfinished-order {
    height: 10.4rem;
  }
  .finished-order {
    height: 16.5rem;
  }
}
</style>
