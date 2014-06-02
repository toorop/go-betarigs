betarigs
===========

A Go package to consume Betarigs API.


###Suggestions for betarigs team

* Use same type for algoID (int). int in "GET /api/v1/algorithms.json" response and string in "GET /api/v1/algorithm/[rigID].json"  
* add POST rent/[rigID]/[duration] which return a JSON object {"btcAddress":string, "price": float}

## Usage
~~~ go
package main

import (
	"log"
	"github.com/toorop/go-betarigs"
)

const (
	API_KEY    = "YOUR_API_KEY"
)

func main() {
	// Betarigs client
	btr := betarigs.New(API_KEY)

	// Get algorithms	
	algos, err := btr.GetAlgorithms()
	log.Println(err, algos)
	

	// Get info about one algo
	algoInfo, err := btr.GetAlgorithm(5)
	log.Println(err, algoInfo)
	

	// Get rigs
	rigs, err := btr.GetRigs(5, "available", 1)
	log.Println(err, rigs)
	

	// Get rig info
	rig, err := btr.GetRig(4568)
	log.Println(err, rig)
	

	// UpdateRigPricePerSpeedUnit
	success, err := btr.UpdateRigPricePerSpeedUnit(4568, 0.0018)
	log.Println(err, success)
	

	// UpdateRigPricePerTotalByDay
	success, err := btr.UpdateRigPricePerTotalByDay(4568, 0.0012)
	log.Println(err, success)

}
~~~	

See ["Examples" folder for more... examples](https://github.com/Toorop/go-betarigs/blob/master/examples/basic/basic.go)

## Documentation
[![GoDoc](https://godoc.org/github.com/Toorop/go-betarigs?status.png)](https://godoc.org/github.com/Toorop/go-betarigs)


## Stay tuned
[Follow me on Twitter](https://twitter.com/poroot)

Donate
------

![Donation QR](http://api.qrserver.com/v1/create-qr-code/?size=200x200&data=bitcoin:1HgpsmxV52eAjDcoNpVGpYEhGfgN7mM1JB%3Flabel%3DToorop)

[1HgpsmxV52eAjDcoNpVGpYEhGfgN7mM1JB](http://tinyurl.com/mccsoez)