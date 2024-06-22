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

// HanaConnectorNonEthMetaData contains all meta data concerning the HanaConnectorNonEth contract.
var HanaConnectorNonEthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hanaTokenAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tssAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tssAddressUpdater_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauserAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTss\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssOrUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HanaTranserError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hanaTxSenderAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hanaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"HanaReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingHanaValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"HanaReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceTxOriginAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hanaValueAndGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hanaParams\",\"type\":\"bytes\"}],\"name\":\"HanaSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updaterAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTssAddress\",\"type\":\"address\"}],\"name\":\"PauserAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTssAddress\",\"type\":\"address\"}],\"name\":\"TSSAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLockedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hanaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"hanaTxSenderAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"onReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hanaTxSenderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingHanaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"onRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceTssAddressUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"hanaValueAndGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hanaParams\",\"type\":\"bytes\"}],\"internalType\":\"structHanaInterfaces.SendInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply_\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauserAddress_\",\"type\":\"address\"}],\"name\":\"updatePauserAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tssAddress_\",\"type\":\"address\"}],\"name\":\"updateTssAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6003553480156200003557600080fd5b50604051620022e7380380620022e783398181016040528101906200005b9190620002a8565b8383838360008060006101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480620000e15750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b80620001195750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b80620001515750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b1562000189576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b8152505082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050505050506200036d565b600081519050620002a28162000353565b92915050565b60008060008060808587031215620002c557620002c46200034e565b5b6000620002d58782880162000291565b9450506020620002e88782880162000291565b9350506040620002fb8782880162000291565b92505060606200030e8782880162000291565b91505092959194509250565b600062000327826200032e565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200035e816200031a565b81146200036a57600080fd5b50565b60805160601c611f31620003b6600039600081816102a1015281816103e4015281816104d2015281816107fd01528181610ebe01528181610fac01526111db0152611f316000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80636f8b44b011610097578063942a5e1611610066578063942a5e1614610229578063d5abeb0114610245578063ec02690114610263578063f7fb869b1461027f57610100565b80636f8b44b0146101dd578063779e3b63146101f95780638456cb59146102035780639122c3441461020d57610100565b80635b112591116100d35780635b112591146101675780635c975abb146101855780635e694a92146101a35780636128480f146101c157610100565b8063252bc8861461010557806329dd214d14610123578063328a01d01461013f5780633f4ba83a1461015d575b600080fd5b61010d61029d565b60405161011a9190611c51565b60405180910390f35b61013d6004803603810190610138919061165f565b61034d565b005b6101476106fd565b60405161015491906119e4565b60405180910390f35b610165610723565b005b61016f6107bf565b60405161017c91906119e4565b60405180910390f35b61018d6107e5565b60405161019a9190611b69565b60405180910390f35b6101ab6107fb565b6040516101b891906119e4565b60405180910390f35b6101db60048036038101906101d69190611550565b61081f565b005b6101f760048036038101906101f29190611777565b610995565b005b610201610a31565b005b61020b610bb1565b005b61022760048036038101906102229190611550565b610c4d565b005b610243600480360381019061023e919061157d565b610e1f565b005b61024d6111cb565b60405161025a9190611c51565b60405180910390f35b61027d6004803603810190610278919061172e565b6111d1565b005b610287611302565b60405161029491906119e4565b60405180910390f35b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016102f891906119e4565b60206040518083038186803b15801561031057600080fd5b505afa158015610324573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034891906117a4565b905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103df57336040517fff70ace20000000000000000000000000000000000000000000000000000000081526004016103d691906119e4565b60405180910390fd5b6003547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561044857600080fd5b505afa15801561045c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061048091906117a4565b8561048b9190611d0d565b11156104d0576003546040517f3d3dbc830000000000000000000000000000000000000000000000000000000081526004016104c79190611c51565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631e458bee8686846040518463ffffffff1660e01b815260040161052d93929190611acd565b600060405180830381600087803b15801561054757600080fd5b505af115801561055b573d6000803e3d6000fd5b50505050600083839050111561069b578473ffffffffffffffffffffffffffffffffffffffff1663b204be936040518060a001604052808b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018981526020018873ffffffffffffffffffffffffffffffffffffffff16815260200187815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152506040518263ffffffff1660e01b81526004016106689190611c0d565b600060405180830381600087803b15801561068257600080fd5b505af1158015610696573d6000803e3d6000fd5b505050505b808573ffffffffffffffffffffffffffffffffffffffff16877fb29e5f376a4a399e83dce3b4e7f71bc32f1ab109ccf94574f8f4870cf4aa341f8b8b8989896040516106eb959493929190611b84565b60405180910390a45050505050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107b557336040517f4677a0d30000000000000000000000000000000000000000000000000000000081526004016107ac91906119e4565b60405180910390fd5b6107bd611328565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900460ff16905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108b157336040517f4677a0d30000000000000000000000000000000000000000000000000000000081526004016108a891906119e4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610918576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397338260405161098a9291906119ff565b60405180910390a150565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a2757336040517fff70ace2000000000000000000000000000000000000000000000000000000008152600401610a1e91906119e4565b60405180910390fd5b8060038190555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac357336040517fe700765e000000000000000000000000000000000000000000000000000000008152600401610aba91906119e4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610b4c576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c4357336040517f4677a0d3000000000000000000000000000000000000000000000000000000008152600401610c3a91906119e4565b60405180910390fd5b610c4b61138a565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610cf95750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610d3b57336040517fcdfcef97000000000000000000000000000000000000000000000000000000008152600401610d3291906119e4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610da2576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff3382604051610e149291906119ff565b60405180910390a150565b610e276113ec565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610eb957336040517fff70ace2000000000000000000000000000000000000000000000000000000008152600401610eb091906119e4565b60405180910390fd5b6003547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610f2257600080fd5b505afa158015610f36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5a91906117a4565b85610f659190611d0d565b1115610faa576003546040517f3d3dbc83000000000000000000000000000000000000000000000000000000008152600401610fa19190611c51565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631e458bee8a86846040518463ffffffff1660e01b815260040161100793929190611acd565b600060405180830381600087803b15801561102157600080fd5b505af1158015611035573d6000803e3d6000fd5b50505050600083839050111561117b578873ffffffffffffffffffffffffffffffffffffffff16634204cf9b6040518060c001604052808c73ffffffffffffffffffffffffffffffffffffffff1681526020018b81526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200188815260200187815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152506040518263ffffffff1660e01b81526004016111489190611c2f565b600060405180830381600087803b15801561116257600080fd5b505af1158015611176573d6000803e3d6000fd5b505050505b80857fa0589272400cb97d69e64e37e4dbd1f2aaa79d5b3aa0688e6f2d992ccf5009218b8b8b8b8a8a8a6040516111b89796959493929190611b04565b60405180910390a3505050505050505050565b60035481565b6111d96113ec565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379cc67903383608001356040518363ffffffff1660e01b8152600401611238929190611aa4565b600060405180830381600087803b15801561125257600080fd5b505af1158015611266573d6000803e3d6000fd5b5050505080600001353373ffffffffffffffffffffffffffffffffffffffff167f7a0dd478962bca8db7a72b684fbfd3404be69978e5f8f1c8eab2cfe4724ebaee328480602001906112b89190611c6c565b866080013587604001358880606001906112d29190611c6c565b8a8060a001906112e29190611c6c565b6040516112f799989796959493929190611a28565b60405180910390a350565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611330611436565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61137361147f565b60405161138091906119e4565b60405180910390a1565b6113926113ec565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586113d561147f565b6040516113e291906119e4565b60405180910390a1565b6113f46107e5565b15611434576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161142b90611bed565b60405180910390fd5b565b61143e6107e5565b61147d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161147490611bcd565b60405180910390fd5b565b600033905090565b60008135905061149681611eb6565b92915050565b6000813590506114ab81611ecd565b92915050565b60008083601f8401126114c7576114c6611e2b565b5b8235905067ffffffffffffffff8111156114e4576114e3611e26565b5b602083019150836001820283011115611500576114ff611e3f565b5b9250929050565b600060c0828403121561151d5761151c611e35565b5b81905092915050565b60008135905061153581611ee4565b92915050565b60008151905061154a81611ee4565b92915050565b60006020828403121561156657611565611e4e565b5b600061157484828501611487565b91505092915050565b600080600080600080600080600060e08a8c03121561159f5761159e611e4e565b5b60006115ad8c828d01611487565b99505060206115be8c828d01611526565b98505060408a013567ffffffffffffffff8111156115df576115de611e49565b5b6115eb8c828d016114b1565b975097505060606115fe8c828d01611526565b955050608061160f8c828d01611526565b94505060a08a013567ffffffffffffffff8111156116305761162f611e49565b5b61163c8c828d016114b1565b935093505060c061164f8c828d0161149c565b9150509295985092959850929598565b60008060008060008060008060c0898b03121561167f5761167e611e4e565b5b600089013567ffffffffffffffff81111561169d5761169c611e49565b5b6116a98b828c016114b1565b985098505060206116bc8b828c01611526565b96505060406116cd8b828c01611487565b95505060606116de8b828c01611526565b945050608089013567ffffffffffffffff8111156116ff576116fe611e49565b5b61170b8b828c016114b1565b935093505060a061171e8b828c0161149c565b9150509295985092959890939650565b60006020828403121561174457611743611e4e565b5b600082013567ffffffffffffffff81111561176257611761611e49565b5b61176e84828501611507565b91505092915050565b60006020828403121561178d5761178c611e4e565b5b600061179b84828501611526565b91505092915050565b6000602082840312156117ba576117b9611e4e565b5b60006117c88482850161153b565b91505092915050565b6117da81611d63565b82525050565b6117e981611d63565b82525050565b6117f881611d75565b82525050565b61180781611d81565b82525050565b60006118198385611ceb565b9350611826838584611db5565b61182f83611e53565b840190509392505050565b600061184582611ccf565b61184f8185611cda565b935061185f818560208601611dc4565b61186881611e53565b840191505092915050565b6000611880601483611cfc565b915061188b82611e64565b602082019050919050565b60006118a3601083611cfc565b91506118ae82611e8d565b602082019050919050565b600060a08301600083015184820360008601526118d6828261183a565b91505060208301516118eb60208601826119c6565b5060408301516118fe60408601826117d1565b50606083015161191160608601826119c6565b5060808301518482036080860152611929828261183a565b9150508091505092915050565b600060c08301600083015161194e60008601826117d1565b50602083015161196160208601826119c6565b5060408301518482036040860152611979828261183a565b915050606083015161198e60608601826119c6565b5060808301516119a160808601826119c6565b5060a083015184820360a08601526119b9828261183a565b9150508091505092915050565b6119cf81611dab565b82525050565b6119de81611dab565b82525050565b60006020820190506119f960008301846117e0565b92915050565b6000604082019050611a1460008301856117e0565b611a2160208301846117e0565b9392505050565b600060c082019050611a3d600083018c6117e0565b8181036020830152611a50818a8c61180d565b9050611a5f60408301896119d5565b611a6c60608301886119d5565b8181036080830152611a7f81868861180d565b905081810360a0830152611a9481848661180d565b90509a9950505050505050505050565b6000604082019050611ab960008301856117e0565b611ac660208301846119d5565b9392505050565b6000606082019050611ae260008301866117e0565b611aef60208301856119d5565b611afc60408301846117fe565b949350505050565b600060a082019050611b19600083018a6117e0565b611b2660208301896119d5565b8181036040830152611b3981878961180d565b9050611b4860608301866119d5565b8181036080830152611b5b81848661180d565b905098975050505050505050565b6000602082019050611b7e60008301846117ef565b92915050565b60006060820190508181036000830152611b9f81878961180d565b9050611bae60208301866119d5565b8181036040830152611bc181848661180d565b90509695505050505050565b60006020820190508181036000830152611be681611873565b9050919050565b60006020820190508181036000830152611c0681611896565b9050919050565b60006020820190508181036000830152611c2781846118b9565b905092915050565b60006020820190508181036000830152611c498184611936565b905092915050565b6000602082019050611c6660008301846119d5565b92915050565b60008083356001602003843603038112611c8957611c88611e3a565b5b80840192508235915067ffffffffffffffff821115611cab57611caa611e30565b5b602083019250600182023603831315611cc757611cc6611e44565b5b509250929050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611d1882611dab565b9150611d2383611dab565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611d5857611d57611df7565b5b828201905092915050565b6000611d6e82611d8b565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611de2578082015181840152602081019050611dc7565b83811115611df1576000848401525b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b611ebf81611d63565b8114611eca57600080fd5b50565b611ed681611d81565b8114611ee157600080fd5b50565b611eed81611dab565b8114611ef857600080fd5b5056fea2646970667358221220776953d34304b5d912bbd65b2826d35d4e7b56993dfb086feb364262285d143864736f6c63430008070033",
}

// HanaConnectorNonEthABI is the input ABI used to generate the binding from.
// Deprecated: Use HanaConnectorNonEthMetaData.ABI instead.
var HanaConnectorNonEthABI = HanaConnectorNonEthMetaData.ABI

// HanaConnectorNonEthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HanaConnectorNonEthMetaData.Bin instead.
var HanaConnectorNonEthBin = HanaConnectorNonEthMetaData.Bin

// DeployHanaConnectorNonEth deploys a new Ethereum contract, binding an instance of HanaConnectorNonEth to it.
func DeployHanaConnectorNonEth(auth *bind.TransactOpts, backend bind.ContractBackend, hanaTokenAddress_ common.Address, tssAddress_ common.Address, tssAddressUpdater_ common.Address, pauserAddress_ common.Address) (common.Address, *types.Transaction, *HanaConnectorNonEth, error) {
	parsed, err := HanaConnectorNonEthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HanaConnectorNonEthBin), backend, hanaTokenAddress_, tssAddress_, tssAddressUpdater_, pauserAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HanaConnectorNonEth{HanaConnectorNonEthCaller: HanaConnectorNonEthCaller{contract: contract}, HanaConnectorNonEthTransactor: HanaConnectorNonEthTransactor{contract: contract}, HanaConnectorNonEthFilterer: HanaConnectorNonEthFilterer{contract: contract}}, nil
}

// HanaConnectorNonEth is an auto generated Go binding around an Ethereum contract.
type HanaConnectorNonEth struct {
	HanaConnectorNonEthCaller     // Read-only binding to the contract
	HanaConnectorNonEthTransactor // Write-only binding to the contract
	HanaConnectorNonEthFilterer   // Log filterer for contract events
}

// HanaConnectorNonEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type HanaConnectorNonEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorNonEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HanaConnectorNonEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorNonEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HanaConnectorNonEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HanaConnectorNonEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HanaConnectorNonEthSession struct {
	Contract     *HanaConnectorNonEth // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HanaConnectorNonEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HanaConnectorNonEthCallerSession struct {
	Contract *HanaConnectorNonEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// HanaConnectorNonEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HanaConnectorNonEthTransactorSession struct {
	Contract     *HanaConnectorNonEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// HanaConnectorNonEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type HanaConnectorNonEthRaw struct {
	Contract *HanaConnectorNonEth // Generic contract binding to access the raw methods on
}

// HanaConnectorNonEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HanaConnectorNonEthCallerRaw struct {
	Contract *HanaConnectorNonEthCaller // Generic read-only contract binding to access the raw methods on
}

// HanaConnectorNonEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HanaConnectorNonEthTransactorRaw struct {
	Contract *HanaConnectorNonEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHanaConnectorNonEth creates a new instance of HanaConnectorNonEth, bound to a specific deployed contract.
func NewHanaConnectorNonEth(address common.Address, backend bind.ContractBackend) (*HanaConnectorNonEth, error) {
	contract, err := bindHanaConnectorNonEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorNonEth{HanaConnectorNonEthCaller: HanaConnectorNonEthCaller{contract: contract}, HanaConnectorNonEthTransactor: HanaConnectorNonEthTransactor{contract: contract}, HanaConnectorNonEthFilterer: HanaConnectorNonEthFilterer{contract: contract}}, nil
}

// NewHanaConnectorNonEthCaller creates a new read-only instance of HanaConnectorNonEth, bound to a specific deployed contract.
func NewHanaConnectorNonEthCaller(address common.Address, caller bind.ContractCaller) (*HanaConnectorNonEthCaller, error) {
	contract, err := bindHanaConnectorNonEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorNonEthCaller{contract: contract}, nil
}

// NewHanaConnectorNonEthTransactor creates a new write-only instance of HanaConnectorNonEth, bound to a specific deployed contract.
func NewHanaConnectorNonEthTransactor(address common.Address, transactor bind.ContractTransactor) (*HanaConnectorNonEthTransactor, error) {
	contract, err := bindHanaConnectorNonEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorNonEthTransactor{contract: contract}, nil
}

// NewHanaConnectorNonEthFilterer creates a new log filterer instance of HanaConnectorNonEth, bound to a specific deployed contract.
func NewHanaConnectorNonEthFilterer(address common.Address, filterer bind.ContractFilterer) (*HanaConnectorNonEthFilterer, error) {
	contract, err := bindHanaConnectorNonEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorNonEthFilterer{contract: contract}, nil
}

// bindHanaConnectorNonEth binds a generic wrapper to an already deployed contract.
func bindHanaConnectorNonEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HanaConnectorNonEthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaConnectorNonEth *HanaConnectorNonEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaConnectorNonEth.Contract.HanaConnectorNonEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaConnectorNonEth *HanaConnectorNonEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.HanaConnectorNonEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaConnectorNonEth *HanaConnectorNonEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.HanaConnectorNonEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HanaConnectorNonEth *HanaConnectorNonEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HanaConnectorNonEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.contract.Transact(opts, method, params...)
}

// GetLockedAmount is a free data retrieval call binding the contract method 0x252bc886.
//
// Solidity: function getLockedAmount() view returns(uint256)
func (_HanaConnectorNonEth *HanaConnectorNonEthCaller) GetLockedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HanaConnectorNonEth.contract.Call(opts, &out, "getLockedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedAmount is a free data retrieval call binding the contract method 0x252bc886.
//
// Solidity: function getLockedAmount() view returns(uint256)
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) GetLockedAmount() (*big.Int, error) {
	return _HanaConnectorNonEth.Contract.GetLockedAmount(&_HanaConnectorNonEth.CallOpts)
}

// GetLockedAmount is a free data retrieval call binding the contract method 0x252bc886.
//
// Solidity: function getLockedAmount() view returns(uint256)
func (_HanaConnectorNonEth *HanaConnectorNonEthCallerSession) GetLockedAmount() (*big.Int, error) {
	return _HanaConnectorNonEth.Contract.GetLockedAmount(&_HanaConnectorNonEth.CallOpts)
}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaConnectorNonEth *HanaConnectorNonEthCaller) HanaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorNonEth.contract.Call(opts, &out, "hanaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) HanaToken() (common.Address, error) {
	return _HanaConnectorNonEth.Contract.HanaToken(&_HanaConnectorNonEth.CallOpts)
}

// HanaToken is a free data retrieval call binding the contract method 0x5e694a92.
//
// Solidity: function hanaToken() view returns(address)
func (_HanaConnectorNonEth *HanaConnectorNonEthCallerSession) HanaToken() (common.Address, error) {
	return _HanaConnectorNonEth.Contract.HanaToken(&_HanaConnectorNonEth.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_HanaConnectorNonEth *HanaConnectorNonEthCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HanaConnectorNonEth.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) MaxSupply() (*big.Int, error) {
	return _HanaConnectorNonEth.Contract.MaxSupply(&_HanaConnectorNonEth.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_HanaConnectorNonEth *HanaConnectorNonEthCallerSession) MaxSupply() (*big.Int, error) {
	return _HanaConnectorNonEth.Contract.MaxSupply(&_HanaConnectorNonEth.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HanaConnectorNonEth *HanaConnectorNonEthCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HanaConnectorNonEth.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) Paused() (bool, error) {
	return _HanaConnectorNonEth.Contract.Paused(&_HanaConnectorNonEth.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HanaConnectorNonEth *HanaConnectorNonEthCallerSession) Paused() (bool, error) {
	return _HanaConnectorNonEth.Contract.Paused(&_HanaConnectorNonEth.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_HanaConnectorNonEth *HanaConnectorNonEthCaller) PauserAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorNonEth.contract.Call(opts, &out, "pauserAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) PauserAddress() (common.Address, error) {
	return _HanaConnectorNonEth.Contract.PauserAddress(&_HanaConnectorNonEth.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_HanaConnectorNonEth *HanaConnectorNonEthCallerSession) PauserAddress() (common.Address, error) {
	return _HanaConnectorNonEth.Contract.PauserAddress(&_HanaConnectorNonEth.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_HanaConnectorNonEth *HanaConnectorNonEthCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorNonEth.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) TssAddress() (common.Address, error) {
	return _HanaConnectorNonEth.Contract.TssAddress(&_HanaConnectorNonEth.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_HanaConnectorNonEth *HanaConnectorNonEthCallerSession) TssAddress() (common.Address, error) {
	return _HanaConnectorNonEth.Contract.TssAddress(&_HanaConnectorNonEth.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_HanaConnectorNonEth *HanaConnectorNonEthCaller) TssAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HanaConnectorNonEth.contract.Call(opts, &out, "tssAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) TssAddressUpdater() (common.Address, error) {
	return _HanaConnectorNonEth.Contract.TssAddressUpdater(&_HanaConnectorNonEth.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_HanaConnectorNonEth *HanaConnectorNonEthCallerSession) TssAddressUpdater() (common.Address, error) {
	return _HanaConnectorNonEth.Contract.TssAddressUpdater(&_HanaConnectorNonEth.CallOpts)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes hanaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 hanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactor) OnReceive(opts *bind.TransactOpts, hanaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, hanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorNonEth.contract.Transact(opts, "onReceive", hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes hanaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 hanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) OnReceive(hanaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, hanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.OnReceive(&_HanaConnectorNonEth.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes hanaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 hanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactorSession) OnReceive(hanaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, hanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.OnReceive(&_HanaConnectorNonEth.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, hanaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address hanaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingHanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactor) OnRevert(opts *bind.TransactOpts, hanaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingHanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorNonEth.contract.Transact(opts, "onRevert", hanaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingHanaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address hanaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingHanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) OnRevert(hanaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingHanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.OnRevert(&_HanaConnectorNonEth.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingHanaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address hanaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingHanaValue, bytes message, bytes32 internalSendHash) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactorSession) OnRevert(hanaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingHanaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.OnRevert(&_HanaConnectorNonEth.TransactOpts, hanaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingHanaValue, message, internalSendHash)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorNonEth.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) Pause() (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.Pause(&_HanaConnectorNonEth.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactorSession) Pause() (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.Pause(&_HanaConnectorNonEth.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactor) RenounceTssAddressUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorNonEth.contract.Transact(opts, "renounceTssAddressUpdater")
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.RenounceTssAddressUpdater(&_HanaConnectorNonEth.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactorSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.RenounceTssAddressUpdater(&_HanaConnectorNonEth.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactor) Send(opts *bind.TransactOpts, input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnectorNonEth.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) Send(input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.Send(&_HanaConnectorNonEth.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactorSession) Send(input HanaInterfacesSendInput) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.Send(&_HanaConnectorNonEth.TransactOpts, input)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactor) SetMaxSupply(opts *bind.TransactOpts, maxSupply_ *big.Int) (*types.Transaction, error) {
	return _HanaConnectorNonEth.contract.Transact(opts, "setMaxSupply", maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.SetMaxSupply(&_HanaConnectorNonEth.TransactOpts, maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactorSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.SetMaxSupply(&_HanaConnectorNonEth.TransactOpts, maxSupply_)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HanaConnectorNonEth.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) Unpause() (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.Unpause(&_HanaConnectorNonEth.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactorSession) Unpause() (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.Unpause(&_HanaConnectorNonEth.TransactOpts)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactor) UpdatePauserAddress(opts *bind.TransactOpts, pauserAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorNonEth.contract.Transact(opts, "updatePauserAddress", pauserAddress_)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) UpdatePauserAddress(pauserAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.UpdatePauserAddress(&_HanaConnectorNonEth.TransactOpts, pauserAddress_)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactorSession) UpdatePauserAddress(pauserAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.UpdatePauserAddress(&_HanaConnectorNonEth.TransactOpts, pauserAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactor) UpdateTssAddress(opts *bind.TransactOpts, tssAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorNonEth.contract.Transact(opts, "updateTssAddress", tssAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthSession) UpdateTssAddress(tssAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.UpdateTssAddress(&_HanaConnectorNonEth.TransactOpts, tssAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_HanaConnectorNonEth *HanaConnectorNonEthTransactorSession) UpdateTssAddress(tssAddress_ common.Address) (*types.Transaction, error) {
	return _HanaConnectorNonEth.Contract.UpdateTssAddress(&_HanaConnectorNonEth.TransactOpts, tssAddress_)
}

// HanaConnectorNonEthHanaReceivedIterator is returned from FilterHanaReceived and is used to iterate over the raw logs and unpacked data for HanaReceived events raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthHanaReceivedIterator struct {
	Event *HanaConnectorNonEthHanaReceived // Event containing the contract specifics and raw log

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
func (it *HanaConnectorNonEthHanaReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorNonEthHanaReceived)
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
		it.Event = new(HanaConnectorNonEthHanaReceived)
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
func (it *HanaConnectorNonEthHanaReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorNonEthHanaReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorNonEthHanaReceived represents a HanaReceived event raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthHanaReceived struct {
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
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) FilterHanaReceived(opts *bind.FilterOpts, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (*HanaConnectorNonEthHanaReceivedIterator, error) {

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

	logs, sub, err := _HanaConnectorNonEth.contract.FilterLogs(opts, "HanaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorNonEthHanaReceivedIterator{contract: _HanaConnectorNonEth.contract, event: "HanaReceived", logs: logs, sub: sub}, nil
}

// WatchHanaReceived is a free log subscription operation binding the contract event 0xb29e5f376a4a399e83dce3b4e7f71bc32f1ab109ccf94574f8f4870cf4aa341f.
//
// Solidity: event HanaReceived(bytes hanaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 hanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) WatchHanaReceived(opts *bind.WatchOpts, sink chan<- *HanaConnectorNonEthHanaReceived, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HanaConnectorNonEth.contract.WatchLogs(opts, "HanaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorNonEthHanaReceived)
				if err := _HanaConnectorNonEth.contract.UnpackLog(event, "HanaReceived", log); err != nil {
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
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) ParseHanaReceived(log types.Log) (*HanaConnectorNonEthHanaReceived, error) {
	event := new(HanaConnectorNonEthHanaReceived)
	if err := _HanaConnectorNonEth.contract.UnpackLog(event, "HanaReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorNonEthHanaRevertedIterator is returned from FilterHanaReverted and is used to iterate over the raw logs and unpacked data for HanaReverted events raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthHanaRevertedIterator struct {
	Event *HanaConnectorNonEthHanaReverted // Event containing the contract specifics and raw log

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
func (it *HanaConnectorNonEthHanaRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorNonEthHanaReverted)
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
		it.Event = new(HanaConnectorNonEthHanaReverted)
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
func (it *HanaConnectorNonEthHanaRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorNonEthHanaRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorNonEthHanaReverted represents a HanaReverted event raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthHanaReverted struct {
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
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) FilterHanaReverted(opts *bind.FilterOpts, destinationChainId []*big.Int, internalSendHash [][32]byte) (*HanaConnectorNonEthHanaRevertedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _HanaConnectorNonEth.contract.FilterLogs(opts, "HanaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorNonEthHanaRevertedIterator{contract: _HanaConnectorNonEth.contract, event: "HanaReverted", logs: logs, sub: sub}, nil
}

// WatchHanaReverted is a free log subscription operation binding the contract event 0xa0589272400cb97d69e64e37e4dbd1f2aaa79d5b3aa0688e6f2d992ccf500921.
//
// Solidity: event HanaReverted(address hanaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingHanaValue, bytes message, bytes32 indexed internalSendHash)
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) WatchHanaReverted(opts *bind.WatchOpts, sink chan<- *HanaConnectorNonEthHanaReverted, destinationChainId []*big.Int, internalSendHash [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _HanaConnectorNonEth.contract.WatchLogs(opts, "HanaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorNonEthHanaReverted)
				if err := _HanaConnectorNonEth.contract.UnpackLog(event, "HanaReverted", log); err != nil {
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
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) ParseHanaReverted(log types.Log) (*HanaConnectorNonEthHanaReverted, error) {
	event := new(HanaConnectorNonEthHanaReverted)
	if err := _HanaConnectorNonEth.contract.UnpackLog(event, "HanaReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorNonEthHanaSentIterator is returned from FilterHanaSent and is used to iterate over the raw logs and unpacked data for HanaSent events raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthHanaSentIterator struct {
	Event *HanaConnectorNonEthHanaSent // Event containing the contract specifics and raw log

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
func (it *HanaConnectorNonEthHanaSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorNonEthHanaSent)
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
		it.Event = new(HanaConnectorNonEthHanaSent)
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
func (it *HanaConnectorNonEthHanaSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorNonEthHanaSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorNonEthHanaSent represents a HanaSent event raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthHanaSent struct {
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
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) FilterHanaSent(opts *bind.FilterOpts, hanaTxSenderAddress []common.Address, destinationChainId []*big.Int) (*HanaConnectorNonEthHanaSentIterator, error) {

	var hanaTxSenderAddressRule []interface{}
	for _, hanaTxSenderAddressItem := range hanaTxSenderAddress {
		hanaTxSenderAddressRule = append(hanaTxSenderAddressRule, hanaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _HanaConnectorNonEth.contract.FilterLogs(opts, "HanaSent", hanaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &HanaConnectorNonEthHanaSentIterator{contract: _HanaConnectorNonEth.contract, event: "HanaSent", logs: logs, sub: sub}, nil
}

// WatchHanaSent is a free log subscription operation binding the contract event 0x7a0dd478962bca8db7a72b684fbfd3404be69978e5f8f1c8eab2cfe4724ebaee.
//
// Solidity: event HanaSent(address sourceTxOriginAddress, address indexed hanaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 hanaValueAndGas, uint256 destinationGasLimit, bytes message, bytes hanaParams)
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) WatchHanaSent(opts *bind.WatchOpts, sink chan<- *HanaConnectorNonEthHanaSent, hanaTxSenderAddress []common.Address, destinationChainId []*big.Int) (event.Subscription, error) {

	var hanaTxSenderAddressRule []interface{}
	for _, hanaTxSenderAddressItem := range hanaTxSenderAddress {
		hanaTxSenderAddressRule = append(hanaTxSenderAddressRule, hanaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _HanaConnectorNonEth.contract.WatchLogs(opts, "HanaSent", hanaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorNonEthHanaSent)
				if err := _HanaConnectorNonEth.contract.UnpackLog(event, "HanaSent", log); err != nil {
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
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) ParseHanaSent(log types.Log) (*HanaConnectorNonEthHanaSent, error) {
	event := new(HanaConnectorNonEthHanaSent)
	if err := _HanaConnectorNonEth.contract.UnpackLog(event, "HanaSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorNonEthPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthPausedIterator struct {
	Event *HanaConnectorNonEthPaused // Event containing the contract specifics and raw log

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
func (it *HanaConnectorNonEthPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorNonEthPaused)
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
		it.Event = new(HanaConnectorNonEthPaused)
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
func (it *HanaConnectorNonEthPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorNonEthPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorNonEthPaused represents a Paused event raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) FilterPaused(opts *bind.FilterOpts) (*HanaConnectorNonEthPausedIterator, error) {

	logs, sub, err := _HanaConnectorNonEth.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorNonEthPausedIterator{contract: _HanaConnectorNonEth.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HanaConnectorNonEthPaused) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorNonEth.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorNonEthPaused)
				if err := _HanaConnectorNonEth.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) ParsePaused(log types.Log) (*HanaConnectorNonEthPaused, error) {
	event := new(HanaConnectorNonEthPaused)
	if err := _HanaConnectorNonEth.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorNonEthPauserAddressUpdatedIterator is returned from FilterPauserAddressUpdated and is used to iterate over the raw logs and unpacked data for PauserAddressUpdated events raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthPauserAddressUpdatedIterator struct {
	Event *HanaConnectorNonEthPauserAddressUpdated // Event containing the contract specifics and raw log

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
func (it *HanaConnectorNonEthPauserAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorNonEthPauserAddressUpdated)
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
		it.Event = new(HanaConnectorNonEthPauserAddressUpdated)
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
func (it *HanaConnectorNonEthPauserAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorNonEthPauserAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorNonEthPauserAddressUpdated represents a PauserAddressUpdated event raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthPauserAddressUpdated struct {
	UpdaterAddress common.Address
	NewTssAddress  common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPauserAddressUpdated is a free log retrieval operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address updaterAddress, address newTssAddress)
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) FilterPauserAddressUpdated(opts *bind.FilterOpts) (*HanaConnectorNonEthPauserAddressUpdatedIterator, error) {

	logs, sub, err := _HanaConnectorNonEth.contract.FilterLogs(opts, "PauserAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorNonEthPauserAddressUpdatedIterator{contract: _HanaConnectorNonEth.contract, event: "PauserAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchPauserAddressUpdated is a free log subscription operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address updaterAddress, address newTssAddress)
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) WatchPauserAddressUpdated(opts *bind.WatchOpts, sink chan<- *HanaConnectorNonEthPauserAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorNonEth.contract.WatchLogs(opts, "PauserAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorNonEthPauserAddressUpdated)
				if err := _HanaConnectorNonEth.contract.UnpackLog(event, "PauserAddressUpdated", log); err != nil {
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
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) ParsePauserAddressUpdated(log types.Log) (*HanaConnectorNonEthPauserAddressUpdated, error) {
	event := new(HanaConnectorNonEthPauserAddressUpdated)
	if err := _HanaConnectorNonEth.contract.UnpackLog(event, "PauserAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorNonEthTSSAddressUpdatedIterator is returned from FilterTSSAddressUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdated events raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthTSSAddressUpdatedIterator struct {
	Event *HanaConnectorNonEthTSSAddressUpdated // Event containing the contract specifics and raw log

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
func (it *HanaConnectorNonEthTSSAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorNonEthTSSAddressUpdated)
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
		it.Event = new(HanaConnectorNonEthTSSAddressUpdated)
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
func (it *HanaConnectorNonEthTSSAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorNonEthTSSAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorNonEthTSSAddressUpdated represents a TSSAddressUpdated event raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthTSSAddressUpdated struct {
	HanaTxSenderAddress common.Address
	NewTssAddress       common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdated is a free log retrieval operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address hanaTxSenderAddress, address newTssAddress)
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) FilterTSSAddressUpdated(opts *bind.FilterOpts) (*HanaConnectorNonEthTSSAddressUpdatedIterator, error) {

	logs, sub, err := _HanaConnectorNonEth.contract.FilterLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorNonEthTSSAddressUpdatedIterator{contract: _HanaConnectorNonEth.contract, event: "TSSAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdated is a free log subscription operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address hanaTxSenderAddress, address newTssAddress)
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) WatchTSSAddressUpdated(opts *bind.WatchOpts, sink chan<- *HanaConnectorNonEthTSSAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorNonEth.contract.WatchLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorNonEthTSSAddressUpdated)
				if err := _HanaConnectorNonEth.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
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
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) ParseTSSAddressUpdated(log types.Log) (*HanaConnectorNonEthTSSAddressUpdated, error) {
	event := new(HanaConnectorNonEthTSSAddressUpdated)
	if err := _HanaConnectorNonEth.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HanaConnectorNonEthUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthUnpausedIterator struct {
	Event *HanaConnectorNonEthUnpaused // Event containing the contract specifics and raw log

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
func (it *HanaConnectorNonEthUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HanaConnectorNonEthUnpaused)
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
		it.Event = new(HanaConnectorNonEthUnpaused)
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
func (it *HanaConnectorNonEthUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HanaConnectorNonEthUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HanaConnectorNonEthUnpaused represents a Unpaused event raised by the HanaConnectorNonEth contract.
type HanaConnectorNonEthUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) FilterUnpaused(opts *bind.FilterOpts) (*HanaConnectorNonEthUnpausedIterator, error) {

	logs, sub, err := _HanaConnectorNonEth.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &HanaConnectorNonEthUnpausedIterator{contract: _HanaConnectorNonEth.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HanaConnectorNonEthUnpaused) (event.Subscription, error) {

	logs, sub, err := _HanaConnectorNonEth.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HanaConnectorNonEthUnpaused)
				if err := _HanaConnectorNonEth.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_HanaConnectorNonEth *HanaConnectorNonEthFilterer) ParseUnpaused(log types.Log) (*HanaConnectorNonEthUnpaused, error) {
	event := new(HanaConnectorNonEthUnpaused)
	if err := _HanaConnectorNonEth.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
