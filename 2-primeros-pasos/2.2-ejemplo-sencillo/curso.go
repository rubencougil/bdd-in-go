package curso

type Curso struct {
	Name string
	Votes int
}

func NewCurso(name string) *Curso {
	return &Curso{
		Name: name,
		Votes: 0,
	}
}

func (c *Curso) AddScore() {
	c.Votes++
}