# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest
EXPOSE 5000

ENV SRC_DIR=/go/src/gitlab.com/karte/healthrecord-repository/
# Add the source code:
COPY . $SRC_DIR
WORKDIR $SRC_DIR

# Build it:
RUN export PATH=$PATH:$GOPATH/bin
RUN go get ./...
RUN go get -v github.com/kevinburke/go-bindata
RUN go get -v github.com/kevinburke/go-bindata/...
RUN go generate ./schema
RUN go build cmd/health_record_repository/main.go
ENTRYPOINT ["health_record_repository"]
