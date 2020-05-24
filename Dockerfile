FROM golang:1.14.3-alpine3.11 as build

WORKDIR /pantogram

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o badger ./cmd

FROM alpine

WORKDIR /pantogram

COPY --from=build /pantogram/badger ./badger
COPY --from=build /pantogram/store/prefecture.json ./store/prefecture.json

CMD ["./badger"]