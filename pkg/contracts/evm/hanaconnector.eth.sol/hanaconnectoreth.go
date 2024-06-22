// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hanaconnector

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

// HanaInterfacesSendInput is an auto generated low-level Go binding around an user-defined struct.
type HanaInterfacesSendInput struct {
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	DestinationGasLimit *big.Int
	Message             []byte
	HanaValueAndGas     *big.Int
	HanaParams          []byte
}

// HanaConnectorEthMetaData contains all meta data concerning the HanaConnectorEth contract.
var HanaConnectorEthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hanaToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tssAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tssAddressUpdater_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauserAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTss\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssOrUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HanaTranserError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hanaTxSenderAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hanaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"HanaReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingHanaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"HanaReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceTxOriginAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hanaValueAndGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hanaParams\",\"type\":\"bytes\"}],\"name\":\"HanaSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updaterAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTssAddress\",\"type\":\"address\"}],\"name\":\"PauserAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTssAddress\",\"type\":\"address\"}],\"name\":\"TSSAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLockedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hanaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"hanaTxSenderAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"onReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingHanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"onRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceTssAddressUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"hanaValueAndGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hanaParams\",\"type\":\"bytes\"}],\"internalType\":\"structHanaInterfaces.SendInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauserAddress_\",\"type\":\"address\"}],\"name\":\"updatePauserAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tssAddress_\",\"type\":\"address\"}],\"name\":\"updateTssAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200208b3803806200208b833981810160405281019062000037919062000284565b8383838360008060006101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480620000bd5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b80620000f55750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b806200012d5750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b1562000165576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b8152505082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050505050505062000349565b6000815190506200027e816200032f565b92915050565b60008060008060808587031215620002a157620002a06200032a565b5b6000620002b1878288016200026d565b9450506020620002c4878288016200026d565b9350506040620002d7878288016200026d565b9250506060620002ea878288016200026d565b91505092959194509250565b600062000303826200030a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200033a81620002f6565b81146200034657600080fd5b50565b60805160601c611d076200038460003960008181610251015281816103930152818161071601528181610d3a0152610fbd0152611d076000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80636128480f1161008c5780639122c344116100665780639122c344146101db578063942a5e16146101f7578063ec02690114610213578063f7fb869b1461022f576100ea565b80636128480f146101ab578063779e3b63146101c75780638456cb59146101d1576100ea565b80633f4ba83a116100c85780633f4ba83a146101475780635b112591146101515780635c975abb1461016f5780635e694a921461018d576100ea565b8063252bc886146100ef57806329dd214d1461010d578063328a01d014610129575b600080fd5b6100f761024d565b6040516101049190611a95565b60405180910390f35b610127600480360381019061012291906114df565b6102fd565b005b610131610616565b60405161013e9190611828565b60405180910390f35b61014f61063c565b005b6101596106d8565b6040516101669190611828565b60405180910390f35b6101776106fe565b60405161018491906119ad565b60405180910390f35b610195610714565b6040516101a29190611828565b60405180910390f35b6101c560048036038101906101c091906113a3565b610738565b005b6101cf6108ae565b005b6101d9610a2e565b005b6101f560048036038101906101f091906113a3565b610aca565b005b610211600480360381019061020c91906113d0565b610c9c565b005b61022d600480360381019061022891906115ae565b610fb1565b005b610237611140565b6040516102449190611828565b60405180910390f35b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016102a89190611828565b60206040518083038186803b1580156102c057600080fd5b505afa1580156102d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102f891906115f7565b905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461038f57336040517fff70ace20000000000000000000000000000000000000000000000000000000081526004016103869190611828565b60405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb87876040518363ffffffff1660e01b81526004016103ec92919061191f565b602060405180830381600087803b15801561040657600080fd5b505af115801561041a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043e91906114b2565b905080610477576040517f1071d9da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008484905011156105b3578573ffffffffffffffffffffffffffffffffffffffff1663b204be936040518060a001604052808c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018a81526020018973ffffffffffffffffffffffffffffffffffffffff16815260200188815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152506040518263ffffffff1660e01b81526004016105809190611a51565b600060405180830381600087803b15801561059a57600080fd5b505af11580156105ae573d6000803e3d6000fd5b505050505b818673ffffffffffffffffffffffffffffffffffffffff16887fb29e5f376a4a399e83dce3b4e7f71bc32f1ab109ccf94574f8f4870cf4aa341f8c8c8a8a8a6040516106039594939291906119c8565b60405180910390a4505050505050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ce57336040517f4677a0d30000000000000000000000000000000000000000000000000000000081526004016106c59190611828565b60405180910390fd5b6106d6611166565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900460ff16905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107ca57336040517f4677a0d30000000000000000000000000000000000000000000000000000000081526004016107c19190611828565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610831576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d039733826040516108a3929190611843565b60405180910390a150565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461094057336040517fe700765e0000000000000000000000000000000000000000000000000000000081526004016109379190611828565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156109c9576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac057336040517f4677a0d3000000000000000000000000000000000000000000000000000000008152600401610ab79190611828565b60405180910390fd5b610ac86111c8565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610b765750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610bb857336040517fcdfcef97000000000000000000000000000000000000000000000000000000008152600401610baf9190611828565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610c1f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff3382604051610c91929190611843565b60405180910390a150565b610ca461122a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d3657336040517fff70ace2000000000000000000000000000000000000000000000000000000008152600401610d2d9190611828565b60405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8b876040518363ffffffff1660e01b8152600401610d9392919061191f565b602060405180830381600087803b158015610dad57600080fd5b505af1158015610dc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610de591906114b2565b905080610e1e576040517f1071d9da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000848490501115610f60578973ffffffffffffffffffffffffffffffffffffffff16634204cf9b6040518060c001604052808d73ffffffffffffffffffffffffffffffffffffffff1681526020018c81526020018b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200189815260200188815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152506040518263ffffffff1660e01b8152600401610f2d9190611a73565b600060405180830381600087803b158015610f4757600080fd5b505af1158015610f5b573d6000803e3d6000fd5b505050505b81867fa0589272400cb97d69e64e37e4dbd1f2aaa79d5b3aa0688e6f2d992ccf5009218c8c8c8c8b8b8b604051610f9d9796959493929190611948565b60405180910390a350505050505050505050565b610fb961122a565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd333085608001356040518463ffffffff1660e01b815260040161101c9392919061186c565b602060405180830381600087803b15801561103657600080fd5b505af115801561104a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061106e91906114b2565b9050806110a7576040517f1071d9da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600001353373ffffffffffffffffffffffffffffffffffffffff167f7a0dd478962bca8db7a72b684fbfd3404be69978e5f8f1c8eab2cfe4724ebaee328580602001906110f59190611ab0565b8760800135886040013589806060019061110f9190611ab0565b8b8060a0019061111f9190611ab0565b604051611134999897969594939291906118a3565b60405180910390a35050565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61116e611274565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6111b16112bd565b6040516111be9190611828565b60405180910390a1565b6111d061122a565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586112136112bd565b6040516112209190611828565b60405180910390a1565b6112326106fe565b15611272576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161126990611a31565b60405180910390fd5b565b61127c6106fe565b6112bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112b290611a11565b60405180910390fd5b565b600033905090565b6000813590506112d481611c75565b92915050565b6000815190506112e981611c8c565b92915050565b6000813590506112fe81611ca3565b92915050565b60008083601f84011261131a57611319611bea565b5b8235905067ffffffffffffffff81111561133757611336611be5565b5b60208301915083600182028301111561135357611352611bfe565b5b9250929050565b600060c082840312156113705761136f611bf4565b5b81905092915050565b60008135905061138881611cba565b92915050565b60008151905061139d81611cba565b92915050565b6000602082840312156113b9576113b8611c0d565b5b60006113c7848285016112c5565b91505092915050565b600080600080600080600080600060e08a8c0312156113f2576113f1611c0d565b5b60006114008c828d016112c5565b99505060206114118c828d01611379565b98505060408a013567ffffffffffffffff81111561143257611431611c08565b5b61143e8c828d01611304565b975097505060606114518c828d01611379565b95505060806114628c828d01611379565b94505060a08a013567ffffffffffffffff81111561148357611482611c08565b5b61148f8c828d01611304565b935093505060c06114a28c828d016112ef565b9150509295985092959850929598565b6000602082840312156114c8576114c7611c0d565b5b60006114d6848285016112da565b91505092915050565b60008060008060008060008060c0898b0312156114ff576114fe611c0d565b5b600089013567ffffffffffffffff81111561151d5761151c611c08565b5b6115298b828c01611304565b9850985050602061153c8b828c01611379565b965050604061154d8b828c016112c5565b955050606061155e8b828c01611379565b945050608089013567ffffffffffffffff81111561157f5761157e611c08565b5b61158b8b828c01611304565b935093505060a061159e8b828c016112ef565b9150509295985092959890939650565b6000602082840312156115c4576115c3611c0d565b5b600082013567ffffffffffffffff8111156115e2576115e1611c08565b5b6115ee8482850161135a565b91505092915050565b60006020828403121561160d5761160c611c0d565b5b600061161b8482850161138e565b91505092915050565b61162d81611b51565b82525050565b61163c81611b51565b82525050565b61164b81611b63565b82525050565b600061165d8385611b2f565b935061166a838584611ba3565b61167383611c12565b840190509392505050565b600061168982611b13565b6116938185611b1e565b93506116a3818560208601611bb2565b6116ac81611c12565b840191505092915050565b60006116c4601483611b40565b91506116cf82611c23565b602082019050919050565b60006116e7601083611b40565b91506116f282611c4c565b602082019050919050565b600060a083016000830151848203600086015261171a828261167e565b915050602083015161172f602086018261180a565b5060408301516117426040860182611624565b506060830151611755606086018261180a565b506080830151848203608086015261176d828261167e565b9150508091505092915050565b600060c0830160008301516117926000860182611624565b5060208301516117a5602086018261180a565b50604083015184820360408601526117bd828261167e565b91505060608301516117d2606086018261180a565b5060808301516117e5608086018261180a565b5060a083015184820360a08601526117fd828261167e565b9150508091505092915050565b61181381611b99565b82525050565b61182281611b99565b82525050565b600060208201905061183d6000830184611633565b92915050565b60006040820190506118586000830185611633565b6118656020830184611633565b9392505050565b60006060820190506118816000830186611633565b61188e6020830185611633565b61189b6040830184611819565b949350505050565b600060c0820190506118b8600083018c611633565b81810360208301526118cb818a8c611651565b90506118da6040830189611819565b6118e76060830188611819565b81810360808301526118fa818688611651565b905081810360a083015261190f818486611651565b90509a9950505050505050505050565b60006040820190506119346000830185611633565b6119416020830184611819565b9392505050565b600060a08201905061195d600083018a611633565b61196a6020830189611819565b818103604083015261197d818789611651565b905061198c6060830186611819565b818103608083015261199f818486611651565b905098975050505050505050565b60006020820190506119c26000830184611642565b92915050565b600060608201905081810360008301526119e3818789611651565b90506119f26020830186611819565b8181036040830152611a05818486611651565b90509695505050505050565b60006020820190508181036000830152611a2a816116b7565b9050919050565b60006020820190508181036000830152611a4a816116da565b9050919050565b60006020820190508181036000830152611a6b81846116fd565b905092915050565b60006020820190508181036000830152611a8d818461177a565b905092915050565b6000602082019050611aaa6000830184611819565b92915050565b60008083356001602003843603038112611acd57611acc611bf9565b5b80840192508235915067ffffffffffffffff821115611aef57611aee611bef565b5b602083019250600182023603831315611b0b57611b0a611c03565b5b509250929050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611b5c82611b79565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611bd0578082015181840152602081019050611bb5565b83811115611bdf576000848401525b50505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b611c7e81611b51565b8114611c8957600080fd5b50565b611c9581611b63565b8114611ca057600080fd5b50565b611cac81611b6f565b8114611cb757600080fd5b50565b611cc381611b99565b8114611cce57600080fd5b5056fea2646970667358221220dca621e232f9a3241f68a20662b675533859ef752a56b4b52672bdc2d522761364736f6c63430008070033",
}

// HanaConnectorEthABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaConnectorEthMetaData.ABI instead.
var HanaConnectorEthABI = HanaConnectorEthMetaData.ABI

// HanaConnectorEthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HanaConnectorEthMetaData.Bin instead.
var HanaConnectorEthBin = HanaConnectorEthMetaData.Bin

// DeployHanaConnectorEth deploys a new Ethereum contract, binding an instance of HanaConnectorEth to it.
func DeployHanaConnectorEth(auth *bind.TransactOpts, backend bind.ContractBackend, hanaToken_ common.Address, tssAddress_ common.Address, tssAddressUpdater_ common.Address, pauserAddress_ common.Address) (common.Address, *types.Transaction, *HanaConnectorEth, error) {
	parsed, err := HanaConnectorEthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HanaConnectorEthBin), backend, hanaToken_, tssAddress_, tssAddressUpdater_, pauserAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HanaConnectorEth{HanaConnectorEthCaller: HanaConnectorEthCaller{contract: contract}, HanaConnectorEthTransactor: HanaConnectorEthTransactor{contract: contract}, HanaConnectorEthFilterer: HanaConnectorEthFilterer{contract: contract}}, nil
}

// HanaConnectorEth is an auto generated Go binding around an Ethereum contract.
type HanaConnectorEth struct {
	HanaConnectorEthCaller     // Read-only binding to the contract
	HanaConnectorEthTransactor // Write-only binding to the contract
	HanaConnectorEthFilterer   // Log filterer for contract events
}

// HanaConnectorEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaConnectorEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaConnectorEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaConnectorEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaConnectorEthSession struct {
	Contract     *HanaConnectorEth // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HanaConnectorEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaConnectorEthCallerSession struct {
	Contract *HanaConnectorEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// HanaConnectorEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaConnectorEthTransactorSession struct {
	Contract     *HanaConnectorEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// HanaConnectorEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaConnectorEthRaw struct {
	Contract *HanaConnectorEth // Generic contract binding to access the raw methods on
}

// HanaConnectorEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaConnectorEthCallerRaw struct {
	Contract *HanaConnectorEthCaller // Generic read-only contract binding to access the raw methods on
}

// HanaConnectorEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaConnectorEthTransactorRaw struct {
	Contract *HanaConnectorEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaConnectorEth creates a new instance of HanaConnectorEth, bound to a specific deployed contract.
func NewHanaConnectorEth(address common.Address, backend bind.ContractBackend) (*HanaConnectorEth, error) {
	contract, err := bindHanaConnectorEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorEth{HanaConnectorEthCaller: HanaConnectorEthCaller{contract: contract}, HanaConnectorEthTransactor: HanaConnectorEthTransactor{contract: contract}, HanaConnectorEthFilterer: HanaConnectorEthFilterer{contract: contract}}, nil
}

// NewHanaConnectorEthCaller creates a new read-only instance of HanaConnectorEth, bound to a specific deployed contract.
func NewHanaConnectorEthCaller(address common.Address, caller bind.ContractCaller) (*HanaConnectorEthCaller, error) {
	contract, err := bindHanaConnectorEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorEthCaller{contract: contract}, nil
}

// NewHanaConnectorEthTransactor creates a new write-only instance of HanaConnectorEth, bound to a specific deployed contract.
func NewHanaConnectorEthTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaConnectorEthTransactor, error) {
	contract, err := bindHanaConnectorEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorEthTransactor{contract: contract}, nil
}

// NewHanaConnectorEthFilterer creates a new log filterer instance of HanaConnectorEth, bound to a specific deployed contract.
func NewHanaConnectorEthFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaConnectorEthFilterer, error) {
	contract, err := bindHanaConnectorEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorEthFilterer{contract: contract}, nil
}

// bindHanaConnectorEth binds a generic wrapper to an already deployed contract.
func bindHanaConnectorEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaConnectorEthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaConnectorEth *HanaConnectorEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaConnectorEth.Contract.HanaConnectorEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaConnectorEth *HanaConnectorEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.HanaConnectorEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaConnectorEth *HanaConnectorEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.HanaConnectorEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaConnectorEth *HanaConnectorEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaConnectorEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaConnectorEth *HanaConnectorEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaConnectorEth *HanaConnectorEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.contract.Transact(opts, method, params...)
}

// GetLockedAmount is a free data retrieval call binding the contract method 0x252bc886.
//
// Solidity: function getLockedAmount() view returns(uint256)
func (_HanaConnectorEth *HanaConnectorEthCaller) GetLockedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HanaConnectorEth.contract.Call(opts, &out, "getLockedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedAmount is a free data retrieval call binding the contract method 0x252bc886.
//
// Solidity: function getLockedAmount() view returns(uint256)
func (_HanaConnectorEth *HanaConnectorEthSession) GetLockedAmount() (*big.Int, error) {
	return _HanaConnectorEth.Contract.GetLockedAmount(&_HanaConnectorEth.CallOpts)
}

// GetLockedAmount is a free data retrieval call binding the contract method 0x252bc886.
//
// Solidity: function getLockedAmount() view returns(uint256)
func (_HanaConnectorEth *HanaConnectorEthCallerSession) GetLockedAmount() (*big.Int, error) {
	return _HanaConnectorEth.Contract.GetLockedAmount(&_HanaConnectorEth.CallOpts)
}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaConnectorEth *HanaConnectorEthCaller) HanaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorEth.contract.Call(opts, &out, "hanaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaConnectorEth *HanaConnectorEthSession) HanaToken() (common.Address, error) {
	return _HanaConnectorEth.Contract.HanaToken(&_HanaConnectorEth.CallOpts)
}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaConnectorEth *HanaConnectorEthCallerSession) HanaToken() (common.Address, error) {
	return _HanaConnectorEth.Contract.HanaToken(&_HanaConnectorEth.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HanaConnectorEth *HanaConnectorEthCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HanaConnectorEth.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HanaConnectorEth *HanaConnectorEthSession) Paused() (bool, error) {
	return _HanaConnectorEth.Contract.Paused(&_HanaConnectorEth.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HanaConnectorEth *HanaConnectorEthCallerSession) Paused() (bool, error) {
	return _HanaConnectorEth.Contract.Paused(&_HanaConnectorEth.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_HanaConnectorEth *HanaConnectorEthCaller) PauserAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorEth.contract.Call(opts, &out, "pauserAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_HanaConnectorEth *HanaConnectorEthSession) PauserAddress() (common.Address, error) {
	return _HanaConnectorEth.Contract.PauserAddress(&_HanaConnectorEth.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_HanaConnectorEth *HanaConnectorEthCallerSession) PauserAddress() (common.Address, error) {
	return _HanaConnectorEth.Contract.PauserAddress(&_HanaConnectorEth.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_HanaConnectorEth *HanaConnectorEthCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorEth.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_HanaConnectorEth *HanaConnectorEthSession) TssAddress() (common.Address, error) {
	return _HanaConnectorEth.Contract.TssAddress(&_HanaConnectorEth.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_HanaConnectorEth *HanaConnectorEthCallerSession) TssAddress() (common.Address, error) {
	return _HanaConnectorEth.Contract.TssAddress(&_HanaConnectorEth.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_HanaConnectorEth *HanaConnectorEthCaller) TssAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorEth.contract.Call(opts, &out, "tssAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_HanaConnectorEth *HanaConnectorEthSession) TssAddressUpdater() (common.Address, error) {
	return _HanaConnectorEth.Contract.TssAddressUpdater(&_HanaConnectorEth.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_HanaConnectorEth *HanaConnectorEthCallerSession) TssAddressUpdater() (common.Address, error) {
	return _HanaConnectorEth.Contract.TssAddressUpdater(&_HanaConnectorEth.CallOpts)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes hanaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 hanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorEth *HanaConnectorEthTransactor) OnReceive(opts *bind.TransactOpts, hanaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, hanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorEth.contract.Transact(opts, "onReceive", hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes hanaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 hanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorEth *HanaConnectorEthSession) OnReceive(hanaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, hanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.OnReceive(&_HanaConnectorEth.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes hanaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 hanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorEth *HanaConnectorEthTransactorSession) OnReceive(hanaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, hanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.OnReceive(&_HanaConnectorEth.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address hanaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingHanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorEth *HanaConnectorEthTransactor) OnRevert(opts *bind.TransactOpts, hanaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingHanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorEth.contract.Transact(opts, "onRevert", hanaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingHanaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address hanaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingHanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorEth *HanaConnectorEthSession) OnRevert(hanaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingHanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.OnRevert(&_HanaConnectorEth.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingHanaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address hanaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingHanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorEth *HanaConnectorEthTransactorSession) OnRevert(hanaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingHanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.OnRevert(&_HanaConnectorEth.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingHanaValue, message, internalSendHash)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HanaConnectorEth *HanaConnectorEthTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorEth.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HanaConnectorEth *HanaConnectorEthSession) Pause() (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.Pause(&_HanaConnectorEth.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HanaConnectorEth *HanaConnectorEthTransactorSession) Pause() (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.Pause(&_HanaConnectorEth.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_HanaConnectorEth *HanaConnectorEthTransactor) RenounceTssAddressUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorEth.contract.Transact(opts, "renounceTssAddressUpdater")
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_HanaConnectorEth *HanaConnectorEthSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.RenounceTssAddressUpdater(&_HanaConnectorEth.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_HanaConnectorEth *HanaConnectorEthTransactorSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.RenounceTssAddressUpdater(&_HanaConnectorEth.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnectorEth *HanaConnectorEthTransactor) Send(opts *bind.TransactOpts, input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnectorEth.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnectorEth *HanaConnectorEthSession) Send(input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.Send(&_HanaConnectorEth.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnectorEth *HanaConnectorEthTransactorSession) Send(input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.Send(&_HanaConnectorEth.TransactOpts, input)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HanaConnectorEth *HanaConnectorEthTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorEth.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HanaConnectorEth *HanaConnectorEthSession) Unpause() (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.Unpause(&_HanaConnectorEth.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HanaConnectorEth *HanaConnectorEthTransactorSession) Unpause() (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.Unpause(&_HanaConnectorEth.TransactOpts)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_HanaConnectorEth *HanaConnectorEthTransactor) UpdatePauserAddress(opts *bind.TransactOpts, pauserAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorEth.contract.Transact(opts, "updatePauserAddress", pauserAddress_)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_HanaConnectorEth *HanaConnectorEthSession) UpdatePauserAddress(pauserAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.UpdatePauserAddress(&_HanaConnectorEth.TransactOpts, pauserAddress_)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_HanaConnectorEth *HanaConnectorEthTransactorSession) UpdatePauserAddress(pauserAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.UpdatePauserAddress(&_HanaConnectorEth.TransactOpts, pauserAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_HanaConnectorEth *HanaConnectorEthTransactor) UpdateTssAddress(opts *bind.TransactOpts, tssAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorEth.contract.Transact(opts, "updateTssAddress", tssAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_HanaConnectorEth *HanaConnectorEthSession) UpdateTssAddress(tssAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.UpdateTssAddress(&_HanaConnectorEth.TransactOpts, tssAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_HanaConnectorEth *HanaConnectorEthTransactorSession) UpdateTssAddress(tssAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorEth.Contract.UpdateTssAddress(&_HanaConnectorEth.TransactOpts, tssAddress_)
}

// HanaConnectorEthHanaReceivedIterator is returned from FilterHanaReceived and is used to iterate over the raw logs and unpacked data for HanaReceived events raised by the HanaConnectorEth contract.
type HanaConnectorEthHanaReceivedIterator struct {
	Event *HanaConnectorEthHanaReceived // Event containing the contract specifics and raw log

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
func (it *HanaConnectorEthHanaReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorEthHanaReceived)
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
		it.Event = new(HanaConnectorEthHanaReceived)
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
func (it *HanaConnectorEthHanaReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorEthHanaReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorEthHanaReceived represents a HanaReceived event raised by the HanaConnectorEth contract.
type HanaConnectorEthHanaReceived struct {
	HanaTxSenderAddress []byte
	SourceChainId       *big.Int
	DestinationAddress  common.Address
	HanaValue           *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterHanaReceived is a free log retrieval operation binding the contract event 0xb29e5f376a4a399e83dce3b4e7f71bc32f1ab109ccf94574f8f4870cf4aa341f.
//
// Solidity: event HanaReceived(bytes hanaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 hanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorEth *HanaConnectorEthFilterer) FilterHanaReceived(opts *bind.FilterOpts, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (*HanaConnectorEthHanaReceivedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _HanaConnectorEth.contract.FilterLogs(opts, "HanaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorEthHanaReceivedIterator{contract: _HanaConnectorEth.contract, event: "HanaReceived", logs: logs, sub: sub}, nil
}

// WatchHanaReceived is a free log subscription operation binding the contract event 0xb29e5f376a4a399e83dce3b4e7f71bc32f1ab109ccf94574f8f4870cf4aa341f.
//
// Solidity: event HanaReceived(bytes hanaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 hanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorEth *HanaConnectorEthFilterer) WatchHanaReceived(opts *bind.WatchOpts, sink chan<- *HanaConnectorEthHanaReceived, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _HanaConnectorEth.contract.WatchLogs(opts, "HanaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorEthHanaReceived)
				if err := _HanaConnectorEth.contract.UnpackLog(event, "HanaReceived", log); err != nil {
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

// ParseHanaReceived is a log parse operation binding the contract event 0xb29e5f376a4a399e83dce3b4e7f71bc32f1ab109ccf94574f8f4870cf4aa341f.
//
// Solidity: event HanaReceived(bytes hanaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 hanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorEth *HanaConnectorEthFilterer) ParseHanaReceived(log types.Log) (*HanaConnectorEthHanaReceived, error) {
	event := new(HanaConnectorEthHanaReceived)
	if err := _HanaConnectorEth.contract.UnpackLog(event, "HanaReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorEthHanaRevertedIterator is returned from FilterHanaReverted and is used to iterate over the raw logs and unpacked data for HanaReverted events raised by the HanaConnectorEth contract.
type HanaConnectorEthHanaRevertedIterator struct {
	Event *HanaConnectorEthHanaReverted // Event containing the contract specifics and raw log

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
func (it *HanaConnectorEthHanaRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorEthHanaReverted)
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
		it.Event = new(HanaConnectorEthHanaReverted)
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
func (it *HanaConnectorEthHanaRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorEthHanaRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorEthHanaReverted represents a HanaReverted event raised by the HanaConnectorEth contract.
type HanaConnectorEthHanaReverted struct {
	HanaTxSenderAddress common.Address
	SourceChainId       *big.Int
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	RemainingHanaValue  *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterHanaReverted is a free log retrieval operation binding the contract event 0xa0589272400cb97d69e64e37e4dbd1f2aaa79d5b3aa0688e6f2d992ccf500921.
//
// Solidity: event HanaReverted(address hanaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingHanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorEth *HanaConnectorEthFilterer) FilterHanaReverted(opts *bind.FilterOpts, destinationChainId []*big.Int, internalSendHash [][32]byte) (*HanaConnectorEthHanaRevertedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _HanaConnectorEth.contract.FilterLogs(opts, "HanaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorEthHanaRevertedIterator{contract: _HanaConnectorEth.contract, event: "HanaReverted", logs: logs, sub: sub}, nil
}

// WatchHanaReverted is a free log subscription operation binding the contract event 0xa0589272400cb97d69e64e37e4dbd1f2aaa79d5b3aa0688e6f2d992ccf500921.
//
// Solidity: event HanaReverted(address hanaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingHanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorEth *HanaConnectorEthFilterer) WatchHanaReverted(opts *bind.WatchOpts, sink chan<- *HanaConnectorEthHanaReverted, destinationChainId []*big.Int, internalSendHash [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _HanaConnectorEth.contract.WatchLogs(opts, "HanaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorEthHanaReverted)
				if err := _HanaConnectorEth.contract.UnpackLog(event, "HanaReverted", log); err != nil {
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

// ParseHanaReverted is a log parse operation binding the contract event 0xa0589272400cb97d69e64e37e4dbd1f2aaa79d5b3aa0688e6f2d992ccf500921.
//
// Solidity: event HanaReverted(address hanaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingHanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorEth *HanaConnectorEthFilterer) ParseHanaReverted(log types.Log) (*HanaConnectorEthHanaReverted, error) {
	event := new(HanaConnectorEthHanaReverted)
	if err := _HanaConnectorEth.contract.UnpackLog(event, "HanaReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorEthHanaSentIterator is returned from FilterHanaSent and is used to iterate over the raw logs and unpacked data for HanaSent events raised by the HanaConnectorEth contract.
type HanaConnectorEthHanaSentIterator struct {
	Event *HanaConnectorEthHanaSent // Event containing the contract specifics and raw log

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
func (it *HanaConnectorEthHanaSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorEthHanaSent)
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
		it.Event = new(HanaConnectorEthHanaSent)
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
func (it *HanaConnectorEthHanaSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorEthHanaSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorEthHanaSent represents a HanaSent event raised by the HanaConnectorEth contract.
type HanaConnectorEthHanaSent struct {
	SourceTxOriginAddress common.Address
	HanaTxSenderAddress   common.Address
	DestinationChainId    *big.Int
	DestinationAddress    []byte
	HanaValueAndGas       *big.Int
	DestinationGasLimit   *big.Int
	Message               []byte
	HanaParams            []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterHanaSent is a free log retrieval operation binding the contract event 0x7a0dd478962bca8db7a72b684fbfd3404be69978e5f8f1c8eab2cfe4724ebaee.
//
// Solidity: event HanaSent(address sourceTxOriginAddress, address indexed hanaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 hanaValueAndGas, uint256 destinationGasLimit, bytes message, bytes hanaParams)
func (_HanaConnectorEth *HanaConnectorEthFilterer) FilterHanaSent(opts *bind.FilterOpts, hanaTxSenderAddress []common.Address, destinationChainId []*big.Int) (*HanaConnectorEthHanaSentIterator, error) {

	var hanaTxSenderAddressRule []interface{}
	for _, hanaTxSenderAddressItem := range hanaTxSenderAddress {
		hanaTxSenderAddressRule = append(hanaTxSenderAddressRule, hanaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _HanaConnectorEth.contract.FilterLogs(opts, "HanaSent", hanaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorEthHanaSentIterator{contract: _HanaConnectorEth.contract, event: "HanaSent", logs: logs, sub: sub}, nil
}

// WatchHanaSent is a free log subscription operation binding the contract event 0x7a0dd478962bca8db7a72b684fbfd3404be69978e5f8f1c8eab2cfe4724ebaee.
//
// Solidity: event HanaSent(address sourceTxOriginAddress, address indexed hanaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 hanaValueAndGas, uint256 destinationGasLimit, bytes message, bytes hanaParams)
func (_HanaConnectorEth *HanaConnectorEthFilterer) WatchHanaSent(opts *bind.WatchOpts, sink chan<- *HanaConnectorEthHanaSent, hanaTxSenderAddress []common.Address, destinationChainId []*big.Int) (event.Subscription, error) {

	var hanaTxSenderAddressRule []interface{}
	for _, hanaTxSenderAddressItem := range hanaTxSenderAddress {
		hanaTxSenderAddressRule = append(hanaTxSenderAddressRule, hanaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _HanaConnectorEth.contract.WatchLogs(opts, "HanaSent", hanaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorEthHanaSent)
				if err := _HanaConnectorEth.contract.UnpackLog(event, "HanaSent", log); err != nil {
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

// ParseHanaSent is a log parse operation binding the contract event 0x7a0dd478962bca8db7a72b684fbfd3404be69978e5f8f1c8eab2cfe4724ebaee.
//
// Solidity: event HanaSent(address sourceTxOriginAddress, address indexed hanaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 hanaValueAndGas, uint256 destinationGasLimit, bytes message, bytes hanaParams)
func (_HanaConnectorEth *HanaConnectorEthFilterer) ParseHanaSent(log types.Log) (*HanaConnectorEthHanaSent, error) {
	event := new(HanaConnectorEthHanaSent)
	if err := _HanaConnectorEth.contract.UnpackLog(event, "HanaSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorEthPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the HanaConnectorEth contract.
type HanaConnectorEthPausedIterator struct {
	Event *HanaConnectorEthPaused // Event containing the contract specifics and raw log

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
func (it *HanaConnectorEthPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorEthPaused)
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
		it.Event = new(HanaConnectorEthPaused)
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
func (it *HanaConnectorEthPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorEthPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorEthPaused represents a Paused event raised by the HanaConnectorEth contract.
type HanaConnectorEthPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HanaConnectorEth *HanaConnectorEthFilterer) FilterPaused(opts *bind.FilterOpts) (*HanaConnectorEthPausedIterator, error) {

	logs, sub, err := _HanaConnectorEth.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorEthPausedIterator{contract: _HanaConnectorEth.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HanaConnectorEth *HanaConnectorEthFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HanaConnectorEthPaused) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorEth.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorEthPaused)
				if err := _HanaConnectorEth.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HanaConnectorEth *HanaConnectorEthFilterer) ParsePaused(log types.Log) (*HanaConnectorEthPaused, error) {
	event := new(HanaConnectorEthPaused)
	if err := _HanaConnectorEth.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorEthPauserAddressUpdatedIterator is returned from FilterPauserAddressUpdated and is used to iterate over the raw logs and unpacked data for PauserAddressUpdated events raised by the HanaConnectorEth contract.
type HanaConnectorEthPauserAddressUpdatedIterator struct {
	Event *HanaConnectorEthPauserAddressUpdated // Event containing the contract specifics and raw log

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
func (it *HanaConnectorEthPauserAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorEthPauserAddressUpdated)
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
		it.Event = new(HanaConnectorEthPauserAddressUpdated)
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
func (it *HanaConnectorEthPauserAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorEthPauserAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorEthPauserAddressUpdated represents a PauserAddressUpdated event raised by the HanaConnectorEth contract.
type HanaConnectorEthPauserAddressUpdated struct {
	UpdaterAddress common.Address
	NewTssAddress  common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPauserAddressUpdated is a free log retrieval operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address updaterAddress, address newTssAddress)
func (_HanaConnectorEth *HanaConnectorEthFilterer) FilterPauserAddressUpdated(opts *bind.FilterOpts) (*HanaConnectorEthPauserAddressUpdatedIterator, error) {

	logs, sub, err := _HanaConnectorEth.contract.FilterLogs(opts, "PauserAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorEthPauserAddressUpdatedIterator{contract: _HanaConnectorEth.contract, event: "PauserAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchPauserAddressUpdated is a free log subscription operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address updaterAddress, address newTssAddress)
func (_HanaConnectorEth *HanaConnectorEthFilterer) WatchPauserAddressUpdated(opts *bind.WatchOpts, sink chan<- *HanaConnectorEthPauserAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorEth.contract.WatchLogs(opts, "PauserAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorEthPauserAddressUpdated)
				if err := _HanaConnectorEth.contract.UnpackLog(event, "PauserAddressUpdated", log); err != nil {
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

// ParsePauserAddressUpdated is a log parse operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address updaterAddress, address newTssAddress)
func (_HanaConnectorEth *HanaConnectorEthFilterer) ParsePauserAddressUpdated(log types.Log) (*HanaConnectorEthPauserAddressUpdated, error) {
	event := new(HanaConnectorEthPauserAddressUpdated)
	if err := _HanaConnectorEth.contract.UnpackLog(event, "PauserAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorEthTSSAddressUpdatedIterator is returned from FilterTSSAddressUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdated events raised by the HanaConnectorEth contract.
type HanaConnectorEthTSSAddressUpdatedIterator struct {
	Event *HanaConnectorEthTSSAddressUpdated // Event containing the contract specifics and raw log

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
func (it *HanaConnectorEthTSSAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorEthTSSAddressUpdated)
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
		it.Event = new(HanaConnectorEthTSSAddressUpdated)
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
func (it *HanaConnectorEthTSSAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorEthTSSAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorEthTSSAddressUpdated represents a TSSAddressUpdated event raised by the HanaConnectorEth contract.
type HanaConnectorEthTSSAddressUpdated struct {
	HanaTxSenderAddress common.Address
	NewTssAddress       common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdated is a free log retrieval operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address hanaTxSenderAddress, address newTssAddress)
func (_HanaConnectorEth *HanaConnectorEthFilterer) FilterTSSAddressUpdated(opts *bind.FilterOpts) (*HanaConnectorEthTSSAddressUpdatedIterator, error) {

	logs, sub, err := _HanaConnectorEth.contract.FilterLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorEthTSSAddressUpdatedIterator{contract: _HanaConnectorEth.contract, event: "TSSAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdated is a free log subscription operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address hanaTxSenderAddress, address newTssAddress)
func (_HanaConnectorEth *HanaConnectorEthFilterer) WatchTSSAddressUpdated(opts *bind.WatchOpts, sink chan<- *HanaConnectorEthTSSAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorEth.contract.WatchLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorEthTSSAddressUpdated)
				if err := _HanaConnectorEth.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
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

// ParseTSSAddressUpdated is a log parse operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address hanaTxSenderAddress, address newTssAddress)
func (_HanaConnectorEth *HanaConnectorEthFilterer) ParseTSSAddressUpdated(log types.Log) (*HanaConnectorEthTSSAddressUpdated, error) {
	event := new(HanaConnectorEthTSSAddressUpdated)
	if err := _HanaConnectorEth.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorEthUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the HanaConnectorEth contract.
type HanaConnectorEthUnpausedIterator struct {
	Event *HanaConnectorEthUnpaused // Event containing the contract specifics and raw log

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
func (it *HanaConnectorEthUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorEthUnpaused)
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
		it.Event = new(HanaConnectorEthUnpaused)
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
func (it *HanaConnectorEthUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorEthUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorEthUnpaused represents a Unpaused event raised by the HanaConnectorEth contract.
type HanaConnectorEthUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HanaConnectorEth *HanaConnectorEthFilterer) FilterUnpaused(opts *bind.FilterOpts) (*HanaConnectorEthUnpausedIterator, error) {

	logs, sub, err := _HanaConnectorEth.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorEthUnpausedIterator{contract: _HanaConnectorEth.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HanaConnectorEth *HanaConnectorEthFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HanaConnectorEthUnpaused) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorEth.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorEthUnpaused)
				if err := _HanaConnectorEth.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HanaConnectorEth *HanaConnectorEthFilterer) ParseUnpaused(log types.Log) (*HanaConnectorEthUnpaused, error) {
	event := new(HanaConnectorEthUnpaused)
	if err := _HanaConnectorEth.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
