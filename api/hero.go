package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/game/core"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"io/ioutil"
	"net/http"
)

func FetchHero(web *http.Client, host string, k hsk.Key) (core.Hero, error) {
	url := fmt.Sprintf("%s/heroes/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Hero{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return core.Hero{}, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := core.Hero{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchAllHeroes(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/heroes/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := records.NewResultPage(core.Hero{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
