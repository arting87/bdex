<template>
  <div id="trade-view">
  </div>
</template>

<script>
import { widget as TvWidget } from '../../../static/tradeview/charting_library/charting_library.min.js'
import Socket from './datafeeds/socket.js'
import Datafeeds from './datafeeds/datafees.js'
export default {
  data () {
    return {
      widget: null,
      socket: new Socket(),
      datafeeds: new Datafeeds(this),
      symbol: null,
      interval: null,
      marketCode: null,
      cacheData: {},
      lastTime: 0,
      getBarTimer: null,
      isLoading: true,
      lastPrice: 1000
    }
  },
  computed: {
    // timeType() {
    //   return this.interval
    // }
  },
  created () {
    this.socket.doOpen()
    this.socket.on('message', this.onMessage)
  },
  methods: {
    // widght个人定制
    init (symbol = 'DDD', interval = 5, marketCode = 2, lang = 'zh', changeLang = false) {
      if (!this.widget || changeLang) {
        this.widget = new TvWidget({
          // widget的尺寸
          // widget: 600,
          // height: 460,
          // 初始代币
          symbol: symbol,
          theme: 'Dark',
          // 周期
          interval: interval,
          // fullscreen: true,
          autosize: true,
          // 指定要包含widget的DOM元素id
          container_id: 'trade-view',
          // JavaScript对象的实现接口 JS API 以反馈图表及数据
          datafeed: this.datafeeds,
          // static文件夹的路径
          library_path: 'static/tradeview/charting_library/',
          // 功能在默认情况下启用/禁用名称的数组
          disabled_features: [
            // 商品搜索
            'header_symbol_search',
            // 截图按钮
            'header_screenshot',
            // 比较按钮
            'header_compare',
            // 撤销恢复按钮
            'header_undo_redo',
            // 左边工具栏
            'left_toolbar',
            // 日期跳转按钮
            'go_to_date',
            // k line 左右控制按钮
            'control_bar',
            // footer工具
            'timeframes_toolbar',
            // 设置
            'header_settings'
            // K line border
            // 'border_around_the_chart',
            // 'remove_library_container_border'
          ],
          enabled_features: [
            'hide_last_na_study_output'
            // 'charting_library_debug_mode'
          ],
          // 图表的初始时区
          timezone: 'Asia/Shanghai',
          // hide_side_toolbar: true,
          // 图表库的本地化处理：语言
          locale: lang, // 支持语言切换
          // 图表将详细的API日志写入控制台
          debug: false
        })
      } else {
        const self = this
        this.widget.onChartReady(function () {
          self.widget.setSymbol(symbol, interval, function (data) {
          })
        })
      }
      this.symbol = symbol
      this.interval = interval
      this.marketCode = marketCode
      this.sendMessage(this.getParams())
    },
    getParams () {
      let timeType = 1
      if (this.interval === '1D') {
        timeType = 60 * 24
      } else if (this.interval === '1W') {
        timeType = 60 * 24 * 7
      } else if (this.interval === '1M') {
        timeType = 60 * 24 * 30
      } else {
        timeType = this.interval
      }
      const params = {
        namespace: 'api',
        action: 'k',
        data: {
          timeType: parseInt(timeType),
          symbol: this.symbol,
          marketCode: this.marketCode
        }
      }
      return params
    },
    sendMessage (data) {
      if (this.socket.checkOpen()) {
        this.socket.send(data)
      } else {
        this.socket.on('open', () => {
          this.socket.send(data)
        })
      }
    },
    unSubscribe (interval) {
      // this.params.data.timeType = interval
      // this.sendMessage(this.params)
    },
    subscribe () {
      // console.log(this.getParams())
      // this.sendMessage(this.getParams())
    },
    onMessage (data) {
      const ticker = `${this.symbol}-${this.interval}`
      if (data.data && data.data.length) {
        if (data.action && data.action.indexOf('INIT') !== -1) {
          const list = []
          data.data.forEach(function (element) {
            list.push({
              time: this.interval !== 'D' || this.interval !== '1D' ? element.gapStart : element.gapStart / 1000,
              open: element.openPrice,
              high: element.highPrice,
              low: element.lowerPrice,
              close: element.closePrice,
              volume: element.quoteVol
            })
          }, this)
          this.cacheData[ticker] = list
          this.lastTime = list[list.length - 1].time
          this.lastPrice = list[list.length - 1].close
          // this.subscribe()
        } else if (data.action && data.action.indexOf('UPDATE') !== -1) {
          const newData = data.data[0]
          const newBar = {
            time: newData.gapStart,
            open: newData.openPrice,
            high: newData.highPrice,
            low: newData.lowerPrice,
            close: newData.closePrice,
            volume: newData.quoteVol
          }
          if (newBar.time >= this.lastTime && this.cacheData[ticker] && this.cacheData[ticker].length) {
            this.cacheData[ticker][this.cacheData[ticker].length - 1] = newBar
          }
          this.datafeeds.barsUpdater.updateData()
        }
      }
    },
    getBars (symbolInfo, resolution, rangeStartDate, rangeEndDate, onLoadedCallback) {
      if (this.interval.toString() !== resolution.toString()) { // shj修改, this.interval != resolution
        this.unSubscribe(this.interval)
        this.interval = resolution
        this.sendMessage(this.getParams())
      }
      const ticker = `${this.symbol}-${this.interval}`
      if (this.cacheData[ticker] && this.cacheData[ticker].length) { // 分辨率改变，则ticker改变，原本存储的数据还在，但是没有cacheData[ticker]数据
        this.isLoading = false
        const newBars = []
        this.cacheData[ticker].forEach(item => {
          if (item.time >= rangeStartDate * 1000 && item.time <= rangeEndDate * 1000) {
            newBars.push(item)
          }
        })
        onLoadedCallback(newBars)
      } else {
        const self = this
        this.getBarTimer = setTimeout(function () {
          self.getBars(symbolInfo, resolution, rangeStartDate, rangeEndDate, onLoadedCallback)
        }, 2000)
      }
    },
    getSymbol () {
      return {
        'name': this.symbol,
        'pricescale': this.adjustScale(),
        'fractional': true,
        'description': this.symbol,
        'ticker': this.symbol + '-' + this.interval
      }
    },
    adjustScale () {
      if (this.lastPrice > 100) {
        return 100
      } else {
        return 10000
      }
    }
  }
}
</script>

<style scoped>
#trade-view {
  height: 39.2rem;
  width: 37.5rem
}
.layout__area--top {
  background-color: #131624;
}
</style>
