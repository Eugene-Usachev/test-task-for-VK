To generate proto files call `make` in the root of the project.

Before pushing to the repository:

- call `docker exec -i <container name> pg_dump -U admin -d container_monitoring --schema-only > ./db/schema.sql`;

- call `golangci-lint run --config=../golangci.yml` in `backend` and `pinger` directories.