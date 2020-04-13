package ethutil

import "testing"

func TestEthTrieHash(t *testing.T) {
	values := make(map[string]string, 0)
	values["test_1"] = "1"
	values["test_2"] = "2"
	values["test_3"] = "3"
	values["test_4"] = "4"

	hashVal := EthTrieHash(values)
	t.Log(hashVal)
}
