/*
Copyright © 2020-2023 Hannes Hayashi

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

package gsmdrivelabels

import (
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/drivelabels/v2"
	"google.golang.org/api/googleapi"
)

// Get the constraints on the structure of a Label; such as, the maximum number of Fields allowed and maximum length of the label title.
func GetLabelLimits(name, fields string) (*drivelabels.GoogleAppsDriveLabelsV2LabelLimits, error) {
	srv := getLimitsService()
	c := srv.GetLabel().Name(name)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drivelabels.GoogleAppsDriveLabelsV2LabelLimits)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
