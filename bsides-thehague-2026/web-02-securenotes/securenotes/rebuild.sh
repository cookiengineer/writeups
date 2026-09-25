#!/bin/bash

podman build -f Dockerfile -t securenotes .;
podman run -it --rm -p 8080:80 securenotes;

