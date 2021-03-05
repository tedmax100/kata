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
	pairs = NumberCollactor(dices)

	if r.hasPairs(dices, pairs) {
		for diceNum := range pairs {
			if pairs[diceNum] > 1 && diceNum > maxNum {
				maxNum = diceNum
			}
		}

		return maxNum * 2
	}
	return 0
}

func (r *Rule) TwoPairs(dices []int) int {
	var pairs map[int]int = make(map[int]int)
	var result int = 0
	pairs = NumberCollactor(dices)
	var pairCount int = 0
	if r.hasPairs(dices, pairs) {
		for diceNum := range pairs {
			if pairs[diceNum] > 1 {
				pairCount++
				result += 2 * diceNum
			}

		}
		if pairCount == 2 {
			return result
		}
	}
	return 0
}

func (r *Rule) ThreeOfAKind(dices []int) int {
	var pairs map[int]int = make(map[int]int)
	pairs = NumberCollactor(dices)

	if r.hasPairs(dices, pairs) {
		for diceNum := range pairs {
			if pairs[diceNum] > 2 {
				return diceNum * 3
			}
		}
	}
	return 0
}

func (r *Rule) FourOfAKind(dices []int) int {
	var pairs map[int]int = make(map[int]int)
	var number int = 0
	pairs = NumberCollactor(dices)

	if r.hasPairs(dices, pairs) {
		for diceNum := range pairs {
			if pairs[diceNum] > 3 {
				return diceNum * 4
			}
		}
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
	// 將dice做點數上的分堆
	pairs = NumberCollactor(dices)

	if r.hasPairs(dices, pairs) {
		for diceNum := range pairs {
			if pairs[diceNum] == 2 {
				onePairNum = diceNum
			}
			if pairs[diceNum] == 3 {
				threOfKindNum = diceNum
			}
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

// NumberCollactor : 將numbers做數字上的分堆整理
func NumberCollactor(numbers []int) map[int]int {
	var pairs map[int]int = make(map[int]int)

	for idx := range numbers {
		if _, exist := pairs[numbers[idx]]; exist {
			pairs[numbers[idx]]++
		} else {
			pairs[numbers[idx]] = 1
		}
	}

	return pairs
}

// hasPairs : check length of dices is large than length of pairs
func (r *Rule) hasPairs(dices []int, pairs map[int]int) bool {
	return len(dices) > len(pairs)
}
