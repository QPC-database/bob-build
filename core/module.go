/*
 * Copyright 2019 Arm Limited.
 * SPDX-License-Identifier: Apache-2.0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Declare some minimal interfaces that Bob uses so that both Soong
// and Blueprint can call into the same functions. Note that if a Bob
// function uses these types, they can't necessarily call back into
// Blueprint or Soong

package core

import (
	"text/scanner"

	"github.com/google/blueprint/pathtools"
)

// Common functions in blueprint.BaseModuleContext and android.BaseModuleContext
type commonModuleContext interface {
	ModuleName() string
	ModuleDir() string

	ContainsProperty(name string) bool
	Errorf(pos scanner.Position, fmt string, args ...interface{})
	ModuleErrorf(fmt string, args ...interface{})
	PropertyErrorf(property, fmt string, args ...interface{})
	Failed() bool

	GlobWithDeps(pattern string, excludes []string) ([]string, error)

	Fs() pathtools.FileSystem
	AddNinjaFileDeps(deps ...string)
}
