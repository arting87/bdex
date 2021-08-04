pragma solidity >=0.4.21 <0.6.0;

import "../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./ExchangeOwners.sol";
import "./ExchangeData.sol";
import "./ERC20.sol";
//交易所安全控制合约
contract ExchangeSafeControl is ExchangeOwners {

    using SafeMath for *;
    bool  public startContract;                              //是否对外开放合约（仅控制用户调用接口）
    ExchangeData internal exchangeDataInstance;              //合约数据实例

    constructor()public{
        startContract = true;
    }
    //设置数据合约
    function setExchangeData(address addr) public isContractOwner(msg.sender) {
        exchangeDataInstance = ExchangeData(addr);
    }
    //暂停合约(最高权限调用)
    function pauseContract() public isContractOwner(msg.sender) {
        require(startContract == true, "ExchangeSafeControl:23");
        startContract = false;
    }
    //重启合约(最高权限调用)
    function restartContract() public isContractOwner(msg.sender) {
        require(startContract == false, "ExchangeSafeControl:28");
        startContract = true;
    }
    //摧毁合约(最高权限调用)
    function destroyContract() public isContractOwner(msg.sender) {
        selfdestruct(address(uint160(contractOwner)));
    }
    //提取合约中所有eth和erc20代币(最高权限调用)
    function withdrawAllContractMoney() public isContractOwner(msg.sender) {
        get_money_eth(address(this).balance);
        get_money_erc20();
    }
    //提取指定数量eth或者erc20代币(最高权限调用)
    function withdrawalSomeContractMoney(uint256 amount, address tokenContract) public isContractOwner(msg.sender) {
        require(amount > 0, "ExchangeSafeControl:42");
        if (tokenContract == address(0)) {
            get_money_eth(amount);
        } else {
            uint256 balanceAmount = ERC20(tokenContract).balanceOf(address(this));
            require(balanceAmount >= amount, "ExchangeSafeControl:47");
            ERC20(tokenContract).transfer(contractOwner, amount);
        }
    }
    //查询合约eth余额或erc20余额(公开接口)
    function getExchangeBalance(address tokenContract) public view returns (uint256) {
        if (tokenContract == address(0)) {
            return address(this).balance;

        } else {
            return ERC20(tokenContract).balanceOf(address(this));
        }
    }
    //转走eth(内部调用)
    function get_money_eth(uint256 amount) internal {
        require(address(this).balance >= amount, "ExchangeSafeControl:62");
        address(uint160(contractOwner)).transfer(amount);
    }

    //转走所有erc20(内部调用)
    function get_money_erc20() internal {
        address[] memory tokenArray = exchangeDataInstance.getTokenArray();
        for (uint256 i = 0; i < tokenArray.length; i++) {
            uint256 amount = ERC20(tokenArray[i]).balanceOf(address(this));
            if (amount > 0) {
                ERC20(tokenArray[i]).transfer(contractOwner, amount);
            }
        }
    }


}
