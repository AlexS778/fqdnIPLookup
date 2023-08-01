# fqdnIPLookup
Для запуска:

```
docker-compose build
docker-compose up
```
для сваггера:
```
docker pull swaggerapi/swagger-ui
sudo docker run -p 90:8080 -e SWAGGER_JSON=/foo/swagger.json -v .:/foo swaggerapi/swagger-ui
```


для теста можно послать запросы через curl:

получить все IP из fqdn
```
curl -X 'POST' \
  'http://localhost:8080/v1/FQDNToIP' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[
  "ads.yahoo.com"
]'
```

получить все fqdn из IP
```
curl -X 'POST' \
  'http://localhost:8080/v1/IPToFQDN' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[
  "87.248.119.251"
]'
```

получить whois информацию для sld
```
curl -X 'POST' \
  'http://localhost:8080/v1/whoishere' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[
  "ads.yahoo.com"
]'
```


если вы запустили сервер на другом порте, вам нужно поменять это значние в swagger.json
