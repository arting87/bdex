// eos 配置
// const isProduction = process.env.NODE_ENV === 'production'
// 交易所上线主网配置
const myEosContract = '' // TODO 替换部署合约的eos账户名
const myEosNet = { // scatter中提供的eos主网配置
  blockchain: 'eos',
  protocol: 'https',
  host: 'nodes.get-scatter.com',
  port: 443,
  chainId: 'aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906'
}

// 交易所测试网 配置
// const myEosContract = 'tbcexchange1' // 交易所合约
// const myEosNet = {
//   blockchain: 'eos',
//   protocol: 'http',
//   host: '47.102.110.128',
//   port: 8888,
//   chainId: '6cbecff836a9fa60da53bf97a0a180103b2e76041d4414693d11bf39e2341547'
// }

export {myEosContract, myEosNet}
