package yahtzee

type Rule struct {
}

func (r *Rule) Chance(dices []int) int {
	return SumRecursive(dices)
}

func SumRecursive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + SumRecursive(nums[1:])
}

func (r *Rule) Once(dices []int) int {
	var targetNum int = 1
	return r.CountTarget(dices, targetNum) * targetNum
}

func (r *Rule) Twos(dices []int) int {
	var targetNum int = 2
	return r.CountTarget(dices, targetNum) * targetNum
}

func (r *Rule) Threes(dices []int) int {
	var targetNum int = 3
	return r.CountTarget(dices, targetNum) * targetNum
}

func (r *Rule) CountTarget(nums []int, targetNum int) int {
	var count int = 0
	for idx := range nums {
		if nums[idx] == targetNum {
			count++
		}
	}
	return count
}
