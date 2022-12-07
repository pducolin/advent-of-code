package filesystem

type File struct {
	name string
	size int
}

func NewFile(name string, size int) File {
	return File{
		name: name,
		size: size,
	}
}

type Directory struct {
	Name   string
	Parent *Directory

	Subdirectories map[string]*Directory

	Files []File

	size int
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		Name:           name,
		Parent:         parent,
		Subdirectories: map[string]*Directory{},
		Files:          []File{},
		size:           -1,
	}
}

func (directory *Directory) AddSubdirectory(name string) {
	directory.Subdirectories[name] = NewDirectory(name, directory)
}

func (directory *Directory) AddFile(file File) {
	directory.Files = append(directory.Files, file)
}

func (directory *Directory) GetOrEvaluateSize() int {
	if directory.size != -1 {
		return directory.size
	}

	size := 0
	for _, file := range directory.Files {
		size += file.size
	}
	for _, subdir := range directory.Subdirectories {
		size += subdir.GetOrEvaluateSize()
	}

	directory.size = size
	return size
}
