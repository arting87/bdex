<template>
  <div id="app">
    <router-view/>
    <tabbar :selectNavIndex.sync='index' v-show="$route.meta.needTabbar" />
  </div>
</template>

<script>
import Tabbar from '@/components/Tabbar'

export default {
  name: 'App',
  data () {
    return {
      index: 0
    }
  },
  components: {
    Tabbar
  },
  mounted () {
  },
  created () {
  }
}
</script>

<style lang="scss">
html {
  font-size: 10px;
  background-color: $bg;
}

body {
  color: #2c3e50;
  padding: 0;
  margin: 0;
  display: block;
  // font: "Microsoft YaHei", tahoma, arial, sans-serif;
  min-width: 320px;
  position: relative;
}
#app {
  font-family: -apple-system, BlinkMacSystemFont, "PingFang SC","Helvetica Neue",STHeiti,"Microsoft Yahei",Tahoma,Simsun,sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  width: 100%;
}
img {
  border: none;
}
form,
input {
  background: none;
  border: none;
}

</style>
