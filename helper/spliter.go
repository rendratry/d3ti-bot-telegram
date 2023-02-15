package helper

import "strings"

type SplitResult struct {
	IdRegis string
	Uuid    string
}

func Spliter(msg string) SplitResult {
	parts := strings.Split(msg, ":")
	resultSplit := SplitResult{}
	if len(parts) > 1 {
		result := parts[1]
		parts := strings.Split(result, ".")
		if len(parts) > 1 {
			resultSplit.IdRegis = parts[0]
			resultSplit.Uuid = parts[1]
			return resultSplit
		}
	}
	return resultSplit
}
