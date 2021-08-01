package main

type Spell func(int, int) bool

type spells []Spell

func newSpells() spells {
	spells := make(spells, 0)
	spells = append(spells, s1)
	spells = append(spells, s2)
	spells = append(spells, s3)
	spells = append(spells, s4)
	spells = append(spells, s5)
	spells = append(spells, s6)
	spells = append(spells, s7)
	spells = append(spells, s8)
	spells = append(spells, s9)
	spells = append(spells, s10)
	spells = append(spells, s11)
	spells = append(spells, s20)
	spells = append(spells, s23)
	spells = append(spells, s24)
	spells = append(spells, s25)
	return spells
}

func s1(x, y int) bool {
	return x <= y
}

func s2(x, y int) bool {
	return x == y
}

func s3(x, y int) bool {
	return x == 24-y
}

func s4(x, y int) bool {
	return x <= 24+5-y
}

func s5(x, y int) bool {
	return x == y/2
}

func s6(x, y int) bool {
	return x < 10 || y < 10
}

func s7(x, y int) bool {
	return x > 16 && y > 16
}

func s8(x, y int) bool {
	return x == 0 || y == 0
}

func s9(x, y int) bool {
	return x >= 11+y || x <= y-11
}

func s10(x, y int) bool {
	return !(x >= y || x <= y/2-1)
}

func s11(x, y int) bool {
	return x == 1 || x == 23 || y == 1 || y == 23
}

func s20(x, y int) bool {
	return x%2-y%2 == 0
}

func s23(x, y int) bool {
	return x%3 == 0 && y%2 == 0
}

func s24(x, y int) bool {
	return x == y || x == 24-y
}

func s25(x, y int) bool {
	return x%6 == 0 || y%6 == 0
}
