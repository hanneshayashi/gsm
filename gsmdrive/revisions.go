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
	"os"

	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// DeleteRevision permanently deletes a file version.
// You can only delete revisions for files with binary content in Google Drive, like images or videos.
// Revisions for other files, like Google Docs or Sheets, and the last remaining file version can't be deleted.
func DeleteRevision(fileID, revisionID string) (bool, error) {
	srv := getRevisionsService()
	err := srv.Delete(fileID, revisionID).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetRevision gets a revision's metadata or content by ID.
func GetRevision(fileID, revisionID, fields string) (*drive.Revision, error) {
	srv := getRevisionsService()
	c := srv.Get(fileID, revisionID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// DownloadRevision downloads a file revision from drive
func DownloadRevision(fileID, revisionID string, acknowledgeAbuse bool) (string, error) {
	srv := getRevisionsService()
	file, err := GetRevision(fileID, revisionID, "")
	if err != nil {
		return "", err
	}
	r, err := srv.Get(fileID, revisionID).AcknowledgeAbuse(acknowledgeAbuse).Download()
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	fileLocal, err := os.Create(file.OriginalFilename)
	if err != nil {
		return "", err
	}
	defer fileLocal.Close()
	_, err = io.Copy(fileLocal, r.Body)
	return file.OriginalFilename, err
}

func makeListRevisionsCallAndAppend(c *drive.RevisionsListCall, revisions []*drive.Revision) ([]*drive.Revision, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, p := range r.Revisions {
		revisions = append(revisions, p)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		revisions, err = makeListRevisionsCallAndAppend(c, revisions)
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
	revisions, err := makeListRevisionsCallAndAppend(c, revisions)
	return revisions, err
}

// UpdateRevision updates a revision with patch semantics.
func UpdateRevision(fileID, revisionID, fields string, revision *drive.Revision) (*drive.Revision, error) {
	srv := getRevisionsService()
	c := srv.Update(fileID, revisionID, revision)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}
