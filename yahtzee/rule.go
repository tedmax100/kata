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
	return SumRecursive(r.CollectTarget(dices, targetNum))
}

func (r *Rule) Twos(dices []int) int {
	var targetNum int = 2
	return SumRecursive(r.CollectTarget(dices, targetNum))
}

func (r *Rule) Threes(dices []int) int {
	var targetNum int = 3
	return SumRecursive(r.CollectTarget(dices, targetNum))
}

func (r *Rule) Fours(dices []int) int {
	var targetNum int = 4
	return SumRecursive(r.CollectTarget(dices, targetNum))
}

func (r *Rule) Fives(dices []int) int {
	var targetNum int = 5
	return SumRecursive(r.CollectTarget(dices, targetNum))
}

func (r *Rule) Sixes(dices []int) int {
	var targetNum int = 6
	return SumRecursive(r.CollectTarget(dices, targetNum))
}

func (r *Rule) CollectTarget(nums []int, targetNum int) []int {
	var collected []int = make([]int, 0)
	for idx := range nums {
		if nums[idx] == targetNum {
			collected = append(collected, nums[idx])
		}
	}
	return collected
}
