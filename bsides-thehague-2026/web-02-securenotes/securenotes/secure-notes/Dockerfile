FROM docker.io/library/python:3.11-slim

WORKDIR /app
COPY . /app
RUN pip install -r requirements.txt

EXPOSE 3000

CMD python wait_for_db.py && python main.py
