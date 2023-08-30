include .env
export

.PHONY: all back_to_school encoded_string rot13 uncompress_me
PROGRAMMING_DIR:=./programming
GOBIN:=./bin

all:

back_to_school:
	@echo "run $@..."
	go run $(PROGRAMMING_DIR)/back_to_school/main.go -a $(PROG_TASK1_ADDRESS)
encoded_string:
	@echo "run $@..."
	go run $(PROGRAMMING_DIR)/encoded_string/main.go -a $(PROG_TASK2_ADDRESS)
rot13:
	@echo "run $@..."
	go run $(PROGRAMMING_DIR)/rot13/main.go -a $(PROG_TASK3_ADDRESS)
uncompress_me:
	@echo "run $@..."
	go run $(PROGRAMMING_DIR)/uncompress_me/main.go -a $(PROG_TASK4_ADDRESS)

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1 run ./... --max-same-issues 0

stats: build-stats
	./bin/stats -a $(ROOT_ME_API_URL) -k $(ROOT_ME_API_KEY) -u $(ROOT_ME_API_UID)

build-stats:
	bash scripts/build_stats.sh
