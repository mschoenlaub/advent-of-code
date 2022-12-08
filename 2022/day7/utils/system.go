package utils

const (
	CommandCd          = "cd"
	DirectoryLineStart = 'd'
	CommandStart       = '$'
)

type System struct {
	Root *Node
	Cwd  *Node
}

func (s *System) Cd(path string) {
	switch path {
	case "/":
		s.Cwd = s.Root
	case "..":
		s.Cwd = s.Cwd.Parent
	case ".":
	default:
		s.Cwd = s.Cwd.Children[path]
	}
}
