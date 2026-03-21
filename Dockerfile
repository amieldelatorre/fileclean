FROM --platform=$BUILDPLATFORM golang:1.26.1-alpine3.23 AS build
WORKDIR /build
COPY --chown=app:app go.mod ./
RUN go mod download
COPY ./ ./
RUN go build -o fileclean .

FROM alpine:3.23 AS final
COPY --from=build /build/fileclean /usr/local/bin/fileclean

ENTRYPOINT ["fileclean"]
