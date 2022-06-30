package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ssummers02/invest-api-go-sdk/pkg"
	pb "github.com/ssummers02/invest-api-go-sdk/pkg/investapi"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	cfg := pkg.Config{
		Token:     "",
		AccountID: []string{"2179550658"},
	}

	//Инициализируем SDK
	srv, err := pkg.NewServicePool(cfg)
	if err != nil {
		log.Println(err)
	}
	// печатаем последние цены по всем бумагам:
	lp, err := srv.GetLastPricesForAll()
	if err != nil {
		log.Fatalf("Невозможно получить последнюю цену по всем бумагам, ошибка - %s", err)
	}
	fmt.Printf("%v\n", lp)
	//Получаем список акций, доступных для торгов через API
	shares, err := srv.GetSharesBase()
	if err != nil {
		log.Fatalf("Невозможно получить список акций, ошибка - %s", err)
	}
	for _, v := range shares {
		fmt.Printf("%v\n", v)
	}

	start, _ := time.Parse("2006-01-02", "2022-01-01")
	to := time.Now()

	//Получаем список фондов, доступных для торгов через API
	etfs, err := srv.GetETFsBase()
	if err != nil {
		log.Fatalf("Невозможно получить список фондов, ошибка - %s", err)
	}

	for i := range etfs {
		fmt.Printf("%v\n", etfs[i])
	}

	//Получаем список облигаций, доступных для торгов через API
	bonds, err := srv.GetBondsBase()
	if err != nil {
		log.Fatalf("Невозможно получить список облигаций, ошибка - %s", err)
	}

	for i := range bonds {
		fmt.Printf("%v\n", bonds[i])
	}

	//Получаем список фьючерсов, доступных для торгов через API
	futures, err := srv.GetFuturesBase()
	if err != nil {
		log.Fatalf("Невозможно получить список фьючерсов, ошибка - %s", err)
	}

	for i := range futures {
		fmt.Printf("%v\n", futures[i])
	}

	//Получаем список счетов
	accounts, err := srv.GetAccounts()
	if err != nil {
		log.Fatalf("Невозможно получить список счетов, ошибка - %s", err)
	}
	for _, v := range accounts {
		fmt.Printf("%v\n", v)
	}

	//Получаем список исполненых операций
	start, _ = time.Parse("2006-01-02", "2021-01-01")
	to = time.Now()
	operations, err := srv.GetOperations(cfg.AccountID[0], timestamppb.New(start), timestamppb.New(to), pb.OperationState_OPERATION_STATE_EXECUTED, "")
	if err != nil {
		log.Fatalf("Невозможно получить список операций, ошибка - %s", err)
	}

	for i := range operations {
		fmt.Printf("%v\n", operations[i])
	}

	//Получаем портфолио
	_, err = srv.GetPortfolio(cfg.AccountID[0])
	if err != nil {
		log.Fatalf("Невозможно получить портфолио, ошибка - %s", err)
	}

	//Получаем позиции
	_, err = srv.GetPositions(cfg.AccountID[0])
	if err != nil {
		log.Fatalf("Невозможно получить список позиций, ошибка - %s", err)
	}

	//Получаем доступный остаток для вывода средств.
	withdrawLimit, err := srv.GetWithdrawLimits(cfg.AccountID[0])
	if err != nil {
		log.Fatalf("Невозможно получить список аккаунтов, ошибка - %s", err)
	}
	fmt.Println(withdrawLimit)
}
