package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yaska1706/AfyaPevu/patient-registration/internal/api/router"
	"github.com/yaska1706/AfyaPevu/patient-registration/internal/pkg/config"
	"github.com/yaska1706/AfyaPevu/patient-registration/internal/pkg/db"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
