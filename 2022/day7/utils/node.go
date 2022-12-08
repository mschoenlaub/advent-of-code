package utils

type Node struct {
	Parent   *Node
	Children map[string]*Node
	Path     string
	size     int
}

func NewNode(parent *Node, path string, size int) Node {
	return Node{Parent: parent, Path: path, size: size, Children: make(map[string]*Node)}
}

func (n *Node) Size() int {
	if n.size == 0 {
		size := 0
		for _, child := range n.Children {
			size += child.Size()
		}
		n.size = size
	}
	return n.size
}

func (n *Node) AddChild(path string, size int) {
	child := NewNode(n, path, size)
	n.Children[path] = &child
}
