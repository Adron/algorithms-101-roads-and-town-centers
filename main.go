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
	start := time.Now()

	data, _ := dataRead()
	calculatePlan(data)

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Printf("\nIt took %s to complete processing the cities.\n", elapsed)
	fmt.Println("- - - - - - - - - - - - - - - - - -\n")
}

func dataRead() ([]nation, error) {
	nationStates := ingestNations()
	return nationStates, nil
}

func ingestNations() []nation {
	var nationStates []nation
	reader := bufio.NewReader(os.Stdin)
	textLine, _, _ := reader.ReadLine()
	nations, _ := strconv.Atoi(string(textLine))

	for i := 1; i < nations+1; i++ {
		newNation := buildNation(getSpaceSeparatedValues(textLine, reader))
		ingestRoads(newNation, textLine, reader)
		nationStates = append(nationStates, newNation)
	}
	return nationStates
}

func ingestRoads(newNation nation, textLine []byte, reader *bufio.Reader) {
	for r := 1; r < newNation.roadCount+1; r++ {
		newRoad := buildRoad(getSpaceSeparatedValues(textLine, reader))
		newNation.roads = append(newNation.roads, newRoad)
	}
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

func calculatePlan(data []nation) (bestPlanPrice int) {

	fmt.Println("Calculating")

	fmt.Print(data)

	return 4
}

func fixRoads(cost int) {

}

func fixTownCenters(cost int) {

}
