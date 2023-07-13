package main

import "GoAuth/database"

func main(){
	database.Connect("root:root@tcp(localhost:3306)/GoAuth?parseTime=true")
	database.Migrate()
}