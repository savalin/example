package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"runtime"
	"runtime/debug"

	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

type edge struct {
	ID     int64   `json:"id"`
	From   int64   `json:"from"`
	To     int64   `json:"to"`
	Weight float64 `json:"weight"`
}

func main() {
	var file, err = os.Open("edges.json")
	if err != nil {
		panic(err)
	}

	var b []byte
	if b, err = ioutil.ReadAll(file); err != nil {
		panic(err)
	}

	var edges []edge
	if err = json.Unmarshal(b, &edges); err != nil {
		panic(err)
	}

	log("JSON parsed! %d edges found.", len(edges))

	printMemUsage()

	i := 1

	input := bufio.NewScanner(os.Stdin)
	for {
		log("Press Enter to start loading ⏎")
		input.Scan()

		log("  => loading graph: %d time", i)

		dg := simple.NewWeightedDirectedGraph(0, math.Inf(1))

		for _, e := range edges {
			var weight float64 = 99999 // default weight
			if e.Weight > 0 {
				weight = e.Weight
			}

			w, ok := dg.Weight(e.From, e.To)
			if ok && w <= weight {
				continue
			}

			weightedEdge := dg.NewWeightedEdge(simple.Node(e.From), simple.Node(e.To), weight)
			dg.SetWeightedEdge(weightedEdge)
		}

		log("  => DG created!")

		allShortest := path.DijkstraAllPaths(dg)

		log("  => all shortest created!")
		printMemUsage()

		// TRY TO FREE HEAP SYS MANUALLY
		debug.FreeOSMemory()
		log("after debug.FreeOSMemory():")
		printMemUsage()
		// =============================

		p, weight, unique := allShortest.Between(2, 3)

		log("  => path: %v, weight: %f, unique: %t\n", p, weight, unique)
		log("\n==============================================\n")

		i++
	}
}

func log(s string, args ...interface{}) {
	fmt.Printf(s+"\n", args...)
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("==================================\n")
	//fmt.Printf("\tAlloc\t|")
	//fmt.Printf("\tTotalAlloc\t")
	//fmt.Printf("\tSys\t")
	//fmt.Printf("\tHeapInuse\t")
	//fmt.Printf("\tHeapSys\t")
	//fmt.Printf("\tHeapIdle\t")
	//fmt.Printf("\tHeapReleased\t")
	//fmt.Printf("\tStackInuse\t")
	//fmt.Printf("\tStackSys\t")
	//fmt.Printf("\tNumGC\n", m.NumGC)
	//
	//fmt.Printf("\t%v\t|", bToMb(m.Alloc))
	//fmt.Printf("\t%v\t|", bToMb(m.TotalAlloc))
	//fmt.Printf("\t%v\t|", bToMb(m.Sys))
	//fmt.Printf("\t%v\t|", bToMb(m.HeapInuse))
	//fmt.Printf("\t%v\t|", bToMb(m.HeapSys))
	//fmt.Printf("\t%v\t|", bToMb(m.HeapIdle))
	//fmt.Printf("\t%v\t|", bToMb(m.HeapReleased))
	//fmt.Printf("\t%v\t|", bToMb(m.StackInuse))
	//fmt.Printf("\t%v\t|", bToMb(m.StackSys))
	//fmt.Printf("\t%v\n", m.NumGC)

	fmt.Printf("Alloc = %v MiB\t", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB\t", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB\t", bToMb(m.Sys))
	fmt.Printf("\tHeapInuse = %v MiB\t", bToMb(m.HeapInuse))
	fmt.Printf("\tHeapSys = %v MiB\t", bToMb(m.HeapSys))
	fmt.Printf("\tHeapIdle = %v MiB\t", bToMb(m.HeapIdle))
	fmt.Printf("\tHeapReleased = %v MiB\t", bToMb(m.HeapReleased))
	fmt.Printf("\tStackInuse = %v MiB\t", bToMb(m.StackInuse))
	fmt.Printf("\tStackSys = %v MiB\t", bToMb(m.StackSys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	fmt.Printf("==================================\n")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
