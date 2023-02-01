package main

import (
	"fmt"
	sort2 "sort"
	"time"
)

var arr1 = []int{9, 16, 3, 12, 8, 4, 3, 2, 9, 5}
var arr2 = []int{9, 16, 3, 12, 8, 4, 3, 2, 9, 5}

func main() {
	start := time.Now()
	sort2.Ints(arr1)
	elapsed := time.Since(start)
	fmt.Printf("Standart Go 'Qucik Sort' based algo took %s\n", elapsed)

	start = time.Now()
	heapSort(arr2)
	elapsed = time.Since(start)
	fmt.Printf("Heap Sort took %s", elapsed*1000)
	fmt.Println(arr2)

}
func heapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		// Перемещаем текущий корень в конец
		temp := arr[0]
		arr[0] = arr[i]
		arr[i] = temp

		// Вызываем процедуру heapify на уменьшенной куче
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {

	largest := i // Инициализируем наибольший элемент как корень
	l := 2*i + 1 // левый = 2*i + 1
	r := 2*i + 2 // правый = 2*i + 2

	// Если левый дочерний элемент больше корня
	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	// Если правый дочерний элемент больше, чем самый большой элемент на данный момент
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	// Если самый большой элемент не корень
	if largest != i {
		swap := arr[i]
		arr[i] = arr[largest]
		arr[largest] = swap

		// Рекурсивно преобразуем в двоичную кучу затронутое поддерево
		heapify(arr, n, largest)
	}
}
