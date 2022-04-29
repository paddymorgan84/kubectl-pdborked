package ui

import v1 "k8s.io/api/policy/v1"

// Renderer is an interface I use to abstract away any future render implementations I may want, such as JSON
type Renderer interface {
	Render(pdbs []v1.PodDisruptionBudget)
}
