/*
	Copyright 2020 Alexander Vollschwitz <xelalex@gmx.net>

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

package main

import (
	"testing"
	"time"

	"github.com/xelalexv/dregsy/internal/pkg/test"
)

//
func TestSkopeo(t *testing.T) {
	th := test.NewTestHelper(t)
	ret := runDregsy(th, 1, 0, "-config="+th.GetFixture("e2e/skopeo.yaml"))
	th.AssertEqual(0, ret)
}

//
func TestDocker(t *testing.T) {
	th := test.NewTestHelper(t)
	ret := runDregsy(th, 1, 0, "-config="+th.GetFixture("e2e/docker.yaml"))
	th.AssertEqual(0, ret)
}

//
func runDregsy(th *test.TestHelper, ticks int, wait time.Duration,
	args ...string) int {

	syncInstance = nil
	inTestRound = true
	testArgs = args

	go func() {
		main()
	}()

	for syncInstance == nil {
	}

	shutdown := (ticks > 0 || wait > 0)

	for i := ticks; i > 0; i-- {
		syncInstance.WaitForTick()
	}

	if wait > 0 {
		time.Sleep(time.Second * wait)
	}

	if syncInstance != nil {
		if shutdown {
			syncInstance.Shutdown()
		}
		for syncInstance != nil {
		}
		return dregsyExitCode
	}

	return -1
}
