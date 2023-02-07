package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lgcmotta/s3-poc/files"
)

type Response struct {
	Url string `json:"url"`
}

// Get Presigned URL without OTEL
// @Summary Get Presigned URL without OTEL
// @Description Get Presigned URL without OTEL
// @ID get.url.without.otel
// @Produce json
// @Param name path string true "The file name"
// @Success 200 {string} Response
// @Router /s3/files/{name} [get]
func GetFileURL() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileName := ctx.Param("name")
		url := files.GetFileURLWithoutOTEL(ctx.Request.Context(), fileName)

		response := Response{
			Url: url,
		}

		ctx.IndentedJSON(http.StatusOK, response)
	}
}

// Get Presigned URL with OTEL
// @Summary Get Presigned URL with OTEL
// @Description Get Presigned URL with OTEL
// @ID get.url.with.otel
// @Produce json
// @Param name path string true "The file name"
// @Success 200 {string} Response
// @Router /s3/files/otel/{name} [get]
func GetFileURLOtel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileName := ctx.Param("name")

		url := files.GetFileURLWithOTEL(ctx.Request.Context(), fileName)

		response := Response{
			Url: url,
		}

		ctx.IndentedJSON(http.StatusOK, response)
	}
}
