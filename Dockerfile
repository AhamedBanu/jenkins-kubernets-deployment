FROM golang:1.20-alpine3.16
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o main .
CMD ["/app/main"]


