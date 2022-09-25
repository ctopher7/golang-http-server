prerequisites :
- go v1.17

how to run:

- berat server:
```
go mod vendor && go run main.go berat development
```


api for berat server:
- get berat by id
```
curl --request GET \
  --url 'http://localhost:9000/berat?id=1'
```

- get all berat
```
curl --request GET \
  --url 'http://localhost:9000/berat
```

- insert berat
```
curl --request POST \
  --url http://localhost:9000/berat \
  --header 'Content-Type: application/json' \
  --data '{
	"tanggal":"2022-02-02T00:00:00Z",
	"max":40,
	"min":30
}'
```