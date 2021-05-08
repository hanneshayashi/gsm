/*
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

// FolderSize represents the size of a Drive folder or Shared Drive
type FolderSize struct {
	Files   int64 `json:"files,omitempty"`
	Folders int64 `json:"folders,omitempty"`
	Size    int64 `json:"size,omitempty"`
}

// isFolder returns true if the file object is a folder, otherwise false
// Make sure that the MimeType property is actually set.
func isFolder(f *drive.File) bool {
	return f.MimeType == folderMimetype
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
func CopyFoldersAndReturnFilesWithNewParents(folderID, destination string, results chan *drive.File, excludeFolders []string, threads int) (<-chan *drive.File, error) {
	root, err := GetFolder(folderID)
	if err != nil {
		return nil, fmt.Errorf("error getting folder: %v", err)
	}
	folderMap := make(map[string]string)
	newRoot, err := createFolder(destination, root.Name)
	if err != nil {
		return nil, err
	}
	folderMap[root.Id] = newRoot.Id
	items := ListFilesRecursive(folderID, "files(id,parents,mimeType,name),nextPageToken", excludeFolders, threads)
	files := make(chan *drive.File, threads)
	results <- newRoot
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

// ListFilesRecursive lists all files and foldes in a parent folder recursively
func ListFilesRecursive(id, fields string, excludeFolders []string, threads int) <-chan *drive.File {
	wg := &sync.WaitGroup{}
	folders := make(chan string, threads)
	files := make(chan *drive.File, threads)
	wg.Add(1)
	folders <- id
	go func() {
		for i := 0; i < threads; i++ {
			go func() {
				for id := range folders {
					result, err := ListFiles(fmt.Sprintf("'%s' in parents and trashed = false", id), "", "allDrives", "", "", "", fields, true, threads)
					go func() {
						for f := range result {
							if isFolder(f) {
								if !gsmhelpers.Contains(f.Id, excludeFolders) {
									wg.Add(1)
									files <- f
									folders <- f.Id
								}
							} else {
								files <- f
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
	for i := range possibleFlags {
		if flags[possibleFlags[i]].IsSet() {
			set++
		}
	}
	if set != 1 {
		return "", fmt.Errorf("exactly one of %s must be set", strings.Join(possibleFlags, ", "))
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
			return "", fmt.Errorf("can't find a matching rule for the specified trustee")
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
			return "", fmt.Errorf("can't find a matching rule for the specified trustee")
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

// CountFilesAndFolders returns the number of files in a channel and their size
func CountFilesAndFolders(filesCh <-chan *drive.File) (folderSize *FolderSize) {
	folderSize = &FolderSize{}
	for f := range filesCh {
		if isFolder(f) {
			folderSize.Folders++
		} else {
			folderSize.Files++
			folderSize.Size += f.Size
		}
	}
	return folderSize
}
