package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2019/utils"
)

type reaction struct {
	product         string
	productQuantity int
	reactants       map[string]int
}

func (r *Runner) Day14Part1(in string) string {
	reactions := convertInputToReactions(in)
	oreForFuel := calcOreForProduct(reactions, "FUEL", 1)
	return strconv.Itoa(oreForFuel)
}

func (r *Runner) Day14Part2(in string) string {
	reactions := convertInputToReactions(in)
	oreForOneFuel := calcOreForProduct(reactions, "FUEL", 1)

	tryFuelAmount := 1000000000000 / oreForOneFuel
	currentMaxFuel := 0
	for {
		oreForTestFuel := calcOreForProduct(reactions, "FUEL", tryFuelAmount)
		if r.verbose {
			fmt.Printf("%d fuel requires %d ore\n", tryFuelAmount, oreForTestFuel)
		}
		if oreForTestFuel > 1000000000000 {
			break
		}
		currentMaxFuel = tryFuelAmount
		incrementFuel := (1000000000000 - oreForTestFuel) / oreForOneFuel
		if incrementFuel == 0 {
			incrementFuel = 1
		}
		tryFuelAmount += incrementFuel
	}
	return strconv.Itoa(currentMaxFuel)
}

func convertInputToReactions(in string) map[string]reaction {
	//5 QNDT, 13 WDGM, 13 NTHXM, 10 NBGZ, 14 GTGRP, 14 KFWM, 3 HDWSV, 5 LSWQ => 1 FUEL
	matches := utils.AllStringsFromRegex(in, `([0-9]+) ([A-Z]+)`)

	reactions := map[string]reaction{}

	for _, lineMatch := range matches {
		productDetails := lineMatch[len(lineMatch)-1]
		product := productDetails[2]
		productQuantity := utils.MustAtoi(productDetails[1])

		inputReactants := map[string]int{}
		for _, r := range lineMatch[0 : len(lineMatch)-1] {
			inputReactants[r[2]] = utils.MustAtoi(r[1])
		}
		reactions[product] = reaction{
			product:         product,
			productQuantity: productQuantity,
			reactants:       inputReactants,
		}
	}
	return reactions
}

func calcOreForProduct(reactions map[string]reaction, product string, quantity int) int {

	spares := map[string]int{}
	oreCount := 0

	currentReactants := map[string]int{
		product: quantity,
	}
	for len(currentReactants) > 0 {
		newReactants := map[string]int{}
		for chemical, amount := range currentReactants {

			if chemical == "ORE" {
				oreCount += amount
				continue
			}

			curSpare := spares[chemical]
			if curSpare > 0 {
				if curSpare >= amount {
					spares[chemical] -= amount
					continue
				} else {
					amount -= curSpare
					spares[chemical] = 0
				}
			}

			reqReaction := reactions[chemical]
			multiplier, remainder := amount/reqReaction.productQuantity, amount%reqReaction.productQuantity
			if remainder != 0 {
				multiplier++
			}
			for thisChemical, thisAmount := range reqReaction.reactants {
				newReactants[thisChemical] += thisAmount * multiplier
			}

			producedAmount := reqReaction.productQuantity * multiplier
			if producedAmount > amount {
				spares[chemical] += producedAmount - amount
			}
		}
		currentReactants = newReactants
	}
	return oreCount
}

func hashSpares(in map[string]int) string {
	spareChemicals := []string{}
	for chem := range in {
		spareChemicals = append(spareChemicals, chem)
	}
	sort.Strings(spareChemicals)

	hash := strings.Builder{}
	for _, chem := range spareChemicals {
		fmt.Fprintf(&hash, "%s:%d", chem, in[chem])
	}
	return hash.String()
}
