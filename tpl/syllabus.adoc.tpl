= {{.Title}}
:favicon: ./resources/favicon.jpg
:nofooter:
:toc: left
ifdef::backend-html5[]
endif::[]
ifdef::backend-pdf[]
:notitle:
:pdf-style: theme.yml
endif::[]
:icons: font
:book: link:https://www.amazon.com/dp/013394302X/
:0001: link:https://falconair.github.io/2015/01/30/composingcontracts.html

ifdef::backend-pdf[= {{.Title}}]

== Information

**Instructor**: {{.Instructor}}

**Office**: {{.Office}}

**Phone**: {{.Phone}}

**Email**: mailto:{{.Email}}[{{.Email}}] (**Email is best!**)

**Meetings**: {{.Meetings}}

*Text*: {{.Text}}

include::status.adoc[opts=optional]

== Course

{{.Description}}

== Schedule

Note: *Many of these links may be dead or locked.* They will be live when we reach the
content represented. Please ask via email if there is a
particular file missing from the course that should be live.
 
include::schedule.adoc[]
 
=== Assignments

include::assignments.adoc[]

This schedule is tentative.

{{if .Resources}}
== Resources

{{.Resources}}
{{end}}

== Grading Scale

[%header,format=psv]
|===
| Grade Percentage | Letter Grade 
| 90-100           | A            
| 80-89            | B            
| 70-79            | C            
| 60-69            | D            
| 0-59             | F            
|===

A +/- will be added for the upper/lower two points of each grade respectively.

Note that a C is the minimum grade accepted for Natural Science Degrees.

== Course Evaluation

[%header,format=psv]
|===
| Category | Percentage
{{- range .Evaluation}}
| **{{.Title}}**   | {{.Value}}
{{- end}}
|===

{{if .EvalText}}
{{.EvalText}}
{{else}}
== Homework

. Homework assignments cover areas you need to know and practice.
. Submit homework at the beginning of class or by time posted in Canvas if online.
. Your submission must be your own work.

== In-Class

. Each class may have an in-class element that counts towards the in-class evaluation.
. These assignments will not be graded, but participation points will be awarded.
. Class participation will also be assigned in this category.

{{end}}

{{.Legal}}

== Code of Student Rights, Responsibilities and Conduct

You are responsible for knowing the IU Code of Student Rights, Responsibilities, and Conduct. http://www.iu.edu/~code/[]

Student responsibilities outlined in the code include Academic Misconduct and Personal Misconduct. Academic Misconduct includes cheating, fabrication, plagiarism, interference, violation of course rules, and facilitating academic dishonesty. Personal Misconduct includes acts of personal misconduct both on and off university property. Ignorance of the rules is not a defense.

=== Plagiarism

Plagiarism is defined as presenting someone else's work, including the work of other students, as one's own. Any ideas or materials taken from another source for either written or oral use must be fully acknowledged, unless the information is common knowledge. What is considered "common knowledge" may differ from course to course.

A student must not adopt or reproduce ideas, opinions, theories, formulas, graphics, or pictures of another person without acknowledgment. A student must give credit to the originality of others and acknowledge indebtedness whenever: Directly quoting another person’s actual words, whether oral or written; Using another person’s ideas, opinions, or theories; Paraphrasing the words, ideas, opinions, or theories of others, whether oral or written; Borrowing facts, statistics, or illustrative material; or Offering materials assembled or collected by others in the form of projects or collections without acknowledgment. Cheating:

Cheating is considered to be an attempt to use or provide unauthorized assistance, materials, information, or study aids in any form and in any academic exercise or environment.

=== Penalties

* Cheating on homework assignments - loss of all points for that homework and if severe, a failing grade for the course.
* Cheating on quizzes - loss of all points for that quiz and if severe, a failing grade for the course.
* Cheating on exams - loss of all points for that exam and if severe, a failing grade for the course.

== Course Policies

* Please do not hesitate to email me at mailto:cwsexton@ius.edu[cwsexton@ius.edu] if you need to get in touch outside of campus or outside of office hours. I will be happy to resolve issues via email or set up a video call. I will do my best to respond to your emails within 24 hours during business hours Monday through Thursday. It may take me longer to respond on the weekends.
* Feel free to use your laptop during class. Taking notes and following along with exercises is a great way to stay engaged with class. Please avoid allowing any technology use to become a distraction. Participation points may be lowered in this case.
* Please keep your personal computing devices on silent or vibrate.
* If problems occur or if you become ill, please contact me immediately so we can determine your best options.
* If you have problems with equipment, please let me know but you should also contact a person at the computer helpdesk immediately at (812) 941-2447. Technical issues are **not** an excuse for late work. Assignments are given with plenty of lead time to proactively solve technical issues.
* Please proof all assignments and email messages to ensure the use of Standard English, proper grammar, and correct spelling. You will lose points for problems in this area. If you have concerns about your writing, contact the https://www.ius.edu/writing-center/[Writing Center] for a consultation.
* Class Attendance:  At IU Southeast, attendance is required. Participation points are non-recoverable for absences.
* You’re probably used to seeing many policy statements on a syllabus.  Faculty include these statements to ensure you understand course expectations so that you can succeed in your courses.  At IU Southeast, we have placed all university policies on a single website easily accessed from every Canvas course site. Simply look at the left navigation bar and click on https://www.ius.edu/get-help/[Succeed at IU Southeast]. You can find links to sites with a great deal of useful information including

** How to avoid plagiarism and cheating
** Disability Services
** FLAGS
** Tutoring centers
** Canvas Guides
** Financial Aid
** Sexual Misconduct
** Counseling
** Writing Center and much more!
 
+
My expectation is that you review university policies carefully to ensure you understand the policy and possible consequences for violating the policy.  Please contact me if you have any questions about any university policy.

* All labs/assignments/forums/quizzes/tests/etc. will be open for a time window, at least a number of hours and in many cases a number of days.
+
Documented illness is the only acceptable excuse for not completing an assignment during its open window. Other reasons for not completing labs/assignments/forums/quizzes/tests/etc. must be explained to the satisfaction of the instructor, who will decide whether missed assignments may be made up. Being sick on the last few days of an assignment's due date is *NOT* an excuse.

== Disclaimer

Although every effort has been made to make the above listing complete and accurate, minor adjustments to the schedule are sometimes necessary due to weather, or other problems that crop up. The grading scale and late policy will remain constant.
