// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License included
// in the file licenses/BSL.txt and at www.mariadb.com/bsl11.
//
// Change Date: 2022-10-01
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt and at
// https://www.apache.org/licenses/LICENSE-2.0

package sysutil

import (
	"os/exec"
	"testing"
)

func TestExitStatus(t *testing.T) {
	cmd := exec.Command("sh", "-c", "exit 42")
	err := cmd.Run()
	if err == nil {
		t.Fatalf("%s did not return error", cmd.Args)
	}
	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("%s returned error of type %T, but expected *exec.ExitError", cmd.Args, err)
	}
	if status := ExitStatus(exitErr); status != 42 {
		t.Fatalf("expected exit status 42, but got %d", status)
	}
}
