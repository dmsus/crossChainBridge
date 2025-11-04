// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
	_ = abi.ConvertType
)

// BridgeEthereumMetaData contains all meta data concerning the BridgeEthereum contract.
var BridgeEthereumMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockTokens\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"targetChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseBridge\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseBridge\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BridgePaused\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeUnpaused\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensLocked\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"targetAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"targetChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BridgeIsPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAllowance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainId\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161119d38038061119d83398181016040528101906100319190610243565b60015f8190555061005461004961011b60201b60201c565b61012260201b60201c565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100b9576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250505f60035f6101000a81548160ff0219169083151502179055506101153361012260201b60201c565b5061026e565b5f33905090565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610212826101e9565b9050919050565b61022281610208565b811461022c575f5ffd5b50565b5f8151905061023d81610219565b92915050565b5f60208284031215610258576102576101e5565b5b5f6102658482850161022f565b91505092915050565b608051610f0261029b5f395f81816102dd015281816103ae01528181610481015261079a0152610f025ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c8063a82f143c1161006f578063a82f143c14610117578063affed0e014610121578063b187bd261461013f578063d087d2881461015d578063f2fde38b1461017b578063fc0c546a14610197576100a7565b8063154b3db0146100ab5780635c975abb146100c7578063715018a6146100e55780637dd0480f146100ef5780638da5cb5b146100f9575b5f5ffd5b6100c560048036038101906100c091906109eb565b6101b5565b005b6100cf6105d8565b6040516100dc9190610a55565b60405180910390f35b6100ed6105ea565b005b6100f76105fd565b005b610101610664565b60405161010e9190610a7d565b60405180910390f35b61011f61068c565b005b6101296106f2565b6040516101369190610aa5565b60405180910390f35b6101476106f8565b6040516101549190610a55565b60405180910390f35b61016561070d565b6040516101729190610aa5565b60405180910390f35b61019560048036038101906101909190610abe565b610716565b005b61019f610798565b6040516101ac9190610b44565b60405180910390f35b6101bd6107bc565b60035f9054906101000a900460ff1615610203576040517f7cf03c4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f830361023c576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036102a1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f81036102da576040517fc84885d400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b827f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016103349190610a7d565b602060405180830381865afa15801561034f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103739190610b71565b10156103ab576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b827f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401610407929190610b9c565b602060405180830381865afa158015610422573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104469190610b71565b101561047e576040517f13be252b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330876040518463ffffffff1660e01b81526004016104dc93929190610bc3565b6020604051808303815f875af11580156104f8573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061051c9190610c22565b90508061055e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055590610ca7565b60405180910390fd5b60025f81548092919061057090610cf2565b91905055506002543373ffffffffffffffffffffffffffffffffffffffff167fc65d896fc2ca9fd58f5384f9b647c3ceb103fd3029e659645afa90f0a1520ef98686866040516105c293929190610d39565b60405180910390a3506105d3610809565b505050565b60035f9054906101000a900460ff1681565b6105f2610812565b6105fb5f610890565b565b610605610812565b600160035f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f80c6ef1d6f41faaf47f10c6da1ae6d45d4773d53d4c79bc1ef2737b10dbed57960405160405180910390a2565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610694610812565b5f60035f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167fa668d29cf81d3acc05e4a35a5bf898bb2c919fc7a4366e536bbd43522cb670f060405160405180910390a2565b60025481565b5f60035f9054906101000a900460ff16905090565b5f600254905090565b61071e610812565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361078c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078390610dde565b60405180910390fd5b61079581610890565b50565b7f000000000000000000000000000000000000000000000000000000000000000081565b60025f5403610800576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f790610e46565b60405180910390fd5b60025f81905550565b60015f81905550565b61081a610953565b73ffffffffffffffffffffffffffffffffffffffff16610838610664565b73ffffffffffffffffffffffffffffffffffffffff161461088e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088590610eae565b60405180910390fd5b565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b5f5ffd5b5f819050919050565b6109708161095e565b811461097a575f5ffd5b50565b5f8135905061098b81610967565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6109ba82610991565b9050919050565b6109ca816109b0565b81146109d4575f5ffd5b50565b5f813590506109e5816109c1565b92915050565b5f5f5f60608486031215610a0257610a0161095a565b5b5f610a0f8682870161097d565b9350506020610a20868287016109d7565b9250506040610a318682870161097d565b9150509250925092565b5f8115159050919050565b610a4f81610a3b565b82525050565b5f602082019050610a685f830184610a46565b92915050565b610a77816109b0565b82525050565b5f602082019050610a905f830184610a6e565b92915050565b610a9f8161095e565b82525050565b5f602082019050610ab85f830184610a96565b92915050565b5f60208284031215610ad357610ad261095a565b5b5f610ae0848285016109d7565b91505092915050565b5f819050919050565b5f610b0c610b07610b0284610991565b610ae9565b610991565b9050919050565b5f610b1d82610af2565b9050919050565b5f610b2e82610b13565b9050919050565b610b3e81610b24565b82525050565b5f602082019050610b575f830184610b35565b92915050565b5f81519050610b6b81610967565b92915050565b5f60208284031215610b8657610b8561095a565b5b5f610b9384828501610b5d565b91505092915050565b5f604082019050610baf5f830185610a6e565b610bbc6020830184610a6e565b9392505050565b5f606082019050610bd65f830186610a6e565b610be36020830185610a6e565b610bf06040830184610a96565b949350505050565b610c0181610a3b565b8114610c0b575f5ffd5b50565b5f81519050610c1c81610bf8565b92915050565b5f60208284031215610c3757610c3661095a565b5b5f610c4484828501610c0e565b91505092915050565b5f82825260208201905092915050565b7f546f6b656e207472616e73666572206661696c656400000000000000000000005f82015250565b5f610c91601583610c4d565b9150610c9c82610c5d565b602082019050919050565b5f6020820190508181035f830152610cbe81610c85565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610cfc8261095e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610d2e57610d2d610cc5565b5b600182019050919050565b5f606082019050610d4c5f830186610a96565b610d596020830185610a6e565b610d666040830184610a96565b949350505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f610dc8602683610c4d565b9150610dd382610d6e565b604082019050919050565b5f6020820190508181035f830152610df581610dbc565b9050919050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c005f82015250565b5f610e30601f83610c4d565b9150610e3b82610dfc565b602082019050919050565b5f6020820190508181035f830152610e5d81610e24565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f610e98602083610c4d565b9150610ea382610e64565b602082019050919050565b5f6020820190508181035f830152610ec581610e8c565b905091905056fea264697066735822122020b6f70528d923e78dbb99110558d9458ee0aa46ab94b4a5f39f7a533352466164736f6c634300081e0033",
}

// BridgeEthereumABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeEthereumMetaData.ABI instead.
var BridgeEthereumABI = BridgeEthereumMetaData.ABI

// BridgeEthereumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeEthereumMetaData.Bin instead.
var BridgeEthereumBin = BridgeEthereumMetaData.Bin

// DeployBridgeEthereum deploys a new Ethereum contract, binding an instance of BridgeEthereum to it.
func DeployBridgeEthereum(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *BridgeEthereum, error) {
	parsed, err := BridgeEthereumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeEthereumBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeEthereum{BridgeEthereumCaller: BridgeEthereumCaller{contract: contract}, BridgeEthereumTransactor: BridgeEthereumTransactor{contract: contract}, BridgeEthereumFilterer: BridgeEthereumFilterer{contract: contract}}, nil
}

// BridgeEthereum is an auto generated Go binding around an Ethereum contract.
type BridgeEthereum struct {
	BridgeEthereumCaller     // Read-only binding to the contract
	BridgeEthereumTransactor // Write-only binding to the contract
	BridgeEthereumFilterer   // Log filterer for contract events
}

// BridgeEthereumCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeEthereumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeEthereumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeEthereumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeEthereumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeEthereumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeEthereumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeEthereumSession struct {
	Contract     *BridgeEthereum   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeEthereumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeEthereumCallerSession struct {
	Contract *BridgeEthereumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeEthereumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeEthereumTransactorSession struct {
	Contract     *BridgeEthereumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeEthereumRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeEthereumRaw struct {
	Contract *BridgeEthereum // Generic contract binding to access the raw methods on
}

// BridgeEthereumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeEthereumCallerRaw struct {
	Contract *BridgeEthereumCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeEthereumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeEthereumTransactorRaw struct {
	Contract *BridgeEthereumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeEthereum creates a new instance of BridgeEthereum, bound to a specific deployed contract.
func NewBridgeEthereum(address common.Address, backend bind.ContractBackend) (*BridgeEthereum, error) {
	contract, err := bindBridgeEthereum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereum{BridgeEthereumCaller: BridgeEthereumCaller{contract: contract}, BridgeEthereumTransactor: BridgeEthereumTransactor{contract: contract}, BridgeEthereumFilterer: BridgeEthereumFilterer{contract: contract}}, nil
}

// NewBridgeEthereumCaller creates a new read-only instance of BridgeEthereum, bound to a specific deployed contract.
func NewBridgeEthereumCaller(address common.Address, caller bind.ContractCaller) (*BridgeEthereumCaller, error) {
	contract, err := bindBridgeEthereum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumCaller{contract: contract}, nil
}

// NewBridgeEthereumTransactor creates a new write-only instance of BridgeEthereum, bound to a specific deployed contract.
func NewBridgeEthereumTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeEthereumTransactor, error) {
	contract, err := bindBridgeEthereum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumTransactor{contract: contract}, nil
}

// NewBridgeEthereumFilterer creates a new log filterer instance of BridgeEthereum, bound to a specific deployed contract.
func NewBridgeEthereumFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeEthereumFilterer, error) {
	contract, err := bindBridgeEthereum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumFilterer{contract: contract}, nil
}

// bindBridgeEthereum binds a generic wrapper to an already deployed contract.
func bindBridgeEthereum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeEthereumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeEthereum *BridgeEthereumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeEthereum.Contract.BridgeEthereumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeEthereum *BridgeEthereumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeEthereumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeEthereum *BridgeEthereumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeEthereumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeEthereum *BridgeEthereumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeEthereum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeEthereum *BridgeEthereumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeEthereum *BridgeEthereumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumSession) GetNonce() (*big.Int, error) {
	return _BridgeEthereum.Contract.GetNonce(&_BridgeEthereum.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCallerSession) GetNonce() (*big.Int, error) {
	return _BridgeEthereum.Contract.GetNonce(&_BridgeEthereum.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_BridgeEthereum *BridgeEthereumCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) IsPaused() (bool, error) {
	return _BridgeEthereum.Contract.IsPaused(&_BridgeEthereum.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_BridgeEthereum *BridgeEthereumCallerSession) IsPaused() (bool, error) {
	return _BridgeEthereum.Contract.IsPaused(&_BridgeEthereum.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumSession) Nonce() (*big.Int, error) {
	return _BridgeEthereum.Contract.Nonce(&_BridgeEthereum.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCallerSession) Nonce() (*big.Int, error) {
	return _BridgeEthereum.Contract.Nonce(&_BridgeEthereum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeEthereum *BridgeEthereumCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeEthereum *BridgeEthereumSession) Owner() (common.Address, error) {
	return _BridgeEthereum.Contract.Owner(&_BridgeEthereum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeEthereum *BridgeEthereumCallerSession) Owner() (common.Address, error) {
	return _BridgeEthereum.Contract.Owner(&_BridgeEthereum.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeEthereum *BridgeEthereumCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) Paused() (bool, error) {
	return _BridgeEthereum.Contract.Paused(&_BridgeEthereum.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeEthereum *BridgeEthereumCallerSession) Paused() (bool, error) {
	return _BridgeEthereum.Contract.Paused(&_BridgeEthereum.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgeEthereum *BridgeEthereumCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgeEthereum *BridgeEthereumSession) Token() (common.Address, error) {
	return _BridgeEthereum.Contract.Token(&_BridgeEthereum.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgeEthereum *BridgeEthereumCallerSession) Token() (common.Address, error) {
	return _BridgeEthereum.Contract.Token(&_BridgeEthereum.CallOpts)
}

// LockTokens is a paid mutator transaction binding the contract method 0x154b3db0.
//
// Solidity: function lockTokens(uint256 amount, address targetAddress, uint256 targetChainId) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) LockTokens(opts *bind.TransactOpts, amount *big.Int, targetAddress common.Address, targetChainId *big.Int) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "lockTokens", amount, targetAddress, targetChainId)
}

// LockTokens is a paid mutator transaction binding the contract method 0x154b3db0.
//
// Solidity: function lockTokens(uint256 amount, address targetAddress, uint256 targetChainId) returns()
func (_BridgeEthereum *BridgeEthereumSession) LockTokens(amount *big.Int, targetAddress common.Address, targetChainId *big.Int) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.LockTokens(&_BridgeEthereum.TransactOpts, amount, targetAddress, targetChainId)
}

// LockTokens is a paid mutator transaction binding the contract method 0x154b3db0.
//
// Solidity: function lockTokens(uint256 amount, address targetAddress, uint256 targetChainId) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) LockTokens(amount *big.Int, targetAddress common.Address, targetChainId *big.Int) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.LockTokens(&_BridgeEthereum.TransactOpts, amount, targetAddress, targetChainId)
}

// PauseBridge is a paid mutator transaction binding the contract method 0x7dd0480f.
//
// Solidity: function pauseBridge() returns()
func (_BridgeEthereum *BridgeEthereumTransactor) PauseBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "pauseBridge")
}

// PauseBridge is a paid mutator transaction binding the contract method 0x7dd0480f.
//
// Solidity: function pauseBridge() returns()
func (_BridgeEthereum *BridgeEthereumSession) PauseBridge() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PauseBridge(&_BridgeEthereum.TransactOpts)
}

// PauseBridge is a paid mutator transaction binding the contract method 0x7dd0480f.
//
// Solidity: function pauseBridge() returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) PauseBridge() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PauseBridge(&_BridgeEthereum.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeEthereum *BridgeEthereumTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeEthereum *BridgeEthereumSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RenounceOwnership(&_BridgeEthereum.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RenounceOwnership(&_BridgeEthereum.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeEthereum *BridgeEthereumSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.TransferOwnership(&_BridgeEthereum.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.TransferOwnership(&_BridgeEthereum.TransactOpts, newOwner)
}

// UnpauseBridge is a paid mutator transaction binding the contract method 0xa82f143c.
//
// Solidity: function unpauseBridge() returns()
func (_BridgeEthereum *BridgeEthereumTransactor) UnpauseBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "unpauseBridge")
}

// UnpauseBridge is a paid mutator transaction binding the contract method 0xa82f143c.
//
// Solidity: function unpauseBridge() returns()
func (_BridgeEthereum *BridgeEthereumSession) UnpauseBridge() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.UnpauseBridge(&_BridgeEthereum.TransactOpts)
}

// UnpauseBridge is a paid mutator transaction binding the contract method 0xa82f143c.
//
// Solidity: function unpauseBridge() returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) UnpauseBridge() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.UnpauseBridge(&_BridgeEthereum.TransactOpts)
}

// BridgeEthereumBridgePausedIterator is returned from FilterBridgePaused and is used to iterate over the raw logs and unpacked data for BridgePaused events raised by the BridgeEthereum contract.
type BridgeEthereumBridgePausedIterator struct {
	Event *BridgeEthereumBridgePaused // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumBridgePaused)
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
		it.Event = new(BridgeEthereumBridgePaused)
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
func (it *BridgeEthereumBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumBridgePaused represents a BridgePaused event raised by the BridgeEthereum contract.
type BridgeEthereumBridgePaused struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgePaused is a free log retrieval operation binding the contract event 0x80c6ef1d6f41faaf47f10c6da1ae6d45d4773d53d4c79bc1ef2737b10dbed579.
//
// Solidity: event BridgePaused(address indexed admin)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterBridgePaused(opts *bind.FilterOpts, admin []common.Address) (*BridgeEthereumBridgePausedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "BridgePaused", adminRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumBridgePausedIterator{contract: _BridgeEthereum.contract, event: "BridgePaused", logs: logs, sub: sub}, nil
}

// WatchBridgePaused is a free log subscription operation binding the contract event 0x80c6ef1d6f41faaf47f10c6da1ae6d45d4773d53d4c79bc1ef2737b10dbed579.
//
// Solidity: event BridgePaused(address indexed admin)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchBridgePaused(opts *bind.WatchOpts, sink chan<- *BridgeEthereumBridgePaused, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "BridgePaused", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumBridgePaused)
				if err := _BridgeEthereum.contract.UnpackLog(event, "BridgePaused", log); err != nil {
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

// ParseBridgePaused is a log parse operation binding the contract event 0x80c6ef1d6f41faaf47f10c6da1ae6d45d4773d53d4c79bc1ef2737b10dbed579.
//
// Solidity: event BridgePaused(address indexed admin)
func (_BridgeEthereum *BridgeEthereumFilterer) ParseBridgePaused(log types.Log) (*BridgeEthereumBridgePaused, error) {
	event := new(BridgeEthereumBridgePaused)
	if err := _BridgeEthereum.contract.UnpackLog(event, "BridgePaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumBridgeUnpausedIterator is returned from FilterBridgeUnpaused and is used to iterate over the raw logs and unpacked data for BridgeUnpaused events raised by the BridgeEthereum contract.
type BridgeEthereumBridgeUnpausedIterator struct {
	Event *BridgeEthereumBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumBridgeUnpaused)
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
		it.Event = new(BridgeEthereumBridgeUnpaused)
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
func (it *BridgeEthereumBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumBridgeUnpaused represents a BridgeUnpaused event raised by the BridgeEthereum contract.
type BridgeEthereumBridgeUnpaused struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgeUnpaused is a free log retrieval operation binding the contract event 0xa668d29cf81d3acc05e4a35a5bf898bb2c919fc7a4366e536bbd43522cb670f0.
//
// Solidity: event BridgeUnpaused(address indexed admin)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterBridgeUnpaused(opts *bind.FilterOpts, admin []common.Address) (*BridgeEthereumBridgeUnpausedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "BridgeUnpaused", adminRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumBridgeUnpausedIterator{contract: _BridgeEthereum.contract, event: "BridgeUnpaused", logs: logs, sub: sub}, nil
}

// WatchBridgeUnpaused is a free log subscription operation binding the contract event 0xa668d29cf81d3acc05e4a35a5bf898bb2c919fc7a4366e536bbd43522cb670f0.
//
// Solidity: event BridgeUnpaused(address indexed admin)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchBridgeUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeEthereumBridgeUnpaused, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "BridgeUnpaused", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumBridgeUnpaused)
				if err := _BridgeEthereum.contract.UnpackLog(event, "BridgeUnpaused", log); err != nil {
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

// ParseBridgeUnpaused is a log parse operation binding the contract event 0xa668d29cf81d3acc05e4a35a5bf898bb2c919fc7a4366e536bbd43522cb670f0.
//
// Solidity: event BridgeUnpaused(address indexed admin)
func (_BridgeEthereum *BridgeEthereumFilterer) ParseBridgeUnpaused(log types.Log) (*BridgeEthereumBridgeUnpaused, error) {
	event := new(BridgeEthereumBridgeUnpaused)
	if err := _BridgeEthereum.contract.UnpackLog(event, "BridgeUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeEthereum contract.
type BridgeEthereumOwnershipTransferredIterator struct {
	Event *BridgeEthereumOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumOwnershipTransferred)
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
		it.Event = new(BridgeEthereumOwnershipTransferred)
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
func (it *BridgeEthereumOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeEthereum contract.
type BridgeEthereumOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeEthereumOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumOwnershipTransferredIterator{contract: _BridgeEthereum.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeEthereumOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumOwnershipTransferred)
				if err := _BridgeEthereum.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeEthereum *BridgeEthereumFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeEthereumOwnershipTransferred, error) {
	event := new(BridgeEthereumOwnershipTransferred)
	if err := _BridgeEthereum.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumTokensLockedIterator is returned from FilterTokensLocked and is used to iterate over the raw logs and unpacked data for TokensLocked events raised by the BridgeEthereum contract.
type BridgeEthereumTokensLockedIterator struct {
	Event *BridgeEthereumTokensLocked // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumTokensLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumTokensLocked)
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
		it.Event = new(BridgeEthereumTokensLocked)
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
func (it *BridgeEthereumTokensLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumTokensLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumTokensLocked represents a TokensLocked event raised by the BridgeEthereum contract.
type BridgeEthereumTokensLocked struct {
	User          common.Address
	Amount        *big.Int
	Nonce         *big.Int
	TargetAddress common.Address
	TargetChainId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokensLocked is a free log retrieval operation binding the contract event 0xc65d896fc2ca9fd58f5384f9b647c3ceb103fd3029e659645afa90f0a1520ef9.
//
// Solidity: event TokensLocked(address indexed user, uint256 amount, uint256 indexed nonce, address targetAddress, uint256 targetChainId)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterTokensLocked(opts *bind.FilterOpts, user []common.Address, nonce []*big.Int) (*BridgeEthereumTokensLockedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "TokensLocked", userRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumTokensLockedIterator{contract: _BridgeEthereum.contract, event: "TokensLocked", logs: logs, sub: sub}, nil
}

// WatchTokensLocked is a free log subscription operation binding the contract event 0xc65d896fc2ca9fd58f5384f9b647c3ceb103fd3029e659645afa90f0a1520ef9.
//
// Solidity: event TokensLocked(address indexed user, uint256 amount, uint256 indexed nonce, address targetAddress, uint256 targetChainId)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchTokensLocked(opts *bind.WatchOpts, sink chan<- *BridgeEthereumTokensLocked, user []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "TokensLocked", userRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumTokensLocked)
				if err := _BridgeEthereum.contract.UnpackLog(event, "TokensLocked", log); err != nil {
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

// ParseTokensLocked is a log parse operation binding the contract event 0xc65d896fc2ca9fd58f5384f9b647c3ceb103fd3029e659645afa90f0a1520ef9.
//
// Solidity: event TokensLocked(address indexed user, uint256 amount, uint256 indexed nonce, address targetAddress, uint256 targetChainId)
func (_BridgeEthereum *BridgeEthereumFilterer) ParseTokensLocked(log types.Log) (*BridgeEthereumTokensLocked, error) {
	event := new(BridgeEthereumTokensLocked)
	if err := _BridgeEthereum.contract.UnpackLog(event, "TokensLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
