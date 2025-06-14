package controllers

import (
	"backend/config"
	"backend/models"
	request "backend/request/check-cibil"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func OtpRequest(context *gin.Context) {

	var req request.RequestOtp

	if errReq := context.ShouldBindJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	fmt.Println(req)

	otp := GenerateOTP()

	identifierType := DetectType(req.UserIdentifier)

	saveOtp := models.Otp{
		Value:  req.UserIdentifier,
		OTP:    otp,
		Type:   identifierType,
		Status: "true",
	}

	fmt.Println(saveOtp)

	result := config.DB.Create(&saveOtp)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	context.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "opt has been sent to mobile ",
			"otp":     otp,
		},
	)
}

func OtpVerify(context *gin.Context) {

	var req request.ValidateOtp

	if errReq := context.ShouldBindJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	var otp models.Otp

	if err := config.DB.Where("value = ? AND otp = ? AND status = ?", req.UserIdentifier, req.Otp, "true").First(&otp).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "invalid otp",
			},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp verify"})
}

func OtpIndex(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp Index",
		},
	)
}

func OtpShow(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp show",
		},
	)
}

func OtpStore(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp Store",
		},
	)
}

func OtpUpdate(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp Update",
		},
	)
}

func OtpDestroy(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp Destroy",
		},
	)
}

func GenerateOTP() string {
	const otpLength = 6
	const digits = "0123456789"
	otp := make([]byte, otpLength)

	for i := 0; i < otpLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "" // silently return empty string on error
		}
		otp[i] = digits[num.Int64()]
	}

	return string(otp)
}

func DetectType(input string) string {
	if strings.Contains(input, "@") {
		return "email"
	}
	return "mobile"
}
