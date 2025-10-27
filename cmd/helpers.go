package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var cmdPrefix = "calculator-"

var ErrPPluginNotFound = errors.New("cannot find plugin in PATH")

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func StringListToIntList(listStrings []string) []int {
	n := len(listStrings)
	numbers := make([]int, n)

	for i, s := range listStrings {
		numbers[i] = StringToInt(s)
	}

	return numbers
}

func FindPlugins() []string {
	pluginlist := []string{}
	paths := strings.Split(os.Getenv("PATH"), ":")
	// go over all paths in the PATH environment
	// and add them to the completion command
	for _, path := range paths {
		// list all files in path
		files, err := os.ReadDir(path)
		if err != nil {
			continue
		}
		// add all files that start with rootCmd-
		for _, file := range files {
			if strings.HasPrefix(file.Name(), cmdPrefix) {
				basep := strings.ReplaceAll(file.Name(), cmdPrefix, "")
				fpath := filepath.Join(path, file.Name())

				info, err := os.Stat(fpath)
				if err != nil {
					continue
				}

				if info.Mode()&0o111 != 0 {
					pluginlist = append(pluginlist, basep)
				}
			}
		}
	}

	return pluginlist
}

// FindPlugin find a binary in plugin homedir directory or user paths.
func FindPlugin(pluginame string) (string, error) {
	cmd := cmdPrefix + pluginame
	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {
		// list all files in path
		files, err := os.ReadDir(path)
		if err != nil {
			continue
		}

		// add all files that start with rootCmd-
		for _, file := range files {
			if strings.EqualFold(file.Name(), cmd) {
				return filepath.Join(path, file.Name()), nil
			}
		}
	}

	return "", ErrPPluginNotFound
}
