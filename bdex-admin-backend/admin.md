## 1

**简要描述：** 

- admin登陆接口
用户身份验证：采用cookie
name:用户名
pwd：密码
用户登陆成功后，写入cookie，后端所有请求根据cookie验证用户身份权限

**请求URL：** 
- ` http://xx.com/api/admin/login `

**请求方式：**

- POST 

**参数：** 

| 参数名    | 必选 | 类型   | 说明   |
| :-------- | :--- | :----- | ------ |
| name      | 是   | string | 用户名 |
| pass_word | 是   | string | 密码   |

 **返回示例**

``` 
  {
    "error_code": 0,
    "data": {
      "success": true,
	  "prmission": 1
    }
  }
```

 **返回参数说明** 

| 参数名     | 类型 | 说明                        |
| :--------- | :--- | --------------------------- |
| success    | bool | 登陆状态                    |
| permission | int  | 1. 超级管理员 2. 普通管理员 |



## 2

**简要描述：** 

- 添加token交易（最高用户权限）

**请求URL：** 

- ` http://xx.com/api/admin/addToken `
  

**请求方式：**
- POST 

**参数：** 

| 参数名         | 必选 | 类型   | 说明             |
| :------------- | :--- | :----- | ---------------- |
| symbol         | 是   | string | 代币名称         |
| token_contract | 是   | string | 代币地址         |
| decimal        | 是   | int    | 代币精度         |
| market_code    | 是   | int    | 基本交易对的类型 |



 **返回示例**

``` 
  {
    "error_code": 0,
    "data": {
		"success": true
    }
  }
```

 **返回参数说明** 

| 参数名  | 类型 | 说明         |
| :------ | :--- | ------------ |
| success | bool | 添加是否成功 |

## 3


**简要描述：** 

- 删除交易对（最高用户权限）

**请求URL：** 
- ` http://xx.com/api/admin/deleteToken `
  

**请求方式：**
- POST 

**参数：** 

|参数名|必选|类型|说明|
| :------------ | :--- | :----- | ---------------- |
|token_contract |是  |string |代币地址   |
|market_code |是  |int |基本交易对的类型   |


 **返回示例**

``` 
  {
    "error_code": 0,
    "data": {
		"success": true
    }
  }
```

 **返回参数说明** 

| 参数名  | 类型 | 说明         |
| :------ | :--- | ------------ |
| success | bool | 删除是否成功 |

## 4


**简要描述：** 

- 查找代币

**请求URL：** 
- ` http://xx.com/api/admin/getToken`
  

**请求方式：**
- POST 

**参数：** 

| 参数名         | 必选 | 类型   | 说明             |
| :------------- | :--- | ------ | ---------------- |
| market_code    | 否   | int    | 基本交易对的类型 |
| token_contract | 否   | string | 代币地址         |
| symbol         | 否   | string | 代币名称         |
| decimal        | 否   | int    | 代币精度         |
| offset         | 否   | int    | 偏移量           |
| limit          | 否   | int    | 取的个数         |


 **返回示例**

``` 
  {
    "error_code": 0,
    "data": {
		    "id":1,
			"token_id":1,
			"symbol":"SOS",
			"token_contract":"0xC18cc1E5BE460DE4FBa8eD50d7E7979dFA7acC17",
			"market_code":2,
			"decimal":18,
			"order_decimal":8,
			"is_cancel":true
    }
  }
```

 **返回参数说明** 

| 参数名         | 类型   | 说明                                  |
| :------------- | :----- | ------------------------------------- |
| id             | int    | 代币在数据库中唯一id                  |
| token_id       | uint64 | 代币所在链上id                        |
| symbol         | string | 代币名称                              |
| token_contract | string | 代币地址                              |
| market_code    | int    | 代币类型                              |
| decimal        | int    | 代币精度                              |
| order_decimal  | int    | 数据库代币处理精度                    |
| is_cancel      | bool   | true是取消了的代币，false是上了的代币 |

## 5

**简要描述：** 

- 从数据库中删除订单

**请求URL：** 
- ` http://xx.com/api/admin/dbDeleteOrder `
  

**请求方式：**

- POST 

**参数：** 

| 参数名   | 必选 | 类型   | 说明   |
| :------- | :--- | :----- | ------ |
| order_id | 是   | string | 订单id |


 **返回示例**

``` 
  {
    "error_code": 0,
    "data": {
      "success": true
    }
  }
```

 **返回参数说明** 

| 参数名  | 类型 | 说明             |
| :------ | :--- | ---------------- |
| success | bool | 删除订单是否成功 |

## 6

**简要描述：分页查询所有订单** 

- 分页查询所有指定条件的订单，根据操作按照指定内容查询

**请求URL：** 
- ` http://xx.com/api/admin/getAllDeleteOrder `
  

**请求方式：**
- POST 

**参数：** 

| 参数名        | 必选 | 类型    | 说明                                                         |
| :------------ | :--- | :------ | ------------------------------------------------------------ |
| time_start    | 否   | string  | 查询时间区间的开始时间                                       |
| time_end      | 否   | string  | 查询时间区间的结束时间                                       |
| time_type     | 否   | bool    | 标注时间区间是订单创建时间还是订单更新时间                   |
| market_type   | 否   | int     | 交易类型    1=eos交易区内，2=eth交易区内，3=bde交易区内，4=跨链交易 |
| match_symbol  | 否   | string  | 交易对名称                                                   |
| price_start   | 否   | float64 | 价格区间的开始                                               |
| price_end     | 否   | float64 | 价格区间的结束                                               |
| withdraw_able | 否   | bool    | 订单的可否撤销状态                                           |
| success       | 否   | bool    | 订单的是否完成状态                                           |
| status        | 否   | int     | 订单状态                                                     |
| name          | 否   | string  | 订单发送者名称                                               |

 **返回示例**

``` 
  {
    "error_code": 0,
    "data": {
	[
		{
			"order_id":213213123213123,
			"name":"abcd",
			"name_chain_code":"1",
			"trade_type":"1",
			"market_type":"1",
			"market_symbol":"ETH,8",
			"dele_vol":123.654,
			"sell_token_symbol":"ETH,8",
			"sell_token_amount":396,
			"sell_token_contract":"0x0000000000000000000000000000000000000000",
			"buy_token_symbol":"ERC,8",
			"buy_token_amount":39878,
			"buy_token_contract":"0x1234500000000000000000000000000000000000",
			"price":123456789,
			"fee":0.000456，
			"withdraw_able":0,
			"ex_success":1,
			"target_address":"0x60C278c444f4a0a65441f1bE51a1B63DE608FBC0",
			"target_chain_code":1,
			"create_at":"2019-07-11 11:56:40",
			"update_at":"2019-07-11 12:36:29",
			"tx_hash":"0x8d6c03bfd79be6a51662c41c57cf98e11f7531dd49253b76b1185dad0d4cfd06",
			"status":2
		}
	]
    }
  }
```

 **返回参数说明** 

| 参数名              | 类型    | 说明                                                         |
| :------------------ | :------ | ------------------------------------------------------------ |
| order_id            | string  | 订单Id                                                       |
| name                | string  | 订单发送者                                                   |
| name_chain_code     | int     | 发送者属于哪一条链                                           |
| trade_type          | int     | 交易类型 1:卖单 0：买单                                      |
| market_type         | int     | 市场类型 1:eos 2:eth 3:Bdex 4:跨链                           |
| market_symbol       | string  | 无重复交易对名称 EOS,1:EOS链上的EOS代币                      |
| dele_vol            | float64 | 下单时的总量                                                 |
| sell_token_symbol   | string  | 卖出的代币名称                                               |
| sell_token_amount   | float64 | 卖出的代币数量                                               |
| sell_token_contract | string  | 卖出的代币合约                                               |
| buy_token_symbol    | string  | 买的代币名称                                                 |
| buy_token_amount    | float64 | 买的代币数量                                                 |
| buy_token_contract  | string  | 买的代币合约                                                 |
| price               | float64 | 下单价格                                                     |
| fee                 | float64 | 目前扣除的手续费                                             |
| withdraw_able       | bool    | 是否可撤单                                                   |
| ex_success          | bool    | 订单是否完成                                                 |
| target_address      | string  | 接受代币的账户地址                                           |
| target_chain_code   | int     | 接受地址属于那一条链                                         |
| create_at           | string  | 订单的创建时间                                               |
| update_at           | time    | 订单的更新时间                                               |
| tx_hash             | time    | 订单下单的交易Hash                                           |
| status              | int     | 订单当前的状态  1:完全撮合成功 2:未完全撮合 3:已被撤单 4:下单失败 5:下单成功 6:错误订单 |



## 7

**简要描述：24小时内交易所状态** 

- 展示指定日期的当天交易所状态，如果未指定日期或日期错误则展示当天的状态

**请求URL：** 

- ` http://xx.com/api/admin/getDataInDay `

**请求方式：**
- POST 

**参数：** 

| 参数名 | 必选 | 类型 | 说明       |
| :----- | :--- | :--- | ---------- |
| day    | 否   | time | 查询的日期 |

 **返回示例**

``` 
  {
    "error_code": 0,
    "data": {
      "order_count": 22,
      "eos_in_deal_count": 12345,
	  "eos_Cross_Deal_Count": 12345,
	  "eth_in_deal_count": 12345,
	  "bdex_in_deal_count": 12345,
      "err_order_count": 0,
      "withdraw_order_count": 213 ,
      "deal_order_count": 555
    }
  }
```

 **返回参数说明** 

| 参数名               | 类型 | 说明                         |
| :------------------- | :--- | ---------------------------- |
| order_count          | int  | 当天24小时内下单数量         |
| eos_in_deal_count    | int  | 当天24小时内eos链内交易数量  |
| eos_Cross_Deal_Count | int  | 当天24小时内eos跨链交易数量  |
| eth_in_deal_count    | int  | 当天24小时内eth链内交易数量  |
| bde_in_deal_count    | int  | 当天24小时内 bde链内交易数量 |
| err_order_count      | int  | 当天24小时内错误订单数量     |
| withdraw_order_count | int  | 当天24小时内撤单数量         |
| deal_order_count     | int  | 当天24小时内完成订单数量     |

 

## 8

**简要描述：获取所有新闻

**请求URL：** 

- ` http://xx.com/api/admin/getNews `

**请求方式：**

- get

**参数：** 无

 **返回示例**

```
  {
    "error_code": 0,
    "data": [{
      "id": 1,
      "title": "惊天秘闻",
	  "typeCode": 1,
	  "content": "",
	  "time": “2019-07-11 11:56:40”
    }]
  }
```

 **返回参数说明** 

| 参数名    | 类型   | 说明                                                     |
| :-------- | :----- | -------------------------------------------------------- |
| id        | int    | 新闻唯一id                                               |
| title     | string | 新闻标题                                                 |
| type_code | int    | 新闻类型；1 重要公告，2 活动公告，3 新币上线，4 其它公告 |
| content   | string | 新闻内容                                                 |
| time      | time   | 时间                                                     |



## 9

**简要描述：添加新闻（最高用户权限）

**请求URL：** 

- ` http://xx.com/api/admin/addNews `

**请求方式：**

- post

**参数：** 

**参数：** 

| 参数名    | 必选 | 类型   | 说明     |
| :-------- | :--- | :----- | -------- |
| title     | 是   | string | 新闻标题 |
| type_code | 是   | int    | 新闻类型 |
| content   | 是   | string | 新闻内容 |

 **返回示例**

```
   {
    "error_code": 0,
    "data": {
      "success": true
    }
  }
```

 **返回参数说明** 

| 参数名  | 类型 | 说明             |
| :------ | :--- | ---------------- |
| success | bool | 添加新闻是否成功 |

## 10

**简要描述：删除新闻（最高用户权限）

**请求URL：** 

- ` http://xx.com/api/admin/deleteNews `

**请求方式：**

- post

**参数：** 

**参数：** 

| 参数名 | 必选 | 类型 | 说明   |
| :----- | :--- | :--- | ------ |
| id     | 是   | int  | 新闻id |

 **返回示例**

```
   {
    "error_code": 0,
    "data": {
      "success": true
    }
  }
```

 **返回参数说明** 

| 参数名  | 类型 | 说明             |
| :------ | :--- | ---------------- |
| success | bool | 删除新闻是否成功 |

**备注** 

- 更多返回错误代码请看首页的错误代码描述







## 11

**简要描述：查询所有后台管理用户（最高用户权限）

**请求URL：** 

- ` http://xx.comapi/admin/user/getUser `

**请求方式：**

- get

**参数：** 

**参数：** 无

 **返回示例**

```
  "error_code": 0,
  "data": [
    {
      "id": 1,
      "name": "admin",
      "pass_word": "123456",
      "permission": 1
    },
    {
      "id": 2,
      "name": "admin01",
      "pass_word": "123456",
      "permission": 1
    },
    {
      "id": 3,
      "name": "user01",
      "pass_word": "123456",
      "permission": 2
    },
    {
      "id": 4,
      "name": "test01",
      "pass_word": "123456",
      "permission": 2
    }
  ]
}
```

 **返回参数说明** 

| 参数名     | 类型   | 说明                         |
| :--------- | :----- | ---------------------------- |
| id         | int    | 用户唯一id，数据库自增       |
| name       | string | 用户名                       |
| pass_word  | string | 密码                         |
| permission | int    | 权限等级   1=最高级 ，2=普通 |



## 12

**简要描述：删除用户（最高用户权限）

**请求URL：** 

- ` http://xx.com/api/admin/user/delete `

**请求方式：**

- post

**参数：** 

**参数：** 

| 参数名 | 必选 | 类型 | 说明   |
| :----- | :--- | :--- | ------ |
| id     | 是   | int  | 用户id |

 **返回示例**

```
   {
    "error_code": 0,
    "data": {
      "success": true
    }
  }
```

 **返回参数说明** 

| 参数名  | 类型 | 说明             |
| :------ | :--- | ---------------- |
| success | bool | 删除用户是否成功 |



## 13

**简要描述：更新用户（最高用户权限）

**请求URL：** 

- ` http://xx.com/api/admin/user/update `

**请求方式：**

- post

**参数：** 

**参数：** 

| 参数名     | 必选 | 类型   | 说明                                      |
| :--------- | ---- | :----- | ----------------------------------------- |
| id         | 是   | int    | 用户唯一id，数据库自增                    |
| name       | 是   | string | 用户名   唯一（注意，不能重名，否则报错） |
| pass_word  | 是   | string | 密码                                      |
| permission | 是   | int    | 权限等级   1=最高级 ，2=普通              |



```
   {
  "error_code": 0,
  "data": {
    "success": true,
    "permission": 1
  }
}
```

 **返回参数说明** 

| 参数名  | 类型 | 说明         |
| :------ | :--- | ------------ |
| success | bool | 跟新是否成功 |





## 14

**简要描述：创建用户（最高用户权限）

注意：

**请求URL：** 

- ` http://xx.com/api/admin/user/create `

**请求方式：**

- post

**参数：** 

**参数：** 

| 参数名     | 必选 | 类型   | 说明                                      |
| :--------- | ---- | :----- | ----------------------------------------- |
| name       | 是   | string | 用户名   唯一（注意，不能重名，否则报错） |
| pass_word  | 是   | string | 密码                                      |
| permission | 是   | int    | 权限等级   1=最高级 ，2=普通              |



 **返回示例**

```
  {
  "error_code": 0,
  "data": {
    "success": true,
    "permission": 1
  }
}
```

 **返回参数说明** 

| 参数名  | 类型 | 说明             |
| :------ | :--- | ---------------- |
| success | bool | 创建用户是否成功 |

**备注** 

- 更多返回错误代码请看首页的错误代码描述