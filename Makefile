all: clean build

build:
	go fmt
	go build
test:
	go test
clean:
	rm -f ./journey-planner
bench:
	go test -bench=.
cover:
	go test -cover
coverfunc:
	go test --coverprofile=cover.out
	go tool cover -func=cover.out
	rm -f cover.out
	install: clean deps build
	go install
	package: clean deps build
