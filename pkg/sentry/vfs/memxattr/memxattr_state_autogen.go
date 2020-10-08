// automatically generated by stateify.

package memxattr

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (s *SimpleExtendedAttributes) StateTypeName() string {
	return "pkg/sentry/vfs/memxattr.SimpleExtendedAttributes"
}

func (s *SimpleExtendedAttributes) StateFields() []string {
	return []string{
		"xattrs",
	}
}

func (s *SimpleExtendedAttributes) beforeSave() {}

func (s *SimpleExtendedAttributes) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.xattrs)
}

func (s *SimpleExtendedAttributes) afterLoad() {}

func (s *SimpleExtendedAttributes) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.xattrs)
}

func init() {
	state.Register((*SimpleExtendedAttributes)(nil))
}
