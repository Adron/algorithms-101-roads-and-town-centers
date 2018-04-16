package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type nation struct {
	cities     int
	roadCount  int
	costCenter int
	costRoad   int
	roads      []road
}

type road struct {
	originCity      int
	destinationCity int
}

func main() {
	DataRead()
}

func DataRead() ([]nation, error) {

	reader := bufio.NewReader(os.Stdin)
	textLine, _, _ := reader.ReadLine()
	nations, _ := strconv.Atoi(string(textLine))
	var nationStates []nation

	start := time.Now()

	for i := 1; i < nations+1; i++ {
		fmt.Printf("..Processing Nation %d\n", i)

		nationLineValues := getSpaceSeparatedValues(textLine, reader)
		newNation := buildNation(nationLineValues)

		//fmt.Printf("Cities: %d\nRoads: %d\nTown Center Cost: %d\nRoad Cost: %d\n",
		//	newNation.cities,
		//	newNation.roadCount,
		//	newNation.costCenter,
		//	newNation.costRoad)

		for r := 1; r < newNation.roadCount+1; r++ {
			roadLineValues := getSpaceSeparatedValues(textLine, reader)
			newRoad := buildRoad(roadLineValues)

			newNation.roads = append(newNation.roads, newRoad)

			//fmt.Printf("Road %d: City Origin %d and Destination %d.\n", r, newRoad.originCity, newRoad.destinationCity)
		}

		fmt.Println(len(newNation.roads))

		nationStates = append(nationStates, newNation)
	}

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Printf("It took %s to complete processing the cities.\n", elapsed)
	fmt.Println("- - - - - - - - - - - - - - - - - -\n\n")

	return nationStates, nil
}

func buildRoad(roadLineValues []string) road {
	var newRoad road
	newRoad.originCity, _ = strconv.Atoi(roadLineValues[0])
	newRoad.destinationCity, _ = strconv.Atoi(roadLineValues[1])
	return newRoad
}

func buildNation(nationLineValues []string) nation {
	var newNation nation
	newNation.cities, _ = strconv.Atoi(nationLineValues[0])
	newNation.roadCount, _ = strconv.Atoi(nationLineValues[1])
	newNation.costCenter, _ = strconv.Atoi(nationLineValues[2])
	newNation.costRoad, _ = strconv.Atoi(nationLineValues[3])
	return newNation
}

func getSpaceSeparatedValues(textLine []byte, reader *bufio.Reader) []string {
	textLine, _, _ = reader.ReadLine()
	values := strings.Split(string(textLine), " ")
	return values
}
