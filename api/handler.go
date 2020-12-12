package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/delivery-much/dm-go/logger"
	"github.com/delivery-much/dm-go/render"
	"github.com/g1stavo/taco-api/config"
	"github.com/g1stavo/taco-api/entity"
)

func getResponse(url string) (data []byte) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Errorf("Error in GET: ", err)
		return
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("Error in read: ", err)
	}

	return
}

func getTaco(w http.ResponseWriter, r *http.Request) {
	resp := getResponse(config.TacoURL)

	var t entity.Taco
	err := json.Unmarshal(resp, &t)
	if err != nil {
		render.RespondError(w, http.StatusInternalServerError, err)
	}

	var (
		seCh = make(chan string)
		coCh = make(chan string)
		miCh = make(chan string)
		blCh = make(chan string)
		shCh = make(chan string)
	)

	go getDescription(t.Seasoning.URL, seCh)
	go getDescription(t.Condiment.URL, coCh)
	go getDescription(t.Mixin.URL, miCh)
	go getDescription(t.BaseLayer.URL, blCh)
	go getDescription(t.Shell.URL, shCh)

	t.Seasoning.SetVegFields(<-seCh)
	t.Condiment.SetVegFields(<-coCh)
	t.Mixin.SetVegFields(<-miCh)
	t.BaseLayer.SetVegFields(<-blCh)
	t.Shell.SetVegFields(<-shCh)

	render.RespondJSON(w, http.StatusOK, t)
}

func getDescription(url string, ch chan string) {
	resp := string(getResponse(url))
	ch <- resp
}
