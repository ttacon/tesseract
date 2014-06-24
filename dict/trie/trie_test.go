package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := From([]string{"hello", "world", "hermano"})
	if !trie.Exists("hello") {
		t.Error("trie should have contained 'hello'")
	}

	if !trie.Exists("world") {
		t.Error("trie should have contained 'world'")
	}
}

type splitTest struct {
	input1, input2 string
	output         int
}

var splitTests = []splitTest{
	splitTest{
		input1: "hello",
		input2: "hello2",
		output: 5,
	},
	splitTest{
		input1: "hello2",
		input2: "hello",
		output: 5,
	},
	splitTest{
		input1: "",
		input2: "hello2",
		output: 0,
	},
	splitTest{
		input1: "yolo",
		input2: "hello2",
		output: 0,
	},
	splitTest{
		input1: "hello2",
		input2: "hello2",
		output: 6,
	},
}

func TestSplit(t *testing.T) {
	for i, test := range splitTests {
		if v := split(test.input1, test.input2); test.output != v {
			t.Errorf("test %d failed, got: %d, %v\n", i, v, test)
		}
	}
}

type commonEdgeTest struct {
	inputNodes  []*node
	inputString string
	outputNode  *node
}

var (
	helloNode = &node{val: "hello"}

	// TODO(ttacon): add more tests, nodes for tests
	commonEdgeTests = []commonEdgeTest{
		commonEdgeTest{
			inputNodes: []*node{
				helloNode,
			},
			inputString: "hel",
			outputNode:  helloNode,
		},
	}
)

func TestCommonEdge(t *testing.T) {
	for i, test := range commonEdgeTests {
		n := findCommonEdge(test.inputNodes, test.inputString)
		if n != test.outputNode {
			t.Errorf("Test %d failed, n: %v, test: %v\n", i, n, test)
		}
	}
}
