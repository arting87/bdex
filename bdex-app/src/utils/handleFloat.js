import Bignumber from 'bignumber.js'
// const Bignumber = require('bignumber.js')
const handleEthPrice = function (price) { // eth price处理，放大e8倍
  let str = price.toString()
  let priceArr = str.split('.')
  let newPrice = null
  if (priceArr.length === 2) {
    let len = priceArr[1].length
    if (len > 8) { // price精度大于8,返回null，提示用户数据错误
      return null
    } else {
      newPrice = parseInt(str.replace('.', '')) // 去掉小数点，转换成整数
      let pow = 8 - len
      newPrice = newPrice * (10 ** pow) // 若是去掉小数点后放大倍数小于e8,则继续放大
      return newPrice
    }
  } else if (priceArr.length === 1) {
    newPrice = price * (10 ** 8) // price为整数时，执行
    return newPrice
  } else {
    return null
  }
}

const handlePrice = function (price, times) { // bde price处理，放大times倍
  let str = price.toString()
  let priceArr = str.split('.')
  let newPrice = null
  if (priceArr.length === 2) {
    let len = priceArr[1].length
    if (len > 8) { // price精度大于8,返回null，提示用户数据错误
      return null
    } else {
      newPrice = parseInt(str.replace('.', '')) // 去掉小数点，转换成整数
      let pow = 8 - len
      newPrice = newPrice * (10 ** pow) // 若是去掉小数点后放大倍数小于e8,则继续放大
      return newPrice
    }
  } else if (priceArr.length === 1) {
    newPrice = price * (10 ** 8) // price为整数时，执行
    return newPrice
  } else {
    return null
  }
}
// 浮点数精度调整处理
// js具有浮点数精度丢失问题，因此使用字符串处理方式控制精度

const numFloat = function (number, len) { // number需要处理的数字，len处理后的精度长度
  if ((typeof number === 'number' || typeof number === 'string') && typeof len === 'number') {
    // number = parseFloat(number)
    len = parseInt(len)
    if (len > 10) {
      len = 10
    }
    if (number || parseFloat(number) === 0) { // 防止number为NaN
      let numStr = number.toString()
      let numArr = numStr.split('.')
      let newNumber = null
      if (len > 0) { // toFixed() 为了防止精度丢失，暂时只支持1-10位精度
        if (numArr.length === 2) {
          let numInt = numArr[0]
          let numFloat = numArr[1]
          if (numFloat.length >= len) {
            newNumber = parseFloat(numInt + '.' + numFloat.substring(0, len))
          } else {
            let zeroLen = len - numFloat.length
            for (let i = 1; i <= zeroLen; i++) {
              numFloat += '0'
            }
            newNumber = parseFloat(numInt + '.' + numFloat)
          }
          return newNumber.toFixed(len)
        } else { // 如果是整数，则小数部分补零
          let numInt = numArr[0]
          let numFloat = null
          for (let j = 1; j <= len; j++) {
            numFloat += '0'
          }
          newNumber = parseFloat(numInt + '.' + numFloat)
          return newNumber.toFixed(len)
        }
      } else { // len <= 0时, 返回整数
        newNumber = parseInt(numArr[0], 10)
        return newNumber
      }
    }
  } else {
    return null
  }
}

const handBig = {
  plus: function (a, b) { // 加法
    let x = new Bignumber(a)
    return parseFloat(x.plus(b).toString())
  },
  minus: function (a, b) { // 减法
    let x = new Bignumber(a)
    return parseFloat(x.minus(b).toString())
  },
  times: function (a, b) { // 乘法
    let x = new Bignumber(a)
    return parseFloat(x.times(b).toString())
  },
  divide: function (a, b) { // 除法
    let x = new Bignumber(a)
    return parseFloat(x.div(b).toString())
  }
}

// console.log(parseFloat(handBig.minus(-32, 21)))
// console.log(handBig.times())
export {handleEthPrice, numFloat, handBig, handlePrice}
