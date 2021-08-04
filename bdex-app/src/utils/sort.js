const orderSort = function (arr) { // 挂单数组排序，大->小
  for (let j = 0; j < arr.length - 1; j++) {
    for (let i = 0; i < arr.length - j - 1; i++) {
      if (arr[i].price < arr[i + 1].price) {
        let midData = arr[i]
        arr[i] = arr[i + 1]
        arr[i + 1] = midData
      }
    }
  }
  return arr
}
// 涨幅榜的数组排序,专用  大到小
const changeListSort = function (arr) {
  for (let i = 0; i < arr.length - 1; i++) {
    for (let j = 0; j < arr.length - i - 1; j++) {
      if (arr[j].change < arr[j + 1].change) {
        let mid = arr[j] // 中间变量
        arr[j] = arr[j + 1]
        arr[j + 1] = mid
      }
    }
  }
  let newArr = arr.filter(el => el.change >= 0)
  return newArr.slice(0, 10)
}

// 成交量榜的数组排序, 专用 大到小
const volListSort = function (arr, ethPrice, bdePrice) {
  // 成交量统一为EOS
  for (let i = 0; i < arr.length; i++) {
    if (arr[i].code === '2') {
      const vol = (arr[i].vol * ethPrice).toFixed(4)
      arr[i].vol = parseFloat(vol)
    } else if (arr[i].code === '3') {
      const vol = (arr[i].vol * bdePrice * ethPrice).toFixed(4)
      arr[i].vol = parseFloat(vol)
    }
  }
  for (let i = 0; i < arr.length - 1; i++) {
    for (let j = 0; j < arr.length - i - 1; j++) {
      if (arr[j].vol < arr[j + 1].vol) {
        let mid = arr[j] // 中间变量
        arr[j] = arr[j + 1]
        arr[j + 1] = mid
      }
    }
  }
  let newArr = arr.filter(el => el.vol >= 0)
  return newArr.slice(0, 10)
}
export {orderSort, changeListSort, volListSort}
