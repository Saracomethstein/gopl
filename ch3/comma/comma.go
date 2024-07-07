package main

import (
	"bytes"
	"sort"
)

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaNonRecursion(s string) string {
	var buf bytes.Buffer
	n := len(s)
	start := n % 3
	if start > 0 {
		buf.WriteString(s[:start])
	}

	for i := start; i < n; i += 3 {
		if buf.Len() > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}

func commaFloat(s string) string {
	var buf bytes.Buffer
	n := len(s)
	start := n % 3
	end := 0

	if start > 0 {
		buf.WriteString(s[:start])
	}

	for i := start; i < n; i += 3 {
		if buf.Len() > 0 {
			if s[i] == '.' {
				end = i
				break
			}
			buf.WriteByte(',')
		}
		buf.WriteString(s[i : i+3])
	}

	for i := end; i < len(s); i++ {
		buf.WriteByte(s[i])
	}

	return buf.String()
}

func isAnagram(str1, str2 string) bool {
	if str1 == str2 {
		return false
	}

	if len(str1) != len(str2) {
		return false
	}

	if sortString(str1) == sortString(str2) {
		return true
	}

	return false
}

func sortString(s string) string {
	arr := []byte(s)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return string(arr)
}
