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
	"io"
	"net/http"
	"os"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// DeleteRevision permanently deletes a file version.
// You can only delete revisions for files with binary content in Google Drive, like images or videos.
// Revisions for other files, like Google Docs or Sheets, and the last remaining file version can't be deleted.
func DeleteRevision(fileID, revisionID string) (bool, error) {
	srv := getRevisionsService()
	c := srv.Delete(fileID, revisionID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(fileID, revisionID), func() error {
		return c.Do()
	})
	return result, err
}

// GetRevision gets a revision's metadata or content by ID.
func GetRevision(fileID, revisionID, fields string) (*drive.Revision, error) {
	srv := getRevisionsService()
	c := srv.Get(fileID, revisionID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID, revisionID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.Revision)
	return r, nil
}

// DownloadRevision downloads a file revision from drive
func DownloadRevision(fileID, revisionID string, acknowledgeAbuse bool) (string, error) {
	srv := getRevisionsService()
	file, err := GetRevision(fileID, revisionID, "")
	if err != nil {
		return "", err
	}
	c := srv.Get(fileID, revisionID).AcknowledgeAbuse(acknowledgeAbuse)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID, revisionID), func() (interface{}, error) {
		return c.Download()
	})
	if err != nil {
		return "", err
	}
	r, _ := result.(*http.Response)
	defer r.Body.Close()
	fileLocal, err := os.Create(file.OriginalFilename)
	if err != nil {
		return "", err
	}
	defer fileLocal.Close()
	_, err = io.Copy(fileLocal, r.Body)
	return file.OriginalFilename, err
}

func makeListRevisionsCallAndAppend(c *drive.RevisionsListCall, revisions []*drive.Revision, errKey string) ([]*drive.Revision, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.RevisionList)
	revisions = append(revisions, r.Revisions...)
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		revisions, err = makeListRevisionsCallAndAppend(c, revisions, errKey)
	}
	return revisions, err
}

// ListRevisions lists a file's revisions.
func ListRevisions(fileID, fields string) ([]*drive.Revision, error) {
	srv := getRevisionsService()
	c := srv.List(fileID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var revisions []*drive.Revision
	revisions, err := makeListRevisionsCallAndAppend(c, revisions, gsmhelpers.FormatErrorKey(fileID))
	return revisions, err
}

// UpdateRevision updates a revision with patch semantics.
func UpdateRevision(fileID, revisionID, fields string, revision *drive.Revision) (*drive.Revision, error) {
	srv := getRevisionsService()
	c := srv.Update(fileID, revisionID, revision)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID, revisionID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.Revision)
	return r, nil
}
