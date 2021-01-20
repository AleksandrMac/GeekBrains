package lesson2

//CalcPerformer -
func CalcPerformer(start int8, end int8, root bool) (count int64) {
	if start == end {
		return
	}
	if root {
		count++
	}
	if start*2 <= end {
		count++
		count += CalcPerformer(start*2, end, false)
	}
	count += CalcPerformer(start+1, end, false)
	return
}
