package sort

func QuickSort(R []int, low, high int) {
	if low >= high {
		return
	}

	i, j, t := low, high, R[low]

	for i != j {
		for i < j && t <= R[j] {
			j--
		}

		if i < j {
			R[i] = R[j]
			i++
		}

		for i < j && R[i] <= t {
			i++
		}

		if i < j {
			R[j] = R[i]
			j--
		}
	}

	R[i] = t

	QuickSort(R, low, i-1)
	QuickSort(R, i+1, high)
}

func MergeSort(R []int) {
	temp := make([]int, len(R))

	sort(R, temp, 0, len(temp)-1)
}

func sort(a, t []int, l, r int) {
	if l < r {
		mid := (l + r) >> 1
		sort(a, t, l, mid)     //左边归并排序，使得左子序列有序
		sort(a, t, mid+1, r)   //右边归并排序，使得右子序列有序
		merge(a, t, l, r, mid) //将两个有序子数组合并操作
	}
}

func merge(a, temp []int, l, r, mid int) {
	i := l
	j := mid + 1 //右序列指针
	t := 0       //临时数组指针

	for i <= mid && j <= r {
		if a[i] <= a[j] {
			temp[t] = a[i]
			t++
			i++
		} else {
			temp[t] = a[j]
			t++
			j++
		}
	}
	for i <= mid { //将左边剩余元素填充进temp中
		temp[t] = a[i]
		t++
		i++
	}
	for j <= r { //将右序列剩余元素填充进temp中
		temp[t] = a[j]
		t++
		j++
	}
	t = 0
	//将temp中的元素全部拷贝到原数组中
	for l <= r {
		a[l] = temp[t]
		l++
		t++
	}
}
