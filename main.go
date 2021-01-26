package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello go for leetcode")
	// [[2,1],[5,4],[3,7],[6,2],[4,4],[1,8],[9,6],[5,3],[7,4],[1,9],[1,1],[6,6],[9,6],[1,3],[9,7],[4,7],[5,1],[6,5],[1,6],[6,1],[1,8],[7,2],[2,4],[1,6],[3,1],[3,9],[3,7],[9,1],[1,9],[8,9]]
	// l := [][]int{{2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {1, 1}, {2, 2}}
	l := [][]int{{2, 1}, {5, 4}, {3, 7}, {6, 2}, {4, 4}, {1, 8}, {9, 6}, {5, 3}, {7, 4}, {1, 9}, {1, 1}, {6, 6}, {9, 6}, {1, 3}, {9, 7}, {4, 7}, {5, 1}, {6, 5}, {1, 6}, {6, 1}, {1, 8}, {7, 2}, {2, 4}, {1, 6}, {3, 1}, {3, 9}, {3, 7}, {9, 1}, {1, 9}, {8, 9}}

	numEquivDominoPairs(l)
}

func numEquivDominoPairs(dominoes [][]int) (ans int) {
	if len(dominoes) < 2 {
		return
	}
	for _, v := range dominoes {
		if v[0] > v[1] {
			v[0], v[1] = v[1], v[0]
		}
	}
	fmt.Println(dominoes)
	sort.Slice(dominoes, func(i, j int) bool {
		if dominoes[i][0] == dominoes[j][0] {
			return dominoes[i][1] < dominoes[j][1]
		}
		return dominoes[i][0] < dominoes[j][0]
	})
	fmt.Println(dominoes)
	for i, v := range dominoes {
		for j := i + 1; j < len(dominoes); j++ {
			if dominoes[j][0] == v[0] && dominoes[j][1] == v[1] {
				ans++
			} else {
				break
			}
		}
	}
	return
}
