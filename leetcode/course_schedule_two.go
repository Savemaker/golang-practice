package leetcode

const (
	COURSE       = 0
	PREREQUISITE = 1
)

func FindOrder(numCourses int, prerequisites [][]int) []int {
	graph := CreateGraph(prerequisites, numCourses)
	result := []int{}
	for {
		graphUpdated := false
		for k, v := range graph {
			if len(v) == 0 {
				delete(graph, k)
				graphUpdated = true
				result = append(result, k)
				for k1, v1 := range graph {
					for k2 := range v1 {
						if v1[k2] == k {
							v1[k2] = v1[len(v1)-1]
							graph[k1] = v1[:len(v1)-1]
						}
					}
				}
			}
		}
		if !graphUpdated {
			return []int{}
		}
		if len(result) == numCourses {
			return result
		}
	}
}

func CreateGraph(prerequisites [][]int, numCourses int) map[int][]int {
	graph := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}
	for _, v := range prerequisites {
		graph[v[COURSE]] = append(graph[v[COURSE]], v[PREREQUISITE])
	}
	return graph
}
