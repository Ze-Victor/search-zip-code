package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Sumary Health API
// @Description Verify API health
// @Tags Health
// @Success 200 {object} SendSuccessAuthResponse
// @Router /health [get]
func HealthAPI(c *gin.Context) {
	c.String(http.StatusOK, "Health check passed")
}
