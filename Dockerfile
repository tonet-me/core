# Build Stage
# First pull Golang image
FROM docker.arvancloud.ir/golang:1.21.6-alpine as builder
 
# Set environment variable
ENV APP_NAME tonet
ENV CMD_PATH main.go

# Add a work directory
WORKDIR /$APP_NAME

## Cache and install dependencies
#COPY go.mod go.sum ./
#RUN go mod download

# Copy app files
COPY . .

# Budild application
RUN CGO_ENABLED=0 go build -mod=vendor -v -o $APP_NAME .

# Run Stage

FROM docker.arvancloud.ir/alpine:3.18 as development

 
# Set environment variable
ENV APP_NAME tonet
 
# Copy only required data into this image
COPY --from=builder /$APP_NAME .
 
# Expose application port
EXPOSE 1313
 
# Start app
CMD ./$APP_NAME