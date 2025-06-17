package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func chek(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sumOfTheDigitsOfHarshadNumber(x int) int {
	numb := strconv.Itoa(x)
	digits := make([]int, len(numb))
	sum := 0
	for i, char := range numb {
		digit, _ := strconv.Atoi(string(char))
		digits[i] = digit
		sum = sum + digits[i]
	}
	if sum != 0 && x%sum == 0 {
		return sum
	} else {
		return -1
	}
}

func twoSum(nums []int, target int) []int {
	numb := make([]int, 0)
	flag := false
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if flag == false && i != j && nums[i]+nums[j] == target {
				numb = append(numb, i, j)
				flag = true
			}
		}
	}
	return numb
}

func twoSum2(nums []int, target int) []int {
	dct := make(map[int]int)

	for i := range nums {
		diff := target - nums[i]
		if v, ok := dct[diff]; ok {
			return []int{i, v}
		}
		dct[nums[i]] = i
	}
	return []int{-1, -1}
}

func isPalindromeInt(x int) bool {
	if x >= 0 {
		numb := strconv.Itoa(x)
		length := len(numb) / 2
		flag := true
		for i := 0; i < length; i++ {
			if flag == true && numb[i] != numb[len(numb)-i-1] {
				flag = false
			}
		}
		return flag
	} else {
		return false
	}
}

func romanToInt(s string) int {
	temp := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'M' {
			temp += 1000
		} else if s[i] == 'D' {
			temp += 500
		} else if s[i] == 'C' {
			if i != len(s)-1 && s[i+1] == 'D' {
				temp += 400
				i++
			} else if i != len(s)-1 && s[i+1] == 'M' {
				temp += 900
				i++
			} else {
				temp += 100
			}
		} else if s[i] == 'L' {
			temp += 50
		} else if s[i] == 'X' {
			if i != len(s)-1 && s[i+1] == 'L' {
				temp += 40
				i++
			} else if i != len(s)-1 && s[i+1] == 'C' {
				temp += 90
				i++
			} else {
				temp += 10
			}
		} else if s[i] == 'V' {
			temp += 5
		} else if s[i] == 'I' {
			if i != len(s)-1 && s[i+1] == 'V' {
				temp += 4
				i++
			} else if i != len(s)-1 && s[i+1] == 'X' {
				temp += 9
				i++
			} else {
				temp += 1
			}
		}
	}
	return temp
}

func isValid2(s string) bool {
	flag := false
	for i := 0; i < len(s)-1; i += 2 {
		if string(s[i]) == "[" && string(s[i+1]) == "]" {
			flag = true
		} else if string(s[i]) == "(" && string(s[i+1]) == ")" {
			flag = true
		} else if string(s[i]) == "{" && string(s[i+1]) == "}" {
			flag = true
		} else {
			flag = false
			break
		}
	}
	return flag
}

func Pop(s []string) []string {
	if len(s) == 0 {
		return nil
	}
	s = s[:len(s)-1]
	return s
}

func isValid3(s string) bool {
	flag := false
	stack := make([]string, 0)
	for i := range s {
		if string(s[i]) == "(" || string(s[i]) == "{" || string(s[i]) == "[" {
			stack = append(stack, string(s[i]))
		} else if len(stack) != 0 && string(s[i]) == ")" && stack[len(stack)-1] == "(" {
			flag = true
			stack = stack[:len(stack)-1]
		} else if len(stack) != 0 && string(s[i]) == "}" && stack[len(stack)-1] == "{" {
			flag = true
			stack = stack[:len(stack)-1]
		} else if len(stack) != 0 && string(s[i]) == "]" && stack[len(stack)-1] == "[" {
			flag = true
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) != 0 {
		flag = false
	}
	return flag
}

func appArr(stack []string, needle string) []string {
	for i := len(needle) - 1; i >= 0; i-- {
		stack = append(stack, string(needle[i]))
	}
	return stack
}

func strStr(haystack string, needle string) int {
	stack := []string{}
	index := []int{}
	stack = appArr(stack, needle)
	for i := 0; i < len(haystack); i++ {
		if string(haystack[i]) == stack[len(stack)-1] {
			index = append(index, i)
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				goto myReturn
			}
		} else {
			stack = stack[:0]
			stack = appArr(stack, needle)
			i = i - len(index)
			index = index[:0]
		}
	}
	if len(stack) != 0 {
		return -1
	}
myReturn:
	return index[0]
}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	start := 1
	end := x

	for start <= end {
		mid := start + (end-start)/2
		sqr := mid * mid

		if sqr > x {
			end = mid - 1
		} else if sqr == x {
			return mid
		} else {
			start = mid + 1
		}
	}

	return end
}

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	sp_1 := 1
	sp_2 := 1
	for i := 2; i < n+1; i++ {
		current := sp_1 + sp_2
		sp_1 = sp_2
		sp_2 = current
	}
	return sp_2
}

func PopInt(s []int) []int {
	if len(s) == 0 {
		return nil
	}
	s = s[:len(s)-1]
	return s
}

func generate(numRows int) [][]int {
	slice := [][]int{}
	arr := []int{}
	mas := []int{}
	for length := 1; length <= numRows; length++ {
		fmt.Println(length, " - length")
		mas = arr
		arr = []int{}
		for j := 1; j <= length; j++ {
			fmt.Println(j, " - j")
			fmt.Println(mas, " - mas")
			fmt.Println(arr, " - arr")
			if j == 1 || j == length {
				arr = append(arr, 1)
			} else {
				temp := mas[j-1] + mas[j-2]
				arr = append(arr, temp)
			}
			fmt.Println(mas, " - mas")
			fmt.Println(arr, " - arr")
		}
		slice = append(slice, arr)
	}
	return slice
}

func getRow(rowIndex int) []int {
	numRows := rowIndex + 1
	slice := [][]int{}
	arr := []int{}
	mas := []int{}
	for length := 1; length <= numRows; length++ {
		mas = arr
		arr = []int{}
		for j := 1; j <= length; j++ {
			if j == 1 || j == length {
				arr = append(arr, 1)
			} else {
				temp := mas[j-1] + mas[j-2]
				arr = append(arr, temp)
			}
		}
		slice = append(slice, arr)
	}
	return slice[rowIndex]
}

func maxProfit2(prices []int) int {
	profit := 0
	min := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if prices[i]-min > profit {
				profit = prices[i] - min
			}
		}
	}
	return profit
}

func isPalindrome2(s string) bool {
	flag := true
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	reg := regexp.MustCompile(`[[:punct:]]`)
	s = reg.ReplaceAllString(s, "")
	count := len(s) / 2
	for i := 0; i < count; i++ {
		if string(s[i]) != string(s[len(s)-1-i]) {
			flag = false
		}
	}
	return flag
}

func singleNumber(nums []int) int {
	res := 0
	for i := range nums {
		fmt.Println(res, " - ", nums[i])
		res = res ^ nums[i]
	}

	return res
}

func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 1 {
		prefix = string(strs[0])
	} else if len(strs) > 1 {
		first := strs[0]
		second := strs[1]
		for i := 0; i < len(first) && i < len(second); i++ {
			if first[i] != second[i] {
				break
			} else {
				prefix = prefix + string(first[i])
			}
		}
		fmt.Println(prefix, " - prefix")
		for i := 2; i < len(strs); i++ {
			arr := strs[i]
			count := false
			if len(arr) != 0 {
				fmt.Println(arr, " - arr")
				fmt.Println(count, " - count")
				for j := 0; j < len(arr) && j < len(prefix); j++ {
					fmt.Println(prefix, " - prefix")
					fmt.Println(string(prefix[j]), " = prefix[j]")
					fmt.Println(string(arr[j]), " = arr[j]")
					fmt.Println(j, " = j")
					if string(prefix[j]) != string(arr[j]) && count == false {
						fmt.Println("FirstIf")
						prefix = ""
						break
					} else if string(prefix[j]) != string(arr[j]) && count == true {
						fmt.Println("SecondIf")
						prefix = prefix[:j]
						fmt.Println(prefix, " - prefix")
						break
					} else if arr[j] == prefix[j] {
						if j == len(prefix)-1 {
							break
						} else if j == len(arr)-1 {
							prefix = arr
						} else {
							count = true
						}
					}
				}
			} else {
				prefix = ""
			}
		}
	}
	return prefix
}

func longestCommonPrefix_great(strs []string) string {
	p := strs[0]
	for _, s := range strs {
		fmt.Println(s, " - s")
		i := 0
		for ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {
			fmt.Println("WOW")
		}
		p = p[:i]
		fmt.Println(p, " - p")
	}
	return p
}

func addBinary(a string, b string) string {
	answer := ""
	per := 0
	length := 0
	if len(a) > len(b) {
		length = len(a) - len(b)
		for i := 0; i < length; i++ {
			b = "0" + b
		}
	} else if len(b) > len(a) {
		length = len(b) - len(a)
		for i := 0; i < length; i++ {
			a = "0" + a
		}
	}
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	for i := len(a) - 1; i >= 0; i-- {
		if string(a[i]) == "0" && string(b[i]) == "0" {
			if per == 0 {
				answer = "0" + answer
				per = 0
			} else {
				answer = "1" + answer
				per = 0
			}
		} else if string(a[i]) == "1" && string(b[i]) == "1" {
			if per == 0 {
				answer = "0" + answer
				per = 1
			} else {
				answer = "1" + answer
				per = 1
			}
		} else {
			if per == 0 {
				answer = "1" + answer
			} else {
				answer = "0" + answer
				per = 1
			}
		}
	}
	if per == 1 {
		answer = "1" + answer
	}
	return answer
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	head := &ListNode{0, nil}
	curr := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	}

	return head.Next
}

func deleteDuplicates(head *ListNode) *ListNode { //11233
	if head == nil {
		return nil
	}
	mapki := head
	for mapki != nil && mapki.Next != nil {
		if mapki.Val == mapki.Next.Val {
			mapki.Next = mapki.Next.Next
		} else {
			mapki = mapki.Next
		}
	}
	return head
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	mapki := head
	current := mapki
	for current != nil && mapki.Next != nil && current.Next != nil {
		if mapki.Next.Val < current.Val {
			mapki = mapki.Next.Next
		} else {
			mapki = mapki.Next
		}
		current = current.Next
	}
	return head
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	number := 0
	number_start := 0
	count_start := 0
	count_end := 0
	if len(nums) == 1 {
		number = nums[0]
	} else {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] == nums[i+1] {
				count_start++
				number_start = nums[i]
			} else {
				if count_start > count_end {
					count_end = count_start
					number = nums[i]
				}
				count_start = 0
			}
		}
		if count_start > count_end {
			count_end = count_start
			number = number_start
		}
	}
	return number
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr3 := []int{}
	for _, letter := range arr2 {
		for i := 0; i < len(arr1); i++ {
			fmt.Println(arr1[i])
			if letter == arr1[i] {
				arr3 = append(arr3, arr1[i])
				arr1 = append(arr1[:i], arr1[i+1:]...)
				i--
			}
		}
	}
	sort.Ints(arr1)
	arr3 = append(arr3, arr1...)
	return arr3
}

func minMovesToSeat(seats []int, students []int) int {
	bubbleSort(seats)
	bubbleSort(students)
	result := 0
	for i := 0; i < len(seats); i++ {
		count := math.Abs(float64(seats[i] - students[i]))
		result = result + int(count)
	}
	return result
}

func reverseBits(num uint32) uint32 {
	binary := strconv.FormatInt(int64(num), 2)
	count := string(binary)
	amount := 32 - len(binary)
	result := 0
	for i := 0; i < len(count); i++ {
		result = result + int(count[i]-'0')*int(math.Pow(2, float64(i+amount)))
	}
	return uint32(result)
}

func hammingWeight(n int) int {
	binary := strconv.FormatInt(int64(n), 2)
	count := string(binary)
	result := 0
	for i := 0; i < len(count); i++ {
		if int(count[i]-'0') == 1 {
			result++
		}
	}
	return result
}

func isDigit(c string) bool {
	_, err := strconv.Atoi(string(c))
	if err == nil {
		return true
	} else {
		return false
	}
}

func clearDigits(s string) string {
	runes := []rune(s)
	index_let := -1
	index_dig := -1

	for i := 0; i < len(runes); i++ {
		fmt.Println("elem = ", string(runes[i]))
		if unicode.IsDigit(runes[i]) {
			index_dig = i
		} else {
			index_let = i
		}
		fmt.Println("index_let = ", index_let, "    index_dig = ", index_dig)
		if index_dig != -1 && index_let != -1 {
			runes = append(runes[:index_let], runes[index_let+1:]...)
			runes = append(runes[:index_dig-1], runes[index_dig:]...)
			index_let = -1
			index_dig = -1
			i = -1
			fmt.Println("iWOW")
		}
		fmt.Println("rune = ", string(runes))
	}

	result := string(runes)
	hasDigit := false

	for _, char := range result {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}
	if hasDigit {
		result = clearDigits(result)
	}

	return result
}

func isHappy(n int) bool {
	k := 0
	for n != 1 {
		result := 0
		str := strconv.Itoa(n)
		count := strings.Split(str, "")
		for _, elem := range count {
			num, _ := strconv.ParseFloat(elem, 64)
			result = result + int(math.Pow(num, 2))
		}
		n = result
		k++
		if k > 10 {
			break
		}
	}
	if n == 1 {
		return true
	} else {
		return false
	}
}

func factorial_2(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial_2(n-1)
}

func countCompleteDayPairs(hours []int) int {
	day := 0
	for i := 0; i < len(hours)-1; i++ {
		for j := i + 1; j < len(hours); j++ {
			if (hours[i]+hours[j])%24 == 0 {
				day++
			}
		}
	}
	return day
}

func isPrefixAndSuffix(str1 string, str2 string) bool {

	var flag bool
	if len(str1) > len(str2) {
		flag = false
	} else if str1 == str2 {
		flag = true
	} else {
		if str1 == str2[:len(str1)] && str1 == str2[len(str2)-len(str1):] {
			fmt.Println("str1 != str2[:len(str1)]")
			flag = true
		}
	}
	return flag
}

func countPrefixSuffixPairs(words []string) int {
	count := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			fmt.Println("word[i] = ", words[i], "   word[j] = ", words[j])
			if isPrefixAndSuffix(words[i], words[j]) == true {
				count++
				fmt.Println("count = ", count)
			}
			fmt.Println()
		}
	}
	return count
}

func minSteps(n int) int {
	step := 0
	if n == 1 {
		return 0
	} else {
	Procc:
		for i := n / 2; i > 0; i-- {
			if n%i == 0 {
				step += n / i
				n = i
				if i != 1 {
					goto Procc
				}
			}
		}
	}
	return step
}

func findTheDifference(s string, t string) byte {
	var result byte
	for i := 0; i < len(t); i++ {
		flag := false
		for j := 0; j < len(s); j++ {
			if t[i] == s[j] {
				flag = true
				s = s[:j] + s[j+1:]
				break
			}
		}
		if flag == false {
			result = byte(t[i])
		}
	}
	return result
}

func findComplement(num int) int {
	binary := strconv.FormatInt(int64(num), 2)
	result := ""
	fmt.Println(binary)
	for _, i := range binary {
		fmt.Println(string(i), "   ", byte(i))
		if string(i) == "0" {
			result += "1"
		} else {
			result += "0"
		}
		fmt.Println(i)
		fmt.Println(binary)
	}
	number, _ := strconv.ParseInt(result, 2, 0)
	return int(number)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	} else {
		hashmap := make(map[rune]int)
		for _, elem := range s {
			hashmap[elem]++
		}
		for _, elem := range t {
			hashmap[elem]--
			if hashmap[elem] < 0 {
				return false
			}
		}
	}
	return true
}

func repeatedSubstringPattern(s string) bool {
	if len(s) < 2 {
		return false
	} else {
		flag := false
		for i := (len(s) / 2) - 1; i > 0; i-- {
			if len(s)%i == 0 {
				fmt.Println("len(s)%i == 0")
				str := s[:i+1]
				fmt.Println("str = ", str)
				for j := i; j < len(s)/i; j += len(str) {
					fmt.Println("i = ", i, "  j = ", j)
					fmt.Println("part_1 = ", s[j+1:], "    part_2 = ", s[:j+len(str)])
					fmt.Println("part = ", s[j+1:j+len(str)+1])
					if j+1+len(str) == len(s) && str == s[j+1:] {
						fmt.Println("1")
						flag = true
					} else if str == s[j+1:j+len(str)+1] {
						fmt.Println("2")
						flag = true
					} else {
						fmt.Println("3")
						flag = false
						break
					}
				}
				if flag == true {
					return true
				}
			}
		}
		return flag
	}
}

func arraySign(nums []int) int {
	negative := 1
	for _, i := range nums {
		if i == 0 {
			return 0
		} else if i < 0 {
			negative *= -1
		}
	}
	return negative
}

func canMakeArithmeticProgression(arr []int) bool {
	count := arr[1] - arr[0]
	for i := 1; i <= len(arr)-1; i++ {
		count = arr[i] - arr[i-1]
		if arr[i]-arr[i-1] != count || (arr[i]-arr[i-1])%count != 0 || count%(arr[i]-arr[i-1]) != 0 {
			return false
		}
	}
	return true
}

func isMonotonic(nums []int) bool {
	flag_neg := false
	flag_poz := true
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			flag_neg = true
		} else if nums[i] < nums[i-1] {
			flag_poz = false
		}
	}
	if flag_neg == true && flag_poz == false {
		return false
	}
	return true
}

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
		fmt.Println(nums)
	}
	fmt.Println(nums)
	return k
}

func removeDuplicates(nums []int) int {
	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(t); i++ {
		if t[i] == s[0] {
			s = s[1:]
		}
		if len(s) == 0 {
			return true
		}
	}
	return false
}

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[string]int)
	for _, elem := range magazine {
		letters[string(elem)]++
	}
	for _, elem := range ransomNote {
		_, exists := letters[string(elem)]
		if exists {
			letters[string(elem)]--
			if letters[string(elem)] == 0 {
				delete(letters, string(elem))
			}
		} else {
			return false
		}
	}
	return true
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	} else {
		hashmap := make(map[rune]string)
		reversemap := make(map[string]rune)
		for i, elem := range s {
			_, exists := hashmap[elem]
			_, have := reversemap[string(t[i])]
			if (exists || have) && hashmap[elem] != string(t[i]) {
				return false
			}
			hashmap[elem] = string(t[i])
			reversemap[string(t[i])] = elem

		}
	}
	return true
}

func wordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")
	if len(pattern) != len(arr) {
		return false
	}
	hashmap := make(map[rune]string)
	reversemap := make(map[string]rune)
	for ind, elem := range pattern {
		_, exists := hashmap[elem]
		_, have := reversemap[arr[ind]]
		if (exists || have) && hashmap[elem] != arr[ind] {
			return false
		}
		hashmap[elem] = arr[ind]
		reversemap[arr[ind]] = elem
	}

	return true
}

func maxProfit(prices []int) int {
	min := []int{prices[0]}
	temp := prices[0]
	for i, el := range prices {
		if el < min[i] {
			temp = el
		}
		min = append(min, temp)
	}
	profit := 0
	for i, el := range prices {
		fmt.Println(el, " - ", prices[i], " = ", el-prices[i])
		if el-prices[i] > profit {
			profit = el - prices[i]
		}
	}
	return profit
}

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			count++
		} else if count != 0 && string(s[i]) == " " {
			goto myReturn
		}
	}
myReturn:
	return count
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 || len(nums) < 2 {
		return false
	}

	indexMap := make(map[int]int)

	for i, num := range nums {
		if prevIndex, exists := indexMap[num]; exists && i-prevIndex <= k {
			return true
		}
		indexMap[num] = i
	}

	return false
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	result := []string{}

	start := nums[0]
	for i := 1; i <= len(nums); i++ {
		if nums[i] != nums[i-1]+1 || i == len(nums) {
			if start == nums[i-1] {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}

			if i != len(nums) {
				start = nums[i]
			}
		}
	}
	return result
}

func longestPalindrome(s string) int {
	result := 0
	hashmap := make(map[rune]int)
	for _, char := range s {
		hashmap[char]++
	}
	hasOdd := false
	for _, value := range hashmap {
		if value%2 == 0 {
			result += value
		} else {
			result += value - 1
			hasOdd = true
		}
	}
	if hasOdd {
		result++
	}
	return result
}

func searchInsert(nums []int, target int) int {
	index := 0
	for i, numb := range nums {
		if numb < target {
			index = i + 1
		} else {
			return i
		}
	}
	return index
}

func reverseStr(s string, k int) {
	stack := []string{}
	count := false
	for i := k - 1; i < len(s); i += k {
		count = !count
		if count {
			for j := i; j > i-k; j-- {
				stack = append(stack, string(s[i]))
			}
		} else {
			for j := i + 1 - k; j <= i; j++ {
				stack = append(stack, string(s[i]))
			}
		}
	}
	fmt.Println(stack)
}

func plusOne(digits []int) []int {
	stack := []int{1}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			break
		} else {
			digits[i] = 0
			if i == 0 {
				stack = append(stack, digits...)
				return stack
			}
		}
	}
	return digits
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	count := 0
	for i := 0; i < len(arr1); i++ {
		count++
		for j := 0; j < len(arr2); j++ {
			if math.Abs(float64(arr1[i]-arr2[j])) <= float64(d) {
				count--
				break
			}
		}
	}
	return count
}

func minOperations(nums []int, k int) int {
	count := 0
	for _, num := range nums {
		if num < k {
			count++
		}
	}
	return count
}

func resultArray(nums []int) []int {
	n := len(nums)
	result := []int{}
	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}
	for i := 2; i < n; i++ {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}
	result = arr1
	result = append(result, arr2...)
	return result
}

func minimumBoxes(apple []int, capacity []int) int {
	apples := 0
	for _, app := range apple {
		apples += app
	}
	fmt.Printf("apples = %d\n", apples)
	bubbleSort(capacity)
	count := 0
	v := 0
	for i := len(capacity) - 1; i >= 0; i-- {
		if v >= apples {
			break
		} else {
			v += capacity[i]
			count++
		}
	}
	return count
}

func encrypt(x uint64) uint64 {
	strx := strconv.Itoa(int(x))
	fmt.Println(len(strx))
	maxnumb := 0
	for _, el := range strx {
		elint, _ := strconv.Atoi(string(el))
		if elint > maxnumb {
			maxnumb = elint
		}
	}
	resnumb := ""
	for i := 0; i < len(strx); i++ {
		number := strconv.Itoa(maxnumb)
		fmt.Println("Yes")
		resnumb = resnumb + number
	}
	result, _ := strconv.Atoi(resnumb)
	return uint64(result)
}

func sumOfEncryptedInt(nums []int) int {
	result := 0
	for _, num := range nums {
		fmt.Printf("olsnum = %d\n", num)
		num = int(encrypt(uint64(num)))
		fmt.Printf("newnum = %d\n", num)
		result += num
		fmt.Printf("result = %d\n", result)
	}
	return result
}

func checkRecord(s string) bool {
	missed := 0
	strike := 0
	for _, point := range s {
		if string(point) == "L" {
			strike++
		} else {
			if string(point) == "A" {
				missed++
			}
			strike = 0
		}
		if strike >= 3 && missed >= 2 {
			return false
		}
	}
	return true
}

func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	amount_day := 0

	hashmap := make(map[string]int)
	hashmap["1"] = 31
	if year%4 == 0 {
		hashmap["2"] = 29
	} else {
		hashmap["2"] = 28
	}
	if year == 1900 {
		hashmap["2"] = 28
	}
	hashmap["3"] = 31
	hashmap["4"] = 30
	hashmap["5"] = 31
	hashmap["6"] = 30
	hashmap["7"] = 31
	hashmap["8"] = 31
	hashmap["9"] = 30
	hashmap["10"] = 31
	hashmap["11"] = 30
	hashmap["12"] = 31

	for i := 1; month-i > 0; i++ {
		monthstr := strconv.Itoa(month - i)
		fmt.Println(hashmap[monthstr])
		amount_day += hashmap[monthstr]
	}

	amount_day += day

	return amount_day
}

func isPossibleToSplit(nums []int) bool {
	hashmap := make(map[int]int)

	for _, num := range nums {
		if _, exist := hashmap[num]; exist {
			hashmap[num]++
		} else {
			hashmap[num] = 1
		}
	}

	fmt.Println(hashmap)

	for _, value := range hashmap {
		if value > 2 {
			return false
		}
	}
	return true
}

func distributeCandies(candyType []int) int {
	hashmap := make(map[int]int)

	for _, candy := range candyType {
		if _, exist := hashmap[candy]; exist {
			hashmap[candy]++
		} else {
			hashmap[candy] = 1
		}
	}

	if len(hashmap) <= len(candyType)/2 {
		return len(hashmap)
	} else {
		return len(candyType) / 2
	}
}

func mostFrequent(nums []int, key int) int {
	hashmap := make(map[int]int)
	target := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			hashmap[nums[i+1]]++
		}
	}
	count := 0
	for key, value := range hashmap {
		if value > count {
			count = value
			target = key
		}
	}
	return target
}

func repeatedNTimes(nums []int) int {
	hashmap := make(map[int]int)
	for _, num := range nums {
		if _, exsist := hashmap[num]; exsist {
			return num
		} else {
			hashmap[num]++
		}
	}
	return 0
}

func reverseInvert(str string) string {
	runes := []rune(str)
	n := len(runes)
	inverted := make([]rune, n)

	for i := 0; i < n; i++ {

		if runes[i] == '0' {
			inverted[n-1-i] = '1'
		} else {
			inverted[n-1-i] = '0'
		}
	}

	return string(inverted)
}

func findKthBit(n int, k int) byte {
	byte_sl := make([]string, n)
	byte_sl[0] = "0"
	byte_str := ""

	if n == 1 {
		byte_str += string(byte_sl[0])
	}

	for i := 1; i < n; i++ {
		byte_sl[i] = byte_sl[i-1] + "1" + reverseInvert(byte_sl[i-1])
		if i == n-1 {
			byte_str += string(byte_sl[i])
		}
	}

	if k > len(byte_str) {
		return 0
	}

	return byte_str[k-1]
}

func specialArray(nums []int) int {
	for i := 0; i <= 10; i++ {
		count := 0
		for _, num := range nums {
			if num >= i {
				count++
			}
		}
		if count == i {
			return count
		}
	}
	return -1
}

func mostFrequentEven(nums []int) int {
	hashmap := make(map[int]int)
	rate := 0
	res_num := -1

	for _, num := range nums {
		hashmap[num]++
	}
	fmt.Println(hashmap)
	fmt.Println(res_num)
	for key, value := range hashmap {
		fmt.Printf("key : %d;   value : %d;    res_num : %d;    rate : %d\n", key, value, res_num, rate)
		if (value > rate || (value == rate && key < res_num)) && key%2 == 0 {
			fmt.Println("YES")
			res_num = key
			rate = value
		}
	}
	return res_num
}

func checkOnesSegment(s string) bool {
	flag_check := false
	for _, num := range s {
		if num == '0' {
			flag_check = true
		} else if num == '1' && flag_check == true {
			return false
		}
	}
	return true
}

func defangIPaddr(address string) string {
	defAdress := ""
	point := 0
	count := 0
	last_part := ""

	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			count++
			defAdress = defAdress + address[point:i] + "[.]"
			point = i + 1
		}
		if count == 3 {
			last_part = address[point:]
			defAdress = defAdress + last_part
			break
		}
	}
	return defAdress
}

func numberOfLines(widths []int, s string) []int {
	result := []int{}
	pixel := 0
	strs := 1

	for i := 0; i < len(widths); i++ {
		// fmt.Printf("pixels : %d;    strs : %d;   letter : %s;   weight : %d\n", pixel, strs, s[i], widths[i])
		if pixel+widths[i] <= 100 {
			pixel += widths[i]
		} else {
			pixel = 0
			strs++
		}
	}

	result = append(result, strs)
	result = append(result, pixel)

	return result
}

func subtractProductAndSum(n int) int {
	str := strconv.Itoa(n)
	sum := 0
	factorial := 1
	for i := 0; i < len(str); i++ {
		number, _ := strconv.Atoi(string(str[i]))
		fmt.Printf("number = %d;   sum = %d;   factorial = %d\n", number, sum, factorial)
		sum += number
		factorial *= number
	}
	return factorial - sum
}

func heightChecker(heights []int) int {
	old_arr := []int{}
	for _, num := range heights {
		old_arr = append(old_arr, num)
	}
	bubbleSort(heights)
	count := 0
	fmt.Println(heights)
	fmt.Println(old_arr)
	for i := 0; i < len(heights); i++ {
		if heights[i] != old_arr[i] {
			count++
		}
	}
	return count
}

func intersection(nums1 []int, nums2 []int) []int {
	type empty struct{}
	hashmap := make(map[int]empty)
	result := []int{}

	for _, num := range nums1 {
		if _, exist := hashmap[num]; !exist {
			hashmap[num] = empty{}
		}
	}

	for _, num := range nums2 {
		if _, exist := hashmap[num]; exist {
			result = append(result, num)
			delete(hashmap, num)
		}
	}

	return result
}

func addStrings(num1 string, num2 string) string {

	dif := 0
	if len(num1) > len(num2) {
		dif = len(num1) - len(num2)
		for i := 0; i < dif; i++ {
			num2 = "0" + num2
		}
	} else if len(num1) < len(num2) {
		dif = len(num2) - len(num1)
		for i := 0; i < dif; i++ {
			num1 = "0" + num1
		}
	}

	fmt.Println(num1)
	fmt.Println(num2)

	result := ""
	count := 0

	for i := len(num1) - 1; i >= 0; i-- {
		n1, _ := strconv.Atoi(string(num1[i]))
		n2, _ := strconv.Atoi(string(num2[i]))
		fmt.Printf("n1 = %d    n2 = %d\n", n1, n2)
		oper := n1 + n2 + count
		str_oper := strconv.Itoa(oper)
		fmt.Printf("oper = %d    str_oper = %s\n", oper, str_oper)
		if len(str_oper) > 1 {
			count = 1
		} else {
			count = 0
		}
		fmt.Printf("count = %d\n", count)

		if i == 0 && count == 1 {
			result += string(str_oper[len(str_oper)-1]) + "1"
		} else {
			result += string(str_oper[len(str_oper)-1])
		}
		fmt.Printf("result = %s\n", string(result))
	}

	reverse_result := ""
	for i := len(result) - 1; i >= 0; i-- {
		reverse_result += string(result[i])
	}

	return reverse_result
}

func pivotInteger(n int) int {
	start_sum := 0
	end_sum := 0

	for i := 0; i <= n/2; i++ {
		start_sum += i
	}

	for i := n / 2; i <= n; i++ {
		end_sum += i
	}

	flags := []int{}
	for i := n / 2; i >= 0 || i <= n; {
		if start_sum > end_sum {
			start_sum -= i
			i--
			end_sum += i
			flags = append(flags, 0)
		} else if start_sum < end_sum {
			end_sum -= i
			i++
			start_sum += i
			flags = append(flags, 1)
		} else {
			return i
		}
		if len(flags) > 1 && flags[len(flags)-1] != flags[len(flags)-2] {
			break
		}
	}
	return -1
}

func countLargestGroup(n int) int {
	hashmap := make(map[int]int)

	for i := 1; i <= n; i++ {
		sum := 0
		strnum := strconv.Itoa(i)
		for j := 0; j < len(strnum); j++ {
			num, _ := strconv.Atoi(string(strnum[j]))
			sum += num
		}
		hashmap[sum]++
	}

	often := 0
	for _, val := range hashmap {
		if val > often {
			often = val
		}
	}

	count := 0
	for _, val := range hashmap {
		if val == often {
			count++
		}
	}

	return count
}

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

func lemonadeChange(bills []int) bool {
	sum5th := 0
	sum10th := 0
	for _, bill := range bills {
		if bill == 5 {
			sum5th++
		} else if bill == 10 {
			if sum5th == 0 {
				return false
			}
			sum5th--
			sum10th++
		} else {
			if sum5th < 3 && (sum10th < 1 || sum5th < 1) {
				return false
			} else if sum10th >= 1 && sum5th >= 1 {
				sum10th--
				sum5th--
			} else if sum5th >= 3 {
				sum5th -= 3
			}
		}
	}
	return true
}

func longestMonotonicSubarray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	increasing := [][]int{}
	decreasing := [][]int{}

	incrow := []int{nums[0]}
	decrow := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		if incrow[len(incrow)-1] < nums[i] {
			incrow = append(incrow, nums[i])
			if i == len(nums)-1 {
				increasing = append(increasing, incrow)
			}
		} else {
			increasing = append(increasing, incrow)
			incrow = []int{nums[i]}
		}

		if decrow[len(decrow)-1] > nums[i] {
			decrow = append(decrow, nums[i])
			if i == len(nums)-1 {
				decreasing = append(decreasing, decrow)
			}
		} else {
			decreasing = append(decreasing, decrow)
			decrow = []int{nums[i]}
		}
	}

	incLength := 0
	decLength := 0

	for _, row := range increasing {
		if len(row) > incLength {
			incLength = len(row)
		}
	}

	for _, row := range decreasing {
		if len(row) > decLength {
			decLength = len(row)
		}
	}

	if incLength > decLength {
		return incLength
	} else {
		return decLength
	}
}

func longestContinuousSubstring(s string) int {
	alfa := "abcdefghijklmnopqrstuvwxyz"
	i := 0
	count := 0
	mascounts := []int{}
	for len(s) != 0 {
		if alfa[i] != s[0] && count != 0 {
			mascounts = append(mascounts, count)
			count = 0
		} else {
			count++
			s = s[1:]
			if len(s) == 0 {
				mascounts = append(mascounts, count)
			}
		}

		if i == 25 {
			i = -1
			mascounts = append(mascounts, count)
			count = 0
		}
		i++
	}

	fmt.Println("mascounts = ", mascounts)

	maxlen := 0

	for _, val := range mascounts {
		if val > maxlen {
			maxlen = val
		}
	}

	return maxlen
}

func validMountainArray(arr []int) bool {
	if len(arr) < 2 || arr[0] > arr[1] {
		return false
	}

	flag := true

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return false
		} else if arr[i] > arr[i+1] {
			flag = false
		} else if arr[i] < arr[i+1] && flag == false {
			return false
		}
	}

	if flag == true {
		return false
	} else {
		return true
	}
}

func numUniqueEmails(emails []string) int {
	localnames := make(map[string]int)
	for _, email := range emails {
		localname := ""
		flag := true
		for i := 0; i < len(email); i++ {
			if email[i] == '.' {
				part1 := email[:i]
				part2 := email[i+1:]
				email = part1 + part2
				i--
			} else if email[i] == '+' {
				flag = false
			} else if email[i] == '@' {
				localname += email[i:]
				break
			} else if flag == true {
				localname += string(email[i])
			}
		}
		localnames[localname]++
	}

	fmt.Println(localnames)

	return len(localnames)
}

func uniqueMorseRepresentations(words []string) int {
	morseCode := map[rune]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
	}

	hashmap := make(map[string]int)

	for _, word := range words {
		morseWord := ""
		for i := 0; i < len(word); i++ {
			morse := morseCode[rune(word[i])]
			morseWord += morse
		}
		hashmap[morseWord]++
	}

	return len(hashmap)
}

func evenOddBit(n int) []int {
	binaryStr := strconv.FormatInt(int64(n), 2)
	even := 0
	odd := 0
	for i, j := len(binaryStr)-1, 0; i >= 0; i-- {
		if binaryStr[i] == '1' {
			if j%2 == 0 {
				even++
			} else {
				odd++
			}
		}
		j++
	}

	result := []int{even, odd}

	return result
}

func equalFrequency(word string) bool {
	hashmap := make(map[rune]int)
	hashfreq := make(map[int]int)
	for _, let := range word {
		hashmap[let]++
	}

	for _, val := range hashmap {
		hashfreq[val]++
	}

	if len(hashfreq) == 1 {
		for num := range hashfreq {
			return num == 1 || hashfreq[num] == 1
		}
	}

	if len(hashfreq) == 2 {
		var f1, f2 int
		for k := range hashfreq {
			if f1 == 0 {
				f1 = k
			} else {
				f2 = k
			}
		}

		if f1 > f2 {
			f1, f2 = f2, f1
		}

		return (f1 == 1 && hashfreq[f1] == 1) || (f2 == f1+1 && hashfreq[f2] == 1)
	}

	return false

}

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 2 {
		return true
	}

	x0, y0 := coordinates[0][0], coordinates[0][1]
	x1, y1 := coordinates[1][0], coordinates[1][1]

	dx := x1 - x0
	dy := y1 - y0

	for i := 2; i < len(coordinates); i++ {
		x, y := coordinates[i][0], coordinates[i][1]
		if dy*(x-x0) != dx*(y-y0) {
			return false
		}
	}

	return true
}

func maxProduct(nums []int) int {
	maxFirstVal := 0
	i := 0
	for index, num := range nums {
		if num > maxFirstVal {
			i = index
			maxFirstVal = num
		}
	}

	nums = append(nums[:i], nums[i+1:]...)

	maxSecondVal := 0
	i = 0
	for index, num := range nums {
		if num > maxSecondVal {
			i = index
			maxSecondVal = num
		}
	}

	return (maxFirstVal - 1) * (maxSecondVal - 1)
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func containsPattern(arr []int, m int, k int) bool {
	count := 1
	part := []int{}
	for i := 0; i < len(arr)-m+1; i++ {
		temp := arr[i : i+m]
		fmt.Println(temp)
		if valid := slicesEqual(temp, part); !valid {
			part = temp
			count = 1
		} else {
			count++
			i += m - 1
		}
		fmt.Printf("count = %d ", count)
		if count == k {
			return true
		}
	}
	return false
}

func fizzBuzz(n int) []string {
	arr := []string{"1"}
	for i := 2; i <= n; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			arr = append(arr, "FizzBuzz")
		} else if i%5 == 0 {
			arr = append(arr, "Buzz")
		} else if i%3 == 0 {
			arr = append(arr, "Fizz")
		} else {
			num := strconv.Itoa(i)
			arr = append(arr, num)
		}
	}

	return arr
}

func uncommonFromSentences(s1 string, s2 string) []string {
	hashmap := make(map[string]int)
	result := []string{}

	slice1 := strings.Split(s1, " ")
	slice2 := strings.Split(s2, " ")

	for _, str := range slice1 {
		hashmap[str]++
	}

	for _, str := range slice2 {
		hashmap[str]++
	}

	for key, val := range hashmap {
		if val == 1 {
			result = append(result, key)
		}
	}

	return result
}

func lastVisitedIntegers(nums []int) []int {
	seen := []int{}
	ans := []int{}

	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			k = 0
			num := []int{nums[i]}
			seen = append(num, seen...)
		} else {
			k++
			if k <= len(seen) {
				ans = append(ans, seen[k-1])
			} else {
				ans = append(ans, -1)
			}
		}
	}

	return ans
}

func maxDivScore(nums []int, divisors []int) int {
	hashmap := make(map[int]int)
	for _, div := range divisors {
		count := 0
		for _, num := range nums {
			if num%div == 0 {
				count++
			}
		}
		hashmap[div] = count
	}

	if len(hashmap) == 0 {
		for _, div := range divisors {
			hashmap[div]++
		}
	}

	maxVal := 0
	minKey := 100

	for key, val := range hashmap {
		if val > maxVal {
			maxVal = val
			minKey = key
		} else if val == maxVal && key < minKey {
			minKey = key
		}
		fmt.Printf("Key: %d     Val: %d\n", minKey, maxVal)
	}

	return minKey

}

func countPairs(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				count++
			}
		}
	}
	return count
}

func splitWordsBySeparator(words []string, separator byte) []string {
	array := []string{}

	for _, word := range words {
		subarray := strings.Split(word, string(separator))
		fmt.Println(subarray)
		array = append(array, subarray...)
	}

	for i := 0; i < len(array); i++ {
		if len(array[i]) == 0 || array[0] == "," {
			array = append(array[:i], array[i+1:]...)
		}
	}
	return array
}

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	} else {
		return edges[0][1]
	}
}

func smallestNumber(num int64) int64 {
	num_str := strconv.FormatInt(num, 10)
	fmt.Printf("num_str: %s\n", num_str)
	num_result := ""
	if string(num_str[0]) == "-" {
		fmt.Println("-")
		for i := 9; i >= 0; i-- {
			for j := 1; j < len(num_str); j++ {
				digit, _ := strconv.Atoi(string(num_str[j]))
				fmt.Printf("digit: %d    i: %d\n", digit, i)
				if digit == i {
					num_result += string(num_str[j])
				}
			}
		}
		num_result = "-" + num_result
	} else {
		fmt.Println("+")
		count := 0
		for i := 1; i <= 9; i++ {
			for j := 0; j < len(num_str); j++ {
				digit, _ := strconv.Atoi(string(num_str[j]))
				fmt.Printf("digit: %d    i: %d\n", digit, i)
				if digit == i {
					num_result += string(num_str[j])
				}
				if digit == 0 {
					count++
				}
			}
		}
		num_result = num_result[count:]
		part_s := string(num_result[0])
		part_e := string(num_result[1:])
		for i := 0; i < count; i++ {
			part_s += "0"
		}
		num_result = part_s + part_e
	}
	fmt.Printf("num_result: %s\n", num_result)
	num_digit, _ := strconv.ParseInt(num_result, 10, 64)
	return num_digit
} //wrong

func mergeAlternately(word1 string, word2 string) string {
	n := 0
	if len(word1) >= len(word2) {
		n = len(word2)
	} else {
		n = len(word1)
	}
	fmt.Printf("word1: %s,   word2: %s\n", word1, word2)
	result := ""
	for i := 0; i < n; i++ {
		if len(word1) != 0 {
			result += string(word1[0])
			if len(word1) == 1 {
				word1 = ""
			} else {
				word1 = word1[1:]
			}
		}
		if len(word2) != 0 {
			result += string(word2[0])
			if len(word2) == 1 {
				word2 = ""
			} else {
				word2 = word2[1:]
			}
		}
		fmt.Printf("word1: %s,   word2: %s\n", word1, word2)
		fmt.Printf("result: %s\n", result)

	}
	if len(word1) != 0 {
		result += string(word1)
	}
	if len(word2) != 0 {
		result += string(word2)
	}
	return result
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	len1 := len(str1)
	len2 := len(str2)

	gcdLength := gcd(len1, len2)

	return str1[:gcdLength]
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max_candies := 0
	for _, candy := range candies {
		if candy >= max_candies {
			max_candies = candy
		}
	}
	fmt.Printf("Max: %d\n", max_candies)
	array_bool := []bool{}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max_candies {
			fmt.Printf("Candies: %d   Candies+extra: %d  TRUE\n", candies[i], candies[i]+extraCandies)
			array_bool = append(array_bool, true)
		} else {
			fmt.Printf("Candies: %d   Candies+extra: %d  FALSE\n", candies[i], candies[i]+extraCandies)
			array_bool = append(array_bool, false)
		}
	}
	return array_bool
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	array := []int{}
	for _, flower := range flowerbed {
		if flower == 1 && count != 0 {
			array = append(array, count)
			count = 0
		} else if flower == 0 {
			count++
		}
	}
	fmt.Println("Array: ", array)
	for _, a := range array {
		if a%2 == 0 {
			n -= a/2 - 1
		} else {
			n -= a / 2
		}
	}
	fmt.Printf("n: %d\n", n)
	if n <= 0 {
		return true
	} else {
		return false
	}
}

func findLucky(arr []int) int {
	hashmap := make(map[int]int)
	for _, num := range arr {
		if _, exist := hashmap[num]; exist {
			hashmap[num]++
		} else {
			hashmap[num] = 1
		}
	}

	max_key := -1
	for key, val := range hashmap {
		if key == val && key > max_key {
			max_key = key
		}
	}

	return max_key
}

func maximizeSum(nums []int, k int) int {
	max_num := 0
	for _, num := range nums {
		if num > max_num {
			max_num = num
		}
	}

	answer := 0
	for i := 1; i <= k; i++ {
		answer += max_num
		max_num++
	}

	return answer
}

func factorial(n int) int {
	result := 1
	if n == 0 {
		return 1
	} else {
		for n > 0 {
			fmt.Println("result = ", result, "   n = ", n)
			result = result * n
			n--
		}
	}
	return result
}

func maxFrequencyElements(nums []int) int {
	freqNum := make(map[int]int)
	maxFreq := 1
	sumFreq := 0

	for _, num := range nums {
		freqNum[num]++
		if freqNum[num] > maxFreq {
			sumFreq = freqNum[num]
			maxFreq = freqNum[num]
		} else if freqNum[num] == maxFreq {
			sumFreq += freqNum[num]
		}
	}

	return sumFreq
}

func uniqueOccurrences(arr []int) bool {
	freqNum := make(map[int]int)
	freqFreq := make(map[int]int)

	for _, el := range arr {
		freqNum[el]++
	}

	for _, val := range freqNum {
		fmt.Println(val)
		if _, ok := freqFreq[val]; ok {
			return false
		} else {
			freqFreq[val]++
		}
	}

	return true

}

func missingNumber(nums []int) int {
	factorialNum := func(n int) int {
		result := 1
		if n == 0 {
			return 1
		} else {
			for n > 0 {
				result = result + n
				n--
			}
		}
		return result
	}(len(nums))
	fmt.Println("Factorial: ", factorialNum)
	sumMums := 0
	for _, num := range nums {
		sumMums += num
	}
	fmt.Println("Sum: ", sumMums)
	return factorialNum - sumMums
}

func checkIfPangram(sentence string) bool {

	letters := make(map[rune]bool)
	for _, let := range sentence {
		if _, exist := letters[let]; !exist {
			letters[let] = true
		}
	}

	if len(letters) == 26 {
		return true
	}
	return false
}

func destCity(paths [][]string) string {
	cityTrip := make(map[string]bool)
	for _, path := range paths {
		cityTrip[path[0]] = true
		if _, exist := cityTrip[path[1]]; !exist {
			cityTrip[path[1]] = false
		}
	}

	for key, val := range cityTrip {
		if val == false {
			return key
		}
	}

	return ""
}

func containsDuplicate(nums []int) bool {
	mapNums := make(map[int]bool)
	for _, num := range nums {
		if _, exist := mapNums[num]; exist {
			return true
		} else {
			mapNums[num] = true
		}
	}
	return false
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	indN1 := m - 1
	indN2 := n - 1
	lastIndN1 := m + n - 1

	for indN2 >= 0 && indN1 >= 0 {
		if nums1[indN1] > nums2[indN2] {
			nums1[lastIndN1] = nums1[indN1]
			indN1--
		} else {
			nums1[lastIndN1] = nums2[indN2]
			indN2--
		}
		lastIndN1--
	}

	for indN2 >= 0 {
		nums1[lastIndN1] = nums2[indN2]
		indN2--
		lastIndN1--
	}
}

func reverseString(s []byte) {
	slices.Reverse(s)
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	l := 0
	r := len(str) - 1
	for l <= r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func sortedSquares(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}
	sort.Ints(nums)
	return nums
}

func reversePrefix(word string, ch byte) string {
	result := word
	for i := range len(word) {
		if word[i] == ch {
			str := word[:i+1]
			arr := strings.Split(str, "")
			slices.Reverse(arr)
			str_reverse := strings.Join(arr, "")
			result = str_reverse + word[i+1:]
			break
		}
	}
	return result
}

func moveZeroes(nums []int) {
	zero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
		}
	}
}

func getCommon(nums1 []int, nums2 []int) int {
	l := 0
	r := 0
	for l < len(nums1) && r < len(nums2) {
		if nums1[l] > nums2[r] {
			r++
		} else if nums1[l] < nums2[r] {
			l++
		} else {
			return nums1[l]
		}
	}
	return -1
}

func reverseOnlyLetters(s string) string {
	str_slice := strings.Split(s, "")
	l := 0
	r := len(str_slice) - 1
	for l < r {
		if unicode.IsLetter(rune(s[l])) && unicode.IsLetter(rune(s[r])) {
			temp := str_slice[l]
			str_slice[l] = str_slice[r]
			str_slice[r] = temp
			l++
			r--
		} else {
			if !unicode.IsLetter(rune(s[l])) {
				l++
			}
			if !unicode.IsLetter(rune(s[r])) {
				r--
			}
		}
	}
	result := strings.Join(str_slice, "")
	return result
}

func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		b_word := []byte(word)
		for l, r := 0, len(word)-1; l < r; l, r = l+1, r-1 {
			b_word[l], b_word[r] = b_word[r], b_word[l]
		}
		words[i] = string(b_word)
	}
	return strings.Join(words, " ")
}

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	return NumArray{arr: prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.arr[right+1] - this.arr[left]
}

func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	prefix := make([]int, len(nums)+1)
	for i := range nums {
		fmt.Printf("sum: %v     prefix[i]: %v\n", sum, prefix[i])
		sum -= nums[i]
		if prefix[i] == sum {
			return i
		}
		prefix[i+1] = prefix[i] + nums[i]
	}
	return -1
}

func largestAltitude(gain []int) int {
	max := 0
	height := make([]int, len(gain)+1)
	for i := range gain {
		height[i+1] = height[i] + gain[i]
		if height[i+1] > max {
			max = height[i+1]
		}
	}
	return max
}

func isValid(s string) bool {

	stack := make([]byte, len(s))
	for i := range s {
		fmt.Printf("stack: %v\n", stack)
		if (s[i] == ')' && stack[len(stack)-1] == '(') || (s[i] == ']' && stack[len(stack)-1] == '[') || (s[i] == '}' && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			return false
		}
	}
	if len(stack) != len(s) {
		return false
	}
	return true
}

func generateParenthesis(n int) []string {
	var result []string
	generateParenthesisHealper(n, &result, 0, 0, "")
	return result
}

func generateParenthesisHealper(n int, result *[]string, open int, closed int, str string) {
	if len(str) == 2*n {
		*result = append(*result, str)
	}

	if open < n {
		generateParenthesisHealper(n, result, open+1, closed, str+"(")
	}

	if closed < open {
		generateParenthesisHealper(n, result, open, closed+1, str+")")
	}
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	i := 0
	minusSign := false

	if s[i] == '+' {
		i = 1
	} else if s[i] == '-' {
		minusSign = true
		i = 1
	}

	result := 0
	for ; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			break
		}
		result = result*10 + int(s[i]-'0')
		if result > math.MaxInt32 || result < math.MinInt32 {
			if minusSign {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	if minusSign {
		result = -result
	}

	return result
}

func reverse(x int) int {
	result := 0
	minusSign := false

	str := strconv.Itoa(x)
	if str[0] == '-' {
		minusSign = true
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	for i := 0; i < len(str); i++ {
		result = result + (int(str[i]-'0') * int(math.Pow(float64(10), float64(i))))
		if result < math.MinInt32 || result > math.MaxInt32 {
			return 0
		}
	}

	if minusSign {
		result = -result
	}

	return result
}

type generic interface {
	int | string | float64
}

func bubbleSort[G generic](arr []G) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func main() {
}
