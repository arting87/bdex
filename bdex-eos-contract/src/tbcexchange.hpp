//#include <eosio/asset.hpp>
//#include <eosio/eosio.hpp>
#include <eosio.token/eosio.token.hpp>
#include <eosio/system.hpp>
#include <nlohmann/json.hpp>
#include <cmath>

using namespace std;
using namespace eosio;
using json = nlohmann::json;

class [[eosio::contract]] tbcexchange : public eosio::contract {

  private:

  struct [[eosio::table]] deleorder{
    uint64_t     order_id;
    name         account_name;          //用户名
    bool         trade_type;            //交易类型 0=>买 1=>卖
    asset        sell_token;            //想要卖出多少代币，此数量代币将转入合约中
    string       sell_token_contract;   //卖出token所属合约
    double       price;                 //用当前代币购买代币的价格 price=sell/buy
    asset        buy_token;             //想要购买的代币数量
    string       buy_token_contract;    //购买token所属合约
    string         target_address;        //用户购买代币的地址
    asset        fee;                   //已扣手续费
    bool         withdrawable;          //是否可以撤销订单
    bool         ex_success;            //交易是否完成
    uint32_t       create_at;             //订单创建时间
    uint32_t       update_at;             //订单修改时间

    auto primary_key() const { return order_id; }
  };

  struct [[eosio::table]] token{
    uint64_t id;
    uint64_t tcontract;
    string   symbol;
    string   decimal;
    asset    balance;
    bool     iscancle;

    auto primary_key() const { return id; }
    uint64_t by_contract() const {return tcontract; }
  };

  struct [[eosio::table]] submitback{
    uint64_t order_id;      //订单Id
    name     to;       //转给谁
    asset    sell; //此订单需要转出的代币数量
    string   sell_contract;
    asset    sell_before;
    asset    buy_before;
    asset    fee_before;
    bool     withdrawable;
    bool     ex_success;
    string   flag;          //订单标记

    auto primary_key() const { return order_id; }
  };


  typedef eosio::multi_index<"deleorder"_n, deleorder> dotable;
  typedef eosio::multi_index<"submitback"_n, submitback> smbacktable;
  typedef eosio::multi_index<"token"_n, token, indexed_by<"tcontract"_n, const_mem_fun<token, uint64_t, &token::by_contract>>> totable;

  dotable _dotable;
  totable _totable;
  smbacktable _smtable;

  public:

    tbcexchange(name receiver, name code, datastream<const char*> ds):contract(receiver, code, ds),
                                    _dotable( receiver, receiver.value ),
                                    _totable( receiver, receiver.value ),
                                    _smtable( receiver, receiver.value ){}

    void transfer(name from, name to, asset quantity, string memo);

    //撮合成功提交交易
    [[eosio::action]]
    void submit(name from, uint64_t order_id, name to, asset count, string asset_contract, uint64_t orderb_id, string dealprice);

    //撤销委托
    [[eosio::action]]
    void withdraw(name from, uint64_t order_id );

    //修改撤销状态
    [[eosio::action]]
    void upstate(name from, uint64_t ordera_id,uint64_t orderb_id, bool withdrawable);

    [[eosio::action]]
    void addtoken( string token_contract,string symbol,string decimal);

    [[eosio::action]]
    void deletoken( string id);

    [[eosio::action]]
    void drawallfee(string boss);

    [[eosio::action]]
    void drawallasset(string boss);

    [[eosio::action]]
    void deleteall(int a);

    [[eosio::action]]
    void transneed( uint64_t order_id );

    [[eosio::action]]
    void backorder( uint64_t order_id );
};
