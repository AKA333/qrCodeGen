package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

type requestInfo struct {
	Link string `json: "link"`
	Size string `json: "size"`
}

func Generate(c *gin.Context) {
	var reqInfo requestInfo
	err := c.BindJSON(&reqInfo)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":   "bad request",
				"details": err.Error(),
			},
		)
		return
	}

	if reqInfo.Size == "" {
		reqInfo.Size = "256"
	}

	size, err := strconv.Atoi(reqInfo.Size)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":   "enter valid size",
				"details": err.Error(),
			},
		)
		return
	}

	// timestamp := time.Now().Format("0601021504") // YYMMDDHHMM
	// filename := fmt.Sprintf("/data/qr_%s.png", timestamp)

	filename := "./data/qr.png"

	err = qrcode.WriteFile(reqInfo.Link, qrcode.Medium, size, filename)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":   "cannot generate qr code",
				"details": err.Error(),
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"OK": fmt.Sprintf("QR code generated with file name %s", filename),
		},
	)

}
