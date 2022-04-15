FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o atest .
RUN sleep 10

FROM alpine:3.14 as stage1
RUN ls
RUN sleep 10
RUN ls
COPY --from=builder /app/atest .
COPY /config/app.env /app/config/app.env

RUN chmod +x atest

ENTRYPOINT [ "./atest" ]