<template>
  <div class="menu">
    <el-button type="text" size="mini" @click="console.log('tab, event');">
      <i class="el-icon-star-off"></i>
    </el-button>
    <input type="button" value="下载" @click="downloadFile()" />
  </div>
</template>

<script>
export default {
  name: 'Menu',
  data () {
    return {
    }
  },
  methods: {
    handleClick () {
      console.log('tab, event')
    },
    downloadFile () {
      var url = 'src/SubmitToken.pdf'
      window.open(url)
    }
  }
}
</script>

<style scoped>
.menu {
  height: 100px;
}
</style>
