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

package test

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

//
type TestHelper struct {
	*testing.T
}

//
func NewTestHelper(t *testing.T) *TestHelper {
	return &TestHelper{t}
}

//
func (t *TestHelper) AssertTrue(got bool) {
	t.AssertEqual(true, got)
}

//
func (t *TestHelper) AssertFalse(got bool) {
	t.AssertEqual(false, got)
}

//
func (t *TestHelper) AssertEqual(want, got interface{}) {
	if want != got {
		t.raiseError("want \"%v\", not \"%v\"", want, got)
	}
}

//
func (t *TestHelper) AssertOneOf(want []string, got string) {
	for _, w := range want {
		if w == got {
			return
		}
	}
	t.raiseError("value \"%v\" is not in wanted set \"%v\"", got, want)
}

//
func (t *TestHelper) AssertEqualSlices(want, got []string) {

	e := false

	if len(want) == len(got) {
		for ix := range want {
			if want[ix] != got[ix] {
				e = true
				break
			}
		}
	}

	if e {
		t.raiseError("want \"%v\", not \"%v\"", want, got)
	}
}

//
func (t *TestHelper) AssertNotEqual(want, got interface{}) {
	if want == got {
		t.raiseError("don't want \"%v\"", want)
	}
}

//
func (t *TestHelper) AssertNil(got interface{}) {
	if nil != got {
		t.raiseError("want nil, not \"%v\"", got)
	}
}

//
func (t *TestHelper) AssertNotNil(got interface{}) {
	if nil == got {
		t.raiseError("want non-nil")
	}
}

//
func (t *TestHelper) AssertNoError(e error) {
	if e != nil {
		t.raiseError("don't want error: %v", e)
	}
}

//
func (t *TestHelper) GetFixture(fx string) string {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("cannot determine fixture path")
	}
	p := path.Join(path.Dir(f), "../../../test/fixtures", fx)
	return p
}

//
func (t *TestHelper) raiseError(format string, args ...interface{}) {
	t.Error(fmt.Sprintf("%s%s", caller(), fmt.Sprintf(format, args...)))
}

//
func caller() string {

	// check where we are
	fpcs := make([]uintptr, 1)
	n := runtime.Callers(2, fpcs)

	if n == 0 {
		return "n/a"
	}

	pc := fpcs[0]
	thisFile, thisLine := runtime.FuncForPC(pc).FileLine(pc)
	var file string
	var line int
	ok := true

	// stop at first file in call stack that's not this file
	for skip := 0; ok == true; skip++ {
		pc, file, line, ok = runtime.Caller(skip)
		if file != thisFile {
			break
		}
	}

	if !ok {
		return "n/a"
	}

	// calculate number of required backspaces
	_, thisFile = filepath.Split(thisFile)
	_, file = filepath.Split(file)
	back := len(thisFile) + len(strconv.Itoa(thisLine)) + 3
	return fmt.Sprintf("%s%s:%d\n", strings.Repeat("\b", back), file, line)
}
