FROM golang:1.20

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
# Set the GOPROXY environment variable
ENV GOPROXY=https://goproxy.io,direct
# Set environment variable allow bypassing the proxy for specified repos (optional)
ENV GOPRIVATE=git.mycompany.com,github.com/my/private

RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app

EXPOSE 8000

CMD ["app"]