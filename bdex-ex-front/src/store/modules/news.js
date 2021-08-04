// 新闻公告处理，现在使用websocket，后期改为http
const state = {
  importantnews: null, // 重要公告, array
  activitynews: null, // 活动公告
  currencyNews: null, // 新币上线公告
  otherNews: null, // 其它公告
  newsContent: null // 新闻详情， {}
}

const mutations = {
  setNewList (state, {importantnews, activitynews, currencyNews, otherNews}) {
    state.importantnews = importantnews
    state.activitynews = activitynews
    state.currencyNews = currencyNews
    state.otherNews = otherNews
  },
  setNewsContent (state, newsContent) {
    if (state.newsContent === null) {
      state.newsContent = {}
    }
    let id = newsContent.id
    state.newsContent[`${id}`] = newsContent // id作为key, newsContent作为value
  }
}

const actions = {
  newsList ({commit}, msg) { // 新闻列表
    let data = msg.data
    let error = data.error
    if (error) {
      console.warn('获取新闻公告发生错误，请检查网络！')
    } else {
      const importantnews = []
      const activitynews = []
      const currencyNews = []
      const otherNews = []
      data.forEach((el) => {
        if (el.typeCode === 1) { // 重要公告
          importantnews.push(el)
        } else if (el.typeCode === 2) {
          activitynews.push(el)
        } else if (el.typeCode === 3) {
          currencyNews.push(el)
        } else {
          otherNews.push(el)
        }
      })
      commit('setNewList', {importantnews, activitynews, currencyNews, otherNews})
    }
  },
  onenews ({commit}, msg) { // 新闻详情
    let data = msg.data
    let error = data.error
    if (error) {
      console.warn('获取新闻公告发生错误，请检查网络!')
    } else {
      commit('setNewsContent', data)
    }
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
