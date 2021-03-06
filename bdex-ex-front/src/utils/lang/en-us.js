import enLocale from 'element-ui/lib/locale/lang/en'

export default {
  header: {
    home: 'Home',
    exchange: 'Exchange',
    news: 'News',
    lang: 'English',
    login: 'EOS Wallet',
    wallet: 'ETH Wallet',
    EosAccount: 'EOS：'
  },
  footer: {
    submitToken: 'SubmitToken',
    download: 'Download',
    androidApk: 'Download the Android app',
    iosApk: 'Download the Apple app',
    businessEmail: 'Business Cooperation Email',
    submitTokenRemindFront: 'SubmitToken Please',
    submitTokenRemindBack: 'application form, please complete all items sent to：',
    contactUs: 'Contact Us'
  },
  home: {
    favorites: 'Favorites',
    eosMain: 'EOS Main',
    eosSide: 'BDE Trade',
    ethSide: 'ETH Side',
    crossChain: 'Cross Chain',
    pairs: 'Pairs',
    lastPrice: 'Last Price',
    '24hChange': '24H Change',
    '24hHigh': '24H High',
    '24hLow': '24H Low',
    '24hVolume': '24H Volume',
    '24hTotal': '24H Total',
    search: 'Search'
  },
  trade: {
    limitorder: 'Limit Order',
    latestDeal: 'Latest Transactions',
    time: 'Time',
    price: 'Price',
    amount: 'Amt',
    total: 'Total',
    // tradeasidetop
    placeholder: 'serach',
    favorite: 'Fav',
    pair: 'Pair',
    lastprice: 'Last Price',
    change: 'Change',
    cross: 'Cross Chain',
    sideChain: 'BDE',
    // tradefootertop
    openOrder: 'Open Order',
    orderHistory: 'Order History',
    direction: 'Direction',
    orderPrice: 'Order Price',
    orderVol: 'OrderVol',
    orderTotal: 'Order Total',
    dealtVol: 'DealtVol',
    pendingVol: 'Pending Vol',
    action: 'Action',
    status: 'Status',
    // tradeMainBottom
    accreditbtn: 'Approve',
    accredit: 'Approve contract',
    beauthorized: 'Be authorized：',
    accreditPlaceholder: 'Please enter the amount',
    buy: 'Buy',
    sell: 'Sell',
    balance: 'Balance',
    buyprice: 'Buy Price',
    buyamount: 'Buy Amt',
    buytotal: 'Total',
    sellprice: 'Sell Price',
    sellamount: 'Sell Amt',
    withdraw: 'Withdraw',
    trade: 'buy/sell'
  },
  eth: {
    accoutStateInfo: 'Please Login ETH',
    accoutStateInfo2: 'Please unlock your wallet',
    newaccount: 'New ETH Account',
    importaccount: 'Import ETH Account',
    blockexplorer: 'Blockchain Explorer',
    unlockwallet: 'Unlock Wallet',
    punlockwallet: 'Please unlock your ETH wallet',
    importkey: 'Export Private Key',
    deleaccount: 'Delete Account',
    confirmDelete: 'Confirm deletion',
    password: 'Password',
    enterpassword: 'Please enter your password',
    enterprivatekey: 'Please enter your private key',
    cancel: 'Cancel',
    createaccount: 'Create',
    ethaccount: 'ETH Account',
    yes: 'Yes',
    part1: 'You just created an Ethereum account:',
    part2: 'Please BACKUP the private key for this account, we will not be saved after the webpage is closed:',
    privateKey: 'Private Key',
    part3: 'You just import an Ethereum account:',
    part4: 'For account:',
    part5: 'The private key is:',
    part6: 'You are about to remove an Ethereum account:',
    part7: 'If it has funds, please BACKUP the private key for this account.'
  },
  news: {
    announcement: 'Announcement of the center',
    importantAnnouncement: 'Important Announcements',
    activity: 'Activity',
    currencyAnnouncement: 'New Token List',
    others: 'Others'
  },
  // {{ $t(' ') }}  "$t(' ')" $t('')
  // tradeasidebottom
  ...enLocale
}
