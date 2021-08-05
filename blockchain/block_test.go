package blockchain

import (
	"reflect"
	"testing"

	"github.com/derekahndev/sadfrogcoin/utils"
)

func TestCreateBlock(t *testing.T) {
	dbStorage = fakeDB{}
	Mempool().Txs["test"] = &Tx{}
	b := createBlock("x", 1, 1)
	if reflect.TypeOf(b) != reflect.TypeOf(&Block{}) {
		t.Error("createBlock should return an instance of a block")
	}
}

func TestFindBlock(t *testing.T) {
	t.Run("Block is found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				b := &Block{}
				return utils.ToBytes(b)
			},
		}
		block, _ := FindBlock("xx")
		if block.Height != 1 {
			t.Error("Block should be found.")
		}
	})
}
