// package cache implemets logic of functions build cache
package cache

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
)

var (
	ErrInvalidSchemeHash = errors.New("invalid scheme hash")
	ErrInvalidCommitID   = errors.New("invalid commit id")
	ErrNoFunction        = errors.New("no function")
)

type sumFileItem struct {
	function   string
	commitID   string
	schemeHash []byte
}

type Manager struct {
	sumFile       io.Reader
	items         []sumFileItem
	cacheLocation string
}

func NewManager(sumFile io.Reader, cacheLocation string) (*Manager, error) {
	r := csv.NewReader(sumFile)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	m := &Manager{
		sumFile:       sumFile,
		items:         []sumFileItem{},
		cacheLocation: cacheLocation,
	}

	for _, row := range rows {
		if len(row) < 3 {
			return nil, errors.New("sum file is invalid")
		}
		item := sumFileItem{
			function:   row[0],
			commitID:   row[1],
			schemeHash: []byte(row[2]),
		}
		m.items = append(m.items, item)
	}
	return m, nil
}

func (m *Manager) Close() error {

}

// GetFunctionCache compares not empty commitID, hash of the composition scheme and returns path to the ZIP containing built function.
func (m *Manager) GetFunctionCache(function, commitID string, schemeHash []byte) (string, error) {
	item := m.find(function)
	if item == nil {
		return "", ErrNoFunction
	}

	if !bytes.Equal(item.schemeHash, schemeHash) {
		return "", ErrInvalidSchemeHash
	}
	if commitID != "" && item.commitID != commitID {
		return "", ErrInvalidCommitID
	}

	return m.getFunction(function, commitID, schemeHash)
}

func (m *Manager) SetFunctionCache(function string, commitID string, schemeHash []byte, file io.Reader) error {
	if _, err := m.GetFunctionCache(function, commitID, schemeHash); err == nil {
		return nil
	}
	return m.saveFunction(function, commitID, schemeHash, file)
}

func (m *Manager) getFunction(function string, commitID string, schemeHash []byte) (string, error) {
	path := m.getFilePath(function, commitID, schemeHash)
	fd, err := os.Open(path)
	if fd != nil {
		defer fd.Close()
	}
	return path, err
}

func (m *Manager) saveFunction(function string, commitID string, schemeHash []byte, file io.Reader) error {
	path := m.getFilePath(function, commitID, schemeHash)
	fd, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = io.Copy(fd, file)
	return err
}

func (m *Manager) getFilePath(function string, commitID string, schemeHash []byte) string {
	fileName := fmt.Sprintf("%s_%s_%s.zip", function, commitID, string(schemeHash))
	return path.Join(m.cacheLocation, fileName)
}

func (m *Manager) find(function string) *sumFileItem {
	for i := range m.items {
		if m.items[i].function == function {
			return &m.items[i]
		}
	}
	return nil
}
