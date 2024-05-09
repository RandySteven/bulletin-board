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

MODEL_NAME := $(shell bash -c 'read -p "Model name : " modelfile; echo $$modelfile')
LOWER_FIRST_CHAR := $(shell echo $(MODELNAME) | cut -c1 | tr '[:upper:]' '[:lower:]')
UPPER_FIRST_CHAR := $(shell echo $(MODELNAME) | cut -c1)
MODEL_FILE := $(subst $(UPPER_FIRST_CHAR),$(LOWER_FIRST_CHAR),$(MODEL_NAME))
PACKAGE_MODEL := "package models"
INIT_MODEL := "type $(MODEL_NAME) struct {}"
MODEL_PATH := "entities/models/$(MODEL_FILE).go"
COMMAND_MODEL := "$(PACKAGE_MODEL)\n\n$(INIT_MODEL)" > $(MODEL_PATH)

make_model:
	@echo $(COMMAND_MODEL)