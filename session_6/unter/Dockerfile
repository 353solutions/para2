FROM golang:1.25 AS build
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o unter

FROM debian:trixie-slim
COPY --from=build /app/unter /usr/local/bin
RUN groupadd -r unter && useradd --no-log-init -r -g unter unter
RUN mkdir /opt/unter
RUN chown unter /opt/unter
ENV UNTER_DB=/opt/unter/unter.db
USER unter
CMD ["unter"]
