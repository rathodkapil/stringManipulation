package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	//	fmt.Println(isPallindrome("racAR"))

	//	fmt.Println(reverseString("abcdef"))

	// b := []byte("hello")
	// reverseString1(b)
	// m := make(map[int]int)
	// coinPiles(2, m)
	//arrayRotaion()
	//leftRotate([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3)
	//rotateRight([]int{1, 2, 3, 4, 5}, 3)
	//replace()
	//splitSamples()
	//builderSample()
	//findIndex()
	//fmt.Println(isAnagram("Listen", "Silent"))
	//fmt.Println(mostFrequentCharacter("hello"))
	//stringPermutations("abc")
	//fmt.Println(binarySearch([]int{2, 15, 20, 30, 40, 50, 72, 90}, 100))
	//stringCompression("aabcaaa")
	//fmt.Println(isStringRotated("waterbot.tle", "erbot.Tlewat"))
	longestSubstring("abab", "baba")

}

// find the longest substring between two given string
// a := ABA , b:= BA
// longest substring is BA
// reference via youtube video :
// https://www.youtube.com/watch?v=aVFWW3pBQFo
func longestSubstring(a, b string) string {

	/*
		creating a 2D slice (table) to hold the substring at a given point
			   A B A
			B [0 0 0]
			A [0 0 0]
	*/

	out := ""
	if len(a) == 0 || len(b) == 0 {
		return out
	}

	cache := make([][]int, len(a))
	for i := range cache {
		cache[i] = make([]int, len(b))
	}
	fmt.Println(cache)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {

			if a[i] == b[j] {
				if i == 0 || j == 0 {
					cache[i][j] = 1 //by default, putting 1 count as cache[0][1] or cache[1][0] no previous poisition is there

				} else {
					cache[i][j] = cache[i-1][j-1] + 1 // going to previous row and fetching the count and incrementing by 1
				}

				if cache[i][j] > len(out) {
					out = a[i-cache[i][j]+1 : i+1]
				}

			}
		}
	}

	fmt.Println(cache)
	fmt.Println(out)
	return out
}

func isStringRotated(str1, str2 string) bool {

	/*
		str1 - "erbottlewat"
		str2 - "waterbottle"

		KMP algo
		t := str1 + str2
		waterbottlewaterbottle
		if  strings.contains()
	*/

	if len(str1) != len(str2) {
		return false
	}

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	return strings.Contains(str1+str2, str1)

}

func stringCompression(str string) {
	//aabcccccaaa , output ::"a2b1c5a3"
	//aabc , output:: a2b1c1

	/*
		i=3 c
		j=4 c   i-1 =b   count=1 , comstr=a1b1c1


	*/

	compressed := ""
	count := 1
	prevChar := rune(0)

	for _, char := range str {

		if char == prevChar {
			count++
		} else {

			if count > 1 {
				compressed += string(prevChar) + strconv.Itoa(count)
			} else {
				compressed += string(prevChar)
			}
			count = 1
			prevChar = char
		}

	}
	// Handle the last character
	if count > 1 {
		compressed += string(prevChar) + strconv.Itoa(count)
	} else {
		compressed += string(prevChar)
	}

	// if len(compressed) >= len(str) {
	// 	return str

	// }

	fmt.Printf(compressed)

}

/*
abcd== 4*3*2*1
a-- abcd, acdb,adbc,abdc,acbd,adcb
b--bacd, bcad, bdac,badc,bcda,bdca
c--cabd, cbad, cdab,cadb,cbda,cdba
d--dabc, dbca,dcab, dcba,dacb,dbac
*/
func stringPermutations(str string) {
	runes := []rune(str)
	l := len(str) - 1
	permuteHelper(runes, 0, l) // ====finding permutation between start index i.e 0 and end index l
}

func permuteHelper(runes []rune, start int, end int) {

	if start == end {
		fmt.Println(string(runes))
	} else {

		for i := start; i <= end; i++ {

			runes[i], runes[start] = runes[start], runes[i]         //swap with start
			permuteHelper(runes, start+1, end)                      // increasing the start index value with +1
			runes[end], runes[start+1] = runes[start+1], runes[end] //swap with end
		}
	}
}

func binarySearch(numbers []int, t int) int {

	l := len(numbers) - 1
	s := 0
	m := (l - s) / 2
	m = m % l
	sort.IntSlice.Sort(numbers)

	//find the t num in the given array
	//[2 15 20 30 40 50 72 90]

	if t < numbers[m] {

		for i := s; i < m; i++ {

			if t == numbers[i] {
				return i
			}
		}
	} else if t == numbers[m] {
		return m
	} else if t > numbers[m] {
		for i := m; i < l; i++ {

			if t == numbers[i] {
				return i
			}

		}
	}
	return -1
}

func mostFrequentCharacter(str string) (rune, int) {

	m := make(map[rune]int)

	for _, r := range str {

		if _, ok := m[r]; ok {
			m[r]++
		} else {
			m[r] = 1
		}
	}

	var c rune
	var count int
	for r, v := range m {
		if v > count {
			c = r
			count = v
		}
	}

	fmt.Println(m)
	return c, count
}

func isAnagram(s1, s2 string) bool {

	var isAna = true

	if len(s1) == 0 || len(s2) == 0 {
		isAna = false
		return isAna
	}
	for _, v := range strings.ToLower(s1) {

		fmt.Println(string(v))
		if strings.Index(strings.ToLower(s2), strings.ToLower(string(v))) == -1 {
			isAna = false
			break
		}
	}

	return isAna

}
func findIndex() {

	s := "hello"

	fmt.Println(strings.Index(s, "j"))

}

func replace() {
	fmt.Println(strings.Replace("oink", "j", "l", 1))
	fmt.Println(strings.ReplaceAll("hello", "l", "jj"))
}

func builderSample() {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())
}

func splitSamples() {
	/*
		output::
		[www google com]
		[www. google. com]
		[http www.google.com]
	*/

	fmt.Println(strings.Split("www.google.com", "."))

	fmt.Println(strings.SplitAfter("www.google.com", "."))

	fmt.Println(strings.SplitN("http://www.google.com", "://", 2))

}

func leftRotate(a []int, d int) {

	// 1,2,3,4,5
	//d=3
	//output :: 4,5,1,2,3
	//fmt.Println(a)
	length := len(a)

	d = d % length
	//fmt.Println(d)

	a = append(a[d:], a[:d]...)
	fmt.Println(a)

}

func rotateRight(a []int, d int) {

	length := len(a)

	d = d % length
	// R--->>>
	//1,2, 3,4,5
	//d=3
	//output :: [3, 4, 5, 1, 2]

	fmt.Println(a[d-1 : len(a)])
	a = append(a[d-1:len(a)], a[:d-1]...)
	fmt.Println("final---", a)

}

func arrayRotaion() {

	a := []int{1, 2, 3, 4, 5, 6, 7}
	d := 4

	temA := make([]int, d)
	for i := 0; i < d; i++ {
		temA[i] = a[i]
	}
	fmt.Println("temA---", temA)
	fmt.Println("a---", a)

	a = a[d:]

	fmt.Println("after adjusting other elements in slice a---", a)
	fmt.Println("after roation a---", len(temA), len(a))

	for i := 0; i < d; i++ {
		a = append(a, temA[i])
		fmt.Println("temA[i]----", temA[i])
	}

	fmt.Println("again adding last element to a slice ----", a)

}

func reverse(a []int) {
	for i := 0; i <= len(a)-1; i++ {
		if a[i] < a[len(a)-1] {
			a[i], a[len(a)-1] = a[len(a)-1], a[i]
		}
	}
	fmt.Println(a)
}

func isPallindrome(s string) bool {

	revStr := ""
	for i, v := range s {

		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			revStr = revStr + strings.ToLower(string(s[len(s)-1-i]))
		}
	}

	fmt.Println(revStr)

	return revStr == strings.ToLower(s)

}

func reverseString(s string) string {

	revStr := ""

	for i, _ := range s {

		revStr = revStr + string(s[len(s)-1-i])
	}

	return revStr

}

func reverseString1(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i] // Swap characters using a single-line swap idiom
	}

	fmt.Println(string(s))
}

// where n is the nos of coins, we need to write a func
//which will return the number of distinct piles it can make.

/*
{5}   ---
{4,1}  ---
{3,2}  ---
{3,1,1} ---
{2,2,1} ---
{1,1,1,2} ---
{1,1,1,1,1} ---
*/
func coinPiles(n int, m map[int]int) int {

	//			n=1  ret 1
	//       n=5    1<5  cp(4)	 + cp(1) ==  cp(4)+1
	//              2 < 5 cp(3) + cp(2)  == cp(3)+ cp(2)
	//              3 <5  cp(2) + cp(3)  == cp(2) + cp(3)
	//               4 <5 cp(1) + cp(4)	=== 1 + cp(4)
	//

	return 0
}
