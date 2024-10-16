.DEFAULT_GOAL := build

.PHONY:fmt vet build clean run
fmt:
	# go fmt ./...
	go fmt ./exercises.go

vet: fmt
	# go vet ./...
	go vet ./exercises.go

build: vet
	# go build
	go build ./exercises.go

clean: build
	go clean

run: build
	./exercises