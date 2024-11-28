package Controllers

import (
	"go-gin-mysql/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRiwayatUpdateNilaiPembimbing(c *gin.Context) {

	dataRiwayat, err := Models.GetRiwayatNilaiPembimbing()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": dataRiwayat,
	})
}

func GetRiwayatUpdateNilaiFasilitator(c *gin.Context) {

	dataRiwayat, err := Models.GetRiwayatNilaiFasilitator()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": dataRiwayat,
	})
}

func GetJumlahPembimbing(c *gin.Context) {
	total, err := Models.GetJumlahPembimbing()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"data":    total,
		"message": "success",
	})
	return
}

func GetJumlahFasilitator(c *gin.Context) {
	total, err := Models.GetJumlahFasilitator()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"data":    total,
		"message": "success",
	})
	return
}

func GetJumlahHubin(c *gin.Context) {
	total, err := Models.GetJumlahHubin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"data":    total,
		"message": "success",
	})
	return
}

func GetJumlahSiswaPkl(c *gin.Context) {
	total, err := Models.GetJumlahSiswaPkl()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"data":    total,
		"message": "success",
	})
	return
}

func GetJumlahPemangku(c *gin.Context) {
	total, err := Models.GetJumlahPemangku()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"data":    total,
		"message": "success",
	})
	return
}

func GetJumlahWakel(c *gin.Context) {
	total, err := Models.GetJumlahWakel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"data":    total,
		"message": "success",
	})
	return
}
