package main

import "log"

func main() {

	services, err := NewServicePool()
	if err != nil {
		log.Println(err)
	}

	accounts, err := services.UsersService.GetUserTariff()
	if err != nil {
		return
	}
	log.Println(accounts.UnaryLimits)
}
