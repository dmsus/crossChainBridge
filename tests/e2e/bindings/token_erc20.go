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

// TokenERC20MetaData contains all meta data concerning the TokenERC20 contract.
var TokenERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040518060400160405280601581526020017f42726964676520546f6b656e20457468657265756d00000000000000000000008152506040518060400160405280600781526020017f4252472d45544800000000000000000000000000000000000000000000000000815250816003908161008b9190610497565b50806004908161009b9190610497565b505050601260055f6101000a81548160ff021916908360ff1602179055506100f13360055f9054906101000a900460ff16600a6100d891906106ce565b620f42406100e69190610718565b6100f660201b60201c565b61082c565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610164576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015b906107b3565b60405180910390fd5b6101755f838361025060201b60201c565b8060025f82825461018691906107d1565b92505081905550805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508173ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516102339190610813565b60405180910390a361024c5f838361025560201b60201c565b5050565b505050565b505050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806102d557607f821691505b6020821081036102e8576102e7610291565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261034a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261030f565b610354868361030f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61039861039361038e8461036c565b610375565b61036c565b9050919050565b5f819050919050565b6103b18361037e565b6103c56103bd8261039f565b84845461031b565b825550505050565b5f5f905090565b6103dc6103cd565b6103e78184846103a8565b505050565b5b8181101561040a576103ff5f826103d4565b6001810190506103ed565b5050565b601f82111561044f57610420816102ee565b61042984610300565b81016020851015610438578190505b61044c61044485610300565b8301826103ec565b50505b505050565b5f82821c905092915050565b5f61046f5f1984600802610454565b1980831691505092915050565b5f6104878383610460565b9150826002028217905092915050565b6104a08261025a565b67ffffffffffffffff8111156104b9576104b8610264565b5b6104c382546102be565b6104ce82828561040e565b5f60209050601f8311600181146104ff575f84156104ed578287015190505b6104f7858261047c565b86555061055e565b601f19841661050d866102ee565b5f5b828110156105345784890151825560018201915060208501945060208101905061050f565b86831015610551578489015161054d601f891682610460565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8160011c9050919050565b5f5f8291508390505b60018511156105e8578086048111156105c4576105c3610566565b5b60018516156105d35780820291505b80810290506105e185610593565b94506105a8565b94509492505050565b5f8261060057600190506106bb565b8161060d575f90506106bb565b8160018114610623576002811461062d5761065c565b60019150506106bb565b60ff84111561063f5761063e610566565b5b8360020a91508482111561065657610655610566565b5b506106bb565b5060208310610133831016604e8410600b84101617156106915782820a90508381111561068c5761068b610566565b5b6106bb565b61069e848484600161059f565b925090508184048111156106b5576106b4610566565b5b81810290505b9392505050565b5f60ff82169050919050565b5f6106d88261036c565b91506106e3836106c2565b92506107107fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846105f1565b905092915050565b5f6107228261036c565b915061072d8361036c565b925082820261073b8161036c565b9150828204841483151761075257610751610566565b5b5092915050565b5f82825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f2061646472657373005f82015250565b5f61079d601f83610759565b91506107a882610769565b602082019050919050565b5f6020820190508181035f8301526107ca81610791565b9050919050565b5f6107db8261036c565b91506107e68361036c565b92508282019050808211156107fe576107fd610566565b5b92915050565b61080d8161036c565b82525050565b5f6020820190506108265f830184610804565b92915050565b6113a1806108395f395ff3fe608060405234801561000f575f5ffd5b50600436106100b2575f3560e01c806340c10f191161006f57806340c10f19146101a057806370a08231146101bc57806395d89b41146101ec578063a457c2d71461020a578063a9059cbb1461023a578063dd62ed3e1461026a576100b2565b806306fdde03146100b6578063095ea7b3146100d457806318160ddd1461010457806323b872dd14610122578063313ce567146101525780633950935114610170575b5f5ffd5b6100be61029a565b6040516100cb9190610c53565b60405180910390f35b6100ee60048036038101906100e99190610d04565b61032a565b6040516100fb9190610d5c565b60405180910390f35b61010c61034c565b6040516101199190610d84565b60405180910390f35b61013c60048036038101906101379190610d9d565b610355565b6040516101499190610d5c565b60405180910390f35b61015a610383565b6040516101679190610e08565b60405180910390f35b61018a60048036038101906101859190610d04565b610398565b6040516101979190610d5c565b60405180910390f35b6101ba60048036038101906101b59190610d04565b6103ce565b005b6101d660048036038101906101d19190610e21565b6103dc565b6040516101e39190610d84565b60405180910390f35b6101f4610421565b6040516102019190610c53565b60405180910390f35b610224600480360381019061021f9190610d04565b6104b1565b6040516102319190610d5c565b60405180910390f35b610254600480360381019061024f9190610d04565b610526565b6040516102619190610d5c565b60405180910390f35b610284600480360381019061027f9190610e4c565b610548565b6040516102919190610d84565b60405180910390f35b6060600380546102a990610eb7565b80601f01602080910402602001604051908101604052809291908181526020018280546102d590610eb7565b80156103205780601f106102f757610100808354040283529160200191610320565b820191905f5260205f20905b81548152906001019060200180831161030357829003601f168201915b5050505050905090565b5f5f6103346105ca565b90506103418185856105d1565b600191505092915050565b5f600254905090565b5f5f61035f6105ca565b905061036c858285610794565b61037785858561081f565b60019150509392505050565b5f60055f9054906101000a900460ff16905090565b5f5f6103a26105ca565b90506103c38185856103b48589610548565b6103be9190610f14565b6105d1565b600191505092915050565b6103d88282610a8b565b5050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60606004805461043090610eb7565b80601f016020809104026020016040519081016040528092919081815260200182805461045c90610eb7565b80156104a75780601f1061047e576101008083540402835291602001916104a7565b820191905f5260205f20905b81548152906001019060200180831161048a57829003601f168201915b5050505050905090565b5f5f6104bb6105ca565b90505f6104c88286610548565b90508381101561050d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050490610fb7565b60405180910390fd5b61051a82868684036105d1565b60019250505092915050565b5f5f6105306105ca565b905061053d81858561081f565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f33905090565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361063f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063690611045565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036106ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a4906110d3565b60405180910390fd5b8060015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516107879190610d84565b60405180910390a3505050565b5f61079f8484610548565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610819578181101561080b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108029061113b565b60405180910390fd5b61081884848484036105d1565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361088d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610884906111c9565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f290611257565b60405180910390fd5b610906838383610bd9565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610989576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610980906112e5565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550815f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a729190610d84565b60405180910390a3610a85848484610bde565b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610af9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af09061134d565b60405180910390fd5b610b045f8383610bd9565b8060025f828254610b159190610f14565b92505081905550805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508173ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610bc29190610d84565b60405180910390a3610bd55f8383610bde565b5050565b505050565b505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610c2582610be3565b610c2f8185610bed565b9350610c3f818560208601610bfd565b610c4881610c0b565b840191505092915050565b5f6020820190508181035f830152610c6b8184610c1b565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ca082610c77565b9050919050565b610cb081610c96565b8114610cba575f5ffd5b50565b5f81359050610ccb81610ca7565b92915050565b5f819050919050565b610ce381610cd1565b8114610ced575f5ffd5b50565b5f81359050610cfe81610cda565b92915050565b5f5f60408385031215610d1a57610d19610c73565b5b5f610d2785828601610cbd565b9250506020610d3885828601610cf0565b9150509250929050565b5f8115159050919050565b610d5681610d42565b82525050565b5f602082019050610d6f5f830184610d4d565b92915050565b610d7e81610cd1565b82525050565b5f602082019050610d975f830184610d75565b92915050565b5f5f5f60608486031215610db457610db3610c73565b5b5f610dc186828701610cbd565b9350506020610dd286828701610cbd565b9250506040610de386828701610cf0565b9150509250925092565b5f60ff82169050919050565b610e0281610ded565b82525050565b5f602082019050610e1b5f830184610df9565b92915050565b5f60208284031215610e3657610e35610c73565b5b5f610e4384828501610cbd565b91505092915050565b5f5f60408385031215610e6257610e61610c73565b5b5f610e6f85828601610cbd565b9250506020610e8085828601610cbd565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610ece57607f821691505b602082108103610ee157610ee0610e8a565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610f1e82610cd1565b9150610f2983610cd1565b9250828201905080821115610f4157610f40610ee7565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f775f8201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b5f610fa1602583610bed565b9150610fac82610f47565b604082019050919050565b5f6020820190508181035f830152610fce81610f95565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f206164645f8201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b5f61102f602483610bed565b915061103a82610fd5565b604082019050919050565b5f6020820190508181035f83015261105c81611023565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f2061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f6110bd602283610bed565b91506110c882611063565b604082019050919050565b5f6020820190508181035f8301526110ea816110b1565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000005f82015250565b5f611125601d83610bed565b9150611130826110f1565b602082019050919050565b5f6020820190508181035f83015261115281611119565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f2061645f8201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b5f6111b3602583610bed565b91506111be82611159565b604082019050919050565b5f6020820190508181035f8301526111e0816111a7565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f20616464725f8201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b5f611241602383610bed565b915061124c826111e7565b604082019050919050565b5f6020820190508181035f83015261126e81611235565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320625f8201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b5f6112cf602683610bed565b91506112da82611275565b604082019050919050565b5f6020820190508181035f8301526112fc816112c3565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f2061646472657373005f82015250565b5f611337601f83610bed565b915061134282611303565b602082019050919050565b5f6020820190508181035f8301526113648161132b565b905091905056fea26469706673582212203586416dd22c3bb11fd1e13b49bcd2ed36e7c576744a57f4f42dc39e939f2e9364736f6c634300081e0033",
}

// TokenERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenERC20MetaData.ABI instead.
var TokenERC20ABI = TokenERC20MetaData.ABI

// TokenERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenERC20MetaData.Bin instead.
var TokenERC20Bin = TokenERC20MetaData.Bin

// DeployTokenERC20 deploys a new Ethereum contract, binding an instance of TokenERC20 to it.
func DeployTokenERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenERC20, error) {
	parsed, err := TokenERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenERC20{TokenERC20Caller: TokenERC20Caller{contract: contract}, TokenERC20Transactor: TokenERC20Transactor{contract: contract}, TokenERC20Filterer: TokenERC20Filterer{contract: contract}}, nil
}

// TokenERC20 is an auto generated Go binding around an Ethereum contract.
type TokenERC20 struct {
	TokenERC20Caller     // Read-only binding to the contract
	TokenERC20Transactor // Write-only binding to the contract
	TokenERC20Filterer   // Log filterer for contract events
}

// TokenERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type TokenERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenERC20Session struct {
	Contract     *TokenERC20       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenERC20CallerSession struct {
	Contract *TokenERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenERC20TransactorSession struct {
	Contract     *TokenERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type TokenERC20Raw struct {
	Contract *TokenERC20 // Generic contract binding to access the raw methods on
}

// TokenERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenERC20CallerRaw struct {
	Contract *TokenERC20Caller // Generic read-only contract binding to access the raw methods on
}

// TokenERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenERC20TransactorRaw struct {
	Contract *TokenERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenERC20 creates a new instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20(address common.Address, backend bind.ContractBackend) (*TokenERC20, error) {
	contract, err := bindTokenERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenERC20{TokenERC20Caller: TokenERC20Caller{contract: contract}, TokenERC20Transactor: TokenERC20Transactor{contract: contract}, TokenERC20Filterer: TokenERC20Filterer{contract: contract}}, nil
}

// NewTokenERC20Caller creates a new read-only instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20Caller(address common.Address, caller bind.ContractCaller) (*TokenERC20Caller, error) {
	contract, err := bindTokenERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenERC20Caller{contract: contract}, nil
}

// NewTokenERC20Transactor creates a new write-only instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*TokenERC20Transactor, error) {
	contract, err := bindTokenERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenERC20Transactor{contract: contract}, nil
}

// NewTokenERC20Filterer creates a new log filterer instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*TokenERC20Filterer, error) {
	contract, err := bindTokenERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenERC20Filterer{contract: contract}, nil
}

// bindTokenERC20 binds a generic wrapper to an already deployed contract.
func bindTokenERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenERC20 *TokenERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenERC20.Contract.TokenERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenERC20 *TokenERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC20.Contract.TokenERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenERC20 *TokenERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenERC20.Contract.TokenERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenERC20 *TokenERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenERC20 *TokenERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenERC20 *TokenERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenERC20 *TokenERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenERC20 *TokenERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenERC20.Contract.Allowance(&_TokenERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenERC20 *TokenERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenERC20.Contract.Allowance(&_TokenERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenERC20 *TokenERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenERC20 *TokenERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _TokenERC20.Contract.BalanceOf(&_TokenERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenERC20 *TokenERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TokenERC20.Contract.BalanceOf(&_TokenERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenERC20 *TokenERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenERC20 *TokenERC20Session) Decimals() (uint8, error) {
	return _TokenERC20.Contract.Decimals(&_TokenERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenERC20 *TokenERC20CallerSession) Decimals() (uint8, error) {
	return _TokenERC20.Contract.Decimals(&_TokenERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenERC20 *TokenERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenERC20 *TokenERC20Session) Name() (string, error) {
	return _TokenERC20.Contract.Name(&_TokenERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenERC20 *TokenERC20CallerSession) Name() (string, error) {
	return _TokenERC20.Contract.Name(&_TokenERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenERC20 *TokenERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenERC20 *TokenERC20Session) Symbol() (string, error) {
	return _TokenERC20.Contract.Symbol(&_TokenERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenERC20 *TokenERC20CallerSession) Symbol() (string, error) {
	return _TokenERC20.Contract.Symbol(&_TokenERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenERC20 *TokenERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenERC20 *TokenERC20Session) TotalSupply() (*big.Int, error) {
	return _TokenERC20.Contract.TotalSupply(&_TokenERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenERC20 *TokenERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _TokenERC20.Contract.TotalSupply(&_TokenERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Approve(&_TokenERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Approve(&_TokenERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenERC20 *TokenERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenERC20 *TokenERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.DecreaseAllowance(&_TokenERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenERC20 *TokenERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.DecreaseAllowance(&_TokenERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenERC20 *TokenERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenERC20 *TokenERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.IncreaseAllowance(&_TokenERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenERC20 *TokenERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.IncreaseAllowance(&_TokenERC20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TokenERC20 *TokenERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TokenERC20 *TokenERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Mint(&_TokenERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TokenERC20 *TokenERC20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Mint(&_TokenERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Transfer(&_TokenERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Transfer(&_TokenERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.TransferFrom(&_TokenERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TokenERC20 *TokenERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.TransferFrom(&_TokenERC20.TransactOpts, from, to, amount)
}

// TokenERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TokenERC20 contract.
type TokenERC20ApprovalIterator struct {
	Event *TokenERC20Approval // Event containing the contract specifics and raw log

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
func (it *TokenERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC20Approval)
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
		it.Event = new(TokenERC20Approval)
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
func (it *TokenERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC20Approval represents a Approval event raised by the TokenERC20 contract.
type TokenERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenERC20 *TokenERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC20ApprovalIterator{contract: _TokenERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenERC20 *TokenERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC20Approval)
				if err := _TokenERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenERC20 *TokenERC20Filterer) ParseApproval(log types.Log) (*TokenERC20Approval, error) {
	event := new(TokenERC20Approval)
	if err := _TokenERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenERC20 contract.
type TokenERC20TransferIterator struct {
	Event *TokenERC20Transfer // Event containing the contract specifics and raw log

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
func (it *TokenERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC20Transfer)
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
		it.Event = new(TokenERC20Transfer)
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
func (it *TokenERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC20Transfer represents a Transfer event raised by the TokenERC20 contract.
type TokenERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenERC20 *TokenERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC20TransferIterator{contract: _TokenERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenERC20 *TokenERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC20Transfer)
				if err := _TokenERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenERC20 *TokenERC20Filterer) ParseTransfer(log types.Log) (*TokenERC20Transfer, error) {
	event := new(TokenERC20Transfer)
	if err := _TokenERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
