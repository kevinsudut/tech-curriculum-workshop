gobuild:
	@go build -v -o tech-curriculum-bgp cmd/tech-curriculum-bgp/*.go

gorun:
	make gobuild
	@./tech-curriculum-bgp

gomigrate:
	@go build -v -o migration cmd/migration/*go
	@./migration
	@sleep 1
	@rm -rf ./migration

composeup:
	@docker compose up -d

composedown:
	@docker compose down
