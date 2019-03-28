package blockchain

import (
	"testing"
)

func TestBlockChain(t *testing.T) {
	c := &BlockChain{}
	b := c.CreateBlock(1, "0")

	if p := c.PreviousBlock(); p != b {
		t.Errorf("BlockChain.ProofOfWork() = %v, want %v", p.GetProof(), b.GetProof())
	}

	for index := 0; index < 10; index++ {
		b = c.CreateBlock(c.ProofOfWork(b.GetProof()), b.Hash())
		if !c.IsChainValid() {
			t.Errorf("BlockChain.CreateBlock() = %v, want %v", c.IsChainValid(), true)
		}
	}

	if got := c.IsChainValid(); got != true {
		t.Errorf("BlockChain.IsChainValid() = %v, want %v", got, true)
	}

	c.CreateBlock(10, b.Hash())
	if got := c.IsChainValid(); got != false {
		t.Errorf("BlockChain.IsChainValid() = %v, want %v", got, false)
	}

}
