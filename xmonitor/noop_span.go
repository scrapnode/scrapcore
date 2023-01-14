package xmonitor

import "github.com/scrapnode/scrapcore/xmonitor/attributes"

type NoopSpan struct{}

func (span *NoopSpan) SetAttributes(attributes attributes.Attributes) {}
func (span *NoopSpan) OK(desc string)                                 {}
func (span *NoopSpan) KO(desc string)                                 {}
func (span *NoopSpan) End()                                           {}
