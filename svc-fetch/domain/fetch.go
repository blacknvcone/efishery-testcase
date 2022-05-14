package domain

import "context"

type (
	SteinDataRes struct {
		UUID        string `json:"uuid"`
		Commodity   string `json:"komoditas"`
		Province    string `json:"area_provinsi"`
		City        string `json:"area_kota"`
		Size        string `json:"size"`
		Price       string `json:"price"`
		TimeParsing string `json:"tgl_parsed"`
		Timestamp   string `json:"timestamp"`
	}

	FetchRes struct {
		UUID        string  `json:"uuid"`
		Commodity   string  `json:"commodity"`
		Province    string  `json:"province"`
		City        string  `json:"city"`
		Size        string  `json:"size"`
		Price       string  `json:"price"`
		PriceUSD    float64 `json:"price_usd"`
		TimeParsing string  `json:"time_parsing"`
		Timestamp   string  `json:"timestamp"`
	}

	RestCurrencyRes struct {
		Data CurrencyUSD `json:"data"`
	}

	CurrencyUSD struct {
		USD Currency `json:"USD"`
	}

	Currency struct {
		Code  string  `json:"code"`
		Value float64 `json:"value"`
	}

	AggregationRes struct {
		Year     string          `json:"year"`
		Month    string          `json:"month"`
		Week     string          `json:"week"`
		Province string          `json:"province"`
		Data     int             `json:"total_data"`
		Size     AggregationSize `json:"size"`
		Price    AggregationSize `json:"price"`
	}

	AggregationSize struct {
		Maximal float64 `json:"maximal"`
		Minimal float64 `json:"minimal"`
		Median  float64 `json:"median"`
		Average float64 `json:"average"`
	}
)

type IFetchUseCase interface {
	FetchAndCustom(ctx context.Context) ([]FetchRes, error)
}

type IFetchRepository interface {
	GetSampleData() ([]SteinDataRes, error)
	ConvertIDRtoUSD() (RestCurrencyRes, error)
}
