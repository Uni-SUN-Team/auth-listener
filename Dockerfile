FROM golang:1.17.9

ENV NODE=production
ENV PORT=8080
ENV CONTEXT_PATH=/auth
ENV VERSION=1.0.0
ENV HOST=api.unisun.dynu.com
ENV GIN_MODE=release
ENV JWT_SECRET=aSiAZgPRmmw7gN7p9WeQxQ==

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

EXPOSE 8080

CMD ["app"]