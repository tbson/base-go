package main

import (
	configSchema "app/service/config/schema"
	"app/util/db_util"
	"fmt"
)

func main() {
	fmt.Println("[+] START MIGRATION...")
	db_util.InitDb()
	db := db_util.Db()
	db.AutoMigrate(&configSchema.Variable{})
	fmt.Println("[+] DONE")
}
