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

// BridgePolygonMetaData contains all meta data concerning the BridgePolygon contract.
var BridgePolygonMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBridgeBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDigest\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isNonceProcessed\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseBridge\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processedNonces\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseTokens\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseBridge\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BridgePaused\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeUnpaused\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensReleased\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BridgeIsPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonceAlreadyProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
	Bin: "0x6101a0604052348015610010575f5ffd5b506040516126dd3803806126dd83398181016040528101906100329190610502565b6040518060400160405280601081526020017f43726f7373436861696e427269646765000000000000000000000000000000008152506040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506100ba6100af6102c760201b60201c565b6102ce60201b60201c565b6100ce60018361038f60201b90919060201c565b61012081815250506100ea60028261038f60201b90919060201c565b6101408181525050818051906020012060e08181525050808051906020012061010081815250504660a081815250506101276103dc60201b60201c565b608081815250503073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff168152505050505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036101c9576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361022e576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166101608173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff166101808173ffffffffffffffffffffffffffffffffffffffff16815250505f60045f6101000a81548160ff0219169083151502179055506102c0336102ce60201b60201c565b50506109ec565b5f33905090565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f6020835110156103b0576103a98361043660201b60201c565b90506103d6565b826103c08361049b60201b60201c565b5f0190816103ce919061077d565b5060ff5f1b90505b92915050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60e05161010051463060405160200161041b959493929190610882565b60405160208183030381529060405280519060200120905090565b5f5f829050601f8151111561048257826040517f305a27a90000000000000000000000000000000000000000000000000000000081526004016104799190610939565b60405180910390fd5b80518161048e90610986565b5f1c175f1b915050919050565b5f819050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104d1826104a8565b9050919050565b6104e1816104c7565b81146104eb575f5ffd5b50565b5f815190506104fc816104d8565b92915050565b5f5f60408385031215610518576105176104a4565b5b5f610525858286016104ee565b9250506020610536858286016104ee565b9150509250929050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806105bb57607f821691505b6020821081036105ce576105cd610577565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026106307fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826105f5565b61063a86836105f5565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61067e61067961067484610652565b61065b565b610652565b9050919050565b5f819050919050565b61069783610664565b6106ab6106a382610685565b848454610601565b825550505050565b5f5f905090565b6106c26106b3565b6106cd81848461068e565b505050565b5b818110156106f0576106e55f826106ba565b6001810190506106d3565b5050565b601f82111561073557610706816105d4565b61070f846105e6565b8101602085101561071e578190505b61073261072a856105e6565b8301826106d2565b50505b505050565b5f82821c905092915050565b5f6107555f198460080261073a565b1980831691505092915050565b5f61076d8383610746565b9150826002028217905092915050565b61078682610540565b67ffffffffffffffff81111561079f5761079e61054a565b5b6107a982546105a4565b6107b48282856106f4565b5f60209050601f8311600181146107e5575f84156107d3578287015190505b6107dd8582610762565b865550610844565b601f1984166107f3866105d4565b5f5b8281101561081a578489015182556001820191506020850194506020810190506107f5565b868310156108375784890151610833601f891682610746565b8355505b6001600288020188555050505b505050505050565b5f819050919050565b61085e8161084c565b82525050565b61086d81610652565b82525050565b61087c816104c7565b82525050565b5f60a0820190506108955f830188610855565b6108a26020830187610855565b6108af6040830186610855565b6108bc6060830185610864565b6108c96080830184610873565b9695505050505050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61090b82610540565b61091581856108d3565b93506109258185602086016108e3565b61092e816108f1565b840191505092915050565b5f6020820190508181035f8301526109518184610901565b905092915050565b5f81519050919050565b5f819050602082019050919050565b5f61097d825161084c565b80915050919050565b5f61099082610959565b8261099a84610963565b90506109a581610972565b925060208210156109e5576109e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff836020036008026105f5565b831692505b5050919050565b60805160a05160c05160e0516101005161012051610140516101605161018051611c77610a665f395f81816102e001526107c901525f8181610878015281816109b00152610acf01525f6103f801525f6103c401525f610ecf01525f610eae01525f610cf901525f610d4f01525f610d780152611c775ff3fe608060405234801561000f575f5ffd5b50600436106100fe575f3560e01c8063a82f143c11610095578063ef63058711610064578063ef6305871461026a578063f0bc86de14610286578063f2fde38b146102a4578063fc0c546a146102c0576100fe565b8063a82f143c146101f4578063b187bd26146101fe578063b8d06bc41461021c578063ed24911d1461024c576100fe565b80637dd0480f116100d15780637dd0480f1461017857806384b0196e146101825780638da5cb5b146101a657806391dae519146101c4576100fe565b8063238ac933146101025780635c975abb146101205780636b3f4b3f1461013e578063715018a61461016e575b5f5ffd5b61010a6102de565b6040516101179190611235565b60405180910390f35b610128610302565b6040516101359190611268565b60405180910390f35b610158600480360381019061015391906112bc565b610314565b6040516101659190611268565b60405180910390f35b61017661033a565b005b61018061034d565b005b61018a6103b4565b60405161019d979695949392919061146f565b60405180910390f35b6101ae6104b1565b6040516101bb9190611235565b60405180910390f35b6101de60048036038101906101d991906112bc565b6104d8565b6040516101eb9190611268565b60405180910390f35b6101fc6104f5565b005b61020661055b565b6040516102139190611268565b60405180910390f35b6102366004803603810190610231919061151b565b610570565b604051610243919061156b565b60405180910390f35b6102546105d4565b604051610261919061156b565b60405180910390f35b610284600480360381019061027f91906115e5565b6105e2565b005b61028e6109ad565b60405161029b9190611669565b60405180910390f35b6102be60048036038101906102b99190611682565b610a4b565b005b6102c8610acd565b6040516102d59190611708565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b60045f9054906101000a900460ff1681565b5f60035f8381526020019081526020015f205f9054906101000a900460ff169050919050565b610342610af1565b61034b5f610b6f565b565b610355610af1565b600160045f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f80c6ef1d6f41faaf47f10c6da1ae6d45d4773d53d4c79bc1ef2737b10dbed57960405160405180910390a2565b5f6060805f5f5f60606103f160017f0000000000000000000000000000000000000000000000000000000000000000610c3090919063ffffffff16565b61042560027f0000000000000000000000000000000000000000000000000000000000000000610c3090919063ffffffff16565b46305f5f1b5f67ffffffffffffffff81111561044457610443611721565b5b6040519080825280602002602001820160405280156104725781602001602082028036833780820191505090505b507f0f00000000000000000000000000000000000000000000000000000000000000959493929190965096509650965096509650965090919293949596565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6003602052805f5260405f205f915054906101000a900460ff1681565b6104fd610af1565b5f60045f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167fa668d29cf81d3acc05e4a35a5bf898bb2c919fc7a4366e536bbd43522cb670f060405160405180910390a2565b5f60045f9054906101000a900460ff16905090565b5f5f7f629bf0bb24071f41b4accc58d7ed3feadb3149563f72fc06f490e4093aedf7868585856040516020016105a9949392919061174e565b6040516020818303038152906040528051906020012090506105ca81610cdd565b9150509392505050565b5f6105dd610cf6565b905090565b60045f9054906101000a900460ff1615610628576040517f7cf03c4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff160361068d576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f84036106c6576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035f8481526020019081526020015f205f9054906101000a900460ff161561071b576040517f752eea4f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f629bf0bb24071f41b4accc58d7ed3feadb3149563f72fc06f490e4093aedf786868686604051602001610753949392919061174e565b6040516020818303038152906040528051906020012090505f61077582610cdd565b90505f6107c58286868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050610dac565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461084c576040517f8baa579f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160035f8881526020019081526020015f205f6101000a81548160ff0219169083151502179055505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8a8a6040518363ffffffff1660e01b81526004016108d1929190611791565b6020604051808303815f875af11580156108ed573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061091191906117e2565b905080610953576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094a90611857565b60405180910390fd5b868973ffffffffffffffffffffffffffffffffffffffff167fc5c52c2a9175470464d5ea4429889e7df2ea88630a3d32f4d0d3d2d4486562108a60405161099a9190611669565b60405180910390a3505050505050505050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610a079190611235565b602060405180830381865afa158015610a22573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a469190611889565b905090565b610a53610af1565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ac1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ab890611924565b60405180910390fd5b610aca81610b6f565b50565b7f000000000000000000000000000000000000000000000000000000000000000081565b610af9610dd1565b73ffffffffffffffffffffffffffffffffffffffff16610b176104b1565b73ffffffffffffffffffffffffffffffffffffffff1614610b6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b649061198c565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b606060ff5f1b8314610c4c57610c4583610dd8565b9050610cd7565b818054610c58906119d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610c84906119d7565b8015610ccf5780601f10610ca657610100808354040283529160200191610ccf565b820191905f5260205f20905b815481529060010190602001808311610cb257829003601f168201915b505050505090505b92915050565b5f610cef610ce9610cf6565b83610e4a565b9050919050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16148015610d7157507f000000000000000000000000000000000000000000000000000000000000000046145b15610d9e577f00000000000000000000000000000000000000000000000000000000000000009050610da9565b610da6610e8a565b90505b90565b5f5f5f610db98585610f1f565b91509150610dc681610f6b565b819250505092915050565b5f33905090565b60605f610de4836110d0565b90505f602067ffffffffffffffff811115610e0257610e01611721565b5b6040519080825280601f01601f191660200182016040528015610e345781602001600182028036833780820191505090505b5090508181528360208201528092505050919050565b5f6040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000004630604051602001610f04959493929190611a07565b60405160208183030381529060405280519060200120905090565b5f5f6041835103610f5c575f5f5f602086015192506040860151915060608601515f1a9050610f508782858561111e565b94509450505050610f64565b5f6002915091505b9250929050565b5f6004811115610f7e57610f7d611a58565b5b816004811115610f9157610f90611a58565b5b03156110cd5760016004811115610fab57610faa611a58565b5b816004811115610fbe57610fbd611a58565b5b03610ffe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff590611acf565b60405180910390fd5b6002600481111561101257611011611a58565b5b81600481111561102557611024611a58565b5b03611065576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161105c90611b37565b60405180910390fd5b6003600481111561107957611078611a58565b5b81600481111561108c5761108b611a58565b5b036110cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110c390611bc5565b60405180910390fd5b5b50565b5f5f60ff835f1c169050601f811115611115576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80915050919050565b5f5f7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0835f1c1115611156575f6003915091506111ed565b5f6001878787876040515f81526020016040526040516111799493929190611bfe565b6020604051602081039080840390855afa158015611199573d5f5f3e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036111e5575f600192509250506111ed565b805f92509250505b94509492505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61121f826111f6565b9050919050565b61122f81611215565b82525050565b5f6020820190506112485f830184611226565b92915050565b5f8115159050919050565b6112628161124e565b82525050565b5f60208201905061127b5f830184611259565b92915050565b5f5ffd5b5f5ffd5b5f819050919050565b61129b81611289565b81146112a5575f5ffd5b50565b5f813590506112b681611292565b92915050565b5f602082840312156112d1576112d0611281565b5b5f6112de848285016112a8565b91505092915050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b61131b816112e7565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61136382611321565b61136d818561132b565b935061137d81856020860161133b565b61138681611349565b840191505092915050565b61139a81611289565b82525050565b5f819050919050565b6113b2816113a0565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6113ea81611289565b82525050565b5f6113fb83836113e1565b60208301905092915050565b5f602082019050919050565b5f61141d826113b8565b61142781856113c2565b9350611432836113d2565b805f5b8381101561146257815161144988826113f0565b975061145483611407565b925050600181019050611435565b5085935050505092915050565b5f60e0820190506114825f83018a611312565b81810360208301526114948189611359565b905081810360408301526114a88188611359565b90506114b76060830187611391565b6114c46080830186611226565b6114d160a08301856113a9565b81810360c08301526114e38184611413565b905098975050505050505050565b6114fa81611215565b8114611504575f5ffd5b50565b5f81359050611515816114f1565b92915050565b5f5f5f6060848603121561153257611531611281565b5b5f61153f86828701611507565b9350506020611550868287016112a8565b9250506040611561868287016112a8565b9150509250925092565b5f60208201905061157e5f8301846113a9565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126115a5576115a4611584565b5b8235905067ffffffffffffffff8111156115c2576115c1611588565b5b6020830191508360018202830111156115de576115dd61158c565b5b9250929050565b5f5f5f5f5f608086880312156115fe576115fd611281565b5b5f61160b88828901611507565b955050602061161c888289016112a8565b945050604061162d888289016112a8565b935050606086013567ffffffffffffffff81111561164e5761164d611285565b5b61165a88828901611590565b92509250509295509295909350565b5f60208201905061167c5f830184611391565b92915050565b5f6020828403121561169757611696611281565b5b5f6116a484828501611507565b91505092915050565b5f819050919050565b5f6116d06116cb6116c6846111f6565b6116ad565b6111f6565b9050919050565b5f6116e1826116b6565b9050919050565b5f6116f2826116d7565b9050919050565b611702816116e8565b82525050565b5f60208201905061171b5f8301846116f9565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f6080820190506117615f8301876113a9565b61176e6020830186611226565b61177b6040830185611391565b6117886060830184611391565b95945050505050565b5f6040820190506117a45f830185611226565b6117b16020830184611391565b9392505050565b6117c18161124e565b81146117cb575f5ffd5b50565b5f815190506117dc816117b8565b92915050565b5f602082840312156117f7576117f6611281565b5b5f611804848285016117ce565b91505092915050565b7f546f6b656e207472616e73666572206661696c656400000000000000000000005f82015250565b5f61184160158361132b565b915061184c8261180d565b602082019050919050565b5f6020820190508181035f83015261186e81611835565b9050919050565b5f8151905061188381611292565b92915050565b5f6020828403121561189e5761189d611281565b5b5f6118ab84828501611875565b91505092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f61190e60268361132b565b9150611919826118b4565b604082019050919050565b5f6020820190508181035f83015261193b81611902565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f61197660208361132b565b915061198182611942565b602082019050919050565b5f6020820190508181035f8301526119a38161196a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806119ee57607f821691505b602082108103611a0157611a006119aa565b5b50919050565b5f60a082019050611a1a5f8301886113a9565b611a2760208301876113a9565b611a3460408301866113a9565b611a416060830185611391565b611a4e6080830184611226565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f45434453413a20696e76616c6964207369676e617475726500000000000000005f82015250565b5f611ab960188361132b565b9150611ac482611a85565b602082019050919050565b5f6020820190508181035f830152611ae681611aad565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e677468005f82015250565b5f611b21601f8361132b565b9150611b2c82611aed565b602082019050919050565b5f6020820190508181035f830152611b4e81611b15565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c5f8201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b5f611baf60228361132b565b9150611bba82611b55565b604082019050919050565b5f6020820190508181035f830152611bdc81611ba3565b9050919050565b5f60ff82169050919050565b611bf881611be3565b82525050565b5f608082019050611c115f8301876113a9565b611c1e6020830186611bef565b611c2b60408301856113a9565b611c3860608301846113a9565b9594505050505056fea26469706673582212200cd9de450fc137a5dba1fa4dacee5c12aa988da0ddfc75e5ce07ac70d18c169764736f6c634300081e0033",
}

// BridgePolygonABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgePolygonMetaData.ABI instead.
var BridgePolygonABI = BridgePolygonMetaData.ABI

// BridgePolygonBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgePolygonMetaData.Bin instead.
var BridgePolygonBin = BridgePolygonMetaData.Bin

// DeployBridgePolygon deploys a new Ethereum contract, binding an instance of BridgePolygon to it.
func DeployBridgePolygon(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _signer common.Address) (common.Address, *types.Transaction, *BridgePolygon, error) {
	parsed, err := BridgePolygonMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgePolygonBin), backend, _token, _signer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgePolygon{BridgePolygonCaller: BridgePolygonCaller{contract: contract}, BridgePolygonTransactor: BridgePolygonTransactor{contract: contract}, BridgePolygonFilterer: BridgePolygonFilterer{contract: contract}}, nil
}

// BridgePolygon is an auto generated Go binding around an Ethereum contract.
type BridgePolygon struct {
	BridgePolygonCaller     // Read-only binding to the contract
	BridgePolygonTransactor // Write-only binding to the contract
	BridgePolygonFilterer   // Log filterer for contract events
}

// BridgePolygonCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgePolygonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgePolygonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgePolygonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgePolygonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgePolygonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgePolygonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgePolygonSession struct {
	Contract     *BridgePolygon    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgePolygonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgePolygonCallerSession struct {
	Contract *BridgePolygonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BridgePolygonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgePolygonTransactorSession struct {
	Contract     *BridgePolygonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BridgePolygonRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgePolygonRaw struct {
	Contract *BridgePolygon // Generic contract binding to access the raw methods on
}

// BridgePolygonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgePolygonCallerRaw struct {
	Contract *BridgePolygonCaller // Generic read-only contract binding to access the raw methods on
}

// BridgePolygonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgePolygonTransactorRaw struct {
	Contract *BridgePolygonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgePolygon creates a new instance of BridgePolygon, bound to a specific deployed contract.
func NewBridgePolygon(address common.Address, backend bind.ContractBackend) (*BridgePolygon, error) {
	contract, err := bindBridgePolygon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgePolygon{BridgePolygonCaller: BridgePolygonCaller{contract: contract}, BridgePolygonTransactor: BridgePolygonTransactor{contract: contract}, BridgePolygonFilterer: BridgePolygonFilterer{contract: contract}}, nil
}

// NewBridgePolygonCaller creates a new read-only instance of BridgePolygon, bound to a specific deployed contract.
func NewBridgePolygonCaller(address common.Address, caller bind.ContractCaller) (*BridgePolygonCaller, error) {
	contract, err := bindBridgePolygon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgePolygonCaller{contract: contract}, nil
}

// NewBridgePolygonTransactor creates a new write-only instance of BridgePolygon, bound to a specific deployed contract.
func NewBridgePolygonTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgePolygonTransactor, error) {
	contract, err := bindBridgePolygon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgePolygonTransactor{contract: contract}, nil
}

// NewBridgePolygonFilterer creates a new log filterer instance of BridgePolygon, bound to a specific deployed contract.
func NewBridgePolygonFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgePolygonFilterer, error) {
	contract, err := bindBridgePolygon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgePolygonFilterer{contract: contract}, nil
}

// bindBridgePolygon binds a generic wrapper to an already deployed contract.
func bindBridgePolygon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgePolygonMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgePolygon *BridgePolygonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgePolygon.Contract.BridgePolygonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgePolygon *BridgePolygonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgePolygon.Contract.BridgePolygonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgePolygon *BridgePolygonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgePolygon.Contract.BridgePolygonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgePolygon *BridgePolygonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgePolygon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgePolygon *BridgePolygonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgePolygon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgePolygon *BridgePolygonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgePolygon.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BridgePolygon *BridgePolygonCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BridgePolygon.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BridgePolygon *BridgePolygonSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BridgePolygon.Contract.Eip712Domain(&_BridgePolygon.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BridgePolygon *BridgePolygonCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BridgePolygon.Contract.Eip712Domain(&_BridgePolygon.CallOpts)
}

// GetBridgeBalance is a free data retrieval call binding the contract method 0xf0bc86de.
//
// Solidity: function getBridgeBalance() view returns(uint256)
func (_BridgePolygon *BridgePolygonCaller) GetBridgeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgePolygon.contract.Call(opts, &out, "getBridgeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBridgeBalance is a free data retrieval call binding the contract method 0xf0bc86de.
//
// Solidity: function getBridgeBalance() view returns(uint256)
func (_BridgePolygon *BridgePolygonSession) GetBridgeBalance() (*big.Int, error) {
	return _BridgePolygon.Contract.GetBridgeBalance(&_BridgePolygon.CallOpts)
}

// GetBridgeBalance is a free data retrieval call binding the contract method 0xf0bc86de.
//
// Solidity: function getBridgeBalance() view returns(uint256)
func (_BridgePolygon *BridgePolygonCallerSession) GetBridgeBalance() (*big.Int, error) {
	return _BridgePolygon.Contract.GetBridgeBalance(&_BridgePolygon.CallOpts)
}

// GetDigest is a free data retrieval call binding the contract method 0xb8d06bc4.
//
// Solidity: function getDigest(address user, uint256 amount, uint256 nonce) view returns(bytes32)
func (_BridgePolygon *BridgePolygonCaller) GetDigest(opts *bind.CallOpts, user common.Address, amount *big.Int, nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgePolygon.contract.Call(opts, &out, "getDigest", user, amount, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDigest is a free data retrieval call binding the contract method 0xb8d06bc4.
//
// Solidity: function getDigest(address user, uint256 amount, uint256 nonce) view returns(bytes32)
func (_BridgePolygon *BridgePolygonSession) GetDigest(user common.Address, amount *big.Int, nonce *big.Int) ([32]byte, error) {
	return _BridgePolygon.Contract.GetDigest(&_BridgePolygon.CallOpts, user, amount, nonce)
}

// GetDigest is a free data retrieval call binding the contract method 0xb8d06bc4.
//
// Solidity: function getDigest(address user, uint256 amount, uint256 nonce) view returns(bytes32)
func (_BridgePolygon *BridgePolygonCallerSession) GetDigest(user common.Address, amount *big.Int, nonce *big.Int) ([32]byte, error) {
	return _BridgePolygon.Contract.GetDigest(&_BridgePolygon.CallOpts, user, amount, nonce)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BridgePolygon *BridgePolygonCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgePolygon.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BridgePolygon *BridgePolygonSession) GetDomainSeparator() ([32]byte, error) {
	return _BridgePolygon.Contract.GetDomainSeparator(&_BridgePolygon.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BridgePolygon *BridgePolygonCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _BridgePolygon.Contract.GetDomainSeparator(&_BridgePolygon.CallOpts)
}

// IsNonceProcessed is a free data retrieval call binding the contract method 0x6b3f4b3f.
//
// Solidity: function isNonceProcessed(uint256 nonce) view returns(bool)
func (_BridgePolygon *BridgePolygonCaller) IsNonceProcessed(opts *bind.CallOpts, nonce *big.Int) (bool, error) {
	var out []interface{}
	err := _BridgePolygon.contract.Call(opts, &out, "isNonceProcessed", nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNonceProcessed is a free data retrieval call binding the contract method 0x6b3f4b3f.
//
// Solidity: function isNonceProcessed(uint256 nonce) view returns(bool)
func (_BridgePolygon *BridgePolygonSession) IsNonceProcessed(nonce *big.Int) (bool, error) {
	return _BridgePolygon.Contract.IsNonceProcessed(&_BridgePolygon.CallOpts, nonce)
}

// IsNonceProcessed is a free data retrieval call binding the contract method 0x6b3f4b3f.
//
// Solidity: function isNonceProcessed(uint256 nonce) view returns(bool)
func (_BridgePolygon *BridgePolygonCallerSession) IsNonceProcessed(nonce *big.Int) (bool, error) {
	return _BridgePolygon.Contract.IsNonceProcessed(&_BridgePolygon.CallOpts, nonce)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_BridgePolygon *BridgePolygonCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgePolygon.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_BridgePolygon *BridgePolygonSession) IsPaused() (bool, error) {
	return _BridgePolygon.Contract.IsPaused(&_BridgePolygon.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_BridgePolygon *BridgePolygonCallerSession) IsPaused() (bool, error) {
	return _BridgePolygon.Contract.IsPaused(&_BridgePolygon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgePolygon *BridgePolygonCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgePolygon.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgePolygon *BridgePolygonSession) Owner() (common.Address, error) {
	return _BridgePolygon.Contract.Owner(&_BridgePolygon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgePolygon *BridgePolygonCallerSession) Owner() (common.Address, error) {
	return _BridgePolygon.Contract.Owner(&_BridgePolygon.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgePolygon *BridgePolygonCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgePolygon.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgePolygon *BridgePolygonSession) Paused() (bool, error) {
	return _BridgePolygon.Contract.Paused(&_BridgePolygon.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgePolygon *BridgePolygonCallerSession) Paused() (bool, error) {
	return _BridgePolygon.Contract.Paused(&_BridgePolygon.CallOpts)
}

// ProcessedNonces is a free data retrieval call binding the contract method 0x91dae519.
//
// Solidity: function processedNonces(uint256 ) view returns(bool)
func (_BridgePolygon *BridgePolygonCaller) ProcessedNonces(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _BridgePolygon.contract.Call(opts, &out, "processedNonces", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedNonces is a free data retrieval call binding the contract method 0x91dae519.
//
// Solidity: function processedNonces(uint256 ) view returns(bool)
func (_BridgePolygon *BridgePolygonSession) ProcessedNonces(arg0 *big.Int) (bool, error) {
	return _BridgePolygon.Contract.ProcessedNonces(&_BridgePolygon.CallOpts, arg0)
}

// ProcessedNonces is a free data retrieval call binding the contract method 0x91dae519.
//
// Solidity: function processedNonces(uint256 ) view returns(bool)
func (_BridgePolygon *BridgePolygonCallerSession) ProcessedNonces(arg0 *big.Int) (bool, error) {
	return _BridgePolygon.Contract.ProcessedNonces(&_BridgePolygon.CallOpts, arg0)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_BridgePolygon *BridgePolygonCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgePolygon.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_BridgePolygon *BridgePolygonSession) Signer() (common.Address, error) {
	return _BridgePolygon.Contract.Signer(&_BridgePolygon.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_BridgePolygon *BridgePolygonCallerSession) Signer() (common.Address, error) {
	return _BridgePolygon.Contract.Signer(&_BridgePolygon.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgePolygon *BridgePolygonCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgePolygon.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgePolygon *BridgePolygonSession) Token() (common.Address, error) {
	return _BridgePolygon.Contract.Token(&_BridgePolygon.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgePolygon *BridgePolygonCallerSession) Token() (common.Address, error) {
	return _BridgePolygon.Contract.Token(&_BridgePolygon.CallOpts)
}

// PauseBridge is a paid mutator transaction binding the contract method 0x7dd0480f.
//
// Solidity: function pauseBridge() returns()
func (_BridgePolygon *BridgePolygonTransactor) PauseBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgePolygon.contract.Transact(opts, "pauseBridge")
}

// PauseBridge is a paid mutator transaction binding the contract method 0x7dd0480f.
//
// Solidity: function pauseBridge() returns()
func (_BridgePolygon *BridgePolygonSession) PauseBridge() (*types.Transaction, error) {
	return _BridgePolygon.Contract.PauseBridge(&_BridgePolygon.TransactOpts)
}

// PauseBridge is a paid mutator transaction binding the contract method 0x7dd0480f.
//
// Solidity: function pauseBridge() returns()
func (_BridgePolygon *BridgePolygonTransactorSession) PauseBridge() (*types.Transaction, error) {
	return _BridgePolygon.Contract.PauseBridge(&_BridgePolygon.TransactOpts)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0xef630587.
//
// Solidity: function releaseTokens(address user, uint256 amount, uint256 nonce, bytes signature) returns()
func (_BridgePolygon *BridgePolygonTransactor) ReleaseTokens(opts *bind.TransactOpts, user common.Address, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _BridgePolygon.contract.Transact(opts, "releaseTokens", user, amount, nonce, signature)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0xef630587.
//
// Solidity: function releaseTokens(address user, uint256 amount, uint256 nonce, bytes signature) returns()
func (_BridgePolygon *BridgePolygonSession) ReleaseTokens(user common.Address, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _BridgePolygon.Contract.ReleaseTokens(&_BridgePolygon.TransactOpts, user, amount, nonce, signature)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0xef630587.
//
// Solidity: function releaseTokens(address user, uint256 amount, uint256 nonce, bytes signature) returns()
func (_BridgePolygon *BridgePolygonTransactorSession) ReleaseTokens(user common.Address, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _BridgePolygon.Contract.ReleaseTokens(&_BridgePolygon.TransactOpts, user, amount, nonce, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgePolygon *BridgePolygonTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgePolygon.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgePolygon *BridgePolygonSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgePolygon.Contract.RenounceOwnership(&_BridgePolygon.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgePolygon *BridgePolygonTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgePolygon.Contract.RenounceOwnership(&_BridgePolygon.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgePolygon *BridgePolygonTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgePolygon.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgePolygon *BridgePolygonSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgePolygon.Contract.TransferOwnership(&_BridgePolygon.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgePolygon *BridgePolygonTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgePolygon.Contract.TransferOwnership(&_BridgePolygon.TransactOpts, newOwner)
}

// UnpauseBridge is a paid mutator transaction binding the contract method 0xa82f143c.
//
// Solidity: function unpauseBridge() returns()
func (_BridgePolygon *BridgePolygonTransactor) UnpauseBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgePolygon.contract.Transact(opts, "unpauseBridge")
}

// UnpauseBridge is a paid mutator transaction binding the contract method 0xa82f143c.
//
// Solidity: function unpauseBridge() returns()
func (_BridgePolygon *BridgePolygonSession) UnpauseBridge() (*types.Transaction, error) {
	return _BridgePolygon.Contract.UnpauseBridge(&_BridgePolygon.TransactOpts)
}

// UnpauseBridge is a paid mutator transaction binding the contract method 0xa82f143c.
//
// Solidity: function unpauseBridge() returns()
func (_BridgePolygon *BridgePolygonTransactorSession) UnpauseBridge() (*types.Transaction, error) {
	return _BridgePolygon.Contract.UnpauseBridge(&_BridgePolygon.TransactOpts)
}

// BridgePolygonBridgePausedIterator is returned from FilterBridgePaused and is used to iterate over the raw logs and unpacked data for BridgePaused events raised by the BridgePolygon contract.
type BridgePolygonBridgePausedIterator struct {
	Event *BridgePolygonBridgePaused // Event containing the contract specifics and raw log

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
func (it *BridgePolygonBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePolygonBridgePaused)
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
		it.Event = new(BridgePolygonBridgePaused)
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
func (it *BridgePolygonBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePolygonBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePolygonBridgePaused represents a BridgePaused event raised by the BridgePolygon contract.
type BridgePolygonBridgePaused struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgePaused is a free log retrieval operation binding the contract event 0x80c6ef1d6f41faaf47f10c6da1ae6d45d4773d53d4c79bc1ef2737b10dbed579.
//
// Solidity: event BridgePaused(address indexed admin)
func (_BridgePolygon *BridgePolygonFilterer) FilterBridgePaused(opts *bind.FilterOpts, admin []common.Address) (*BridgePolygonBridgePausedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _BridgePolygon.contract.FilterLogs(opts, "BridgePaused", adminRule)
	if err != nil {
		return nil, err
	}
	return &BridgePolygonBridgePausedIterator{contract: _BridgePolygon.contract, event: "BridgePaused", logs: logs, sub: sub}, nil
}

// WatchBridgePaused is a free log subscription operation binding the contract event 0x80c6ef1d6f41faaf47f10c6da1ae6d45d4773d53d4c79bc1ef2737b10dbed579.
//
// Solidity: event BridgePaused(address indexed admin)
func (_BridgePolygon *BridgePolygonFilterer) WatchBridgePaused(opts *bind.WatchOpts, sink chan<- *BridgePolygonBridgePaused, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _BridgePolygon.contract.WatchLogs(opts, "BridgePaused", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePolygonBridgePaused)
				if err := _BridgePolygon.contract.UnpackLog(event, "BridgePaused", log); err != nil {
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
func (_BridgePolygon *BridgePolygonFilterer) ParseBridgePaused(log types.Log) (*BridgePolygonBridgePaused, error) {
	event := new(BridgePolygonBridgePaused)
	if err := _BridgePolygon.contract.UnpackLog(event, "BridgePaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePolygonBridgeUnpausedIterator is returned from FilterBridgeUnpaused and is used to iterate over the raw logs and unpacked data for BridgeUnpaused events raised by the BridgePolygon contract.
type BridgePolygonBridgeUnpausedIterator struct {
	Event *BridgePolygonBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgePolygonBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePolygonBridgeUnpaused)
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
		it.Event = new(BridgePolygonBridgeUnpaused)
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
func (it *BridgePolygonBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePolygonBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePolygonBridgeUnpaused represents a BridgeUnpaused event raised by the BridgePolygon contract.
type BridgePolygonBridgeUnpaused struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgeUnpaused is a free log retrieval operation binding the contract event 0xa668d29cf81d3acc05e4a35a5bf898bb2c919fc7a4366e536bbd43522cb670f0.
//
// Solidity: event BridgeUnpaused(address indexed admin)
func (_BridgePolygon *BridgePolygonFilterer) FilterBridgeUnpaused(opts *bind.FilterOpts, admin []common.Address) (*BridgePolygonBridgeUnpausedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _BridgePolygon.contract.FilterLogs(opts, "BridgeUnpaused", adminRule)
	if err != nil {
		return nil, err
	}
	return &BridgePolygonBridgeUnpausedIterator{contract: _BridgePolygon.contract, event: "BridgeUnpaused", logs: logs, sub: sub}, nil
}

// WatchBridgeUnpaused is a free log subscription operation binding the contract event 0xa668d29cf81d3acc05e4a35a5bf898bb2c919fc7a4366e536bbd43522cb670f0.
//
// Solidity: event BridgeUnpaused(address indexed admin)
func (_BridgePolygon *BridgePolygonFilterer) WatchBridgeUnpaused(opts *bind.WatchOpts, sink chan<- *BridgePolygonBridgeUnpaused, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _BridgePolygon.contract.WatchLogs(opts, "BridgeUnpaused", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePolygonBridgeUnpaused)
				if err := _BridgePolygon.contract.UnpackLog(event, "BridgeUnpaused", log); err != nil {
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
func (_BridgePolygon *BridgePolygonFilterer) ParseBridgeUnpaused(log types.Log) (*BridgePolygonBridgeUnpaused, error) {
	event := new(BridgePolygonBridgeUnpaused)
	if err := _BridgePolygon.contract.UnpackLog(event, "BridgeUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePolygonEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BridgePolygon contract.
type BridgePolygonEIP712DomainChangedIterator struct {
	Event *BridgePolygonEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BridgePolygonEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePolygonEIP712DomainChanged)
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
		it.Event = new(BridgePolygonEIP712DomainChanged)
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
func (it *BridgePolygonEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePolygonEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePolygonEIP712DomainChanged represents a EIP712DomainChanged event raised by the BridgePolygon contract.
type BridgePolygonEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BridgePolygon *BridgePolygonFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BridgePolygonEIP712DomainChangedIterator, error) {

	logs, sub, err := _BridgePolygon.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BridgePolygonEIP712DomainChangedIterator{contract: _BridgePolygon.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BridgePolygon *BridgePolygonFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BridgePolygonEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BridgePolygon.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePolygonEIP712DomainChanged)
				if err := _BridgePolygon.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BridgePolygon *BridgePolygonFilterer) ParseEIP712DomainChanged(log types.Log) (*BridgePolygonEIP712DomainChanged, error) {
	event := new(BridgePolygonEIP712DomainChanged)
	if err := _BridgePolygon.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePolygonOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgePolygon contract.
type BridgePolygonOwnershipTransferredIterator struct {
	Event *BridgePolygonOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgePolygonOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePolygonOwnershipTransferred)
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
		it.Event = new(BridgePolygonOwnershipTransferred)
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
func (it *BridgePolygonOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePolygonOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePolygonOwnershipTransferred represents a OwnershipTransferred event raised by the BridgePolygon contract.
type BridgePolygonOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgePolygon *BridgePolygonFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgePolygonOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgePolygon.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgePolygonOwnershipTransferredIterator{contract: _BridgePolygon.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgePolygon *BridgePolygonFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgePolygonOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgePolygon.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePolygonOwnershipTransferred)
				if err := _BridgePolygon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgePolygon *BridgePolygonFilterer) ParseOwnershipTransferred(log types.Log) (*BridgePolygonOwnershipTransferred, error) {
	event := new(BridgePolygonOwnershipTransferred)
	if err := _BridgePolygon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePolygonTokensReleasedIterator is returned from FilterTokensReleased and is used to iterate over the raw logs and unpacked data for TokensReleased events raised by the BridgePolygon contract.
type BridgePolygonTokensReleasedIterator struct {
	Event *BridgePolygonTokensReleased // Event containing the contract specifics and raw log

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
func (it *BridgePolygonTokensReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePolygonTokensReleased)
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
		it.Event = new(BridgePolygonTokensReleased)
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
func (it *BridgePolygonTokensReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePolygonTokensReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePolygonTokensReleased represents a TokensReleased event raised by the BridgePolygon contract.
type BridgePolygonTokensReleased struct {
	User   common.Address
	Amount *big.Int
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensReleased is a free log retrieval operation binding the contract event 0xc5c52c2a9175470464d5ea4429889e7df2ea88630a3d32f4d0d3d2d448656210.
//
// Solidity: event TokensReleased(address indexed user, uint256 amount, uint256 indexed nonce)
func (_BridgePolygon *BridgePolygonFilterer) FilterTokensReleased(opts *bind.FilterOpts, user []common.Address, nonce []*big.Int) (*BridgePolygonTokensReleasedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _BridgePolygon.contract.FilterLogs(opts, "TokensReleased", userRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &BridgePolygonTokensReleasedIterator{contract: _BridgePolygon.contract, event: "TokensReleased", logs: logs, sub: sub}, nil
}

// WatchTokensReleased is a free log subscription operation binding the contract event 0xc5c52c2a9175470464d5ea4429889e7df2ea88630a3d32f4d0d3d2d448656210.
//
// Solidity: event TokensReleased(address indexed user, uint256 amount, uint256 indexed nonce)
func (_BridgePolygon *BridgePolygonFilterer) WatchTokensReleased(opts *bind.WatchOpts, sink chan<- *BridgePolygonTokensReleased, user []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _BridgePolygon.contract.WatchLogs(opts, "TokensReleased", userRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePolygonTokensReleased)
				if err := _BridgePolygon.contract.UnpackLog(event, "TokensReleased", log); err != nil {
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

// ParseTokensReleased is a log parse operation binding the contract event 0xc5c52c2a9175470464d5ea4429889e7df2ea88630a3d32f4d0d3d2d448656210.
//
// Solidity: event TokensReleased(address indexed user, uint256 amount, uint256 indexed nonce)
func (_BridgePolygon *BridgePolygonFilterer) ParseTokensReleased(log types.Log) (*BridgePolygonTokensReleased, error) {
	event := new(BridgePolygonTokensReleased)
	if err := _BridgePolygon.contract.UnpackLog(event, "TokensReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
