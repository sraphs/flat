package flat

import (
	"sort"
	"strings"
)

func Unflatten(src map[string]interface{}) (dst map[string]interface{}) {
	return DefaultOption.Unflatten(src)
}

// unflatten
func (opt Option) Unflatten(src map[string]interface{}) (dst map[string]interface{}) {
	if src == nil {
		return nil
	}

	dst = make(map[string]interface{})
	opt.unflatten("", src, dst)
	return dst
}

func (opt Option) unflatten(prefix string, src map[string]interface{}, dst map[string]interface{}) {
	if prefix != "" {
		prefix += opt.GetSeparator()
	}

	sortedKeys := sortKeys(src)

	for _, k := range sortedKeys {
		if !strings.HasPrefix(k, prefix) {
			continue
		}

		key := k[len(prefix):]
		idx := strings.Index(key, opt.GetSeparator())
		if idx != -1 {
			key = key[:idx]
		}
		if _, ok := dst[key]; ok {
			continue
		}
		if idx == -1 {
			dst[opt.Case.to(key)] = src[k]
			continue
		}

		// It contains a period, so it is a more complex structure
		m := make(map[string]interface{})
		opt.unflatten(k[:len(prefix)+len(key)], src, m)

		if strings.HasSuffix(key, "]") && key[0] != '[' {
			// It is an array
			indexLeft := strings.Index(key, "[")
			keyLeft := key[:indexLeft]

			if _, ok := dst[keyLeft]; !ok {
				dst[keyLeft] = []interface{}{m}
			} else {
				dst[keyLeft] = append(dst[keyLeft].([]interface{}), m)
			}
		} else {
			opt.unflatten(k[:len(prefix)+len(key)], src, m)
			dst[opt.Case.to(key)] = m
		}
	}
}

func sortKeys(m map[string]interface{}) []string {
	var sortedKeys []string
	for key := range m {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	return sortedKeys
}
