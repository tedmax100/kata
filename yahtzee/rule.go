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
	var count int = 0
	for idx := range dices {
		if dices[idx] == 1 {
			count++
		}
	}
	return count * 1
}

func (r *Rule) Twos(dices []int) int {
	var count int = 0
	for idx := range dices {
		if dices[idx] == 2 {
			count++
		}
	}
	return count * 2
}
