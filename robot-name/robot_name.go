package robotname

import (
	"fmt"
	"time"
)

type Robot struct {
	name string
}
func (robot *Robot) Name() string {
	if robot.name == "" {
		robot.Reset()
	}

	return robot.name
}

func generateName() string {
	n := time.Now().Nanosecond()
	return string(fmt.Sprintf("%c%c%d%d%d", 'A'+n%2, 'A'+n%26000/1000, n%1000/100, n%100/10, n%10))
}

func (robot *Robot) Reset() {
	robot.name = generateName()
	fmt.Println(robot.name )
}