package day13

type Vector struct {
	X int
	Y int
}

type Machine struct {
	ButtonA Vector
	ButtonB Vector
	Prize   Vector
}

func CountTokens(machines []Machine, maxPresses int) int {
	total := 0
	for _, machine := range machines {
		total += countMachineTokens(machine, maxPresses)
	}
	return total
}

/*
*
Let a = # button a presses
Let b = # button b presses

We have:
- a * ButtonA.X + b * ButtonB.X = Prize.X
- a * ButtonA.Y + b * ButtonB.Y = Prize.Y

Multiply first equation by ButtonA.Y
Multiply second equation by ButtonA.X
Subtract second from first to get:

b * (ButtonB.X * ButtonA.Y - ButtonA.X * ButtonB.Y) = Prize.X * ButtonA.Y - Prize.Y * ButtonA.X
=> b = (Prize.X * ButtonA.Y - Prize.Y * ButtonA.X) / (ButtonB.X * ButtonA.Y - ButtonA.X * ButtonB.Y)

Likewise, we can get
a = (Prize.X * ButtonB.Y - Prize.Y * ButtonB.X) / (ButtonA.X * ButtonB.Y - ButtonA.Y * ButtonB.X)

	= (Prize.Y * ButtonB.X - Prize.X * ButtonB.Y) / (ButtonB.X * ButtonA.Y - ButtonA.X * ButtonB.Y)

It's only not possible to get any solution if button A and B are collinear
We also throw out the solution if a or b are not integers in [0, 100]
*
*/
func countMachineTokens(machine Machine, maxPresses int) int {
	determinant := machine.ButtonB.X*machine.ButtonA.Y - machine.ButtonA.X*machine.ButtonB.Y
	if determinant == 0 {
		return 0
	}
	aNumerator := machine.Prize.Y*machine.ButtonB.X - machine.Prize.X*machine.ButtonB.Y
	if aNumerator%determinant != 0 {
		return 0
	}
	bNumerator := machine.Prize.X*machine.ButtonA.Y - machine.Prize.Y*machine.ButtonA.X
	if bNumerator%determinant != 0 {
		return 0
	}
	a := int(float64(aNumerator) / float64(determinant))
	b := int(float64(bNumerator) / float64(determinant))
	if a < 0 || b < 0 {
		return 0
	}
	if maxPresses > 0 && (a > maxPresses || b > maxPresses) {
		return 0
	}
	return 3*a + b
}
