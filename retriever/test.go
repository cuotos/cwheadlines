package retriever

type Test struct{}

func (Test) GetSolution() (string, error) {
	return "test", nil
}
