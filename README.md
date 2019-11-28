# example

This example demonstrates the issues:
- https://github.com/golang/go/issues/35890
- https://github.com/golang/go/issues/35858.
- https://github.com/gonum/gonum/issues/1174

## Usage
Make sure you have go v1.13.4 set as default (or edit Makefile) and Docker installed on your system.
```
$ make build
$ make run
```

You will see something like this:
```
$ make build
go version go1.13.4 darwin/amd64
CGOENABLED=0 GOOS=linux GOARCH=amd64 go build ./main.go
docker build --memory 10Gb --memory-swap 0 -t mem_leak_issue_example .
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



$ make run
docker run -i -t --rm --name="mem_leak_issue_example" mem_leak_issue_example
  => JSON parsed! 13695 edges found (data size: 1 Mb)
        ---
        Alloc = 4 MiB           TotalAlloc = 5 MiB              StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 68 MiB
        HeapInuse = 4 MiB               HeapSys = 63 MiB                HeapIdle = 59 MiB               HeapReleased = 59 MiB           NumGC = 1
        ---
  => loading graph: 1 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 4778 MiB                TotalAlloc = 4784 MiB           StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 4999 MiB
        HeapInuse = 4780 MiB            HeapSys = 4799 MiB              HeapIdle = 19 MiB               HeapReleased = 19 MiB           NumGC = 5
        ---
  => time spent for #1 iteration: 43.563816952s

  => loading graph: 2 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 7958 MiB                TotalAlloc = 9564 MiB           StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 8327 MiB
        HeapInuse = 7960 MiB            HeapSys = 7999 MiB              HeapIdle = 39 MiB               HeapReleased = 39 MiB           NumGC = 6
        ---
  => time spent for #2 iteration: 50.665379005s

  => loading graph: 3 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 7198 MiB                TotalAlloc = 14344 MiB          StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 12463 MiB
        HeapInuse = 7200 MiB            HeapSys = 11967 MiB             HeapIdle = 4766 MiB             HeapReleased = 58 MiB           NumGC = 7
        ---
  => time spent for #3 iteration: 1m13.81873211s

  => loading graph: 4 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 11978 MiB               TotalAlloc = 19124 MiB          StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 14859 MiB
        HeapInuse = 11981 MiB           HeapSys = 14271 MiB             HeapIdle = 2290 MiB             HeapReleased = 927 MiB          NumGC = 8
        ---
  => time spent for #4 iteration: 1m52.920715546s

  => loading graph: 5 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 11236 MiB               TotalAlloc = 23903 MiB          StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 15664 MiB
        HeapInuse = 11238 MiB           HeapSys = 15039 MiB             HeapIdle = 3800 MiB             HeapReleased = 1672 MiB         NumGC = 8
        ---
  => time spent for #5 iteration: 1m34.722615361s

  => loading graph: 6 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 7957 MiB                TotalAlloc = 28683 MiB          StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 18043 MiB
        HeapInuse = 7962 MiB            HeapSys = 17343 MiB             HeapIdle = 9381 MiB             HeapReleased = 3340 MiB         NumGC = 9
        ---
  => time spent for #6 iteration: 1m25.087806114s

  => loading graph: 7 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 7480 MiB                TotalAlloc = 33462 MiB          StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 18043 MiB
        HeapInuse = 7484 MiB            HeapSys = 17343 MiB             HeapIdle = 9859 MiB             HeapReleased = 3635 MiB         NumGC = 10
        ---
  => time spent for #7 iteration: 1m15.462949042s

  => loading graph: 8 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 12259 MiB               TotalAlloc = 38242 MiB          StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 20424 MiB
        HeapInuse = 12263 MiB           HeapSys = 19647 MiB             HeapIdle = 7384 MiB             HeapReleased = 5758 MiB         NumGC = 11
        ---
  => time spent for #8 iteration: 2m7.055378632s

  => loading graph: 9 time
  => main dijkstra graph created!
  => all shortest created!
        ---
        Alloc = 11368 MiB               TotalAlloc = 43022 MiB          StackInuse = 0 MiB              StackSys = 0 MiB                Sys = 20426 MiB
        HeapInuse = 11371 MiB           HeapSys = 19647 MiB             HeapIdle = 8276 MiB             HeapReleased = 5914 MiB         NumGC = 11
        ---
  => time spent for #9 iteration: 1m33.04776087s

  => loading graph: 10 time
  => main dijkstra graph created!
make: *** [run] Error 137
```
Container is terminated by OOM killer.
