# Build stage
FROM golang:latest as builder

# Install tzdata for timezone support in the builder stage (not strictly necessary but good practice)
RUN apt-get update && apt-get install -y tzdata

RUN mkdir /app

ADD . /app/
WORKDIR /app
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Final stage
FROM alpine:latest

# Install tzdata for timezone support in the final stage
RUN apk --no-cache add tzdata

WORKDIR /app/
COPY --from=builder /app .

# Set the timezone environment variable (optional)
ENV TZ=UTC 

CMD ["/app/app"]
