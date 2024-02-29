// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package builder

import (
	"github.com/cilium/cilium-cli/utils/features"

	"github.com/cilium/cilium/pkg/cilium-cli/connectivity/check"
	"github.com/cilium/cilium/pkg/cilium-cli/connectivity/tests"
)

type fromCidrHostNetns struct{}

func (t fromCidrHostNetns) build(ct *check.ConnectivityTest, templates map[string]string) {
	newTest("from-cidr-host-netns", ct).
		WithCondition(func() bool { return ct.Params().IncludeUnsafeTests }).
		WithFeatureRequirements(features.RequireEnabled(features.NodeWithoutCilium)).
		WithCiliumPolicy(templates["echoIngressFromCIDRYAML"]).
		WithIPRoutesFromOutsideToPodCIDRs().
		WithScenarios(tests.FromCIDRToPod()).
		WithExpectations(func(_ *check.Action) (egress, ingress check.Result) {
			return check.ResultOK, check.ResultNone
		})
}
