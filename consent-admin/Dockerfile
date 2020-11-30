FROM golang:1.15.2

USER root

RUN apt-get update
RUN curl -sL https://deb.nodesource.com/setup_14.x | bash -
RUN apt -y install nodejs

WORKDIR /app/web/app

COPY consent-admin/web/app/package.json package.json
COPY consent-admin/web/app/package-lock.json package-lock.json

RUN npm install --no-audit --prefer-offline

COPY consent-admin/web/app/tsconfig.json tsconfig.json
COPY consent-admin/web/app/public public
COPY consent-admin/web/app/src src

RUN npm run build

WORKDIR /app

COPY consent-admin .
COPY vendor vendor
COPY go.mod .
COPY go.sum .
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
