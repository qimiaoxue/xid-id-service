FROM golang:1.18 as base

FROM base as built

WORKDIR /opt/app/api
COPY . .

RUN go get -d -v ./...
RUN go build -o /tmp/api-server ./*.go

FROM busybox

COPY --from=built /tmp/api-server /usr/bin/api-server

CMD ["api-server"]