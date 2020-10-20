package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	i "github.com/savalin/example/internal"
)

// PROBLEM:
// The app calculates all shortest paths and destroys data in loop (by same dataset loaded from json ~1Mb).
// Using docker limits memory=10Gb and memory-swap=10Gb we faced OOM killing after a few iterations.
//
// reproduced using configuration:
// Makefile (for docker container build):
//   - MEM_LIMIT=10Gb
//   - MEM_SWAP=10Gb
// Go version:
//   - >=go1.13
// Docker:
//   - Server: Docker Engine - Community v19.03.5
//   - Client: Docker Engine - Community v19.03.5
// Reproducible on Linux (tested on 5.4/5.4 kernels) and MacOS (tested on Catalina)

const edgesLimit = 9000 // please don't change this value. It's 'optimal' for success reproducing

var (
	current i.Loader
	old     i.Loader
)

func main() {
	var edges, err = readEdgesFile()
	if err != nil {
		panic(err)
	}

	i.PrintMemUsage()

	var n = 1
	for {
		i.Log("loading graph: %d time", n)
		start := time.Now()

		l := i.NewLoader(edgesLimit)
		l.Load(edges)

		i.PrintMemUsage()

		// some 'old' pointer to previous graph version
		old = current

		//doFakeUsage(old)

		// set new current graph
		current = l

		//doFakeUsage(current)

		// remove old pointer after all previous graph consumers graceful termination
		old = nil

		//// TRY TO FREE HEAP SYS MANUALLY
		//t := time.Now()
		//debug.FreeOSMemory()
		//i.Log("after debug.FreeOSMemory() (%v):", time.Since(t))
		//i.PrintMemUsage()

		i.Log("time spent for #%d iteration: %v\n", n, time.Since(start))
		n++

		//time.Sleep(time.Minute * 5)
	}
}

func readEdgesFile() ([]i.Edge, error) {
	var file, err = os.Open("edges.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var b []byte
	if b, err = ioutil.ReadAll(file); err != nil {
		return nil, err
	}

	var edges []i.Edge
	if err = json.Unmarshal(b, &edges); err != nil {
		return nil, err
	}

	i.Log("JSON parsed! %d edges found (data size: %d Mb)", len(edges), len(b)/1024/1024)

	return edges, nil
}

func doFakeUsage(l i.Loader) {
	if l == nil {
		return
	}

	var weight float64
	for k:=0; k<100;k++ {
		_,weight = l.RoutesByAllShortest(11, 12915)
	}
	i.Log("graph fake usage passed. %f", weight)
}