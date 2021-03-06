// Copyright © 2018 The ELPS authors

package lisp_test

import (
	"testing"

	"github.com/luthersystems/elps/elpstest"
)

func TestPackages(t *testing.T) {
	tests := elpstest.TestSuite{
		{"basic namespace evaluation", elpstest.TestSequence{
			// The lisp package contains all default builtins
			{"(lisp:+ 2 3)", "5", ""},
			{"(+ 2 3)", "5", ""},
			{"(lisp:list 1 2 3)", "'(1 2 3)", ""},
			{"(list 1 2 3)", "'(1 2 3)", ""},
		}},
		{"in-package", elpstest.TestSequence{
			// Switch into a new package and define a function
			{"(in-package 'new-package)", "()", ""},
			{"(defun fun (x) (+ x 1))", "()", ""},
			{"(in-package 'other-package)", "()", ""},
			{"(defun fun (x) (- x 1))", "()", ""},
			{"(in-package 'user)", "()", ""},
			{`(new-package:fun 2)`, `3`, ""},
			{`(other-package:fun 2)`, `1`, ""},
		}},
		{"use-package", elpstest.TestSequence{
			// Define fun in a new-package
			{"(in-package 'new-package)", "()", ""},
			{"(defun fun (x) (internal x))", "()", ""},
			{"(defun internal (x) (+ x 1))", "()", ""},
			{"(export 'fun)", "()", ""},
			// Also define fun in other-package
			{"(in-package 'other-package)", "()", ""},
			{"(defun fun (x) (internal x))", "()", ""},
			{"(defun internal (x) (- x 1))", "()", ""},
			{"(export 'fun)", "()", ""},
			// Use new-package and check the result of an unqualified fun call.
			{"(in-package 'user)", "()", ""},
			{"(use-package 'new-package)", "()", ""},
			{`(fun 2)`, `3`, ""},
			{`(other-package:fun 2)`, `1`, ""},
		}},
	}
	elpstest.RunTestSuite(t, tests)
}
