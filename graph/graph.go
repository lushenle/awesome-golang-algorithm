package graph

import (
	"fmt"
	"sync"

	"github.com/lushenle/awesome-golang-algorithm/generic"
)

// Item the type of the binary search tree
type Item generic.Type

// Node a single node that composes the tree
type Node struct{ value Item }

// String format the value to string
func (n *Node) String() string { return fmt.Sprintf("%v", n.value) }

// ItemGraph the Items graph
type ItemGraph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(node *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, node)
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(node1, node2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*node1] = append(g.edges[*node1], node2)
	g.edges[*node2] = append(g.edges[*node2], node1)
	g.lock.Unlock()
}

// String for inspection purposes
func (g *ItemGraph) String() {
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}
