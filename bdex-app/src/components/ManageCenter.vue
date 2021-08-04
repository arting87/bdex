<template>
  <div class="manage-center">
    <div @click="goto('/my/manageResource')">资源管理</div>
    <div @click="goto('/my/manageRam')">内存管理</div>
    <div>待开发</div>
  </div>
</template>

<script>
/* eslint-disable */
import store from '../store'

export default {
  name: 'manage-center',
  props: [],
  data() {
    return {
      
    }
  },
  created() {

  },
  mounted() {
    
  },
  methods: {
    goto (url) {
      this.$router.push(url)
    }
  }
}
</script>

<style lang="scss" scoped>
.manage-center {
  height: 9rem;
  display: flex;
  justify-content: space-around;
  align-items: center;
  font-size: 1.2rem;
  color: #8ca1ff;
  div {
    height: 4.9rem;
    background-repeat: no-repeat;
    background-size: 2.8rem 2.8rem;
    background-position-x: center;
    display: flex;
    align-items: flex-end;
    justify-content: space-around;
  }
  :nth-child(1) {
    background-image: url('../../static/images/资源管理@2x.png');
  }
  :nth-child(2) {
    background-image: url('../../static/images/内存管理@2x.png');
  }
  :nth-child(3) {
    background-image: url('../../static/images/待开发@2x.png');
  }
}
</style>