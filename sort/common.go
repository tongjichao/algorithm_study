package mysort

func less(i, j int64) bool {

	if i < j {
		return true
	}

	return false
}

func exch(a *[]int64, i, j int64) {
	tmp := (*a)[i]
	(*a)[i] = (*a)[j]
	(*a)[j] = tmp
}
