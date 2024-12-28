package utils

type reduceFunc[TSlice any, TResult any] func(total TResult, currentValue TSlice) TResult
type mapFunc[TSlice any, TResult any] func(currentValue TSlice) TResult

func SliceReduce[TSlice any, TResult any](s []TSlice, initialValue TResult, cb reduceFunc[TSlice, TResult]) TResult {
	total := initialValue
	for _, currentValue := range s {
		total = cb(total, currentValue)
	}

	return total
}

func SliceMap[TSlice any, TResult any](s []TSlice, cb mapFunc[TSlice, TResult]) []TResult {
	var result []TResult
	for _, currentValue := range s {
		elem := cb(currentValue)
		result = append(result, elem)
	}

	return result
}

func SliceFilter[TSlice any](s []TSlice, cb func(currentValue TSlice) bool) []TSlice {
	var result []TSlice
	for _, currentValue := range s {
		if cb(currentValue) {
			result = append(result, currentValue)
		}
	}

	return result
}

func SliceFind[TSlice any](s []TSlice, cb func(currentValue TSlice) bool) *TSlice {
	for _, currentValue := range s {
		if cb(currentValue) {
			return &currentValue
		}
	}

	return nil
}
