package binance

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// GetAccountService get account info
type GetMiningService struct {
	c          *Client
	algo       string
	userName   string
	coin       *string
	startDate  *string
	endDate    *string
	pageIndex  *string
	pageSize   *string
	recvWindow *string
}

func (s *GetMiningService) Algo(algo string) *GetMiningService {
	s.algo = algo
	return s
}

func (s *GetMiningService) UserName(userName string) *GetMiningService {
	s.userName = userName
	return s
}

func (s *GetMiningService) Coin(coin string) *GetMiningService {
	s.coin = &coin
	return s
}

func (s *GetMiningService) StartDate(startDate string) *GetMiningService {
	s.startDate = &startDate
	return s
}

func (s *GetMiningService) EndDate(endDate string) *GetMiningService {
	s.endDate = &endDate
	return s
}

func (s *GetMiningService) PageIndex(pageIndex string) *GetMiningService {
	s.pageIndex = &pageIndex
	return s
}

func (s *GetMiningService) PageSize(pageSize string) *GetMiningService {
	s.pageSize = &pageSize
	return s
}

func (s *GetMiningService) RecvWindow(recvWindow string) *GetMiningService {
	s.recvWindow = &recvWindow
	return s
}

type MiningResponse struct {
	AccountProfits []struct {
		Time           int64       `json:"time"`
		Type           int         `json:"type"`
		HashTransfer   interface{} `json:"hashTransfer"`
		TransferAmount interface{} `json:"transferAmount"`
		DayHashRate    int64       `json:"dayHashRate"`
		ProfitAmount   float64     `json:"profitAmount"`
		CoinName       string      `json:"coinName"`
		Status         int         `json:"status"`
	} `json:"accountProfits"`
	TotalNum int `json:"totalNum"`
	PageSize int `json:"pageSize"`
}

// Do send request
func (s *GetMiningService) Do(ctx context.Context, opts ...RequestOption) (res *MiningResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/mining/payment/list",
		secType:  secTypeSigned,
	}
	if s.algo != "" {
		r.setParam("algo", s.algo)
	}
	if s.userName != "" {
		r.setParam("userName", s.userName)
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.startDate != nil {
		r.setParam("startDate", *s.startDate)
	}

	if s.endDate != nil {
		r.setParam("endDate", *s.endDate)
	}
	if s.pageIndex != nil {
		r.setParam("pageIndex", *s.pageIndex)
	}

	if s.pageSize != nil {
		r.setParam("pageSize", *s.pageSize)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	log.Println("data", string(data))
	res = new(MiningResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
