package golanggaselection

type Selector interface {
	Select(Individuals) (Individuals, error)
}
