package permissions

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = `{"address":"ed9d02e382b34818e88b88a309c7fe71e65f419d","crypto":{"cipher":"aes-128-ctr","ciphertext":"4e77046ba3f699e744acb4a89c36a3ea1158a1bd90a076d36675f4c883864377","cipherparams":{"iv":"a8932af2a3c0225ee8e872bc0e462c11"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8ca49552b3e92f79c51f2cd3d38dfc723412c212e702bd337a3724e8937aff0f"},"mac":"6d1354fef5aa0418389b1a5d1f5ee0050d7273292a1171c51fd02f9ecff55264"},"id":"a65d1ac3-db7e-445d-a1cc-b6c5eeaa05e0","version":3}`

const enode1 = "ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef"
const enode2 = "0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416"
const addr = "0xd9d64b7dc034fafdba5dc2902875a67b5d586420"

func TestContract_ProposeNode(t *testing.T) {
	conn, err := ethclient.Dial("/home/vagrant/quorum-examples/examples/7nodes/qdata/dd1/geth.ipc")
	if err != nil {
		t.Errorf("Failed to connect to the Ethereum client: %v", err)
	}

	var contractAddr = common.HexToAddress(addr)

	permissions, err := NewPermissions(contractAddr, conn)
	if err != nil {
		t.Errorf("Failed to instantiate a Permissions contract: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "")
	if err != nil {
		t.Errorf("Failed to create authorized transactor: %v", err)
	}

	tx, err := permissions.ProposeNode(auth, enode2, true, true)
	if err != nil {
		t.Errorf("Failed to propose node: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
}

func TestContract_GetNodeIndexForNode(t *testing.T) {
	conn, err := ethclient.Dial("/home/vagrant/quorum-examples/examples/7nodes/qdata/dd1/geth.ipc")
	if err != nil {
		t.Errorf("Failed to connect to the Ethereum client: %v", err)
	}

	var contractAddr = common.HexToAddress(addr)

	permissions, err := NewPermissions(contractAddr, conn)
	if err != nil {
		t.Errorf("Failed to instantiate a Permissions contract: %v", err)
	}

	nodeIndex, err := permissions.GetNodeIndexForNode(nil, enode2)

	if err != nil {
		t.Errorf("Failed to create authorized transactor: %v", err)
	} else {
		fmt.Printf("nodeIndex: %v\n", nodeIndex)
	}

}
