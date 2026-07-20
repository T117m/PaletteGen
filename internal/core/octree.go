package core

import (
	"image"
	"image/color"
)

const depth = 7

type (
	ocTreeNode struct {
		children  [8]*ocTreeNode
		count     int
		r, g, b   int
		isLeaf    bool
		level     int
		next      *ocTreeNode
		leafCount int
	}

	ocTree struct {
		root      *ocTreeNode
		reducible [depth]*ocTreeNode
		leafs     int
	}
)

// OcTree строит палитру из k цветов с помощью октодерева
func OcTree(img image.Image, k int) color.Palette {
	var (
		tree = &ocTree{root: nil}

		width, height = getBounds(img)

		k57 = k
		p color.Palette
	)

	for y := range height {
		for x := range width {
			tree.insert(&tree.root, img.At(x, y), 0)
		}
	}

	if k < 8 && k != 1 {
		k = 8
	}

	for tree.leafs > k {
		tree.reduce()
	}

	p = tree.root.getPalette()

	if k57 >= 5 && k57 <= 7 {
		return p[:k57]
	}

	return p
}

func (node *ocTreeNode) getPalette() color.Palette {
	if node == nil {
		return nil
	}

	if node.isLeaf {
		return color.Palette{
			color.RGBA{
				uint8(node.r / node.count),
				uint8(node.g / node.count),
				uint8(node.b / node.count),
				255,
			},
		}
	}

	var p color.Palette

	for _, child := range node.children {
		if child != nil {
			p = append(p, child.getPalette()...)
		}
	}

	return p
}

func (tree *ocTree) reduce() {
	level := depth - 1

	for level >= 0 && tree.reducible[level] == nil {
		level--
	}
	if level < 0 {
		return
	}

	node := tree.reducible[level]
	tree.reducible[level] = node.next

	sum := 0
	for i := range 8 {
		child := node.children[i]
		if child == nil {
			continue
		}
		sum += child.leafCount
		if !child.isLeaf {
			tree.removeFromReducible(level+1, child)
		}
		node.children[i] = nil
	}

	node.isLeaf = true
	node.leafCount = 1
	tree.leafs = tree.leafs - sum + 1
}

func (tree *ocTree) removeFromReducible(level int, node *ocTreeNode) {
	if level < 0 || level >= depth || node == nil {
		return
	}

	if tree.reducible[level] == node {
		tree.reducible[level] = node.next
	} else {
		prev := tree.reducible[level]

		for prev != nil && prev.next != node {
			prev = prev.next
		}

		if prev != nil {
			prev.next = node.next
		}
	}

	if !node.isLeaf {
		for _, child := range node.children {
			if child != nil && !child.isLeaf {
				tree.removeFromReducible(level+1, child)
			}
		}
	}
}

func (tree *ocTree) insert(node **ocTreeNode, c color.Color, level int) {
	if *node == nil {
		*node = tree.newNode(level)
	}
	n := *node

	r, g, b, _ := c.RGBA()
	r8, g8, b8 := int(r>>8), int(g>>8), int(b>>8)

	n.count++
	n.r += r8
	n.g += g8
	n.b += b8

	if n.isLeaf {
		return
	}

	bit := 7 - level
	index := ((r8>>bit)&1)<<2 | ((g8>>bit)&1)<<1 | ((b8 >> bit) & 1)
	tree.insert(&n.children[index], c, level+1)
}

func (tree *ocTree) newNode(level int) *ocTreeNode {
	node := &ocTreeNode{
		level:     level,
		leafCount: 0,
	}

	if level == depth {
		node.isLeaf = true
		node.leafCount = 1
		tree.leafs++
	} else {
		node.next = tree.reducible[level]
		tree.reducible[level] = node
	}

	return node
}
