/*
Package gsmdrive implements the Drive API
Copyright © 2020 Hannes Hayashi

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package gsmdrive

import (
	"fmt"
	"sync"
	"time"

	drive "google.golang.org/api/drive/v3"
)

// Folder represents a structure useful for copying Drive folder trees to a new destination
type Folder struct {
	NewID     string
	OldParent string
	NewParent string
	Name      string
	Root      bool
}

const folderMimetype = "application/vnd.google-apps.folder"

// IsFolder returns true if the file object is a folder, otherwise false
// Make sure that the MimeType property is actually set.
func IsFolder(f *drive.File) bool {
	if f.MimeType == folderMimetype {
		return true
	}
	return false
}

// CopyFolders creates a copy of a Drive folder structure at a new destination
func CopyFolders(folderMap map[string]*Folder, destination string) error {
	var err error
	for k := range folderMap {
		if folderMap[k].NewID == "" {
			folderMap[k].NewID, err = getNewID(k, folderMap)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func createFolder(parent, name string) (*drive.File, error) {
	f := &drive.File{
		MimeType: folderMimetype,
		Parents:  []string{parent},
		Name:     name,
	}
	newFolder, err := CreateFile(f, nil, false, false, false, "", "", "id,mimeType,name")
	if err != nil {
		return nil, err
	}
	return newFolder, nil
}

func getNewID(oldID string, folderMap map[string]*Folder) (string, error) {
	if folderMap[oldID].NewID != "" {
		return folderMap[oldID].NewID, nil
	}
	var err error
	if !folderMap[oldID].Root {
		folderMap[oldID].NewParent, err = getNewID(folderMap[oldID].OldParent, folderMap)
		if err != nil {
			return "", err
		}
	}
	newFolder, err := createFolder(folderMap[oldID].NewParent, folderMap[oldID].Name)
	time.Sleep(200 * time.Millisecond)
	if err != nil {
		return "", err
	}
	folderMap[oldID].NewID = newFolder.Id
	return folderMap[oldID].NewID, nil
}

// GetFilesAndFolders recursively gets all files and folders below a parent folder and separates them,
// returning a map[string]*folder for folders and a simple list for files.
func GetFilesAndFolders(folderID string, threads int) (folderMap map[string]*Folder, files []*drive.File, err error) {
	folderMap = make(map[string]*Folder)
	root, err := GetFile(folderID, "id,mimeType,name,parents", "")
	if err != nil {
		return nil, nil, err
	}
	if !IsFolder(root) {
		return nil, nil, fmt.Errorf("%s is not a folder", folderID)
	}
	items, err := ListFilesRecursive(folderID, "files(id,parents,mimeType,name),nextPageToken", threads)
	if err != nil {
		return nil, nil, err
	}
	for _, i := range items {
		if IsFolder(i) {
			folderMap[i.Id] = &Folder{OldParent: i.Parents[0], Name: i.Name}
		} else {
			files = append(files, i)
		}
	}
	folderMap[folderID] = &Folder{Name: root.Name, OldParent: root.Parents[0], Root: true}
	return folderMap, files, nil
}

func listFilesRecursive(id, fields string, folders chan string, files chan *drive.File, wgFolders, wgFiles *sync.WaitGroup) error {
	result, err := ListFiles(fmt.Sprintf("'%s' in parents and trashed = false", id), "", "allDrives", "", "", "", fields, true)
	if err != nil {
		return err
	}
	for _, f := range result {
		wgFiles.Add(1)
		files <- f
		if IsFolder(f) {
			wgFolders.Add(1)
			folders <- f.Id
		}
	}
	return nil
}

// ListFilesRecursive lists all files and foldes in a parent folder recursively
func ListFilesRecursive(id, fields string, threads int) ([]*drive.File, error) {
	final := []*drive.File{}
	wgFolders := &sync.WaitGroup{}
	wgFiles := &sync.WaitGroup{}
	result, err := ListFiles(fmt.Sprintf("'%s' in parents and trashed = false", id), "", "allDrives", "", "", "", fields, true)
	if err != nil {
		return nil, err
	}
	folders := make(chan string, threads)
	files := make(chan *drive.File, threads)
	go func() {
		for id := range folders {
			time.Sleep(200 * time.Millisecond)
			go func(id string) {
				listFilesRecursive(id, fields, folders, files, wgFolders, wgFiles)
				wgFolders.Done()
			}(id)
		}
	}()
	go func() {
		for f := range files {
			final = append(final, f)
			wgFiles.Done()
		}
	}()
	for _, f := range result {
		wgFiles.Add(1)
		files <- f
		if IsFolder(f) {
			wgFolders.Add(1)
			folders <- f.Id
		}
	}
	wgFolders.Wait()
	close(folders)
	wgFiles.Wait()
	close(files)
	return final, nil
}