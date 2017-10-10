package rules

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/uswitch/klint/engine"
	extv1b1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	prefix = "com.uswitch.heimdall"
)

var IngressNeedsAnnotation = engine.NewRule(
	func(old runtime.Object, new runtime.Object, ctx *engine.RuleHandlerContext) {
		ingress := new.(*extv1b1.Ingress)
		logger := log.WithFields(log.Fields{"name": ingress.Name, "namespace": ingress.Namespace, "rule": "IngressNeedsAnnotation"})
		annotation := ingress.GetAnnotations()
		hasAnnotation := false
		for key := range annotation {
			if strings.HasPrefix(key, prefix) {
				logger.Debugf("Checking annotation %s", key)
				hasAnnotation = true
				break
			}
		}
		if !hasAnnotation {
			ctx.Alertf(new, "You don't have any alerts set up for your ingress: %s.%s. You may want to check https://github.com/uswitch/heimdall for more info.", ingress.Namespace, ingress.Name)
		}
	}, engine.WantIngress)
