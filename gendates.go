package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"

	"github.com/BurntSushi/toml"
)

const layout = "2006-01-02"

var tplMap = map[string]string{
	"assignments": "assignments.adoc.tpl",
	"schedule":    "schedule.adoc.tpl",
	"syllabus":    "syllabus.adoc.tpl",
	"course":      "course.task.tpl",
}

type DayMap map[int]time.Time

type Eval struct {
	Title string
	Value string
}

type Link struct {
	// Text of Link
	Title string

	// Link target
	URL string

	Level int
}

func (l Link) Stars() string {
	out := "*"
	for i := 0; i < l.Level; i++ {
		out += "*"
	}
	return out
}

func (l Link) Slug() string {
	out := strings.ReplaceAll(l.Title, " ", "-")
	out = strings.ReplaceAll(out, "#", "")
	out = strings.ReplaceAll(out, "&", "")
	return out
}

type Assignment struct {
	// Assignment name
	Title string

	// Assignment Link plus supplemental material
	Links []Link

	// Day of the semester assignment is due (must be a class day)
	Due int

	// Day of the semester assignment is assigned (must be a class day)
	Assigned int
}

type Day struct {
	// Assignment name
	Title string

	// Assignment Link plus supplemental material
	Links []Link

	Num int
}

type Config struct {
	Title       string
	Instructor  string
	Office      string
	Phone       string
	Email       string
	Meetings    string
	Text        string
	Description string
	Legal       string

	Start   time.Time
	End     time.Time
	DaysOff []time.Time
	DueTime string
	Project string

	Days        []Day
	Assignments []Assignment
	Resources   string
	Evaluation  []Eval
	EvalText    string

	Dates DayMap
}

func (c Config) GetDate(day int) string {
	return c.Dates[day].Format("01-02")
}

var box *packr.Box

func main() {
	flag.Parse()

	box = packr.New("templates", "./tpl")

	log.Println("semestergen 0.02")

	box.Walk(func(s string, file packd.File) error {
		log.Printf("box file: %s", s)
		return nil
	})

	for i := 0; i < flag.NArg(); i++ {
		c := mkConfig(flag.Arg(i))
		if err := writeSyllabus(c); err != nil {
			panic(err)
		}
		if err := writeSchedule(c); err != nil {
			panic(err)
		}
		if err := writeAssignments(c); err != nil {
			panic(err)
		}
		if err := writeTaskPaper(c); err != nil {
			panic(err)
		}
	}
}

func writeTaskPaper(c Config) error {
	f, err := os.Create("course.task")
	defer f.Close()
	if err != nil {
		return err
	}
	funcs := template.FuncMap{
		"getDate": c.GetDate,
		"dueTime": func() string { return c.DueTime },
	}
	tplName := tplMap["course"]
	src, _ := box.FindString(tplName)
	tpl, err := template.New(tplName).Funcs(funcs).Parse(src)
	if err != nil {
		return err
	}
	err = tpl.Funcs(funcs).Execute(f, c)
	return err
}

func writeAssignments(c Config) error {
	f, err := os.Create("assignments.adoc")
	defer f.Close()
	if err != nil {
		return err
	}
	funcs := template.FuncMap{
		"getDate": c.GetDate,
	}
	tplName := tplMap["assignments"]
	src, _ := box.FindString(tplName)
	tpl, err := template.New(tplName).Funcs(funcs).Parse(src)
	if err != nil {
		return err
	}
	err = tpl.Funcs(funcs).Execute(f, c)
	return err
}

func writeSchedule(c Config) error {
	f, err := os.Create("schedule.adoc")
	defer f.Close()
	if err != nil {
		return err
	}
	tplName := tplMap["schedule"]
	src, err := box.FindString(tplName)
	if err != nil {
		return fmt.Errorf("error finding template: %w", err)
	}
	tpl, err := template.New(tplName).Parse(src)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
		return err
	}
	err = tpl.Execute(f, c)
	return err
}

func writeSyllabus(c Config) error {
	f, err := os.Create("syllabus.adoc")
	defer f.Close()
	if err != nil {
		return err
	}
	tplName := tplMap["syllabus"]
	src, _ := box.FindString(tplName)
	tpl, err := template.New(tplName).Parse(src)
	if err != nil {
		return err
	}
	err = tpl.Execute(f, c)
	return err
}

func mkConfig(path string) Config {
	var c Config
	if _, err := toml.DecodeFile(path, &c); err != nil {
		panic(err)
	}
	for i := range c.Days {
		c.Days[i].Num = i + 1
	}
	c.Dates = mkDates(c.Start, c.End, c.DaysOff)
	return c
}

func mkDates(begin, end time.Time, daysOff []time.Time) DayMap {
	if begin.Year() == 1 || end.Year() == 1 {
		return DayMap{}
	}
	out := make(DayMap)
	for i, d := 1, begin; d.Before(end.AddDate(0, 0, 1)); d = nextD(d) {
		if !in(daysOff, d) {
			out[i] = d
			i++
		}
	}
	return out
}

func in(days []time.Time, d0 time.Time) bool {
	for _, d := range days {
		if d0 == d {
			return true
		}
	}
	return false
}

func nextD(d0 time.Time) time.Time {
	if w := d0.Weekday(); w == time.Monday || w == time.Tuesday {
		return d0.AddDate(0, 0, 2)
	}
	return d0.AddDate(0, 0, 5)
}
