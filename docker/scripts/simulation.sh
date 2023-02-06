#!/bin/bash

curl --location --request POST '127.0.0.1:9999/scooters/connect' \
--header 'x-api-key: scooter_api_key' \
--header 'Content-Type: application/json' \
--data-raw '{
    "scooter_uuid": "cd651482-f10e-47d1-9f31-a77fd1fa343d",
    "mobile_uuid": "20587b2c-3969-49b6-add1-27fe09006ef9"
}'

x=1
while [ $x -le 10 ]
do
  docker-compose logs --tail=1 app

  x=$(( x + 1 ))

  sleep 2
done

curl --location --request POST '127.0.0.1:9999/scooters/disconnect' \
--header 'x-api-key: scooter_api_key' \
--header 'Content-Type: application/json' \
--data-raw '{
    "scooter_uuid": "cd651482-f10e-47d1-9f31-a77fd1fa343d",
    "mobile_uuid": "20587b2c-3969-49b6-add1-27fe09006ef9"
}'
