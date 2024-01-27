package scan

import (
	"fmt"
	"gitcontrib/internal/utils"
	"io"
	"log"
	"os"
	"strings"
)

// scan given a path crawls it and its subfolders
// searching for Git repositories
func Scan(folder string) {
	fmt.Printf("Found folders:\n\n")
	repositories := recursiveScanFolder(folder)
	filePath := utils.GetDotFilePath()
	addNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func addNewSliceElementsToFile(filePath string, newRepos []string) {
	existingRepos := utils.ParseFileLinesToSlice(filePath)
	repos := joinSlices(newRepos, existingRepos)
	dumpStringsSliceToFile(repos, filePath)
}

// writes content to the file in path `filePath` (overwriting existing content)
func dumpStringsSliceToFile(repos []string, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content := strings.Join(repos, "\n")
	_, err = io.WriteString(f, content)
	if err != nil {
		panic(err)
	}
}

func joinSlices(new []string, existing []string) []string {
	for _, i := range new {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing
}

// sliceContains returns true if `slice` contains `value`
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}

// scanGitFolders returns a list of subfolders of `folder` ending with `.git`.
// Returns the base folder of the repo, the .git folder parent.
// Recursively searches in the subfolders by passing an existing `folders` slice.
func scanGitFolders(folders []string, folder string) []string {
	// trim last '/'
	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var path string
	
	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}
			if inWhitelist(file.Name()){
				continue
			}
			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}

func inWhitelist(fileName string) bool {
	whitelistedFolders := []string{"vendor", "node_modules"}
	for _,value := range whitelistedFolders {
		if value == fileName {
			return true
		}
	}
	return false
}