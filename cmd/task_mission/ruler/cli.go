package main

type (
	CliAction interface {
		createModel()
		createRepository()
	}

	cli struct {
		commands []string
	}
)
