package assertion

import (
	"fmt"
	"reflect"
)

const panicFormatWithType = "expected %v (type %v) bot got %v (type %v)"
const panicFormat = "expected %v bot got %v"

func Equals(expected, actual interface{}) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		panic(fmt.Sprintf(panicFormatWithType, expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual)))
	}
	if !isEqual(expected, actual) {
		panic(fmt.Sprintf(panicFormat, expected, actual))
	}
}

func isEqual(expected, actual interface{}) bool {
	switch expected.(type) {
	case int:
		return expected.(int) == actual.(int)
	case int8:
		return expected.(int8) == actual.(int8)
	case int16:
		return expected.(int16) == actual.(int16)
	case int32:
		return expected.(int32) == actual.(int32)
	case int64:
		return expected.(int64) == actual.(int64)
	case string:
		return expected.(string) == actual.(string)
	case uint8:
		return expected.(uint8) == actual.(uint8)
	case uint16:
		return expected.(uint16) == actual.(uint16)
	case uint32:
		return expected.(uint32) == actual.(uint32)
	case uint64:
		return expected.(uint64) == actual.(uint64)
	case float32:
		return expected.(float32) == actual.(float32)
	case float64:
		return expected.(float64) == actual.(float64)
	case uintptr:
		return expected.(uintptr) == actual.(uintptr)
	case []int:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]int)
		actualSlice := actual.([]int)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []int8:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]int8)
		actualSlice := actual.([]int8)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []int16:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]int16)
		actualSlice := actual.([]int16)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []int32:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]int32)
		actualSlice := actual.([]int32)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []int64:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]int64)
		actualSlice := actual.([]int64)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []string:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]string)
		actualSlice := actual.([]string)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []uint:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]uint)
		actualSlice := actual.([]uint)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []uint8:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]uint8)
		actualSlice := actual.([]uint8)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []uint16:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]uint16)
		actualSlice := actual.([]uint16)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []uint32:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]uint32)
		actualSlice := actual.([]uint32)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []uint64:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]uint64)
		actualSlice := actual.([]uint64)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []float32:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]float32)
		actualSlice := actual.([]float32)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []float64:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]float64)
		actualSlice := actual.([]float64)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if expectedSlice[i] != actualSlice[i] {
				return false
			}
		}
		return true
	case []interface{}:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.([]interface{})
		actualSlice := actual.([]interface{})
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if !isEqual(expectedSlice[i], actualSlice[i]) {
				return false
			}
		}
		return true
	case map[int]int:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]int)
		actualSlice := actual.(map[int]int)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[int]int8:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]int8)
		actualSlice := actual.(map[int]int8)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[int]int16:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]int16)
		actualSlice := actual.(map[int]int16)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[int]int32:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]int32)
		actualSlice := actual.(map[int]int32)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[int]int64:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]int64)
		actualSlice := actual.(map[int]int64)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[int]string:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]string)
		actualSlice := actual.(map[int]string)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[int]uint:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]uint)
		actualSlice := actual.(map[int]uint)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[int]uint8:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]uint8)
		actualSlice := actual.(map[int]uint8)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[int]uint16:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]uint16)
		actualSlice := actual.(map[int]uint16)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[int]uint32:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]uint32)
		actualSlice := actual.(map[int]uint32)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[int]uint64:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]uint64)
		actualSlice := actual.(map[int]uint64)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[int]float32:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]float32)
		actualSlice := actual.(map[int]float32)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[int]float64:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[int]float64)
		actualSlice := actual.(map[int]float64)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]int:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]int)
		actualSlice := actual.(map[string]int)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]int8:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]int8)
		actualSlice := actual.(map[string]int8)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]int16:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]int16)
		actualSlice := actual.(map[string]int16)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]int32:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]int32)
		actualSlice := actual.(map[string]int32)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]int64:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]int64)
		actualSlice := actual.(map[string]int64)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]string:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]string)
		actualSlice := actual.(map[string]string)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]uint:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]uint)
		actualSlice := actual.(map[string]uint)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]uint8:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]uint8)
		actualSlice := actual.(map[string]uint8)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]uint16:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]uint16)
		actualSlice := actual.(map[string]uint16)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]uint32:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]uint32)
		actualSlice := actual.(map[string]uint32)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]uint64:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]uint64)
		actualSlice := actual.(map[string]uint64)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]float32:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]float32)
		actualSlice := actual.(map[string]float32)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case map[string]float64:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[string]float64)
		actualSlice := actual.(map[string]float64)
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if expectedSlice[k] != actualSlice[k] {
				return false
			}
		}
		return true
	case []map[interface{}]interface{}:
		if expected == nil && actual == nil {
			return true
		}
		if expected == nil || actual == nil {
			return false
		}
		expectedSlice := expected.(map[interface{}]interface{})
		actualSlice := actual.(map[interface{}]interface{})
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for k := range expectedSlice {
			if !isEqual(expectedSlice[k], actualSlice[k]) {
				return false
			}
		}
		return true
	case interface{}:
		expectedVal := reflect.ValueOf(expected)
		actualVal := reflect.ValueOf(actual)
		switch reflect.TypeOf(expected).Kind() {
		case reflect.Array, reflect.Slice:
			if expectedVal.Len() != actualVal.Len() {
				return false
			}
			for i := 0; i < expectedVal.Len(); i++ {
				if !isEqual(expectedVal.Index(i).Interface(), actualVal.Index(i).Interface()) {
					return false
				}
			}
			return true
		case reflect.Map:
			if expectedVal.Len() != actualVal.Len() {
				return false
			}
			for _, key := range expectedVal.MapKeys() {
				if !isEqual(expectedVal.MapIndex(key).Interface(), actualVal.MapIndex(key).Interface()) {
					return false
				}
			}
			return true
		default:
			return expected == actual
		}
	default:
		return expected == actual
	}
}
