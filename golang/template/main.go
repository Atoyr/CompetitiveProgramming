package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	nextReader = newScanner()

	N := nextInt()
	s := nextString()
	sheet := make([]int, N)
	N = N - 1
	for i := range s {
		no := i + 1
		if s[i] == 'L' {
			if sheet[0] == 0 {
				sheet[0] = no
			} else {
				first := -1
				before := -1
				ok := false
				for j := 0; j < len(sheet); j++ {
					if sheet[j] == 0 {
						if first < 0 {
							first = j
						}
						if before == 0 {
							sheet[j] = no
							ok = true
							break
						}

					}
					before = sheet[j]
				}
				if ok == false {
					sheet[first] = no
				}
			}

		} else {
			if sheet[N] == 0 {
				sheet[N] = no
			} else {
				first := -1
				before := -1
				ok := false
				for j := 0; j < len(sheet); j++ {
					if sheet[N-j] == 0 {
						if first < 0 {
							first = N - j
						}
						if before == 0 {
							sheet[N-j] = no
							ok = true
							break
						}
					}
					before = sheet[N-j]
				}
				if ok == false {
					sheet[first] = no
				}
			}
		}
	}
	for i := range sheet {
		fmt.Printf("%d\n", sheet[i])
	}
}

var nextReader func() string

func newScanner() func() string {
	r := bufio.NewScanner(os.Stdin)
	r.Buffer(make([]byte, 1024), int(1e+11))
	r.Split(bufio.ScanWords)
	return func() string {
		r.Scan()
		return r.Text()
	}
}
func nextString() string {
	return nextReader()
}
func nextInt64() int64 {
	v, _ := strconv.ParseInt(nextReader(), 10, 64)
	return v
}
func nextInt() int {
	v, _ := strconv.Atoi(nextReader())
	return v
}
func nextInts(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = nextInt()
	}
	return r
}
func nextInt64s(n int) []int64 {
	r := make([]int64, n)
	for i := 0; i < n; i++ {
		r[i] = nextInt64()
	}
	return r
}
func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(nextReader(), 64)
	return f
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
