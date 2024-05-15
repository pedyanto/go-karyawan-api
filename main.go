package main

import (
	"go-api/controllers"
	"go-api/database"

	"github.com/gin-gonic/gin"
)

func main() {
    database.InitDB()

    r := gin.Default()

    r.GET("/karyawans", controllers.GetKaryawans)
    r.POST("/karyawans", controllers.CreateKaryawan)
    r.GET("/karyawans/:id", controllers.GetKaryawanByID)
    r.PUT("/karyawans/:id", controllers.UpdateKaryawan)
    r.DELETE("/karyawans/:id", controllers.DeleteKaryawan)

    r.Run(":8080")
}
