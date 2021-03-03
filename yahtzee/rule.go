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

func (r *Rule) SmallStraight(dices []int) int {
	var set map[int]int = make(map[int]int)
	for idx := range dices {
		set[dices[idx]] = 1
	}
	if len(set) < 5 {
		return 0
	}
	if set[1] == 0 || set[5] == 0 {
		return 0
	}
	return 15
}

func (r *Rule) LargeStraight(dices []int) int {
	var set map[int]int = make(map[int]int)
	for idx := range dices {
		set[dices[idx]] = 1
	}
	if len(set) < 5 {
		return 0
	}
	if set[2] == 0 || set[6] == 0 {
		return 0
	}
	return 20
}

func (r *Rule) FullHouse(dices []int) int {
	// 找出 ThreeOfAKind
	// 再把剩下的元素, 去找OnePair
	var pairs map[int]int = make(map[int]int)
	var threOfKindNum, onePairNum int
	for idx := range dices {
		if _, exist := pairs[dices[idx]]; exist {
			pairs[dices[idx]]++
		} else {
			pairs[dices[idx]] = 1
		}
	}

	if len(pairs) < 2 {
		return 0
	}

	for pair := range pairs {
		if pairs[pair] == 2 {
			onePairNum = pair
		}
		if pairs[pair] == 3 {
			threOfKindNum = pair
		}
	}
	if threOfKindNum > 0 && onePairNum > 0 {
		return threOfKindNum*3 + onePairNum*2
	}
	return 0
}

func (r *Rule) Yahtzee(dices []int) int {
	var preDice int = dices[0]
	for idx := range dices {
		if preDice == dices[idx] {
			continue
		}
		return 0
	}

	return 50
}
