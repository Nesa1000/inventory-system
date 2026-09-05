# --- Stage 1: Build the binary ---
FROM golang:1.26-alpine AS builder

# set a directory for the app
WORKDIR /app

# copy requirements file and install dependencies (for optimal layer caching)
COPY go.mod go.sum ./
RUN go mod download

# copy all remaining files to the container
COPY . .

# build binary file
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --- Stage 2: Run the binary ---
FROM alpine:latest

WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/main .

# run the app
EXPOSE 8080
CMD ["./main"]
