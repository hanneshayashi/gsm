/*
Package gsmdrive implements the Drive API
Copyright © 2020-2021 Hannes Hayashi

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
	"log"
	"strings"
	"sync"

	"github.com/hanneshayashi/gsm/gsmhelpers"
	drive "google.golang.org/api/drive/v3"
)

const folderMimetype = "application/vnd.google-apps.folder"

// isFolder returns true if the file object is a folder, otherwise false
// Make sure that the MimeType property is actually set.
func isFolder(f *drive.File) bool {
	if f.MimeType == folderMimetype {
		return true
	}
	return false
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

// CopyFoldersAndReturnFilesWithNewParents creates copy of each folder in the supplied channel,
// adds the new ID to the parents propertie of the files in the source folder and returns the files in a channel.
func CopyFoldersAndReturnFilesWithNewParents(folderID, destination string, results chan *drive.File, threads int) (<-chan *drive.File, error) {
	root, err := GetFolder(folderID)
	if err != nil {
		return nil, fmt.Errorf("Error getting folder: %v", err)
	}
	folderMap := make(map[string]string)
	newRoot, err := createFolder(destination, root.Name)
	if err != nil {
		return nil, err
	}
	folderMap[root.Id] = newRoot.Id
	items := ListFilesRecursive(folderID, "files(id,parents,mimeType,name),nextPageToken", threads)
	files := make(chan *drive.File, threads)
	go func() {
		for i := range items {
			if isFolder(i) {
				newF, err := createFolder(folderMap[i.Parents[0]], i.Name)
				if err != nil {
					log.Println(err)
				} else {
					results <- newF
					folderMap[i.Id] = newF.Id
				}
			} else {
				i.Parents = append(i.Parents, folderMap[i.Parents[0]])
				files <- i
			}
		}
		close(files)
	}()
	return files, nil
}

func listFilesRecursive(id, fields string, folders chan string, files chan *drive.File, wg *sync.WaitGroup, cap int) {
	result, err := ListFiles(fmt.Sprintf("'%s' in parents and trashed = false", id), "", "allDrives", "", "", "", fields, true, cap)
	wg.Add(1)
	go func() {
		for f := range result {
			files <- f
			if isFolder(f) {
				wg.Add(1)
				folders <- f.Id
			}
		}
		wg.Done()
	}()
	go func() {
		for e := range err {
			log.Println(e)
		}
	}()
}

// ListFilesRecursive lists all files and foldes in a parent folder recursively
func ListFilesRecursive(id, fields string, threads int) <-chan *drive.File {
	wg := &sync.WaitGroup{}
	folders := make(chan string, threads)
	files := make(chan *drive.File, threads)
	wg.Add(1)
	folders <- id
	go func() {
		for i := 0; i < threads; i++ {
			go func() {
				for id := range folders {
					listFilesRecursive(id, fields, folders, files, wg, threads)
					wg.Done()
				}
			}()
		}
	}()
	go func() {
		wg.Wait()
		close(folders)
		close(files)
	}()
	return files
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
		permissions, err := ListPermissions(fileID, "", "permissions(emailAddress,id)", flags["useDomainAdminAccess"].GetBool(), gsmhelpers.MaxThreads(0))
		pFound := false
		for p := range permissions {
			if strings.ToLower(p.EmailAddress) == emailAddress {
				permissionID = p.Id
				pFound = true
				break
			}
		}
		e := <-err
		if e != nil {
			return "", e
		}
		if !pFound {
			return "", fmt.Errorf("Can't find a matching rule for the specified trustee")
		}
	} else {
		domain := strings.ToLower(flags["domain"].GetString())
		permissions, err := ListPermissions(fileID, "", "permissions(domain,id)", flags["useDomainAdminAccess"].GetBool(), gsmhelpers.MaxThreads(0))
		pFound := false
		for p := range permissions {
			if strings.ToLower(p.Domain) == domain {
				permissionID = p.Id
				pFound = true
				break
			}
		}
		e := <-err
		if e != nil {
			return "", e
		}
		if !pFound {
			return "", fmt.Errorf("Can't find a matching rule for the specified trustee")
		}
	}
	return permissionID, nil
}

// GetFolder returns the file if it can be found AND is a folder, otherwise, it returns an error
func GetFolder(folderID string) (*drive.File, error) {
	folder, err := GetFile(folderID, "id,name,mimeType,parents", "")
	if err != nil {
		return nil, err
	}
	if !isFolder(folder) {
		return nil, fmt.Errorf("%s is not a folder", folderID)
	}
	return folder, nil
}
