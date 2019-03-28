package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

// Block : one block in a blockchain
type Block struct {
	Index     int       `json:"index"`
	TimeStamp time.Time `json:"timeStamp"`
	PHash     string    `json:"pHash"`
	Proof     int       `json:"proof"`
}

// BlockChain : chain of blocks
type BlockChain struct {
	Chain []Block `json:"chain"`
}

// GetProof : get Proof
func (b *Block) GetProof() int {
	return b.Proof
}

// Hash : get hash of a block
func (b *Block) Hash() string {
	d, err := json.Marshal(b)
	if err != nil {
		log.Println("error:", err)
	}
	h := sha256.Sum256([]byte(d))
	return hex.EncodeToString(h[:])
}

// CreateBlock : Create block
func (c *BlockChain) CreateBlock(proof int, pHash string) Block {
	block := Block{
		Index:     len(c.Chain) + 1,
		PHash:     pHash,
		Proof:     proof,
		TimeStamp: time.Now(),
	}
	c.Chain = append(c.Chain, block)
	return block
}

// GetChain : Get Chain
func (c *BlockChain) GetChain() []Block {
	return c.Chain
}

// PreviousBlock : get prevoid block
func (c *BlockChain) PreviousBlock() Block {
	return c.Chain[len(c.Chain)-1]
}

// ProofOfWork : check proof of work
func (c *BlockChain) ProofOfWork(pProof int) int {
	nProof := int(1)
	proved := false
	for proved == false {
		d := strconv.Itoa(int(math.Pow(float64(pProof), 2) - math.Pow(float64(nProof), 2)))
		h := sha256.Sum256([]byte(d))
		if strings.HasPrefix(hex.EncodeToString(h[:]), "0000") {
			proved = true
		} else {
			nProof++
		}
	}
	return nProof
}

// IsChainValid : checks if chain is valid
func (c *BlockChain) IsChainValid() bool {
	l := len(c.Chain)
	pBlock := c.Chain[0]
	for i := 1; i < l; i++ {
		block := c.Chain[i]
		if block.PHash != pBlock.Hash() {
			return false
		}
		pProof := pBlock.Proof
		proof := block.Proof
		d := strconv.Itoa(int(math.Pow(float64(pProof), 2) - math.Pow(float64(proof), 2)))
		h := sha256.Sum256([]byte(d))
		if !strings.HasPrefix(hex.EncodeToString(h[:]), "0000") {
			return false
		}
		pBlock = block
	}
	return true
}
