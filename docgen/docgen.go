// based on doc-extract<https://github.com/joeshaw/doc-extract/blob/master/parser.go>
package docgen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path"
	"sort"
	"strings"
)

type file struct {
	file *ast.File
	name string
}
// sort in lexigraphical order except main.go
func sortedFiles(pkg *ast.Package) []file {
	files := make([]file, 0, len(pkg.Files))

	for name, f := range pkg.Files {
		files = append(files, file{file: f, name: path.Base(name)})
	}

	sort.Slice(files, func(i, j int) bool {
		ni, nj := files[i].name, files[j].name

		if ni == "main.go" {
			return true
		}

		if nj == "main.go" {
			return false
		}

		return ni < nj
	})

	return files
}

func extractComment(cgrp *ast.CommentGroup) (string, bool) {
	s := cgrp.Text()
	parts := strings.SplitN(s, "\n", 2)
	if strings.TrimSpace(parts[0]) == "[docgen]" {
		return parts[1], true
	}
	return "", false
}

func extractPackageComments(pkg *ast.Package) []string {
	files := sortedFiles(pkg)

	var comments []string
	for _, f := range files {
		for _, c := range f.file.Comments {
			s, ok := extractComment(c)
			if ok {
				comments = append(comments, s)
			}
		}
	}

	return comments
}

func main() {

	srcDir := os.Args[1]
	outFile := os.Args[2]

	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, srcDir, nil, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing files in %s: %s\n", srcDir, err)
		os.Exit(1)
	}

	var out io.Writer
	if outFile == "-" {
		out = os.Stdout
	} else {
		f, err := os.Create(outFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file %s: %s\n", outFile, err)
			os.Exit(1)
		}
		defer f.Close()

		out = f
	}

	for _, pkg := range pkgs {
		comments := extractPackageComments(pkg)
		for _, c := range comments {
			fmt.Fprintln(out, c)
		}
	}
}
