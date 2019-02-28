package canFinish

func dfs(from int, g map[int][]int, visiting map[int]bool) bool {
	visiting[from] = true

	for _, v := range g[from] {
		if _, ok := visiting[v]; ok || !dfs(v, g, visiting) {
			return false
		}
	}

	delete(visiting, from)
	delete(g, from)

	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := map[int][]int{}
	visiting := map[int]bool{}

	// build graph
	for _, p := range prerequisites {
		g[p[0]] = append(g[p[0]], p[1])
	}

	// we are going to find cycle, so then starting point
	// is not important at all.
	// in fact, it is still not important even if we
	// want to get the topological sort order, see:
	// https://www.geeksforgeeks.org/topological-sorting/
	for i := 0; i < numCourses; i++ {
		if !dfs(i, g, visiting) {
			return false
		}
	}

	return true
}
