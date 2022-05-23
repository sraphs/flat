// package flat providers flatten/unflatten nested map or struct(only flatten support  struct).
package flat

import (
	"fmt"
	"reflect"
	"strings"
)

func Flatten(src interface{}) map[string]interface{} {
	return DefaultOption.Flatten(src)
}

func (opt Option) Flatten(src interface{}) map[string]interface{} {
	if src == nil {
		return nil
	}

	dst := make(map[string]interface{})
	opt.flatten("", src, dst)
	return dst
}

func (opt Option) flatten(prefix string, src interface{}, dst map[string]interface{}) {
	base := ""

	if prefix != "" {
		base = prefix + opt.GetSeparator()
	}

	if src == nil {
		dst[opt.Case.to(prefix)] = nil
		return
	}

	original := reflect.ValueOf(src)
	kind := original.Kind()

	if kind == reflect.Ptr || kind == reflect.Interface {
		original = reflect.Indirect(original)
		kind = original.Kind()
	}

	t := original.Type()

	switch kind {
	case reflect.Map:
		if t.Key().Kind() != reflect.String {
			break
		}
		for _, key := range original.MapKeys() {
			v := original.MapIndex(key)
			if !v.CanInterface() {
				continue
			}
			n := base + key.String()
			opt.flatten(n, v.Interface(), dst)
		}
	case reflect.Struct:
		for i := 0; i < original.NumField(); i += 1 {
			v := original.Field(i)
			if !v.CanInterface() {
				continue
			}
			n := base + t.Field(i).Name
			opt.flatten(n, v.Interface(), dst)
		}
	case reflect.Slice, reflect.Array:
		switch src := src.(type) {
		case []string, []int, []int8, []int16, []int32, []int64, []float32, []float64, []interface{}:
			var values []string
			for i := 0; i < original.Len(); i += 1 {
				v := original.Index(i)
				if !v.CanInterface() {
					continue
				}
				values = append(values, fmt.Sprintf("%v", v.Interface()))
			}
			dst[opt.Case.to(prefix)] = strings.Join(values, ",")
		case []byte:
			dst[opt.Case.to(prefix)] = string(src)
		default:
			for i := 0; i < original.Len(); i += 1 {
				v := original.Index(i)
				if !v.CanInterface() {
					continue
				}
				n := prefix + fmt.Sprintf("[%d]", i)
				opt.flatten(n, v.Interface(), dst)
			}
		}
	default:
		dst[opt.Case.to(prefix)] = src
	}
}
