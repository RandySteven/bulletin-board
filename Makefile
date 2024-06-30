include /Users/randy.steven/others/bulletin-board/files/env/.env
export


yaml_file = ./files/yml/configs/task.local.yml
cmd_folder = ./cmd/task_mission/
gorun = @go run

ifeq ($(ENV),prod)
	yaml_file = ./files/yml/configs/task.prod.yml
else ifeq ($(ENV),staging)
	yaml_file = ./files/yml/configs/task.docker.yml
else ifeq ($(ENV),dev)
	yaml_file = ./files/yml/configs/task.local.yml
else
	$(error unknown variable in .env file)
endif

run:
	${gorun} ${cmd_folder}cmd -config ${yaml_file}

migration:
	${gorun} ${cmd_folder}migration -config ${yaml_file}

seed:
	${gorun} ${cmd_folder}seed -config ${yaml_file}

drop:
	${gorun} ${cmd_folder}drop -config ${yaml_file}

test_env:
	${yaml_file}

env_check:
	$(ENV)

refresh: drop migration seed

run-docker:
	docker compose up --build -d

stop-docker:
	docker compose down

run-docker-migration:
	docker compose up --build -d bulletin-migration

usecase_mockery: interfaces/usecases/*.go
	mockery --dir=interfaces/usecases \
           --name=$(shell basename $< -ext) \  # Extract filename without extension
           --filename=user_usecase.go \
           --output=mocks/usecasemocks \
           --outpkg=usecasemocks
.PHONY: usecase_mockery

test:
	go test -coverprofile=coverage.out -v ./... && go tool cover -html=coverage.out -o coverage.html

refresh_volume:
	docker system prune -a
#
#MODELNAME := $(shell bash -c 'read -p "Model name : " modelfile; echo $$modelfile')
#LOWER_FIRST_CHAR := $(shell echo $(MODELNAME) | cut -c1 | tr '[:upper:]' '[:lower:]')
#UPPER_FIRST_CHAR := $(shell echo $(MODELNAME) | cut -c1 )
#MODELFILE := $(subst $(UPPER_FIRST_CHAR),$(LOWER_FIRST_CHAR),$(MODELNAME))
#REPOSITORYFILE := "$(MODELFILE)_repository.go"
#PACKAGE_MODEL := "package models"
#PACKAGE_REPOSITORY := package repositories
#IMPORT_MODEL_PACKAGE := import(\n"\"task_mission/entities/models\""\n)
#INIT_MODEL := "type $(MODELNAME) struct {}"
#INIT_REPOSITORY_INTERFACE := type I$(MODELNAME)Repository interface{\nIRepository[models.$(MODELNAME)]\n}
#
#make_model:
#	@echo "$(PACKAGE_MODEL)\n\n$(INIT_MODEL)" > "entities/models/$(MODELFILE).go"
#
#make_repository:
#	@echo "$(PACKAGE_REPOSITORY)\n\n$(IMPORT_MODEL_PACKAGE)\n\n$(INIT_REPOSITORY_INTERFACE)" > "interfaces/repositories/$(MODELFILE)_repository.go"
#
#make_model_repo: make_model make_repository
