# debian lts with node lts
FROM golang:1.17.6-bullseye AS build

WORKDIR /go/src/slate-1

COPY . .

RUN go build -o /go/bin/slate-1 ./cmd/cloud


FROM debian
COPY --from=build /go/bin/slate-1 /bin/slate-1

EXPOSE 6969

COPY . .

ENTRYPOINT ["/bin/slate-1"]
