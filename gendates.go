package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/rs/zerolog/log"

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
	ID string

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
	out := strings.ReplaceAll(l.ID, " ", "-")
	out = strings.ReplaceAll(out, "#", "")
	out = strings.ReplaceAll(out, "&", "")
	out = strings.ReplaceAll(out, ":", "")
	out = strings.ReplaceAll(out, "'", "")
	out = strings.ReplaceAll(out, `""`, "")
	out = strings.ReplaceAll(out, "$", "")
	out = strings.ReplaceAll(out, "%", "")
	out = strings.ReplaceAll(out, "^", "")
	log.Debug().Str("l.ID", l.ID).Str("out", out).Msgf("slug")
	return out
}

type Assignment struct {
	// Assignment name
	Title string

	// Assignment Link plus supplemental material
	Links []Link

	// Day of the semester assignment is due (must be a class day)
	Due int

	// DueDate date in YYYY-MM-DD HH:mm format
	DueDate *time.Time

	// Day of the semester assignment is assigned (must be a class day)
	Assigned int

	// AssignedDate date in YYYY-MM-DD format
	AssignedDate *time.Time
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
	ICalLink    string

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

func (c Config) GetDate(day int, override *time.Time) string {
	d := c.Dates[day].Format("01-02")
	if override != nil {
		d = override.Format("01-02")
	}
	return d
}

func (c Config) GetDateNum(day int) string {
	return c.GetDate(day, nil)
}

var box *packr.Box

var syllabusFile = flag.String("syllabus", "readme.adoc", "name of the syllabus file")

func main() {
	flag.Parse()

	box = packr.New("templates", "./tpl")

	log.Debug().Msgf("semestergen 0.03")

	box.Walk(func(s string, file packd.File) error {
		log.Debug().Msgf("box file: %s", s)
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
		"getDate":    c.GetDate,
		"getDateNum": c.GetDateNum,
		"dueTime":    func() string { return c.DueTime },
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
		"getDate":    c.GetDate,
		"getDateNum": c.GetDateNum,
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
	f, err := os.Create(*syllabusFile)
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

var linkChecker = func() func(Link) Link {
	ids := map[string]bool{}
	return func(l Link) Link {
		if !ids[l.Title] {
			ids[l.Title] = true
			l.ID = l.Title
			return l
		}

		ext := 1
		potential := fmt.Sprintf("%s-%d", l.Title, ext)

		for ids[potential] {
			ext += 1
			potential = fmt.Sprintf("%s-%d", l.Title, ext)
		}

		l.ID = potential
		ids[potential] = true

		return l
	}
}()

func mkConfig(path string) Config {
	var c Config
	if _, err := toml.DecodeFile(path, &c); err != nil {
		panic(err)
	}
	for i := range c.Days {
		c.Days[i].Num = i + 1
	}
	c.Dates = mkDates(c.Start, c.End, c.DaysOff)
	for i, assn := range c.Assignments {
		for j, l := range assn.Links {
			c.Assignments[i].Links[j] = linkChecker(l)
		}
	}
	for i, day := range c.Days {
		for j, l := range day.Links {
			c.Days[i].Links[j] = linkChecker(l)
		}
	}
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
