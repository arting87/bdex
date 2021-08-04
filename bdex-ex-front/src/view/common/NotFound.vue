<template>
  <div class="notfound">
    <el-image :src="imgurl" fit="cover"></el-image>
    <el-button round class="backhome" @click="$router.push({path: '/'})">回到首页</el-button>
  </div>
</template>

<script>
export default {
  name: 'NotFound',
  data () {
    return {
      imgurl: require('../../assets/img/404.jpg')
    }
  }
}
</script>

<style scoped>
/* .notfound {
  margin: 0;
  padding: 0;
} */
.backhome {
  position: absolute;
  left: 43%;
  top: 72%;
  border: 1px solid #e18b26;
  background: rgba(255, 255, 255, 0);
  color: #ffffff;
  letter-spacing: 5px;
  font-size: 16px;
  width: 15%;
}
</style>
