package models

type Profile struct {
	NickName string
	Score    string
	Position int
	Rank     string
	Solved   map[string][]Task
}

func (p *Profile) AddSolved(rubric string, task Task) {
	p.Solved[rubric] = append(p.Solved[rubric], task)
}

func NewTask(title, difficulty, date string) Task {
	return Task{
		Title:      title,
		Difficulty: difficulty,
		Date:       date,
	}
}

type Task struct {
	Title      string
	Difficulty string
	Date       string
}
