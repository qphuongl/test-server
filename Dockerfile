FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o atest .

FROM alpine:3.14
COPY --from=builder /app/atest .
COPY /config/app.env /config/app.env

RUN chmod +x atest

ENTRYPOINT [ "./atest" ]