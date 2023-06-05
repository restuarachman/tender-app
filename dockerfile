FROM golang:1.17-alpine3.15 as building-stage

WORKDIR /building-stage

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN  go build -o dockerized-myapp ./cmd/myapp

FROM alpine:3.15
WORKDIR /app
COPY --from=building-stage /building-stage/dockerized-myapp .
COPY --from=building-stage /building-stage/.env .

EXPOSE 9090

CMD [ "/app/dockerized-myapp" ]