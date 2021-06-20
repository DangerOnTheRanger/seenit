FROM golang:1.15-buster AS builder

ENV GO111MODULE=on
WORKDIR /go/src/seenit
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
COPY cmd/ ./cmd/
WORKDIR /go/src/seenit/cmd/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o seenit .

FROM gcr.io/distroless/base:latest

WORKDIR /seenit
COPY --from=builder /go/src/seenit/cmd/seenit ./seenit
COPY templates/ templates/

CMD ["/seenit/seenit"]
