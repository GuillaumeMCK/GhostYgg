package utils

import "github.com/sqweek/dialog"

// PickFilePath opens a file picker dialog.go and returns the path of the file
func PickFilePath(msg string) (path string, err error) {
	if dialog.Message("%s", msg).YesNo() {
		path, err = dialog.File().Filter("Torrent files", "torrent").Load()
	}
	return
}
