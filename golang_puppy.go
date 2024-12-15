package golang_puppy

import "github.com/ajayshinde007/golang_dog"

func Bark() string {
	return "woof v4!"
}

func Barks() string {
	return "woof! woof! woof! v4"
}

func BigBark() string {
	return golang_dog.WhenGrownUp(Barks())
}
