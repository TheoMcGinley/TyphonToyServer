package db

// mimics a real DB with in-memory map of post ID to post contents
var internalDB map[int]string

func Delete(postID int) {
	internalDB[postID] = ""
}

func Post(postID int, postContents string) error {
	internalDB[postID] = postContents
	return nil
}

func Get(postID int) string {
	return internalDB[postID]
}

func init() {
	internalDB = make(map[int]string)
}