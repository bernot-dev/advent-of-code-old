package main

import (
	"encoding/json"
	"log"
	"reflect"
)

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	var blob interface{}
	err := json.Unmarshal(input, &blob)
	if err != nil {
		log.Fatal(err)
	}
	return Solution(sumAll(blob))
}

func sumAll(node interface{}) (sum int64) {
	v := reflect.ValueOf(node)
	switch v.Kind() {
	case reflect.Map:
		m := node.(map[string]interface{})
		for _, v := range m {
			sum += sumAll(v)
		}
	case reflect.Slice:
		s := node.([]interface{})
		for _, v := range s {
			sum += sumAll(v)
		}
	case reflect.Float64:
		sum += int64(v.Float())
	case reflect.Int, reflect.String:
	default:
		log.Fatalf("%v", v.Type())
	}
	return sum
}
