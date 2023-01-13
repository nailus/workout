FROM golang:1.20rc2-buster

COPY . /workout
WORKDIR /workout

RUN dpkg --install build/debs/migrate.linux-arm64.deb