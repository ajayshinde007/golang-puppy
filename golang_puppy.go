package golang_puppy

import "github.com/ajayshinde007/golang_dog"

func Bark() string {
	return "woof v1!"
}

func Barks() string {
	return "woof! woof! woof! v1"
}

func BigBark() string {
	return golang_dog.WhenGrownUp(Barks())
}
