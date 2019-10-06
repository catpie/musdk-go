package musdk

import (
	"os"
	"strconv"
)

func env(k string) string  {
	return os.Getenv(k)
}

func envInt(k string) int {
	v := env(k)
	i,_ := strconv.Atoi(v)
	return i
}