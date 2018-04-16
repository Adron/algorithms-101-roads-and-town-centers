package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type nation struct {
	cities     int
	roads      int
	costCenter int
	costRoad   int
}

type road struct {
	originCity      int
	destinationCity int
}

func main() {
	DataRead()
}

func DataRead() (int, error) {

	reader := bufio.NewReader(os.Stdin)

	textLine, _, _ := reader.ReadLine()
	nations, _ := strconv.Atoi(string(textLine))

	fmt.Printf("Nations: %d\n", nations)

	for i := 1; i < nations+1; i++ {
		fmt.Printf("..Processing Nation %d\n", i)

		nationLineValues := getSpaceSeparatedValues(textLine, reader)
		fmt.Println(nationLineValues)

		var newNation nation
		newNation.roads, _ = strconv.Atoi(nationLineValues[0])
		newNation.cities, _ = strconv.Atoi(nationLineValues[1])
		newNation.costCenter, _ = strconv.Atoi(nationLineValues[2])
		newNation.costRoad, _ = strconv.Atoi(nationLineValues[3])

		fmt.Printf("Cities: %d\nRoads: %d\nTown Center Cost: %d\nRoad Cost: %d\n",
			newNation.cities,
			newNation.roads,
			newNation.costCenter,
			newNation.costRoad)

		for r := 1; r < newNation.roads+1; r++ {
			roadLineValues := getSpaceSeparatedValues(textLine, reader)
			var newRoad road
			newRoad.originCity, _ = strconv.Atoi(roadLineValues[0])
			newRoad.destinationCity, _ = strconv.Atoi(roadLineValues[1])

			fmt.Printf("Road %d: City Origin %d and Destination %d.\n", r, newRoad.originCity, newRoad.destinationCity)
		}
		fmt.Println("- - - - - - - - -")
	}

	//for reader.ReadLine() {
	//
	//	fmt.Printf("Nations: %d", nations)
	//
	//	for i := 1; i < nations; i++ {
	//
	//		textLine = scanner.Text()
	//
	//		values := strings.Split(textLine, " ")
	//
	//		fmt.Printf("Nation: %s", values[0])
	//
	//
	//
	//
	//
	//	}
	//	fmt.Println("Completed processing.")
	//}

	return 42, nil
}

func getSpaceSeparatedValues(textLine []byte, reader *bufio.Reader) []string {
	textLine, _, _ = reader.ReadLine()
	values := strings.Split(string(textLine), " ")
	return values
}
