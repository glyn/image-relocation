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

package bundle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal/image-relocation/pkg/bundle"
	"github.com/pivotal/image-relocation/pkg/image"
)

var _ = Describe("FromYaml", func() {
	var (
		b   bundle.Bundle
		err error
		in  string
	)

	JustBeforeEach(func() {
		b, err = bundle.FromYaml([]byte(in))
	})

	Context("when the input is a valid sequence of image names", func() {
		BeforeEach(func() {
			in = `- x:v1
- y@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef`
		})

		It("should produce a bundle with the correct image names", func() {
			Expect(err).NotTo(HaveOccurred())
			expected := []image.Name{imageName("x:v1"), imageName("y@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")}
			Expect(b.Images()).To(ConsistOf(expected))
		})
	})

	Context("when the input contains duplicate image names", func() {
		BeforeEach(func() {
			in = `- x
- x
- y
- library/y`
		})

		It("should produce a bundle with no duplicate image names", func() {
			Expect(err).NotTo(HaveOccurred())
			Expect(b.Images()).To(ConsistOf(imageName("x"), imageName("y")))
		})
	})

	// TODO: test error cases
})

func imageName(i string) image.Name {
	img, err := image.NewName(i)
	Expect(err).NotTo(HaveOccurred())
	return img
}
