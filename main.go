package main

func main() {
	// puzzle := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"
	puzzle := "..53.....8......2..7..1.5..4....53...1..7...6..32...8..6.5....9..4....3......97.."
	output := Solve(puzzle) // ParseGrid(puzzle)
	Display(output)
}
