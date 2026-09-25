package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

import "github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"

func main() {

	// The string we want to transform (prepend your seed/FLAG here)
	s := "L}CJFOLRc6CJFL3c1LnCJFOLmc6CJFL5c1LFCJFOLfc6CJFL7c1LqCJFOLkc6CJFL9c1LICJFOLLc6CJFLbc1LMCJFOLmc6CJFLdc1LdCJFOLzc6CJFLfc1L8CJFOLyc6CJFLhc1LLCJFOLSc6CJFLjc1LzCJFOLoc6CJFLlc1LkCJFOL5c6CJFLnc1LFCJFOLyc6CJFLpc1LbCJFOLjc6CJFLrc1LhCJFOLAc6CJFLtc1LzCJFOLsc6CJFLvc1LRCJFOLec6CJFLxc1LuCJFOLhc6CJFLzc1LjCJFOL8c6CJFLBc1LrCJFLDc1LQCJFL{CJFLGCJFLACJFLLCJFLFCJFS#f1L0sdcL0CjmrUsP3Asrsrf4DMsHDMsHDMsHDMsHDMsHDMsHrf5L2c1AL2c1AL2c1AL2c1AL2c1Af6c4Y7c4Y7AMpbY6AMp9Y5AMp7Y4AMp5Y3AMp3Y2AMY2Y3Y4Y5Y6c5"

	// Open the .htaccess file
	file, err := os.Open("htstaccess")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read lines into a slice
	var rules []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// Only keep lines starting with RewriteRule
		if strings.HasPrefix(line, "RewriteRule") {
			rules = append(rules, line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	changed := true
	for changed {
		changed = false

		// Iterate over all RewriteRule lines
		for r, rule := range rules {
			// Example line: RewriteRule "^L(.)(.*)$" "$2$1" [N]
			parts := strings.Split(rule, " ")
			if len(parts) < 3 {
				continue
			}
			pattern := parts[1]
			subst := parts[2]

			// Remove quotes around pattern/subst
			pattern = strings.Trim(pattern, `"`)
			subst = strings.Trim(subst, `"`)

			re, err := pcre.Compile(pattern, 0)
			if err != nil {
				fmt.Println("Invalid regex:", pattern)
				continue
			}

			matcher := re.MatcherString(s, 0)

			if matcher.Matches() {

				if r == 267 {
					s = s[1:]
					changed = true
				} else {

					reGroup := regexp.MustCompile(`\$(\d+)`)
					newS := reGroup.ReplaceAllStringFunc(subst, func(m string) string {
						numStr := m[1:]                     // e.g., "10" from "$10"
						n, _ := strconv.Atoi(numStr)        // parse number
						if n == 0 {                         // $0 is usually whole match
							return matcher.GroupString(0)
						}
						if n <= matcher.Groups() {
							// If group exists, return its value
							return matcher.GroupString(n)
						}
						// Group does NOT exist → append the "extra digits" to $1
						// e.g., "$10" when only 1 group → $1 + "0"
						if matcher.Groups() >= 1 {
							return matcher.GroupString(1) + numStr[1:] // combine first digit + remaining suffix inline
						}
						// fallback: no group exists → leave literal
						return m
					})



					// Stuck in loop for whatever reason
					// newS := subst
					// // for i := 1; i <= matcher.Groups(); i++ {
					// for i := matcher.Groups(); i >= 1; i-- {
					// 	newS = strings.ReplaceAll(newS, fmt.Sprintf("$%d", i), matcher.GroupString(i))
					// }

					fmt.Println("pattern #" + strconv.Itoa(r) + ": " + rule)
					fmt.Println("string #" + strconv.Itoa(r) + ": " + s)

					if s != newS {
						changed = true
						s = newS
						break // go back to the top for [N] behavior
					} else {
						break
					}

				}

			}
		}
	}

	fmt.Println("Final string:", s)
}
