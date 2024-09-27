package tools

import (
	"fmt"
	"os"
	"path"
)

type TraverseCallback func(name string) bool

type Traverse struct {
	Root        string
	EnterFolder TraverseCallback
	ExitFolder  TraverseCallback
	ProcessFile TraverseCallback
}

func NewTraverse(root string) *Traverse {
	return &Traverse{Root: root}
}

func (t *Traverse) WithOnEnterFolder(OnEnter TraverseCallback) *Traverse {
	t.EnterFolder = OnEnter
	return t
}

func (t *Traverse) WithOnExitFolder(OnExit TraverseCallback) *Traverse {
	t.ExitFolder = OnExit
	return t
}

func (t *Traverse) WithProcessFile(Process TraverseCallback) *Traverse {
	t.ProcessFile = Process
	return t
}

func (t *Traverse) Run() error {
	return traverseFolder(t.Root, t.EnterFolder, t.ExitFolder, t.ProcessFile)
}

// Traverse recursively the specified root folder
// Gather the list of files and the list of subfolders
// If there are files to be processed it calls ProduceEmbedGo()
func traverseFolder(root string, onEnter TraverseCallback, onExit TraverseCallback, process TraverseCallback) error {
	files := make([]string, 0)
	folders := make([]string, 0)

	// Get the contents of the current folder
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	// Gather all files and folders of the current folder
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		} else {
			files = append(files, path.Base(entry.Name()))
		}
	}

	// If no files nor folders found, we return
	if len(files)+len(folders) == 0 {
		return nil
	}

	// Tell the caller we just entered the folder
	// We then process the files only if the caller returns true
	// This allows the caller to skip processing certain folders
	if onEnter == nil || onEnter(root) {
		// Now process all files
		for _, file := range files {
			// If user callback is defined and returns false, we end the loop
			if process != nil && !process(path.Join(root, file)) {
				break
			}
		}

		// Now process all the folders.
		for _, folder := range folders {
			err := traverseFolder(path.Join(root, folder), onEnter, onExit, process)
			if err != nil {
				return fmt.Errorf("%s", err.Error())
			}
		}
	}

	// Now tell caller that we are leaving the folder
	if onExit != nil {
		onExit(root)
	}

	return nil
}
