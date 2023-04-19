FROM golang:1.17 AS builder

ARG VERSION
ARG COMMIT

ADD . $GOPATH/src/github.com/salsadigitalauorg/lagoon-csp-collector/

WORKDIR $GOPATH/src/github.com/salsadigitalauorg/lagoon-csp-collector

ENV CGO_ENABLED 0

RUN apt-get install ca-certificates

RUN go build -ldflags="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}" -o build/lagoon-csp-collector

FROM scratch

ARG PORT=3000

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/salsadigitalauorg/lagoon-csp-collector/build/lagoon-csp-collector /usr/local/bin/lagoon-csp-collector

EXPOSE $PORT

CMD [ "lagoon-csp-collector", "-port", $PORT ]