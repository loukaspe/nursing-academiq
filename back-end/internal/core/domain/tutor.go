package domain

type Tutor struct {
	User
	AcademicRank string
	Courses      []Course
}
