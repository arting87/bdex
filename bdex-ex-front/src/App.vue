<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
body {
  margin: 0;
  padding: 0;
  color: #2c3e50;
  position: relative;
  background-color: #1A1D2C;
  /* 页面最小宽度 */
  min-width: 1260px;
  overflow-x: auto;
}
::-webkit-scrollbar {
  width: 4px;
  height: 2px;
}

::-webkit-scrollbar-thumb {
  background-color: #1a1d2c;
  border-radius: 3px;
}

</style>
