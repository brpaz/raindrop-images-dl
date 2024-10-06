# syntax = docker/dockerfile:1-experimental

# ==================================
# Base image
# ===================================
FROM --platform=$BUILDPLATFORM golang:1.22-alpine3.19 as base
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# ==================================
# Builder image
# ===================================
FROM base as builder

ARG VERSION
ARG GIT_COMMIT
ARG BUILD_DATE
ARG GC_FLAGS

ARG TARGETOS
ARG TARGETARCH

RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -gcflags "${GC_FLAGS}" \
    -ldflags="-w -s \
    -X github.com/brpaz/raindrop-images-dl/internal/version.Version=${VERSION} \
    -X github.com/brpaz/raindrop-images-dl/internal/version.GitCommit=${GIT_COMMIT} \
    -X github.com/brpaz/raindrop-images-dl/internal/version.BuildDate=${BUILD_DATE} \
    -extldflags '-static'" \
    -o /go/bin/raindrop-images-dl main.go

# ===================================
# production image
# ===================================
FROM alpine:3.20.3 as production

ARG UID=1000
ARG GID=1000

RUN apk add --no-cache curl ca-certificates && \
    addgroup -g ${GID} app && \
    adduser -D -u ${UID} -G app app

COPY --from=builder --chown=app:app /go/bin/raindrop-images-dl /bin/raindrop-images-dl

ENTRYPOINT [ "/bin/raindrop-images-dl" ]

USER app

LABEL org.opencontainers.image.title "Raindrop Images Downloader"
LABEL org.opencontainers.image.description "A cli tool to download images from Raindrop.io"
LABEL org.opencontainers.image.authors "Bruno Paz"
LABEL org.opencontainers.image.url "https://github.com/brpaz/raindrop-images-dl"
