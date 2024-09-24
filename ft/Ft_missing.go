package TPWOUHAIBIMAHDI

func Ft_missing(nums []int) int {
	n := len(nums)
	expected_sum := n * (n + 1) / 2
	actual_sum := 0
	for _, num := range nums {
		actual_sum += num
	}
	return expected_sum - actual_sum
}
