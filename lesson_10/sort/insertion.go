package sort

func InsertionSort(ar []int64) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}
