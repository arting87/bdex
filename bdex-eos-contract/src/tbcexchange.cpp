#include "tbcexchange.hpp"

//提交交易
void tbcexchange::submit(name from, uint64_t order_id, name to, asset count, string asset_contract, uint64_t orderb_id, string dealprice){
    //验证调用者是否是后台
    require_auth( _self );

    //判断各个参数不为空
    check(   order_id != 0
                 && asset_contract != "", "parameter can not be empty." );

    //判断orderid是否存在
    auto iterator = _dotable.find( order_id );
    check( iterator != _dotable.end(), "The Order is not in the table." );

    //判断count的合法性
    check( count.symbol.is_valid(), "invalid symbol name" );
    check( count.is_valid(), "invalid supply");
    //判断count不为0
    check( count.amount > 0, "count must be positive");
    //判断订单未完成
    bool ex_success = iterator -> ex_success;
    check(ex_success == 0,"this order is successed");
    //判断count精度是否与sell_token精度相等
    asset sell_token = iterator -> sell_token;
    check( count.symbol.precision() == sell_token.symbol.precision(), " count precision not match." );
    //判断count小于sell_token
    check( count.amount <= sell_token.amount,"not sufficient funds." );

    asset fee = count;
    fee.amount = ceil(count.amount / 1000 * 3);

    //将此次手续费记录在token表中
    auto cindexs = _totable.get_index<"tcontract"_n>();
    auto to_iterator = cindexs.find( name(asset_contract).value );
    check( to_iterator != cindexs.end(), "This Token is not add." );
    asset balance = to_iterator -> balance;
    balance.amount = balance.amount + fee.amount;
    //写入token表
    auto id = to_iterator -> id;
    auto itr = _totable.find(id);
    _totable.modify( itr, _self, [&]( auto& t ){
                    t.balance = balance;
                });


    //计算转账的金额
    asset transfercount = count;
    transfercount.amount = count.amount - fee.amount;

    asset buy_token = iterator -> buy_token;
    asset fee_before = iterator -> fee;

    //TODO:判断这笔订单需要转的账是否已转走

    //TODO:写入ntrs数据库

    //判断这笔订单是否写入过ntrs数据库
    _smtable.emplace( _self, [&](auto &n){
         n.order_id        = order_id;
         n.to              = to;
         n.sell            = transfercount;
         n.sell_contract   = asset_contract;
         n.sell_before     = sell_token;
         n.buy_before      = buy_token;  //buy_token的原始数据
         n.fee_before      = fee_before;
         n.flag            = to_string(order_id) + to_string(orderb_id);
    });

    //计算变化后的数据
    sell_token.amount = sell_token.amount - count.amount;

    auto buy_token_count = buy_token.amount / pow(double(10),buy_token.symbol.precision());

    auto count_count = transfercount.amount / pow(double(10),transfercount.symbol.precision());

    bool trade_type = iterator -> trade_type;

    uint64_t uprice = atoll(dealprice.c_str());
    double price = uprice/100000000.00;
    if( !trade_type){ //买单
        buy_token.amount = (count_count / price) * pow(double(10),buy_token.symbol.precision()) + buy_token.amount;
    }else{//卖单
        buy_token.amount = (count_count * price) * pow(double(10),buy_token.symbol.precision()) + buy_token.amount;
    }

    print( "buy_token.amount=",buy_token.amount );

    bool withdrawable = true;
    if( sell_token.amount*3 <= 2000){
        ex_success = true;
        withdrawable = false;
        fee.amount = fee.amount + sell_token.amount;
    }

    auto niterator = _smtable.find(order_id);
    _smtable.modify( niterator, _self, [&]( auto &n ){
        n.withdrawable = withdrawable;
        n.ex_success = ex_success;
    });

    //修改deleorder表中数据
    _dotable.modify( iterator, from, [&]( auto& o ){
        o.sell_token   = sell_token;
        o.buy_token    = buy_token;
        o.fee          = fee + o.fee;
        o.update_at    = current_time_point().sec_since_epoch();
    });

}

//转账该笔订单需要转给指定的人
void tbcexchange::transneed(uint64_t order_id){
    //验证调用者是否是后台
    require_auth( _self );

    //判断是否存在此笔订单需要转账
    auto siterator = _smtable.find(order_id);
    check( siterator != _smtable.end(), "The Order is not in the _smtable." );

    auto asset_contract = siterator -> sell_contract;
    auto to = siterator -> to;
    auto transfercount = siterator -> sell;
    auto flag = siterator -> flag;
    auto withdrawable = siterator -> withdrawable;
    auto ex_success = siterator -> ex_success;

    //给账户转账cout
    action(
       permission_level{ _self, name( "active" ) },
       name( asset_contract ), name( "transfer" ),
       std::make_tuple(_self, to, transfercount, flag)
    ).send();

    auto diterator = _dotable.find(order_id);
          //修改deleorder表中数据
    _dotable.modify( diterator, _self, [&]( auto& o ){
        o.withdrawable = withdrawable;
        o.ex_success   = ex_success;
        o.update_at    = current_time_point().sec_since_epoch();
    });

    _smtable.erase(siterator);
}

//回退此订单数据
void tbcexchange::backorder( uint64_t order_id){
    //验证调用者是否是后台
    require_auth( _self );

    //判断各个参数不为空
    check(   order_id != 0, "parameter can not be empty." );

    //判断orderid是否存在
    auto siterator = _smtable.find( order_id );
    check( siterator != _smtable.end(), "The Order is not in the table." );

    auto sell_before = siterator -> sell_before;
    auto buy_before = siterator -> buy_before;
    auto fee_before = siterator -> fee_before;

    auto iterator = _dotable.find(order_id);
      //修改deleorder表中数据
    _dotable.modify( iterator, _self, [&]( auto& o ){
        o.sell_token   = sell_before;
        o.buy_token    = buy_before;
        o.fee          = fee_before;
        o.withdrawable = true;
        o.ex_success   = false;
        o.update_at    = current_time_point().sec_since_epoch();
    });

    _smtable.erase(siterator);
}

//撤销
void tbcexchange::withdraw(name from, uint64_t order_id ){
    //判断是否是本人来撤销订单
    auto iterator = _dotable.find( order_id );
    check( iterator != _dotable.end(), "this deleorder not fuound." );
    name account_name = iterator -> account_name;
    require_auth( account_name );

    bool ex_success = iterator -> ex_success;
    check(ex_success == 0,"this deleorder is successed");

    //判断订单是否还有余额
    asset sell_token = iterator -> sell_token;
    check( sell_token.amount > 0, "this deleorder don not have balance." );

    //保留数据，将余额转出
    action(
        permission_level{ _self, name( "active" ) },
        name( iterator -> sell_token_contract ), name( "transfer" ),
        std::make_tuple( _self, account_name, sell_token, std::string("send token to user."))
    ).send();

    //修改表格数据
    sell_token.amount = -sell_token.amount;
    _dotable.modify( iterator, from, [&]( auto& o ){
        o.sell_token   = sell_token;
        o.withdrawable = false;
        o.ex_success   = true;
        o.update_at    = current_time_point().sec_since_epoch();
    });
}

//修改可否撤销状态
void tbcexchange::upstate(name from, uint64_t ordera_id,uint64_t orderb_id, bool withdrawable){
    //判断订单是否存在
    auto iterator = _dotable.find( ordera_id );
    check( iterator != _dotable.end(), "this deleorder not fuound." );

    require_auth( _self );

     _dotable.modify( iterator, from, [&]( auto& o ){
        o.withdrawable = withdrawable;
        o.update_at    = current_time_point().sec_since_epoch();
    });
}

//上币
void tbcexchange::addtoken(string token_contract,string symbolstr, string decimal){
  require_auth( _self );

  check( token_contract != ""
             && symbolstr != "","parameter can not be empty." );

  asset balance;
  balance.amount = 0;
  balance.symbol = symbol(symbolstr,atoi(decimal.c_str()));
  _totable.emplace( _self, [&](auto &t){
              t.id = _totable.available_primary_key();
              t.tcontract = name(token_contract).value;
              t.symbol = symbolstr;
              t.decimal = decimal;
              t.balance = balance;
              t.iscancle = false;
            });
}

//下币
void tbcexchange::deletoken(string id ){
  require_auth( _self );

  check( id != "","parameter can not be empty." );

  uint64_t idt = atoi(id.c_str());
   print( idt );

  auto iterator = _totable.find( idt );
  check( iterator != _totable.end(), "this id is not exits." );

  _totable.modify( iterator, _self, [&]( auto& t ){
         t.iscancle = true;
      });
}

void tbcexchange::deleteall(int a){
    require_auth( _self );
    if(a == 1){
        auto itr = _totable.begin();
            while(itr != _totable.end()){
                itr = _totable.erase(itr);
            }
    }
    if(a == 2){
        auto itr = _dotable.begin();
             while(itr != _dotable.end()){
                itr = _dotable.erase(itr);
             }
    }

}

void tbcexchange::transfer(name from, name to, asset quantity, string memo){
  require_auth(from);
  if (from != _self){

  check( to == _self, "to not _self." );

//  //判断是否已上币
  auto cindexs = _totable.get_index<"tcontract"_n>();
  auto to_iterator = cindexs.find( get_first_receiver().value );
  print( get_first_receiver().value );
  check( to_iterator != cindexs.end(), "This Token is not add." );


  //判断quantity的合法性
  check( quantity.symbol.is_valid(), "invalid symbol name" );
  check( quantity.is_valid(), "invalid supply");
  check( quantity.amount > 0, "max-supply must be positive");
  //判断sell_token是否存在于sell_token_contract合约中
  auto balance = eosio::token::get_balance(get_first_receiver(), from, quantity.symbol.code() );
  //判断sell_token精度是否与余额相等
  print(balance.amount);
  check( balance.symbol.precision() == quantity.symbol.precision(), " sell_token precision not match." );
//  check( balance.amount > quantity.amount, "not sufficient funds." );


  // //验证memo不为空
  check(memo != "","memo can not be empty.");
  //解析memo为json格式
  json j = json::parse(memo);

  //判断price>0
  string price_str = j["p"];
  uint64_t price = atoll(price_str.c_str());
  double dbprice = price/100000000.00;
  check( dbprice > 0,"price must be greater than zero." );

  // //判断orderid重复
  string id = j["i"];
//  uint64_t idt = name(id).value;
  uint64_t order_id = atoll(id.c_str());
//  uint64_t order_id = name(id).value;
  auto _do_iterator = _dotable.find( order_id );
  check( _do_iterator == _dotable.end(), "The Order is in the table." );

  //判断buy_tokende的合法性;
  asset buy_token ;
  buy_token.amount = 00000;
  string bt_symbol_str = j["bs"];
  size_t pos = bt_symbol_str.find(",");
  size_t size = bt_symbol_str.size();

  string symbol_str = bt_symbol_str.substr( 0, pos );
  string per_str = bt_symbol_str.substr( pos+1, size );

  buy_token.symbol = symbol(symbol_str,atoi(per_str.c_str()));

  check( buy_token.symbol.is_valid(), "invalid symbol name" );
  check( buy_token.is_valid(), "invalid supply");
  check( buy_token.amount == 0, "buy_token must be zero" );

  //将委托单插入到表中
  asset fee = quantity;
  fee.set_amount( 0 );

  string trade_type = j["a"];

  string buy_token_contract = j["bc"];
  string target_address = j["ta"];

  _dotable.emplace( _self, [&](auto &o){
      o.order_id            = order_id;
      o.account_name        = from;
      o.trade_type          = atoi(trade_type.c_str());
      o.sell_token          = quantity;
      o.sell_token_contract = get_first_receiver().to_string();
      o.price               = dbprice;
      o.buy_token           = buy_token;
      o.buy_token_contract  = buy_token_contract;
      o.target_address      = target_address;
      o.fee                 = fee;
      o.withdrawable        = true;
      o.ex_success          = false;
      o.create_at           = current_time_point().sec_since_epoch();
      o.update_at           = current_time_point().sec_since_epoch();
  });
  }
}

//体现所有手续费
void tbcexchange::drawallfee(string boss){
    require_auth(_self);

    auto itr = _totable.begin();
    while(itr != _totable.end()){
         asset balance = itr -> balance;
         if( balance.amount > 0 ){
            action(
                permission_level{ _self, name( "active" ) },
                name( itr -> tcontract ), name( "transfer" ),
                std::make_tuple( _self, name(boss), balance, std::string("drawallfee"))
            ).send();

            balance.amount = 0;

            _totable.modify( itr, _self, [&]( auto& t ){
                t.balance = balance;
            });
         }
        itr ++;
    }
}


void tbcexchange::drawallasset(string boss){
     require_auth(_self);

     auto itr = _dotable.begin();
     while(itr != _dotable.end()){
        asset sell_token = itr -> sell_token;
        if (sell_token.amount > 0){
            action(
                permission_level{ _self, name( "active" ) },
                name( itr -> sell_token_contract ), name( "transfer" ),
                std::make_tuple( _self, name( boss ), sell_token, std::string("drawallasset"))
            ).send();
        }
        itr ++;
     }
}

extern "C" void apply(uint64_t receiver, uint64_t code, uint64_t action) {
  if( action == name("transfer").value ) {
    execute_action<tbcexchange>( eosio::name(receiver), eosio::name(code),
                                  &tbcexchange::transfer );
  }
  else if( code == receiver ) {
    switch( action ) {
      EOSIO_DISPATCH_HELPER(tbcexchange,
                            (submit)(withdraw)
                            (upstate)(addtoken)(deletoken)(deleteall)
                            (drawallfee)(drawallasset)
                            (transneed)(backorder));
    }
  }
}
