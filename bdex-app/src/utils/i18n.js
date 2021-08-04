import Vue from 'vue'
import ElementLocale from 'element-ui/lib/locale'
import VueI18n from 'vue-i18n'
import en from './lang/en-us'
import zh from './lang/zh-cn'

Vue.use(VueI18n)

let lang = localStorage.lang
if (lang !== 'zh' && lang !== 'en') {
  lang = ''
}

const i18n = new VueI18n({
  locale: lang || 'zh',
  fallbackLocale: 'zh',
  messages: {
    en: en,
    zh: zh
  }
})

ElementLocale.i18n((key, value) => i18n.t(key, value))

export default i18n

/* html调用说明:
* 调用自定义语言包：{{ $t('message.xxxx')}}
* 调用element的语言包: {{ $t('el.xxxx.xxxx') }}
* */
