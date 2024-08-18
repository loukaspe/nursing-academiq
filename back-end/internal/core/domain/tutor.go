package domain

type Tutor struct {
	User
	ID           uint
	AcademicRank string
	Courses      []Course
}
