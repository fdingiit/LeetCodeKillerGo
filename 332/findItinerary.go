package findItinerary

import (
	"sort"
)

func findItinerary(tickets [][]string) []string {
	var stack, res []string
	var routines map[string][]string

	routines = map[string][]string{}
	for _, t := range tickets {
		from, to := t[0], t[1]
		routines[from] = append(routines[from], to)
	}

	// sorted by lexical
	for from, toList := range routines {
		sort.Strings(toList)
		routines[from] = toList
	}

	stack = append(stack, "JFK")
	for len(stack) > 0 {
		for {
			from := stack[len(stack)-1]
			toList, ok := routines[from]

			if ok && len(toList) > 0 {
				stack = append(stack, toList[0])
				routines[from] = toList[1:]
			} else {
				res = append(res, from)
				stack = stack[:len(stack)-1]
				break
			}
		}
	}

	// reverse
	for i := len(res)/2 - 1; i >= 0; i-- {
		opp := len(res) - 1 - i
		res[i], res[opp] = res[opp], res[i]
	}

	return res
}
