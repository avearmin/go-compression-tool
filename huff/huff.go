package huff

type Node interface {
  IsLeaf() bool
}

type LeafNode struct {
  Element rune
  Weight int
}

func (ln LeafNode) IsLeaf() bool {
  return true
}

type InternalNode struct {
  Left Node
  Right Node
  Weight int
}

func (in InternalNode) IsLeaf() bool {
  return false
}
