FROM docker.io/library/debian:11-slim
RUN mkdir /app
WORKDIR /app
RUN apt update
RUN apt install curl -y
RUN apt install build-essential -y
RUN curl --proto '=https' --tlsv1.3 https://sh.rustup.rs -sSf -o rustup.sh
RUN chmod +x ./rustup.sh
RUN ./rustup.sh -y
ADD ./src /app/src
ADD ./static /app/static
ADD ./webfonts /app/webfonts
COPY Cargo.toml /app
COPY Cargo.lock /app
RUN /root/.cargo/bin/cargo build --release
RUN cp /app/target/release/shell_phone /app/shell_phone
CMD unset FLAG && \
    /app/shell_phone
