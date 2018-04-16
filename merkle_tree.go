package merkle_tree

import (
	"errors"
	"hash"
	"io"
)

// MerkleTreeNode is the type of MerkleTree
type MerkleTreeNode struct {
	hashValue []byte
	left      *MerkleTreeNode
	right     *MerkleTreeNode
	height    int
}

// GetValue returns the hash value of the node
func (m *MerkleTreeNode) GetValue() []byte {
	return m.hashValue
}

var emptyNode = &MerkleTreeNode{
	hashValue: []byte(""),
}

// MerkleTree is a merkle tree =)
// Detail: https://en.wikipedia.org/wiki/Merkle_tree
type MerkleTree struct {
	hasher        hash.Hash
	bytesPerBlock int

	root *MerkleTreeNode
}

// NewMerkleTree returns a merkle tree by splitting data into several blocks of bytesPerBlock
// and building the merkel tree of the data
func NewMerkleTree(data io.Reader, bytesPerBlock int, hasher hash.Hash) (*MerkleTree, error) {
	if bytesPerBlock <= 0 {
		return nil, errors.New("")
	}

	if hasher == nil {
		return nil, errors.New("")
	}

	mTree := &MerkleTree{
		hasher:        hasher,
		bytesPerBlock: bytesPerBlock,
	}

	root, err := mTree.buildMerkleTree(data)
	if err != nil {
		return nil, err
	}
	mTree.root = root
	return mTree, nil
}

// createParent takes two tree node and calculate the hash value of their
// parent. Make them as children of their parent
func (m *MerkleTree) createParent(n1, n2 *MerkleTreeNode) *MerkleTreeNode {
	return &MerkleTreeNode{
		left:      n1,
		right:     n2,
		hashValue: m.hashBytes(n1.hashValue, n2.hashValue),
		height:    n1.height + 1,
	}
}

// hashBytes takes an array of []byte and calculate hash value by
// concating all elements
func (m *MerkleTree) hashBytes(byteArrays ...[]byte) []byte {
	m.hasher.Reset()
	for _, b := range byteArrays {
		m.hasher.Write(b)
	}
	return m.hasher.Sum(nil)
}

// buildMerkleTree contructs the merkle tree with the given data bottom up
func (m *MerkleTree) buildMerkleTree(data io.Reader) (*MerkleTreeNode, error) {
	buf := make([]byte, m.bytesPerBlock)
	que := make([]*MerkleTreeNode, 0)
	for {
		_, err := data.Read(buf)
		if err != nil {
			break
		}
		que = append(que, &MerkleTreeNode{
			hashValue: m.hashBytes(buf),
			height:    0,
		})
	}
	if len(que) == 0 {
		return nil, errors.New("")
	}
	for len(que) > 1 {
		n1 := que[0]
		var n2 *MerkleTreeNode
		if que[0].height == que[1].height {
			n2 = que[1]
			que = que[2:]
		} else {
			n2 = emptyNode
			que = que[1:]
		}
		que = append(que, m.createParent(n1, n2))
	}
	return que[0], nil
}

// Update updates the current merkle tree with the new data
// (TODO) Figure out how not to re-build the whole tree
func (m *MerkleTree) Update(data io.Reader) error {
	root, err := m.buildMerkleTree(data)
	if err != nil {
		return err
	}
	m.root = root
	return nil
}

func (m *MerkleTree) Verify(data io.Reader) bool {
	return false
}

// Root returns the root the merkle tree
func (m *MerkleTree) Root() *MerkleTreeNode {
	return m.root
}
