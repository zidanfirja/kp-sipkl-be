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

}

func GetJumlahPemangku(c *gin.Context) {
	total, err := Models.GetTotalPemangku()
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

}

func JumlahSetiapPemangku(c *gin.Context) {

	dataJumlahPemangku, total, err := Models.GetJumlahSetiapPemangku()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	type RespJumlahPemangku struct {
		Total          int                     `json:"total"`
		DaftarPemangku []Models.JumlahPemangku `json:"daftar_pemangku"`
	}

	data := RespJumlahPemangku{
		Total:          total,
		DaftarPemangku: dataJumlahPemangku,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func TotalSiswaPklJurusan(c *gin.Context) {
	dataJumlahSiswa, total, err := Models.GetTotalSiswaPklJurusan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	type RespJumlahSiswa struct {
		Total       int                            `json:"total"`
		DaftarSiswa []Models.JumlahSiswaPklJurusan `json:"daftar_siswa"`
	}

	data := RespJumlahSiswa{
		Total:       total,
		DaftarSiswa: dataJumlahSiswa,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
