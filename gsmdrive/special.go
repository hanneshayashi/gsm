/*
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
	"iter"
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
	newFolder, err := CreateFile(f, nil, false, false, false, "", "", "", "id,mimeType,name")
	if err != nil {
		return nil, err
	}
	return newFolder, nil
}

// CopyFoldersAndReturnFilesWithNewParents creates copy of each folder in the supplied channel,
// adds the new ID to the parents property of the files in the source folder and returns the files in a channel.
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
	items := ListFilesRecursive(folderID, "files(id,parents,mimeType,name),nextPageToken", excludeFolders, false, threads)
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

// ListFilesRecursive lists all files and folders in a parent folder recursively.
//
// The implementation uses a mutex-protected folder queue (instead of a channel)
// to avoid deadlocks that occur when workers block on a bounded folders channel
// while the files channel is also full.
func ListFilesRecursive(id, fields string, excludeFolders []string, includeRoot bool, threads int) <-chan *drive.File {
	files := make(chan *drive.File, threads)

	go func() {
		defer close(files)

		if includeRoot {
			root, err := GetFile(id, "*", "")
			if err != nil {
				log.Println(err)
			} else {
				files <- root
			}
		}

		// Mutex-protected folder queue — no bounded channel, no deadlock.
		var mu sync.Mutex
		queue := []string{id}

		sem := make(chan struct{}, threads) // semaphore for bounded parallelism

		// BFS level-by-level: dispatch all current folders, wait for workers to
		// finish, then check if new folders were discovered. This avoids the race
		// where the queue appears empty while workers are still in flight.
		for len(queue) > 0 {
			// Snapshot current queue and reset it—workers will append to it.
			mu.Lock()
			batch := queue
			queue = nil
			mu.Unlock()

			var wg sync.WaitGroup
			for _, folderID := range batch {
				wg.Add(1)
				sem <- struct{}{} // acquire semaphore slot

				go func(fid string) {
					defer func() {
						<-sem // release semaphore slot
						wg.Done()
					}()
					var discovered []string
					for f, err := range ListFiles(fmt.Sprintf("'%s' in parents and trashed = false", fid), "", "allDrives", "", "", "", fields, true) {
						if err != nil {
							log.Println(err)
							continue
						}
						if isFolder(f) {
							if !gsmhelpers.Contains(f.Id, excludeFolders) {
								files <- f
								discovered = append(discovered, f.Id)
							}
						} else {
							files <- f
						}
					}
					if len(discovered) > 0 {
						mu.Lock()
						queue = append(queue, discovered...)
						mu.Unlock()
					}
				}(folderID)
			}
			wg.Wait()
		}
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
		permissionFound := false
		for p, err := range ListPermissions(fileID, "", "permissions(emailAddress,id)", flags["useDomainAdminAccess"].GetBool()) {
			if err != nil {
				return "", err
			}
			if strings.ToLower(p.EmailAddress) == emailAddress {
				permissionID = p.Id
				permissionFound = true
				break
			}
		}
		if !permissionFound {
			return "", fmt.Errorf("can't find a matching rule for the specified trustee")
		}
	} else {
		domain := strings.ToLower(flags["domain"].GetString())
		permissionFound := false
		for p, err := range ListPermissions(fileID, "", "permissions(domain,id)", flags["useDomainAdminAccess"].GetBool()) {
			if err != nil {
				return "", err
			}
			if strings.ToLower(p.Domain) == domain {
				permissionID = p.Id
				permissionFound = true
				break
			}
		}
		if !permissionFound {
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

// CountFilesAndFolders returns the number of files in an iterator and their size
func CountFilesAndFolders(files iter.Seq2[*drive.File, error]) (folderSize *FolderSize) {
	folderSize = &FolderSize{}
	for f, err := range files {
		if err != nil {
			log.Println(err)
			continue
		}
		if isFolder(f) {
			folderSize.Folders++
		} else {
			folderSize.Files++
			folderSize.Size += f.Size
		}
	}
	return folderSize
}
