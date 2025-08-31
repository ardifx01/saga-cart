.PHONY: all gateway order_service product_service auth_service

all:
	@trap 'kill 0' SIGINT; \
	make gateway & \
	make order_service & \
	make product_service & \
	make payment_service & \
	make auth_service & \
	wait

gateway:
	@cd gateway && go run main.go

order_service:
	@cd order_service/cmd && go run main.go

product_service:
	@cd product_service/cmd && go run main.go

payment_service:
	@cd payment_service/cmd && go run main.go

auth_service:
	@cd auth_service/cmd && go run main.go
