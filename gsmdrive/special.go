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
	"gsm/gsmhelpers"
	"sync"
	"time"

	drive "google.golang.org/api/drive/v3"
)

type folder struct {
	NewID     string
	OldParent string
	NewParent string
	Name      string
	Root      bool
}

func copyFolders(folderMap map[string]*folder, destination string) error {
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
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{parent},
		Name:     name,
	}
	newFolder, err := CreateFile(f, nil, false, false, false, "", "", "id,mimeType,name")
	if err != nil {
		return nil, err
	}
	return newFolder, nil
}

func getNewID(oldID string, folderMap map[string]*folder) (string, error) {
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

func getFilesAndFolders(folderID string, threads int) (folderMap map[string]*folder, files []*drive.File, err error) {
	folderMap = make(map[string]*folder)
	root, err := GetFile(folderID, "id,mimeType,name,parents", "")
	if err != nil {
		return nil, nil, err
	}
	if root.MimeType != "application/vnd.google-apps.folder" {
		return nil, nil, fmt.Errorf("%s is not a folder", folderID)
	}
	items, err := ListFilesRecursive(folderID, "files(id,parents,mimeType,name),nextPageToken", threads)
	if err != nil {
		return nil, nil, err
	}
	for _, i := range items {
		if i.MimeType == "application/vnd.google-apps.folder" {
			folderMap[i.Id] = &folder{OldParent: i.Parents[0], Name: i.Name}
		} else {
			files = append(files, i)
		}
	}
	folderMap[folderID] = &folder{Name: root.Name, OldParent: root.Parents[0], Root: true}
	return folderMap, files, nil
}

// CopyFolderRecursive recursively copies a folder to a new destination
func CopyFolderRecursive(folderID, destination string, threads int) ([]*drive.File, []error) {
	threads = gsmhelpers.MaxThreads(threads)
	folderMap, files, err := getFilesAndFolders(folderID, threads)
	if err != nil {
		return nil, []error{err}
	}
	filesChan := make(chan *drive.File, threads)
	finalChan := make(chan *drive.File, threads)
	errChan := make(chan error, threads)
	finalErr := []error{}
	final := []*drive.File{}
	wgFiles := &sync.WaitGroup{}
	wgErrors := &sync.WaitGroup{}
	wgFinal := &sync.WaitGroup{}
	folderMap[folderID].NewParent = destination
	wgFiles.Add(1)
	go func() {
		for _, f := range files {
			filesChan <- f
		}
		close(filesChan)
		wgFiles.Done()
	}()
	err = copyFolders(folderMap, "")
	if err != nil {
		return nil, []error{err}
	}
	for i := 0; i < threads; i++ {
		wgFiles.Add(1)
		go func() {
			for f := range filesChan {
				folder := folderMap[f.Parents[0]]
				c, err := CopyFile(f.Id, "", "", "id,name,mimeType,parents", &drive.File{Parents: []string{folder.NewID}, Name: f.Name}, false, false)
				if err != nil {
					errChan <- err
				} else {
					finalChan <- c
				}
				time.Sleep(200 * time.Millisecond)
			}
			wgFiles.Done()
		}()
	}
	wgErrors.Add(1)
	go func() {
		for e := range errChan {
			finalErr = append(finalErr, e)
		}
		wgErrors.Done()
	}()
	wgFinal.Add(1)
	go func() {
		for r := range finalChan {
			final = append(final, r)
		}
		wgFinal.Done()
	}()
	wgFiles.Wait()
	close(errChan)
	close(finalChan)
	wgErrors.Wait()
	wgFinal.Wait()
	return final, finalErr
}

// MoveFolderToSharedDrive migrates a folder to a drive
func MoveFolderToSharedDrive(folderID, destination string, threads int) ([]*drive.File, []error) {
	threads = gsmhelpers.MaxThreads(threads)
	folderMap, files, err := getFilesAndFolders(folderID, threads)
	if err != nil {
		return nil, []error{err}
	}
	filesChan := make(chan *drive.File, threads)
	finalChan := make(chan *drive.File, threads)
	errChan := make(chan error, threads)
	finalErr := []error{}
	final := []*drive.File{}
	wgFiles := &sync.WaitGroup{}
	wgErrors := &sync.WaitGroup{}
	wgFinal := &sync.WaitGroup{}
	folderMap[folderID].NewParent = destination
	wgFiles.Add(1)
	go func() {
		for _, f := range files {
			filesChan <- f
		}
		close(filesChan)
		wgFiles.Done()
	}()
	err = copyFolders(folderMap, "")
	if err != nil {
		return nil, []error{err}
	}
	for i := 0; i < threads; i++ {
		wgFiles.Add(1)
		go func() {
			for f := range filesChan {
				folder := folderMap[f.Parents[0]]
				u, err := UpdateFile(f.Id, folder.NewID, folder.OldParent, "", "", "id", nil, nil, false, false)
				if err != nil {
					errChan <- err
				} else {
					finalChan <- u
				}
				time.Sleep(200 * time.Millisecond)
			}
			wgFiles.Done()
		}()
	}
	wgErrors.Add(1)
	go func() {
		for e := range errChan {
			finalErr = append(finalErr, e)
		}
		wgErrors.Done()
	}()
	wgFinal.Add(1)
	go func() {
		for r := range finalChan {
			final = append(final, r)
		}
		wgFinal.Done()
	}()
	wgFiles.Wait()
	close(errChan)
	close(finalChan)
	wgErrors.Wait()
	wgFinal.Wait()
	return final, finalErr
}

func listFilesRecursive(id, fields string, folders chan string, files chan *drive.File, wgFolders, wgFiles *sync.WaitGroup) error {
	result, err := ListFiles(fmt.Sprintf("'%s' in parents and trashed = false", id), "", "allDrives", "", "", "", fields, true)
	if err != nil {
		return err
	}
	for _, f := range result {
		wgFiles.Add(1)
		files <- f
		if f.MimeType == "application/vnd.google-apps.folder" {
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
		if f.MimeType == "application/vnd.google-apps.folder" {
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

// CreatePermissionRecursive recursively grants permissions on a folder
func CreatePermissionRecursive(fileIds []string, emailMessage, fields string, useDomainAdminAccess, sendNotificationEmail, transferOwnership, moveToNewOwnersRoot bool, permission *drive.Permission, threads int) ([]*drive.Permission, []error) {
	ids := make(chan string, threads)
	results := make(chan *drive.Permission, threads)
	errChan := make(chan error, threads)
	final := []*drive.Permission{}
	finalErr := []error{}
	wgIDs := &sync.WaitGroup{}
	wgPermissions := &sync.WaitGroup{}
	wgErrors := &sync.WaitGroup{}
	wgIDs.Add(1)
	go func() {
		for _, id := range fileIds {
			ids <- id
		}
		wgIDs.Done()
	}()
	for i := 0; i < threads; i++ {
		wgPermissions.Add(1)
		go func() {
			for id := range ids {
				r, err := CreatePermission(id, emailMessage, fields, useDomainAdminAccess, sendNotificationEmail, transferOwnership, moveToNewOwnersRoot, permission)
				if err != nil {
					errChan <- err
				} else {
					results <- r
				}
			}
			wgPermissions.Done()
		}()
	}
	wgErrors.Add(1)
	go func() {
		for e := range errChan {
			finalErr = append(finalErr, e)
		}
		wgErrors.Done()
	}()
	wgErrors.Add(1)
	go func() {
		for r := range results {
			final = append(final, r)
		}
		wgErrors.Done()
	}()
	wgIDs.Wait()
	close(ids)
	wgPermissions.Wait()
	close(errChan)
	close(results)
	wgErrors.Wait()
	return final, finalErr
}
