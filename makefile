
build:
	docker-compose -f deployments/docker-compose.yml down -v # 볼륨까지 한꺼번에 삭제
	docker-compose -f deployments/docker-compose.yml up --build

# test
