import zhLocale from 'element-ui/lib/locale/lang/zh-CN'

export default {
  header: {
    home: '首页',
    exchange: '交易',
    news: '新闻公告',
    lang: '简体中文',
    login: 'EOS钱包',
    wallet: 'ETH钱包',
    EosAccount: 'EOS账号：'
  },
  footer: {
    submitToken: '上币申请',
    download: '下载',
    androidApk: '下载安卓端',
    iosApk: '下载苹果端',
    businessEmail: '商务合作邮箱',
    submitTokenRemindFront: '提交token请',
    submitTokenRemindBack: '申请表单，请完整填写表格发送至：',
    contactUs: '联系我们'
  },
  home: {
    favorites: '自选区',
    eosMain: 'EOS 链内',
    eosSide: 'BDE 交易',
    ethSide: 'ETH 链内',
    crossChain: '跨链',
    pairs: '交易对',
    lastPrice: '最新价',
    '24hChange': '24H 涨跌幅',
    '24hHigh': '24H 最高价',
    '24hLow': '24H 最低价',
    '24hVolume': '24H 成交量',
    '24hTotal': '24H 成交额',
    search: '输入关键字搜索'
  },
  trade: {
    limitorder: '限价交易',
    latestDeal: '最新价格',
    time: '委托时间',
    price: '价格',
    amount: '数量',
    total: '金额',
    // tradeasidetop
    placeholder: '币种搜索',
    favorite: '自选',
    pair: '交易对',
    lastprice: '最新价格',
    change: '涨跌',
    cross: '跨链',
    sideChain: 'BDE',
    // tradefootertop
    openOrder: '当前委托',
    orderHistory: '历史委托',
    direction: '方向',
    orderPrice: '委托价',
    orderVol: '委托量',
    orderTotal: '委托金额',
    dealtVol: '成交量',
    pendingVol: '未成交',
    action: '操作',
    status: '状态',
    // tradeMainBottom
    accreditbtn: '授权',
    accredit: '授权合约',
    beauthorized: '已授权：',
    accreditPlaceholder: '请输入要授权的额度',
    buy: '买入',
    sell: '卖出',
    balance: '余额',
    buyprice: '买入价',
    buyamount: '买入量',
    buytotal: '交易额',
    sellprice: '卖出价',
    sellamount: '卖出量',
    withdraw: '撤单',
    trade: '买卖'
  },
  eth: {
    accoutStateInfo: '请登录ETH',
    accoutStateInfo2: '请解锁您的钱包',
    newaccount: '新建账户',
    importaccount: '导入账户',
    blockexplorer: '区块链浏览器',
    unlockwallet: '解锁钱包',
    punlockwallet: '请解锁ETH钱包',
    importkey: '导出私钥',
    deleaccount: '删除账户',
    confirmDelete: '确认删除',
    password: '密码',
    enterpassword: '请输入钱包密码',
    enterprivatekey: '请输入私钥',
    cancel: '取 消',
    createaccount: '创建账号',
    ethaccount: 'ETH账号',
    yes: '确 定',
    part1: '你创建了一个以太坊账户:',
    part2: '请备份你的私钥，我们网页端关闭后不会保存私钥:',
    privateKey: '私钥',
    part3: '你导入了一个以太坊账户:',
    part4: '账户地址:',
    part5: '私钥:',
    part6: '删除以太坊账户:',
    part7: '删除账户前请备份你的私钥。'
  },
  news: {
    announcement: '公告中心',
    importantAnnouncement: '重要公告',
    activity: '活动公告',
    currencyAnnouncement: '新币上线',
    others: '其他公告'

  },
  // {{ $t(' ') }}  "$t(' ')" $t('')
  // tradeasidebottom
  ...zhLocale
}
