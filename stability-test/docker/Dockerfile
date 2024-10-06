ARG ANUMAD_IMAGE
ARG ANUMAMINER_IMAGE

FROM ${ANUMAD_IMAGE} as anumad
FROM ${ANUMAMINER_IMAGE} as anumaminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/arumanetwork/anumad

WORKDIR /go/src/github.com/arumanetwork/anumad

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=anumad /app/ /app/
COPY --from=anumaminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/arumanetwork/anumad/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]