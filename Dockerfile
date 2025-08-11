# Stage 1 - Build Go binary and install Python deps
FROM golang:1.24 AS builder

# Install Python + pip
RUN apt-get update && apt-get install -y python3 python3-pip && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy Python requirements first for caching
COPY requirments.txt .
RUN pip3 install --no-cache-dir --break-system-packages -r requirments.txt

# Copy Go modules first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build Go app
RUN go build -o server main.go

# Stage 2 - Final runtime image (lighter)
FROM python:3.12-slim

WORKDIR /app

# Copy from builder
COPY --from=builder /app /app

# Install Python requirements at runtime
RUN pip install --no-cache-dir -r requirments.txt

EXPOSE 8000

CMD ["./server"]

