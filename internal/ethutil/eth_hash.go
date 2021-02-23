package ethutil

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
)

func newHashTrie() *trie.Trie {
	db := memorydb.New()
	trieDb, _ := trie.New(common.Hash{}, trie.NewDatabase(db))
	return trieDb
}

func updateHashTrie(trie *trie.Trie, k, v string) {
	trie.Update([]byte(k), []byte(v))
}

func EthTrieHash(values map[string]string) string {
	if len(values) == 0 {
		return ""
	}
	trieDb := newHashTrie()
	for key, val := range values {
		updateHashTrie(trieDb, key, val)
	}
	hash := trieDb.Hash()

	return hash.String()
}
