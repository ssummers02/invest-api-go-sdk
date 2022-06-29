package main

import (
	"invest-api-go-sdk/pkg"
	"log"
)

func main() {
	cfg := pkg.Config{
		Token:     "t.Qh5RbKdX8aVcIeP1tFzhtTJCRdZhnkbXdIS8TLGrt0WNJXAFjaGjrFjRse0yp90Ic-BTnAH71nY5JfXChICsEg",
		AccountID: []string{"invest-api-go-sdk-test-account"},
	}

	services, err := pkg.NewServicePool(cfg)
	if err != nil {
		log.Println(err)
	}

	accounts, err := services.GetSandboxAccounts()
	if err != nil {
		log.Println(err)
	}

	log.Println(accounts)
}
