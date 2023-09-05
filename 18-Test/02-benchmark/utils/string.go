package utils

func ConcatStringWithoutBuilder(a, b string) string {
	var s string = ""
	for i := 0; i < 1000; i++ {
		s = s + "a"
	}
	return s
}
