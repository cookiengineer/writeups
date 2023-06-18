#!/bin/bash

ROOT="$(pwd)";

build_app() {

	cd "${ROOT}/source";

	env CGO_ENABLED=1 GOOS=linux go build -x -a -ldflags '-Wl-rpath,/usr/lib' -ldflags '-extldflags "-static"' -o static ./main.go;

}

build_app;

