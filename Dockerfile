FROM golang:1.13-alpine

WORKDIR /go/src/neptune

#set go path
RUN export GOPATH=/go
RUN export PATH=$GOPATH:$PATH

# build
RUN GOOS=linux go build -ldflags "-X 'git.redbus.com/foo/foo/pkg/auth.GitCommit=$(git log -1 --pretty=%h)'" -a -o foo .

# run
ENTRYPOINT ["./foo"]
