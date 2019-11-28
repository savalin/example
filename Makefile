APP_NAME=mem_leak_issue_example

## Hard limit for OOM killer
MEM_LIMIT=10Gb
## MEM_SWAP=-1 - unlimited swapping (global limit used instead)
## MEM_SWAP = MEM_LIMIT + swap
MEM_SWAP=10Gb

build:
	$(info ******** PLEASE CHECK YOUR GO VERSION: must be go1.13.4 to reproduce OOM killing after ~10-12 iterations ******** )
	$(info ******** HeapSys must be about 20Gb when OOM killer terminates the container ******** )
	go version
	CGOENABLED=0 GOOS=linux GOARCH=amd64 go build ./main.go
	docker build --memory $(MEM_LIMIT) --memory-swap $(MEM_SWAP) -t $(APP_NAME) .

build-force:
	docker build --no-cache --memory $(MEM_LIMIT) --memory-swap $(MEM_SWAP) -t $(APP_NAME) .

run:
	docker run -i -t --rm --name="$(APP_NAME)" $(APP_NAME)
