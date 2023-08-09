# fqdnIPLookup

Сервис для получения IP адресов из полных доменных имен (FQDN), для получения полных доменных имен из IP адресов, и для получения данных whois для домена 2 уровня.

Для запуска:

```bash
docker-compose build
docker-compose up
```

для запуска сваггера на "<http://localhost:90>":

```bash
docker pull swaggerapi/swagger-ui
sudo docker run -p 90:8080 -e SWAGGER_JSON=/foo/swagger.json -v .:/foo swaggerapi/swagger-ui
```

для теста можно послать запросы через curl:

получить все IP из fqdn

```bash
curl -X 'POST' \
  'http://localhost:8080/v1/fqdntoip' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[
  "ads.yahoo.com"
]'
```

получить все fqdn из IP

```bash
curl -X 'POST' \
  'http://localhost:8080/v1/iptofqdn' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[
  "87.248.119.251"
]'
```

получить whois информацию для sld

```bash
curl -X 'POST' \
  'http://localhost:8080/v1/whoishere' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[
  "ads.yahoo.com"
]'
```
