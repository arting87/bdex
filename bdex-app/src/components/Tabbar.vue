<template>
  <div class="tabbar" v-if="navList.length > 0">
    <div class="item" v-for="(item, index) in navList"
        @click="selectNavItem(index, item.pagePath)"
        :key="index">
      <div class="item-image">
        <img :src="selectNavIndex === index ? item.selectedIconPath : item.iconPath" alt="iconPath">
      </div>
      <span :class="selectNavIndex === index ? 'item-text item-text-active' : 'item-text' ">
        {{item.text}}
      </span>
    </div>
  </div>
</template>

<script>
/* eslint-disable */
import store from '../store'

export default {
  name: 'tabbar',
  props: ['selectNavIndex'],
  data () {
    return {
      navList: [
        {
          pagePath: '/index',
          iconPath: require('../../static/images/首页-未选中@2x.png'),
          selectedIconPath: require('../../static/images/首页-已选中@2x.png'),
          text: '首页'
        },
        {
          pagePath: '/market',
          iconPath: require('../../static/images/行情-未选中@2x.png'),
          selectedIconPath: require('../../static/images/行情-已选中@2x.png'),
          text: '行情'
        },
        {
          pagePath: '/trade',
          iconPath: require('../../static/images/交易-未选中@2x.png'),
          selectedIconPath: require('../../static/images/交易-已选中@2x.png'),
          text: '交易'
        },
        {
          pagePath: '/my',
          iconPath: require('../../static/images/我的-未选中@2x.png'),
          selectedIconPath: require('../../static/images/我的-已选中@2x.png'),
          text: '我的'
        }
      ],
    }
  },
  created () {
    console.log('tab bar created')
  },
  mounted () {
    if (this.selectNavIndex !== 0) {
      this.$router.push(this.navList[this.selectNavIndex].pagePath)
    }
  },
  methods: {
    selectNavItem(index, pagePath) {
      if (index === this.selectNavIndex) {
        return false;
      }
      this.$emit('update:selectNavIndex', index)
      this.$router.push(pagePath)
    }
  },
  created() {
    this.$bus.$on('updateIndex', index => {
      if (index !== this.selectNavIndex) {
        this.$emit('update:selectNavIndex', index)
      }
    })
  }
}
</script>

<style lang="scss" scoped>
  .tabbar {
    z-index: 1000;
    position: fixed;
    bottom: 0;
    width: 100%;
    height: 5rem;
    background-color: #131624;
    box-shadow: 0rem -0.2rem 0.5rem 0rem 
    rgba(0, 0, 0, 0.55);
    display: flex;
    justify-content: space-around;
    align-items: center;
    .item {
      flex: 1;
      text-align: center;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      .item-text {
        font-size: 1rem;
        font-weight: normal;
        font-stretch: normal;
        letter-spacing: 0.06rem;
        color: #6798FF;
      }
      .item-image {
        width: 2.2rem;
        height: 2.1rem;
        transition: .24s linear;
        & img {
          width: 100%;
          height: 100%;
        }
      }
    }
  }
</style>