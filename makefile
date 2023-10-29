create-db-container:
	docker run --name blog-db -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=blog-db --network my-network -p 5432:5432 -d postgres

stop-db: 
	docker stop blog-db
	docker rm blog-db

postgres:
	docker exec -it blog-db psql -U postgres
