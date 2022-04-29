package ui

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	v1 "k8s.io/api/policy/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// TableRenderer is a concrete implementation of the Renderer interface
type TableRenderer struct {
}

// Render the outpit from the available pod disruption budgets
func (t TableRenderer) Render(pdbs []v1.PodDisruptionBudget) {
	tab := table.NewWriter()

	tab.SetOutputMirror(os.Stdout)

	tab.AppendHeader(table.Row{
		text.FgCyan.Sprint("PDB"),
		text.FgCyan.Sprint("Namespace"),
		text.FgCyan.Sprint("Minimum Available"),
		text.FgCyan.Sprint("Maximum Unavailable"),
		text.FgCyan.Sprint("Current Healthy"),
		text.FgCyan.Sprint("Desired Healthy"),
		text.FgCyan.Sprint("Expected Pods"),
	})

	for _, pdb := range pdbs {
		if pdb.Status.DisruptionsAllowed == 0 {
			if pdb.Spec.MaxUnavailable == nil && pdb.Spec.MinAvailable == nil {
				tab.Style().Color.Row = text.Colors{text.Reset, text.FgYellow}
			}

			tab.AppendRow([]interface{}{
				pdb.Name,
				pdb.Namespace,
				checkNil(pdb.Spec.MinAvailable),
				checkNil(pdb.Spec.MinAvailable),
				pdb.Status.CurrentHealthy,
				pdb.Status.DesiredHealthy,
				pdb.Status.ExpectedPods,
			})
			tab.Style().Color.Row = text.Colors{text.Reset, text.Reset}
		}
	}

	tab.Render()
}

func checkNil(is *intstr.IntOrString) string {
	if is == nil {
		return "-"
	}

	return is.String()
}
