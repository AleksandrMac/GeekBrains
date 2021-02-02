package lesson4

import "fmt"

//Point -
type Point struct {
	x, y int
}

//Steps -
type Steps struct {
	number       int32
	prev         *Steps
	point        Point
	historyPoint []Point
}

var notPoint Point = Point{
	x: -1,
	y: -1,
}

//Horse -
func Horse(m int) [][]int32 {

	var slice [][]int32
	slice = make([][]int32, m, m)
	for i := 0; i < m; i++ {
		slice[i] = make([]int32, m)
	}
	steps := Steps{
		number: 1,
		prev:   nil,
		point: Point{
			x: 0,
			y: 0,
		},
		historyPoint: []Point{},
	}
	slice[steps.point.y][steps.point.x] = steps.number
	for {
		step := makeStep(steps, slice)
		if !equal(step, notPoint) {
			slice[step.y][step.x] = steps.number + 1
			steps.historyPoint = append(steps.historyPoint, step)
			s := steps
			steps = Steps{
				number:       steps.number + 1,
				prev:         &s,
				point:        step,
				historyPoint: []Point{},
			}
		} else if steps.number == int32(m*m) {
			return slice
		} else if steps.prev == nil {
			return [][]int32{
				{-1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1},
			}
		} else {
			slice[steps.point.y][steps.point.x] = 0
			steps = *steps.prev
		}
		fmt.Println()
		for i := 0; i < len(slice); i++ {
			fmt.Println(slice[i])
		}
	} /*

		for i := 0; i < m; i++ {
			fmt.Println(slice[i])
		}*/
	//return slice
}

func makeStep(st Steps, slice [][]int32) Point {

	step := stepRU(st.point, slice)
	if !equal(step, notPoint) && !contains(st.historyPoint, step) {
		return step
	}
	step = stepRD(st.point, slice)
	if !equal(step, notPoint) && !contains(st.historyPoint, step) {
		return step
	}
	step = stepDR(st.point, slice)
	if !equal(step, notPoint) && !contains(st.historyPoint, step) {
		return step
	}
	step = stepDL(st.point, slice)
	if !equal(step, notPoint) && !contains(st.historyPoint, step) {
		return step
	}
	step = stepLD(st.point, slice)
	if !equal(step, notPoint) && !contains(st.historyPoint, step) {
		return step
	}
	step = stepLU(st.point, slice)
	if !equal(step, notPoint) && !contains(st.historyPoint, step) {
		return step
	}
	step = stepUL(st.point, slice)
	if !equal(step, notPoint) && !contains(st.historyPoint, step) {
		return step
	}
	step = stepUR(st.point, slice)
	if !equal(step, notPoint) && !contains(st.historyPoint, step) {
		return step
	}
	return notPoint
}

//equal - возвращает равны ли точки
func equal(p1 Point, p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

//contains - возвращает true, в случае если points содержит point, иначе false
func contains(points []Point, point Point) bool {
	for _, p := range points {
		if equal(p, point) {
			return true
		}
	}
	return false
}

//step to 2right and 1up
func stepRU(position Point, slice [][]int32) Point {
	if position.x+2 < len(slice) &&
		position.y-1 >= 0 &&
		slice[position.y-1][position.x+2] == 0 {
		return Point{
			x: position.x + 2,
			y: position.y - 1,
		}
	}
	return notPoint
}

//step to 2right and 1down
func stepRD(position Point, slice [][]int32) Point {
	if position.x+2 < len(slice) &&
		position.y+1 < len(slice[position.y]) &&
		slice[position.y+1][position.x+2] == 0 {
		return Point{
			x: position.x + 2,
			y: position.y + 1,
		}
	}
	return notPoint
}

//step to 2down and 1right
func stepDR(position Point, slice [][]int32) Point {
	if position.x+1 < len(slice) &&
		position.y+2 < len(slice[position.y]) &&
		slice[position.y+2][position.x+1] == 0 {
		return Point{
			x: position.x + 1,
			y: position.y + 2,
		}
	}
	return notPoint
}

//step to 2down and 1left
func stepDL(position Point, slice [][]int32) Point {
	if position.x-1 >= 0 &&
		position.y+2 < len(slice[position.y]) &&
		slice[position.y+2][position.x-1] == 0 {
		return Point{
			x: position.x - 1,
			y: position.y + 2,
		}
	}
	return notPoint
}

//step to 2left and 1down
func stepLD(position Point, slice [][]int32) Point {
	if position.x-2 >= 0 &&
		position.y+1 < len(slice[position.y]) &&
		slice[position.y+1][position.x-2] == 0 {
		return Point{
			x: position.x - 2,
			y: position.y + 1,
		}
	}
	return notPoint
}

//step to 2left and 1up
func stepLU(position Point, slice [][]int32) Point {
	if position.x-2 >= 0 &&
		position.y-1 >= 0 &&
		slice[position.y-1][position.x-2] == 0 {
		return Point{
			x: position.x - 2,
			y: position.y - 1,
		}
	}
	return notPoint
}

//step to 2up and 1left
func stepUL(position Point, slice [][]int32) Point {
	if position.x-1 >= 0 &&
		position.y-2 >= 0 &&
		slice[position.y-2][position.x-1] == 0 {
		return Point{
			x: position.x - 1,
			y: position.y - 2,
		}
	}
	return notPoint
}

//step to 2up and 1right
func stepUR(position Point, slice [][]int32) Point {
	if position.x+1 < len(slice) &&
		position.y-2 >= 0 &&
		slice[position.y-2][position.x+1] == 0 {
		return Point{
			x: position.x + 1,
			y: position.y - 2,
		}
	}
	return notPoint
}
