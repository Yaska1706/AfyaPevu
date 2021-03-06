package main

import (
	_ "github.com/yaska1706/AfyaPevu/patient-registration/docs"
	"github.com/yaska1706/AfyaPevu/patient-registration/internal/api"
)

// @AFYAPEVU
// @version 1.0
// @description API REST for AFYAPEVU patient registration

// @contact.name AfyaPevu
// @contact.url
// @contact.email

// @license.name MIT
// @license.url https://github.com/yaska1706/AfyaPevu/patient-registration/blob/master/LICENSE

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	api.Run("")
}
