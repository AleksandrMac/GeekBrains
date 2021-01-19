package lesson2

//CalcPerformer -
func CalcPerformer(start int8, end int8) (count int64) {

	if start*2 < end {
		count++
		count += CalcPerformer(start+1, end)
		count += CalcPerformer(start*2, end)
	} else {
		count++
		return
	}
	return
}
