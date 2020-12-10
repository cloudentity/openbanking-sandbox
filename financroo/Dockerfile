FROM golang:1.15.2

USER root

RUN apt-get update
RUN curl -sL https://deb.nodesource.com/setup_14.x | bash -
RUN apt -y install nodejs

WORKDIR /app/web/app

COPY financroo/web/app/package.json package.json
COPY financroo/web/app/package-lock.json package-lock.json

RUN npm install --no-audit --prefer-offline

COPY financroo/web/app/tsconfig.json tsconfig.json
COPY financroo/web/app/public public
COPY financroo/web/app/src src

RUN npm run build

WORKDIR /app

COPY financroo .
COPY vendor vendor
COPY client client
COPY models models
COPY go.mod .
COPY go.sum .
RUN go build -o main .

EXPOSE 8091

CMD ["./main"]
