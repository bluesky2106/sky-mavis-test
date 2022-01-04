FROM golang:1.15.0-alpine3.12 AS build

ENV GO111MODULE=on

# Arguments
ARG WDIR=$GOPATH/src/github.com/bluesky2106/sky-mavis-test/backend

# install important packages
RUN apk add --no-cache ca-certificates git bash

# Create app directory in image filesystem 
RUN mkdir -p ${WDIR}
WORKDIR ${WDIR}
COPY . ${WDIR}

RUN CGO_ENABLED=0 go build -o server && \
    chmod u+x server

RUN mkdir -p /app
RUN cp ${WDIR}/server ${WDIR}/logs/.gitkeep ${WDIR}/resources/GeoLite2-City.mmdb /app/

FROM buildpack-deps:stretch-scm

ENV host="0.0.0.0"
ENV port="6000"
ENV env="production"
ENV mysql_conn_url="skymavis:skymavis@tcp(skymavis-mysql:3306)/skymavis_db?charset=utf8mb4&parseTime=True&loc=UTC"

COPY --from=build /app/server /server
COPY --from=build /app/.gitkeep /logs/.gitkeep
COPY --from=build /app/GeoLite2-City.mmdb /resources/GeoLite2-City.mmdb

CMD ["./server"]