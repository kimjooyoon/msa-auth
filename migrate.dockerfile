# Builder
FROM golang:1.19.4-buster AS builder

WORKDIR /usr/src/app
COPY . .
RUN go mod tidy
RUN go get
RUN go install github.com/google/wire/cmd/wire@latest
RUN make di

RUN cd cmd/migrate;CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main .
RUN cd cmd/migrate;chmod +x main

# RUNNING
FROM scratch AS Production
COPY --from=builder /usr/src/app/cmd/migrate/main /main
COPY --from=builder /usr/src/app/.env /.env

CMD ["./main"]
