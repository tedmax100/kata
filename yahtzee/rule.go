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

func (r *Rule) OnePair(dices []int) int {
	var pairs map[int]int = make(map[int]int)
	var maxNum int = 0
	for idx := range dices {
		if _, exist := pairs[dices[idx]]; exist {
			pairs[dices[idx]]++
			if dices[idx] > maxNum {
				maxNum = dices[idx]
			}
		} else {
			pairs[dices[idx]] = 1
		}
	}

	if maxNum == 0 {
		return 0
	}
	return maxNum * 2
}

func (r *Rule) TwoPairs(dices []int) int {
	var pairs map[int]int = make(map[int]int)
	var result int = 0
	for idx := range dices {
		if _, exist := pairs[dices[idx]]; exist {
			pairs[dices[idx]]++
		} else {
			pairs[dices[idx]] = 1
		}
	}

	// remove item of pairs that value = 1
	for number := range pairs {
		if pairs[number] == 1 {
			delete(pairs, number)
		}
	}

	if len(pairs) < 2 {
		return 0
	}

	for pair := range pairs {
		result += 2 * pair
	}
	return result
}

func (r *Rule) ThreeOfAKind(dices []int) int {
	var pairs map[int]int = make(map[int]int)
	var number int = 0
	for idx := range dices {
		if _, exist := pairs[dices[idx]]; exist {
			pairs[dices[idx]]++
		} else {
			pairs[dices[idx]] = 1
		}
	}

	// remove item of pairs that value = 1
	for number := range pairs {
		if pairs[number] < 3 {
			delete(pairs, number)
		}
	}

	if len(pairs) < 1 {
		return 0
	}

	for pair := range pairs {
		number = pair
	}
	return number * 3
}

func (r *Rule) FourOfAKind(dices []int) int {
	var pairs map[int]int = make(map[int]int)
	var number int = 0
	for idx := range dices {
		if _, exist := pairs[dices[idx]]; exist {
			pairs[dices[idx]]++
		} else {
			pairs[dices[idx]] = 1
		}
	}

	// remove item of pairs that value = 1
	for number := range pairs {
		if pairs[number] < 4 {
			delete(pairs, number)
		}
	}

	if len(pairs) < 1 {
		return 0
	}

	for pair := range pairs {
		number = pair
	}
	return number * 4
}
