// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package render

import (
	"strconv"
	"strings"

	"github.com/derailed/tcell/v2"
)

// HorizontalPodAutoscaler renders a K8s HorizontalPodAutoscaler to screen.
type HorizontalPodAutoscaler struct {
	Generic
}

// ColorerFunc colors a resource row.
func (hpa HorizontalPodAutoscaler) ColorerFunc() ColorerFunc {
	return func(ns string, h Header, re RowEvent) tcell.Color {
		c := DefaultColorer(ns, h, re)

		maxPodsIndex := h.IndexOf("MAXPODS", true)
		currentReplicasIndex := h.IndexOf("REPLICAS", true)
		maxPodsS := strings.TrimSpace(re.Row.Fields[maxPodsIndex])
		currentReplicasS := strings.TrimSpace(re.Row.Fields[currentReplicasIndex])

		maxPods, err := strconv.Atoi(maxPodsS)
		if err != nil {
			return c
		}
		currentReplicas, err := strconv.Atoi(currentReplicasS)
		if err != nil {
			return c
		}
		if currentReplicas >= maxPods {
			c = ErrColor
			return c
		}
		return c
	}
}
