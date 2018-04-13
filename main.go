package main

import "fmt"

type nation struct {
}

type city struct {
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

	fmt.Printf("Nations: %d", queries)

	data := make([]int64, queries)

	for i := range data {
		_, err := fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, err
		}
	}

	fmt.Println(data)

	return data, nil
}

func GetNation() {

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
