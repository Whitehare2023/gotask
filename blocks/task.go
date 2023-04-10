// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blocks

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TaskInfo is an auto generated low-level Go binding around an user-defined struct.
type TaskInfo struct {
	Issuer    string
	Worker    string
	Desc      string
	Bonus     *big.Int
	Status    uint8
	Timestamp *big.Int
	Comment   string
}

// IERC2000MetaData contains all meta data concerning the IERC2000 contract.
var IERC2000MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_owner\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"35ee5f87": "balanceOf(string)",
		"18160ddd": "totalSupply()",
		"9b80b050": "transfer(string,string,uint256)",
	},
}

// IERC2000ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC2000MetaData.ABI instead.
var IERC2000ABI = IERC2000MetaData.ABI

// Deprecated: Use IERC2000MetaData.Sigs instead.
// IERC2000FuncSigs maps the 4-byte function signature to its string representation.
var IERC2000FuncSigs = IERC2000MetaData.Sigs

// IERC2000 is an auto generated Go binding around an Ethereum contract.
type IERC2000 struct {
	IERC2000Caller     // Read-only binding to the contract
	IERC2000Transactor // Write-only binding to the contract
	IERC2000Filterer   // Log filterer for contract events
}

// IERC2000Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC2000Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC2000Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC2000Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC2000Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC2000Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC2000Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC2000Session struct {
	Contract     *IERC2000         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC2000CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC2000CallerSession struct {
	Contract *IERC2000Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IERC2000TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC2000TransactorSession struct {
	Contract     *IERC2000Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC2000Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC2000Raw struct {
	Contract *IERC2000 // Generic contract binding to access the raw methods on
}

// IERC2000CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC2000CallerRaw struct {
	Contract *IERC2000Caller // Generic read-only contract binding to access the raw methods on
}

// IERC2000TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC2000TransactorRaw struct {
	Contract *IERC2000Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC2000 creates a new instance of IERC2000, bound to a specific deployed contract.
func NewIERC2000(address common.Address, backend bind.ContractBackend) (*IERC2000, error) {
	contract, err := bindIERC2000(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC2000{IERC2000Caller: IERC2000Caller{contract: contract}, IERC2000Transactor: IERC2000Transactor{contract: contract}, IERC2000Filterer: IERC2000Filterer{contract: contract}}, nil
}

// NewIERC2000Caller creates a new read-only instance of IERC2000, bound to a specific deployed contract.
func NewIERC2000Caller(address common.Address, caller bind.ContractCaller) (*IERC2000Caller, error) {
	contract, err := bindIERC2000(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC2000Caller{contract: contract}, nil
}

// NewIERC2000Transactor creates a new write-only instance of IERC2000, bound to a specific deployed contract.
func NewIERC2000Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC2000Transactor, error) {
	contract, err := bindIERC2000(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC2000Transactor{contract: contract}, nil
}

// NewIERC2000Filterer creates a new log filterer instance of IERC2000, bound to a specific deployed contract.
func NewIERC2000Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC2000Filterer, error) {
	contract, err := bindIERC2000(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC2000Filterer{contract: contract}, nil
}

// bindIERC2000 binds a generic wrapper to an already deployed contract.
func bindIERC2000(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC2000ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC2000 *IERC2000Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC2000.Contract.IERC2000Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC2000 *IERC2000Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC2000.Contract.IERC2000Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC2000 *IERC2000Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC2000.Contract.IERC2000Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC2000 *IERC2000CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC2000.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC2000 *IERC2000TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC2000.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC2000 *IERC2000TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC2000.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) view returns(uint256 balance)
func (_IERC2000 *IERC2000Caller) BalanceOf(opts *bind.CallOpts, _owner string) (*big.Int, error) {
	var out []interface{}
	err := _IERC2000.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) view returns(uint256 balance)
func (_IERC2000 *IERC2000Session) BalanceOf(_owner string) (*big.Int, error) {
	return _IERC2000.Contract.BalanceOf(&_IERC2000.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) view returns(uint256 balance)
func (_IERC2000 *IERC2000CallerSession) BalanceOf(_owner string) (*big.Int, error) {
	return _IERC2000.Contract.BalanceOf(&_IERC2000.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC2000 *IERC2000Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC2000.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC2000 *IERC2000Session) TotalSupply() (*big.Int, error) {
	return _IERC2000.Contract.TotalSupply(&_IERC2000.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC2000 *IERC2000CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC2000.Contract.TotalSupply(&_IERC2000.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_IERC2000 *IERC2000Transactor) Transfer(opts *bind.TransactOpts, _from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _IERC2000.contract.Transact(opts, "transfer", _from, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_IERC2000 *IERC2000Session) Transfer(_from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _IERC2000.Contract.Transfer(&_IERC2000.TransactOpts, _from, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_IERC2000 *IERC2000TransactorSession) Transfer(_from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _IERC2000.Contract.Transfer(&_IERC2000.TransactOpts, _from, _to, _value)
}

// IUSERMetaData contains all meta data concerning the IUSER contract.
var IUSERMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"58467dbc": "login(string,string)",
		"3ffbd47f": "register(string,string)",
	},
}

// IUSERABI is the input ABI used to generate the binding from.
// Deprecated: Use IUSERMetaData.ABI instead.
var IUSERABI = IUSERMetaData.ABI

// Deprecated: Use IUSERMetaData.Sigs instead.
// IUSERFuncSigs maps the 4-byte function signature to its string representation.
var IUSERFuncSigs = IUSERMetaData.Sigs

// IUSER is an auto generated Go binding around an Ethereum contract.
type IUSER struct {
	IUSERCaller     // Read-only binding to the contract
	IUSERTransactor // Write-only binding to the contract
	IUSERFilterer   // Log filterer for contract events
}

// IUSERCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUSERCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUSERTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUSERTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUSERFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUSERFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUSERSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUSERSession struct {
	Contract     *IUSER            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUSERCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUSERCallerSession struct {
	Contract *IUSERCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IUSERTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUSERTransactorSession struct {
	Contract     *IUSERTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUSERRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUSERRaw struct {
	Contract *IUSER // Generic contract binding to access the raw methods on
}

// IUSERCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUSERCallerRaw struct {
	Contract *IUSERCaller // Generic read-only contract binding to access the raw methods on
}

// IUSERTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUSERTransactorRaw struct {
	Contract *IUSERTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUSER creates a new instance of IUSER, bound to a specific deployed contract.
func NewIUSER(address common.Address, backend bind.ContractBackend) (*IUSER, error) {
	contract, err := bindIUSER(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUSER{IUSERCaller: IUSERCaller{contract: contract}, IUSERTransactor: IUSERTransactor{contract: contract}, IUSERFilterer: IUSERFilterer{contract: contract}}, nil
}

// NewIUSERCaller creates a new read-only instance of IUSER, bound to a specific deployed contract.
func NewIUSERCaller(address common.Address, caller bind.ContractCaller) (*IUSERCaller, error) {
	contract, err := bindIUSER(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUSERCaller{contract: contract}, nil
}

// NewIUSERTransactor creates a new write-only instance of IUSER, bound to a specific deployed contract.
func NewIUSERTransactor(address common.Address, transactor bind.ContractTransactor) (*IUSERTransactor, error) {
	contract, err := bindIUSER(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUSERTransactor{contract: contract}, nil
}

// NewIUSERFilterer creates a new log filterer instance of IUSER, bound to a specific deployed contract.
func NewIUSERFilterer(address common.Address, filterer bind.ContractFilterer) (*IUSERFilterer, error) {
	contract, err := bindIUSER(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUSERFilterer{contract: contract}, nil
}

// bindIUSER binds a generic wrapper to an already deployed contract.
func bindIUSER(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUSERABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUSER *IUSERRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUSER.Contract.IUSERCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUSER *IUSERRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUSER.Contract.IUSERTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUSER *IUSERRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUSER.Contract.IUSERTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUSER *IUSERCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUSER.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUSER *IUSERTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUSER.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUSER *IUSERTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUSER.Contract.contract.Transact(opts, method, params...)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) view returns(bool)
func (_IUSER *IUSERCaller) Login(opts *bind.CallOpts, username string, passwd string) (bool, error) {
	var out []interface{}
	err := _IUSER.contract.Call(opts, &out, "login", username, passwd)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) view returns(bool)
func (_IUSER *IUSERSession) Login(username string, passwd string) (bool, error) {
	return _IUSER.Contract.Login(&_IUSER.CallOpts, username, passwd)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) view returns(bool)
func (_IUSER *IUSERCallerSession) Login(username string, passwd string) (bool, error) {
	return _IUSER.Contract.Login(&_IUSER.CallOpts, username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_IUSER *IUSERTransactor) Register(opts *bind.TransactOpts, username string, passwd string) (*types.Transaction, error) {
	return _IUSER.contract.Transact(opts, "register", username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_IUSER *IUSERSession) Register(username string, passwd string) (*types.Transaction, error) {
	return _IUSER.Contract.Register(&_IUSER.TransactOpts, username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_IUSER *IUSERTransactorSession) Register(username string, passwd string) (*types.Transaction, error) {
	return _IUSER.Contract.Register(&_IUSER.TransactOpts, username, passwd)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209fe799cfbe95e0c3a8f368dc60c2a7f2c0892d239f4f1d1558bbfd36bc1e0df164736f6c634300060a0033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TaskMetaData contains all meta data concerning the Task contract.
var TaskMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_owner\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"worker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"confirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qryAllTasks\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"worker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"internalType\":\"structTaskInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"qryOneTask\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"worker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"take\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"35ee5f87": "balanceOf(string)",
		"89dc99d2": "commit(string,string,uint256)",
		"fc18cb62": "confirm(string,string,uint256,string,uint8)",
		"679a5e14": "issue(string,string,string,uint256)",
		"58467dbc": "login(string,string)",
		"056b01ce": "mint(string,uint256)",
		"8a3f7bb3": "qryAllTasks()",
		"9b3add75": "qryOneTask(uint256)",
		"3ffbd47f": "register(string,string)",
		"5b4c0b54": "take(string,string,uint256)",
		"18160ddd": "totalSupply()",
		"9b80b050": "transfer(string,string,uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b50604051620031b5380380620031b5833981016040819052620000349162000118565b60028054336001600160a01b031991821617909155600380549091166001600160a01b0383161790556040516200006b90620000fc565b604051809103906000f08015801562000088573d6000803e3d6000fd5b50600080546001600160a01b0319166001600160a01b0392909216919091179055604051620000b7906200010a565b604051809103906000f080158015620000d4573d6000803e3d6000fd5b50600180546001600160a01b0319166001600160a01b03929092169190911790555062000148565b6108c980620020d483390190565b610818806200299d83390190565b6000602082840312156200012a578081fd5b81516001600160a01b038116811462000141578182fd5b9392505050565b611f7c80620001586000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063679a5e1411610071578063679a5e141461014557806389dc99d2146101585780638a3f7bb31461016b5780639b3add75146101805780639b80b050146101a4578063fc18cb62146101b7576100b4565b8063056b01ce146100b957806318160ddd146100ce57806335ee5f87146100ec5780633ffbd47f146100ff57806358467dbc146101125780635b4c0b5414610132575b600080fd5b6100cc6100c7366004611949565b6101ca565b005b6100d661027d565b6040516100e39190611f31565b60405180910390f35b6100d66100fa3660046116d8565b61030a565b6100cc61010d366004611718565b610394565b610125610120366004611718565b610402565b6040516100e39190611b25565b6100cc610140366004611820565b610492565b6100cc610153366004611781565b610620565b6100cc610166366004611820565b610915565b610173610b62565b6040516100e39190611a3f565b61019361018e366004611993565b610e40565b6040516100e3959493929190611c9c565b6101256101b2366004611820565b6110a4565b6100cc6101c5366004611891565b611139565b6002546001600160a01b03163314806101ed57506003546001600160a01b031633145b6102125760405162461bcd60e51b815260040161020990611e9b565b60405180910390fd5b6000546040516302b580e760e11b81526001600160a01b039091169063056b01ce9061024690869086908690600401611c65565b600060405180830381600087803b15801561026057600080fd5b505af1158015610274573d6000803e3d6000fd5b50505050505050565b60008060009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156102cc57600080fd5b505afa1580156102e0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030491906119ab565b90505b90565b600080546040516335ee5f8760e01b81526001600160a01b03909116906335ee5f879061033d9086908690600401611b30565b60206040518083038186803b15801561035557600080fd5b505afa158015610369573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061038d91906119ab565b9392505050565b600154604051633ffbd47f60e01b81526001600160a01b0390911690633ffbd47f906103ca908790879087908790600401611b4c565b600060405180830381600087803b1580156103e457600080fd5b505af11580156103f8573d6000803e3d6000fd5b5050505050505050565b6001546040516316119f6f60e21b81526000916001600160a01b0316906358467dbc90610439908890889088908890600401611b4c565b60206040518083038186803b15801561045157600080fd5b505afa158015610465573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061048991906116b8565b95945050505050565b6001546040516316119f6f60e21b81526001600160a01b03909116906358467dbc906104c8908890889088908890600401611b4c565b60206040518083038186803b1580156104e057600080fd5b505afa1580156104f4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061051891906116b8565b6105345760405162461bcd60e51b815260040161020990611df6565b60006004828154811061054357fe5b906000526020600020906007020160050154116105725760405162461bcd60e51b815260040161020990611d5e565b6004818154811061057f57fe5b600091825260209091206004600790920201015460ff16156105b35760405162461bcd60e51b815260040161020990611e64565b8484600483815481106105c257fe5b906000526020600020906007020160010191906105e092919061152b565b506001600482815481106105f057fe5b906000526020600020906007020160040160006101000a81548160ff021916908360ff1602179055505050505050565b6001546040516316119f6f60e21b81526001600160a01b03909116906358467dbc90610656908a908a908a908a90600401611b4c565b60206040518083038186803b15801561066e57600080fd5b505afa158015610682573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a691906116b8565b6106c25760405162461bcd60e51b815260040161020990611d88565b600081116106e25760405162461bcd60e51b815260040161020990611ed2565b6000546040516335ee5f8760e01b815282916001600160a01b0316906335ee5f8790610714908b908b90600401611b30565b60206040518083038186803b15801561072c57600080fd5b505afa158015610740573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076491906119ab565b10156107825760405162461bcd60e51b815260040161020990611efa565b61078a6115a9565b6040518060e0016040528089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093855250506040805160208181018352938152838501528051601f8901849004840281018401825288815293019291889150879081908401838280828437600092018290525093855250505060208083018690526040808401839052426060850152805180830190915282815260809093019290925260048054600181018255915282518051939450849360079092027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0192610889928492909101906115e9565b5060208281015180516108a292600185019201906115e9565b50604082015180516108be9160028401916020909101906115e9565b5060608201516003820155608082015160048201805460ff191660ff90921691909117905560a0820151600582015560c082015180516109089160068401916020909101906115e9565b5050505050505050505050565b6001546040516316119f6f60e21b81526001600160a01b03909116906358467dbc9061094b908890889088908890600401611b4c565b60206040518083038186803b15801561096357600080fd5b505afa158015610977573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099b91906116b8565b6109b75760405162461bcd60e51b815260040161020990611df6565b6000600482815481106109c657fe5b906000526020600020906007020160050154116109f55760405162461bcd60e51b815260040161020990611d5e565b60048181548110610a0257fe5b600091825260209091206004600790920201015460ff16600114610a385760405162461bcd60e51b815260040161020990611cf0565b610b3760048281548110610a4857fe5b90600052602060002090600702016001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610aed5780601f10610ac257610100808354040283529160200191610aed565b820191906000526020600020905b815481529060010190602001808311610ad057829003601f168201915b505050505086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6114c8169050565b610b535760405162461bcd60e51b815260040161020990611d27565b6002600482815481106105f057fe5b60606004805480602002602001604051908101604052809291908181526020016000905b82821015610e3757600084815260209081902060408051600786029092018054600260018216156101009081026000190190921604601f81018690049095028401810190925260e0830184815292939092849290918491840182828015610c2e5780601f10610c0357610100808354040283529160200191610c2e565b820191906000526020600020905b815481529060010190602001808311610c1157829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cd05780601f10610ca557610100808354040283529160200191610cd0565b820191906000526020600020905b815481529060010190602001808311610cb357829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f81018390048302850183019091528084529381019390830182828015610d625780601f10610d3757610100808354040283529160200191610d62565b820191906000526020600020905b815481529060010190602001808311610d4557829003601f168201915b50505091835250506003820154602080830191909152600483015460ff1660408084019190915260058401546060840152600684018054825160026101006001841615026000190190921691909104601f810185900485028201850190935282815260809094019392909190830182828015610e1f5780601f10610df457610100808354040283529160200191610e1f565b820191906000526020600020905b815481529060010190602001808311610e0257829003601f168201915b50505050508152505081526020019060010190610b86565b50505050905090565b60608060006060600060048681548110610e5657fe5b906000526020600020906007020160000160048781548110610e7457fe5b906000526020600020906007020160020160048881548110610e9257fe5b90600052602060002090600702016003015460048981548110610eb157fe5b906000526020600020906007020160010160048a81548110610ecf57fe5b600091825260209182902060046007909202010154855460408051601f6002600019600186161561010002019094169390930492830185900485028101850190915281815260ff90921692879190830182828015610f6e5780601f10610f4357610100808354040283529160200191610f6e565b820191906000526020600020905b815481529060010190602001808311610f5157829003601f168201915b5050875460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959a5089945092508401905082828015610ffc5780601f10610fd157610100808354040283529160200191610ffc565b820191906000526020600020905b815481529060010190602001808311610fdf57829003601f168201915b5050855460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529599508794509250840190508282801561108a5780601f1061105f5761010080835404028352916020019161108a565b820191906000526020600020905b81548152906001019060200180831161106d57829003601f168201915b505050505091509450945094509450945091939590929450565b600080546040516309b80b0560e41b81526001600160a01b0390911690639b80b050906110dd9089908990899089908990600401611b7e565b602060405180830381600087803b1580156110f757600080fd5b505af115801561110b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061112f91906116b8565b5095945050505050565b6001546040516316119f6f60e21b81526001600160a01b03909116906358467dbc9061116f908b908b908b908b90600401611b4c565b60206040518083038186803b15801561118757600080fd5b505afa15801561119b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111bf91906116b8565b6111db5760405162461bcd60e51b815260040161020990611df6565b6000600485815481106111ea57fe5b906000526020600020906007020160050154116112195760405162461bcd60e51b815260040161020990611d5e565b6004848154811061122657fe5b600091825260209091206007909102016004015460ff1660021461125c5760405162461bcd60e51b815260040161020990611dbf565b6113496004858154811061126c57fe5b6000918252602091829020600790910201805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156112ff5780601f106112d4576101008083540402835291602001916112ff565b820191906000526020600020905b8154815290600101906020018083116112e257829003601f168201915b505050505089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6114c8169050565b6113655760405162461bcd60e51b815260040161020990611e2d565b8060ff166003148061137a57508060ff166001145b61138357600080fd5b806004858154811061139157fe5b906000526020600020906007020160040160006101000a81548160ff021916908360ff1602179055508282600486815481106113c957fe5b906000526020600020906007020160060191906113e792919061152b565b508060ff16600314156103f857600054600480546001600160a01b0390921691639b80b050918b918b91908990811061141c57fe5b90600052602060002090600702016001016004898154811061143a57fe5b9060005260206000209060070201600301546040518563ffffffff1660e01b815260040161146b9493929190611bb8565b602060405180830381600087803b15801561148557600080fd5b505af1158015611499573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114bd91906116b8565b505050505050505050565b600080836040516020016114dc9190611c89565b6040516020818303038152906040528051906020012090506000836040516020016115079190611c89565b60408051601f19818403018152919052805160209091012091909114949350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061156c5782800160ff19823516178555611599565b82800160010185558215611599579182015b8281111561159957823582559160200191906001019061157e565b506115a5929150611657565b5090565b6040518060e0016040528060608152602001606081526020016060815260200160008152602001600060ff16815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061162a57805160ff1916838001178555611599565b82800160010185558215611599579182015b8281111561159957825182559160200191906001019061163c565b61030791905b808211156115a5576000815560010161165d565b60008083601f840112611682578182fd5b50813567ffffffffffffffff811115611699578182fd5b6020830191508360208285010111156116b157600080fd5b9250929050565b6000602082840312156116c9578081fd5b8151801515811461038d578182fd5b600080602083850312156116ea578081fd5b823567ffffffffffffffff811115611700578182fd5b61170c85828601611671565b90969095509350505050565b6000806000806040858703121561172d578182fd5b843567ffffffffffffffff80821115611744578384fd5b61175088838901611671565b90965094506020870135915080821115611768578384fd5b5061177587828801611671565b95989497509550505050565b60008060008060008060006080888a03121561179b578283fd5b873567ffffffffffffffff808211156117b2578485fd5b6117be8b838c01611671565b909950975060208a01359150808211156117d6578485fd5b6117e28b838c01611671565b909750955060408a01359150808211156117fa578485fd5b506118078a828b01611671565b989b979a50959894979596606090950135949350505050565b600080600080600060608688031215611837578081fd5b853567ffffffffffffffff8082111561184e578283fd5b61185a89838a01611671565b90975095506020880135915080821115611872578283fd5b5061187f88828901611671565b96999598509660400135949350505050565b60008060008060008060008060a0898b0312156118ac578081fd5b883567ffffffffffffffff808211156118c3578283fd5b6118cf8c838d01611671565b909a50985060208b01359150808211156118e7578283fd5b6118f38c838d01611671565b909850965060408b0135955060608b0135915080821115611912578283fd5b5061191f8b828c01611671565b909450925050608089013560ff81168114611938578182fd5b809150509295985092959890939650565b60008060006040848603121561195d578283fd5b833567ffffffffffffffff811115611973578384fd5b61197f86828701611671565b909790965060209590950135949350505050565b6000602082840312156119a4578081fd5b5035919050565b6000602082840312156119bc578081fd5b5051919050565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b60008151808452815b81811015611a12576020818501810151868301820152016119f6565b81811115611a235782602083870101525b50601f01601f19169290920160200192915050565b60ff169052565b60208082528251828201819052600091906040908185019080840286018301878501865b83811015611b1757603f19898403018552815160e08151818652611a89828701826119ed565b8a84015192508681038b880152611aa081846119ed565b91505088830151915085810389870152611aba81836119ed565b606084810151908801526080808501519093509150611adb83880183611a38565b60a0848101519088015260c080850151888303828a01529093509150611b0181836119ed565b988b019896505050928801925050600101611a63565b509098975050505050505050565b901515815260200190565b600060208252611b446020830184866119c3565b949350505050565b600060408252611b606040830186886119c3565b8281036020840152611b738185876119c3565b979650505050505050565b600060608252611b926060830187896119c3565b8281036020840152611ba58186886119c3565b9150508260408301529695505050505050565b600060608252611bcc6060830186886119c3565b602083820381850152828654600180821660008114611bf25760018114611c1257611c4d565b611c02607f600285041687611f31565b60ff198416815285019350611c4d565b60028304611c208188611f31565b611c298c611f3a565b895b83811015611c4457815483820152908501908801611c2b565b91909101955050505b50505080935050505082604083015295945050505050565b600060408252611c796040830185876119c3565b9050826020830152949350505050565b60006020825261038d60208301846119ed565b600060a08252611caf60a08301886119ed565b8281036020840152611cc181886119ed565b8660408501528381036060850152611cd981876119ed565b9250505060ff831660808301529695505050505050565b60208082526017908201527f7461736b277320737461747573206d757374203d3d2031000000000000000000604082015260600190565b6020808252601b908201527f7461736b277320776f726b73206d75737420697320776f726b65720000000000604082015260600190565b60208082526010908201526f7461736b206d7573742065786973747360801b604082015260600190565b6020808252601a908201527f697373756572206d757374206861766520686973207269676874000000000000604082015260600190565b60208082526017908201527f7461736b277320737461747573206d757374203d3d2032000000000000000000604082015260600190565b6020808252601a908201527f776f726b6572206d757374206861766520686973207269676874000000000000604082015260600190565b6020808252601c908201527f7461736b277320697373756572206d7573742069732069737375657200000000604082015260600190565b60208082526017908201527f7461736b277320737461747573206d757374203d3d2030000000000000000000604082015260600190565b60208082526017908201527f4f6e6c792061646d696e2063616e20646f207468697321000000000000000000604082015260600190565b6020808252600e908201526d0626f6e7573206d757374203e20360941b604082015260600190565b6020808252601a908201527f697373756572277320626f6e7573206d75737420656e6f756768000000000000604082015260600190565b90815260200190565b6000908152602090209056fea2646970667358221220ebe892d0a078026fa1fcf9a40b8aa78095cfaa5370e175366b989ee1a93f349964736f6c634300060a0033608060405234801561001057600080fd5b5060008055600180546001600160a01b03191633179055610893806100366000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063056b01ce1461005157806318160ddd146100c157806335ee5f87146100db5780639b80b05014610149575b600080fd5b6100bf6004803603604081101561006757600080fd5b810190602081018135600160201b81111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460018302840111600160201b831117156100b457600080fd5b91935091503561021b565b005b6100c9610444565b60408051918252519081900360200190f35b6100c9600480360360208110156100f157600080fd5b810190602081018135600160201b81111561010b57600080fd5b82018360208201111561011d57600080fd5b803590602001918460018302840111600160201b8311171561013e57600080fd5b50909250905061044a565b6102076004803603606081101561015f57600080fd5b810190602081018135600160201b81111561017957600080fd5b82018360208201111561018b57600080fd5b803590602001918460018302840111600160201b831117156101ac57600080fd5b919390929091602081019035600160201b8111156101c957600080fd5b8201836020820111156101db57600080fd5b803590602001918460018302840111600160201b831117156101fc57600080fd5b91935091503561047b565b604080519115158252519081900360200190f35b6001546001600160a01b0316331461027a576040805162461bcd60e51b815260206004820152601760248201527f4f6e6c792061646d696e2063616e20646f207468697321000000000000000000604482015290519081900360640190fd5b6102d26040518060200160405280600081525084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff610709169050565b15610316576040805162461bcd60e51b815260206004820152600f60248201526e2a379036bab9ba1032bc34b9ba399760891b604482015290519081900360640190fd5b6000811161035d576040805162461bcd60e51b815260206004820152600f60248201526e2b30b63ab29036bab9ba101f10181760891b604482015290519081900360640190fd5b600054610370908263ffffffff61082f16565b6000819055506103ae81600285856040518083838082843791909101948552505060405192839003602001909220549291505063ffffffff61082f16565b60028484604051808383808284379190910194855250506040519283900360200183209390935550849150839080838380828437604080519390910183900383206573797374656d60d01b8452815193849003600601842088855291519096509094507f5358be4df107be4d9b023fc323f41d7109610225c6ef211b9d375b9fbd7ccc4f935091829003602001919050a3505050565b60005490565b6000600283836040518083838082843791909101948552505060405192839003602001909220549250505092915050565b60006104d56040518060200160405280600081525085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff610709169050565b15610519576040805162461bcd60e51b815260206004820152600f60248201526e2a379036bab9ba1032bc34b9ba399760891b604482015290519081900360640190fd5b60008211610560576040805162461bcd60e51b815260206004820152600f60248201526e2b30b63ab29036bab9ba101f10181760891b604482015290519081900360640190fd5b816002878760405180838380828437808301925050509250505090815260200160405180910390205410156105d3576040805162461bcd60e51b81526020600482015260146024820152732130b630b731b29036bab9ba1032b737bab3b41760611b604482015290519081900360640190fd5b61060b82600288886040518083838082843791909101948552505060405192839003602001909220549291505063ffffffff61084816565b6002878760405180838380828437808301925050509250505090815260200160405180910390208190555061066e82600286866040518083838082843791909101948552505060405192839003602001909220549291505063ffffffff61082f16565b600285856040518083838082843791909101948552505060405192839003602001832093909355508591508490808383808284376040519201829003822094508a93508992508190508383808284376040805193909101839003832089845290519095507f5358be4df107be4d9b023fc323f41d7109610225c6ef211b9d375b9fbd7ccc4f945091829003602001925050a395945050505050565b600080836040516020018080602001828103825283818151815260200191508051906020019080838360005b8381101561074d578181015183820152602001610735565b50505050905090810190601f16801561077a5780820380516001836020036101000a031916815260200191505b50925050506040516020818303038152906040528051906020012090506000836040516020018080602001828103825283818151815260200191508051906020019080838360005b838110156107da5781810151838201526020016107c2565b50505050905090810190601f1680156108075780820380516001836020036101000a031916815260200191505b5060408051808303601f19018152919052805160209091012094909414979650505050505050565b60008282018381101561084157600080fd5b9392505050565b60008282111561085757600080fd5b5090039056fea26469706673582212204e9eb7f1376d7d4a8ca22f130ce09c15a217d3467c5a49bdb983e67f80cef77964736f6c634300060a0033608060405234801561001057600080fd5b50600180546001600160a01b031916331790556107e6806100326000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633ffbd47f1461003b57806358467dbc146100fb575b600080fd5b6100f96004803603604081101561005157600080fd5b810190602081018135600160201b81111561006b57600080fd5b82018360208201111561007d57600080fd5b803590602001918460018302840111600160201b8311171561009e57600080fd5b919390929091602081019035600160201b8111156100bb57600080fd5b8201836020820111156100cd57600080fd5b803590602001918460018302840111600160201b831117156100ee57600080fd5b5090925090506101cd565b005b6101b96004803603604081101561011157600080fd5b810190602081018135600160201b81111561012b57600080fd5b82018360208201111561013d57600080fd5b803590602001918460018302840111600160201b8311171561015e57600080fd5b919390929091602081019035600160201b81111561017b57600080fd5b82018360208201111561018d57600080fd5b803590602001918460018302840111600160201b831117156101ae57600080fd5b50909250905061046a565b604080519115158252519081900360200190f35b6102256040518060200160405280600081525085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6105ef169050565b15610277576040805162461bcd60e51b815260206004820152601760248201527f557365726e616d65206d757374206e6f74206e756c6c2e000000000000000000604482015290519081900360640190fd5b6102cf6040518060200160405280600081525083838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6105ef169050565b15610319576040805162461bcd60e51b81526020600482015260156024820152742830b9b9bbb21036bab9ba103737ba10373ab6361760591b604482015290519081900360640190fd5b6103e56040518060200160405280600081525060008686604051808383808284379190910194855250506040805160209481900385018120805460026001821615610100026000190190911604601f810187900487028301870190935282825290949093509091508301828280156103d25780601f106103a7576101008083540402835291602001916103d2565b820191906000526020600020905b8154815290600101906020018083116103b557829003601f168201915b50505050506105ef90919063ffffffff16565b61042d576040805162461bcd60e51b81526020600482015260146024820152732ab9b2b91036bab9ba103737ba1032bc34b9ba1760611b604482015290519081900360640190fd5b81816000868660405180838380828437808301925050509250505090815260200160405180910390209190610463929190610715565b5050505050565b60006104c46040518060200160405280600081525086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6105ef169050565b8061052257506105226040518060200160405280600081525084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6105ef169050565b1561052f575060006105e7565b6105e483838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052506040519093508a9250899150808383808284379190910194855250506040805160209481900385018120805460026001821615610100026000190190911604601f810187900487028301870190935282825290949093509091508301828280156103d25780601f106103a7576101008083540402835291602001916103d2565b90505b949350505050565b600080836040516020018080602001828103825283818151815260200191508051906020019080838360005b8381101561063357818101518382015260200161061b565b50505050905090810190601f1680156106605780820380516001836020036101000a031916815260200191505b50925050506040516020818303038152906040528051906020012090506000836040516020018080602001828103825283818151815260200191508051906020019080838360005b838110156106c05781810151838201526020016106a8565b50505050905090810190601f1680156106ed5780820380516001836020036101000a031916815260200191505b5060408051808303601f19018152919052805160209091012094909414979650505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106107565782800160ff19823516178555610783565b82800160010185558215610783579182015b82811115610783578235825591602001919060010190610768565b5061078f929150610793565b5090565b6107ad91905b8082111561078f5760008155600101610799565b9056fea26469706673582212204c10f3ec26a661a91a5821e7ed19afb6de76b1677f7e2cdb771ffc2fa788bc0264736f6c634300060a0033",
}

// TaskABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskMetaData.ABI instead.
var TaskABI = TaskMetaData.ABI

// Deprecated: Use TaskMetaData.Sigs instead.
// TaskFuncSigs maps the 4-byte function signature to its string representation.
var TaskFuncSigs = TaskMetaData.Sigs

// TaskBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TaskMetaData.Bin instead.
var TaskBin = TaskMetaData.Bin

// DeployTask deploys a new Ethereum contract, binding an instance of Task to it.
func DeployTask(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address) (common.Address, *types.Transaction, *Task, error) {
	parsed, err := TaskMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TaskBin), backend, owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// Task is an auto generated Go binding around an Ethereum contract.
type Task struct {
	TaskCaller     // Read-only binding to the contract
	TaskTransactor // Write-only binding to the contract
	TaskFilterer   // Log filterer for contract events
}

// TaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskSession struct {
	Contract     *Task             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskCallerSession struct {
	Contract *TaskCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskTransactorSession struct {
	Contract     *TaskTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskRaw struct {
	Contract *Task // Generic contract binding to access the raw methods on
}

// TaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskCallerRaw struct {
	Contract *TaskCaller // Generic read-only contract binding to access the raw methods on
}

// TaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskTransactorRaw struct {
	Contract *TaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTask creates a new instance of Task, bound to a specific deployed contract.
func NewTask(address common.Address, backend bind.ContractBackend) (*Task, error) {
	contract, err := bindTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// NewTaskCaller creates a new read-only instance of Task, bound to a specific deployed contract.
func NewTaskCaller(address common.Address, caller bind.ContractCaller) (*TaskCaller, error) {
	contract, err := bindTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskCaller{contract: contract}, nil
}

// NewTaskTransactor creates a new write-only instance of Task, bound to a specific deployed contract.
func NewTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskTransactor, error) {
	contract, err := bindTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskTransactor{contract: contract}, nil
}

// NewTaskFilterer creates a new log filterer instance of Task, bound to a specific deployed contract.
func NewTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskFilterer, error) {
	contract, err := bindTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskFilterer{contract: contract}, nil
}

// bindTask binds a generic wrapper to an already deployed contract.
func bindTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Task.Contract.TaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Task.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) view returns(uint256 balance)
func (_Task *TaskCaller) BalanceOf(opts *bind.CallOpts, _owner string) (*big.Int, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) view returns(uint256 balance)
func (_Task *TaskSession) BalanceOf(_owner string) (*big.Int, error) {
	return _Task.Contract.BalanceOf(&_Task.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) view returns(uint256 balance)
func (_Task *TaskCallerSession) BalanceOf(_owner string) (*big.Int, error) {
	return _Task.Contract.BalanceOf(&_Task.CallOpts, _owner)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) view returns(bool)
func (_Task *TaskCaller) Login(opts *bind.CallOpts, username string, passwd string) (bool, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "login", username, passwd)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) view returns(bool)
func (_Task *TaskSession) Login(username string, passwd string) (bool, error) {
	return _Task.Contract.Login(&_Task.CallOpts, username, passwd)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) view returns(bool)
func (_Task *TaskCallerSession) Login(username string, passwd string) (bool, error) {
	return _Task.Contract.Login(&_Task.CallOpts, username, passwd)
}

// QryAllTasks is a free data retrieval call binding the contract method 0x8a3f7bb3.
//
// Solidity: function qryAllTasks() view returns((string,string,string,uint256,uint8,uint256,string)[])
func (_Task *TaskCaller) QryAllTasks(opts *bind.CallOpts) ([]TaskInfo, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "qryAllTasks")

	if err != nil {
		return *new([]TaskInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]TaskInfo)).(*[]TaskInfo)

	return out0, err

}

// QryAllTasks is a free data retrieval call binding the contract method 0x8a3f7bb3.
//
// Solidity: function qryAllTasks() view returns((string,string,string,uint256,uint8,uint256,string)[])
func (_Task *TaskSession) QryAllTasks() ([]TaskInfo, error) {
	return _Task.Contract.QryAllTasks(&_Task.CallOpts)
}

// QryAllTasks is a free data retrieval call binding the contract method 0x8a3f7bb3.
//
// Solidity: function qryAllTasks() view returns((string,string,string,uint256,uint8,uint256,string)[])
func (_Task *TaskCallerSession) QryAllTasks() ([]TaskInfo, error) {
	return _Task.Contract.QryAllTasks(&_Task.CallOpts)
}

// QryOneTask is a free data retrieval call binding the contract method 0x9b3add75.
//
// Solidity: function qryOneTask(uint256 taskID) view returns(string, string, uint256, string, uint8)
func (_Task *TaskCaller) QryOneTask(opts *bind.CallOpts, taskID *big.Int) (string, string, *big.Int, string, uint8, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "qryOneTask", taskID)

	if err != nil {
		return *new(string), *new(string), *new(*big.Int), *new(string), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return out0, out1, out2, out3, out4, err

}

// QryOneTask is a free data retrieval call binding the contract method 0x9b3add75.
//
// Solidity: function qryOneTask(uint256 taskID) view returns(string, string, uint256, string, uint8)
func (_Task *TaskSession) QryOneTask(taskID *big.Int) (string, string, *big.Int, string, uint8, error) {
	return _Task.Contract.QryOneTask(&_Task.CallOpts, taskID)
}

// QryOneTask is a free data retrieval call binding the contract method 0x9b3add75.
//
// Solidity: function qryOneTask(uint256 taskID) view returns(string, string, uint256, string, uint8)
func (_Task *TaskCallerSession) QryOneTask(taskID *big.Int) (string, string, *big.Int, string, uint8, error) {
	return _Task.Contract.QryOneTask(&_Task.CallOpts, taskID)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Task *TaskCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Task *TaskSession) TotalSupply() (*big.Int, error) {
	return _Task.Contract.TotalSupply(&_Task.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Task *TaskCallerSession) TotalSupply() (*big.Int, error) {
	return _Task.Contract.TotalSupply(&_Task.CallOpts)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskTransactor) Commit(opts *bind.TransactOpts, worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "commit", worker, passwd, taskID)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskSession) Commit(worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Commit(&_Task.TransactOpts, worker, passwd, taskID)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskTransactorSession) Commit(worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Commit(&_Task.TransactOpts, worker, passwd, taskID)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string passwd, uint256 taskID, string comment, uint8 status) returns()
func (_Task *TaskTransactor) Confirm(opts *bind.TransactOpts, issuer string, passwd string, taskID *big.Int, comment string, status uint8) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "confirm", issuer, passwd, taskID, comment, status)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string passwd, uint256 taskID, string comment, uint8 status) returns()
func (_Task *TaskSession) Confirm(issuer string, passwd string, taskID *big.Int, comment string, status uint8) (*types.Transaction, error) {
	return _Task.Contract.Confirm(&_Task.TransactOpts, issuer, passwd, taskID, comment, status)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string passwd, uint256 taskID, string comment, uint8 status) returns()
func (_Task *TaskTransactorSession) Confirm(issuer string, passwd string, taskID *big.Int, comment string, status uint8) (*types.Transaction, error) {
	return _Task.Contract.Confirm(&_Task.TransactOpts, issuer, passwd, taskID, comment, status)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string passwd, string desc, uint256 bonus) returns()
func (_Task *TaskTransactor) Issue(opts *bind.TransactOpts, issuer string, passwd string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "issue", issuer, passwd, desc, bonus)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string passwd, string desc, uint256 bonus) returns()
func (_Task *TaskSession) Issue(issuer string, passwd string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Issue(&_Task.TransactOpts, issuer, passwd, desc, bonus)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string passwd, string desc, uint256 bonus) returns()
func (_Task *TaskTransactorSession) Issue(issuer string, passwd string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Issue(&_Task.TransactOpts, issuer, passwd, desc, bonus)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _to, uint256 _value) returns()
func (_Task *TaskTransactor) Mint(opts *bind.TransactOpts, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _to, uint256 _value) returns()
func (_Task *TaskSession) Mint(_to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Mint(&_Task.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _to, uint256 _value) returns()
func (_Task *TaskTransactorSession) Mint(_to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Mint(&_Task.TransactOpts, _to, _value)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_Task *TaskTransactor) Register(opts *bind.TransactOpts, username string, passwd string) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "register", username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_Task *TaskSession) Register(username string, passwd string) (*types.Transaction, error) {
	return _Task.Contract.Register(&_Task.TransactOpts, username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_Task *TaskTransactorSession) Register(username string, passwd string) (*types.Transaction, error) {
	return _Task.Contract.Register(&_Task.TransactOpts, username, passwd)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskTransactor) Take(opts *bind.TransactOpts, worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "take", worker, passwd, taskID)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskSession) Take(worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Take(&_Task.TransactOpts, worker, passwd, taskID)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskTransactorSession) Take(worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Take(&_Task.TransactOpts, worker, passwd, taskID)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_Task *TaskTransactor) Transfer(opts *bind.TransactOpts, _from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "transfer", _from, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_Task *TaskSession) Transfer(_from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Transfer(&_Task.TransactOpts, _from, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_Task *TaskTransactorSession) Transfer(_from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Transfer(&_Task.TransactOpts, _from, _to, _value)
}

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_from\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_owner\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"35ee5f87": "balanceOf(string)",
		"056b01ce": "mint(string,uint256)",
		"18160ddd": "totalSupply()",
		"9b80b050": "transfer(string,string,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060008055600180546001600160a01b03191633179055610893806100366000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063056b01ce1461005157806318160ddd146100c157806335ee5f87146100db5780639b80b05014610149575b600080fd5b6100bf6004803603604081101561006757600080fd5b810190602081018135600160201b81111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460018302840111600160201b831117156100b457600080fd5b91935091503561021b565b005b6100c9610444565b60408051918252519081900360200190f35b6100c9600480360360208110156100f157600080fd5b810190602081018135600160201b81111561010b57600080fd5b82018360208201111561011d57600080fd5b803590602001918460018302840111600160201b8311171561013e57600080fd5b50909250905061044a565b6102076004803603606081101561015f57600080fd5b810190602081018135600160201b81111561017957600080fd5b82018360208201111561018b57600080fd5b803590602001918460018302840111600160201b831117156101ac57600080fd5b919390929091602081019035600160201b8111156101c957600080fd5b8201836020820111156101db57600080fd5b803590602001918460018302840111600160201b831117156101fc57600080fd5b91935091503561047b565b604080519115158252519081900360200190f35b6001546001600160a01b0316331461027a576040805162461bcd60e51b815260206004820152601760248201527f4f6e6c792061646d696e2063616e20646f207468697321000000000000000000604482015290519081900360640190fd5b6102d26040518060200160405280600081525084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff610709169050565b15610316576040805162461bcd60e51b815260206004820152600f60248201526e2a379036bab9ba1032bc34b9ba399760891b604482015290519081900360640190fd5b6000811161035d576040805162461bcd60e51b815260206004820152600f60248201526e2b30b63ab29036bab9ba101f10181760891b604482015290519081900360640190fd5b600054610370908263ffffffff61082f16565b6000819055506103ae81600285856040518083838082843791909101948552505060405192839003602001909220549291505063ffffffff61082f16565b60028484604051808383808284379190910194855250506040519283900360200183209390935550849150839080838380828437604080519390910183900383206573797374656d60d01b8452815193849003600601842088855291519096509094507f5358be4df107be4d9b023fc323f41d7109610225c6ef211b9d375b9fbd7ccc4f935091829003602001919050a3505050565b60005490565b6000600283836040518083838082843791909101948552505060405192839003602001909220549250505092915050565b60006104d56040518060200160405280600081525085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff610709169050565b15610519576040805162461bcd60e51b815260206004820152600f60248201526e2a379036bab9ba1032bc34b9ba399760891b604482015290519081900360640190fd5b60008211610560576040805162461bcd60e51b815260206004820152600f60248201526e2b30b63ab29036bab9ba101f10181760891b604482015290519081900360640190fd5b816002878760405180838380828437808301925050509250505090815260200160405180910390205410156105d3576040805162461bcd60e51b81526020600482015260146024820152732130b630b731b29036bab9ba1032b737bab3b41760611b604482015290519081900360640190fd5b61060b82600288886040518083838082843791909101948552505060405192839003602001909220549291505063ffffffff61084816565b6002878760405180838380828437808301925050509250505090815260200160405180910390208190555061066e82600286866040518083838082843791909101948552505060405192839003602001909220549291505063ffffffff61082f16565b600285856040518083838082843791909101948552505060405192839003602001832093909355508591508490808383808284376040519201829003822094508a93508992508190508383808284376040805193909101839003832089845290519095507f5358be4df107be4d9b023fc323f41d7109610225c6ef211b9d375b9fbd7ccc4f945091829003602001925050a395945050505050565b600080836040516020018080602001828103825283818151815260200191508051906020019080838360005b8381101561074d578181015183820152602001610735565b50505050905090810190601f16801561077a5780820380516001836020036101000a031916815260200191505b50925050506040516020818303038152906040528051906020012090506000836040516020018080602001828103825283818151815260200191508051906020019080838360005b838110156107da5781810151838201526020016107c2565b50505050905090810190601f1680156108075780820380516001836020036101000a031916815260200191505b5060408051808303601f19018152919052805160209091012094909414979650505050505050565b60008282018381101561084157600080fd5b9392505050565b60008282111561085757600080fd5b5090039056fea26469706673582212204e9eb7f1376d7d4a8ca22f130ce09c15a217d3467c5a49bdb983e67f80cef77964736f6c634300060a0033",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// Deprecated: Use TokenMetaData.Sigs instead.
// TokenFuncSigs maps the 4-byte function signature to its string representation.
var TokenFuncSigs = TokenMetaData.Sigs

// TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenMetaData.Bin instead.
var TokenBin = TokenMetaData.Bin

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) view returns(uint256 balance)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner string) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) view returns(uint256 balance)
func (_Token *TokenSession) BalanceOf(_owner string) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) view returns(uint256 balance)
func (_Token *TokenCallerSession) BalanceOf(_owner string) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _to, uint256 _value) returns()
func (_Token *TokenTransactor) Mint(opts *bind.TransactOpts, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _to, uint256 _value) returns()
func (_Token *TokenSession) Mint(_to string, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _to, uint256 _value) returns()
func (_Token *TokenTransactorSession) Mint(_to string, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _from, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_Token *TokenSession) Transfer(_from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _from, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_Token *TokenTransactorSession) Transfer(_from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _from, _to, _value)
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Hash
	To    common.Hash
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x5358be4df107be4d9b023fc323f41d7109610225c6ef211b9d375b9fbd7ccc4f.
//
// Solidity: event Transfer(string indexed _from, string indexed _to, uint256 _value)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []string, _to []string) (*TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x5358be4df107be4d9b023fc323f41d7109610225c6ef211b9d375b9fbd7ccc4f.
//
// Solidity: event Transfer(string indexed _from, string indexed _to, uint256 _value)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, _from []string, _to []string) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0x5358be4df107be4d9b023fc323f41d7109610225c6ef211b9d375b9fbd7ccc4f.
//
// Solidity: event Transfer(string indexed _from, string indexed _to, uint256 _value)
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserMetaData contains all meta data concerning the User contract.
var UserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"58467dbc": "login(string,string)",
		"3ffbd47f": "register(string,string)",
	},
	Bin: "0x608060405234801561001057600080fd5b50600180546001600160a01b031916331790556107e6806100326000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633ffbd47f1461003b57806358467dbc146100fb575b600080fd5b6100f96004803603604081101561005157600080fd5b810190602081018135600160201b81111561006b57600080fd5b82018360208201111561007d57600080fd5b803590602001918460018302840111600160201b8311171561009e57600080fd5b919390929091602081019035600160201b8111156100bb57600080fd5b8201836020820111156100cd57600080fd5b803590602001918460018302840111600160201b831117156100ee57600080fd5b5090925090506101cd565b005b6101b96004803603604081101561011157600080fd5b810190602081018135600160201b81111561012b57600080fd5b82018360208201111561013d57600080fd5b803590602001918460018302840111600160201b8311171561015e57600080fd5b919390929091602081019035600160201b81111561017b57600080fd5b82018360208201111561018d57600080fd5b803590602001918460018302840111600160201b831117156101ae57600080fd5b50909250905061046a565b604080519115158252519081900360200190f35b6102256040518060200160405280600081525085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6105ef169050565b15610277576040805162461bcd60e51b815260206004820152601760248201527f557365726e616d65206d757374206e6f74206e756c6c2e000000000000000000604482015290519081900360640190fd5b6102cf6040518060200160405280600081525083838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6105ef169050565b15610319576040805162461bcd60e51b81526020600482015260156024820152742830b9b9bbb21036bab9ba103737ba10373ab6361760591b604482015290519081900360640190fd5b6103e56040518060200160405280600081525060008686604051808383808284379190910194855250506040805160209481900385018120805460026001821615610100026000190190911604601f810187900487028301870190935282825290949093509091508301828280156103d25780601f106103a7576101008083540402835291602001916103d2565b820191906000526020600020905b8154815290600101906020018083116103b557829003601f168201915b50505050506105ef90919063ffffffff16565b61042d576040805162461bcd60e51b81526020600482015260146024820152732ab9b2b91036bab9ba103737ba1032bc34b9ba1760611b604482015290519081900360640190fd5b81816000868660405180838380828437808301925050509250505090815260200160405180910390209190610463929190610715565b5050505050565b60006104c46040518060200160405280600081525086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6105ef169050565b8061052257506105226040518060200160405280600081525084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6105ef169050565b1561052f575060006105e7565b6105e483838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052506040519093508a9250899150808383808284379190910194855250506040805160209481900385018120805460026001821615610100026000190190911604601f810187900487028301870190935282825290949093509091508301828280156103d25780601f106103a7576101008083540402835291602001916103d2565b90505b949350505050565b600080836040516020018080602001828103825283818151815260200191508051906020019080838360005b8381101561063357818101518382015260200161061b565b50505050905090810190601f1680156106605780820380516001836020036101000a031916815260200191505b50925050506040516020818303038152906040528051906020012090506000836040516020018080602001828103825283818151815260200191508051906020019080838360005b838110156106c05781810151838201526020016106a8565b50505050905090810190601f1680156106ed5780820380516001836020036101000a031916815260200191505b5060408051808303601f19018152919052805160209091012094909414979650505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106107565782800160ff19823516178555610783565b82800160010185558215610783579182015b82811115610783578235825591602001919060010190610768565b5061078f929150610793565b5090565b6107ad91905b8082111561078f5760008155600101610799565b9056fea26469706673582212204c10f3ec26a661a91a5821e7ed19afb6de76b1677f7e2cdb771ffc2fa788bc0264736f6c634300060a0033",
}

// UserABI is the input ABI used to generate the binding from.
// Deprecated: Use UserMetaData.ABI instead.
var UserABI = UserMetaData.ABI

// Deprecated: Use UserMetaData.Sigs instead.
// UserFuncSigs maps the 4-byte function signature to its string representation.
var UserFuncSigs = UserMetaData.Sigs

// UserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UserMetaData.Bin instead.
var UserBin = UserMetaData.Bin

// DeployUser deploys a new Ethereum contract, binding an instance of User to it.
func DeployUser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *User, error) {
	parsed, err := UserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// User is an auto generated Go binding around an Ethereum contract.
type User struct {
	UserCaller     // Read-only binding to the contract
	UserTransactor // Write-only binding to the contract
	UserFilterer   // Log filterer for contract events
}

// UserCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserSession struct {
	Contract     *User             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCallerSession struct {
	Contract *UserCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserTransactorSession struct {
	Contract     *UserTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRaw struct {
	Contract *User // Generic contract binding to access the raw methods on
}

// UserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCallerRaw struct {
	Contract *UserCaller // Generic read-only contract binding to access the raw methods on
}

// UserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserTransactorRaw struct {
	Contract *UserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUser creates a new instance of User, bound to a specific deployed contract.
func NewUser(address common.Address, backend bind.ContractBackend) (*User, error) {
	contract, err := bindUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// NewUserCaller creates a new read-only instance of User, bound to a specific deployed contract.
func NewUserCaller(address common.Address, caller bind.ContractCaller) (*UserCaller, error) {
	contract, err := bindUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserCaller{contract: contract}, nil
}

// NewUserTransactor creates a new write-only instance of User, bound to a specific deployed contract.
func NewUserTransactor(address common.Address, transactor bind.ContractTransactor) (*UserTransactor, error) {
	contract, err := bindUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserTransactor{contract: contract}, nil
}

// NewUserFilterer creates a new log filterer instance of User, bound to a specific deployed contract.
func NewUserFilterer(address common.Address, filterer bind.ContractFilterer) (*UserFilterer, error) {
	contract, err := bindUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserFilterer{contract: contract}, nil
}

// bindUser binds a generic wrapper to an already deployed contract.
func bindUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.UserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.contract.Transact(opts, method, params...)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) view returns(bool)
func (_User *UserCaller) Login(opts *bind.CallOpts, username string, passwd string) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "login", username, passwd)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) view returns(bool)
func (_User *UserSession) Login(username string, passwd string) (bool, error) {
	return _User.Contract.Login(&_User.CallOpts, username, passwd)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) view returns(bool)
func (_User *UserCallerSession) Login(username string, passwd string) (bool, error) {
	return _User.Contract.Login(&_User.CallOpts, username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_User *UserTransactor) Register(opts *bind.TransactOpts, username string, passwd string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "register", username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_User *UserSession) Register(username string, passwd string) (*types.Transaction, error) {
	return _User.Contract.Register(&_User.TransactOpts, username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_User *UserTransactorSession) Register(username string, passwd string) (*types.Transaction, error) {
	return _User.Contract.Register(&_User.TransactOpts, username, passwd)
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220dd575a179343c8e9eddcf25d83dadb87fbfe44e270cf709c609be74a9128ce4164736f6c634300060a0033",
}

// StringsABI is the input ABI used to generate the binding from.
// Deprecated: Use StringsMetaData.ABI instead.
var StringsABI = StringsMetaData.ABI

// StringsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StringsMetaData.Bin instead.
var StringsBin = StringsMetaData.Bin

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := StringsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}
