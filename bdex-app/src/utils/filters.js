exports.time = function (txdate, needSeconds = false) {
  const d = new Date(txdate)
  const month = (d.getMonth() + 1) < 10 ? '0' + (d.getMonth() + 1) : '' + (d.getMonth() + 1)
  const day = d.getDate() < 10 ? '0' + d.getDate() : '' + d.getDate()
  const hour = d.getHours() < 10 ? '0' + d.getHours() : '' + d.getHours()
  const minutes = d.getMinutes() < 10 ? '0' + d.getMinutes() : '' + d.getMinutes()
  const seconds = d.getSeconds() < 10 ? '0' + d.getSeconds() : '' + d.getSeconds()
  if (needSeconds) {
    return hour + ':' + minutes + ':' + seconds
  }
  return month + '-' + day + ' ' + hour + ':' + minutes
}

exports.fixed = function (val) {
  if (typeof val === 'string') {
    if (val.indexOf('.') > -1) {
      const valList = val.split('.')
      valList[1] = valList[1].slice(0, 4)
      return valList.join('.')
    } else {
      return val.slice(0, 4)
    }
  }
  return val && val.toFixed(4)
}

exports.getCpu = function (value) {
  return (value / 1000).toFixed(2)
}

exports.getKB = function (value) {
  return (value / 1024).toFixed(2)
}
