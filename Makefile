.PHONY: grpc
grpc:
	docker build -t pantogram-badger --build-arg PROTOCOL="grpc" .

.PHONY: http
http:
	docker build -t pantogram-badger-http --build-arg PROTOCOL="http" .