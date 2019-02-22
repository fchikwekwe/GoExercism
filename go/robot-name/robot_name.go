package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot has a single attribute called name that is a string
type Robot struct {
	name string
}

// robotNames keeps track of which names are already in use
var robotNames = make(map[string]bool)

// Name is....
func (robot *Robot) Name() (string, error) {
	if robot.name == "" {
		robot.name = ""
		rand.Seed(time.Now().UnixNano())
		for inUse, ok := string(robot.name) == "", true; ok && inUse; inUse, ok = robotNames[robot.name] {
			robot.name = letter() + letter() + number()
		}
		robotNames[robot.name] = true
	}
	return string(robot.name), nil
}

// Reset is a function that resets the robot's current name when called
func (robot *Robot) Reset() {
	robot.name = ""
}

// letter generates a random letter
func letter() string {
	return string(rand.Intn('Z'-'A') + 'A')
}

// number generates a random three digit number
func number() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}
