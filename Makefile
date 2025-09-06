DB_URL=postgres://electrohub:secret@localhost:5432/electrohub?sslmode=disable

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down 1

migrate-force:
	migrate -path ./migrations -database "$(DB_URL)" force 1