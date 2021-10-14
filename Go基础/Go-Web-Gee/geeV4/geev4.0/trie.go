package geev4_0

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{ pattern = %s, part = %s, isWild= %t\n", n.pattern, n.part, n.isWild)
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height)

}

func (n *node) search(parts []string, heigh int) *node {
	if len(parts) == heigh||strings.HasPrefix(n.part,"*") {
		return nil
	}
	return n
}
// travels  the nodes
func (n *node) travel(list *[]*node)  {
	if n.pattern!="" {
		*list=append(*list,n)
	}
	for _, child := range n.children {
	    child.travel(list)
	}

}
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}
func (n *node) matchChildren(part string) []*node {
	nodes:=make([]*node,0)
	for _,child := range n.children {
		if child.part == part||child.isWild {
			nodes=append(nodes,child)
		}

	}
	return nodes
}
