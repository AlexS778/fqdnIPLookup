package utils

func FindDifferentValues(map1, map2 map[string][]string) map[string][]string {
	differentValues := make(map[string][]string)

	for key, values2 := range map2 {
		values1, exists := map1[key]

		if !exists {
			differentValues[key] = values2
		} else {
			diffValues := make([]string, 0)
			for _, val := range values2 {
				if !Contains(values1, val) {
					diffValues = append(diffValues, val)
				}
			}
			differentValues[key] = diffValues
		}
	}
	return differentValues
}

func Contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
