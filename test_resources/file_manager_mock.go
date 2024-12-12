package testresources

import (
	"errors"
	"fmt"
)

type FileManagerMock struct {
	CheckIfPathExistsReturns    map[string]interface{}
	CheckIfPathExistsCalledWith []string
	CheckIfPathExistsCalled     bool

	ReadFileContentReturns     map[string]interface{}
	ReadFileContentCalleddWith []string
	ReadFileContentCalled      bool

	GetHistoryContentReturns string
	GetHistoryContentCalled  bool

	WriteHistoryContentWrittenContent string

	BasicSetupReturnValue error
	BasicSetupCalled      bool

	GetCurrentDirectoryNameReturnValue string
	GetCurrentDirectoryNameCalled      bool
}

func (m *FileManagerMock) CheckIfPathExists(path string) (bool, error) {
	m.CheckIfPathExistsCalled = true
	m.CheckIfPathExistsCalledWith = append(m.CheckIfPathExistsCalledWith, path)

	if result, exists := m.CheckIfPathExistsReturns[path]; exists {
		if success, ok := result.(bool); ok {
			return success, nil
		}
		if err, ok := result.(error); ok {
			return false, err
		}
		return false, errors.New("invalid mock value for CheckIfPathExists")
	}
	return false, fmt.Errorf("path not found in mock: %s", path)
}

func (m *FileManagerMock) ReadFileContent(filePath string) (string, error) {
	m.ReadFileContentCalled = true
	m.ReadFileContentCalleddWith = append(m.ReadFileContentCalleddWith, filePath)

	if result, exists := m.ReadFileContentReturns[filePath]; exists {
		if success, ok := result.(string); ok {
			return success, nil
		}
		if err, ok := result.(error); ok {
			return "", err
		}
		return "", errors.New("invalid mock value for CheckIfPathExists")
	}
	return "", fmt.Errorf("path not found in mock: %s", filePath)
}

func (m *FileManagerMock) GetHistoryContent() (string, error) {
	m.GetHistoryContentCalled = true
	return m.GetHistoryContentReturns, nil
}

func (m *FileManagerMock) WriteHistoryContent(content string) error {
	m.WriteHistoryContentWrittenContent = content
	return nil
}

func (m *FileManagerMock) BasicSetup() error {
	m.BasicSetupCalled = true
	return m.BasicSetupReturnValue
}

func (m *FileManagerMock) GetCurrentDirectoryName() (string, error) {
	m.GetCurrentDirectoryNameCalled = true
	return m.GetCurrentDirectoryNameReturnValue, nil
}
