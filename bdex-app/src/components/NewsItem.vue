<template>
  <div class="news-item" @click="$router.push('/index/newsdetails')">
    <div class="news-item-container">
      <div class="star"></div>
      <div class="title"><span>{{title}}</span></div>
      <div class="right-arrow"></div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'news-item',
  props: ['title'],
  data () {
    return {
    }
  },
  methods: {
  }
}
</script>

<style lang="scss" scoped>
.news-item {
  width: 100%;
  display: flex;
  justify-content: center;
  border-bottom: solid .1rem #f0f0f0;
  .news-item-container {
    width: 33.5rem;
    height: 5rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    .title {
      width: 26.4rem;
      overflow: hidden;
      justify-self: flex-start;
      color: #333333;
      font-size: 1.4rem;
    }
    .star {
      width: 1.9rem;
      height: 1.7rem;
      background-image: url('../../static/images/公告中心-星星@2x.png');
      background-repeat: no-repeat;
      background-size: contain;
    }
    .right-arrow {
      width: .7rem;
      height: 1.2rem;
      background-image: url('../../static/images/公告中心-右箭头@2x.png');
      background-repeat: no-repeat;
      background-size: contain;
    }
  }
}
</style>
