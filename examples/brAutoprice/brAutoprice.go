package main

import (
	//"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/toorop/go-betarigs"
	"log"
	"os"
	"time"
)

const (
	NL = "\r\n"
)

// Betarigs api returns an average price of the 20 cheapest rigs for a given algo
// It not optimal for our usage
// We will get the cheapest price in a first time and see if it's better
func getMarketPrice(algo uint32, rigId uint32, btr *betarigs.Betarigs) (maketPrice float64, err error) {
	rigs, err := btr.GetRigs(algo, "available", 1)
	if err != nil {
		return
	}
	if len(rigs) == 0 {
		maketPrice = 0
		return
	}
	// Preventing to return our rig price.If it's the case rig price never goes up
	for _, rig := range rigs {
		if rig.Id != rigId {
			maketPrice = rig.Price.PerSpeedUnit.Value
			break
		}
	}
	return
}

func main() {
	app := cli.NewApp()
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [options] [arguments...]

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}
OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`
	app.Name = "brAutoprice"
	app.Usage = "brAutoprice is a tool that helps you to rent your rig at the best price."
	app.Version = "0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "apiKey", Value: "", Usage: "Your Betarig API key. (required)"},
		cli.IntFlag{Name: "rigId", Value: 0, Usage: "Your rig ID. (required)"},
		cli.Float64Flag{Name: "minPrice", Value: 0, Usage: "The min price per speed unit for your rig. brAutoprice will never set rental price below this limit. (required)"},
		cli.Float64Flag{Name: "priceDiff", Value: 0, Usage: "The diff between the market price and the price you want to apply. Ex: if you set 10, your price will be 10 percent higher than market price, if you set -10 it will be 10 percent lower."},
	}

	app.Action = func(c *cli.Context) {
		// check options
		flagOptionErrors := false
		apiKey := c.String("apiKey")
		if apiKey == "" {
			flagOptionErrors = true
			fmt.Println("ERROR : Option --apiKey is missing")
		}
		rigId := c.Int("rigId")
		if rigId == 0 {
			flagOptionErrors = true
			fmt.Println("ERROR : Option --rigId is missing")
		}
		minPrice := c.Float64("minPrice")
		if minPrice == 0 {
			flagOptionErrors = true
			fmt.Println("ERROR : Option --minPrice is missing")
		}
		if flagOptionErrors {
			fmt.Println(NL + "Usage:" + NL)
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
		priceDiff := c.Float64("priceDiff")

		// Go go go
		log.Println("Type ctrl+c to quit")
		btr := betarigs.New(apiKey)

		// start main loop
		var algo betarigs.Algorithm
		var rig betarigs.Rig
		var err error
		for {
			rig, err = btr.GetRig(uint32(rigId))
			if err != nil {
				log.Fatalln("Unable to fetch Rig Algoritm for rig ", fmt.Sprintf("%d", rigId), ". Check if you haven't make a mistake on the rig ID", " Error:", err)
			}
			log.Println(fmt.Sprintf("Current rig price: %f %s", rig.Price.PerSpeedUnit.Value, rig.Price.PerSpeedUnit.Unit))
			// Get current market price
			algo, err = btr.GetAlgorithm(rig.Algorithm.Id)
			if err != nil {
				log.Println(fmt.Sprintf(" ERROR: Unable to fetch Algoritm %d. %v", rig.Algorithm.Id, err))
				time.Sleep(30 * time.Second)
				continue
			}
			marketPrice, err := getMarketPrice(rig.Algorithm.Id, uint32(rigId), btr)
			if err != nil {
				log.Println("ERROR: Unable to get Market Price.", err)
				time.Sleep(30 * time.Second)
				continue
			}
			log.Println(fmt.Sprintf("Current market price: %f %s", marketPrice, algo.MarketPrice.Unit))

			// Change price ?
			// if market price == 0 or (me==me)
			if marketPrice != 0 && marketPrice != rig.Price.PerSpeedUnit.Value {
				newPrice := marketPrice + (priceDiff * marketPrice / 100)
				if newPrice > minPrice && newPrice != rig.Price.PerSpeedUnit.Value {
					success, err := btr.UpdateRigPricePerSpeedUnit(uint32(rigId), newPrice)
					if err != nil || !success {
						log.Println("ERROR: Unable to update rig Price.", err)
					} else {
						log.Println(fmt.Sprintf("Rig prince changed to: %f %s", newPrice, algo.MarketPrice.Unit))
					}
				} else if rig.Price.PerSpeedUnit.Value < minPrice {
					success, err := btr.UpdateRigPricePerSpeedUnit(uint32(rigId), minPrice)
					if err != nil || !success {
						log.Println("ERROR: Unable to update rig Price.", err)
					} else {
						log.Println(fmt.Sprintf("Rig prince changed to: %f %s", minPrice, algo.MarketPrice.Unit))
					}
				}
			}
			time.Sleep(30 * time.Second)
		}
	}
	app.Run(os.Args)
}
