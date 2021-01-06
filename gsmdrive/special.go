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
	"strings"
	"sync"

	"github.com/hanneshayashi/gsm/gsmhelpers"
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

// GetPermissionID returns the permissionId from a flag set if either the permissionId itself, or the emailAddress is set.
// Otherwise, it will return an error.
func GetPermissionID(flags map[string]*gsmhelpers.Value) (string, error) {
	set := 0
	possibleFlags := []string{
		"permissionId",
		"emailAddress",
		"domain",
	}
	for _, pf := range possibleFlags {
		if flags[pf].IsSet() {
			set++
		}
	}
	if set != 1 {
		return "", fmt.Errorf("Exactly one of %s must be set", strings.Join(possibleFlags, ", "))
	}
	if flags["permissionId"].IsSet() {
		return flags["permissionId"].GetString(), nil
	}
	var permissionID string
	var fileID string
	if flags["folderId"].IsSet() {
		fileID = flags["folderId"].GetString()
	} else {
		fileID = flags["fileId"].GetString()
	}
	if flags["emailAddress"].IsSet() {
		emailAddress := strings.ToLower(flags["emailAddress"].GetString())
		permissions, err := ListPermissions(fileID, "", "permissions(emailAddress,id)", flags["useDomainAdminAccess"].GetBool())
		if err != nil {
			return "", err
		}
		pFound := false
		for _, p := range permissions {
			if strings.ToLower(p.EmailAddress) == emailAddress {
				permissionID = p.Id
				pFound = true
				break
			}
		}
		if !pFound {
			return "", fmt.Errorf("Can't find a matching rule for the specified trustee")
		}
	} else {
		domain := strings.ToLower(flags["domain"].GetString())
		permissions, err := ListPermissions(fileID, "", "permissions(domain,id)", flags["useDomainAdminAccess"].GetBool())
		if err != nil {
			return "", err
		}
		pFound := false
		for _, p := range permissions {
			if strings.ToLower(p.Domain) == domain {
				permissionID = p.Id
				pFound = true
				break
			}
		}
		if !pFound {
			return "", fmt.Errorf("Can't find a matching rule for the specified trustee")
		}
	}
	return permissionID, nil
}

// GetFolder returns the file if it can be found AND is a folder, otherwise, it returns an error
func GetFolder(folderID string) (*drive.File, error) {
	folder, err := GetFile(folderID, "id,mimeType", "")
	if err != nil {
		return nil, err
	}
	if !IsFolder(folder) {
		return nil, fmt.Errorf("%s is not a folder", folderID)
	}
	return folder, nil
}
