package flatten

func Flatten(nested any) []any {
	res := []any{}

	switch v := nested.(type) {
	case []any:
		for _, item := range v {
			if item != nil {
				res = append(res, Flatten(item)...)
			}
		}
	default:
		if nested != nil {
			res = append(res, nested)
		}
	}
	return res
}
