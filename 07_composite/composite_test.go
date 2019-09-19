package composite

import "testing"

func TestCompositePattern(t *testing.T) {
	shark := Shark{
		&Animal{},
		&SwimmerImpl{},
	}
	shark.Eat()
	shark.Swim()


	phelps := Phelps{
		&TrainerImpl{},
		&SwimmerImpl{},
	}
	phelps.Train()
	phelps.Swim()
}
