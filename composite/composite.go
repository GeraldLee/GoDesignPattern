package composite

import "fmt"

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type Eater interface {
	Eat()
}

type Athlete struct{}
func (a *Athlete) Train() {
	fmt.Println("Training")
}

type Animal struct{}
func (r *Animal)Eat() {
	println("Eating")
}

type SwimmerImpl struct{}
func (s *SwimmerImpl) Swim(){
	println("Swimming!")
}

type TrainerImpl struct{}
func (s *TrainerImpl) Train(){
	println("Training!")
}

type Shark struct{
	Eater
	Swimmer
}

type Phelps struct{
	Trainer
	Swimmer
}