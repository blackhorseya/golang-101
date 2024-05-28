package strategy

// SortStrategy 定義排序算法的接口
type SortStrategy interface {
	Sort([]int) []int
}

// BubbleSortStrategy 是冒泡排序算法的具體策略
type BubbleSortStrategy struct{}

func (b *BubbleSortStrategy) Sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// QuickSortStrategy 是快速排序算法的具體策略
type QuickSortStrategy struct{}

func (q *QuickSortStrategy) Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := len(arr) / 2
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	q.Sort(arr[:left])
	q.Sort(arr[left+1:])
	return arr
}

// SortContext 是上下文，持有一個 SortStrategy 的引用
type SortContext struct {
	strategy SortStrategy
}

func NewSortContext(strategy SortStrategy) *SortContext {
	return &SortContext{strategy: strategy}
}

func (c *SortContext) SetStrategy(strategy SortStrategy) {
	c.strategy = strategy
}

func (c *SortContext) Sort(arr []int) []int {
	return c.strategy.Sort(arr)
}
