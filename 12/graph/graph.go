package graph

type Graph struct {
	nodes map[int]*Node
}

func New() *Graph {
	g := Graph{}
	g.nodes = map[int]*Node{}
	return &g
}

type Node struct {
	id       int
	children []*Node
}

func (g *Graph) LinkNodes(id int, children []int) {
	parentNode, ok := g.nodes[id]
	if !ok {
		parentNode = &Node{id: id}
		g.nodes[id] = parentNode
	}

	for _, child := range children {
		childNode, ok := g.nodes[child]
		if !ok {
			childNode = &Node{id: child}
			g.nodes[child] = childNode
		}
		childNode.children = append(childNode.children, parentNode)
		parentNode.children = append(parentNode.children, childNode)
	}
}

func (g *Graph) Size(containing int) int {
	visited := map[int]bool{}
	queue := []int{containing}
	count := 0

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		if _, ok := visited[next]; ok {
			continue
		}
		visited[next] = true

		node := g.nodes[next]
		count++
		for _, c := range node.children {
			queue = append(queue, c.id)
		}
	}
	return count
}

func (g *Graph) Groups() int {
	visited := map[int]bool{}
	count := 0
	for id, _ := range g.nodes {
		if visited[id] {
			continue
		}
		count++

		queue := []int{id}

		for len(queue) > 0 {
			next := queue[0]
			queue = queue[1:]
			if visited[next] {
				continue
			}
			visited[next] = true

			node := g.nodes[next]
			for _, c := range node.children {
				queue = append(queue, c.id)
			}
		}
	}
	return count
}
