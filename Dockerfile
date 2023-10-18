# syntax=docker/dockerfile:1

FROM golang:1.20 as build
WORKDIR /app

ARG VERSION=v0.0.0
ENV CGO_ENABLED=1

COPY ./ ./
RUN go mod download
RUN go build -ldflags="-s -w -X https://github.com/AbderrahimBouhdida/AzH/constants.Version=$VERSION+docker"

FROM gcr.io/distroless/base-debian12:nonroot
LABEL org.opencontainers.image.source https://github.com/BloodHoundAD/AzureHound
COPY --from=build /app/azurehound /
ENTRYPOINT ["/azurehound"]
