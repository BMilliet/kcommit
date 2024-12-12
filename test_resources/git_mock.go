package testresources

type GitMock struct {
	GetCurrentBranchReturnValue string
	GetCurrentBranchCalled      bool

	GitCommitReturnValue string
	GitCommitCalled      bool

	IsGitRepositoryReturnValue bool
	IsGitRepositoryCalled      bool
}

func (g *GitMock) GetCurrentBranch() (string, error) {
	g.GetCurrentBranchCalled = true
	return g.GetCurrentBranchReturnValue, nil
}

func (g *GitMock) GitCommit(msg string) (string, error) {
	g.GitCommitCalled = true
	g.GitCommitReturnValue = msg
	return g.GitCommitReturnValue, nil
}

func (g *GitMock) IsGitRepository() bool {
	g.IsGitRepositoryCalled = true
	return g.IsGitRepositoryReturnValue
}
