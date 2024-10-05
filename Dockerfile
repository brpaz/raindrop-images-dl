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

ARG BUILDTIME
ARG VERSION
ARG REVISION
ARG GC_FLAGS

ARG TARGETOS
ARG TARGETARCH

RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -gcflags "${GC_FLAGS}" \
    -ldflags="-w -s \
    -X github.com/brpaz/raindrop-images-dl/internal/core/version/version.Version=${VERSION} \
    -X github.com/brpaz/raindrop-images-dl/internal/core/version/version.Revision=${REVISION} \
    -X github.com/brpaz/raindrop-images-dl/internal/core/version/version.BuildTime=${BUILDTIME} \
    -extldflags '-static'" \
    -o /go/bin/api ./cmd/api

RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -gcflags "${GC_FLAGS}" \
    -ldflags="-w -s \
    -X github.com/brpaz/raindrop-images-dl/internal/core/version/version.Version=${VERSION} \
    -X github.com/brpaz/raindrop-images-dl/internal/core/version/version.Revision=${REVISION} \
    -X github.com/brpaz/raindrop-images-dl/internal/core/version/version.BuildDate=${BUILDTIME} \
    -extldflags '-static'" \
    -o /go/bin/migrate ./cmd/migrations

# ==================================
# dlv image
# By being a sperate image, we can cache it
# ===================================
FROM --platform=$BUILDPLATFORM golang:1.22-alpine3.18 as delve

SHELL ["/bin/ash", "-o", "pipefail", "-c"]

RUN apk add --no-cache curl ca-certificates

RUN go install -ldflags "-extldflags '-static'" \
    -v github.com/go-delve/delve/cmd/dlv@v1.21.0

# ==================================
# Dev image
# ===================================
FROM base as dev

ARG UID=1000
ARG GID=1000

ENV GOMODCACHE=/tmp/go/pkg/mod

RUN apk add --no-cache curl && \
    addgroup -g ${GID} app && \
    adduser -D -u ${UID} -G app app

COPY --from=delve /go/bin/dlv /usr/local/bin/dlv
COPY --from=cosmtrek/air:v1.52.3 /go/bin/air /usr/local/bin/air

COPY docker/docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]

RUN chown -R app:app /src
USER app

VOLUME [ "/src" ]

CMD ["air"]

# ===================================
# production image
# ===================================
FROM alpine:3.20.3 as production

ARG UID=1000
ARG GID=1000

RUN apk add --no-cache curl ca-certificates && \
    addgroup -g ${GID} app && \
    adduser -D -u ${UID} -G app app

COPY --from=builder --chown=app:app /go/bin/api /bin/api

ENTRYPOINT [ "/bin/api" ]

USER app

LABEL org.opencontainers.image.title "Raindrop Images Downloader"
LABEL org.opencontainers.image.description "A cli tool to download images from Raindrop.io"
LABEL org.opencontainers.image.authors "Bruno Paz"
LABEL org.opencontainers.image.url "https://github.com/brpaz/raindrop-images-dl"
