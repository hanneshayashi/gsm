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
	"gsm/gsmhelpers"
	"io"
	"net/http"
	"os"

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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.File)
	return r, nil
}

// CreateFile creates a new file.
func CreateFile(file *drive.File, content *os.File, ignoreDefaultVisibility, keepRevisionForever, useContentAsIndexableText bool, includePermissionsForView, ocrLanguage, fields string) (*drive.File, error) {
	srv := getFilesService()
	if content == nil {
		file.MimeType = "application/vnd.google-apps.folder"
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(file.Name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.File)
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

// ExportFile exports a Google Doc to the requested MIME type and returns the exported content.
// Please note that the exported content is limited to 10MB.
func ExportFile(fileID, mimeType string) (string, error) {
	srv := getFilesService()
	file, err := GetFile(fileID, "name", "")
	if err != nil {
		return "", err
	}
	c := srv.Export(fileID, mimeType)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (interface{}, error) {
		return c.Download()
	})
	if err != nil {
		return "", err
	}
	r, _ := result.(*http.Response)
	defer r.Body.Close()
	fileLocal, err := os.Create(file.Name)
	if err != nil {
		return "", err
	}
	defer fileLocal.Close()
	_, err = io.Copy(fileLocal, r.Body)
	return file.Name, err // TODO: Extension based on mimeType
}

// GenerateFileIDs generates a set of file IDs which can be provided in create or copy requests.
func GenerateFileIDs(count int64, space string) ([]string, error) {
	srv := getFilesService()
	c := srv.GenerateIds().Count(count)
	if space != "" {
		c = c.Space(space)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey("Generate File Ids"), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.GeneratedIds)
	return r.Ids, nil
}

// DownloadFile downloads a file from drive
func DownloadFile(fileID string, acknowledgeAbuse bool) (string, error) {
	srv := getFilesService()
	file, err := GetFile(fileID, "name", "")
	if err != nil {
		return "", err
	}
	c := srv.Get(fileID).SupportsAllDrives(true).AcknowledgeAbuse(acknowledgeAbuse)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (interface{}, error) {
		return c.Download()
	})
	if err != nil {
		return "", err
	}
	r, _ := result.(*http.Response)
	defer r.Body.Close()
	fileLocal, err := os.Create(file.Name)
	if err != nil {
		return "", err
	}
	defer fileLocal.Close()
	_, err = io.Copy(fileLocal, r.Body)
	return file.Name, err
}

// RandomFile gets a file's metadata or content by ID.
// func RandomFile(fileID, fields, includePermissionsForView string) (*drive.File, error) {
// 	randomFile := func() (*drive.File, error) {
// 		r := &drive.File{
// 			Id: fileID,
// 		}
// 		err := &googleapi.Error{}
// 		rand.Seed(time.Now().UnixNano())
// 		random := rand.Intn(10)
// 		if random%2 == 0 {
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
// 		if random%3 == 0 {
// 			err.Message = "File not found"
// 			err.Code = 404
// 			return nil, err
// 		}
// 		return r, nil
// 	}
// 	var err error
// 	var result *drive.File
// 	errKey := fmt.Sprintf("%s", fileID)
// 	operation := func() error {
// 		result, err = randomFile()
// 		if gsmhelpers.RetryLog(err, errKey) {
// 			return err
// 		}
// 		return nil
// 	}
// 	gsmhelpers.StandarRetrier.Run(operation)
// 	if err != nil {
// 		return nil, gsmhelpers.FormatError(err, errKey)
// 	}
// 	return result, nil
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.File)
	return r, nil
}

func makeListFilesCallAndAppend(c *drive.FilesListCall, files []*drive.File, errKey string) ([]*drive.File, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, gsmhelpers.FormatError(err, errKey)
	}
	r, _ := result.(*drive.FileList)
	for _, f := range r.Files {
		files = append(files, f)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		files, err = makeListFilesCallAndAppend(c, files, errKey)
	}
	return files, err
}

// ListFiles lists or searches files.
// This method accepts the q parameter, which is a search query combining one or more search terms.
// For more information, see https://developers.google.com/drive/api/v3/search-files.
func ListFiles(q, driveID, corpora, includePermissionsForView, orderBy, spaces, fields string, includeItemsFromAllDrives bool) ([]*drive.File, error) {
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
	var files []*drive.File
	files, err := makeListFilesCallAndAppend(c, files, gsmhelpers.FormatErrorKey("List files"))
	return files, err
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.File)
	return r, nil
}
