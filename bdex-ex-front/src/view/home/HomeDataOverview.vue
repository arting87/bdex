<template>
  <div class="homedataoverview">
    <el-col :span="7" :offset="1">
      <span class="label"></span>
      <div class="value"></div>
    </el-col>
    <el-col :span="7" :offset="0">
      <span class="label"></span>
      <div class="value"></div>
    </el-col>
    <el-col :span="7" :offset="0">
      <span class="label"></span>
      <div class="value"></div>
    </el-col>
  </div>
</template>

<script>
export default {
  name: 'HomeDataOverview'
}
</script>

<style scoped>
.label {
  font-size: 13px;
  font-weight: 900;
  text-transform: uppercase;
}
.value {
  font-size: 21px;
  line-height: 23px;
  text-transform: uppercase;
}
</style>
