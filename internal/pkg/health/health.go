package pkg

import (
	"net/http"
	"time"

	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// @BasePath /api/v1
// @Sumary Health API
// @Description Verify API health
// @Produce json
// @Tags Health
// @Success 200 {object} SendSuccessAuthResponse
// @Router /health [get]
func CheckApplicationHealth(c *gin.Context) {

	logger := config.GetLogger("Health")

	cpuUsage, err := GetCpuUsage()
	if err != nil {
		logger.Errorf("Failed to get CPU usage: %v", err)
		return
	}

	memoryUsage, err := GetMemoryUsage()
	if err != nil {
		logger.Errorf("Failed to get memory usage: %v", err)
		return
	}

	if cpuUsage < 80 && memoryUsage < 80 {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "unhealthy"})
	}

}

func GetCpuUsage() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}

func GetMemoryUsage() (float64, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return memInfo.UsedPercent, nil
}
