package main

import (
	"flag"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/BurntSushi/toml"
)

const layout = "2006-01-02"

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
	Resources   []Link
	Evaluation  []Eval
	EvalText    string

	Dates DayMap
}

func (c Config) GetDate(day int) string {
	return c.Dates[day].Format("01-02")
}

func main() {
	flag.Parse()
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
	//tpl := template.Must(template.ParseFiles("assignments.adoc.tpl"))
	tpl, err := template.New("course.task.tpl").Funcs(funcs).ParseFiles("course.task.tpl")
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
	//tpl := template.Must(template.ParseFiles("assignments.adoc.tpl"))
	tpl, err := template.New("assignments.adoc.tpl").Funcs(funcs).ParseFiles("assignments.adoc.tpl")
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
	tpl := template.Must(template.ParseFiles("schedule.adoc.tpl"))
	err = tpl.Execute(f, c)
	return err
}

func writeSyllabus(c Config) error {
	f, err := os.Create("syllabus.adoc")
	defer f.Close()
	if err != nil {
		return err
	}
	tpl := template.Must(template.ParseFiles("syllabus.adoc.tpl"))
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
