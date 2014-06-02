package main

import (
	"github.com/Toorop/go-betarigs"
	"log"
)

const (
	API_KEY = ""
)

func main() {
	// Betarigs client
	btr := betarigs.New(API_KEY)

	// Get algorithms
	/*
		algos, err := btr.GetAlgorithms()
		log.Println(err, algos)
	*/

	// Get info about one algo
	/*
		algoInfo, err := btr.GetAlgorithm(5)
		log.Println(err, algoInfo)
	*/

	// Get rigs
	/*
		rigs, err := btr.GetRigs(5, "available", 1)
		log.Println(err, rigs)
	*/

	// Get rig info
	/*
		rig, err := btr.GetRig(4568)
		log.Println(err, rig)
	*/

	// UpdateRigPricePerSpeedUnit
	/*
		success, err := btr.UpdateRigPricePerSpeedUnit(4568, 0.0018)
		log.Println(err, success)
	*/

	// UpdateRigPricePerTotalByDay
	success, err := btr.UpdateRigPricePerTotalByDay(4568, 0.0012)
	log.Println(err, success)

}
