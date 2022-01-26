app-name = "emobi-service"

.PHONY: clean install unittest build docker run stop vendor migrate

# make install evn=init-local
install:
	sh scripts/install.sh
	sql-migrate up -env=$(env) -config=configs/dbconfig.yml

#make run env=dev
run:
	go run cmd/$(app-name)/main.go --e=$(env)

run-gateway:
	mvn -f ../../IdeaProjects/laporgub-api-gateway/pom.xml spring-boot:run

###### DATABASE MIGRATION ######

# make create-factories name="name"
create-factories:
	sql-migrate new -config=configs/dbconfig.yml -env=default-factories $(name)

# make create-migrations name="name"
create-migration:
	sql-migrate new -config=configs/dbconfig.yml -env=default-migrations $(name)

# make create-seeds name="name"
create-seeds:
	sql-migrate new -config=configs/dbconfig.yml -env=default-seeds $(name)

# make migrate env="default"
migrate:
	sql-migrate up -env=$(env)-init -config=configs/dbconfig.yml
	sql-migrate up -env=$(env)-factories -config=configs/dbconfig.yml
	sql-migrate up -env=$(env)-migrations -config=configs/dbconfig.yml
	sql-migrate up -env=$(env)-seeds -config=configs/dbconfig.yml

###### BUILD DEPLOY ######

#make build env="default"
build:
	git pull origin
	go mod tidy
	sql-migrate up -env=$(env)-init -config=configs/dbconfig.yml
	sql-migrate up -env=$(env)-factories -config=configs/dbconfig.yml
	sql-migrate up -env=$(env)-migrations -config=configs/dbconfig.yml
	sql-migrate up -env=$(env)-seeds -config=configs/dbconfig.yml
	go build cmd/$(app-name)/main.go

#make deploy env=apa
deploy:
	pm2 start --silent './main --e=$(env)' --watch --ignore-watch ./logs --name $(app-name)-$(env) --log=logs/app.log

###### TEST ######

test:
	go test ./...

coverage:
	go test -coverprofile cp.out ./...
	go tool cover -html=cp.out

mock:
	sh scripts/mock.sh

lint:
	golangci-lint run -D=typecheck
	gosec -exclude=G304 ./...
	revive ./...
	gocritic check ./...
	go-consistent -v ./...

# Sonar

sonar-start:
	/Users/default/Apps/sonarqube-8.2.0.32929/bin/macosx-universal-64/sonar.sh start

sonar-log:
	/Users/default/Apps/sonarqube-8.2.0.32929/bin/macosx-universal-64/sonar.sh start

sonar:
	  /Users/default/Apps/sonar-scanner-4.2.0.1873-macosx/bin/sonar-scanner   -Dsonar.projectKey=organization \
                                                                                  -Dsonar.sources=. \
                                                                                  -Dsonar.host.url=http://localhost:9000 \
                                                                                  -Dsonar.login=e8ceb53356694dcf67b0af096be298497358c05d