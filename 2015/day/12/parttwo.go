package main

import (
	"encoding/json"
	"log"
	"reflect"
)

// PartTwo sovles the second part of the problem.
func PartTwo(input Input) (solution Solution) {
	var blob interface{}
	err := json.Unmarshal(input, &blob)
	if err != nil {
		log.Fatal(err)
	}
	return Solution(sumNonRed(blob))
}

func sumNonRed(node interface{}) (sum int64) {
	v := reflect.ValueOf(node)
	switch v.Kind() {
	case reflect.Map:
		m := node.(map[string]interface{})
		for _, v := range m {
			if s, ok := v.(string); ok && s == "red" {
				return 0
			}
		}
		for _, v := range m {
			sum += sumNonRed(v)
		}
	case reflect.Slice:
		s := node.([]interface{})
		for _, v := range s {
			sum += sumNonRed(v)
		}
	case reflect.Float64:
		sum += int64(v.Float())
	case reflect.Int, reflect.String:
	default:
		log.Fatalf("%v", v.Type())
	}
	return sum
}
