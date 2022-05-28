FROM golang:1.17.9

ENV NODE=production
ENV PORT=8080
ENV JWT_SECRET=aSiAZgPRmmw7gN7p9WeQxQ==
ENV CONTEXT_PATH=/auth
ENV HOST_STRAPI_SERVICE=http://strapi-information-gateway:8080
ENV PATH_STRAPI_INFORMATION_GATEWAY=/strapi-information-gateway/api/strapi
ENV PATH_STRAPI_SIGNIN=/api/auth/local
ENV PATH_STRAPI_REFRESHTOKEN=/api/auth/refreshToken
ENV PATH_STRAPI_FORGET_PASSWORD=/api/auth/forgot-password
ENV PATH_STRAPI_RESET_PASSWORD=/api/auth/reset-password
ENV PATH_STRAPI_REGISTER=/api/auth/local/register
ENV AUTHEN_GATEWAY_HOST=http://authen_gateway:8080
ENV AUTHEN_GATEWAY_PATH_SIGNIN=/authen-listening/api/validate/call-signin
ENV AUTHEN_GATEWAY_PATH_GET_TOKENVERSION=/authen-listening/api/validate/token-version/
ENV AUTHEN_GATEWAY_PATH_CALL_REVOKE=/authen-listening/api/validate/call-revoke
ENV AUTHEN_GATEWAY_PATH_CALL_REFRESHTOKEN=/authen-listening/api/validate/call-check-refreshtoken
ENV LOG_PATH=/Users/ns/Documents/UniSUN/auth-listener/tmp/app.log
ENV PATH_STRAPI_CALLBACK_GOOGLE=/api/auth/google/callback
ENV PATH_STRAPI_CALLBACK_FACEBOOK=/api/auth/facebook/callback

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

EXPOSE 8080

CMD ["app"]