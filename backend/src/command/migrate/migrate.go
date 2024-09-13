package main

import (
	schema "app/module/config/repo/schema"
	"app/util/dbutil"
	"fmt"
)

func main() {
	fmt.Println("[+] START MIGRATION...")
	dbutil.InitDb()
	db := dbutil.Db()
	db.AutoMigrate(&schema.Variable{})
	fmt.Println("[+] DONE")
}
