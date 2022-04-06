FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o phuongne .

FROM alpine:3.14
EXPOSE 8080
COPY --from=builder /app/phuongne .

RUN chmod +x phuongne

ENTRYPOINT [ "./phuongne" ]