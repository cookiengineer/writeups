# docker build -t inspect-element . && docker run -p 1337:1337 -t inspect-element
FROM node:20-buster-slim

# Install Google Chrome
RUN apt-get update \
    && apt-get install -y wget gnupg \
    && wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
    && sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list' \
    && apt-get update \
    && apt-get install -y google-chrome-stable fonts-ipafont-gothic fonts-wqy-zenhei fonts-thai-tlwg fonts-kacst fonts-freefont-ttf libxss1 libx11-xcb1 \
    --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

# Install socat
RUN apt-get update && apt-get install -y socat

ARG FLAG=GPNCTF{fake_flag}
RUN echo "$FLAG" > /flag

# Using socat to bypass Chrome filtering out non-local connections
CMD \
socat tcp-listen:1337,fork tcp:localhost:13370 & \
google-chrome --remote-debugging-port=13370 --disable-gpu --headless=new --no-sandbox google.com

# Expose port 1337
EXPOSE 1337