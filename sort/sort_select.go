package mysort

//选择排序
func SelectSort(a *[]int64) {
	N := int64(len(*a))
	for i := int64(0); i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if less((*a)[j], (*a)[min]) {
				min = j
			}
		}
		exch(a, i, min)
	}
}
