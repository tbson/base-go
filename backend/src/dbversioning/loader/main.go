package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"

	account "src/module/account/schema"
	config "src/module/config/schema"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
		&config.Variable{},
		&account.Tenant{},
		&account.AuthClient{},
		&account.User{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
