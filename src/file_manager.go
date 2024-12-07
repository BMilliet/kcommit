package src

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type FileManager struct {
	HomeDir        string
	KcommitDir     string
	KcommitHistory string
}

func NewFileManager() (*FileManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting home directory: %v", err)
	}

	KcommitDir := filepath.Join(homeDir, ".kcommit")
	KcommitHistory := filepath.Join(KcommitDir, ".kcommit_history.json")

	return &FileManager{
		HomeDir:        homeDir,
		KcommitDir:     KcommitDir,
		KcommitHistory: KcommitHistory,
	}, nil
}

func (m *FileManager) ensureKcommitDir() error {
	if _, err := os.Stat(m.KcommitDir); os.IsNotExist(err) {
		err := os.Mkdir(m.KcommitDir, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}
	}
	return nil
}

func (m *FileManager) CheckIfPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("error checking if path exists: %v", err)
}

func (m *FileManager) checkAndCreateFile(filePath string) error {
	exists, err := m.CheckIfPathExists(filePath)
	if err != nil {
		return err
	}
	if !exists {
		_, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating %s: %v", filePath, err)
		}
	}
	return nil
}

func (m *FileManager) ReadFileContent(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %v", filePath, err)
	}
	return string(data), nil
}

func ParseJSONContent[T any](jsonString string, targetStruct *T) error {
	err := json.Unmarshal([]byte(jsonString), targetStruct)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}
	return nil
}

func (m *FileManager) BasicSetup() error {
	if err := m.ensureKcommitDir(); err != nil {
		return err
	}

	files := []string{
		m.KcommitHistory,
	}

	for _, file := range files {
		if err := m.checkAndCreateFile(file); err != nil {
			return err
		}
	}

	return nil
}
