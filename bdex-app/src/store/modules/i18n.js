import i18n from '../../utils/i18n'

const state = {
  lang: null
}

const mutations = {
  switchLang (state, lang) {
    state.lang = lang
    i18n.locale = lang
    localStorage.lang = lang
  }
}

export default {
  namespaced: true,
  state,
  mutations
}
