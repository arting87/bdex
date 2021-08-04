// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainConn

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExchangeDataABI is the input ABI used to generate the binding from.
const ExchangeDataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"rollbackOrderMap\",\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"name\":\"buyAmount\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint64\"}],\"name\":\"rollbackSubmit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint64\"}],\"name\":\"withdrawOrder\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"userAddress\",\"type\":\"address\"},{\"name\":\"tokenContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"name\":\"token_name\",\"type\":\"string\"},{\"name\":\"token_address\",\"type\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"market_code_array\",\"type\":\"uint64[50]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenTable\",\"outputs\":[{\"name\":\"token_name\",\"type\":\"string\"},{\"name\":\"token_address\",\"type\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"name\":\"buyAmount\",\"type\":\"uint256\"},{\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"addSubmit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint64\"},{\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"name\":\"buyAmount\",\"type\":\"uint256\"},{\"name\":\"fees\",\"type\":\"uint256\"},{\"name\":\"withdrawAble\",\"type\":\"bool\"},{\"name\":\"exSuccess\",\"type\":\"bool\"}],\"name\":\"setOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint64\"}],\"name\":\"getTransferInfo\",\"outputs\":[{\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"accessAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"access\",\"type\":\"bool\"}],\"name\":\"setAccessAllow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"setContractOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint64\"},{\"name\":\"userAddress\",\"type\":\"address\"},{\"name\":\"userReceiver\",\"type\":\"string\"},{\"name\":\"tradeType\",\"type\":\"bool\"},{\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"name\":\"sellContract\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"buyContract\",\"type\":\"address\"},{\"name\":\"market_code\",\"type\":\"uint64\"}],\"name\":\"createOrders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token_name\",\"type\":\"string\"},{\"name\":\"token_address\",\"type\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"market_code\",\"type\":\"uint64\"}],\"name\":\"addTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"orderIdToUsers\",\"outputs\":[{\"name\":\"order_id\",\"type\":\"uint64\"},{\"name\":\"user_address\",\"type\":\"address\"},{\"name\":\"user_receiver\",\"type\":\"string\"},{\"name\":\"trade_type\",\"type\":\"bool\"},{\"name\":\"sell_amount\",\"type\":\"uint256\"},{\"name\":\"sell_contract\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"buy_amount\",\"type\":\"uint256\"},{\"name\":\"buy_contract\",\"type\":\"address\"},{\"name\":\"fees\",\"type\":\"uint256\"},{\"name\":\"withdraw_able\",\"type\":\"bool\"},{\"name\":\"ex_success\",\"type\":\"bool\"},{\"name\":\"create_at\",\"type\":\"uint256\"},{\"name\":\"update_at\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token_address\",\"type\":\"address\"},{\"name\":\"market_code\",\"type\":\"uint64\"}],\"name\":\"deleteToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint64\"},{\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"changeWithdrawStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ExchangeData is an auto generated Go binding around an Ethereum contract.
type ExchangeData struct {
	ExchangeDataCaller     // Read-only binding to the contract
	ExchangeDataTransactor // Write-only binding to the contract
	ExchangeDataFilterer   // Log filterer for contract events
}

// ExchangeDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeDataSession struct {
	Contract     *ExchangeData     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeDataCallerSession struct {
	Contract *ExchangeDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExchangeDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeDataTransactorSession struct {
	Contract     *ExchangeDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExchangeDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeDataRaw struct {
	Contract *ExchangeData // Generic contract binding to access the raw methods on
}

// ExchangeDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeDataCallerRaw struct {
	Contract *ExchangeDataCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeDataTransactorRaw struct {
	Contract *ExchangeDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeData creates a new instance of ExchangeData, bound to a specific deployed contract.
func NewExchangeData(address common.Address, backend bind.ContractBackend) (*ExchangeData, error) {
	contract, err := bindExchangeData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeData{ExchangeDataCaller: ExchangeDataCaller{contract: contract}, ExchangeDataTransactor: ExchangeDataTransactor{contract: contract}, ExchangeDataFilterer: ExchangeDataFilterer{contract: contract}}, nil
}

// NewExchangeDataCaller creates a new read-only instance of ExchangeData, bound to a specific deployed contract.
func NewExchangeDataCaller(address common.Address, caller bind.ContractCaller) (*ExchangeDataCaller, error) {
	contract, err := bindExchangeData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeDataCaller{contract: contract}, nil
}

// NewExchangeDataTransactor creates a new write-only instance of ExchangeData, bound to a specific deployed contract.
func NewExchangeDataTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeDataTransactor, error) {
	contract, err := bindExchangeData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeDataTransactor{contract: contract}, nil
}

// NewExchangeDataFilterer creates a new log filterer instance of ExchangeData, bound to a specific deployed contract.
func NewExchangeDataFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeDataFilterer, error) {
	contract, err := bindExchangeData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeDataFilterer{contract: contract}, nil
}

// bindExchangeData binds a generic wrapper to an already deployed contract.
func bindExchangeData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeData *ExchangeDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeData.Contract.ExchangeDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeData *ExchangeDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeData.Contract.ExchangeDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeData *ExchangeDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeData.Contract.ExchangeDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeData *ExchangeDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeData *ExchangeDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeData *ExchangeDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeData.Contract.contract.Transact(opts, method, params...)
}

// AccessAllowed is a free data retrieval call binding the contract method 0x72501976.
//
// Solidity: function accessAllowed(address ) constant returns(bool)
func (_ExchangeData *ExchangeDataCaller) AccessAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ExchangeData.contract.Call(opts, out, "accessAllowed", arg0)
	return *ret0, err
}

// AccessAllowed is a free data retrieval call binding the contract method 0x72501976.
//
// Solidity: function accessAllowed(address ) constant returns(bool)
func (_ExchangeData *ExchangeDataSession) AccessAllowed(arg0 common.Address) (bool, error) {
	return _ExchangeData.Contract.AccessAllowed(&_ExchangeData.CallOpts, arg0)
}

// AccessAllowed is a free data retrieval call binding the contract method 0x72501976.
//
// Solidity: function accessAllowed(address ) constant returns(bool)
func (_ExchangeData *ExchangeDataCallerSession) AccessAllowed(arg0 common.Address) (bool, error) {
	return _ExchangeData.Contract.AccessAllowed(&_ExchangeData.CallOpts, arg0)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() constant returns(address)
func (_ExchangeData *ExchangeDataCaller) ContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExchangeData.contract.Call(opts, out, "contractOwner")
	return *ret0, err
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() constant returns(address)
func (_ExchangeData *ExchangeDataSession) ContractOwner() (common.Address, error) {
	return _ExchangeData.Contract.ContractOwner(&_ExchangeData.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() constant returns(address)
func (_ExchangeData *ExchangeDataCallerSession) ContractOwner() (common.Address, error) {
	return _ExchangeData.Contract.ContractOwner(&_ExchangeData.CallOpts)
}

// GetTokenArray is a free data retrieval call binding the contract method 0xcf37d74e.
//
// Solidity: function getTokenArray() constant returns(address[])
func (_ExchangeData *ExchangeDataCaller) GetTokenArray(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ExchangeData.contract.Call(opts, out, "getTokenArray")
	return *ret0, err
}

// GetTokenArray is a free data retrieval call binding the contract method 0xcf37d74e.
//
// Solidity: function getTokenArray() constant returns(address[])
func (_ExchangeData *ExchangeDataSession) GetTokenArray() ([]common.Address, error) {
	return _ExchangeData.Contract.GetTokenArray(&_ExchangeData.CallOpts)
}

// GetTokenArray is a free data retrieval call binding the contract method 0xcf37d74e.
//
// Solidity: function getTokenArray() constant returns(address[])
func (_ExchangeData *ExchangeDataCallerSession) GetTokenArray() ([]common.Address, error) {
	return _ExchangeData.Contract.GetTokenArray(&_ExchangeData.CallOpts)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address tokenAddress) constant returns(string token_name, address token_address, uint8 decimals, uint64[50] market_code_array)
func (_ExchangeData *ExchangeDataCaller) GetTokenInfo(opts *bind.CallOpts, tokenAddress common.Address) (struct {
	TokenName       string
	TokenAddress    common.Address
	Decimals        uint8
	MarketCodeArray [50]uint64
}, error) {
	ret := new(struct {
		TokenName       string
		TokenAddress    common.Address
		Decimals        uint8
		MarketCodeArray [50]uint64
	})
	out := ret
	err := _ExchangeData.contract.Call(opts, out, "getTokenInfo", tokenAddress)
	return *ret, err
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address tokenAddress) constant returns(string token_name, address token_address, uint8 decimals, uint64[50] market_code_array)
func (_ExchangeData *ExchangeDataSession) GetTokenInfo(tokenAddress common.Address) (struct {
	TokenName       string
	TokenAddress    common.Address
	Decimals        uint8
	MarketCodeArray [50]uint64
}, error) {
	return _ExchangeData.Contract.GetTokenInfo(&_ExchangeData.CallOpts, tokenAddress)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address tokenAddress) constant returns(string token_name, address token_address, uint8 decimals, uint64[50] market_code_array)
func (_ExchangeData *ExchangeDataCallerSession) GetTokenInfo(tokenAddress common.Address) (struct {
	TokenName       string
	TokenAddress    common.Address
	Decimals        uint8
	MarketCodeArray [50]uint64
}, error) {
	return _ExchangeData.Contract.GetTokenInfo(&_ExchangeData.CallOpts, tokenAddress)
}

// OrderIdToUsers is a free data retrieval call binding the contract method 0xd1f769f8.
//
// Solidity: function orderIdToUsers(uint64 ) constant returns(uint64 order_id, address user_address, string user_receiver, bool trade_type, uint256 sell_amount, address sell_contract, uint256 price, uint256 buy_amount, address buy_contract, uint256 fees, bool withdraw_able, bool ex_success, uint256 create_at, uint256 update_at)
func (_ExchangeData *ExchangeDataCaller) OrderIdToUsers(opts *bind.CallOpts, arg0 uint64) (struct {
	OrderId      uint64
	UserAddress  common.Address
	UserReceiver string
	TradeType    bool
	SellAmount   *big.Int
	SellContract common.Address
	Price        *big.Int
	BuyAmount    *big.Int
	BuyContract  common.Address
	Fees         *big.Int
	WithdrawAble bool
	ExSuccess    bool
	CreateAt     *big.Int
	UpdateAt     *big.Int
}, error) {
	ret := new(struct {
		OrderId      uint64
		UserAddress  common.Address
		UserReceiver string
		TradeType    bool
		SellAmount   *big.Int
		SellContract common.Address
		Price        *big.Int
		BuyAmount    *big.Int
		BuyContract  common.Address
		Fees         *big.Int
		WithdrawAble bool
		ExSuccess    bool
		CreateAt     *big.Int
		UpdateAt     *big.Int
	})
	out := ret
	err := _ExchangeData.contract.Call(opts, out, "orderIdToUsers", arg0)
	return *ret, err
}

// OrderIdToUsers is a free data retrieval call binding the contract method 0xd1f769f8.
//
// Solidity: function orderIdToUsers(uint64 ) constant returns(uint64 order_id, address user_address, string user_receiver, bool trade_type, uint256 sell_amount, address sell_contract, uint256 price, uint256 buy_amount, address buy_contract, uint256 fees, bool withdraw_able, bool ex_success, uint256 create_at, uint256 update_at)
func (_ExchangeData *ExchangeDataSession) OrderIdToUsers(arg0 uint64) (struct {
	OrderId      uint64
	UserAddress  common.Address
	UserReceiver string
	TradeType    bool
	SellAmount   *big.Int
	SellContract common.Address
	Price        *big.Int
	BuyAmount    *big.Int
	BuyContract  common.Address
	Fees         *big.Int
	WithdrawAble bool
	ExSuccess    bool
	CreateAt     *big.Int
	UpdateAt     *big.Int
}, error) {
	return _ExchangeData.Contract.OrderIdToUsers(&_ExchangeData.CallOpts, arg0)
}

// OrderIdToUsers is a free data retrieval call binding the contract method 0xd1f769f8.
//
// Solidity: function orderIdToUsers(uint64 ) constant returns(uint64 order_id, address user_address, string user_receiver, bool trade_type, uint256 sell_amount, address sell_contract, uint256 price, uint256 buy_amount, address buy_contract, uint256 fees, bool withdraw_able, bool ex_success, uint256 create_at, uint256 update_at)
func (_ExchangeData *ExchangeDataCallerSession) OrderIdToUsers(arg0 uint64) (struct {
	OrderId      uint64
	UserAddress  common.Address
	UserReceiver string
	TradeType    bool
	SellAmount   *big.Int
	SellContract common.Address
	Price        *big.Int
	BuyAmount    *big.Int
	BuyContract  common.Address
	Fees         *big.Int
	WithdrawAble bool
	ExSuccess    bool
	CreateAt     *big.Int
	UpdateAt     *big.Int
}, error) {
	return _ExchangeData.Contract.OrderIdToUsers(&_ExchangeData.CallOpts, arg0)
}

// RollbackOrderMap is a free data retrieval call binding the contract method 0x108a03b9.
//
// Solidity: function rollbackOrderMap(uint64 ) constant returns(uint64 orderId, address to, uint256 sellAmount, uint256 buyAmount, uint256 fee)
func (_ExchangeData *ExchangeDataCaller) RollbackOrderMap(opts *bind.CallOpts, arg0 uint64) (struct {
	OrderId    uint64
	To         common.Address
	SellAmount *big.Int
	BuyAmount  *big.Int
	Fee        *big.Int
}, error) {
	ret := new(struct {
		OrderId    uint64
		To         common.Address
		SellAmount *big.Int
		BuyAmount  *big.Int
		Fee        *big.Int
	})
	out := ret
	err := _ExchangeData.contract.Call(opts, out, "rollbackOrderMap", arg0)
	return *ret, err
}

// RollbackOrderMap is a free data retrieval call binding the contract method 0x108a03b9.
//
// Solidity: function rollbackOrderMap(uint64 ) constant returns(uint64 orderId, address to, uint256 sellAmount, uint256 buyAmount, uint256 fee)
func (_ExchangeData *ExchangeDataSession) RollbackOrderMap(arg0 uint64) (struct {
	OrderId    uint64
	To         common.Address
	SellAmount *big.Int
	BuyAmount  *big.Int
	Fee        *big.Int
}, error) {
	return _ExchangeData.Contract.RollbackOrderMap(&_ExchangeData.CallOpts, arg0)
}

// RollbackOrderMap is a free data retrieval call binding the contract method 0x108a03b9.
//
// Solidity: function rollbackOrderMap(uint64 ) constant returns(uint64 orderId, address to, uint256 sellAmount, uint256 buyAmount, uint256 fee)
func (_ExchangeData *ExchangeDataCallerSession) RollbackOrderMap(arg0 uint64) (struct {
	OrderId    uint64
	To         common.Address
	SellAmount *big.Int
	BuyAmount  *big.Int
	Fee        *big.Int
}, error) {
	return _ExchangeData.Contract.RollbackOrderMap(&_ExchangeData.CallOpts, arg0)
}

// TokenArray is a free data retrieval call binding the contract method 0xfaa32f5d.
//
// Solidity: function tokenArray(uint256 ) constant returns(address)
func (_ExchangeData *ExchangeDataCaller) TokenArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExchangeData.contract.Call(opts, out, "tokenArray", arg0)
	return *ret0, err
}

// TokenArray is a free data retrieval call binding the contract method 0xfaa32f5d.
//
// Solidity: function tokenArray(uint256 ) constant returns(address)
func (_ExchangeData *ExchangeDataSession) TokenArray(arg0 *big.Int) (common.Address, error) {
	return _ExchangeData.Contract.TokenArray(&_ExchangeData.CallOpts, arg0)
}

// TokenArray is a free data retrieval call binding the contract method 0xfaa32f5d.
//
// Solidity: function tokenArray(uint256 ) constant returns(address)
func (_ExchangeData *ExchangeDataCallerSession) TokenArray(arg0 *big.Int) (common.Address, error) {
	return _ExchangeData.Contract.TokenArray(&_ExchangeData.CallOpts, arg0)
}

// TokenTable is a free data retrieval call binding the contract method 0x22d2ca35.
//
// Solidity: function tokenTable(address ) constant returns(string token_name, address token_address, uint8 decimals)
func (_ExchangeData *ExchangeDataCaller) TokenTable(opts *bind.CallOpts, arg0 common.Address) (struct {
	TokenName    string
	TokenAddress common.Address
	Decimals     uint8
}, error) {
	ret := new(struct {
		TokenName    string
		TokenAddress common.Address
		Decimals     uint8
	})
	out := ret
	err := _ExchangeData.contract.Call(opts, out, "tokenTable", arg0)
	return *ret, err
}

// TokenTable is a free data retrieval call binding the contract method 0x22d2ca35.
//
// Solidity: function tokenTable(address ) constant returns(string token_name, address token_address, uint8 decimals)
func (_ExchangeData *ExchangeDataSession) TokenTable(arg0 common.Address) (struct {
	TokenName    string
	TokenAddress common.Address
	Decimals     uint8
}, error) {
	return _ExchangeData.Contract.TokenTable(&_ExchangeData.CallOpts, arg0)
}

// TokenTable is a free data retrieval call binding the contract method 0x22d2ca35.
//
// Solidity: function tokenTable(address ) constant returns(string token_name, address token_address, uint8 decimals)
func (_ExchangeData *ExchangeDataCallerSession) TokenTable(arg0 common.Address) (struct {
	TokenName    string
	TokenAddress common.Address
	Decimals     uint8
}, error) {
	return _ExchangeData.Contract.TokenTable(&_ExchangeData.CallOpts, arg0)
}

// AddSubmit is a paid mutator transaction binding the contract method 0x41e4be54.
//
// Solidity: function addSubmit(uint64 orderId, address to, uint256 sellAmount, uint256 buyAmount, uint256 feeAmount) returns()
func (_ExchangeData *ExchangeDataTransactor) AddSubmit(opts *bind.TransactOpts, orderId uint64, to common.Address, sellAmount *big.Int, buyAmount *big.Int, feeAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeData.contract.Transact(opts, "addSubmit", orderId, to, sellAmount, buyAmount, feeAmount)
}

// AddSubmit is a paid mutator transaction binding the contract method 0x41e4be54.
//
// Solidity: function addSubmit(uint64 orderId, address to, uint256 sellAmount, uint256 buyAmount, uint256 feeAmount) returns()
func (_ExchangeData *ExchangeDataSession) AddSubmit(orderId uint64, to common.Address, sellAmount *big.Int, buyAmount *big.Int, feeAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeData.Contract.AddSubmit(&_ExchangeData.TransactOpts, orderId, to, sellAmount, buyAmount, feeAmount)
}

// AddSubmit is a paid mutator transaction binding the contract method 0x41e4be54.
//
// Solidity: function addSubmit(uint64 orderId, address to, uint256 sellAmount, uint256 buyAmount, uint256 feeAmount) returns()
func (_ExchangeData *ExchangeDataTransactorSession) AddSubmit(orderId uint64, to common.Address, sellAmount *big.Int, buyAmount *big.Int, feeAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeData.Contract.AddSubmit(&_ExchangeData.TransactOpts, orderId, to, sellAmount, buyAmount, feeAmount)
}

// AddTokens is a paid mutator transaction binding the contract method 0xbf5c4b1e.
//
// Solidity: function addTokens(string token_name, address token_address, uint8 decimals, uint64 market_code) returns()
func (_ExchangeData *ExchangeDataTransactor) AddTokens(opts *bind.TransactOpts, token_name string, token_address common.Address, decimals uint8, market_code uint64) (*types.Transaction, error) {
	return _ExchangeData.contract.Transact(opts, "addTokens", token_name, token_address, decimals, market_code)
}

// AddTokens is a paid mutator transaction binding the contract method 0xbf5c4b1e.
//
// Solidity: function addTokens(string token_name, address token_address, uint8 decimals, uint64 market_code) returns()
func (_ExchangeData *ExchangeDataSession) AddTokens(token_name string, token_address common.Address, decimals uint8, market_code uint64) (*types.Transaction, error) {
	return _ExchangeData.Contract.AddTokens(&_ExchangeData.TransactOpts, token_name, token_address, decimals, market_code)
}

// AddTokens is a paid mutator transaction binding the contract method 0xbf5c4b1e.
//
// Solidity: function addTokens(string token_name, address token_address, uint8 decimals, uint64 market_code) returns()
func (_ExchangeData *ExchangeDataTransactorSession) AddTokens(token_name string, token_address common.Address, decimals uint8, market_code uint64) (*types.Transaction, error) {
	return _ExchangeData.Contract.AddTokens(&_ExchangeData.TransactOpts, token_name, token_address, decimals, market_code)
}

// ChangeWithdrawStatus is a paid mutator transaction binding the contract method 0xf2cf5a1f.
//
// Solidity: function changeWithdrawStatus(uint64 orderId, bool status) returns()
func (_ExchangeData *ExchangeDataTransactor) ChangeWithdrawStatus(opts *bind.TransactOpts, orderId uint64, status bool) (*types.Transaction, error) {
	return _ExchangeData.contract.Transact(opts, "changeWithdrawStatus", orderId, status)
}

// ChangeWithdrawStatus is a paid mutator transaction binding the contract method 0xf2cf5a1f.
//
// Solidity: function changeWithdrawStatus(uint64 orderId, bool status) returns()
func (_ExchangeData *ExchangeDataSession) ChangeWithdrawStatus(orderId uint64, status bool) (*types.Transaction, error) {
	return _ExchangeData.Contract.ChangeWithdrawStatus(&_ExchangeData.TransactOpts, orderId, status)
}

// ChangeWithdrawStatus is a paid mutator transaction binding the contract method 0xf2cf5a1f.
//
// Solidity: function changeWithdrawStatus(uint64 orderId, bool status) returns()
func (_ExchangeData *ExchangeDataTransactorSession) ChangeWithdrawStatus(orderId uint64, status bool) (*types.Transaction, error) {
	return _ExchangeData.Contract.ChangeWithdrawStatus(&_ExchangeData.TransactOpts, orderId, status)
}

// CreateOrders is a paid mutator transaction binding the contract method 0xbda3d1b1.
//
// Solidity: function createOrders(uint64 orderId, address userAddress, string userReceiver, bool tradeType, uint256 sellAmount, address sellContract, uint256 price, address buyContract, uint64 market_code) returns()
func (_ExchangeData *ExchangeDataTransactor) CreateOrders(opts *bind.TransactOpts, orderId uint64, userAddress common.Address, userReceiver string, tradeType bool, sellAmount *big.Int, sellContract common.Address, price *big.Int, buyContract common.Address, market_code uint64) (*types.Transaction, error) {
	return _ExchangeData.contract.Transact(opts, "createOrders", orderId, userAddress, userReceiver, tradeType, sellAmount, sellContract, price, buyContract, market_code)
}

// CreateOrders is a paid mutator transaction binding the contract method 0xbda3d1b1.
//
// Solidity: function createOrders(uint64 orderId, address userAddress, string userReceiver, bool tradeType, uint256 sellAmount, address sellContract, uint256 price, address buyContract, uint64 market_code) returns()
func (_ExchangeData *ExchangeDataSession) CreateOrders(orderId uint64, userAddress common.Address, userReceiver string, tradeType bool, sellAmount *big.Int, sellContract common.Address, price *big.Int, buyContract common.Address, market_code uint64) (*types.Transaction, error) {
	return _ExchangeData.Contract.CreateOrders(&_ExchangeData.TransactOpts, orderId, userAddress, userReceiver, tradeType, sellAmount, sellContract, price, buyContract, market_code)
}

// CreateOrders is a paid mutator transaction binding the contract method 0xbda3d1b1.
//
// Solidity: function createOrders(uint64 orderId, address userAddress, string userReceiver, bool tradeType, uint256 sellAmount, address sellContract, uint256 price, address buyContract, uint64 market_code) returns()
func (_ExchangeData *ExchangeDataTransactorSession) CreateOrders(orderId uint64, userAddress common.Address, userReceiver string, tradeType bool, sellAmount *big.Int, sellContract common.Address, price *big.Int, buyContract common.Address, market_code uint64) (*types.Transaction, error) {
	return _ExchangeData.Contract.CreateOrders(&_ExchangeData.TransactOpts, orderId, userAddress, userReceiver, tradeType, sellAmount, sellContract, price, buyContract, market_code)
}

// DeleteToken is a paid mutator transaction binding the contract method 0xdedb6e40.
//
// Solidity: function deleteToken(address token_address, uint64 market_code) returns()
func (_ExchangeData *ExchangeDataTransactor) DeleteToken(opts *bind.TransactOpts, token_address common.Address, market_code uint64) (*types.Transaction, error) {
	return _ExchangeData.contract.Transact(opts, "deleteToken", token_address, market_code)
}

// DeleteToken is a paid mutator transaction binding the contract method 0xdedb6e40.
//
// Solidity: function deleteToken(address token_address, uint64 market_code) returns()
func (_ExchangeData *ExchangeDataSession) DeleteToken(token_address common.Address, market_code uint64) (*types.Transaction, error) {
	return _ExchangeData.Contract.DeleteToken(&_ExchangeData.TransactOpts, token_address, market_code)
}

// DeleteToken is a paid mutator transaction binding the contract method 0xdedb6e40.
//
// Solidity: function deleteToken(address token_address, uint64 market_code) returns()
func (_ExchangeData *ExchangeDataTransactorSession) DeleteToken(token_address common.Address, market_code uint64) (*types.Transaction, error) {
	return _ExchangeData.Contract.DeleteToken(&_ExchangeData.TransactOpts, token_address, market_code)
}

// GetTransferInfo is a paid mutator transaction binding the contract method 0x5a7705b7.
//
// Solidity: function getTransferInfo(uint64 orderId) returns(uint256 sellAmount, address to, address tokenContract)
func (_ExchangeData *ExchangeDataTransactor) GetTransferInfo(opts *bind.TransactOpts, orderId uint64) (*types.Transaction, error) {
	return _ExchangeData.contract.Transact(opts, "getTransferInfo", orderId)
}

// GetTransferInfo is a paid mutator transaction binding the contract method 0x5a7705b7.
//
// Solidity: function getTransferInfo(uint64 orderId) returns(uint256 sellAmount, address to, address tokenContract)
func (_ExchangeData *ExchangeDataSession) GetTransferInfo(orderId uint64) (*types.Transaction, error) {
	return _ExchangeData.Contract.GetTransferInfo(&_ExchangeData.TransactOpts, orderId)
}

// GetTransferInfo is a paid mutator transaction binding the contract method 0x5a7705b7.
//
// Solidity: function getTransferInfo(uint64 orderId) returns(uint256 sellAmount, address to, address tokenContract)
func (_ExchangeData *ExchangeDataTransactorSession) GetTransferInfo(orderId uint64) (*types.Transaction, error) {
	return _ExchangeData.Contract.GetTransferInfo(&_ExchangeData.TransactOpts, orderId)
}

// RollbackSubmit is a paid mutator transaction binding the contract method 0x1638550b.
//
// Solidity: function rollbackSubmit(uint64 orderId) returns()
func (_ExchangeData *ExchangeDataTransactor) RollbackSubmit(opts *bind.TransactOpts, orderId uint64) (*types.Transaction, error) {
	return _ExchangeData.contract.Transact(opts, "rollbackSubmit", orderId)
}

// RollbackSubmit is a paid mutator transaction binding the contract method 0x1638550b.
//
// Solidity: function rollbackSubmit(uint64 orderId) returns()
func (_ExchangeData *ExchangeDataSession) RollbackSubmit(orderId uint64) (*types.Transaction, error) {
	return _ExchangeData.Contract.RollbackSubmit(&_ExchangeData.TransactOpts, orderId)
}

// RollbackSubmit is a paid mutator transaction binding the contract method 0x1638550b.
//
// Solidity: function rollbackSubmit(uint64 orderId) returns()
func (_ExchangeData *ExchangeDataTransactorSession) RollbackSubmit(orderId uint64) (*types.Transaction, error) {
	return _ExchangeData.Contract.RollbackSubmit(&_ExchangeData.TransactOpts, orderId)
}

// SetAccessAllow is a paid mutator transaction binding the contract method 0x9bab223e.
//
// Solidity: function setAccessAllow(address owner, bool access) returns()
func (_ExchangeData *ExchangeDataTransactor) SetAccessAllow(opts *bind.TransactOpts, owner common.Address, access bool) (*types.Transaction, error) {
	return _ExchangeData.contract.Transact(opts, "setAccessAllow", owner, access)
}

// SetAccessAllow is a paid mutator transaction binding the contract method 0x9bab223e.
//
// Solidity: function setAccessAllow(address owner, bool access) returns()
func (_ExchangeData *ExchangeDataSession) SetAccessAllow(owner common.Address, access bool) (*types.Transaction, error) {
	return _ExchangeData.Contract.SetAccessAllow(&_ExchangeData.TransactOpts, owner, access)
}

// SetAccessAllow is a paid mutator transaction binding the contract method 0x9bab223e.
//
// Solidity: function setAccessAllow(address owner, bool access) returns()
func (_ExchangeData *ExchangeDataTransactorSession) SetAccessAllow(owner common.Address, access bool) (*types.Transaction, error) {
	return _ExchangeData.Contract.SetAccessAllow(&_ExchangeData.TransactOpts, owner, access)
}

// SetContractOwner is a paid mutator transaction binding the contract method 0xa34d42b8.
//
// Solidity: function setContractOwner(address newAddr) returns()
func (_ExchangeData *ExchangeDataTransactor) SetContractOwner(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _ExchangeData.contract.Transact(opts, "setContractOwner", newAddr)
}

// SetContractOwner is a paid mutator transaction binding the contract method 0xa34d42b8.
//
// Solidity: function setContractOwner(address newAddr) returns()
func (_ExchangeData *ExchangeDataSession) SetContractOwner(newAddr common.Address) (*types.Transaction, error) {
	return _ExchangeData.Contract.SetContractOwner(&_ExchangeData.TransactOpts, newAddr)
}

// SetContractOwner is a paid mutator transaction binding the contract method 0xa34d42b8.
//
// Solidity: function setContractOwner(address newAddr) returns()
func (_ExchangeData *ExchangeDataTransactorSession) SetContractOwner(newAddr common.Address) (*types.Transaction, error) {
	return _ExchangeData.Contract.SetContractOwner(&_ExchangeData.TransactOpts, newAddr)
}

// SetOrder is a paid mutator transaction binding the contract method 0x5492c9f9.
//
// Solidity: function setOrder(uint64 orderId, uint256 sellAmount, uint256 buyAmount, uint256 fees, bool withdrawAble, bool exSuccess) returns()
func (_ExchangeData *ExchangeDataTransactor) SetOrder(opts *bind.TransactOpts, orderId uint64, sellAmount *big.Int, buyAmount *big.Int, fees *big.Int, withdrawAble bool, exSuccess bool) (*types.Transaction, error) {
	return _ExchangeData.contract.Transact(opts, "setOrder", orderId, sellAmount, buyAmount, fees, withdrawAble, exSuccess)
}

// SetOrder is a paid mutator transaction binding the contract method 0x5492c9f9.
//
// Solidity: function setOrder(uint64 orderId, uint256 sellAmount, uint256 buyAmount, uint256 fees, bool withdrawAble, bool exSuccess) returns()
func (_ExchangeData *ExchangeDataSession) SetOrder(orderId uint64, sellAmount *big.Int, buyAmount *big.Int, fees *big.Int, withdrawAble bool, exSuccess bool) (*types.Transaction, error) {
	return _ExchangeData.Contract.SetOrder(&_ExchangeData.TransactOpts, orderId, sellAmount, buyAmount, fees, withdrawAble, exSuccess)
}

// SetOrder is a paid mutator transaction binding the contract method 0x5492c9f9.
//
// Solidity: function setOrder(uint64 orderId, uint256 sellAmount, uint256 buyAmount, uint256 fees, bool withdrawAble, bool exSuccess) returns()
func (_ExchangeData *ExchangeDataTransactorSession) SetOrder(orderId uint64, sellAmount *big.Int, buyAmount *big.Int, fees *big.Int, withdrawAble bool, exSuccess bool) (*types.Transaction, error) {
	return _ExchangeData.Contract.SetOrder(&_ExchangeData.TransactOpts, orderId, sellAmount, buyAmount, fees, withdrawAble, exSuccess)
}

// WithdrawOrder is a paid mutator transaction binding the contract method 0x1712046f.
//
// Solidity: function withdrawOrder(uint64 orderId) returns(uint256 amount, address userAddress, address tokenContract)
func (_ExchangeData *ExchangeDataTransactor) WithdrawOrder(opts *bind.TransactOpts, orderId uint64) (*types.Transaction, error) {
	return _ExchangeData.contract.Transact(opts, "withdrawOrder", orderId)
}

// WithdrawOrder is a paid mutator transaction binding the contract method 0x1712046f.
//
// Solidity: function withdrawOrder(uint64 orderId) returns(uint256 amount, address userAddress, address tokenContract)
func (_ExchangeData *ExchangeDataSession) WithdrawOrder(orderId uint64) (*types.Transaction, error) {
	return _ExchangeData.Contract.WithdrawOrder(&_ExchangeData.TransactOpts, orderId)
}

// WithdrawOrder is a paid mutator transaction binding the contract method 0x1712046f.
//
// Solidity: function withdrawOrder(uint64 orderId) returns(uint256 amount, address userAddress, address tokenContract)
func (_ExchangeData *ExchangeDataTransactorSession) WithdrawOrder(orderId uint64) (*types.Transaction, error) {
	return _ExchangeData.Contract.WithdrawOrder(&_ExchangeData.TransactOpts, orderId)
}
