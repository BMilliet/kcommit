package src

import (
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
		return nil, fmt.Errorf("NewFileManager -> %v", err)
	}

	KcommitDir := filepath.Join(homeDir, KcommitDirName)
	KcommitHistory := filepath.Join(KcommitDir, KcommitHistoryFileName)

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
			return fmt.Errorf("ensureKcommitDir -> %v", err)
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
	return false, fmt.Errorf("CheckIfPathExists -> %v", err)
}

func (m *FileManager) checkAndCreateFile(filePath string) error {
	exists, err := m.CheckIfPathExists(filePath)
	if err != nil {
		return err
	}
	if !exists {
		_, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("checkAndCreateFile -> %s %v", filePath, err)
		}
	}
	return nil
}

func (m *FileManager) ReadFileContent(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("ReadFileContent -> %s %v", filePath, err)
	}
	return string(data), nil
}

func (m *FileManager) GetHistoryContent() (string, error) {
	str, err := m.ReadFileContent(m.KcommitHistory)
	if err != nil {
		return "", fmt.Errorf("GetHistoryContent -> %s %v", m.KcommitHistory, err)
	}
	return str, nil
}

func (m *FileManager) WriteHistoryContent(content string) error {
	err := m.writeFileContent(m.KcommitHistory, content)
	if err != nil {
		return fmt.Errorf("WriteHistoryContent -> %s: %v", m.KcommitHistory, err)
	}
	return nil
}

func (m *FileManager) writeFileContent(filePath, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("writeFileContent -> %s %v", filePath, err)
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
