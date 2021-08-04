pragma solidity >=0.4.21 <0.6.0;

import "./ExchangeSafeControl.sol";

//交易所合约
contract Exchange is ExchangeSafeControl {
    //提交订单（用户调用，erc20需先授权）
    function depositsOrder(
        uint64 orderId,                       //订单id
        string  memory userReceiver,          //接收地址
        bool tradeType,                      //买卖类型
        uint256 sellAmount,                  //卖的数量
        address sellContract,              //卖方合约地址
        uint256 price,                     //单价
        address buyContract,             //买方合约地址
        uint64 market_code)           //交易类型
    public payable
    {
        require(startContract, "Exchange:19");
        require(bytes(userReceiver).length != 0, "Exchange:20");
        require(sellAmount > 0, "Exchange:21");
        require(price > 0, "Exchange:22");
        require(market_code != 0, "Exchange:23");
        if (sellContract == address(0)) {
            require(msg.value >= sellAmount, "Exchange:25");
        } else {
            require(ERC20(sellContract).allowance(msg.sender, address(this)) >= sellAmount, "Exchange:27");
            ERC20(sellContract).transferFrom(msg.sender, address(this), sellAmount);
        }
        exchangeDataInstance.createOrders(orderId, msg.sender, userReceiver, tradeType, sellAmount, sellContract, price, buyContract, market_code);
    }

    //撤销订单（用户调用）
    function withdrawal(uint64 orderId)
    public
    {
        require(startContract, "Exchange:37");
        (uint256 amount,address userAddress,address tokenContract) = exchangeDataInstance.withdrawOrder(orderId);
        require(userAddress == msg.sender, "Exchange:39");
        if (tokenContract == address(0)) {
            msg.sender.transfer(amount);
        } else {
            ERC20(tokenContract).transfer(msg.sender, amount);
        }
    }
    //更新订单状态(普通权限调用)
    function updateOrderState(uint64 orderId, bool status)
    public
    isOwners(msg.sender)
    {
        exchangeDataInstance.changeWithdrawStatus(orderId, status);
    }
    //撮合后修改订单数据(普通权限调用)
    function submit(uint64 orderId, address to, uint256 sellAmount, uint256 buyAmount, uint256 feeAmount)
    public
    isOwners(msg.sender)
    {
        require(to != address(0), "Exchange:58");
        require(sellAmount > 0 && buyAmount > 0 && feeAmount > 0, "Exchange:59");
        exchangeDataInstance.addSubmit(orderId, to, sellAmount, buyAmount, feeAmount);
    }
    //回滚订单(普通权限调用)
    function rollBackOrderByOrderId(uint64 orderId)
    public
    isOwners(msg.sender)
    {
        exchangeDataInstance.rollbackSubmit(orderId);
    }
    //订单成交并修改数据成功后转账(普通权限调用)
    function TransferMoney(uint64 orderId)
    public isOwners(msg.sender)
    {
        (uint256 sellAmount,address to,address tokenContract) = exchangeDataInstance.getTransferInfo(orderId);
        if (tokenContract == address(0)) {
            address(uint160(to)).transfer(sellAmount);
        } else {
            ERC20(tokenContract).transfer(to, sellAmount);
        }
    }
}
