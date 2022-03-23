FROM python:3.9.2-slim-buster

RUN pip install flask gunicorn --no-cache-dir

WORKDIR /app
COPY app.py flag.txt ./
COPY templates templates
RUN mkdir /app/static && \
    chmod -R 775 . && \
    chmod 1773 static templates/errors && \
    mv flag.txt flag-$(cat /proc/sys/kernel/random/uuid).txt

CMD ["gunicorn", "-w16", "-t5", "--graceful-timeout", "0", "-unobody", "-gnogroup", "-b0.0.0.0", "app:app"]
