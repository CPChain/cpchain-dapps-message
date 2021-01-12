// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package message

import (
	"math/big"
	"strings"

	cpchain "bitbucket.org/cpchain/chain"
	"bitbucket.org/cpchain/chain/accounts/abi"
	"bitbucket.org/cpchain/chain/accounts/abi/bind"
	"bitbucket.org/cpchain/chain/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// MessageABI is the input ABI used to generate the binding from.
const MessageABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enableContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"receivedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"disableContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"sentCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"message\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sentID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"recvID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"message\",\"type\":\"string\"}],\"name\":\"NewMessage\",\"type\":\"event\"}]"

// MessageBin is the compiled bytecode used for deploying new contracts.
const MessageBin = `0x60806040526000805460a060020a60ff0219167401000000000000000000000000000000000000000017815560015534801561003a57600080fd5b5060008054600160a060020a031916331790556103d78061005c6000396000f3006080604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306661abd8114610087578063238dafe0146100ae578063367edd32146100d757806360792829146100ee578063894ba8331461010f578063d6d8a23a14610124578063de6f24bb14610145575b600080fd5b34801561009357600080fd5b5061009c61019f565b60408051918252519081900360200190f35b3480156100ba57600080fd5b506100c36101a5565b604080519115158252519081900360200190f35b3480156100e357600080fd5b506100ec6101c6565b005b3480156100fa57600080fd5b5061009c600160a060020a0360043516610214565b34801561011b57600080fd5b506100ec61022f565b34801561013057600080fd5b5061009c600160a060020a0360043516610266565b60408051602060046024803582810135601f81018590048502860185019096528585526100ec958335600160a060020a03169536956044949193909101919081908401838280828437509497506102819650505050505050565b60015490565b60005474010000000000000000000000000000000000000000900460ff1681565b600054600160a060020a031633146101dd57600080fd5b6000805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600160a060020a031660009081526003602052604090205490565b600054600160a060020a0316331461024657600080fd5b6000805474ff000000000000000000000000000000000000000019169055565b600160a060020a031660009081526002602052604090205490565b60005474010000000000000000000000000000000000000000900460ff1615156102aa57600080fd5b6001805481018155336000818152600260209081526040808320805486018155600160a060020a0388168085526003845282852080549097019687905590548251868152808501929092529181018290526060810186905260a06080820181815288519183019190915287517fb11b470cd1e7e24bdb3a4b2a0e4b1e5de9221663e94920c122f585cb1e5f9027978a96949590948a94939260c08501928601918190849084905b83811015610369578181015183820152602001610351565b50505050905090810190601f1680156103965780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a150505600a165627a7a723058208c3cc1e4938a33403850bc02f4378de5d0279f5fb337a77905f9e6ce7984facb0029`

// DeployMessage deploys a new cpchain contract, binding an instance of Message to it.
func DeployMessage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Message, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MessageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Message{MessageCaller: MessageCaller{contract: contract}, MessageTransactor: MessageTransactor{contract: contract}, MessageFilterer: MessageFilterer{contract: contract}}, nil
}

// Message is an auto generated Go binding around an cpchain contract.
type Message struct {
	MessageCaller     // Read-only binding to the contract
	MessageTransactor // Write-only binding to the contract
	MessageFilterer   // Log filterer for contract events
}

// MessageCaller is an auto generated read-only Go binding around an cpchain contract.
type MessageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTransactor is an auto generated write-only Go binding around an cpchain contract.
type MessageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageFilterer is an auto generated log filtering Go binding around an cpchain contract events.
type MessageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageSession is an auto generated Go binding around an cpchain contract,
// with pre-set call and transact options.
type MessageSession struct {
	Contract     *Message          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageCallerSession is an auto generated read-only Go binding around an cpchain contract,
// with pre-set call options.
type MessageCallerSession struct {
	Contract *MessageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MessageTransactorSession is an auto generated write-only Go binding around an cpchain contract,
// with pre-set transact options.
type MessageTransactorSession struct {
	Contract     *MessageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MessageRaw is an auto generated low-level Go binding around an cpchain contract.
type MessageRaw struct {
	Contract *Message // Generic contract binding to access the raw methods on
}

// MessageCallerRaw is an auto generated low-level read-only Go binding around an cpchain contract.
type MessageCallerRaw struct {
	Contract *MessageCaller // Generic read-only contract binding to access the raw methods on
}

// MessageTransactorRaw is an auto generated low-level write-only Go binding around an cpchain contract.
type MessageTransactorRaw struct {
	Contract *MessageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessage creates a new instance of Message, bound to a specific deployed contract.
func NewMessage(address common.Address, backend bind.ContractBackend) (*Message, error) {
	contract, err := bindMessage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Message{MessageCaller: MessageCaller{contract: contract}, MessageTransactor: MessageTransactor{contract: contract}, MessageFilterer: MessageFilterer{contract: contract}}, nil
}

// NewMessageCaller creates a new read-only instance of Message, bound to a specific deployed contract.
func NewMessageCaller(address common.Address, caller bind.ContractCaller) (*MessageCaller, error) {
	contract, err := bindMessage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageCaller{contract: contract}, nil
}

// NewMessageTransactor creates a new write-only instance of Message, bound to a specific deployed contract.
func NewMessageTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageTransactor, error) {
	contract, err := bindMessage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTransactor{contract: contract}, nil
}

// NewMessageFilterer creates a new log filterer instance of Message, bound to a specific deployed contract.
func NewMessageFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageFilterer, error) {
	contract, err := bindMessage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageFilterer{contract: contract}, nil
}

// bindMessage binds a generic wrapper to an already deployed contract.
func bindMessage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Message *MessageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Message.Contract.MessageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Message *MessageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Message.Contract.MessageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Message *MessageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Message.Contract.MessageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Message *MessageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Message.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Message *MessageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Message.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Message *MessageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Message.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_Message *MessageCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Message.contract.Call(opts, out, "count")
	return *ret0, err
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_Message *MessageSession) Count() (*big.Int, error) {
	return _Message.Contract.Count(&_Message.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_Message *MessageCallerSession) Count() (*big.Int, error) {
	return _Message.Contract.Count(&_Message.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_Message *MessageCaller) Enabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Message.contract.Call(opts, out, "enabled")
	return *ret0, err
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_Message *MessageSession) Enabled() (bool, error) {
	return _Message.Contract.Enabled(&_Message.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_Message *MessageCallerSession) Enabled() (bool, error) {
	return _Message.Contract.Enabled(&_Message.CallOpts)
}

// ReceivedCount is a free data retrieval call binding the contract method 0x60792829.
//
// Solidity: function receivedCount(addr address) constant returns(uint256)
func (_Message *MessageCaller) ReceivedCount(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Message.contract.Call(opts, out, "receivedCount", addr)
	return *ret0, err
}

// ReceivedCount is a free data retrieval call binding the contract method 0x60792829.
//
// Solidity: function receivedCount(addr address) constant returns(uint256)
func (_Message *MessageSession) ReceivedCount(addr common.Address) (*big.Int, error) {
	return _Message.Contract.ReceivedCount(&_Message.CallOpts, addr)
}

// ReceivedCount is a free data retrieval call binding the contract method 0x60792829.
//
// Solidity: function receivedCount(addr address) constant returns(uint256)
func (_Message *MessageCallerSession) ReceivedCount(addr common.Address) (*big.Int, error) {
	return _Message.Contract.ReceivedCount(&_Message.CallOpts, addr)
}

// SentCount is a free data retrieval call binding the contract method 0xd6d8a23a.
//
// Solidity: function sentCount(addr address) constant returns(uint256)
func (_Message *MessageCaller) SentCount(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Message.contract.Call(opts, out, "sentCount", addr)
	return *ret0, err
}

// SentCount is a free data retrieval call binding the contract method 0xd6d8a23a.
//
// Solidity: function sentCount(addr address) constant returns(uint256)
func (_Message *MessageSession) SentCount(addr common.Address) (*big.Int, error) {
	return _Message.Contract.SentCount(&_Message.CallOpts, addr)
}

// SentCount is a free data retrieval call binding the contract method 0xd6d8a23a.
//
// Solidity: function sentCount(addr address) constant returns(uint256)
func (_Message *MessageCallerSession) SentCount(addr common.Address) (*big.Int, error) {
	return _Message.Contract.SentCount(&_Message.CallOpts, addr)
}

// DisableContract is a paid mutator transaction binding the contract method 0x894ba833.
//
// Solidity: function disableContract() returns()
func (_Message *MessageTransactor) DisableContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Message.contract.Transact(opts, "disableContract")
}

// DisableContract is a paid mutator transaction binding the contract method 0x894ba833.
//
// Solidity: function disableContract() returns()
func (_Message *MessageSession) DisableContract() (*types.Transaction, error) {
	return _Message.Contract.DisableContract(&_Message.TransactOpts)
}

// DisableContract is a paid mutator transaction binding the contract method 0x894ba833.
//
// Solidity: function disableContract() returns()
func (_Message *MessageTransactorSession) DisableContract() (*types.Transaction, error) {
	return _Message.Contract.DisableContract(&_Message.TransactOpts)
}

// EnableContract is a paid mutator transaction binding the contract method 0x367edd32.
//
// Solidity: function enableContract() returns()
func (_Message *MessageTransactor) EnableContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Message.contract.Transact(opts, "enableContract")
}

// EnableContract is a paid mutator transaction binding the contract method 0x367edd32.
//
// Solidity: function enableContract() returns()
func (_Message *MessageSession) EnableContract() (*types.Transaction, error) {
	return _Message.Contract.EnableContract(&_Message.TransactOpts)
}

// EnableContract is a paid mutator transaction binding the contract method 0x367edd32.
//
// Solidity: function enableContract() returns()
func (_Message *MessageTransactorSession) EnableContract() (*types.Transaction, error) {
	return _Message.Contract.EnableContract(&_Message.TransactOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0xde6f24bb.
//
// Solidity: function sendMessage(to address, message string) returns()
func (_Message *MessageTransactor) SendMessage(opts *bind.TransactOpts, to common.Address, message string) (*types.Transaction, error) {
	return _Message.contract.Transact(opts, "sendMessage", to, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xde6f24bb.
//
// Solidity: function sendMessage(to address, message string) returns()
func (_Message *MessageSession) SendMessage(to common.Address, message string) (*types.Transaction, error) {
	return _Message.Contract.SendMessage(&_Message.TransactOpts, to, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xde6f24bb.
//
// Solidity: function sendMessage(to address, message string) returns()
func (_Message *MessageTransactorSession) SendMessage(to common.Address, message string) (*types.Transaction, error) {
	return _Message.Contract.SendMessage(&_Message.TransactOpts, to, message)
}

// MessageNewMessageIterator is returned from FilterNewMessage and is used to iterate over the raw logs and unpacked data for NewMessage events raised by the Message contract.
type MessageNewMessageIterator struct {
	Event *MessageNewMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MessageNewMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageNewMessage)
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
		it.Event = new(MessageNewMessage)
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
func (it *MessageNewMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageNewMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageNewMessage represents a NewMessage event raised by the Message contract.
type MessageNewMessage struct {
	From    common.Address
	To      common.Address
	SentID  *big.Int
	RecvID  *big.Int
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewMessage is a free log retrieval operation binding the contract event 0xb11b470cd1e7e24bdb3a4b2a0e4b1e5de9221663e94920c122f585cb1e5f9027.
//
// Solidity: e NewMessage(from address, to address, sentID uint256, recvID uint256, message string)
func (_Message *MessageFilterer) FilterNewMessage(opts *bind.FilterOpts) (*MessageNewMessageIterator, error) {

	logs, sub, err := _Message.contract.FilterLogs(opts, "NewMessage")
	if err != nil {
		return nil, err
	}
	return &MessageNewMessageIterator{contract: _Message.contract, event: "NewMessage", logs: logs, sub: sub}, nil
}

// WatchNewMessage is a free log subscription operation binding the contract event 0xb11b470cd1e7e24bdb3a4b2a0e4b1e5de9221663e94920c122f585cb1e5f9027.
//
// Solidity: e NewMessage(from address, to address, sentID uint256, recvID uint256, message string)
func (_Message *MessageFilterer) WatchNewMessage(opts *bind.WatchOpts, sink chan<- *MessageNewMessage) (event.Subscription, error) {

	logs, sub, err := _Message.contract.WatchLogs(opts, "NewMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageNewMessage)
				if err := _Message.contract.UnpackLog(event, "NewMessage", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058200383508c084362424c2fd8727ed7f805f84d0c8e94faef3b0b450896b72791ec0029`

// DeploySafeMath deploys a new cpchain contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an cpchain contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an cpchain contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an cpchain contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an cpchain contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an cpchain contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an cpchain contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an cpchain contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an cpchain contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an cpchain contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an cpchain contract.
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
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// SetABI is the input ABI used to generate the binding from.
const SetABI = "[]"

// SetBin is the compiled bytecode used for deploying new contracts.
const SetBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582017662ff907862766b9f9896ce4bd8ca53e202f25f23c0c50706f9e8d86e96b900029`

// DeploySet deploys a new cpchain contract, binding an instance of Set to it.
func DeploySet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Set, error) {
	parsed, err := abi.JSON(strings.NewReader(SetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Set{SetCaller: SetCaller{contract: contract}, SetTransactor: SetTransactor{contract: contract}, SetFilterer: SetFilterer{contract: contract}}, nil
}

// Set is an auto generated Go binding around an cpchain contract.
type Set struct {
	SetCaller     // Read-only binding to the contract
	SetTransactor // Write-only binding to the contract
	SetFilterer   // Log filterer for contract events
}

// SetCaller is an auto generated read-only Go binding around an cpchain contract.
type SetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SetTransactor is an auto generated write-only Go binding around an cpchain contract.
type SetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SetFilterer is an auto generated log filtering Go binding around an cpchain contract events.
type SetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SetSession is an auto generated Go binding around an cpchain contract,
// with pre-set call and transact options.
type SetSession struct {
	Contract     *Set              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SetCallerSession is an auto generated read-only Go binding around an cpchain contract,
// with pre-set call options.
type SetCallerSession struct {
	Contract *SetCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SetTransactorSession is an auto generated write-only Go binding around an cpchain contract,
// with pre-set transact options.
type SetTransactorSession struct {
	Contract     *SetTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SetRaw is an auto generated low-level Go binding around an cpchain contract.
type SetRaw struct {
	Contract *Set // Generic contract binding to access the raw methods on
}

// SetCallerRaw is an auto generated low-level read-only Go binding around an cpchain contract.
type SetCallerRaw struct {
	Contract *SetCaller // Generic read-only contract binding to access the raw methods on
}

// SetTransactorRaw is an auto generated low-level write-only Go binding around an cpchain contract.
type SetTransactorRaw struct {
	Contract *SetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSet creates a new instance of Set, bound to a specific deployed contract.
func NewSet(address common.Address, backend bind.ContractBackend) (*Set, error) {
	contract, err := bindSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Set{SetCaller: SetCaller{contract: contract}, SetTransactor: SetTransactor{contract: contract}, SetFilterer: SetFilterer{contract: contract}}, nil
}

// NewSetCaller creates a new read-only instance of Set, bound to a specific deployed contract.
func NewSetCaller(address common.Address, caller bind.ContractCaller) (*SetCaller, error) {
	contract, err := bindSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SetCaller{contract: contract}, nil
}

// NewSetTransactor creates a new write-only instance of Set, bound to a specific deployed contract.
func NewSetTransactor(address common.Address, transactor bind.ContractTransactor) (*SetTransactor, error) {
	contract, err := bindSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SetTransactor{contract: contract}, nil
}

// NewSetFilterer creates a new log filterer instance of Set, bound to a specific deployed contract.
func NewSetFilterer(address common.Address, filterer bind.ContractFilterer) (*SetFilterer, error) {
	contract, err := bindSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SetFilterer{contract: contract}, nil
}

// bindSet binds a generic wrapper to an already deployed contract.
func bindSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Set *SetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Set.Contract.SetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Set *SetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Set.Contract.SetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Set *SetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Set.Contract.SetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Set *SetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Set.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Set *SetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Set.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Set *SetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Set.Contract.contract.Transact(opts, method, params...)
}
