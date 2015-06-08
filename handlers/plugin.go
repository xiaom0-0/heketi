//
// Copyright (c) 2014 The heketi Authors
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

package handlers

// Volume interface for plugins
type Plugin interface {
	VolumeCreate(v *VolumeCreateRequest) (*VolumeInfoResp, error)
	VolumeDelete(id uint64) error
	VolumeInfo(id uint64) (*VolumeInfoResp, error)
	VolumeResize(id uint64) (*VolumeInfoResp, error)
	VolumeList() (*VolumeListResponse, error)

	NodeAdd(v *NodeAddRequest) (*NodeInfoResp, error)
	NodeRemove(id uint64) error
	NodeInfo(id uint64) (*NodeInfoResp, error)
	NodeList() (*NodeListResponse, error)
}
