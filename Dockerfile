FROM --platform=$BUILDPLATFORM golang:1.24.0-alpine3.21 AS build
WORKDIR /build
COPY --chown=app:app go.mod ./
RUN go mod download
COPY ./ ./
RUN go build -o fileclean .

FROM alpine:3.21.3 AS final
COPY --from=build /build/fileclean /usr/local/bin/fileclean

ENTRYPOINT ["fileclean"]
