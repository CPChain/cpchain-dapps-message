package message_test

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"testing"
	"time"

	"message"

	"bitbucket.org/cpchain/chain/accounts/abi/bind"
	"bitbucket.org/cpchain/chain/accounts/abi/bind/backends"
	"bitbucket.org/cpchain/chain/configs"
	"bitbucket.org/cpchain/chain/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ownerKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	ownerAddr   = crypto.PubkeyToAddress(ownerKey.PublicKey)

	candidateKey, _ = crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	candidateAddr   = crypto.PubkeyToAddress(candidateKey.PublicKey)

	candidate2Key, _ = crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	candidate2Addr   = crypto.PubkeyToAddress(candidate2Key.PublicKey)
)

func deploy(prvKey *ecdsa.PrivateKey, backend *backends.SimulatedBackend) (common.Address, *message.Message) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	add, _, instance, err := message.DeployMessage(deployTransactor, backend)
	if err != nil {
		log.Fatal("deploy message failed:", "error", err)
	}
	return add, instance
}

func TestDeployAndRegister(t *testing.T) {
	// deploy contract
	contractBackend := backends.NewDporSimulatedBackend(core.GenesisAlloc{
		ownerAddr:      {Balance: big.NewInt(1000000000000)},
		candidateAddr:  {Balance: new(big.Int).Mul(big.NewInt(1000000), big.NewInt(configs.Cpc))},
		candidate2Addr: {Balance: new(big.Int).Mul(big.NewInt(1000000), big.NewInt(configs.Cpc))}})
	_, instance := deploy(ownerKey, contractBackend)

	ch := make(chan *message.MessageNewMessage)
	done := make(chan struct{})
	if _, err := instance.WatchNewMessage(nil, ch); err != nil {
		t.Fatal(err)
	}

	watched := false

	go func() {
		for {
			select {
			case e := <-ch:
				watched = true
				if e.From.Hex() != candidateAddr.Hex() {
					t.Error("FROM error")
				}
				if e.RecvID.Int64() != 1 {
					t.Error("the first recvID should be 1")
				}
				if e.SentID.Int64() != 1 {
					t.Error("the first sentID should be 1")
				}
			case <-done:
				return
			}
		}
	}()

	// send message
	sendMsg(t, instance, candidateKey, candidate2Addr, "Hello world")
	contractBackend.Commit()

	// check count
	checkNum(t, instance, 1)

	sentCnt, _ := instance.SentCount(nil, candidateAddr)
	if sentCnt.Int64() != 1 {
		t.Error("sentCnt should be 1")
	}

	recvCnt, _ := instance.ReceivedCount(nil, candidate2Addr)
	if recvCnt.Int64() != 1 {
		t.Error("receivedCnt should be 1")
	}

	// wait
	<-time.After(1 * time.Second)
	done <- struct{}{}

	if !watched {
		t.Error("Not watched the event")
	}
}

func sendMsg(t *testing.T, instance *message.Message, key *ecdsa.PrivateKey, to common.Address, message string) {
	txOpts := bind.NewKeyedTransactor(key)
	txOpts.GasLimit = uint64(50000000)
	txOpts.Value = big.NewInt(0)
	_, err := instance.SendMessage(txOpts, to, message)
	if err != nil {
		t.Fatal("register failed:", err)
	}
}

func checkError(t *testing.T, title string, err error) {
	if err != nil {
		t.Fatal(title, ":", err)
	}
}

func checkNum(t *testing.T, instance *message.Message, amount int) {
	num, err := instance.Count(nil)
	checkError(t, "get num", err)

	if num.Cmp(new(big.Int).SetInt64(int64(amount))) != 0 {
		t.Errorf("message's num %d != %d", num, amount)
	}
}
