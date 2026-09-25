FROM debian:11-slim
RUN apt update
RUN apt install -y build-essential socat
RUN mkdir /app
COPY ./main.c /app
WORKDIR /app
EXPOSE 1337
RUN gcc /app/main.c -o /app/babypwn
CMD if [ -z "${FLAG}" ]; then export FLAG="FLAG{dummy}"; fi && \
    echo -n "${FLAG}" > /app/flag.txt && \
    unset FLAG && \
    socat TCP-LISTEN:1337,reuseaddr,fork EXEC:/app/babypwn