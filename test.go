package main

import (
	"fmt"
	"github.com/jack410/factory_patterns/desp/factory"
	"github.com/jack410/factory_patterns/desp/models"
)

func main() {
	user := factory.CreateUser(factory.Aminuser)(123, "Kaka").(*models.Admin)
	fmt.Println(user)
}
