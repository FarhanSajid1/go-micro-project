package main

import (
	"context"
	"log"

	user "github.com/farhansajid1/go-micro-project/user-service/proto/user"
	"github.com/micro/go-micro"
)

func newUser(name string) *user.User {
	temp := &user.User{
		Name:     name,
		Company:  "KpAcademy",
		Password: "random",
	}

	temp.Email = temp.Name + "." + temp.Company + "@gmail.com"

	return temp
}

func main() {
	service := micro.NewService(micro.Name("user.client"))
	service.Init()

	// name of the server is "greeter"
	client := user.NewUserService("go.micro.srv.user", service.Client())

	CreateNewUser := newUser("farhan")

	resp, err := client.Create(context.Background(), CreateNewUser)
	if err != nil {
		log.Printf("The error received was \n%v", err)
	}

	// attempting to login after creating the new user
	loginresp, err := client.Auth(context.Background(), CreateNewUser)
	if err != nil {
		log.Printf("The error received was \n%v", err)
	}

	log.Printf("==> Created user %v", resp.GetUser())
	log.Printf("==> Login successful user %v", loginresp.Token)

}
