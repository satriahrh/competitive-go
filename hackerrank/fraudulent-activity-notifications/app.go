package fraudulent_activity_notifications

func ActivityNotifications(expenditure []int32, d int32) (numberOfNotice int32) {
	dictionary := make(map[int32]int32)

	find := func(id int32) int32 {
		sum := int32(0)
		for i := int32(0); i <= 200; i++ {
			sum += dictionary[i]
			if sum >= id {
				return i
			}
		}
		return 0
	}

	for i := int32(0); i < int32(len(expenditure)); i++ {
		current := expenditure[i]
		dictionary[current] += 1
		if i >= d {
			median := find(d/2 + d%2)
			if d%2 == 0 {
				median += find(d/2 + 1)
			} else {
				median *= 2
			}

			if current >= median {
				numberOfNotice += 1
			}

			toBeRemoved := expenditure[int32(i)-d]
			dictionary[toBeRemoved] -= 1
		}
	}

	return numberOfNotice
}
