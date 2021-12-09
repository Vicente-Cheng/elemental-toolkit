/*
Copyright © 2021 SUSE LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	. "github.com/onsi/gomega"
	v1mock "github.com/rancher-sandbox/elemental-cli/tests/mocks"
	"github.com/spf13/afero"
	"k8s.io/mount-utils"
	"testing"
)

func TestChroot(t *testing.T) {
	RegisterTestingT(t)
	syscallInterface := &v1mock.FakeSyscall{}
	chroot := NewChroot(
		"/whatever",
		WithRunner(&v1mock.FakeRunner{}),
		WithFS(afero.NewMemMapFs()),
		WithSyscall(syscallInterface),
		WithMounter(&mount.FakeMounter{}),
		)
	defer chroot.Close()
	_, err := chroot.Run( "chroot-command")
	Expect(err).To(BeNil())
	Expect(syscallInterface.WasChrootCalledWith("/whatever")).To(BeTrue())
}
