package helper

import (
	"errors"
	"reflect"
)

// SplitArray 切片拆分
// num 指每个切片里最大容量
func SplitArray[T any](arr interface{}, num int) ([][]T, error) {
	slice, ok := createAnyTypeSlice(arr)
	if !ok {
		return nil, errors.New("conversion error")
	}
	var data [][]T
	max := len(slice)
	if max <= num {
		var val []T
		for _, v := range slice {
			if v1, ok := v.(T); ok {
				val = append(val, v1)
			}
		}
		data = append(data, val)
		return data, nil
	}

	var quantity int
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	var start, end, i int
	for i = 1; i <= quantity; i++ {
		end = i * num
		var (
			val  []T
			list []interface{}
		)
		if i != quantity {
			list = slice[start:end]
		} else {
			list = slice[start:]
		}
		for _, v := range list {
			if v1, ok := v.(T); ok {
				val = append(val, v1)
			}
		}
		data = append(data, val)
		start = i * num
	}
	return data, nil
}

// CreateAnyTypeSlice 转interface切片
func createAnyTypeSlice(slice interface{}) ([]interface{}, bool) {
	val, ok := isSlice(slice)
	if !ok {
		return nil, false
	}
	sliceLen := val.Len()
	out := make([]interface{}, sliceLen)
	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}
	return out, true
}

// 校验是否是切片
func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == reflect.Slice {
		ok = true
	}
	return
}
