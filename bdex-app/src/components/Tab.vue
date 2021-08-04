<template>
  <div class="tab">
    <div v-for="(item, index) in tabList" :key="index" :class="{'active': tabIndex === index}" @click="switchTab(index)">{{item}}</div>
  </div>
</template>

<script>
/* eslint-disable */

export default {
  name: 'tab',
  props: ['tabList', 'tabIndex'],
  data() {
    return {
      
    }
  },
  methods: {
    switchTab(index) {
      this.$emit('update:tabIndex', index)
    }
  }
}
</script>

<style lang="scss" scoped>
.tab {
  width: 100%;
  height: 4rem;
  background-color: $bg;
  box-shadow: 0rem 0.2rem 0.4rem 0rem rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: space-around;
  align-items: center;
  div {
    flex: 1;
    font-size: 1.2rem;
    color: #8CA1FF;
    height: 100%;
    display: flex;
    justify-content: space-around;
    align-items: center;
  }
  .active {
    background-color: #21263C;
    color: #fff;
  }
}
</style>
