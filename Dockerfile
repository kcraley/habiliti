# Compile binary in a build layer
FROM golang:1.17 as build
ARG OS=linux
ARG SOURCE=/go/src/github.com/kcraley/habiliti

WORKDIR ${SOURCE}

ADD go.mod .
ADD go.sum .
RUN go mod download

ADD . .
RUN CGO_ENABLED=0 GOOS=${OS} make -C ${SOURCE} build && \
    ls -lah ${SOURCE}

# Copy binary to production image
FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/habiliti /
ENTRYPOINT ["/habiliti"]
