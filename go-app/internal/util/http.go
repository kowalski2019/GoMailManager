package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func SendJsonResponse(c *gin.Context, statusCode int, res gin.H) {
     c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
     c.JSON(statusCode, res)
}

func SendMediaResponse(c *gin.Context, filePath string) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	c.File(filePath)
}
func SendResponse(c *gin.Context, statusCode int, statusName string, description string) {
	var res gin.H
	if description != "" {
		res = gin.H{"statusCode": statusCode, "statusName": statusName, "description": description}
	} else {
		res = gin.H{"statusCode": statusCode, "statusName": statusName}
	}
	SendJsonResponse(c, statusCode, res)
}

func SendOkResponse(c *gin.Context, res string) {
	SendResponse(c, http.StatusOK, "OK", res)
}
func SendBadRequestResponse(c *gin.Context, res string) {
	SendResponse(c, http.StatusBadRequest, "BadRequest", res)
}
func SendUnauthorizedResponse(c *gin.Context, res string) {
	SendResponse(c, http.StatusUnauthorized, "Unauthorized", res)
}
func SendForbiddenResponse(c *gin.Context, res string) {
	SendResponse(c, http.StatusForbidden, "Forbidden", res)
}
func SendNotFoundResponse(c *gin.Context, res string) {
	SendResponse(c, http.StatusNotFound, "NotFound", res)
}
func SendConflictResponse(c *gin.Context, res string) {
	SendResponse(c, http.StatusConflict, "Conflict", res)
}
func SendInternalServerErrorResponse(c *gin.Context, res string) {
	SendResponse(c, http.StatusInternalServerError, "InternalServerError", res)
}