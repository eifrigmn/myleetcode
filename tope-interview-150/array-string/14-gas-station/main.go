package main

func canCompleteCircuit(gas []int, cost []int) int {
	var size = len(gas)
	for i := 0; i < size; i++ {
		var gasLeft = gas[i]
		j := i
		for gasLeft-cost[j] >= 0 {
			gasLeft = gasLeft - cost[j] + gas[(j+1)%size]
			j = (j + 1) % size
			if j == i {
				return i
			}
		}
		if j < i {
			return -1
		}
		i = j

	}
	return -1
}
