#bdex-ex-backend
1.交易所启动命令
    
    bdex-ex-backend serve
    
2.代币同步命令

当完成eos和eth的代币上币之后，需要使用此命令将合约中的代币同步至后端数据库中，否则下单时会报错
同步完成之后后端需要重启或者等待12个小时后后端会自动同步
   
    bdex-ex-backend sync serve
 


#项目账户配置

1.bdex-ex-backend.yml文件
    
    line 11~12替换为eos合约部署账户的名称和私钥
    line 28替换为mysql数据库的连接
    line 37~58行统一替换为eos合约部署账户的名称和私钥和公钥
    
2.configs/bulk_ether_accounts.json
    
    该文件中的账户是循环提交交易撮合的账户，需要保证这些账户中有足够发送交易的余额
    账户数量任意，账户越多，能够瞬时撮合的交易就越多
    
