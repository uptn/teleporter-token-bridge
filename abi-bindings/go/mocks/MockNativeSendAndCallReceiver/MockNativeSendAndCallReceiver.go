// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocknativesendandcallreceiver

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MockNativeSendAndCallReceiverMetaData contains all meta data concerning the MockNativeSendAndCallReceiver contract.
var MockNativeSendAndCallReceiverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"TokensReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"blockSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"blockedSenders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"blocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"receiveTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061035c806100206000396000f3fe6080604052600436106100345760003560e01c8063355fb77c14610039578063487bd69e1461004e5780637f450d8d1461009a575b600080fd5b61004c61004736600461023d565b6100e3565b005b34801561005a57600080fd5b506100866100693660046102c4565b600060208181529281526040808220909352908152205460ff1681565b604051901515815260200160405180910390f35b3480156100a657600080fd5b5061004c6100b53660046102c4565b6000918252602082815260408084206001600160a01b0390931684529190529020805460ff19166001179055565b6000848152602081815260408083206001600160a01b038716845290915290205460ff161561016f5760405162461bcd60e51b815260206004820152602d60248201527f4d6f636b4e617469766553656e64416e6443616c6c52656365697665723a207360448201526c195b99195c88189b1bd8dad959609a1b60648201526084015b60405180910390fd5b826001600160a01b0316847f640996264f389b4827dd450bf3c45efe9db2f5ef790ec42abdc8ab81461a7fe03485856040516101ad939291906102f0565b60405180910390a3600081900361021b5760405162461bcd60e51b815260206004820152602c60248201527f4d6f636b4e617469766553656e64416e6443616c6c52656365697665723a206560448201526b1b5c1d1e481c185e5b1bd85960a21b6064820152608401610166565b50505050565b80356001600160a01b038116811461023857600080fd5b919050565b6000806000806060858703121561025357600080fd5b8435935061026360208601610221565b9250604085013567ffffffffffffffff8082111561028057600080fd5b818701915087601f83011261029457600080fd5b8135818111156102a357600080fd5b8860208285010111156102b557600080fd5b95989497505060200194505050565b600080604083850312156102d757600080fd5b823591506102e760208401610221565b90509250929050565b83815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f191601019291505056fea26469706673582212203ceb6324bd75e00fe7fca4017a2fa346467e805eef5bdae3ad4b89cb505d1c6264736f6c63430008120033",
}

// MockNativeSendAndCallReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MockNativeSendAndCallReceiverMetaData.ABI instead.
var MockNativeSendAndCallReceiverABI = MockNativeSendAndCallReceiverMetaData.ABI

// MockNativeSendAndCallReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockNativeSendAndCallReceiverMetaData.Bin instead.
var MockNativeSendAndCallReceiverBin = MockNativeSendAndCallReceiverMetaData.Bin

// DeployMockNativeSendAndCallReceiver deploys a new Ethereum contract, binding an instance of MockNativeSendAndCallReceiver to it.
func DeployMockNativeSendAndCallReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockNativeSendAndCallReceiver, error) {
	parsed, err := MockNativeSendAndCallReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockNativeSendAndCallReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockNativeSendAndCallReceiver{MockNativeSendAndCallReceiverCaller: MockNativeSendAndCallReceiverCaller{contract: contract}, MockNativeSendAndCallReceiverTransactor: MockNativeSendAndCallReceiverTransactor{contract: contract}, MockNativeSendAndCallReceiverFilterer: MockNativeSendAndCallReceiverFilterer{contract: contract}}, nil
}

// MockNativeSendAndCallReceiver is an auto generated Go binding around an Ethereum contract.
type MockNativeSendAndCallReceiver struct {
	MockNativeSendAndCallReceiverCaller     // Read-only binding to the contract
	MockNativeSendAndCallReceiverTransactor // Write-only binding to the contract
	MockNativeSendAndCallReceiverFilterer   // Log filterer for contract events
}

// MockNativeSendAndCallReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockNativeSendAndCallReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockNativeSendAndCallReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockNativeSendAndCallReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockNativeSendAndCallReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockNativeSendAndCallReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockNativeSendAndCallReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockNativeSendAndCallReceiverSession struct {
	Contract     *MockNativeSendAndCallReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MockNativeSendAndCallReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockNativeSendAndCallReceiverCallerSession struct {
	Contract *MockNativeSendAndCallReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// MockNativeSendAndCallReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockNativeSendAndCallReceiverTransactorSession struct {
	Contract     *MockNativeSendAndCallReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// MockNativeSendAndCallReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockNativeSendAndCallReceiverRaw struct {
	Contract *MockNativeSendAndCallReceiver // Generic contract binding to access the raw methods on
}

// MockNativeSendAndCallReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockNativeSendAndCallReceiverCallerRaw struct {
	Contract *MockNativeSendAndCallReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// MockNativeSendAndCallReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockNativeSendAndCallReceiverTransactorRaw struct {
	Contract *MockNativeSendAndCallReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockNativeSendAndCallReceiver creates a new instance of MockNativeSendAndCallReceiver, bound to a specific deployed contract.
func NewMockNativeSendAndCallReceiver(address common.Address, backend bind.ContractBackend) (*MockNativeSendAndCallReceiver, error) {
	contract, err := bindMockNativeSendAndCallReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockNativeSendAndCallReceiver{MockNativeSendAndCallReceiverCaller: MockNativeSendAndCallReceiverCaller{contract: contract}, MockNativeSendAndCallReceiverTransactor: MockNativeSendAndCallReceiverTransactor{contract: contract}, MockNativeSendAndCallReceiverFilterer: MockNativeSendAndCallReceiverFilterer{contract: contract}}, nil
}

// NewMockNativeSendAndCallReceiverCaller creates a new read-only instance of MockNativeSendAndCallReceiver, bound to a specific deployed contract.
func NewMockNativeSendAndCallReceiverCaller(address common.Address, caller bind.ContractCaller) (*MockNativeSendAndCallReceiverCaller, error) {
	contract, err := bindMockNativeSendAndCallReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockNativeSendAndCallReceiverCaller{contract: contract}, nil
}

// NewMockNativeSendAndCallReceiverTransactor creates a new write-only instance of MockNativeSendAndCallReceiver, bound to a specific deployed contract.
func NewMockNativeSendAndCallReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MockNativeSendAndCallReceiverTransactor, error) {
	contract, err := bindMockNativeSendAndCallReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockNativeSendAndCallReceiverTransactor{contract: contract}, nil
}

// NewMockNativeSendAndCallReceiverFilterer creates a new log filterer instance of MockNativeSendAndCallReceiver, bound to a specific deployed contract.
func NewMockNativeSendAndCallReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MockNativeSendAndCallReceiverFilterer, error) {
	contract, err := bindMockNativeSendAndCallReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockNativeSendAndCallReceiverFilterer{contract: contract}, nil
}

// bindMockNativeSendAndCallReceiver binds a generic wrapper to an already deployed contract.
func bindMockNativeSendAndCallReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockNativeSendAndCallReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockNativeSendAndCallReceiver.Contract.MockNativeSendAndCallReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.MockNativeSendAndCallReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.MockNativeSendAndCallReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockNativeSendAndCallReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.contract.Transact(opts, method, params...)
}

// BlockedSenders is a free data retrieval call binding the contract method 0x487bd69e.
//
// Solidity: function blockedSenders(bytes32 blockchainID, address senderAddress) view returns(bool blocked)
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverCaller) BlockedSenders(opts *bind.CallOpts, blockchainID [32]byte, senderAddress common.Address) (bool, error) {
	var out []interface{}
	err := _MockNativeSendAndCallReceiver.contract.Call(opts, &out, "blockedSenders", blockchainID, senderAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlockedSenders is a free data retrieval call binding the contract method 0x487bd69e.
//
// Solidity: function blockedSenders(bytes32 blockchainID, address senderAddress) view returns(bool blocked)
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverSession) BlockedSenders(blockchainID [32]byte, senderAddress common.Address) (bool, error) {
	return _MockNativeSendAndCallReceiver.Contract.BlockedSenders(&_MockNativeSendAndCallReceiver.CallOpts, blockchainID, senderAddress)
}

// BlockedSenders is a free data retrieval call binding the contract method 0x487bd69e.
//
// Solidity: function blockedSenders(bytes32 blockchainID, address senderAddress) view returns(bool blocked)
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverCallerSession) BlockedSenders(blockchainID [32]byte, senderAddress common.Address) (bool, error) {
	return _MockNativeSendAndCallReceiver.Contract.BlockedSenders(&_MockNativeSendAndCallReceiver.CallOpts, blockchainID, senderAddress)
}

// BlockSender is a paid mutator transaction binding the contract method 0x7f450d8d.
//
// Solidity: function blockSender(bytes32 blockchainID, address senderAddress) returns()
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverTransactor) BlockSender(opts *bind.TransactOpts, blockchainID [32]byte, senderAddress common.Address) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.contract.Transact(opts, "blockSender", blockchainID, senderAddress)
}

// BlockSender is a paid mutator transaction binding the contract method 0x7f450d8d.
//
// Solidity: function blockSender(bytes32 blockchainID, address senderAddress) returns()
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverSession) BlockSender(blockchainID [32]byte, senderAddress common.Address) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.BlockSender(&_MockNativeSendAndCallReceiver.TransactOpts, blockchainID, senderAddress)
}

// BlockSender is a paid mutator transaction binding the contract method 0x7f450d8d.
//
// Solidity: function blockSender(bytes32 blockchainID, address senderAddress) returns()
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverTransactorSession) BlockSender(blockchainID [32]byte, senderAddress common.Address) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.BlockSender(&_MockNativeSendAndCallReceiver.TransactOpts, blockchainID, senderAddress)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x355fb77c.
//
// Solidity: function receiveTokens(bytes32 sourceBlockchainID, address originSenderAddress, bytes payload) payable returns()
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverTransactor) ReceiveTokens(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, payload []byte) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.contract.Transact(opts, "receiveTokens", sourceBlockchainID, originSenderAddress, payload)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x355fb77c.
//
// Solidity: function receiveTokens(bytes32 sourceBlockchainID, address originSenderAddress, bytes payload) payable returns()
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverSession) ReceiveTokens(sourceBlockchainID [32]byte, originSenderAddress common.Address, payload []byte) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.ReceiveTokens(&_MockNativeSendAndCallReceiver.TransactOpts, sourceBlockchainID, originSenderAddress, payload)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x355fb77c.
//
// Solidity: function receiveTokens(bytes32 sourceBlockchainID, address originSenderAddress, bytes payload) payable returns()
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverTransactorSession) ReceiveTokens(sourceBlockchainID [32]byte, originSenderAddress common.Address, payload []byte) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.ReceiveTokens(&_MockNativeSendAndCallReceiver.TransactOpts, sourceBlockchainID, originSenderAddress, payload)
}

// MockNativeSendAndCallReceiverTokensReceivedIterator is returned from FilterTokensReceived and is used to iterate over the raw logs and unpacked data for TokensReceived events raised by the MockNativeSendAndCallReceiver contract.
type MockNativeSendAndCallReceiverTokensReceivedIterator struct {
	Event *MockNativeSendAndCallReceiverTokensReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockNativeSendAndCallReceiverTokensReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockNativeSendAndCallReceiverTokensReceived)
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
		it.Event = new(MockNativeSendAndCallReceiverTokensReceived)
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
func (it *MockNativeSendAndCallReceiverTokensReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockNativeSendAndCallReceiverTokensReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockNativeSendAndCallReceiverTokensReceived represents a TokensReceived event raised by the MockNativeSendAndCallReceiver contract.
type MockNativeSendAndCallReceiverTokensReceived struct {
	SourceBlockchainID  [32]byte
	OriginSenderAddress common.Address
	Amount              *big.Int
	Payload             []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensReceived is a free log retrieval operation binding the contract event 0x640996264f389b4827dd450bf3c45efe9db2f5ef790ec42abdc8ab81461a7fe0.
//
// Solidity: event TokensReceived(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, uint256 amount, bytes payload)
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverFilterer) FilterTokensReceived(opts *bind.FilterOpts, sourceBlockchainID [][32]byte, originSenderAddress []common.Address) (*MockNativeSendAndCallReceiverTokensReceivedIterator, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _MockNativeSendAndCallReceiver.contract.FilterLogs(opts, "TokensReceived", sourceBlockchainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return &MockNativeSendAndCallReceiverTokensReceivedIterator{contract: _MockNativeSendAndCallReceiver.contract, event: "TokensReceived", logs: logs, sub: sub}, nil
}

// WatchTokensReceived is a free log subscription operation binding the contract event 0x640996264f389b4827dd450bf3c45efe9db2f5ef790ec42abdc8ab81461a7fe0.
//
// Solidity: event TokensReceived(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, uint256 amount, bytes payload)
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverFilterer) WatchTokensReceived(opts *bind.WatchOpts, sink chan<- *MockNativeSendAndCallReceiverTokensReceived, sourceBlockchainID [][32]byte, originSenderAddress []common.Address) (event.Subscription, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _MockNativeSendAndCallReceiver.contract.WatchLogs(opts, "TokensReceived", sourceBlockchainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockNativeSendAndCallReceiverTokensReceived)
				if err := _MockNativeSendAndCallReceiver.contract.UnpackLog(event, "TokensReceived", log); err != nil {
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

// ParseTokensReceived is a log parse operation binding the contract event 0x640996264f389b4827dd450bf3c45efe9db2f5ef790ec42abdc8ab81461a7fe0.
//
// Solidity: event TokensReceived(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, uint256 amount, bytes payload)
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverFilterer) ParseTokensReceived(log types.Log) (*MockNativeSendAndCallReceiverTokensReceived, error) {
	event := new(MockNativeSendAndCallReceiverTokensReceived)
	if err := _MockNativeSendAndCallReceiver.contract.UnpackLog(event, "TokensReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
