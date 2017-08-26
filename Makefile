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

binary_run:
	go build ./src/main.go && ./main -a 127.0.0.1:8080 -z data/data.zip

dep:
	go get -u github.com/valyala/fasthttp

post_example_user:
	curl -H "Content-Type: application/json" --data @data/example_new_user.json http://localhost:8080/users/new

post_example_visit:
	curl -H "Content-Type: application/json" --data @data/example_new_visit.json http://localhost:8080/visits/new

post_example_location:
	curl -H "Content-Type: application/json" --data @data/example_new_location.json http://localhost:8080/locations/new

post_all: post_example_user post_example_visit post_example_location