package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	proto "github.com/gr4yha7/maikuro/diablo/proto"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
)

type Users struct{}

var (
	usersService proto.UsersService
)

const (
	webServiceName = "go.micro.api.users"
	serviceName    = "go.micro.service.users"
)

func (u *Users) Index(c *gin.Context) {
	log.Println("received Users.Index API request")
	res := map[string]string{"message": "Index works"}
	c.JSON(http.StatusOK, res)
}

func (u *Users) Register(c *gin.Context) {
	// TODO Get user input and send it to RPC handler
	// TODO Validation?
	var user proto.RegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, user)

	// resp, err := usersService.Register(context.TODO())
}

func main() {
	service := web.NewService(
		web.Name(webServiceName),
	)

	service.Init()

	usersService = proto.NewUsersService(serviceName, client.DefaultClient)

	users := new(Users)
	router := gin.Default()
	router.GET("/users", users.Index)
	router.POST("/users", users.Register)

	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
