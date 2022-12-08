package main

import (
	"advent-of-code/2022/day7/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("2022/day7/input.txt")
	scanner := bufio.NewScanner(f)

	cwd := utils.NewNode(nil, "/", 0)
	system := utils.System{Root: &cwd, Cwd: &cwd}

	for scanner.Scan() {
		line := scanner.Text()
		switch line[0] {
		case utils.CommandStart: // $ cd /abc
			line = line[2:]
			split := strings.SplitN(line, " ", 2) // ["cd", "/we have a space maybe"]
			if split[0] == utils.CommandCd {
				system.Cd(split[1])
			}
		case utils.DirectoryLineStart:
			if system.Cwd.Children[line[4:]] == nil {
				system.Cwd.AddChild(line[4:], 0) // directories have no inherent size
			}
		default: // output of ls for files
			split := strings.SplitN(line, " ", 2) // ["123", "maybe another space.txt"]
			size, _ := strconv.Atoi(split[0])
			path := split[1]
			system.Cwd.AddChild(path, size)
		}
	}

	free := 70000000 - system.Root.Size()
	needed := 30000000 - free

	finder := utils.NewMinSizeVisitor(needed)
	finder.Visit(system.Root)
	fmt.Println(finder.MinSize())
}
