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
	"gsm/gsmhelpers"
	"log"
	"sync"
	"time"

	"github.com/flowchartsman/retry"
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

type parentChildren struct {
	Parent   string
	Children []*drive.File
}

type parent struct {
	Parent string
	Folder *drive.File
}

func folder(folder *drive.File, destination, driveID, fields string, pc chan parentChildren, wg *sync.WaitGroup, retrier *retry.Retrier) {
	file := &drive.File{
		MimeType: "application/vnd.google-apps.folder",
		DriveId:  driveID,
		Parents:  []string{destination},
		Name:     folder.Name,
	}
	f := &drive.File{}
	operation := func() error {
		newFile, err := CreateFile(file, nil, false, false, false, "", "", fields)
		if err != nil {
			gerr := err.(*googleapi.Error)
			if gerr.Code == 403 {
				log.Println(err)
				return err
			}
			log.Println(err)
			return nil
		}
		f = newFile
		return nil
	}
	err := retrier.Run(operation)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(folder.Name, "'s new id is", f.Id)
	operation = func() error {
		ci, err := ListFiles(fmt.Sprintf("'%s' in parents", folder.Id), "", "", "", "", "", fields, false)
		if err != nil {
			gerr := err.(*googleapi.Error)
			if gerr.Code == 403 {
				log.Println(err)
				return err
			}
			log.Println(err)
			return nil
		}
		if len(ci) > 0 {
			wg.Add(1)
			pc <- parentChildren{Parent: f.Id, Children: ci}
		}
		return nil
	}
	err = retrier.Run(operation)
	if err != nil {
		log.Fatal(err)
	}
}

// Migrate migrates a folder to a drive
func Migrate(file *drive.File, destination, driveID string) {
	fields := "*"
	pc := make(chan parentChildren, 1000)
	folders := make(chan parent, 1000)
	var wgFolders sync.WaitGroup
	var wgPc sync.WaitGroup
	var wgGor sync.WaitGroup
	wgFolders.Add(1)
	wgGor.Add(1)
	folders <- parent{Parent: destination, Folder: file}
	retrier := retry.NewRetrier(10, 250*time.Millisecond, 60*time.Second)
	go func() {
		for f := range folders {
			folder(f.Folder, f.Parent, driveID, fields, pc, &wgPc, retrier)
			wgFolders.Done()
		}
	}()
	wgFolders.Wait()
	for i := 0; i < gsmhelpers.MaxThreads(10); i++ {
		go func(i int) {
			var err error
			fmt.Println("staring", i, len(pc))
			for p := range pc {
				log.Println(i, "is moving", len(p.Children), "children to", p.Parent)
				for _, c := range p.Children {
					if c.MimeType == "application/vnd.google-apps.folder" {
						wgFolders.Add(1)
						folders <- parent{Parent: p.Parent, Folder: c}
					} else {
						operation := func() error {
							u := &drive.File{}
							_, err := UpdateFile(c.Id, p.Parent, c.Parents[0], "", "", "", u, nil, false, false)
							if err != nil {
								gerr := err.(*googleapi.Error)
								if gerr.Code == 403 {
									log.Println(err)
									return err
								}
								log.Println(err)
								return nil
							}
							return nil
						}
						err = retrier.Run(operation)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
				log.Println(i, "has moved", len(p.Children), "children to", p.Parent)
				wgPc.Done()
			}
		}(i)
	}
	wgGor.Done()
	wgGor.Wait()
	wgFolders.Wait()
	wgPc.Wait()
	for len(pc) > 0 || len(folders) > 0 {
		wgFolders.Wait()
		wgPc.Wait()
		time.Sleep(100 * time.Millisecond)
	}
	close(folders)
	close(pc)
}
