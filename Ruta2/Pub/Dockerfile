# Start from the latest golang base image
FROM golang:1.15.6-alpine AS build


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY ./RedisPub.go .

# Build the Go app
RUN go build -o /project/rpub


######## Start a new stage from scratch #######
# Stage 1
FROM alpine:latest
COPY --from=build /project /project
WORKDIR /project
ENTRYPOINT ["./rpub"]
