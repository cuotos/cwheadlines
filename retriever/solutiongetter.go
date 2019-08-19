package retriever

type SolutionGetter interface {
	GetSolution() (string, error)
}
