// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OrderMatching

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

// OrderMatchingLease is an auto generated low-level Go binding around an user-defined struct.
type OrderMatchingLease struct {
	ProviderAddress common.Address
	ProviderId      *big.Int
	AcceptedPrice   *big.Int
	StartBlock      *big.Int
	StartTime       *big.Int
	EndBlock        *big.Int
	EndTime         *big.Int
	State           uint8
}

// OrderMatchingMetaData contains all meta data concerning the OrderMatching contract.
var OrderMatchingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeProviderRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_authorizedAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAuthorizedAddress\",\"type\":\"address\"}],\"name\":\"AuthorizedAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"orderId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"orderId\",\"type\":\"uint64\"}],\"name\":\"OrderClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"orderId\",\"type\":\"uint64\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"orderId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptedPrice\",\"type\":\"uint256\"}],\"name\":\"OrderMatched\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"authorizedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bids\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"orderId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_orderId\",\"type\":\"uint64\"}],\"name\":\"closeOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_region\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"_uptime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_reputation\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_slashes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_token\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_spec\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_orderId\",\"type\":\"uint64\"}],\"name\":\"getOrderById\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"providerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"enumOrderMatching.State\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structOrderMatching.Lease\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddress\",\"type\":\"address\"}],\"name\":\"getOrderByProvider\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"leases\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"providerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"enumOrderMatching.State\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_orderId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_providerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_acceptedPrice\",\"type\":\"uint256\"}],\"name\":\"matchOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeProviderRegistry\",\"outputs\":[{\"internalType\":\"contractINodeProviderRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reputation\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"slashes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"enumOrderMatching.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"specs\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_orderId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_bidPrice\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAuthorizedAddress\",\"type\":\"address\"}],\"name\":\"updateAuthorizedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260015f806101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550348015610037575f80fd5b50604051612ccc380380612ccc8339818101604052810190610059919061017e565b3360055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506101bc565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61014d82610124565b9050919050565b61015d81610143565b8114610167575f80fd5b50565b5f8151905061017881610154565b92915050565b5f806040838503121561019457610193610120565b5b5f6101a18582860161016a565b92505060206101b28582860161016a565b9150509250929050565b612b03806101c95f395ff3fe608060405234801561000f575f80fd5b50600436106100cd575f3560e01c8063638d7dec1161008a5780638da5cb5b116100645780638da5cb5b14610216578063a469dffa14610234578063b930807d1461026f578063fd2e5a231461029f576100cd565b8063638d7dec146101ac5780636582972e146101dc5780637a544ad3146101fa576100cd565b806328d389bc146100d157806347d4f8f2146100ed5780634c7f3c791461011f5780634fd2edbb146101565780635539d4001461017257806359e7bf4b14610190575b5f80fd5b6100eb60048036038101906100e69190611862565b6102bb565b005b610107600480360381019061010291906118fd565b6103c4565b60405161011693929190611968565b60405180910390f35b6101396004803603810190610134919061199d565b610436565b60405161014d989796959493929190611a3b565b60405180910390f35b610170600480360381019061016b9190611ab7565b6104a5565b005b61017a61099c565b6040516101879190611b07565b60405180910390f35b6101aa60048036038101906101a59190611c5c565b6109c1565b005b6101c660048036038101906101c19190611862565b610bea565b6040516101d39190611e34565b60405180910390f35b6101e4610cb0565b6040516101f19190611eaf565b60405180910390f35b610214600480360381019061020f919061199d565b610cd5565b005b61021e61106d565b60405161022b9190611b07565b60405180910390f35b61024e6004803603810190610249919061199d565b611092565b6040516102669c9b9a99989796959493929190611f28565b60405180910390f35b6102896004803603810190610284919061199d565b611380565b60405161029691906120c9565b60405180910390f35b6102b960048036038101906102b491906118fd565b61148a565b005b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461034a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034190612153565b60405180910390fd5b8060065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe12f594f8990d16225e1238de64b5ef16fbe78fbb1a70f64e3a82274e9554c0a816040516103b99190611b07565b60405180910390a150565b6002602052815f5260405f2081815481106103dd575f80fd5b905f5260205f2090600302015f9150915050805f015f9054906101000a900467ffffffffffffffff1690806001015490806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b6003602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002015490806003015490806004015490806005015490806006015490806007015f9054906101000a900460ff16905088565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610534576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052b906121e1565b60405180910390fd5b5f60015f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16036105ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b190612249565b60405180910390fd5b5f60038111156105cd576105cc6119c8565b5b60015f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060050160149054906101000a900460ff166003811115610615576106146119c8565b5b14610655576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064c906122b1565b60405180910390fd5b5f8060075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166368efbce3856040518263ffffffff1660e01b81526004016106b19190611b07565b5f60405180830381865afa1580156106cb573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906106f3919061239a565b93505050915080610739576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107309061248a565b60405180910390fd5b5f6040518061010001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018581526020014381526020014281526020015f81526020015f81526020016001600381111561079a576107996119c8565b5b81525090508060035f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e0820151816007015f6101000a81548160ff0219169083600381111561086f5761086e6119c8565b5b021790555090505060045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2086908060018154018082558091505060019003905f5260205f2090600491828204019190066008029091909190916101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600360015f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060050160146101000a81548160ff02191690836003811115610952576109516119c8565b5b02179055507ff50b3cb342f0081a2c26e7b8816717154b65d850ef73d68a185a3613c719f7b48686858760405161098c94939291906124a8565b60405180910390a1505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f81819054906101000a900467ffffffffffffffff16809291906109e690612518565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555090505f60015f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20905081815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555089816001019081610a709190612738565b5088816002015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550878160020160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550868160020160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555085816003018190555084816004019081610b0b9190612738565b5033816005015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f8160050160146101000a81548160ff02191690836003811115610b7557610b746119c8565b5b021790555083816006019081610b8b9190612738565b5082816007019081610b9d9190612738565b504381600801819055507fedcc895bbb0933b024b35f376f6d52ad824c61c10d6c463841fc61daf47e37de82604051610bd69190612807565b60405180910390a150505050505050505050565b606060045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020018280548015610ca457602002820191905f5260205f20905f905b82829054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020019060080190602082600701049283019260010382029150808411610c5f5790505b50505050509050919050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f60015f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1603610d5b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5290612249565b60405180910390fd5b60016003811115610d6f57610d6e6119c8565b5b60035f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206007015f9054906101000a900460ff166003811115610db657610db56119c8565b5b1480610e1957505f6003811115610dd057610dcf6119c8565b5b60035f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206007015f9054906101000a900460ff166003811115610e1757610e166119c8565b5b145b610e58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4f90612890565b60405180910390fd5b60015f8267ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206005015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610f4a575060035f8267ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610f89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f809061291e565b60405180910390fd5b600260035f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206007015f6101000a81548160ff02191690836003811115610fd457610fd36119c8565b5b02179055504360035f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600501819055504260035f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600601819055507f3a5b75e7dbdc051c7b1554f8abe12c04268e5fc8b3fb7cdb71fb826edd1aa476816040516110629190612807565b60405180910390a150565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001602052805f5260405f205f91509050805f015f9054906101000a900467ffffffffffffffff16908060010180546110ca90612574565b80601f01602080910402602001604051908101604052809291908181526020018280546110f690612574565b80156111415780601f1061111857610100808354040283529160200191611141565b820191905f5260205f20905b81548152906001019060200180831161112457829003601f168201915b505050505090806002015f9054906101000a900467ffffffffffffffff16908060020160089054906101000a900467ffffffffffffffff16908060020160109054906101000a900467ffffffffffffffff16908060030154908060040180546111a990612574565b80601f01602080910402602001604051908101604052809291908181526020018280546111d590612574565b80156112205780601f106111f757610100808354040283529160200191611220565b820191905f5260205f20905b81548152906001019060200180831161120357829003601f168201915b505050505090806005015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160149054906101000a900460ff169080600601805461126d90612574565b80601f016020809104026020016040519081016040528092919081815260200182805461129990612574565b80156112e45780601f106112bb576101008083540402835291602001916112e4565b820191905f5260205f20905b8154815290600101906020018083116112c757829003601f168201915b5050505050908060070180546112f990612574565b80601f016020809104026020016040519081016040528092919081815260200182805461132590612574565b80156113705780601f1061134757610100808354040283529160200191611370565b820191905f5260205f20905b81548152906001019060200180831161135357829003601f168201915b505050505090806008015490508c565b611388611792565b60035f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20604051806101000160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015f9054906101000a900460ff16600381111561146d5761146c6119c8565b5b600381111561147f5761147e6119c8565b5b815250509050919050565b60015f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600301548111156114f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114ec90612986565b60405180910390fd5b600860015f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206008015461152a91906129a4565b43111561156c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161156390612a21565b60405180910390fd5b5f60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166368efbce3336040518263ffffffff1660e01b81526004016115c79190611b07565b5f60405180830381865afa1580156115e1573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190611609919061239a565b93505050508061164e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161164590612aaf565b60405180910390fd5b60025f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060405180606001604052808567ffffffffffffffff1681526020018481526020013373ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550602082015181600101556040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050507f3a15efc4974c5d6dc4b1c18daba29785cc2d9ae521ad82862e6d32b414b4f5bd83833360405161178593929190611968565b60405180910390a1505050565b6040518061010001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f60038111156117f1576117f06119c8565b5b81525090565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61183182611808565b9050919050565b61184181611827565b811461184b575f80fd5b50565b5f8135905061185c81611838565b92915050565b5f6020828403121561187757611876611800565b5b5f6118848482850161184e565b91505092915050565b5f67ffffffffffffffff82169050919050565b6118a98161188d565b81146118b3575f80fd5b50565b5f813590506118c4816118a0565b92915050565b5f819050919050565b6118dc816118ca565b81146118e6575f80fd5b50565b5f813590506118f7816118d3565b92915050565b5f806040838503121561191357611912611800565b5b5f611920858286016118b6565b9250506020611931858286016118e9565b9150509250929050565b6119448161188d565b82525050565b611953816118ca565b82525050565b61196281611827565b82525050565b5f60608201905061197b5f83018661193b565b611988602083018561194a565b6119956040830184611959565b949350505050565b5f602082840312156119b2576119b1611800565b5b5f6119bf848285016118b6565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60048110611a0657611a056119c8565b5b50565b5f819050611a16826119f5565b919050565b5f611a2582611a09565b9050919050565b611a3581611a1b565b82525050565b5f61010082019050611a4f5f83018b611959565b611a5c602083018a61194a565b611a69604083018961194a565b611a76606083018861194a565b611a83608083018761194a565b611a9060a083018661194a565b611a9d60c083018561194a565b611aaa60e0830184611a2c565b9998505050505050505050565b5f805f60608486031215611ace57611acd611800565b5b5f611adb868287016118b6565b9350506020611aec8682870161184e565b9250506040611afd868287016118e9565b9150509250925092565b5f602082019050611b1a5f830184611959565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611b6e82611b28565b810181811067ffffffffffffffff82111715611b8d57611b8c611b38565b5b80604052505050565b5f611b9f6117f7565b9050611bab8282611b65565b919050565b5f67ffffffffffffffff821115611bca57611bc9611b38565b5b611bd382611b28565b9050602081019050919050565b828183375f83830152505050565b5f611c00611bfb84611bb0565b611b96565b905082815260208101848484011115611c1c57611c1b611b24565b5b611c27848285611be0565b509392505050565b5f82601f830112611c4357611c42611b20565b5b8135611c53848260208601611bee565b91505092915050565b5f805f805f805f80610100898b031215611c7957611c78611800565b5b5f89013567ffffffffffffffff811115611c9657611c95611804565b5b611ca28b828c01611c2f565b9850506020611cb38b828c016118b6565b9750506040611cc48b828c016118b6565b9650506060611cd58b828c016118b6565b9550506080611ce68b828c016118e9565b94505060a089013567ffffffffffffffff811115611d0757611d06611804565b5b611d138b828c01611c2f565b93505060c089013567ffffffffffffffff811115611d3457611d33611804565b5b611d408b828c01611c2f565b92505060e089013567ffffffffffffffff811115611d6157611d60611804565b5b611d6d8b828c01611c2f565b9150509295985092959890939650565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611daf8161188d565b82525050565b5f611dc08383611da6565b60208301905092915050565b5f602082019050919050565b5f611de282611d7d565b611dec8185611d87565b9350611df783611d97565b805f5b83811015611e27578151611e0e8882611db5565b9750611e1983611dcc565b925050600181019050611dfa565b5085935050505092915050565b5f6020820190508181035f830152611e4c8184611dd8565b905092915050565b5f819050919050565b5f611e77611e72611e6d84611808565b611e54565b611808565b9050919050565b5f611e8882611e5d565b9050919050565b5f611e9982611e7e565b9050919050565b611ea981611e8f565b82525050565b5f602082019050611ec25f830184611ea0565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f611efa82611ec8565b611f048185611ed2565b9350611f14818560208601611ee2565b611f1d81611b28565b840191505092915050565b5f61018082019050611f3c5f83018f61193b565b8181036020830152611f4e818e611ef0565b9050611f5d604083018d61193b565b611f6a606083018c61193b565b611f77608083018b61193b565b611f8460a083018a61194a565b81810360c0830152611f968189611ef0565b9050611fa560e0830188611959565b611fb3610100830187611a2c565b818103610120830152611fc68186611ef0565b9050818103610140830152611fdb8185611ef0565b9050611feb61016083018461194a565b9d9c50505050505050505050505050565b61200581611827565b82525050565b612014816118ca565b82525050565b61202381611a1b565b82525050565b61010082015f82015161203e5f850182611ffc565b506020820151612051602085018261200b565b506040820151612064604085018261200b565b506060820151612077606085018261200b565b50608082015161208a608085018261200b565b5060a082015161209d60a085018261200b565b5060c08201516120b060c085018261200b565b5060e08201516120c360e085018261201a565b50505050565b5f610100820190506120dd5f830184612029565b92915050565b7f4f6e6c79206f776e65722063616e20706572666f726d207468697320616374695f8201527f6f6e000000000000000000000000000000000000000000000000000000000000602082015250565b5f61213d602283611ed2565b9150612148826120e3565b604082019050919050565b5f6020820190508181035f83015261216a81612131565b9050919050565b7f4f6e6c7920617574686f72697a656420616464726573732063616e20706572665f8201527f6f726d207468697320616374696f6e0000000000000000000000000000000000602082015250565b5f6121cb602f83611ed2565b91506121d682612171565b604082019050919050565b5f6020820190508181035f8301526121f8816121bf565b9050919050565b7f4f7264657220646f6573206e6f742065786973740000000000000000000000005f82015250565b5f612233601483611ed2565b915061223e826121ff565b602082019050919050565b5f6020820190508181035f83015261226081612227565b9050919050565b7f4f72646572206973206e6f7420616374697665000000000000000000000000005f82015250565b5f61229b601383611ed2565b91506122a682612267565b602082019050919050565b5f6020820190508181035f8301526122c88161228f565b9050919050565b5f815190506122dd816118d3565b92915050565b5f6122f56122f084611bb0565b611b96565b90508281526020810184848401111561231157612310611b24565b5b61231c848285611ee2565b509392505050565b5f82601f83011261233857612337611b20565b5b81516123488482602086016122e3565b91505092915050565b5f8151905061235f81611838565b92915050565b5f8115159050919050565b61237981612365565b8114612383575f80fd5b50565b5f8151905061239481612370565b92915050565b5f805f80608085870312156123b2576123b1611800565b5b5f6123bf878288016122cf565b945050602085015167ffffffffffffffff8111156123e0576123df611804565b5b6123ec87828801612324565b93505060406123fd87828801612351565b925050606061240e87828801612386565b91505092959194509250565b7f50726f7669646572206973206e6f7420616e20616374697665206e6f646520705f8201527f726f766964657200000000000000000000000000000000000000000000000000602082015250565b5f612474602783611ed2565b915061247f8261241a565b604082019050919050565b5f6020820190508181035f8301526124a181612468565b9050919050565b5f6080820190506124bb5f83018761193b565b6124c86020830186611959565b6124d5604083018561194a565b6124e2606083018461194a565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6125228261188d565b915067ffffffffffffffff820361253c5761253b6124eb565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061258b57607f821691505b60208210810361259e5761259d612547565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026126007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826125c5565b61260a86836125c5565b95508019841693508086168417925050509392505050565b5f61263c612637612632846118ca565b611e54565b6118ca565b9050919050565b5f819050919050565b61265583612622565b61266961266182612643565b8484546125d1565b825550505050565b5f90565b61267d612671565b61268881848461264c565b505050565b5b818110156126ab576126a05f82612675565b60018101905061268e565b5050565b601f8211156126f0576126c1816125a4565b6126ca846125b6565b810160208510156126d9578190505b6126ed6126e5856125b6565b83018261268d565b50505b505050565b5f82821c905092915050565b5f6127105f19846008026126f5565b1980831691505092915050565b5f6127288383612701565b9150826002028217905092915050565b61274182611ec8565b67ffffffffffffffff81111561275a57612759611b38565b5b6127648254612574565b61276f8282856126af565b5f60209050601f8311600181146127a0575f841561278e578287015190505b612798858261271d565b8655506127ff565b601f1984166127ae866125a4565b5f5b828110156127d5578489015182556001820191506020850194506020810190506127b0565b868310156127f257848901516127ee601f891682612701565b8355505b6001600288020188555050505b505050505050565b5f60208201905061281a5f83018461193b565b92915050565b7f4f72646572206973206e6f742070726f766973696f6e6564206f72206d6174635f8201527f6865640000000000000000000000000000000000000000000000000000000000602082015250565b5f61287a602383611ed2565b915061288582612820565b604082019050919050565b5f6020820190508181035f8301526128a78161286e565b9050919050565b7f4f6e6c792063726561746f72206f722070726f76696465722063616e20636c6f5f8201527f736520746865206f726465720000000000000000000000000000000000000000602082015250565b5f612908602c83611ed2565b9150612913826128ae565b604082019050919050565b5f6020820190508181035f830152612935816128fc565b9050919050565b7f4269642070726963652065786365656473206d617820707269636500000000005f82015250565b5f612970601b83611ed2565b915061297b8261293c565b602082019050919050565b5f6020820190508181035f83015261299d81612964565b9050919050565b5f6129ae826118ca565b91506129b9836118ca565b92508282019050808211156129d1576129d06124eb565b5b92915050565b7f42696464696e6720706572696f642068617320656e64656400000000000000005f82015250565b5f612a0b601883611ed2565b9150612a16826129d7565b602082019050919050565b5f6020820190508181035f830152612a38816129ff565b9050919050565b7f426964646572206973206e6f7420616e20616374697665206e6f64652070726f5f8201527f7669646572000000000000000000000000000000000000000000000000000000602082015250565b5f612a99602583611ed2565b9150612aa482612a3f565b604082019050919050565b5f6020820190508181035f830152612ac681612a8d565b905091905056fea264697066735822122052f70dfc81e52f5cc9a40b8998499a860b077636d23907d1c46c896f9d39052864736f6c63430008190033",
}

// OrderMatchingABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderMatchingMetaData.ABI instead.
var OrderMatchingABI = OrderMatchingMetaData.ABI

// OrderMatchingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OrderMatchingMetaData.Bin instead.
var OrderMatchingBin = OrderMatchingMetaData.Bin

// DeployOrderMatching deploys a new Ethereum contract, binding an instance of OrderMatching to it.
func DeployOrderMatching(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeProviderRegistryAddress common.Address, _authorizedAddress common.Address) (common.Address, *types.Transaction, *OrderMatching, error) {
	parsed, err := OrderMatchingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OrderMatchingBin), backend, _nodeProviderRegistryAddress, _authorizedAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrderMatching{OrderMatchingCaller: OrderMatchingCaller{contract: contract}, OrderMatchingTransactor: OrderMatchingTransactor{contract: contract}, OrderMatchingFilterer: OrderMatchingFilterer{contract: contract}}, nil
}

// OrderMatching is an auto generated Go binding around an Ethereum contract.
type OrderMatching struct {
	OrderMatchingCaller     // Read-only binding to the contract
	OrderMatchingTransactor // Write-only binding to the contract
	OrderMatchingFilterer   // Log filterer for contract events
}

// OrderMatchingCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderMatchingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderMatchingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderMatchingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderMatchingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderMatchingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderMatchingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderMatchingSession struct {
	Contract     *OrderMatching    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderMatchingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderMatchingCallerSession struct {
	Contract *OrderMatchingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OrderMatchingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderMatchingTransactorSession struct {
	Contract     *OrderMatchingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OrderMatchingRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderMatchingRaw struct {
	Contract *OrderMatching // Generic contract binding to access the raw methods on
}

// OrderMatchingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderMatchingCallerRaw struct {
	Contract *OrderMatchingCaller // Generic read-only contract binding to access the raw methods on
}

// OrderMatchingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderMatchingTransactorRaw struct {
	Contract *OrderMatchingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderMatching creates a new instance of OrderMatching, bound to a specific deployed contract.
func NewOrderMatching(address common.Address, backend bind.ContractBackend) (*OrderMatching, error) {
	contract, err := bindOrderMatching(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderMatching{OrderMatchingCaller: OrderMatchingCaller{contract: contract}, OrderMatchingTransactor: OrderMatchingTransactor{contract: contract}, OrderMatchingFilterer: OrderMatchingFilterer{contract: contract}}, nil
}

// NewOrderMatchingCaller creates a new read-only instance of OrderMatching, bound to a specific deployed contract.
func NewOrderMatchingCaller(address common.Address, caller bind.ContractCaller) (*OrderMatchingCaller, error) {
	contract, err := bindOrderMatching(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderMatchingCaller{contract: contract}, nil
}

// NewOrderMatchingTransactor creates a new write-only instance of OrderMatching, bound to a specific deployed contract.
func NewOrderMatchingTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderMatchingTransactor, error) {
	contract, err := bindOrderMatching(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderMatchingTransactor{contract: contract}, nil
}

// NewOrderMatchingFilterer creates a new log filterer instance of OrderMatching, bound to a specific deployed contract.
func NewOrderMatchingFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderMatchingFilterer, error) {
	contract, err := bindOrderMatching(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderMatchingFilterer{contract: contract}, nil
}

// bindOrderMatching binds a generic wrapper to an already deployed contract.
func bindOrderMatching(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrderMatchingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderMatching *OrderMatchingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderMatching.Contract.OrderMatchingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderMatching *OrderMatchingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderMatching.Contract.OrderMatchingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderMatching *OrderMatchingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderMatching.Contract.OrderMatchingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderMatching *OrderMatchingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderMatching.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderMatching *OrderMatchingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderMatching.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderMatching *OrderMatchingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderMatching.Contract.contract.Transact(opts, method, params...)
}

// AuthorizedAddress is a free data retrieval call binding the contract method 0x5539d400.
//
// Solidity: function authorizedAddress() view returns(address)
func (_OrderMatching *OrderMatchingCaller) AuthorizedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderMatching.contract.Call(opts, &out, "authorizedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthorizedAddress is a free data retrieval call binding the contract method 0x5539d400.
//
// Solidity: function authorizedAddress() view returns(address)
func (_OrderMatching *OrderMatchingSession) AuthorizedAddress() (common.Address, error) {
	return _OrderMatching.Contract.AuthorizedAddress(&_OrderMatching.CallOpts)
}

// AuthorizedAddress is a free data retrieval call binding the contract method 0x5539d400.
//
// Solidity: function authorizedAddress() view returns(address)
func (_OrderMatching *OrderMatchingCallerSession) AuthorizedAddress() (common.Address, error) {
	return _OrderMatching.Contract.AuthorizedAddress(&_OrderMatching.CallOpts)
}

// Bids is a free data retrieval call binding the contract method 0x47d4f8f2.
//
// Solidity: function bids(uint64 , uint256 ) view returns(uint64 orderId, uint256 bidPrice, address bidder)
func (_OrderMatching *OrderMatchingCaller) Bids(opts *bind.CallOpts, arg0 uint64, arg1 *big.Int) (struct {
	OrderId  uint64
	BidPrice *big.Int
	Bidder   common.Address
}, error) {
	var out []interface{}
	err := _OrderMatching.contract.Call(opts, &out, "bids", arg0, arg1)

	outstruct := new(struct {
		OrderId  uint64
		BidPrice *big.Int
		Bidder   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.BidPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Bidder = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Bids is a free data retrieval call binding the contract method 0x47d4f8f2.
//
// Solidity: function bids(uint64 , uint256 ) view returns(uint64 orderId, uint256 bidPrice, address bidder)
func (_OrderMatching *OrderMatchingSession) Bids(arg0 uint64, arg1 *big.Int) (struct {
	OrderId  uint64
	BidPrice *big.Int
	Bidder   common.Address
}, error) {
	return _OrderMatching.Contract.Bids(&_OrderMatching.CallOpts, arg0, arg1)
}

// Bids is a free data retrieval call binding the contract method 0x47d4f8f2.
//
// Solidity: function bids(uint64 , uint256 ) view returns(uint64 orderId, uint256 bidPrice, address bidder)
func (_OrderMatching *OrderMatchingCallerSession) Bids(arg0 uint64, arg1 *big.Int) (struct {
	OrderId  uint64
	BidPrice *big.Int
	Bidder   common.Address
}, error) {
	return _OrderMatching.Contract.Bids(&_OrderMatching.CallOpts, arg0, arg1)
}

// GetOrderById is a free data retrieval call binding the contract method 0xb930807d.
//
// Solidity: function getOrderById(uint64 _orderId) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint8))
func (_OrderMatching *OrderMatchingCaller) GetOrderById(opts *bind.CallOpts, _orderId uint64) (OrderMatchingLease, error) {
	var out []interface{}
	err := _OrderMatching.contract.Call(opts, &out, "getOrderById", _orderId)

	if err != nil {
		return *new(OrderMatchingLease), err
	}

	out0 := *abi.ConvertType(out[0], new(OrderMatchingLease)).(*OrderMatchingLease)

	return out0, err

}

// GetOrderById is a free data retrieval call binding the contract method 0xb930807d.
//
// Solidity: function getOrderById(uint64 _orderId) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint8))
func (_OrderMatching *OrderMatchingSession) GetOrderById(_orderId uint64) (OrderMatchingLease, error) {
	return _OrderMatching.Contract.GetOrderById(&_OrderMatching.CallOpts, _orderId)
}

// GetOrderById is a free data retrieval call binding the contract method 0xb930807d.
//
// Solidity: function getOrderById(uint64 _orderId) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint8))
func (_OrderMatching *OrderMatchingCallerSession) GetOrderById(_orderId uint64) (OrderMatchingLease, error) {
	return _OrderMatching.Contract.GetOrderById(&_OrderMatching.CallOpts, _orderId)
}

// GetOrderByProvider is a free data retrieval call binding the contract method 0x638d7dec.
//
// Solidity: function getOrderByProvider(address _providerAddress) view returns(uint64[])
func (_OrderMatching *OrderMatchingCaller) GetOrderByProvider(opts *bind.CallOpts, _providerAddress common.Address) ([]uint64, error) {
	var out []interface{}
	err := _OrderMatching.contract.Call(opts, &out, "getOrderByProvider", _providerAddress)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetOrderByProvider is a free data retrieval call binding the contract method 0x638d7dec.
//
// Solidity: function getOrderByProvider(address _providerAddress) view returns(uint64[])
func (_OrderMatching *OrderMatchingSession) GetOrderByProvider(_providerAddress common.Address) ([]uint64, error) {
	return _OrderMatching.Contract.GetOrderByProvider(&_OrderMatching.CallOpts, _providerAddress)
}

// GetOrderByProvider is a free data retrieval call binding the contract method 0x638d7dec.
//
// Solidity: function getOrderByProvider(address _providerAddress) view returns(uint64[])
func (_OrderMatching *OrderMatchingCallerSession) GetOrderByProvider(_providerAddress common.Address) ([]uint64, error) {
	return _OrderMatching.Contract.GetOrderByProvider(&_OrderMatching.CallOpts, _providerAddress)
}

// Leases is a free data retrieval call binding the contract method 0x4c7f3c79.
//
// Solidity: function leases(uint64 ) view returns(address providerAddress, uint256 providerId, uint256 acceptedPrice, uint256 startBlock, uint256 startTime, uint256 endBlock, uint256 endTime, uint8 state)
func (_OrderMatching *OrderMatchingCaller) Leases(opts *bind.CallOpts, arg0 uint64) (struct {
	ProviderAddress common.Address
	ProviderId      *big.Int
	AcceptedPrice   *big.Int
	StartBlock      *big.Int
	StartTime       *big.Int
	EndBlock        *big.Int
	EndTime         *big.Int
	State           uint8
}, error) {
	var out []interface{}
	err := _OrderMatching.contract.Call(opts, &out, "leases", arg0)

	outstruct := new(struct {
		ProviderAddress common.Address
		ProviderId      *big.Int
		AcceptedPrice   *big.Int
		StartBlock      *big.Int
		StartTime       *big.Int
		EndBlock        *big.Int
		EndTime         *big.Int
		State           uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProviderAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ProviderId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AcceptedPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.State = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// Leases is a free data retrieval call binding the contract method 0x4c7f3c79.
//
// Solidity: function leases(uint64 ) view returns(address providerAddress, uint256 providerId, uint256 acceptedPrice, uint256 startBlock, uint256 startTime, uint256 endBlock, uint256 endTime, uint8 state)
func (_OrderMatching *OrderMatchingSession) Leases(arg0 uint64) (struct {
	ProviderAddress common.Address
	ProviderId      *big.Int
	AcceptedPrice   *big.Int
	StartBlock      *big.Int
	StartTime       *big.Int
	EndBlock        *big.Int
	EndTime         *big.Int
	State           uint8
}, error) {
	return _OrderMatching.Contract.Leases(&_OrderMatching.CallOpts, arg0)
}

// Leases is a free data retrieval call binding the contract method 0x4c7f3c79.
//
// Solidity: function leases(uint64 ) view returns(address providerAddress, uint256 providerId, uint256 acceptedPrice, uint256 startBlock, uint256 startTime, uint256 endBlock, uint256 endTime, uint8 state)
func (_OrderMatching *OrderMatchingCallerSession) Leases(arg0 uint64) (struct {
	ProviderAddress common.Address
	ProviderId      *big.Int
	AcceptedPrice   *big.Int
	StartBlock      *big.Int
	StartTime       *big.Int
	EndBlock        *big.Int
	EndTime         *big.Int
	State           uint8
}, error) {
	return _OrderMatching.Contract.Leases(&_OrderMatching.CallOpts, arg0)
}

// NodeProviderRegistry is a free data retrieval call binding the contract method 0x6582972e.
//
// Solidity: function nodeProviderRegistry() view returns(address)
func (_OrderMatching *OrderMatchingCaller) NodeProviderRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderMatching.contract.Call(opts, &out, "nodeProviderRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeProviderRegistry is a free data retrieval call binding the contract method 0x6582972e.
//
// Solidity: function nodeProviderRegistry() view returns(address)
func (_OrderMatching *OrderMatchingSession) NodeProviderRegistry() (common.Address, error) {
	return _OrderMatching.Contract.NodeProviderRegistry(&_OrderMatching.CallOpts)
}

// NodeProviderRegistry is a free data retrieval call binding the contract method 0x6582972e.
//
// Solidity: function nodeProviderRegistry() view returns(address)
func (_OrderMatching *OrderMatchingCallerSession) NodeProviderRegistry() (common.Address, error) {
	return _OrderMatching.Contract.NodeProviderRegistry(&_OrderMatching.CallOpts)
}

// Orders is a free data retrieval call binding the contract method 0xa469dffa.
//
// Solidity: function orders(uint64 ) view returns(uint64 id, string region, uint64 uptime, uint64 reputation, uint64 slashes, uint256 maxPrice, string token, address creator, uint8 state, string specs, string version, uint256 creationBlock)
func (_OrderMatching *OrderMatchingCaller) Orders(opts *bind.CallOpts, arg0 uint64) (struct {
	Id            uint64
	Region        string
	Uptime        uint64
	Reputation    uint64
	Slashes       uint64
	MaxPrice      *big.Int
	Token         string
	Creator       common.Address
	State         uint8
	Specs         string
	Version       string
	CreationBlock *big.Int
}, error) {
	var out []interface{}
	err := _OrderMatching.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Id            uint64
		Region        string
		Uptime        uint64
		Reputation    uint64
		Slashes       uint64
		MaxPrice      *big.Int
		Token         string
		Creator       common.Address
		State         uint8
		Specs         string
		Version       string
		CreationBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Region = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Uptime = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Reputation = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Slashes = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.MaxPrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Creator = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.State = *abi.ConvertType(out[8], new(uint8)).(*uint8)
	outstruct.Specs = *abi.ConvertType(out[9], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[10], new(string)).(*string)
	outstruct.CreationBlock = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa469dffa.
//
// Solidity: function orders(uint64 ) view returns(uint64 id, string region, uint64 uptime, uint64 reputation, uint64 slashes, uint256 maxPrice, string token, address creator, uint8 state, string specs, string version, uint256 creationBlock)
func (_OrderMatching *OrderMatchingSession) Orders(arg0 uint64) (struct {
	Id            uint64
	Region        string
	Uptime        uint64
	Reputation    uint64
	Slashes       uint64
	MaxPrice      *big.Int
	Token         string
	Creator       common.Address
	State         uint8
	Specs         string
	Version       string
	CreationBlock *big.Int
}, error) {
	return _OrderMatching.Contract.Orders(&_OrderMatching.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa469dffa.
//
// Solidity: function orders(uint64 ) view returns(uint64 id, string region, uint64 uptime, uint64 reputation, uint64 slashes, uint256 maxPrice, string token, address creator, uint8 state, string specs, string version, uint256 creationBlock)
func (_OrderMatching *OrderMatchingCallerSession) Orders(arg0 uint64) (struct {
	Id            uint64
	Region        string
	Uptime        uint64
	Reputation    uint64
	Slashes       uint64
	MaxPrice      *big.Int
	Token         string
	Creator       common.Address
	State         uint8
	Specs         string
	Version       string
	CreationBlock *big.Int
}, error) {
	return _OrderMatching.Contract.Orders(&_OrderMatching.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderMatching *OrderMatchingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderMatching.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderMatching *OrderMatchingSession) Owner() (common.Address, error) {
	return _OrderMatching.Contract.Owner(&_OrderMatching.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderMatching *OrderMatchingCallerSession) Owner() (common.Address, error) {
	return _OrderMatching.Contract.Owner(&_OrderMatching.CallOpts)
}

// CloseOrder is a paid mutator transaction binding the contract method 0x7a544ad3.
//
// Solidity: function closeOrder(uint64 _orderId) returns()
func (_OrderMatching *OrderMatchingTransactor) CloseOrder(opts *bind.TransactOpts, _orderId uint64) (*types.Transaction, error) {
	return _OrderMatching.contract.Transact(opts, "closeOrder", _orderId)
}

// CloseOrder is a paid mutator transaction binding the contract method 0x7a544ad3.
//
// Solidity: function closeOrder(uint64 _orderId) returns()
func (_OrderMatching *OrderMatchingSession) CloseOrder(_orderId uint64) (*types.Transaction, error) {
	return _OrderMatching.Contract.CloseOrder(&_OrderMatching.TransactOpts, _orderId)
}

// CloseOrder is a paid mutator transaction binding the contract method 0x7a544ad3.
//
// Solidity: function closeOrder(uint64 _orderId) returns()
func (_OrderMatching *OrderMatchingTransactorSession) CloseOrder(_orderId uint64) (*types.Transaction, error) {
	return _OrderMatching.Contract.CloseOrder(&_OrderMatching.TransactOpts, _orderId)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x59e7bf4b.
//
// Solidity: function createOrder(string _region, uint64 _uptime, uint64 _reputation, uint64 _slashes, uint256 _maxPrice, string _token, string _spec, string _version) returns()
func (_OrderMatching *OrderMatchingTransactor) CreateOrder(opts *bind.TransactOpts, _region string, _uptime uint64, _reputation uint64, _slashes uint64, _maxPrice *big.Int, _token string, _spec string, _version string) (*types.Transaction, error) {
	return _OrderMatching.contract.Transact(opts, "createOrder", _region, _uptime, _reputation, _slashes, _maxPrice, _token, _spec, _version)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x59e7bf4b.
//
// Solidity: function createOrder(string _region, uint64 _uptime, uint64 _reputation, uint64 _slashes, uint256 _maxPrice, string _token, string _spec, string _version) returns()
func (_OrderMatching *OrderMatchingSession) CreateOrder(_region string, _uptime uint64, _reputation uint64, _slashes uint64, _maxPrice *big.Int, _token string, _spec string, _version string) (*types.Transaction, error) {
	return _OrderMatching.Contract.CreateOrder(&_OrderMatching.TransactOpts, _region, _uptime, _reputation, _slashes, _maxPrice, _token, _spec, _version)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x59e7bf4b.
//
// Solidity: function createOrder(string _region, uint64 _uptime, uint64 _reputation, uint64 _slashes, uint256 _maxPrice, string _token, string _spec, string _version) returns()
func (_OrderMatching *OrderMatchingTransactorSession) CreateOrder(_region string, _uptime uint64, _reputation uint64, _slashes uint64, _maxPrice *big.Int, _token string, _spec string, _version string) (*types.Transaction, error) {
	return _OrderMatching.Contract.CreateOrder(&_OrderMatching.TransactOpts, _region, _uptime, _reputation, _slashes, _maxPrice, _token, _spec, _version)
}

// MatchOrder is a paid mutator transaction binding the contract method 0x4fd2edbb.
//
// Solidity: function matchOrder(uint64 _orderId, address _providerAddress, uint256 _acceptedPrice) returns()
func (_OrderMatching *OrderMatchingTransactor) MatchOrder(opts *bind.TransactOpts, _orderId uint64, _providerAddress common.Address, _acceptedPrice *big.Int) (*types.Transaction, error) {
	return _OrderMatching.contract.Transact(opts, "matchOrder", _orderId, _providerAddress, _acceptedPrice)
}

// MatchOrder is a paid mutator transaction binding the contract method 0x4fd2edbb.
//
// Solidity: function matchOrder(uint64 _orderId, address _providerAddress, uint256 _acceptedPrice) returns()
func (_OrderMatching *OrderMatchingSession) MatchOrder(_orderId uint64, _providerAddress common.Address, _acceptedPrice *big.Int) (*types.Transaction, error) {
	return _OrderMatching.Contract.MatchOrder(&_OrderMatching.TransactOpts, _orderId, _providerAddress, _acceptedPrice)
}

// MatchOrder is a paid mutator transaction binding the contract method 0x4fd2edbb.
//
// Solidity: function matchOrder(uint64 _orderId, address _providerAddress, uint256 _acceptedPrice) returns()
func (_OrderMatching *OrderMatchingTransactorSession) MatchOrder(_orderId uint64, _providerAddress common.Address, _acceptedPrice *big.Int) (*types.Transaction, error) {
	return _OrderMatching.Contract.MatchOrder(&_OrderMatching.TransactOpts, _orderId, _providerAddress, _acceptedPrice)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xfd2e5a23.
//
// Solidity: function placeBid(uint64 _orderId, uint256 _bidPrice) returns()
func (_OrderMatching *OrderMatchingTransactor) PlaceBid(opts *bind.TransactOpts, _orderId uint64, _bidPrice *big.Int) (*types.Transaction, error) {
	return _OrderMatching.contract.Transact(opts, "placeBid", _orderId, _bidPrice)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xfd2e5a23.
//
// Solidity: function placeBid(uint64 _orderId, uint256 _bidPrice) returns()
func (_OrderMatching *OrderMatchingSession) PlaceBid(_orderId uint64, _bidPrice *big.Int) (*types.Transaction, error) {
	return _OrderMatching.Contract.PlaceBid(&_OrderMatching.TransactOpts, _orderId, _bidPrice)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xfd2e5a23.
//
// Solidity: function placeBid(uint64 _orderId, uint256 _bidPrice) returns()
func (_OrderMatching *OrderMatchingTransactorSession) PlaceBid(_orderId uint64, _bidPrice *big.Int) (*types.Transaction, error) {
	return _OrderMatching.Contract.PlaceBid(&_OrderMatching.TransactOpts, _orderId, _bidPrice)
}

// UpdateAuthorizedAddress is a paid mutator transaction binding the contract method 0x28d389bc.
//
// Solidity: function updateAuthorizedAddress(address _newAuthorizedAddress) returns()
func (_OrderMatching *OrderMatchingTransactor) UpdateAuthorizedAddress(opts *bind.TransactOpts, _newAuthorizedAddress common.Address) (*types.Transaction, error) {
	return _OrderMatching.contract.Transact(opts, "updateAuthorizedAddress", _newAuthorizedAddress)
}

// UpdateAuthorizedAddress is a paid mutator transaction binding the contract method 0x28d389bc.
//
// Solidity: function updateAuthorizedAddress(address _newAuthorizedAddress) returns()
func (_OrderMatching *OrderMatchingSession) UpdateAuthorizedAddress(_newAuthorizedAddress common.Address) (*types.Transaction, error) {
	return _OrderMatching.Contract.UpdateAuthorizedAddress(&_OrderMatching.TransactOpts, _newAuthorizedAddress)
}

// UpdateAuthorizedAddress is a paid mutator transaction binding the contract method 0x28d389bc.
//
// Solidity: function updateAuthorizedAddress(address _newAuthorizedAddress) returns()
func (_OrderMatching *OrderMatchingTransactorSession) UpdateAuthorizedAddress(_newAuthorizedAddress common.Address) (*types.Transaction, error) {
	return _OrderMatching.Contract.UpdateAuthorizedAddress(&_OrderMatching.TransactOpts, _newAuthorizedAddress)
}

// OrderMatchingAuthorizedAddressUpdatedIterator is returned from FilterAuthorizedAddressUpdated and is used to iterate over the raw logs and unpacked data for AuthorizedAddressUpdated events raised by the OrderMatching contract.
type OrderMatchingAuthorizedAddressUpdatedIterator struct {
	Event *OrderMatchingAuthorizedAddressUpdated // Event containing the contract specifics and raw log

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
func (it *OrderMatchingAuthorizedAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderMatchingAuthorizedAddressUpdated)
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
		it.Event = new(OrderMatchingAuthorizedAddressUpdated)
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
func (it *OrderMatchingAuthorizedAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderMatchingAuthorizedAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderMatchingAuthorizedAddressUpdated represents a AuthorizedAddressUpdated event raised by the OrderMatching contract.
type OrderMatchingAuthorizedAddressUpdated struct {
	NewAuthorizedAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedAddressUpdated is a free log retrieval operation binding the contract event 0xe12f594f8990d16225e1238de64b5ef16fbe78fbb1a70f64e3a82274e9554c0a.
//
// Solidity: event AuthorizedAddressUpdated(address newAuthorizedAddress)
func (_OrderMatching *OrderMatchingFilterer) FilterAuthorizedAddressUpdated(opts *bind.FilterOpts) (*OrderMatchingAuthorizedAddressUpdatedIterator, error) {

	logs, sub, err := _OrderMatching.contract.FilterLogs(opts, "AuthorizedAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &OrderMatchingAuthorizedAddressUpdatedIterator{contract: _OrderMatching.contract, event: "AuthorizedAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorizedAddressUpdated is a free log subscription operation binding the contract event 0xe12f594f8990d16225e1238de64b5ef16fbe78fbb1a70f64e3a82274e9554c0a.
//
// Solidity: event AuthorizedAddressUpdated(address newAuthorizedAddress)
func (_OrderMatching *OrderMatchingFilterer) WatchAuthorizedAddressUpdated(opts *bind.WatchOpts, sink chan<- *OrderMatchingAuthorizedAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _OrderMatching.contract.WatchLogs(opts, "AuthorizedAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderMatchingAuthorizedAddressUpdated)
				if err := _OrderMatching.contract.UnpackLog(event, "AuthorizedAddressUpdated", log); err != nil {
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

// ParseAuthorizedAddressUpdated is a log parse operation binding the contract event 0xe12f594f8990d16225e1238de64b5ef16fbe78fbb1a70f64e3a82274e9554c0a.
//
// Solidity: event AuthorizedAddressUpdated(address newAuthorizedAddress)
func (_OrderMatching *OrderMatchingFilterer) ParseAuthorizedAddressUpdated(log types.Log) (*OrderMatchingAuthorizedAddressUpdated, error) {
	event := new(OrderMatchingAuthorizedAddressUpdated)
	if err := _OrderMatching.contract.UnpackLog(event, "AuthorizedAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderMatchingBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the OrderMatching contract.
type OrderMatchingBidPlacedIterator struct {
	Event *OrderMatchingBidPlaced // Event containing the contract specifics and raw log

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
func (it *OrderMatchingBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderMatchingBidPlaced)
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
		it.Event = new(OrderMatchingBidPlaced)
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
func (it *OrderMatchingBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderMatchingBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderMatchingBidPlaced represents a BidPlaced event raised by the OrderMatching contract.
type OrderMatchingBidPlaced struct {
	OrderId  uint64
	BidPrice *big.Int
	Bidder   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x3a15efc4974c5d6dc4b1c18daba29785cc2d9ae521ad82862e6d32b414b4f5bd.
//
// Solidity: event BidPlaced(uint64 orderId, uint256 bidPrice, address bidder)
func (_OrderMatching *OrderMatchingFilterer) FilterBidPlaced(opts *bind.FilterOpts) (*OrderMatchingBidPlacedIterator, error) {

	logs, sub, err := _OrderMatching.contract.FilterLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return &OrderMatchingBidPlacedIterator{contract: _OrderMatching.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x3a15efc4974c5d6dc4b1c18daba29785cc2d9ae521ad82862e6d32b414b4f5bd.
//
// Solidity: event BidPlaced(uint64 orderId, uint256 bidPrice, address bidder)
func (_OrderMatching *OrderMatchingFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *OrderMatchingBidPlaced) (event.Subscription, error) {

	logs, sub, err := _OrderMatching.contract.WatchLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderMatchingBidPlaced)
				if err := _OrderMatching.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0x3a15efc4974c5d6dc4b1c18daba29785cc2d9ae521ad82862e6d32b414b4f5bd.
//
// Solidity: event BidPlaced(uint64 orderId, uint256 bidPrice, address bidder)
func (_OrderMatching *OrderMatchingFilterer) ParseBidPlaced(log types.Log) (*OrderMatchingBidPlaced, error) {
	event := new(OrderMatchingBidPlaced)
	if err := _OrderMatching.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderMatchingOrderClosedIterator is returned from FilterOrderClosed and is used to iterate over the raw logs and unpacked data for OrderClosed events raised by the OrderMatching contract.
type OrderMatchingOrderClosedIterator struct {
	Event *OrderMatchingOrderClosed // Event containing the contract specifics and raw log

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
func (it *OrderMatchingOrderClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderMatchingOrderClosed)
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
		it.Event = new(OrderMatchingOrderClosed)
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
func (it *OrderMatchingOrderClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderMatchingOrderClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderMatchingOrderClosed represents a OrderClosed event raised by the OrderMatching contract.
type OrderMatchingOrderClosed struct {
	OrderId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderClosed is a free log retrieval operation binding the contract event 0x3a5b75e7dbdc051c7b1554f8abe12c04268e5fc8b3fb7cdb71fb826edd1aa476.
//
// Solidity: event OrderClosed(uint64 orderId)
func (_OrderMatching *OrderMatchingFilterer) FilterOrderClosed(opts *bind.FilterOpts) (*OrderMatchingOrderClosedIterator, error) {

	logs, sub, err := _OrderMatching.contract.FilterLogs(opts, "OrderClosed")
	if err != nil {
		return nil, err
	}
	return &OrderMatchingOrderClosedIterator{contract: _OrderMatching.contract, event: "OrderClosed", logs: logs, sub: sub}, nil
}

// WatchOrderClosed is a free log subscription operation binding the contract event 0x3a5b75e7dbdc051c7b1554f8abe12c04268e5fc8b3fb7cdb71fb826edd1aa476.
//
// Solidity: event OrderClosed(uint64 orderId)
func (_OrderMatching *OrderMatchingFilterer) WatchOrderClosed(opts *bind.WatchOpts, sink chan<- *OrderMatchingOrderClosed) (event.Subscription, error) {

	logs, sub, err := _OrderMatching.contract.WatchLogs(opts, "OrderClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderMatchingOrderClosed)
				if err := _OrderMatching.contract.UnpackLog(event, "OrderClosed", log); err != nil {
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

// ParseOrderClosed is a log parse operation binding the contract event 0x3a5b75e7dbdc051c7b1554f8abe12c04268e5fc8b3fb7cdb71fb826edd1aa476.
//
// Solidity: event OrderClosed(uint64 orderId)
func (_OrderMatching *OrderMatchingFilterer) ParseOrderClosed(log types.Log) (*OrderMatchingOrderClosed, error) {
	event := new(OrderMatchingOrderClosed)
	if err := _OrderMatching.contract.UnpackLog(event, "OrderClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderMatchingOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the OrderMatching contract.
type OrderMatchingOrderCreatedIterator struct {
	Event *OrderMatchingOrderCreated // Event containing the contract specifics and raw log

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
func (it *OrderMatchingOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderMatchingOrderCreated)
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
		it.Event = new(OrderMatchingOrderCreated)
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
func (it *OrderMatchingOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderMatchingOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderMatchingOrderCreated represents a OrderCreated event raised by the OrderMatching contract.
type OrderMatchingOrderCreated struct {
	OrderId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0xedcc895bbb0933b024b35f376f6d52ad824c61c10d6c463841fc61daf47e37de.
//
// Solidity: event OrderCreated(uint64 orderId)
func (_OrderMatching *OrderMatchingFilterer) FilterOrderCreated(opts *bind.FilterOpts) (*OrderMatchingOrderCreatedIterator, error) {

	logs, sub, err := _OrderMatching.contract.FilterLogs(opts, "OrderCreated")
	if err != nil {
		return nil, err
	}
	return &OrderMatchingOrderCreatedIterator{contract: _OrderMatching.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0xedcc895bbb0933b024b35f376f6d52ad824c61c10d6c463841fc61daf47e37de.
//
// Solidity: event OrderCreated(uint64 orderId)
func (_OrderMatching *OrderMatchingFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *OrderMatchingOrderCreated) (event.Subscription, error) {

	logs, sub, err := _OrderMatching.contract.WatchLogs(opts, "OrderCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderMatchingOrderCreated)
				if err := _OrderMatching.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0xedcc895bbb0933b024b35f376f6d52ad824c61c10d6c463841fc61daf47e37de.
//
// Solidity: event OrderCreated(uint64 orderId)
func (_OrderMatching *OrderMatchingFilterer) ParseOrderCreated(log types.Log) (*OrderMatchingOrderCreated, error) {
	event := new(OrderMatchingOrderCreated)
	if err := _OrderMatching.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderMatchingOrderMatchedIterator is returned from FilterOrderMatched and is used to iterate over the raw logs and unpacked data for OrderMatched events raised by the OrderMatching contract.
type OrderMatchingOrderMatchedIterator struct {
	Event *OrderMatchingOrderMatched // Event containing the contract specifics and raw log

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
func (it *OrderMatchingOrderMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderMatchingOrderMatched)
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
		it.Event = new(OrderMatchingOrderMatched)
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
func (it *OrderMatchingOrderMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderMatchingOrderMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderMatchingOrderMatched represents a OrderMatched event raised by the OrderMatching contract.
type OrderMatchingOrderMatched struct {
	OrderId         uint64
	ProviderAddress common.Address
	ProviderId      *big.Int
	AcceptedPrice   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderMatched is a free log retrieval operation binding the contract event 0xf50b3cb342f0081a2c26e7b8816717154b65d850ef73d68a185a3613c719f7b4.
//
// Solidity: event OrderMatched(uint64 orderId, address providerAddress, uint256 providerId, uint256 acceptedPrice)
func (_OrderMatching *OrderMatchingFilterer) FilterOrderMatched(opts *bind.FilterOpts) (*OrderMatchingOrderMatchedIterator, error) {

	logs, sub, err := _OrderMatching.contract.FilterLogs(opts, "OrderMatched")
	if err != nil {
		return nil, err
	}
	return &OrderMatchingOrderMatchedIterator{contract: _OrderMatching.contract, event: "OrderMatched", logs: logs, sub: sub}, nil
}

// WatchOrderMatched is a free log subscription operation binding the contract event 0xf50b3cb342f0081a2c26e7b8816717154b65d850ef73d68a185a3613c719f7b4.
//
// Solidity: event OrderMatched(uint64 orderId, address providerAddress, uint256 providerId, uint256 acceptedPrice)
func (_OrderMatching *OrderMatchingFilterer) WatchOrderMatched(opts *bind.WatchOpts, sink chan<- *OrderMatchingOrderMatched) (event.Subscription, error) {

	logs, sub, err := _OrderMatching.contract.WatchLogs(opts, "OrderMatched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderMatchingOrderMatched)
				if err := _OrderMatching.contract.UnpackLog(event, "OrderMatched", log); err != nil {
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

// ParseOrderMatched is a log parse operation binding the contract event 0xf50b3cb342f0081a2c26e7b8816717154b65d850ef73d68a185a3613c719f7b4.
//
// Solidity: event OrderMatched(uint64 orderId, address providerAddress, uint256 providerId, uint256 acceptedPrice)
func (_OrderMatching *OrderMatchingFilterer) ParseOrderMatched(log types.Log) (*OrderMatchingOrderMatched, error) {
	event := new(OrderMatchingOrderMatched)
	if err := _OrderMatching.contract.UnpackLog(event, "OrderMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
