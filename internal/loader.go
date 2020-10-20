package internal

import (
	"math"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

type Loader interface {
	Load(edges []Edge)
	RoutesByAllShortest(from, to int64) ([][]graph.Node, float64)
}

func NewLoader(edgesLimit int) *loader {
	return &loader{
		edgesLimit:edgesLimit,
	}
}

type loader struct {
	edgesLimit int
	dg *simple.WeightedDirectedGraph
	allShortest path.AllShortest
}

func (l *loader) Load(edges []Edge) {
	l.dg = simple.NewWeightedDirectedGraph(0, math.Inf(1))

	counter := 1
	for i := range edges {
		if l.edgesLimit > 0 && counter > l.edgesLimit {
			break
		}
		var weight float64 = 99999 // default weight
		if edges[i].Weight > 0 {
			weight = edges[i].Weight
		}

		w, ok := l.dg.Weight(edges[i].From, edges[i].To)
		if ok && w <= weight {
			continue
		}

		weightedEdge := l.dg.NewWeightedEdge(simple.Node(edges[i].From), simple.Node(edges[i].To), weight)
		l.dg.SetWeightedEdge(weightedEdge)

		counter++
	}

	Log("main dijkstra graph created!")

	l.allShortest = path.DijkstraAllPaths(l.dg)

	Log("all shortest created!")
}

func (l *loader) RoutesByAllShortest(from, to int64) ([][]graph.Node, float64) {
	return l.allShortest.AllBetween(from, to)
}