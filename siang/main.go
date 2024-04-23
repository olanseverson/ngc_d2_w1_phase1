package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"sort"
)

// NGC 3 WEEK 1 PHASE 1
func soal1() {
	type person struct {
		name string
		age  int
		job  string
	}

	var people []person = []person{
		{name: "Hank", age: 50, job: "Polisi"},
		{name: "Heisenberg", age: 52, job: "Ilmuwan"},
		{name: "Skler", age: 48, job: "Akuntan"},
	}
	for _, v := range people {
		fmt.Printf("Hai, perkenalkan nama saya %s, umur saya %d, dan saya bekerja sebagai %s\n", v.name, v.age, v.job)
	}

}

func soal2() {
	aritmatika := func(data []float64) (rata float64, sum float64, median float64) {
		sort.Float64s(data)
		var x float64 = 0
		for _, v := range data {
			x = x + v
		}
		sum = x
		rata = x / float64(len(data))
		var l int = len(data)
		if l%2 == 0 {
			median = (data[l/2-1] + data[l/2]) / 2
		} else {
			median = data[(l+1)/2-1]
		}
		return
	}

	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	fmt.Println(aritmatika(slice1))
	fmt.Println(aritmatika(slice2))
}

// palindrome
func logic1() bool {
	var input string
	fmt.Print("Masukkan data: ")
	fmt.Scan(&input)

	var l int = len(input)
	var isPalindrome bool = true
	for i := 0; i < l/2; i++ {
		if input[i] != input[l-i-1] {
			isPalindrome = false
		}
		if !isPalindrome {
			break
		}
	}
	return isPalindrome
}

func logic2() bool {
	var data string
	fmt.Print("masukkan data: ")
	fmt.Scan(&data)

	var x, o int = 0, 0
	for _, v := range data {
		switch v {
		case 'x':
			x++
		case 'o':
			o++
		default:
			//
		}
	}

	return x == o
}

func logic3() []int {
	data := []int{}
	for i := 0; i < rand.IntN(20); i++ {
		data = append(data, rand.IntN(100))
	}

	fmt.Println(data)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j] < data[j+1] {
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}
	return data

}

func logic4() {
	x := rand.IntN(10) + 1
	fmt.Println(x)
	for i := 0; i < x; i++ {
		fmt.Println("*")
	}
}

func logic5() {
	// fmt.Println("logic5")
	x := rand.IntN(10) + 1
	fmt.Println(x)
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func logic6() {
	x := rand.IntN(10) + 1
	fmt.Println(x)
	for i := 0; i < x; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func logic7() {
	x := rand.IntN(10) + 1
	fmt.Println(x)
	for i := 0; i < x; i++ {
		for j := x; j > i; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func openfile() {
	file, err := os.Open("C:/Users/olans/Desktop/BOOTCAMP/Phase1/Week1/Day2/siang/hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close() // nutup file

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
func main() {

	// var temp_i string
	fmt.Println("1. Looping 1 \n2. Looping 2 \n3. Palindrome	\n4. XOXO   \n5. Bubble Sort \n6. Asterisk1 \n7. Asterisk2 \n8. Asterisk3 \n9. Asterisk4")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("masukkan nomor: ")
	scanner.Scan()

	// fmt.Scan(&temp_i)
	// check usia is integer or not
	// i, err := strconv.Atoi(temp_i)
	// if err != nil {
	// fmt.Fprintln(os.Stdout, []any{"Age must be a integer value"}...)
	// return
	// }
	i := scanner.Text()
	switch i {
	case "1":
		soal1()
	case "2":
		soal2()
	case "3":
		fmt.Println(logic1())
	case "4":
		fmt.Println(logic2())
	case "5":
		fmt.Println(logic3())
	case "6":
		logic4()
	case "7":
		logic5()
	case "8":
		logic6()
	case "9":
		logic7()
	default:
		fmt.Println("pilihan tidak ada")
	}

	// openfile()
}
