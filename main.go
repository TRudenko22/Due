package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Assn struct {
	Date  time.Time
	Tasks []string
}

type Course struct {
	Name   string
	Assign []Assn
}

func (c Course) GetDiffDays() int {
	difference := c.Assign[0].Date.Sub(time.Now()).Hours()

	return int(difference) / 24
}

func (c Course) OutputDueDates() {
	var days string

	if c.GetDiffDays() != 1 {
		days = fmt.Sprintf("%-3v days\n", c.GetDiffDays())
	} else {
		days = fmt.Sprintf("%-3v day\n", c.GetDiffDays())
	}

	fmt.Printf("\nCourse: %v\n", c.Name)
	fmt.Printf("Assignment due in: %v", days)
	for _, i := range c.Assign[0].Tasks {
		fmt.Println(i)
	}

}

func Unmarshal(file, file_type string) (c Course, err error) {
	v := viper.New()
	v.SetConfigType(file_type)
	v.SetConfigFile(file)

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	err = v.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func main() {
	var courses []Course
	path := "./files/"

	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, i := range files {
		config, err := Unmarshal(path+i.Name(), "yaml")
		if err != nil {
			panic(err)
		}

		courses = append(courses, config)
	}

	for _, course := range courses {
	  course.OutputDueDates()
	}
}
