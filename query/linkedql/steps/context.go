package steps

import (
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/query"
	"github.com/cayleygraph/cayley/query/linkedql"
	"github.com/cayleygraph/cayley/query/path"
	"github.com/cayleygraph/quad/voc"
)

func init() {
	linkedql.Register(&Context{})
}

var _ linkedql.IteratorStep = (*Context)(nil)
var _ linkedql.PathStep = (*Context)(nil)

// Context corresponds to context().
type Context struct {
	From  linkedql.PathStep `json:"from"`
	Rules map[string]string `json:"rules"`
}

// Description implements Step.
func (s *Context) Description() string {
	return "A a set of rules for interpreting identifiers used in the query"
}

// BuildPath implements linkedql.PathStep.
func (s *Context) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	for name, iri := range s.Rules {
		namespace := voc.Namespace{Prefix: name, Full: iri}
		ns.Register(namespace)
	}
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath, nil
}

// BuildIterator implements linkedql.IteratorStep.
func (s *Context) BuildIterator(qs graph.QuadStore, ns *voc.Namespaces) (query.Iterator, error) {
	return linkedql.NewValueIteratorFromPathStep(s, qs, ns)
}
