package gitignore

import (
	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/file"
)

type GitIgnoreService struct {
	cm *manager.ConfigurationManager
	fs *file.FileService
}

func NewGitIgnoreService(cm *manager.ConfigurationManager) *GitIgnoreService {
	return &GitIgnoreService{
		cm: cm,
		fs: file.NewFileService(cm),
	}
}

func (this *GitIgnoreService) Create() error {
	return this.fs.CreateProjectFile(".gitignore", ".gitignore")
}
