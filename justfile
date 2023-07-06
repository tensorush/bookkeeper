docker-test: (up "postgres") sleep migrate test down

up *SERVICES:
    docker compose --env-file ./configs/config.env up --build -d {{ SERVICES }}
    docker compose --env-file ./configs/config.env logs

sleep SECONDS="1":
    sleep {{ SECONDS }}

migrate ADDRESS="postgresql://root:secret@0.0.0.0:5432/bookkeeper?sslmode=disable" DIRECTION="up":
    migrate -path ./internal/db/migrations/ -database "{{ ADDRESS }}" -verbose {{ DIRECTION }}

migration NAME:
    migrate create -ext sql -dir ./internal/db/migrations/ -seq {{ NAME }}

test:
    go test -v -race -cover -short -count=1 -vet=off -coverprofile=cover.out -coverpkg=./... ./...
    go tool cover -html=cover.out -o cover.html

down:
    docker compose --env-file ./configs/config.env down

gen: sqlc protoc mockgen

sqlc:
    sqlc generate -f ./configs/sqlc.yaml

protoc:
    protoc --proto_path=./internal/proto/ --go_out=./internal/protogen/ --go_opt=paths=source_relative \
        --go-grpc_out=./internal/protogen/ --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=./internal/protogen/ --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=./dbdocs/swagger/ --openapiv2_opt=allow_merge=true,merge_file_name=bookkeeper \
        ./internal/proto/*.proto
    statik -src=./dbdocs/swagger/ -dest=./dbdocs/

mockgen:
    mockgen -package mockdb -destination ./internal/db/mock/store.go bookkeeper/internal/db Store
    mockgen -package mockwk -destination ./internal/worker/mock/distributor.go bookkeeper/internal/worker TaskDistributor

dbdocs:
    dbdocs build ./dbdocs/db.dbml
    dbml2sql --postgres -o ./dbdocs/schema.sql ./dbdocs/db.dbml
