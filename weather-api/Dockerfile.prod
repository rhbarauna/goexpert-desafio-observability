FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o weather cmd/main.go cmd/wire_gen.go

FROM scratch
COPY --from=builder /app/cmd/.env ./app/cmd/.env
COPY --from=builder /app/weather .
ENTRYPOINT ["./weather"] 
