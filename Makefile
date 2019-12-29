build:
	protoc -I. --go_out=plugins=micro:. \
      proto/vehicle/vehicle.proto
	docker build -t smartie-vehicle-service .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 smartie-vehicle-service