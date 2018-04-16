package merkle_tree

import (
	"hash"
	"io"
)

type merkleTreeNode struct {
	hashValue     []byte
	left          *merkleTreeNode
	right         *merkleTreeNode
	height        int
	bytesPerBlock int
}

type MerkleTree struct {
	hasher hash.Hash

	root *merkleTreeNode
}

func (m *MerkleTree) buildMerkleTree(data io.Reader, bytesPerBLock int) *merkleTreeNode {
	return nil
}

func NewMerkleTree(data io.Reader, bytesPerBlock int, hasher hash.Hash) (*MerkleTree, error) {
	return nil, nil
}

func (m *MerkleTree) Update(data io.Reader) error {
	return nil
}

func (m *MerkleTree) Verify(data io.Reader) bool {
	return false
}
