build:
	docker build -t hl3 .

run:
	docker run --rm -p 80:80 -t hl3

compile:
	GOOS=linux GOARCH=amd64 go build ./src/main.go

tag: 
	docker tag hl3 stor.highloadcup.ru/travels/oceanic_coral

push: 
	docker push stor.highloadcup.ru/travels/oceanic_coral
