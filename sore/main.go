package main

import "fmt"

func alaysify(data1 string) string {
	data := []rune(data1)
	for i, v := range data {
		switch v {
		case 'a':
			data[i] = '4'
		case 'e':
			data[i] = '3'
		case 'i':
			data[i] = '!'
		case 'l':
			data[i] = '1'
		case 'n':
			data[i] = 'N'
		case 's':
			data[i] = '5'
		case 'x':
			data[i] = '*'
		default:
		}
	}
	return string(data)
}

func alayGen2(data ...string) []string {
	output := []string{}
	for _, v := range data {
		output = append(output, alaysify(v))
	}

	return output
}

func alayGen(data ...string) string {
	var output string
	for _, v := range data {
		output += alaysify(v) + " "
	}
	return output
}

func fibonacci() {
	var n int
	fmt.Print("Masukkan nomor deret: ")
	fmt.Scan(&n)

	var a1, a2 = 0, 1
	if n < 0 {
		fmt.Println("invalid")
	} else if n == 1 {
		fmt.Println(a1)
	} else if n == 2 {
		fmt.Println(a2)
	} else {
		var a3 int
		for i := 2; i < n; i++ {
			a3 = a2 + a1
			a1 = a2
			a2 = a3
		}
		fmt.Println(a3)
	}
}

func main() {
	fmt.Println(alayGen2("hello", "world", "zzz"))
	fmt.Println(alayGen("aku", "sayang", "kamu"))
	fibonacci()
}
