package searches

import "reflect"

type Node struct {
	val   int
	nodes []*Node
}

func DepthFirstSearch(target int, root *Node) bool {
	visited := []*Node{}
	return DepthFirstSearchHelper(target, root, &visited)
}

func DepthFirstSearchHelper(target int, node *Node, visited *[]*Node) bool {
	if node.val == target {
		return true
	}
	*visited = append(*visited, node)

	for _, elem := range node.nodes {
		res := DepthFirstSearchHelper(target, elem, visited)
		if res {
			return res
		}
	}
	return false
}

func contains(val Node, arr []Node) bool {
	for _, elem := range arr {
		if reflect.DeepEqual(val, elem) {
			return true
		}
	}
	return false
}
