#!/bin/bash

podman build -f Dockerfile -t htstaccess .;
podman run -it --rm -p 8080:80 htstaccess;

