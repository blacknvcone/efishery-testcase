package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/blacknvcone/efishery-testcase/svc-fetch/domain"
)

type fetchUseCase struct {
	contextTimeout  time.Duration
	fetchRepository domain.IFetchRepository
}

func NewFetchUseCase(efisheryRepo domain.IFetchRepository, timeout time.Duration) domain.IFetchUseCase {
	return &fetchUseCase{
		fetchRepository: efisheryRepo,
		contextTimeout:  timeout,
	}
}

func (uc *fetchUseCase) FetchAndCustom(ctx context.Context) (res []domain.FetchRes, err error) {

	_, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	idrToUsdCurr, err := uc.fetchRepository.ConvertIDRtoUSD()
	if err != nil {
		return
	}

	eFisheryData, err := uc.fetchRepository.GetSampleData()
	if err != nil {
		return
	}

	for _, value := range eFisheryData {
		priceIDR, err := strconv.Atoi(value.Price)
		if err != nil {
			priceIDR = 0
		}

		res = append(res, domain.FetchRes{
			UUID:        value.UUID,
			Commodity:   value.Commodity,
			Province:    value.Province,
			City:        value.City,
			Size:        value.Size,
			Price:       value.Price,
			PriceUSD:    idrToUsdCurr.Data.USD.Value * float64(priceIDR),
			TimeParsing: value.TimeParsing,
			Timestamp:   value.Timestamp,
		})
	}

	return res, nil
}
