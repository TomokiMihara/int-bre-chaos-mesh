// Copyright Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by chaos-builder. DO NOT EDIT.

package v1alpha1

// this file tests the coupling with all kinds map and each TemplateType
import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestChaosKindMapShouldContainsAWSChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeAWSChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsAzureChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeAzureChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsBlockChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeBlockChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsDNSChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeDNSChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsGCPChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeGCPChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsHelloWorldChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeHelloWorldChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsHTTPChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeHTTPChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsIOChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeIOChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsJVMChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeJVMChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsKernelChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeKernelChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsNetworkChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeNetworkChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsPhysicalMachineChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypePhysicalMachineChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsPodChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypePodChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsStressChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeStressChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}
func TestChaosKindMapShouldContainsTimeChaos(t *testing.T) {
	g := NewGomegaWithT(t)
	var requiredType TemplateType
	requiredType = TypeTimeChaos

	_, ok := all.kinds[string(requiredType)]
	g.Expect(ok).To(Equal(true), "all kinds map should contains this type", requiredType)
}

