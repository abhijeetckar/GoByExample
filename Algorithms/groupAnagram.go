package main

import "fmt"

func main() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	twoD := make([][]string, len(strs))
	twoD =
		fmt.Println(twoD)
	// res := groupAnagrams(strs)
	fmt.Println(strs)
}

func groupAnagrams(strs []string) map[string][]string {

	res := make(map[string][]string)

	for k, v := range strs {
		if v == "" {
			continue
		}
		for i := k + 1; i < len(strs); i++ {
			if strs[i] == "" {
				continue
			}

			if isAnagram(v, strs[i]) {
				res[v] = append(res[v], strs[i])
				strs[i] = ""
			}
		}
	}
	fmt.Println(strs)
	// arrRes := make([][]string, len(strs))

	// for k, v := range strs {
	// 	arrRes[k] = append(arrRes[k], res[v])
	// }

	return res
}

func isAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	arr := make([]int, 26)

	for i := 0; i < len(b); i++ {
		arr[a[i]-'a']++
		arr[b[i]-'a']--
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}
