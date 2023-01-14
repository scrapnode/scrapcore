package noop

import "github.com/scrapnode/scrapcore/xmonitor/attributes"

type Span struct{}

func (span *Span) SetAttributes(attributes attributes.Attributes) {}
func (span *Span) OK(desc string)                                 {}
func (span *Span) KO(desc string)                                 {}
func (span *Span) End()                                           {}
