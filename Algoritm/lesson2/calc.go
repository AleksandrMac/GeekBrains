package lesson2

//CalcPerformer -
func CalcPerformer(start int8, end int8, root bool) (count int64) {
	if start == end {
		//count++
		return
	}
	if start > end {
		//count++
		return
	}
	if start*2 < end {
		if root {
			count++
		}
		count++
		count += CalcPerformer(start+1, end, false)
		count += CalcPerformer(start*2, end, false)
	}
	if start*2 > end {
		if root {
			count++
		}
		count += CalcPerformer(start+1, end, false) //10
	}
	if start*2 == end {
		count += 2 //4
	}
	return
}
