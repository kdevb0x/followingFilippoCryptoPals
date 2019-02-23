// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package set2

import (
	"testing"
)

func TestProblem9(t *testing.T) {
	t.Logf("%x", padPKCS7([]byte("YELLOW SUBMARINE"), 16))
	t.Logf("%x", padPKCS7([]byte("YELLOW SUBMARINE"), 20))

}
