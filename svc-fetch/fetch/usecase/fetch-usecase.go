package usecase

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/blacknvcone/efishery-testcase/svc-fetch/domain"
	"github.com/blacknvcone/efishery-testcase/svc-fetch/fetch/utils"
	"github.com/montanaflynn/stats"
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

func (uc *fetchUseCase) SumAggregate(ctx context.Context) (res []domain.AggregationRes, err error) {

	_, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	//Get Efishery Data Sample
	eFisheryData, err := uc.fetchRepository.GetSampleData()
	if err != nil {
		return
	}

	//Aggregating result data
	var group = map[string][]domain.SteinDataRes{}

	for _, value := range eFisheryData {
		dateConvert := utils.ParseDateString(value.TimeParsing)
		weekGroup := fmt.Sprintf("%d-%d-%d", dateConvert.Year(), int(dateConvert.Month()), int(dateConvert.Weekday()))
		keyGroup := fmt.Sprintf("%s#%s", weekGroup, value.Province)
		group[keyGroup] = append(group[keyGroup], value)
	}

	for i := range group {
		key := strings.Split(i, "#")
		date := strings.Split(key[0], "-")

		var listSize []float64
		var listPrice []float64
		for _, data := range group[i] {
			size, _ := strconv.ParseFloat(data.Size, 64)
			listSize = append(listSize, size)
			price, _ := strconv.ParseFloat(data.Price, 64)
			listPrice = append(listPrice, price)
		}

		medianSize, _ := stats.Median(listSize)
		maxSize, _ := stats.Max(listSize)
		minSize, _ := stats.Min(listSize)
		sumSize := 0.0
		for i := range listSize {
			sumSize += listSize[i]
		}
		averageSize := sumSize / float64(len(listSize))

		medianPrice, _ := stats.Median(listPrice)
		maxPrice, _ := stats.Max(listPrice)
		minPrice, _ := stats.Min(listPrice)
		sumPrice := 0.0
		for i := range listPrice {
			sumPrice += listPrice[i]
		}
		averagePrice := sumPrice / float64(len(listPrice))

		if date[0] == "1" && date[1] == "1" && date[2] == "1" {
			res = append(res, domain.AggregationRes{
				Year:     "",
				Month:    "",
				Week:     "",
				Province: key[1],
				Data:     len(group[i]),
				Size: domain.AggregationSize{
					Maximal: maxSize,
					Minimal: minSize,
					Median:  medianSize,
					Average: averageSize,
				},
				Price: domain.AggregationSize{
					Maximal: maxPrice,
					Minimal: minPrice,
					Median:  medianPrice,
					Average: averagePrice,
				},
			})
		} else {
			res = append(res, domain.AggregationRes{
				Year:     date[0],
				Month:    date[1],
				Week:     date[2],
				Province: key[1],
				Data:     len(group[i]),
				Size: domain.AggregationSize{
					Maximal: maxSize,
					Minimal: minSize,
					Median:  medianSize,
					Average: averageSize,
				},
				Price: domain.AggregationSize{
					Maximal: maxPrice,
					Minimal: minPrice,
					Median:  medianPrice,
					Average: averagePrice,
				},
			})
		}
	}

	return res, nil
}
