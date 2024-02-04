package main

func BubbleSort(arr []int) {
	var flag bool
	flag = false
	for i := 0; i < len(arr); i++ {
		for k := 0; k < len(arr); k++ {
			if arr[i] > arr[k] {
				arr[i], arr[k] = arr[k], arr[i]
				flag = true
			}
			if !flag {
				break
			}
		}

	}
}
