# ---------- Build stage ----------
FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app ./cmd/server

# ---------- Runtime stage ----------
FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=build /bin/app /app
ENV HTTP_ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/app"]