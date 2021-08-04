// import {numFloat, handBig} from './handleFloat'
// const filFloat = function (arr) { // 专门处理挂单数据，即orderAPi的
//   for (let i = 0; i < arr.length; i++) {
//     arr[i].price = numFloat(arr[i].price, 6)
//     arr[i].amt = numFloat(arr[i].amt, 4)
//     console.warn('需要计算的值', arr[i].price, arr[i].amt, handBig.times(arr[i].price, arr[i].amt).toString())
//     arr[i].total = numFloat(handBig.times(arr[i].price, arr[i].amt), 4)
//   }
//   return arr
// }
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

export {orderSort}
