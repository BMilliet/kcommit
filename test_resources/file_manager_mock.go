package testresources

import (
	"errors"
	"fmt"
)

type FileManagerMock struct {
	CheckIfPathExistsReturns    map[string]interface{}
	CheckIfPathExistsCalledWith []string
	CheckIfPathExistsCalled     int

	ReadFileContentReturns     map[string]interface{}
	ReadFileContentCalleddWith []string
	ReadFileContentCalled      int

	GetHistoryContentReturns string
	GetHistoryContentCalled  int

	WriteHistoryContentWrittenContent string

	BasicSetupReturnValue error
	BasicSetupCalled      int

	GetCurrentDirectoryNameReturnValue string
	GetCurrentDirectoryNameCalled      int
}

func (m *FileManagerMock) CheckIfPathExists(path string) (bool, error) {
	m.CheckIfPathExistsCalled += 1
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
	m.ReadFileContentCalled += 1
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
	m.GetHistoryContentCalled += 1
	return m.GetHistoryContentReturns, nil
}

func (m *FileManagerMock) WriteHistoryContent(content string) error {
	m.WriteHistoryContentWrittenContent = content
	return nil
}

func (m *FileManagerMock) BasicSetup() error {
	m.BasicSetupCalled += 1
	return m.BasicSetupReturnValue
}

func (m *FileManagerMock) GetCurrentDirectoryName() (string, error) {
	m.GetCurrentDirectoryNameCalled += 1
	return m.GetCurrentDirectoryNameReturnValue, nil
}
