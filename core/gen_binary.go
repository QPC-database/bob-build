/*
 * Copyright 2018 Arm Limited.
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

package core

import (
	"reflect"

	"github.com/google/blueprint"

	"github.com/ARM-software/bob-build/utils"
)

type generateBinary struct {
	generateLibrary
}

//// Support GenerateLibraryInterface

func (m *generateBinary) libExtension() string {
	return ""
}

//// Support PhonyInterface, DependentInterface

// List of everything generated by this target
func (m *generateBinary) outputs(g generatorBackend) []string {
	out := []string{getLibraryGeneratedPath(m, g)}
	out = append(out, utils.PrefixDirs(m.Properties.Headers, m.outputDir(g))...)
	return out
}

//// Support Installable

func (m *generateBinary) filesToInstall(ctx blueprint.ModuleContext) []string {
	return []string{getLibraryGeneratedPath(m, getBackend(ctx))}
}

//// Support blueprint.Module

func (m *generateBinary) GenerateBuildActions(ctx blueprint.ModuleContext) {
	if isEnabled(m) {
		inouts := inouts(m, ctx)
		getBackend(ctx).genBinaryActions(m, ctx, inouts)
	}
}

//// Factory functions

func genBinaryFactory(config *bobConfig) (blueprint.Module, []interface{}) {
	module := &generateBinary{}
	module.generateCommon.Properties.Features.Init(config.getAvailableFeatures(), reflect.TypeOf(GenerateProps{}))
	return module, []interface{}{
		&module.SimpleName.Properties,
		&module.generateCommon.Properties,
		&module.Properties,
	}
}
