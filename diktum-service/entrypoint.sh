#!/bin/sh

echo "📥 Fetching latest public_key.pem from auth-service..."
# Tunggu auth-service siap
until curl --output /dev/null --silent --head --fail http://auth-service:8081/public-key; do
  echo "⏳ Waiting for auth-service to be ready..."
  sleep 3
done

# Download key
curl -s http://auth-service:8081/public-key -o /app/keys/public_key.pem
echo "✅ Downloaded public_key.pem"

# Jalankan aplikasinya
echo "🚀 Starting diktum-service..."
exec ./diktum-service
