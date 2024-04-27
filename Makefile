yaml_file = ./files/yml/configs/task.local.yml
cmd_folder = ./cmd/task_mission/
gorun = @go run

run:
	${gorun} ${cmd_folder}cmd -config ${yaml_file}

migration:
	${gorun} ${cmd_folder}migration -config ${yaml_file}

seed:
	${gorun} ${cmd_folder}seed -config ${yaml_file}

drop:
	${gorun} ${cmd_folder}drop -config ${yaml_file}

refresh: drop migration seed