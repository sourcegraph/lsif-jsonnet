package refs

import (
	"os"
	"path/filepath"
)

type PathResolver struct {
	basePaths []string
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func NewPathResolver(basePaths []string) (*PathResolver, error) {
	absPaths := make([]string, 0, len(basePaths))

	for _, path := range basePaths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return nil, err
		}
		absPaths = append(absPaths, absPath)
	}
	return &PathResolver{absPaths}, nil
}

func (pr *PathResolver) AddPath(basePath string) error {
	absPath, err := filepath.Abs(basePath)
	if err != nil {
		return err
	}
	pr.basePaths = append(pr.basePaths, absPath)
	return nil
}

func (pr *PathResolver) Resolve(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	for _, bp := range pr.basePaths {
		candidate := filepath.Join(bp, path)

		ex, err := PathExists(candidate)
		if err != nil {
			return "", err
		}

		if ex {
			return candidate, nil
		}
	}
	return "", nil
}
