package main

import (
	"flag"
	"fmt"
)

// StrFind searches the pattern from the text, and if there is any match, return
// the index of the first match. return -1 if no match found.
func StrFind(text string, pattern string) int {
	// short-circuit some cases
	if len(text) < len(pattern) {
		return -1
	}

	table := TableFor(pattern)

	for i := 0; i < len(text); i++ {
		for j := 0; j < len(table); j++ {
			// no possible match
			if i+j >= len(text) {
				return -1
			}

			if trace {
				fmt.Printf("compare patter[%d] vs text[%d]\n", j, i+j)
			}
			match := pattern[j] == text[i+j]
			// exact match
			if j+1 == len(table) && match {
				return i
			}
			if match {
				continue
			}

			// pattern[j]!=text[i+j]
			// and we know that
			//          pattern[0]=pattern[j-table[j]]
			//             ...    =      ...
			// pattern[table[j]-1]=pattern[j-1]
			// which effectively mean:
			//             text[i]=pattern[j-table[j]]
			// so we now move:
			//                   i=i+j-table[j]
			i += (j - table[j]) - 1
			break
		}
	}
	return -1
}

// TableFor returns an failure value array for the pattern to move forward.
func TableFor(pattern string) []int {
	table := make([]int, len(pattern))

	for i := range pattern {
		if i == 0 || i == 1 {
			table[i] = i - 1
			continue
		}

		// define
		//       p[0]=p[j-k]
		//       p[1]=p[j-k+1]
		//        ..    ..
		//     p[k-1]=p[j-1]
		// as
		//       f(j)=k
		// now:
		//        p[k]=p[j] <==> f(j+1)=k+1
		//   p[f(j-1)]=p[j] <==> f(j+1)=f(j-1)+1
		//         ...      <==>     ...
		if pattern[table[i-1]] == pattern[i-1] {
			table[i] = table[i-1] + 1
		} else {
			// now look backward
			backward := table[i-1]
			for backward >= 0 {
				if pattern[backward] == pattern[i-1] {
					break
				} else {
					backward = table[backward]
				}
			}
			table[i] = backward + 1
		}
	}

	return table
}

var pattern string
var text string
var trace bool

func init() {
	flag.StringVar(&pattern, "pattern", "", "the pattern to search")
	flag.StringVar(&text, "text", "", "the text where to search")
	flag.BoolVar(&trace, "trace", false, "if set to true, print the detailed comparison trace log")
}

func main() {
	flag.Parse()

	fmt.Println(StrFind(text, pattern))
}
