# Builder
FROM golang:1.19.4-buster AS builder

WORKDIR /usr/src/app
COPY . .
RUN go mod tidy
RUN go get
RUN go install github.com/google/wire/cmd/wire@latest
RUN make di

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main .
RUN chmod +x main

# RUNNING
FROM alpine AS Production
COPY --from=builder /usr/src/app/main /main
COPY --from=builder /usr/src/app/.env /.env

CMD ["./main"]
