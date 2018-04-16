package merkle_tree

import (
	"hash/fnv"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var randCharSet = "1234567890qwertyuiopasdfghjklzxcvbnm[]_-,./"

func randStr(n int) string {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = randCharSet[rand.Intn(len(randCharSet))]
	}
	return string(buf)
}

func TestCreateParent(t *testing.T) {
	m := &MerkleTree{
		hasher: fnv.New128(),
	}
	n1 := &MerkleTreeNode{
		hashValue: m.hashBytes([]byte("xxx")),
	}

	n2 := &MerkleTreeNode{
		hashValue: m.hashBytes([]byte("zzz")),
	}

	n3 := m.createParent(n1, n2)
	assert.Equal(t, n1, n3.left)
	assert.Equal(t, n2, n3.right)
	assert.Equal(t, m.hashBytes(n1.hashValue, n2.hashValue), n3.hashValue)
}

func TestSingleBlockFile(t *testing.T) {
	rStr := randStr(10)
	m, err := NewMerkleTree(strings.NewReader(rStr), 10, fnv.New128())
	assert.Nil(t, err)
	assert.Nil(t, m.root.left)
	assert.Nil(t, m.root.right)
	assert.Equal(t, m.hashBytes([]byte(rStr)), m.root.hashValue)
}

func TestMultipleBlocksFile(t *testing.T) {
	rStr := randStr(20)
	m, err := NewMerkleTree(strings.NewReader(rStr), 10, fnv.New128())
	assert.Nil(t, err)
	assert.Equal(t, m.hashBytes([]byte(rStr[:10])), m.root.left.hashValue)
	assert.Equal(t, m.hashBytes([]byte(rStr[10:])), m.root.right.hashValue)

	rStr = randStr(10240)
	m, err = NewMerkleTree(strings.NewReader(rStr), 10, fnv.New128())
	assert.Nil(t, err)
	assert.Equal(t, 10, m.Root().height)
}

func TestFileCompare(t *testing.T) {
	m1, _ := NewMerkleTree(strings.NewReader(randStr(200)), 10, fnv.New128())
	m2, _ := NewMerkleTree(strings.NewReader(randStr(150)), 10, fnv.New128())
	assert.NotEqual(t, m1.Root().GetValue(), m2.Root().GetValue())

	rStr := randStr(150)
	m1, _ = NewMerkleTree(strings.NewReader(rStr), 10, fnv.New128())
	m2, _ = NewMerkleTree(strings.NewReader(rStr), 10, fnv.New128())
	assert.Equal(t, m1.Root().GetValue(), m2.Root().GetValue())
}
