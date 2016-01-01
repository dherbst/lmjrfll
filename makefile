GOPATH := ${PWD}

all: test
	echo ${GOPATH}

test:
	GOPATH=${GOPATH} goapp test -coverprofile=coverage.out lmjrfll

get:
	GOPATH=${GOPATH} goapp get golang.org/x/tools/cmd/cover

show-cover:
	GOPATH=${GOPATH} goapp tool cover -html=coverage.out -o coverage.html
	open coverage.html
