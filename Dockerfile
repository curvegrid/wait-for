FROM golang:1.13-alpine AS builder

WORKDIR /src

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o wait-for


FROM scratch

COPY --from=builder /src/wait-for /wait-for

ENTRYPOINT ["/wait-for"]
