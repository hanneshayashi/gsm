/*
Copyright © 2020-2022 Hannes Hayashi

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
	"context"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// CopyFile creates a copy of a file and applies any requested updates with patch semantics. Folders cannot be copied.
func CopyFile(fileID, includePermissionsForView, ocrLanguage, fields string, file *drive.File, ignoreDefaultVisibility, keepRevisionForever bool) (*drive.File, error) {
	srv := getFilesService()
	c := srv.Copy(fileID, file).SupportsAllDrives(true).IgnoreDefaultVisibility(ignoreDefaultVisibility).KeepRevisionForever(keepRevisionForever)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if includePermissionsForView != "" {
		c = c.IncludePermissionsForView(includePermissionsForView)
	}
	if ocrLanguage != "" {
		c = c.OcrLanguage(ocrLanguage)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.File)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// CreateFile creates a new file.
func CreateFile(file *drive.File, content *os.File, ignoreDefaultVisibility, keepRevisionForever, useContentAsIndexableText bool, includePermissionsForView, ocrLanguage, fields string) (*drive.File, error) {
	srv := getFilesService()
	if content == nil {
		file.MimeType = folderMimetype
	}
	c := srv.Create(file).SupportsAllDrives(true).IgnoreDefaultVisibility(ignoreDefaultVisibility).KeepRevisionForever(keepRevisionForever).UseContentAsIndexableText(useContentAsIndexableText)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if content != nil {
		c = c.Media(content)
	}
	if includePermissionsForView != "" {
		c = c.IncludePermissionsForView(includePermissionsForView)
	}
	if ocrLanguage != "" {
		c = c.OcrLanguage(ocrLanguage)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(file.Name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.File)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// DeleteFile permanently deletes a file owned by the user without moving it to the trash.
// If the file belongs to a shared drive the user must be an organizer on the parent.
// If the target is a folder, all descendants owned by the user are also deleted.
func DeleteFile(fileID string) (bool, error) {
	srv := getFilesService()
	c := srv.Delete(fileID).SupportsAllDrives(true)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(fileID), func() error {
		return c.Do()
	})
	return result, err
}

// EmptyTrash permanently deletes all of the user's trashed files.
func EmptyTrash() (bool, error) {
	srv := getFilesService()
	c := srv.EmptyTrash()
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey("Empty Trash"), func() error {
		return c.Do()
	})
	return result, err
}

func getLocalFilePaths(localFilePath string) (folder string, fileName string, err error) {
	if localFilePath != "" {
		stats, err := os.Stat(localFilePath)
		if os.IsNotExist(err) {
			if strings.HasSuffix(localFilePath, string(filepath.Separator)) {
				err = os.MkdirAll(localFilePath, 0777)
				if err != nil {
					return "", "", err
				}
				return localFilePath, "", nil
			}
			dir := filepath.Dir(localFilePath)
			err = os.MkdirAll(dir, 0777)
			if err != nil {
				return "", "", err
			}
			return dir + string(filepath.Separator), filepath.Base(localFilePath), nil
		} else if err != nil {
			return "", "", err
		} else {
			if stats.IsDir() {
				if !strings.HasSuffix(localFilePath, string(filepath.Separator)) {
					localFilePath = localFilePath + string(filepath.Separator)
				}
				return localFilePath, "", nil
			}
			return filepath.Dir(localFilePath) + string(filepath.Separator), filepath.Base(localFilePath), nil
		}
	}
	return "", "", nil
}

// ExportFile exports a Google Doc to the requested MIME type and returns the exported content.
// Please note that the exported content is limited to 10MB.
func ExportFile(fileID, mimeType, localFilePath string) (string, error) {
	srv := getFilesService()
	file, err := GetFile(fileID, "name", "")
	if err != nil {
		return "", err
	}
	c := srv.Export(fileID, mimeType)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (any, error) {
		return c.Download()
	})
	if err != nil {
		return "", err
	}
	r, ok := result.(*http.Response)
	if !ok {
		return "", fmt.Errorf("result unknown")
	}
	defer r.Body.Close()
	folder, fileName, err := getLocalFilePaths(localFilePath)
	if err != nil {
		return "", err
	}
	if fileName == "" {
		fileName = file.Name
		extensions, er := mime.ExtensionsByType(mimeType)
		if er == nil && len(extensions) > 0 {
			fileName = fileName + extensions[0]
		}
	}
	fileName = folder + fileName
	fileLocal, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer fileLocal.Close()
	_, err = io.Copy(fileLocal, r.Body)
	return fileName, err
}

// GenerateFileIDs generates a set of file IDs which can be provided in create or copy requests.
func GenerateFileIDs(count int64, space string) ([]string, error) {
	srv := getFilesService()
	c := srv.GenerateIds().Count(count)
	if space != "" {
		c = c.Space(space)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey("Generate File Ids"), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.GeneratedIds)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r.Ids, nil
}

// DownloadFile downloads a file from drive
func DownloadFile(fileID, localFilePath string, acknowledgeAbuse bool) (string, error) {
	srv := getFilesService()
	file, err := GetFile(fileID, "name", "")
	if err != nil {
		return "", err
	}
	c := srv.Get(fileID).SupportsAllDrives(true).AcknowledgeAbuse(acknowledgeAbuse)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (any, error) {
		return c.Download()
	})
	if err != nil {
		return "", err
	}
	r, ok := result.(*http.Response)
	if !ok {
		return "", fmt.Errorf("result unknown")
	}
	defer r.Body.Close()
	folder, fileName, err := getLocalFilePaths(localFilePath)
	if err != nil {
		return "", err
	}
	if fileName == "" {
		fileName = file.Name
	}
	fileName = folder + fileName
	fileLocal, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer fileLocal.Close()
	_, err = io.Copy(fileLocal, r.Body)
	return fileName, err
}

// RandomFile gets a file's metadata or content by ID.
// func RandomFile(fileID, fields, includePermissionsForView string) (*drive.File, error) {
// 	c := func() (interface{}, error) {
// 		r := &drive.File{
// 			Id: fileID,
// 		}
// 		err := &googleapi.Error{}
// 		random := rand.Intn(100)
// 		if random%99 == 0 {
// 			foo := []string{
// 				"Rate limit reached",
// 				"Quota exceeded",
// 				"Forbidden",
// 			}
// 			rand2 := rand.Intn(3)
// 			err.Message = foo[rand2]
// 			err.Code = 403
// 			return nil, err
// 		}
// 		if random%3 == 20 {
// 			err.Message = "File not found"
// 			err.Code = 404
// 			return nil, err
// 		}
// 		return r, nil
// 	}
// 	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), c)
// 	if err != nil {
// 		return nil, err
// 	}
// 	r, ok := result.(*drive.File)
// if !ok {
// 	return nil, fmt.Errorf("result unknown")
// }
// 	return r, nil
// }

// GetFile gets a file's metadata or content by ID.
func GetFile(fileID, fields, includePermissionsForView string) (*drive.File, error) {
	srv := getFilesService()
	c := srv.Get(fileID).SupportsAllDrives(true)
	if includePermissionsForView != "" {
		c.IncludePermissionsForView(includePermissionsForView)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.File)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListFiles lists or searches files.
// This method accepts the q parameter, which is a search query combining one or more search terms.
// For more information, see https://developers.google.com/drive/api/v3/search-files.
func ListFiles(q, driveID, corpora, includePermissionsForView, orderBy, spaces, fields string, includeItemsFromAllDrives bool, cap int) (<-chan *drive.File, <-chan error) {
	srv := getFilesService()
	c := srv.List().SupportsAllDrives(true).IncludeItemsFromAllDrives(includeItemsFromAllDrives).PageSize(1000)
	if q != "" {
		c = c.Q(q)
	}
	if driveID != "" {
		c = c.DriveId(driveID)
	}
	if corpora != "" {
		c = c.Corpora(corpora)
	}
	if includePermissionsForView != "" {
		c = c.IncludePermissionsForView(includePermissionsForView)
	}
	if orderBy != "" {
		c = c.OrderBy(orderBy)
	}
	if spaces != "" {
		c = c.Spaces(spaces)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *drive.File, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *drive.FileList) error {
			for i := range response.Files {
				ch <- response.Files[i]
			}
			return nil
		})
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// UpdateFile updates a file's metadata and/or content. This method supports patch semantics.
func UpdateFile(fileID, addParents, removeParents, includePermissionsForView, ocrLanguage, fields string, file *drive.File, content *os.File, keepRevisionForever, useContentAsIndexableText bool) (*drive.File, error) {
	srv := getFilesService()
	c := srv.Update(fileID, file).SupportsAllDrives(true).KeepRevisionForever(keepRevisionForever).UseContentAsIndexableText(useContentAsIndexableText)
	if addParents != "" && removeParents != "" {
		c = c.AddParents(addParents).RemoveParents(removeParents)
	}
	if content != nil {
		c = c.Media(content)
	}
	if includePermissionsForView != "" {
		c = c.IncludePermissionsForView(includePermissionsForView)
	}
	if ocrLanguage != "" {
		c = c.OcrLanguage(ocrLanguage)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.File)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListLabels lists the labels on a file.
func ListLabels(fileID, fields string, cap int) (<-chan *drive.Label, <-chan error) {
	srv := getFilesService()
	c := srv.ListLabels(fileID).MaxResults(100)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *drive.Label, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *drive.LabelList) error {
			for i := range response.Labels {
				ch <- response.Labels[i]
			}
			return nil
		})
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// ModifyLabels modifies the set of labels on a file.
func ModifyLabels(fileID, fields string, modifyLabelsRequest *drive.ModifyLabelsRequest) ([]*drive.Label, error) {
	srv := getFilesService()
	c := srv.ModifyLabels(fileID, modifyLabelsRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.ModifyLabelsResponse)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r.ModifiedLabels, nil
}
