# fqdnIPLookup
Для запуска:

бд:
- docker run --name fqdnIPLookup-postgres -e POSTGRES_PASSWORD=mypass -d postgres

env variable для дб:
- export connStr="host=localhost user=postgres password=mypass dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

сервер имеет функцию раз в заданный интервал ( из конфига) обновлять текущую базу адресов из внешних dns серверов, 
для этого нужно указать:
- export waitTime="120s"

сервер имеет следующий функционал:
с флагом port можно указать порт запуска сервера


~/fqdnIPLookup/cmd/server go run . -port 9090

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

или воспользватся сваггером:
```
docker pull swaggerapi/swagger-ui
docker run -p 80:8080 -e SWAGGER_JSON=/foo/swagger.json -v .:/foo swaggerapi/swagger-ui         
```

запустит swagger ui на "http://localhost:80" 

swagger ui будет считать что сервер всегда будет запущен на "http://localhost:8080" 

если вы запустили сервер на другом порте, вам нужно поменять это значние в swagger.json
