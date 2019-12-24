package packages

import "fmt"

func ParseFloat64(val interface{}) (float64, error) {
	switch val.(type) {
	case nil:
		return 1, nil
	default:
		return -1, fmt.Errorf("unable to casting numver %v (type %T)", val, val)
	}
}
