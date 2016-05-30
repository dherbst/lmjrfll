GOPATH := ${PWD}

all: test
	echo ${GOPATH}

test:
	GOPATH=${GOPATH} goapp test -coverprofile=coverage.out lmjrfll

get:
	GOPATH=${GOPATH} goapp get -u golang.org/x/tools/cmd/cover google.golang.org/appengine golang.org/x/net/context

show-cover:
	GOPATH=${GOPATH} goapp tool cover -html=coverage.out -o coverage.html
	open coverage.html
