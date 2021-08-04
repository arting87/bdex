<template>
  <div class="index">
    <Header/>
    <section>
    <router-view class="main"/>
    </section>
    <Footer class="footer" />
  </div>
</template>

<script>
import Footer from './common/Footer'
import Header from './common/Header'
export default {
  name: 'Index',
  data () {
    return {
    }
  },
  components: {
    Footer,
    Header
  }
}
</script>

<style scoped>
.main {
  min-height: 2100px;
  margin-top: 90px;
}
</style>
