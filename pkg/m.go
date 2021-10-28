package main

import (
	"context"
	"fmt"
	"log"

	binance "github.com/joker8023/go-binance"
)

var (
	apiKey    = "R23o966opbfowfC78CHCQLgggYNoh0ggZyKctuIkdfjtDgE1SfFjmRNp45eijwTy"
	secretKey = "dpdixpstGVQFe6eBCIkLYz0mHWveSyZbaGixM3eZZmqUW9DzYWh8alzDqtAMNAiP"
)

func main() {
	log.Println("1", 1)
	client := binance.NewClient(apiKey, secretKey)

	res, err := client.NewGetMiningService().Algo("sha256").UserName("lowmemory").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
