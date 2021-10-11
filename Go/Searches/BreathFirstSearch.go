package searches

import "reflect"

type Node struct {
	val   int
	nodes []*Node
}

func BreathFirstSearch(target int, root *Node) bool {
	if root.val == target {
		return true
	}
	visited := []*Node{root}
	return BreathFirstSearchHelper(target, root, &visited)
}

func BreathFirstSearchHelper(target int, node *Node, visited *[]*Node) bool {
	for _, elem := range node.nodes {
		*visited = append(*visited, elem)
		if elem.val == target {
			return true
		}
	}
	for _, elem := range node.nodes {
		res := BreathFirstSearchHelper(target, elem, visited)
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
