<template>
  <div class="footer">
    <div class="footerlayout">
      <span>
        <el-button
          type="text"
          @click="submitTokenUsDialog = true"
          class="submittoken"
        >{{ $t('footer.submitToken')}}</el-button>
      </span>
      <span class="middleborder"></span>
      <span>
        <el-button
          type="text"
          @click="contactUsDialog = true"
          class="contactus"
        >{{ $t('footer.contactUs') }}</el-button>
      </span>
      <span class="middleborder"></span>
        <a href="static\bdex.apk"
           download="bdex.apk"
           class="download-apk"
        >{{ $t('footer.androidApk') }}</a>
      <span class="middleborder"></span>
        <a href="http://app.dibaqu.top/6zmu"
           class="download-apk"
           target="_Blank"
        >{{ $t('footer.iosApk') }}</a>
      <div class="footer-bottom">Copyright © 2018-2019 bdex | All rights reserved</div>
    </div>
    <!-- 提交token dialog -->
    <el-dialog
      :visible.sync="submitTokenUsDialog"
      width="480px"
      top="32vh"
      :modal="modal"
      class="submittokendialog"
    >
      <div class="submittokenbody">
        <el-image :src="submittokenurl" fit="fill" class="submittokenimg"></el-image>
        <div>
          <span>{{ $t('footer.submitTokenRemindFront') }}</span>
          <a href="static\SubmitToken.pdf"
             download="SubmitToken.pdf">{{ $t('footer.download') }}</a>
          <span>{{ $t('footer.submitTokenRemindBack') }}</span>
          <p>business@bdex.me</p>
        </div>
        <span slot="footer" class="dialog-footer">
          <el-button type="primary" @click="submitTokenUsDialog = false">{{ $t('eth.yes') }}</el-button>
        </span>
      </div>
    </el-dialog>

    <!-- 联系我们dialog -->
    <el-dialog
      :visible.sync="contactUsDialog"
      class="newaccountdialog"
      width="480px"
      top="32vh"
      :modal="modal"
    >
      <div class="newaccountbody">
        <el-image :src="newaccounturl" fit="fill" class="newaccountimg"></el-image>
        <div class="businesscooperation">
          {{ $t('footer.businessEmail') }} :
          <p>business@bdex.me</p>
        </div>
        <span slot="footer" class="dialog-footer">
          <el-button type="primary" @click="contactUsDialog = false">{{ $t('eth.yes') }}</el-button>
        </span>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Footer',
  data () {
    return {
      newaccounturl: require('../../assets/img/newaccount.png'),
      submittokenurl: require('../../assets/img/download.png'),
      modal: false,
      submitTokenUsDialog: false,
      contactUsDialog: false
    }
  },
  methods: {
  }
}
</script>

<style scoped>
  @import '../../assets/css/Footer.css';
</style>
