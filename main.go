package main

import "fmt"

type nation struct {
	// cities n, the number of roads, m, and the cost to build a town center ctown center, and the cost to repair a road croad. Each integer is space seperated.
	cities int
	roads int
	townCenterCost int
	roadRepairCost int
}

type road struct {
	originating int
	destination int
}

func main() {
	DataRead()
}

func DataRead() ([]int64, error) {
	var queries int64

	_, err := fmt.Scanf("%d", &queries)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Nations: %d\n\n", queries)

	return GetQueries(queries)
}

func GetQueries(index int64) ([]int64, error) {
	queries := make([]int64, index)

	for i := range queries {
		_, err := fmt.Scanf("%d", &queries[i])
		if err != nil {
			return nil, err
		}
	}

	return queries, nil
}

func GetQuery(){

}

//
//func SumTotal(list []int64) int64 {
//	var totalSum int64 = 0
//
//	for _, v := range list {
//		totalSum += v
//	}
//
//	return totalSum
//}

//func DataRead() ([]int64, error) {
//	var length int64
//
//	_, err := fmt.Scanf("%d", &length)
//	if err != nil {
//		return nil, err
//	}
//
//
//
//
//
//	data := make([]int64, length)
//
//	for i := range data {
//		_, err := fmt.Scanf("%d", &data[i])
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	return data, nil
//}
