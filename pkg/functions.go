package pkg

import "strconv"

func ConvertToUnit64(value string) (uint64, error) {
	return strconv.ParseUint(value, 10, 64)
}
