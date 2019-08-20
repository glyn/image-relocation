/*
 * Copyright (c) 2019-Present Pivotal Software, Inc. All rights reserved.
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

package bundle

import (
	"github.com/pivotal/image-relocation/pkg/image"
	"gopkg.in/yaml.v2"
)

// Bundle represents a set of image names.
type Bundle interface {
	// Add adds an image name to the bundle.
	Add(image.Name)

	// Images returns a slice of the bundle's image names. The slice contains no duplicates.
	Images() []image.Name

	// Export stores the bundle's images as an OCI image layout at the given path on the file system.
	Export(path string) error

	// TODO: do we need Import too?

	// TODO: do we need Push to push a bundle to a repository
	// Push(repo image.Name) err

	// Relocate pushes the bundle's images to repositories with the given prefix and returns a relocation mapping of the original image names to the relocated names.
	Relocate(repoPrefix string) (RelocationMapping, error)
}

// RelocationMapping is a mapping from original image name to relocated image name.
type RelocationMapping map[image.Name]image.Name

type bundle struct {
	images []image.Name `json:"images,omitempty" yaml:"images,omitempty"`
}

func (b *bundle) Add(im image.Name) {
	for _, i := range b.images {
		if i == im {
			return
		}
	}
	b.images = append(b.images, im)
}

// FromYaml marshals a bundle from a YAML sequence of images.
func FromYaml(in []byte) (*bundle, error) {
	var images []string
	if err := yaml.Unmarshal(in, &images); err != nil {
		return nil, err
	}
	b := bundle {
		images: []image.Name{},
	}
	for _, im := range images {
		i, err := image.NewName(im)
		if err != nil {
			return nil, err
		}
		b.Add(i)
	}
	return &b, nil
}

// TODO: do we need Pull to pull a bundle from a repository
// Pull(image.Name)(Bundle, RelocationMapping, err)

func (b *bundle) Images() []image.Name {
	return append([]image.Name(nil), b.images...)
}

func (*bundle) Export(path string) error {
	panic("implement me")
}

func (*bundle) Relocate(repoPrefix string) (RelocationMapping, error) {
	panic("implement me")
}
