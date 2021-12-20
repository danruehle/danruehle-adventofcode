package dan

import (
	"strings"
	"sync"
	"testing"
)

var day12input = `lg-GW
pt-start
pt-uq
nx-lg
ve-GW
start-nx
GW-start
GW-nx
pt-SM
sx-GW
lg-end
nx-SM
lg-SM
pt-nx
end-ve
ve-SM
TG-uq
end-SM
SM-uq`

type caveKind int

const (
	big caveKind = iota
	small
)

type caveNode struct {
	id     string
	kind   caveKind
	routes []*caveNode
}

func newCaveNode(id string) *caveNode {
	kind := small
	if strings.ToUpper(id) == id {
		kind = big
	}
	return &caveNode{id: id, kind: kind}
}

func (n *caveNode) addRoute(node *caveNode) {
	n.routes = append(n.routes, node)
	node.routes = append(node.routes, n)
}

type caveMap struct {
	nodes map[string]*caveNode
	start *caveNode
	end   *caveNode
}

func newCaveMap(input string) *caveMap {
	m := &caveMap{nodes: make(map[string]*caveNode)}
	for _, line := range strings.Split(input, "\n") {
		ids := strings.Split(line, "-")
		node0 := m.getNode(ids[0])
		node1 := m.getNode(ids[1])
		node0.addRoute(node1)
	}
	return m
}

func (m *caveMap) getNode(id string) *caveNode {
	node, ok := m.nodes[id]
	if ok {
		return node
	}
	node = newCaveNode(id)
	m.nodes[node.id] = node
	if id == "start" {
		m.start = node
	} else if id == "end" {
		m.end = node
	}
	return node
}

type caveRoute []*caveNode

func (r caveRoute) String() string {
	sb := strings.Builder{}
	for _, node := range r {
		sb.WriteString(node.id)
		sb.WriteString("-")
	}
	return sb.String()[0 : sb.Len()-1]
}

func (r caveRoute) Copy() caveRoute {
	c := make(caveRoute, len(r))
	copy(c, r)
	return c
}

func (n *caveNode) traverseToEnd(t *testing.T, route caveRoute) []caveRoute {
	currentRoute := route.Copy()
	currentRoute = append(currentRoute, n)
	if n.id == "end" {
		t.Logf("found full route: %s", currentRoute)
		return []caveRoute{currentRoute}
	}
	if n.kind == small {
		for _, node := range route {
			if n.id == node.id {
				return nil
			}
		}
	}
	var completeRoutes []caveRoute
	for _, node := range n.routes {
		newCompleteRoutes := node.traverseToEnd(t, currentRoute)
		if newCompleteRoutes != nil {
			completeRoutes = append(completeRoutes, newCompleteRoutes...)
			t.Logf("found new complete routes\n\tcurrent route: %s\n\tevaluated node: %s\n\tnewCompleteRoutes: %s\n\tcompleteRoutes: %s", currentRoute, node.id, newCompleteRoutes, completeRoutes)
		}
	}
	if completeRoutes != nil {
		t.Logf("returning complete routes\n\t%s\n\t%s", currentRoute, completeRoutes)
	}
	return completeRoutes
}

func (n *caveNode) traverseToEndB(t *testing.T, route caveRoute, smallCaveLimit int, c chan string) {
	currentRoute := route.Copy()
	currentRoute = append(currentRoute, n)
	if n.id == "end" {
		c <- currentRoute.String()
		return
	}
	if len(currentRoute) > 1 && n.id == "start" {
		return
	}
	if n.kind == small {
		visitedCount := 0
		for _, node := range route {
			if n.id == node.id {
				visitedCount++
				if visitedCount >= smallCaveLimit {
					return
				}
			}
		}
		if visitedCount == 1 && smallCaveLimit == 2 {
			smallCaveLimit = 1
		}
	}
	if len(currentRoute)%5 == 0 {
		var wg sync.WaitGroup
		for _, node := range n.routes {
			wg.Add(1)
			go func(n *caveNode) {
				defer wg.Done()
				n.traverseToEndB(t, currentRoute, smallCaveLimit, c)
			}(node)
		}
		wg.Wait()
	} else {
		for _, node := range n.routes {
			node.traverseToEndB(t, currentRoute, smallCaveLimit, c)
		}
	}
}

func TestDay12a(t *testing.T) {
	caveMap := newCaveMap(day12input)
	routes := caveMap.start.traverseToEnd(t, nil)
	t.Log("routes:")
	for i, route := range routes {
		t.Logf("%d: %s", i, route)
	}
	t.Logf("number of routes: %d", len(routes))
}

func TestDay12b(t *testing.T) {
	caveMap := newCaveMap(day12input)
	completeRoutes := 0
	c := make(chan string, 50)
	go func() {
		caveMap.start.traverseToEndB(t, nil, 2, c)
		close(c)
	}()
	for r := range c {
		completeRoutes++
		t.Logf("route %d: %s", completeRoutes, r)
	}
	t.Logf("number of routes: %d", completeRoutes)
}
