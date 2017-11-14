/*
	二分查找
*/
package search

func Rank(key int, a []int) int {
	lo := int(0)
	hi := int(len(a) - 1)
	for {
		if lo > hi {
			break
		}

		mid := lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
			continue
		}

		if key > a[mid] {
			lo = mid + 1
			continue
		}

		return mid
	}

	return -1
}
