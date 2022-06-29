# invest-api-go-sdk


## Get started

Clone this repository first, and we are ready to go.

```shell script
$ git clone https://github.com/ssummers02/invest-api-go-sdk
```
## Create a first project

```go
func main() {
	cfg := pkg.Config{
		Token:     "token",
		AccountID: []string{"account-id"},
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
```