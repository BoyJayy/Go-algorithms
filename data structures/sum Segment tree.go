package main

import "fmt"

type SegTree struct {
	sz    int
	nodes []int
}

func (tree *SegTree) build(a []int) {
	tree.sz = 1
	for tree.sz < len(a) {
		tree.sz *= 2
	}
	n := tree.sz
	tree.nodes = make([]int, n*2)
	for i := 0; i < len(a); i++ {
		tree.nodes[i+n] = a[i]
	}
	for i := n - 1; i >= 1; i-- {
		tree.nodes[i] = tree.nodes[i<<1] + tree.nodes[i<<1|1]
	}
}

func (tree *SegTree) sum(v int, l int, r int, ls int, rs int) int {
	if l >= ls && r <= rs {
		return tree.nodes[v]
	}
	if r <= ls || l >= rs {
		return 0
	}
	sumr := tree.sum(v*2+1, (l+r)/2, r, ls, rs)
	suml := tree.sum(v*2, l, (l+r)/2, ls, rs)
	return sumr + suml
}

func (tree *SegTree) update(num int, v int, l int, r int, ls int, rs int) {
	if l >= ls && r <= rs {
		tree.nodes[v] = num
		return
	}
	if r <= ls || l >= rs {
		return
	}
	tree.update(num, v*2, l, (r+l)/2, ls, rs)
	tree.update(num, v*2+1, (l+r)/2, r, ls, rs)
	tree.nodes[v] = tree.nodes[v*2+1] + tree.nodes[v*2]
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	tree := SegTree{}
	tree.build(a)
	fmt.Println(tree.sum(1, 0, 4, 0, 4)) // so ofc it's gonna be 15
}
