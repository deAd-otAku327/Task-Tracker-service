FROM golang:1.24.3-alpine AS build

ENV GOPATH=/

WORKDIR /src

COPY ./ ./

RUN go mod download 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o task-tracker ./cmd/main.go

FROM alpine:latest

WORKDIR /root/configs

COPY --from=build /src/configs . 

WORKDIR /root/app

COPY --from=build /src/task-tracker . 

CMD ["./task-tracker"]