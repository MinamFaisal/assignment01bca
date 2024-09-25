package assignment01bca
package main



import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"assignment01bca"
)


type block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

type blockchain struct {
	Blocks []*block
}

// Declare a global blockchain variable
var chain = blockchain{}

// CreateHash is a helper function to calculate the hash of a block.
func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hash[:])
}

// NewBlock function creates a new block and adds it to the blockchain
func NewBlock(transaction string, nonce int, previousHash string) *block {
	// Create the block
	newBlock := &block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	// Calculate the hash for the new block
	blockData := transaction + strconv.Itoa(nonce) + previousHash
	newBlock.Hash = CalculateHash(blockData)

	// Append block to the blockchain
	chain.Blocks = append(chain.Blocks, newBlock)

	return newBlock
}

// ListBlocks prints out all blocks in the blockchain
func ListBlocks() {
	for i, blk := range chain.Blocks {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("\tTransaction: %s\n", blk.Transaction)
		fmt.Printf("\tNonce: %d\n", blk.Nonce)
		fmt.Printf("\tPrevious Hash: %s\n", blk.PreviousHash)
		fmt.Printf("\tHash: %s\n\n", blk.Hash)
	}
}


// ChangeBlock allows changing the transaction of a block at a specific index
func ChangeBlock(index int, newTransaction string) {
	if index < len(chain.Blocks) {
		chain.Blocks[index].Transaction = newTransaction
		// Recalculate the hash since the transaction has changed
		blockData := newTransaction + strconv.Itoa(chain.Blocks[index].Nonce) + chain.Blocks[index].PreviousHash
		chain.Blocks[index].Hash = CalculateHash(blockData)
	}
}


// VerifyChain checks the validity of the blockchain
func VerifyChain() bool {
	for i := 1; i < len(chain.Blocks); i++ {
		previousBlock := chain.Blocks[i-1]
		currentBlock := chain.Blocks[i]

		// Recalculate the hash to verify it matches
		blockData := currentBlock.Transaction + strconv.Itoa(currentBlock.Nonce) + currentBlock.PreviousHash
		calculatedHash := CalculateHash(blockData)

		// If hash doesn't match, blockchain is invalid
		if currentBlock.Hash != calculatedHash || currentBlock.PreviousHash != previousBlock.Hash {
			fmt.Printf("Block %d has been tampered with!\n", i)
			return false
		}
	}

	fmt.Println("Blockchain is valid!")
	return true
}




func main() {
	// Create some blocks
	genesisBlock := assignment01bca.NewBlock("genesis to bob", 1, "")
	fmt.Println("Genesis Block created!")
	assignment01bca.NewBlock("bob to alice", 2, genesisBlock.Hash)
	assignment01bca.NewBlock("alice to eve", 3, genesisBlock.Hash)

	// List the blocks
	fmt.Println("Listing all blocks in the blockchain:")
	assignment01bca.ListBlocks()

	// Change a block and list blocks again
	fmt.Println("Changing the transaction of Block 1...")
	assignment01bca.ChangeBlock(1, "bob to charlie")
	assignment01bca.ListBlocks()

	// Verify the blockchain
	fmt.Println("Verifying the blockchain after modification:")
	assignment01bca.VerifyChain()
}


