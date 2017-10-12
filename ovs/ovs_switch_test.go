// Copyright (c) 2017 Che Wei, Lin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var ovsSwitch *OVSSwitch
var bridgeName string = "ovs_cni"

func TestNewOVSSwitch(t *testing.T) {
	var err error
	ovsSwitch, err = NewOVSSwitch(bridgeName)
	assert.NoError(t, err)
}

func TestAddPort(t *testing.T) {
	err := ovsSwitch.addPort("test")
	assert.NoError(t, err)
}

func TestAddPort_Invalid(t *testing.T) {
	err := ovsSwitch.addPort("")
	assert.Error(t, err)
}

func TestDeleteOVSSwitch(t *testing.T) {
	err := ovsSwitch.Delete()
	assert.NoError(t, err)
}

func TestDeleteOVSSwitch_Invalid(t *testing.T) {
	//wait previous delete
	time.Sleep(300 * time.Millisecond)
	err := ovsSwitch.Delete()
	assert.Error(t, err)
}
