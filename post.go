package quigleyblog

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	title := readMetaLine(scanner, titleMeta)
	description := readMetaLine(scanner, descriptionMeta)
	tags := strings.Split(readMetaLine(scanner, tagsMeta), ",")

	scanner.Scan()

	body := readBody(scanner)

	post := Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}
	return post, nil
}

func readMetaLine(scanner *bufio.Scanner, tagName string) string {
	scanner.Scan()
	return strings.TrimPrefix(scanner.Text(), tagName)
}

func readBody(scanner *bufio.Scanner) string {
	buffer := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buffer, scanner.Text())
	}
	return buffer.String()
}

// Post is a single blog post
type Post struct {
	Title, Description, Body string
	Tags                     []string
}

const (
	titleMeta       = "Title: "
	descriptionMeta = "Description: "
	tagsMeta        = "Tags: "
)
