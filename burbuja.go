package main

func Burbuja(s []int64) {
	var temp int64
	for i := 1; i < len(s); i++ {
		for j := 0; j < (len(s) - i); j++ {
			if s[j] > s[j+1] {
				temp = s[j]
				s[j] = s[j+1]
				s[j+1] = temp
			}
		}
	}
}

func main() {

}
