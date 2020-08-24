// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package auto

// NamedConfig is a base type which provides config name.
type NamedConfig struct {
	name string
}

// Name implements named interface.
func (cfg *NamedConfig) Name() string {
	return cfg.name
}

// CommandConfig sets up settings for command build.
type CommandConfig struct {
	NamedConfig

	DisableImage bool `yaml:"disableImage"`
}