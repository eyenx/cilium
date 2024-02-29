// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package builder

import (
	"github.com/cilium/cilium-cli/utils/features"

	"github.com/cilium/cilium/pkg/cilium-cli/connectivity/check"
	"github.com/cilium/cilium/pkg/cilium-cli/connectivity/tests"
)

type clientEgressToCidrDenyDefault struct{}

func (t clientEgressToCidrDenyDefault) build(ct *check.ConnectivityTest, templates map[string]string) {
	// This test is same as the previous one, but there is no allowed policy.
	// The goal is to test default deny policy
	newTest("client-egress-to-cidr-deny-default", ct).
		WithCiliumPolicy(templates["clientEgressToCIDRExternalDenyPolicyYAML"]).
		WithScenarios(tests.PodToCIDR()). // Denies all traffic to ExternalOtherIP, but allow ExternalIP
		WithExpectations(func(a *check.Action) (egress, ingress check.Result) {
			if a.Destination().Address(features.GetIPFamily(ct.Params().ExternalOtherIP)) == ct.Params().ExternalOtherIP {
				return check.ResultPolicyDenyEgressDrop, check.ResultNone
			}
			if a.Destination().Address(features.GetIPFamily(ct.Params().ExternalIP)) == ct.Params().ExternalIP {
				return check.ResultDefaultDenyEgressDrop, check.ResultNone
			}
			return check.ResultDrop, check.ResultDrop
		})
}
