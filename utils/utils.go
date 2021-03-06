package utils

import (
	"fileshare/bootstrap"
	"log"
	"math/rand"
	"time"

	"github.com/skip2/go-qrcode"
)

func CreateRoute(id,route string) string {
	return "/v1/"+ route + "/" + id
}

func GenerateID(length int) string {
	var chars = []rune("0123456789abcdefghkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	id := make([]rune, length)
	for i := range id {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return string(id)
}

func CreateQR(url string) []byte {
	var png []byte
	qr, err := qrcode.New(bootstrap.App.Config.App.AppUrl+url, qrcode.Medium) 
	if err != nil {
		log.Println("Unable to create QR code: ", err)
	}
	qr.DisableBorder = true
	png, err = qr.PNG(128)
	if err != nil {
		log.Println("Unable to PNG for QR code: ", err)
	}
	return png
}
