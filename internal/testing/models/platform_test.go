package models

import (
	"fmt"
	"noir.github.com/aws_billing/internal/models"
	"testing"
)

func TestPlatform(t *testing.T) {
	models.DB = models.CreateDB()

	var user models.User
	user.Username = "test@highcloud.cn"

	user.Find(user.Username)

	fmt.Printf("%#v\n", user)

	var platform models.Platform
	platform.PlatformID = user.ID
	//_, err := platform.Create("gcp", user.ID)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	platform.Delete(user.ID, "gcp")

}
