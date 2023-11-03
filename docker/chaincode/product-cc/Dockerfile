ARG GO_VER=1.13.8
ARG ALPINE_VER=3.10

FROM golang:${GO_VER}-alpine${ALPINE_VER}

WORKDIR /go/src/github.com/hyperledger/productcontract
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
# RUN go build -o /product-chaincode

EXPOSE 9999
CMD ["productcontract"]