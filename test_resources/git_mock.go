package testresources

type GitMock struct {
	GetCurrentBranchReturnValue string
	GetCurrentBranchCalled      int

	GitCommitReturnValue string
	GitCommitCalled      int

	IsGitRepositoryReturnValue bool
	IsGitRepositoryCalled      int
}

func (g *GitMock) GetCurrentBranch() (string, error) {
	g.GetCurrentBranchCalled += 1
	return g.GetCurrentBranchReturnValue, nil
}

func (g *GitMock) GitCommit(msg string) (string, error) {
	g.GitCommitCalled += 1
	g.GitCommitReturnValue = msg
	return g.GitCommitReturnValue, nil
}

func (g *GitMock) IsGitRepository() bool {
	g.IsGitRepositoryCalled += 1
	return g.IsGitRepositoryReturnValue
}
