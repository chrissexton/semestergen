package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog/log"
)

//go:embed tpl/*
var embeddedFS embed.FS

const layout = "2006-01-02"

var tplMap = map[string]string{
	"assignments": "tpl/assignments.adoc.tpl",
	"ics":         "tpl/assignments.ics.tpl",
	"schedule":    "tpl/schedule.adoc.tpl",
	"syllabus":    "tpl/syllabus.adoc.tpl",
	"course":      "tpl/course.taskpaper.tpl",
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
	forbidden := []string{
		"#",
		"&",
		":",
		"'",
		`"`,
		`/`,
		"$",
		"%",
		"^",
		"(",
		")",
		`’`,
		`|`,
		`@`,
		`+`,
		`\`,
		`<`,
		`>`,
		`?`,
		`|`,
		`[`,
		`]`,
		`{`,
		`}`,
		`,`,
		`.`,
	}
	for _, s := range forbidden {
		out = strings.ReplaceAll(out, s, "")
	}
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

	Num  int
	Date time.Time
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

// Mon Jan 2 15:04:05 MST 2006
// 19971210T080000Z

func GetDTStamp() string {
	return time.Now().Format("20060102T150405Z")
}

func (c Config) ProjectSlug() string {
	return strings.ReplaceAll(c.Project, " ", "-")
}

func (c Config) GetDTStart(day int, override *time.Time) string {
	d := c.Dates[day].Format("20060102")
	if override != nil {
		d = override.Format("20060102")
	}
	return d
}

func (c Config) GetDate(day int, override *time.Time) string {
	if day == -1 {
		return "TBD"
	}
	d := c.Dates[day-1].Format("01-02")
	if override != nil {
		d = override.Format("01-02")
	}
	return d
}

func (c Config) GetDateNum(day int) string {
	if day == -1 {
		return "TBD"
	}
	return c.GetDate(day, nil)
}

var syllabusFile = flag.String("syllabus", "readme.adoc", "name of the syllabus file")

func main() {
	flag.Parse()

	log.Debug().Msgf("semestergen 0.03")

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
		if err := writeICS(c); err != nil {
			panic(err)
		}
		if err := writeTaskPaper(c); err != nil {
			panic(err)
		}
	}
}

func writeTaskPaper(c Config) error {
	f, err := os.Create("course.taskpaper")
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
	src, _ := embeddedFS.ReadFile(tplName)
	tpl, err := template.New(tplName).Funcs(funcs).Parse(string(src))
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
	src, _ := embeddedFS.ReadFile(tplName)
	tpl, err := template.New(tplName).Funcs(funcs).Parse(string(src))
	if err != nil {
		return err
	}
	err = tpl.Funcs(funcs).Execute(f, c)
	return err
}

func writeICS(c Config) error {
	f, err := os.Create(c.ProjectSlug() + ".ics")
	defer f.Close()
	if err != nil {
		return err
	}
	funcs := template.FuncMap{
		"getDate":     c.GetDate,
		"getDateNum":  c.GetDateNum,
		"getDTStamp":  GetDTStamp,
		"getDTStart":  c.GetDTStart,
		"projectSlug": c.ProjectSlug,
	}
	tplName := tplMap["ics"]
	src, _ := embeddedFS.ReadFile(tplName)
	tpl, err := template.New(tplName).Funcs(funcs).Parse(string(src))
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
	src, _ := embeddedFS.ReadFile(tplName)
	tpl, err := template.New(tplName).Parse(string(src))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
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
	src, _ := embeddedFS.ReadFile(tplName)
	tpl, err := template.New(tplName).Parse(string(src))
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
	c.Dates = mkDates(c.Start, c.End, c.DaysOff)
	for i := range c.Days {
		c.Days[i].Num = i + 1
		c.Days[i].Date = c.Dates[i]
	}
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
	for i, d := 0, begin; d.Before(end.AddDate(0, 0, 1)); d = nextD(d) {
		if !in(daysOff, d) {
			out[i] = d
			i++
			log.Printf("Day %d: %s", i, d)
		} else {
			log.Printf("%s is a day off", d)
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
