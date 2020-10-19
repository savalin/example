# example

SEE LAST VERSION IN [MASTER BRANCH](https://github.com/savalin/example/tree/master)

This example demonstrates the issues:
- https://github.com/golang/go/issues/35890
- https://github.com/golang/go/issues/35858.
- https://github.com/gonum/gonum/issues/1174

## Usage
Make sure you have go version >= v1.13 set as default (or edit Makefile) and Docker installed on your system.
```
$ make build
$ make run
```

You will see something like this:
```
$ make build
go version go1.15.1 darwin/amd64
CGOENABLED=0 GOOS=linux GOARCH=amd64 go build ./main.go
docker build --memory 10Gb --memory-swap 10Gb -t mem_leak_issue_example .
Sending build context to Docker daemon  53.38MB
Step 1/4 : FROM ubuntu:latest
 ---> 775349758637
Step 2/4 : COPY ./main .
 ---> f2df7c102986
Step 3/4 : COPY ./edges.json .
 ---> d21215b8d8de
Step 4/4 : CMD ["/main"]
 ---> Running in a1765ec8a72c
Removing intermediate container a1765ec8a72c
 ---> 550ae674af6d
Successfully built 550ae674af6d
Successfully tagged mem_leak_issue_example:latest
```
```
$ make run
docker run -i -t --rm --memory 10Gb --memory-swap 10Gb --name="mem_leak_issue_example" mem_leak_issue_example
  => JSON parsed! 13695 edges found (data size: 1 Mb)
        ---
        Alloc = 4 MiB           TotalAlloc = 5 MiB              StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 69 MiB
        HeapInuse = 4 MiB               HeapSys = 63 MiB                HeapIdle = 59 MiB               HeapReleased = 58 MiB           NumGC = 1
        ---
  => loading graph: 1 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 3110 MiB                TotalAlloc = 3116 MiB           StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 3266 MiB
        HeapInuse = 3113 MiB            HeapSys = 3135 MiB              HeapIdle = 22 MiB               HeapReleased = 22 MiB           NumGC = 5
        ---
        ---
        Alloc = 3110 MiB                TotalAlloc = 3116 MiB           StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 3266 MiB
        HeapInuse = 3113 MiB            HeapSys = 3135 MiB              HeapIdle = 22 MiB               HeapReleased = 22 MiB           NumGC = 5
        ---
  => time spent for #1 iteration: 13.849821349s

  => loading graph: 2 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 5286 MiB                TotalAlloc = 6228 MiB           StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 5525 MiB
        HeapInuse = 5289 MiB            HeapSys = 5311 MiB              HeapIdle = 22 MiB               HeapReleased = 22 MiB           NumGC = 6
        ---
        ---
        Alloc = 5286 MiB                TotalAlloc = 6228 MiB           StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 5525 MiB
        HeapInuse = 5289 MiB            HeapSys = 5311 MiB              HeapIdle = 22 MiB               HeapReleased = 22 MiB           NumGC = 6
        ---
  => time spent for #2 iteration: 15.055379515s

  => loading graph: 3 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 8398 MiB                TotalAlloc = 9339 MiB           StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 8788 MiB
        HeapInuse = 8402 MiB            HeapSys = 8447 MiB              HeapIdle = 44 MiB               HeapReleased = 44 MiB           NumGC = 6
        ---
        ---
        Alloc = 8398 MiB                TotalAlloc = 9339 MiB           StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 8788 MiB
        HeapInuse = 8402 MiB            HeapSys = 8447 MiB              HeapIdle = 44 MiB               HeapReleased = 44 MiB           NumGC = 6
        ---
  => time spent for #3 iteration: 25.699820995s

  => loading graph: 4 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 7539 MiB                TotalAlloc = 12451 MiB          StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 9398 MiB
        HeapInuse = 7543 MiB            HeapSys = 9023 MiB              HeapIdle = 1480 MiB             HeapReleased = 93 MiB           NumGC = 7
        ---
        ---
        Alloc = 7539 MiB                TotalAlloc = 12451 MiB          StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 9398 MiB
        HeapInuse = 7543 MiB            HeapSys = 9023 MiB              HeapIdle = 1480 MiB             HeapReleased = 93 MiB           NumGC = 7
        ---
  => time spent for #4 iteration: 14.367221161s

  => loading graph: 5 time
  => main dijkstra graph created!
make: *** [run] Error 137
```
Container is terminated by OOM killer.
