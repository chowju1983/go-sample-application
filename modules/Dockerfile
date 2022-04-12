FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /src/app/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN CGO_ENABLED=0 go build -o go lambda.go

FROM scratch
COPY --from=builder /src/app/go /var/app/
COPY --from=builder /src/app/db.sql /var/app/
EXPOSE 8080
ENV PORT 8080
ENV HOST <DATABASE_HOST>
CMD ["/var/app/lambda", "/var/app/scripts/db.sql"]
