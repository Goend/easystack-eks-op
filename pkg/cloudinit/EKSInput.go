package cloudinit

/*
Copyright 2019 The Kubernetes Authors.

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

const (
	EKS = `{{.Header}}
{{- template "disk_setup" .DiskSetup}}
{{- template "fs_setup" .DiskSetup}}
runcmd:
{{- template "commands" .PreKubeadmCommands }}
{{- template "commands" .PostKubeadmCommands }}
{{- template "mounts" .Mounts}}
{{template "files" .WriteFiles}}
{{- template "ntp" .NTP }}
{{- template "users" .Users }}
`
)

// EKSInput defines the context to generate a eks instance user data.
type EKSInput struct {
	BaseUserData
}

// NewEKS returns the user data string to be used on a eks instance.
func NewEKS(input *EKSInput) ([]byte, error) {
	input.Header = cloudConfigHeader
	input.WriteFiles = append(input.WriteFiles, input.AdditionalFiles...)
	userData, err := generate("InitEKS", EKS, input)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
