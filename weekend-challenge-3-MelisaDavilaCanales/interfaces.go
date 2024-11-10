package main

type Food interface {
	SatisfyHunger()
	BoostStamina()
	ReduceSleepiness()
}

type Human interface {
	ShowStats() string
	Sleep() string
	Study() string
	Exercise()
	Eat(Food)
}
