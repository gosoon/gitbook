package main

import "fmt"

func romanToInt(s string) int {
	fmt.Println(s)
	var m = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var sp = map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	if value, found := sp[s]; found {
		return value
	}

	var number int
	for i := 0; i < len(s); {
		if i+1 < len(s) {
			subS := s[i : i+2]
			fmt.Println(subS)
			if value, found := sp[subS]; found {
				number += value
				i += 2
			} else {
				number += m[s[i]]
				i++
			}
		} else {
			number += m[s[i]]
			i++
		}
		fmt.Println(number)
	}
	return number
}
