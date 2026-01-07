FROM python:3.12-slim
WORKDIR /app
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

COPY services/python/requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt || true

COPY services/python/notification-service/app.py app.py
EXPOSE 5001
ENV PORT=5001
ENTRYPOINT ["python", "app.py"]


