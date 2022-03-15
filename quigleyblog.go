package quigleyblog

import (
	"io/fs"
)

// NewPostsFromFS return Posts from given directory
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	var posts []Post

	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		post, err := getPost(fileSystem, file.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	post, err := newPost(postFile)
	return post, nil
}
