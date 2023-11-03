ARG GO_VER=1.13.8
ARG ALPINE_VER=3.10

FROM golang:${GO_VER}-alpine${ALPINE_VER}

WORKDIR /go/src/github.com/hyperledger/fabric-samples/asset-transfer-private-data/chaincode-go/chaincode
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
# RUN go build -o /basic

EXPOSE 9999
CMD ["/chaincode"]
