FROM debian:trixie-slim
RUN apt-get update && \
    apt-get install -y python3 python3-pip python3-flask && \
    rm -rf /var/lib/apt/lists/*

# Copy app into a fixed staging directory first
COPY ./static /tmp/app/static
COPY ./paths /tmp/app/paths
COPY ./main.py /tmp/app/main.py

# Create a random /appXXXXXXXXXXXX directory and move files into it
CMD APPDIR="$(mktemp -d /appXXXXXXXXXXXX)" && \
    mv /tmp/app/* "$APPDIR"/ && \
    rmdir /tmp/app && \
    cd "$APPDIR" && \
    if [ -z "${FLAG}" ]; then export FLAG="FLAG{dummy}"; fi && \
    echo -n "${FLAG}" > flag.txt && \
    unset FLAG  && \
    flask --app main run --host 0.0.0.0
