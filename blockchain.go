/*package assignment01bca //requirement of assignment

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv" //conversion between strings and integers
	"strings"
	"time"
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

*/

//assignment_02

package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Transaction struct to store transaction details
type Transaction struct {
	TransactionID              string
	SenderBlockchainAddress    string
	RecipientBlockchainAddress string
	Value                      float32
}

// Block struct is modified to include a list of transactions
type Block struct {
	TransactionPool    []*Transaction
	NonceX             int
	Previous_BlockHash string
	Hash               string
	Timestamp          time.Time
}

// Blockchain struct is modified to include a transaction pool
type Blockchain struct {
	Blocks          []*Block
	TransactionPool []*Transaction
}

// Global blockchain instance
var chain = Blockchain{}

// NewTransaction method creates a new transaction
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	transaction := &Transaction{
		SenderBlockchainAddress:    sender,
		RecipientBlockchainAddress: recipient,
		Value:                      value,
	}
	transactionData := sender + recipient + fmt.Sprintf("%f", value)
	transaction.TransactionID = CalculateHash(transactionData)
	return transaction
}

// AddTransaction adds a new transaction to the transaction pool
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.TransactionPool = append(bc.TransactionPool, t)
}

// CalculateHash method calculates the hash of a given string (unchanged)
func CalculateHash(stringToHash string) string {
	B_hash := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(B_hash[:])
}

// ProofOfWork method generates a nonce based on difficulty level
func ProofOfWork(previousHash string, difficulty int) int {
	nonce := 0
	prefix := strings.Repeat("0", difficulty)

	for {
		data := previousHash + strconv.Itoa(nonce)
		hash := CalculateHash(data)
		if strings.HasPrefix(hash, prefix) {
			break
		}
		nonce++
	}
	return nonce
}

// NewBlock method creates a new block with a transaction pool
func NewBlock(previous_BlockHash string, nonce int) *Block {
	newBlock := &Block{
		TransactionPool:    chain.TransactionPool,
		NonceX:             nonce,
		Previous_BlockHash: previous_BlockHash,
		Timestamp:          time.Now(),
	}
	blockData := previous_BlockHash + strconv.Itoa(nonce)
	newBlock.Hash = CalculateHash(blockData)

	chain.Blocks = append(chain.Blocks, newBlock)

	// Clear the transaction pool once the block is created
	chain.TransactionPool = nil

	return newBlock
}

// ListBlocks method to display the block data along with transactions
func ListBlocks() {
	for i, block := range chain.Blocks {
		fmt.Printf("\033[31mBlock %d:\033[0m\n", i)
		fmt.Printf("\033[32m\tTimestamp: %s\033[0m\n", block.Timestamp)
		fmt.Printf("\033[33m\tNonce(X): %d\033[0m\n", block.NonceX)
		fmt.Printf("\033[34m\tPrevious Block Hash: %s\033[0m\n", block.Previous_BlockHash)
		fmt.Printf("\033[35m\tHash: %s\033[0m\n", block.Hash)

		// Convert transaction pool to JSON format and display
		if len(block.TransactionPool) > 0 {
			transactionsJSON, _ := json.MarshalIndent(block.TransactionPool, "", "    ")
			fmt.Printf("\033[36m\tTransactions: %s\033[0m\n\n", transactionsJSON)
		} else {
			fmt.Println("\033[36m\tNo transactions in this block\033[0m\n")
		}
	}
}

// DisplayBlock displays the details of a specific block by its index
func DisplayBlock(blockNo int) {
	if blockNo < len(chain.Blocks) {
		block := chain.Blocks[blockNo]
		fmt.Printf("\033[31mBlock %d:\033[0m\n", blockNo)
		fmt.Printf("\033[32m\tTimestamp: %s\033[0m\n", block.Timestamp)
		fmt.Printf("\033[33m\tNonce(X): %d\033[0m\n", block.NonceX)
		fmt.Printf("\033[34m\tPrevious Block Hash: %s\033[0m\n", block.Previous_BlockHash)
		fmt.Printf("\033[35m\tHash: %s\033[0m\n", block.Hash)

		transactionsJSON, _ := json.MarshalIndent(block.TransactionPool, "", "    ")
		fmt.Printf("\033[36m\tTransactions: %s\033[0m\n", transactionsJSON)
	}
}
