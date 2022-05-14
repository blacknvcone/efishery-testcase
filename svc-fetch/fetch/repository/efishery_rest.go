package repository

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/blacknvcone/efishery-testcase/svc-fetch/domain"
	"github.com/patrickmn/go-cache"
)

type fetchRepository struct {
	EfisheryUrl string
	CurrencyUrl string
	Cache       *cache.Cache
	CacheKey    string
}

func NewFetchRepository() domain.IFetchRepository {
	return &fetchRepository{
		EfisheryUrl: os.Getenv("EFISHERY_URL"),
		CurrencyUrl: os.Getenv("CURRENCY_URL"),
		Cache:       cache.New(5*time.Minute, 10*time.Minute),
		CacheKey:    os.Getenv("CURRENCY_CACHE_KEY"),
	}

}

func (fr *fetchRepository) GetSampleData() (res []domain.SteinDataRes, err error) {

	req, err := http.NewRequest(http.MethodGet, fr.EfisheryUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBody, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (fr *fetchRepository) ConvertIDRtoUSD() (res domain.RestCurrencyRes, err error) {
	// Get on cache
	dataCache, found := fr.Cache.Get(fr.CacheKey)
	if found {
		res.Data.USD.Value = dataCache.(float64)
		return res, nil
	}

	//Rest Currency and Set Cache
	res, err = fr.restCurrencyAPI()
	if err != nil {
		return res, err
	}

	fr.Cache.Set(fr.CacheKey, res.Data.USD.Value, cache.DefaultExpiration)
	return res, nil
}

func (fr *fetchRepository) restCurrencyAPI() (res domain.RestCurrencyRes, err error) {
	req, err := http.NewRequest(http.MethodGet, fr.CurrencyUrl, nil)
	if err != nil {
		return res, err
	}

	req.Header.Set("content-type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBody, &res)
	if err != nil {
		return res, err
	}

	return res, err
}
