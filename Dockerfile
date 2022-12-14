FROM golang:1.18-alpine3.16 AS build

ENV SEGFAUTILS_PORT 6893

RUN mkdir /segfautils
WORKDIR /segfautils
COPY . /segfautils/
RUN go mod download

EXPOSE 6893

RUN go build -o segfautils
RUN chmod +x segfautils
RUN go clean -modcache

FROM alpine:3.16 AS binary
COPY --from=build /segfautils/ /segfautils
WORKDIR /segfautils

CMD ["./segfautils"]