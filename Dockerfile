###
# Stage 1: Build go binary
###
FROM golang:1.20.5-alpine AS build

WORKDIR /app

# Download and install dependencies first to save on future build times
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY main.go ./
COPY pkg ./pkg

# build binary
RUN CGO_ENABLED=0 go build -o /gowordle

###
# Stage 2: Copy binary into small image 
###
FROM scratch
COPY --from=build /gowordle /gowordle
# Copy word list into image
COPY words.txt ./
COPY solutions.txt ./
ENTRYPOINT ["/gowordle"]
