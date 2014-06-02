package betarigs

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	VERSION            = 0.1
	HTTPCLIENT_TIMEOUT = 30
	API_BASE           = "https://www.betarigs.com/api"
	API_VERSION        = "v1"
)

// Betarigs represents a betarigs client
type Betarigs struct {
	client *client
}

// New return a new instance of Betarigs
func New(apikey ...string) *Betarigs {
	var c *client
	if len(apikey) == 1 {
		c = NewClient(apikey[0])
	} else {
		c = NewClient("")
	}
	return &Betarigs{c}
}

// GetAlgorithms return available algorithms
func (b *Betarigs) GetAlgorithms() (algos []AlgoId, err error) {
	r, err := b.client.do("GET", "algorithms.json", "")
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &algos)
	return
}

// GetAlgorithm returns information about an algorithm
func (b *Betarigs) GetAlgorithm(algoID uint32) (algo Algorithm, err error) {
	ressource := fmt.Sprintf("algorithm/%d.json", algoID)
	r, err := b.client.do("GET", ressource, "")
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &algo)
	return
}

// GetRigs return a list of rigs
func (b *Betarigs) GetRigs(algoID uint32, status string, page uint32) (rigs []Rig, err error) {
	status = strings.ToLower(status)
	if status != "all" && status != "available" {
		status = "all"
	}
	ressource := fmt.Sprintf("rigs.json?algorithm=%d&status=%s&page=%d", algoID, status, page)
	r, err := b.client.do("GET", ressource, "")
	if err != nil {
		return
	}
	type JsonResponse struct {
		Items []Rig `json:"items"`
	}
	var jr JsonResponse
	if err = json.Unmarshal(r, &jr); err != nil {
		return
	}
	rigs = jr.Items
	return
}

// GetRig return info about a rig
func (b *Betarigs) GetRig(rigID uint32) (rig Rig, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("rig/%d.json", rigID), "")
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &rig)
	return
}

// UpdateRigPricePerSpeedUnit update a rig BTC/Mh/Day price
func (b *Betarigs) UpdateRigPricePerSpeedUnit(rigID uint32, price float64) (success bool, err error) {
	success = false
	payload := fmt.Sprintf(`{"price":{"per_speed_unit":{"value":%f,"unit":"BTC/Mh/day"}}}`, price)
	r, err := b.client.do("PUT", fmt.Sprintf("rig/%d", rigID), payload)
	if err != nil {
		return
	}
	type JsonResponse struct {
		Result string `json:"result"`
	}
	var result JsonResponse
	if err = json.Unmarshal(r, &result); err != nil {
		return
	}
	if result.Result == "OK" {
		success = true
	}
	return
}

// UpdateRigPricePerTotalByDay update a rig BTC/Day price
func (b *Betarigs) UpdateRigPricePerTotalByDay(rigID uint32, price float64) (success bool, err error) {
	success = false
	payload := fmt.Sprintf(`{"price":{"total":{"value":%f,"unit":"BTC/day"}}}`, price)
	r, err := b.client.do("PUT", fmt.Sprintf("rig/%d", rigID), payload)
	if err != nil {
		return
	}
	type JsonResponse struct {
		Result string `json:"result"`
	}
	var result JsonResponse
	if err = json.Unmarshal(r, &result); err != nil {
		return
	}
	if result.Result == "OK" {
		success = true
	}
	return
}
