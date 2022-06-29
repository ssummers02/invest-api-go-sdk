# invest-api-go-sdk
Unofficial library for working with [InvestAPI](https://tinkoff.github.io/investAPI/) platforms  [Tinkoff Investments](https://www.tinkoff.ru/sl/AugaFvDlqEP)

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


## Using streams

```go
func listenTradeStream(ctx context.Context, services *pkg.ServicePool) {
	for {
		msg, err := services.OrderStreamInterface.Recv()
		if err != nil {
			log.Println(err)
		}

		orderTrades := msg.GetOrderTrades()
		if orderTrades != nil {
			log.Println(orderTrades)
		}

		select {
		case <-time.After(1 * time.Second):
			// pass
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	cfg := pkg.Config{
		Token:     "token",
		AccountID: []string{"account-id"},
	}

	services, err := pkg.NewServicePool(cfg)
	if err != nil {
		log.Println(err)
	}

	tradeStreamCtx, _ := context.WithCancel(context.Background())
	go listenTradeStream(tradeStreamCtx, services)

}

```
