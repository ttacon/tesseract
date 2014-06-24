package dict

import "fmt"

type node struct {
	edges  []*node
	weight int
	val    string
}

type Trie interface {
	Add(val string)
	Exists(val string) bool
	Print()
}

var IgnoreWeights bool = true

func From(vals []string) Trie {
	trie := New()
	for _, val := range vals {
		trie.Add(val)
	}
	return trie
}

func New() Trie {
	return &node{}
}

func (n *node) Add(val string) {
	n.add(val)
}

func (n *node) Exists(val string) bool {
	return n.exists(val)
}

func (n *node) add(val string) {
	if splitVal := split(n.val, val); n.val != "" && splitVal != 0 { // !root and we need to act
		if len(n.val) == splitVal && len(val) == splitVal {
			// TODO(ttacon): when we keep track of occurences, add here
			return // nothing to add
		}

		if len(n.val) == splitVal { // add the rest of val as an edge
			n.edges = append(n.edges,
				&node{
					val: val[splitVal:],
				},
			)
			return
		}

		// we need to split our current val
		oldVal := n.val
		n.val = n.val[0:splitVal]
		newNode := &node{
			val:   oldVal[splitVal:],
			edges: n.edges,
		}
		n.edges = []*node{
			newNode,
			&node{
				val: val[splitVal:],
			},
		}
		return
	}

	if n.edges == nil {
		n.edges = []*node{
			&node{
				val: val,
			},
		}
		return
	}

	commonEdge := findCommonEdge(n.edges, val)
	if commonEdge == nil { // no common string
		n.edges = append(n.edges,
			&node{
				val: val,
			})
		return
	}

	commonEdge.add(val)
}

func (n *node) Print() {
	n.print("")
}

func (n *node) print(indent string) {
	fmt.Println("+", indent, n.val)
	for _, edge := range n.edges {
		edge.print(indent + "--")
	}
}

func findCommonEdge(edges []*node, val string) *node {
	for _, n := range edges {
		splitVal := split(n.val, val)
		if splitVal > 0 {
			return n
		}
	}
	return nil
}

func (n *node) exists(val string) bool {
	if n == nil {
		// no node was found from findCommonEdge
		return false
	}
	splitVal := split(n.val, val)
	if splitVal == len(val) {
		// check to see if val == the rest of what we're looking for
		return true
	}

	if splitVal == 0 && n.val != "" {
		return false
	}

	val = val[splitVal:]

	commonEdge := findCommonEdge(n.edges, val)
	return commonEdge.exists(val)
}

// TODO(ttacon): this will currently not support unicode,
// it will eventually, but for now don't worry about it
func split(a, b string) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	toRange := a
	if len(a) > len(b) {
		toRange = b
	}

	for i := range toRange {
		if i >= len(a) || i >= len(b) {
			return i - 1
		}
		if a[i] != b[i] {
			return i
		}
	}
	return len(toRange)
}
