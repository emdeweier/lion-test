package main

import (
	"github.com/gin-gonic/gin"
	"lion-test/config"
	"lion-test/helper"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	db := config.NewDB()
	r := gin.Default()
	if helper.GoDotEnvVariable("GO_ENV") != "production" && helper.GoDotEnvVariable("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if helper.GoDotEnvVariable("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	config.NewRouter(r, db)
	err := r.Run(":" + helper.GoDotEnvVariable("GO_PORT"))
	if err != nil {
		return
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
