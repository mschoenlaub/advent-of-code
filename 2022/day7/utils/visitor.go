package utils

import "math"

type Visitor interface {
	Visit(node *Node)
}

type SumSizeVisitor struct {
	sum    int
	Cutoff int
}

func (v *SumSizeVisitor) Visit(node *Node) {
	if len(node.Children) == 0 {
		return
	}
	size := node.Size()
	if size <= v.Cutoff {
		v.sum += size
	}
	for _, child := range node.Children {
		v.Visit(child)
	}
}

func (v *SumSizeVisitor) Sum() int {
	return v.sum
}

type MinSizeVisitor struct {
	Cutoff  int
	minSize int
}

func NewMinSizeVisitor(cutoff int) MinSizeVisitor {
	return MinSizeVisitor{Cutoff: cutoff, minSize: math.MaxInt}
}

func (v *MinSizeVisitor) Visit(node *Node) {
	if len(node.Children) == 0 {
		return
	}
	size := node.Size()
	if size >= v.Cutoff && size < v.minSize {
		v.minSize = size
	}
	for _, child := range node.Children {
		v.Visit(child)
	}
}

func (v *MinSizeVisitor) MinSize() int {
	return v.minSize
}
