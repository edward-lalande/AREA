#!/bin/bash

API_GATEWAY_URL="http://127.0.0.1:8080/ping"
USER_API_URL="http://127.0.0.1:8085/ping"
TIME_API_URL="http://127.0.0.1:8082/ping"
DISCORD_API_URL="http://127.0.0.1:8083/ping"

API_GATEWAY_HTTP_CODE=$(curl -s -o /dev/null -w '%{http_code}' $API_GATEWAY_URL)
USER_SERVICES_HTTP_CODE=$(curl -s -o /dev/null -w '%{http_code}' $USER_API_URL)
TIME_API_HTTP_CODE=$(curl -s -o /dev/null -w '%{http_code}' $TIME_API_URL)
DISCORD_API_HTTP_CODE=$(curl -s -o /dev/null -w '%{http_code}' $DISCORD_API_URL)

SUCCESS=true

if [ "$API_GATEWAY_HTTP_CODE" -ne 200 ]; then
    echo "Test échoué : API Gateway ne renvoie pas un code 200. Code reçu : $API_GATEWAY_HTTP_CODE"
    SUCCESS=false
fi

if [ "$USER_SERVICES_HTTP_CODE" -ne 200 ]; then
    echo "Test échoué : User Services ne renvoie pas un code 200. Code reçu : $USER_SERVICES_HTTP_CODE"
    SUCCESS=false
fi

if [ "$TIME_API_HTTP_CODE" -ne 200 ]; then
    echo "Test échoué : Time Services ne renvoie pas un code 200. Code reçu : $TIME_API_HTTP_CODE"
    SUCCESS=false
fi

if [ "$DISCORD_API_HTTP_CODE" -ne 200 ]; then
    echo "Test échoué : Discord API ne renvoie pas un code 200. Code reçu : $DISCORD_API_HTTP_CODE"
    SUCCESS=false
fi

if [ "$SUCCESS" = true ]; then
    echo "Test réussi : Toutes les routes renvoient un code 200."
    exit 0
else
    echo "Test échoué : Une ou plusieurs routes ne renvoient pas un code 200."
    exit 1
fi
