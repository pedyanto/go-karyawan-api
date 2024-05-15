package controllers

import (
	"net/http"

	"go-api/database"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func GetKaryawans(c *gin.Context) {
    var karyawans []models.Karyawan
    if err := database.DB.Find(&karyawans).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, karyawans)
}

func CreateKaryawan(c *gin.Context) {
    var karyawan models.Karyawan
    if err := c.ShouldBindJSON(&karyawan); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&karyawan).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, karyawan)
}

func GetKaryawanByID(c *gin.Context) {
    var karyawan models.Karyawan
    id := c.Param("id")
    if err := database.DB.First(&karyawan, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Karyawan not found"})
        return
    }
    c.JSON(http.StatusOK, karyawan)
}

func UpdateKaryawan(c *gin.Context) {
    var karyawan models.Karyawan
    id := c.Param("id")
    if err := database.DB.First(&karyawan, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Karyawan not found"})
        return
    }

    if err := c.ShouldBindJSON(&karyawan); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Save(&karyawan).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, karyawan)
}

func DeleteKaryawan(c *gin.Context) {
    var karyawan models.Karyawan
    id := c.Param("id")
    if err := database.DB.Delete(&karyawan, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Karyawan deleted"})
}
