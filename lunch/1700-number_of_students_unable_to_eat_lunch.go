package lunch

func countStudents(students []int, sandwiches []int) int {
	for _, sandwich := range sandwiches {
		for i, studentPref := range students {
			if studentPref == sandwich {
				if i == 0 {
					if len(students) > 1 {
						students = students[1:]
					} else {
						students = make([]int, 0)
					}
				} else if i == len(students)-1 {
					students = students[:len(students)-1]
				} else {
					students = append(students[i+1:], students[:i]...)
				}
				break
			} else {
				if i == len(students)-1 {
					return len(students)
				}
			}
		}
	}
	return 0
}
