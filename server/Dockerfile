FROM golang:1.21.4-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /go-mdeb-ledger

EXPOSE 8080

CMD [ "/go-mdeb-ledger" ]