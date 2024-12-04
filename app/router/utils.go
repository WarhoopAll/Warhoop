package router

import "strconv"

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
