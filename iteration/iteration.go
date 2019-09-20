package iteration

//
func Repeat(c string, num int) string {
	var str string
	for i := 0; i < num; i++ {
		str += c
	}
	return str
}
