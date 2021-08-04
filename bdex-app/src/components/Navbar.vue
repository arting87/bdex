<template>
  <div class="navbar" :class="type">
    <div @click="goBack"></div>
    <div>{{text}}</div>
    <div></div>
  </div>
</template>

<script>
export default {
  name: 'navbar',
  props: ['text', 'type'],
  data () {
    return {
    }
  },
  methods: {
    goBack () {
      this.$router.go(-1)
    }
  }
}
</script>

<style lang="scss" scoped>
.navbar {
  width: 100%;
  height: 4.4rem;
  box-shadow: 0rem 0.2rem 0.4rem 0rem rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: $bg;
  font-size: 1.4rem;
  color: #8ca1ff;
  :first-child {
    margin-left: 2rem;
    height: 1.3rem;
    width: 0.8rem;
    background-repeat: no-repeat;
    background-size: cover;
    background-image: url('../../static/images/头部-返回@2x.png');
  }
  :last-child {
    margin-right: 2rem;
  }
}
.bright {
  color: #333;
  background-color: #fff;
  :first-child {
    background-image: url('../../static/images/公告中心-头部-返回@2x.png');
  }
}
</style>
