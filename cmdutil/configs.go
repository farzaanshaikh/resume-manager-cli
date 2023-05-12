/*
Copyright © 2023 Farzaan Shaikh fvshaikh93@gmail.com

Use of this source code is governed by a GPL
license that can be found in the LICENSE file.
*/
package cmdutil

import "os"

const (
	ConfigDefaultName       = ".reman"
	ConfigDefaultAuthorName = "Reman"
	VersionKey              = "$version"
	AuthorNameKey           = "author.name"
)

// Checks if config file exists
func ConfigExists() bool {
	if _, err := os.Stat(ConfigDefaultName); err != nil {
		return false
	}

	return true
}
