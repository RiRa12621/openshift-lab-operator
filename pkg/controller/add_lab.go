package controller

import (
	"github.com/RiRa12621/openshift-lab-operator/pkg/controller/lab"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, lab.Add)
}
