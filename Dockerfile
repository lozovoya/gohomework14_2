FROM golang:1.14-alpine AS build
ADD . /app
ENV CGO_ENABLED=0
WORKDIR /app
RUN go build -o bank ./cmd/bank

FROM alpine:3.7
COPY --from=build /app/bank /app/bank
ENTRYPOINT ["/app/bank"]

EXPOSE 9999