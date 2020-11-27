package main

import (
	"gindemo/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

)





func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	
	r = CollectRoute(r)


	r.Run(":8080")
}






