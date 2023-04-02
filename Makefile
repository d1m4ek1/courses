node_serve:
	npm run serve

node_build:
	npm run build

go_serve:
	go run ./cmd/main/main.go

go_build:
	go build ./cmd/main/main.go

go_codegangsta:
	gin -i --appPort 5050 --port 3000 -d ./cmd/main