package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	input := getInput()
	temp := strings.Split(input, ",")

	invalid := make([]int, 0)

	for _, x := range temp {
		split := strings.Split(x, "-")

		lower, _ := strconv.Atoi(split[0])
		upper, _ := strconv.Atoi(split[1])

		for i := lower; i <= upper; i++ {
			str := fmt.Sprint(i)
			len := len(str)
			if len%2 != 0 {
				continue
			}

			half := len / 2

			valid := true

			for j := range half {
				if str[j] != str[j+half] {
					valid = false
					break
				}
			}

			if valid {
				invalid = append(invalid, i)
			}
		}
	}

	total := 0

	for _, num := range invalid {
		total += num
	}

	println(total)
}

func part2() {
	input := getInput()
	temp := strings.Split(input, ",")

	invalid := make([]int, 0)

	for _, x := range temp {
		split := strings.Split(x, "-")
		lower, _ := strconv.Atoi(split[0])
		upper, _ := strconv.Atoi(split[1])

		for i := lower; i <= upper; i++ {
			str := fmt.Sprint(i)
			length := len(str)

			for subLen := 1; subLen < length; subLen++ {
				if length%subLen == 0 {
					div := length / subLen
					if div < 2 { //must repeat substring twice or more
						continue
					}

					subStr := str[:subLen]
					repeated := strings.Repeat(subStr, div)

					if repeated == str {
						invalid = append(invalid, i)
						break
					}
				}
			}
		}
	}

	total := 0
	for _, num := range invalid {
		total += num
	}

	println(total)
}

func part2_old() {
	input := getInput()
	temp := strings.Split(input, ",")

	invalid := make([]int, 0)

	for _, x := range temp {
		split := strings.Split(x, "-")

		lower, _ := strconv.Atoi(split[0])
		upper, _ := strconv.Atoi(split[1])

		for i := lower; i <= upper; i++ {
			str := fmt.Sprint(i)
			length := len(str)

			valid := true

			if length%2 == 0 {
				half := length / 2

				for j := range half {
					if str[j] != str[j+half] {
						valid = false
						break
					}
				}

				if valid {
					invalid = append(invalid, i)
				}
			}

			if !valid {
				s := ""
				strFound := false
				for _, jc := range str {
					if !strFound {
						s += string(jc)
						count := 0

						for k := range s {
							offset := len(s) + k
							if offset >= len(str) {
								break
							}
							kStr := string(s[k])
							offsetStr := string(str[offset])
							if kStr != offsetStr {
								break
							}

							count += 1
						}

						if count == len(s) {
							strFound = true
						}
					} else {
						break
					}
				}

				sLen := len(s)

				invalidStr := false

				counter := 0

				for l := 0; l < length; l += sLen {
					substr := str[l : l+sLen]
					if substr != s {
						invalidStr = true
						break
					} else {
						counter++
					}
				}

				if !invalidStr && counter > 1 {
					invalid = append(invalid, i)
				}
			}
		}
	}

	total := 0

	for _, num := range invalid {
		total += num
	}

	println(total)
}

func getInput() string {
	return `1-14,46452718-46482242,16-35,92506028-92574540,1515128146-1515174322,56453-79759,74-94,798-971,49-66,601-752,3428-4981,511505-565011,421819-510058,877942-901121,39978-50500,9494916094-9494978970,7432846301-7432888696,204-252,908772-990423,21425-25165,1030-1285,7685-9644,419-568,474396757-474518094,5252506279-5252546898,4399342-4505058,311262290-311393585,1895-2772,110695-150992,567521-773338,277531-375437,284-364,217936-270837,3365257-3426031,29828-36350`
}
