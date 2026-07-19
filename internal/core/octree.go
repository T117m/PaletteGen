package core

import (
	"image"
	"image/color"
)

type ocTreeNode struct {
	children [8]*ocTreeNode
	count    int
	r, g, b  int
	isLeaf   bool
	next *ocTreeNode
}

const depth = 7

func OcTree(img image.Image, k int) color.Palette {
	var (
		width, height = getBounds(img)

		tree, insert, leafsCount = newOcTree()
	)

	for y := range height {
		for x := range width {
			insert(color.RGBAModel.Convert(img.At(x, y)))
		}
	}

	if *leafsCount > k {
		tree.reduce(leafsCount, k)
	}

	return tree.getPalette()
}

func (tree *ocTreeNode) getPalette() color.Palette {
	if tree == nil {
		return nil
	}

	if tree.isLeaf {
		return color.Palette{
			color.RGBA{
				uint8(tree.r / tree.count),
				uint8(tree.g / tree.count),
				uint8(tree.b / tree.count),
				255,
			},
		}
	}

	var p color.Palette

	for _, child := range tree.children {
		if child != nil {
			p = append(p, child.getPalette()...)
		}
	}

	return p
}

func (tree *ocTreeNode) reduce(leafs *int, k int) {
	for *leafs > k {
		candidate := tree.findReductionCandidate()

		if candidate == nil {
			break
		}

		totalR, totalG, totalB, totalCount, childrenCount := 0, 0, 0, 0, 0

		for _, child := range candidate.children {
			if child != nil {
				totalR += child.r
				totalG += child.g
				totalB += child.b
				totalCount += child.count
				childrenCount++
			}
		}

		candidate.isLeaf = true
		candidate.r = totalR
		candidate.g = totalG
		candidate.b = totalB
		candidate.count = totalCount
		candidate.children = [8]*ocTreeNode{}

		*leafs -= childrenCount - 1
	}
}

func (node *ocTreeNode) findReductionCandidate() *ocTreeNode {
	if node == nil || node.isLeaf {
		return nil
	}

	var (
		allLeaves = true

		best       *ocTreeNode
		bestWeight = int(^uint(0) >> 1)
	)

	for _, child := range node.children {
		if child == nil {
			continue
		}

		if !child.isLeaf {
			allLeaves = false
			candidate := child.findReductionCandidate()

			if candidate != nil && candidate.count < bestWeight {
				best = candidate
				bestWeight = candidate.count
			}
		}
	}

	if allLeaves {
		weight := 0

		for _, child := range node.children {
			if child != nil {
				weight += child.count
			}
		}

		if weight < bestWeight {
			best = node
			bestWeight = weight
		}
	}

	return best
}

func newOcTree() (head *ocTreeNode, insertFunc func(color.Color), leafs *int) {
	var (
		insert func(node *ocTreeNode, c color.Color, lvl, maxLvl int)
		sum    = 0
	)

	insert = func(node *ocTreeNode, c color.Color, lvl, maxLvl int) {
		var (
			r, g, b, _ = c.RGBA()
			r8, g8, b8 = uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)
		)

		if lvl == maxLvl {
			if node.count == 0 {
				sum++
			}

			node.count++
			node.r += int(r8)
			node.g += int(g8)
			node.b += int(b8)
			node.isLeaf = true

			return
		}

		var (
			bit  = depth - lvl
			rBit = checkBit(r8, bit)
			gBit = checkBit(g8, bit)
			bBit = checkBit(b8, bit)

			index = rBit*4 + gBit*2 + bBit
		)

		if node.children[index] == nil {
			node.children[index] = new(ocTreeNode)
		}

		insert(node.children[index], c, lvl+1, maxLvl)
	}

	head = new(ocTreeNode)

	insertFunc = func(c color.Color) {
		insert(head, c, 0, depth)
	}

	leafs = &sum

	return
}

func checkBit(ch uint8, bit int) int {
	return int((ch >> bit) & 1)
}
