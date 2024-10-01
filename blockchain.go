package assignment01bca //requirement of assignment

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv" //conversion between strings and integers
)

type Block struct {
	Transaction        string
	NonceX             int
	Previous_BlockHash string
	Hash               string
}

type Blockchain struct {
	Blocks []*Block
}

// global blockchain variable (Reason: multiple functions isey access kr sky)
var chain = Blockchain{}

// Calculation of Hash of Block
func CalculateHash(stringToHash string) string {
	B_hash := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(B_hash[:])
}

// Newblock creation and then isko hum blockchain mei add kr deingy
func NewBlock(transaction string, noncex int, previous_BlockHash string) *Block {
	newBlock := &Block{
		Transaction:        transaction,
		NonceX:             noncex,
		Previous_BlockHash: previous_BlockHash,
	}
	blockData := transaction + strconv.Itoa(noncex) + previous_BlockHash //new block=previous block ka hash+nonce+current block(all transations) hash
	newBlock.Hash = CalculateHash(blockData)                             //above line se sb add kr k ab uska hash calculate krne k liye calculate hash wale function mei sb jaega

	chain.Blocks = append(chain.Blocks, newBlock) //blockhain mei add krengy
	return newBlock
}

// blockchain k sarey block print hongy
func ListBlocks() {
	for i, block := range chain.Blocks {

		fmt.Printf("\033[31mBlock %d:\033[0m\n", i)

		fmt.Printf("\033[32m\tTransaction: %s\033[0m\n", block.Transaction)

		fmt.Printf("\033[33m\tNonce(X): %d\033[0m\n", block.NonceX)

		fmt.Printf("\033[34m\tPrevious Block Hash: %s\033[0m\n", block.Previous_BlockHash)

		fmt.Printf("\033[35m\tHash: %s\033[0m\n\n", block.Hash)
	}
}

// test krne k liye block ko change krengy
func ChangeBlock(BlockNo int, NewTransaction string) {
	if BlockNo < len(chain.Blocks) {
		//block ki transaction ko change krengy
		block := chain.Blocks[BlockNo]
		block.Transaction = NewTransaction

		//or transaction chnge hone k bad dobara se hash calculate krengy
		blockData := NewTransaction + strconv.Itoa(block.NonceX) + block.Previous_BlockHash
		block.Hash = CalculateHash(blockData)
	}
}

// blockhain ki verfication hogi
func VerifyChain() bool {
	for i := 1; i < len(chain.Blocks); i++ {

		previousBlock := chain.Blocks[i-1] //previous block hash
		currentBlock := chain.Blocks[i]    //cureent block hSH

		//ab dobara current block ka hash calculate krengy
		blockData := currentBlock.Transaction + strconv.Itoa(currentBlock.NonceX) + currentBlock.Previous_BlockHash
		calculatedHash := CalculateHash(blockData)

		// compare krengy current block k hash ko or previous block k hash ko
		if currentBlock.Hash != calculatedHash || currentBlock.Previous_BlockHash != previousBlock.Hash {
			fmt.Printf("\033[31mBlock %d has been tampered!\033[0m\n", i)
			return false
		}
	}

	// Good blockhain,Yahoooo!!!!!
	fmt.Println("\033[1;32m*****************\033[0m")
	fmt.Println("\033[32mBlockchain is valid!\033[0m")
	fmt.Println("\033[1;32m*****************\033[0m")

	return true
}
