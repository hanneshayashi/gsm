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
	"io"
	"math/rand"
	"os"

	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// CopyFile creates a copy of a file and applies any requested updates with patch semantics. Folders cannot be copied.
func CopyFile(fileID, includePermissionsForView, ocrLanguage, fields string, file *drive.File, ignoreDefaultVisibility, keepRevisionForever bool) (*drive.File, error) {
	srv := getFilesService()
	c := srv.Copy(fileID, file).EnforceSingleParent(true).SupportsAllDrives(true).IgnoreDefaultVisibility(ignoreDefaultVisibility).KeepRevisionForever(keepRevisionForever)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if includePermissionsForView != "" {
		c = c.IncludePermissionsForView(includePermissionsForView)
	}
	if ocrLanguage != "" {
		c = c.OcrLanguage(ocrLanguage)
	}
	r, err := c.Do()
	return r, err
}

// CreateFile creates a new file.
func CreateFile(file *drive.File, content *os.File, ignoreDefaultVisibility, keepRevisionForever, useContentAsIndexableText bool, includePermissionsForView, ocrLanguage, fields string) (*drive.File, error) {
	srv := getFilesService()
	if content == nil {
		file.MimeType = "application/vnd.google-apps.folder"
	}
	c := srv.Create(file).EnforceSingleParent(true).SupportsAllDrives(true).IgnoreDefaultVisibility(ignoreDefaultVisibility).KeepRevisionForever(keepRevisionForever).UseContentAsIndexableText(useContentAsIndexableText)
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
	r, err := c.Do()
	return r, err
}

// DeleteFile permanently deletes a file owned by the user without moving it to the trash.
// If the file belongs to a shared drive the user must be an organizer on the parent.
// If the target is a folder, all descendants owned by the user are also deleted.
func DeleteFile(fileID string) (bool, error) {
	srv := getFilesService()
	err := srv.Delete(fileID).SupportsAllDrives(true).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// EmptyTrash permanently deletes all of the user's trashed files.
func EmptyTrash() (bool, error) {
	srv := getFilesService()
	err := srv.EmptyTrash().Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// ExportFile exports a Google Doc to the requested MIME type and returns the exported content.
// Please note that the exported content is limited to 10MB.
func ExportFile(fileID, mimeType string) (string, error) {
	srv := getFilesService()
	file, err := GetFile(fileID, "name", "")
	if err != nil {
		return "", err
	}
	r, err := srv.Export(fileID, mimeType).Download()
	if err != nil {
		return "", err
	}
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
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	return r.Ids, nil
}

// DownloadFile downloads a file from drive
func DownloadFile(fileID string, acknowledgeAbuse bool) (string, error) {
	srv := getFilesService()
	file, err := GetFile(fileID, "name", "")
	if err != nil {
		return "", err
	}
	r, err := srv.Get(fileID).SupportsAllDrives(true).AcknowledgeAbuse(acknowledgeAbuse).Download()
	if err != nil {
		return "", err
	}
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
func RandomFile(fileID, fields, includePermissionsForView string) (*drive.File, error) {
	r := &drive.File{
		Id: fileID,
	}
	err := &googleapi.Error{}
	random := rand.Intn(10)
	if random%10 == 0 {
		return r, nil
	}
	if random%3 == 0 {
		err.Code = 404
		return nil, err
	}
	err.Code = 403
	return nil, err
}

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
	r, err := c.Do()
	return r, err
}

func makeListFilesCallAndAppend(c *drive.FilesListCall, files []*drive.File) ([]*drive.File, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, f := range r.Files {
		files = append(files, f)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		files, err = makeListFilesCallAndAppend(c, files)
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
	files, err := makeListFilesCallAndAppend(c, files)
	return files, err
}

// UpdateFile updates a file's metadata and/or content. This method supports patch semantics.
func UpdateFile(key, addParents, removeParents, includePermissionsForView, ocrLanguage, fields string, file *drive.File, content *os.File, keepRevisionForever, useContentAsIndexableText bool) (*drive.File, error) {
	srv := getFilesService()
	c := srv.Update(key, file).SupportsAllDrives(true).EnforceSingleParent(true).KeepRevisionForever(keepRevisionForever).UseContentAsIndexableText(useContentAsIndexableText)
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
	r, err := c.Do()
	return r, err
}
