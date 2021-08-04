import web3 from './web3'
import Tx from 'ethereumjs-tx'

async function sendEther (from, to, value, pk, callback) {
  if (pk.indexOf('0x') > -1) {
    pk = pk.slice(2)
  }
  const gasPrice = await web3.eth.getGasPrice()
  const nonce = await web3.eth.getTransactionCount(from)
  const privateKey = Buffer.from(pk, 'hex')
  const rawTx = {
    nonce: web3.utils.toHex(nonce),
    gasPrice: web3.utils.toHex(gasPrice), // 2 Gwei 2000000000
    gasLimit: web3.utils.toHex(21000),
    to: to,
    value: web3.utils.toHex(web3.utils.toWei(value, 'ether'))
  }
  const tx = new Tx(rawTx)
  tx.sign(privateKey)
  const serializedTx = tx.serialize()
  console.log(serializedTx.toString('hex'))
  return web3.eth.sendSignedTransaction('0x' + serializedTx.toString('hex'), callback)
}

export { sendEther }
