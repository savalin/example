APP_NAME=mem_leak_issue_example
GOBIN = go

## Hard limit for OOM killer
MEM_LIMIT=10Gb
## MEM_SWAP=-1 - unlimited swapping (global limit used instead)
## MEM_SWAP = MEM_LIMIT + swap
MEM_SWAP=10Gb

build:
	$(info ******** PLEASE CHECK YOUR GO VERSION: must be greater or equal to go1.13 to reproduce OOM killing after a few iterations ******** )
	$(GOBIN) version
	CGOENABLED=0 GOOS=linux GOARCH=amd64 $(GOBIN) build ./main.go
	docker build --memory $(MEM_LIMIT) --memory-swap $(MEM_SWAP) -t $(APP_NAME) .

build-force:
	docker build --no-cache --memory $(MEM_LIMIT) --memory-swap $(MEM_SWAP) -t $(APP_NAME) .

run:
	docker run -i -t --rm --memory $(MEM_LIMIT) --memory-swap $(MEM_SWAP) --name="$(APP_NAME)" $(APP_NAME)
