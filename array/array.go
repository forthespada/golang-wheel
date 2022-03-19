package array

import "reflect"

// 判断某元素是否在Array、Slice中，存在返回 index，否则返回 -1
func Contain(obj interface{}, target interface{}) int {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return i
			}
		}
	}
	return -1
}
