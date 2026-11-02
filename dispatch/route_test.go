// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2015 Prometheus Team.
// Licensed under the Apache License 2.0.

package dispatch

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
)

// testRouteYAML is the shared route configuration used by most route tests.
// It mirrors the upstream Alertmanager test route tree.
const testRouteYAML = `
route:
  receiver: 'notify-def'
  routes:
    - matchers: ['owner="team-A"']
      receiver: 'notify-A'
      routes:
        - matchers: ['env="testing"']
          receiver: 'notify-testing'
          groupBy: ['...']
        - matchers: ['env="production"']
          receiver: 'notify-productionA'
          groupWait: 1m
          continue: true
        - matchers: ['env=~"produ.*"', 'job=~".*"']
          receiver: 'notify-productionB'
          groupWait: 30s
          groupInterval: 5m
          repeatInterval: 1h
          groupBy: ['job']
    - matchers: ['owner=~"team-(B|C)"']
      groupBy: ['foo', 'bar']
      groupWait: 2m
      receiver: 'notify-BC'
    - matchers: ['group_by="role"']
      groupBy: ['role']
      receiver: 'notify-def'
      routes:
        - matchers: ['env="testing"']
          receiver: 'notify-testing'
          routes:
            - matchers: ['wait="long"']
              groupWait: 2m
              receiver: 'notify-testing'
receivers:
  - name: 'notify-def'
  - name: 'notify-A'
  - name: 'notify-testing'
  - name: 'notify-productionA'
  - name: 'notify-productionB'
  - name: 'notify-BC'
`

func loadTestRoute(t *testing.T) *Route {
	t.Helper()
	cfg, err := profile.Load([]byte(testRouteYAML))
	require.NoError(t, err)
	return NewRoute(cfg.Route, nil)
}

func TestRouteMatch(t *testing.T) {
	tree := loadTestRoute(t)

	def := DefaultRouteOpts
	lset := func(labels ...string) map[label.LabelName]struct{} {
		s := map[label.LabelName]struct{}{}
		for _, ls := range labels {
			s[label.LabelName(ls)] = struct{}{}
		}
		return s
	}

	tests := []struct {
		input  label.LabelSet
		result []*RouteOpts
		keys   []string
	}{
		{
			input: label.LabelSet{"owner": "team-A"},
			result: []*RouteOpts{
				{Receiver: "notify-A", GroupBy: def.GroupBy, GroupWait: def.GroupWait, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{owner=\"team-A\"}"},
		},
		{
			input: label.LabelSet{"owner": "team-A", "env": "unset"},
			result: []*RouteOpts{
				{Receiver: "notify-A", GroupBy: def.GroupBy, GroupWait: def.GroupWait, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{owner=\"team-A\"}"},
		},
		{
			input: label.LabelSet{"owner": "team-C"},
			result: []*RouteOpts{
				{Receiver: "notify-BC", GroupBy: lset("foo", "bar"), GroupWait: 2 * time.Minute, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{owner=~\"team-(B|C)\"}"},
		},
		{
			input: label.LabelSet{"owner": "team-A", "env": "testing"},
			result: []*RouteOpts{
				{Receiver: "notify-testing", GroupBy: lset(), GroupByAll: true, GroupWait: def.GroupWait, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{owner=\"team-A\"}/{env=\"testing\"}"},
		},
		{
			input: label.LabelSet{"owner": "team-A", "env": "production"},
			result: []*RouteOpts{
				{Receiver: "notify-productionA", GroupBy: def.GroupBy, GroupWait: time.Minute, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
				{Receiver: "notify-productionB", GroupBy: lset("job"), GroupWait: 30 * time.Second, GroupInterval: 5 * time.Minute, RepeatInterval: time.Hour},
			},
			keys: []string{
				"{}/{owner=\"team-A\"}/{env=\"production\"}",
				"{}/{owner=\"team-A\"}/{env=~\"produ.*\",job=~\".*\"}",
			},
		},
		{
			input: label.LabelSet{"group_by": "role"},
			result: []*RouteOpts{
				{Receiver: "notify-def", GroupBy: lset("role"), GroupWait: def.GroupWait, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{group_by=\"role\"}"},
		},
		{
			input: label.LabelSet{"env": "testing", "group_by": "role"},
			result: []*RouteOpts{
				{Receiver: "notify-testing", GroupBy: lset("role"), GroupWait: def.GroupWait, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{group_by=\"role\"}/{env=\"testing\"}"},
		},
		{
			input: label.LabelSet{"env": "testing", "group_by": "role", "wait": "long"},
			result: []*RouteOpts{
				{Receiver: "notify-testing", GroupBy: lset("role"), GroupWait: 2 * time.Minute, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{group_by=\"role\"}/{env=\"testing\"}/{wait=\"long\"}"},
		},
	}

	for _, test := range tests {
		var matches []*RouteOpts
		var keys []string

		for _, r := range tree.Match(test.input) {
			matches = append(matches, &r.RouteOpts)
			keys = append(keys, r.Key())
		}

		if !reflect.DeepEqual(matches, test.result) {
			t.Errorf("input %v\nexpected:\n%v\ngot:\n%v", test.input, test.result, matches)
		}
		if !reflect.DeepEqual(keys, test.keys) {
			t.Errorf("input %v\nexpected keys:\n%v\ngot:\n%v", test.input, test.keys, keys)
		}
	}
}

func TestRouteWalk(t *testing.T) {
	tree := loadTestRoute(t)

	expected := []string{
		"notify-def",
		"notify-A",
		"notify-testing",
		"notify-productionA",
		"notify-productionB",
		"notify-BC",
		"notify-def",
		"notify-testing",
		"notify-testing",
	}

	var got []string
	tree.Walk(func(r *Route) {
		got = append(got, r.RouteOpts.Receiver)
	})

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\nexpected:\n%v\ngot:\n%v", expected, got)
	}
}

func TestInheritParentGroupByAll(t *testing.T) {
	in := `
route:
  receiver: 'default'
  routes:
    - matchers: ['env="parent"']
      groupBy: ['...']
      routes:
        - matchers: ['env="child1"']
        - matchers: ['env="child2"']
          groupBy: ['foo']
receivers:
  - name: 'default'
`
	cfg, err := profile.Load([]byte(in))
	require.NoError(t, err)
	tree := NewRoute(cfg.Route, nil)

	parent := tree.Routes[0]
	child1 := parent.Routes[0]
	child2 := parent.Routes[1]
	require.True(t, parent.RouteOpts.GroupByAll)
	require.True(t, child1.RouteOpts.GroupByAll)
	require.False(t, child2.RouteOpts.GroupByAll)
}

func TestRouteMatchers(t *testing.T) {
	in := `
route:
  receiver: 'notify-def'
  routes:
    - matchers: ['owner="team-A"', 'level!="critical"']
      receiver: 'notify-A'
      routes:
        - matchers: ['env="testing"', 'baz!~".*quux"']
          receiver: 'notify-testing'
          groupBy: ['...']
        - matchers: ['env="production"']
          receiver: 'notify-productionA'
          groupWait: 1m
          continue: true
        - matchers: ['env=~"produ.*"', 'job=~".*"']
          receiver: 'notify-productionB'
          groupWait: 30s
          groupInterval: 5m
          repeatInterval: 1h
          groupBy: ['job']
    - matchers: ['owner=~"team-(B|C)"']
      groupBy: ['foo', 'bar']
      groupWait: 2m
      receiver: 'notify-BC'
    - matchers: ['group_by="role"']
      groupBy: ['role']
      receiver: 'notify-def'
      routes:
        - matchers: ['env="testing"']
          receiver: 'notify-testing'
          routes:
            - matchers: ['wait="long"']
              groupWait: 2m
              receiver: 'notify-testing'
receivers:
  - name: 'notify-def'
  - name: 'notify-A'
  - name: 'notify-testing'
  - name: 'notify-productionA'
  - name: 'notify-productionB'
  - name: 'notify-BC'
`
	cfg, err := profile.Load([]byte(in))
	require.NoError(t, err)
	tree := NewRoute(cfg.Route, nil)

	def := DefaultRouteOpts
	lset := func(labels ...string) map[label.LabelName]struct{} {
		s := map[label.LabelName]struct{}{}
		for _, ls := range labels {
			s[label.LabelName(ls)] = struct{}{}
		}
		return s
	}

	tests := []struct {
		input  label.LabelSet
		result []*RouteOpts
		keys   []string
	}{
		{
			input: label.LabelSet{"owner": "team-A"},
			result: []*RouteOpts{
				{Receiver: "notify-A", GroupBy: def.GroupBy, GroupWait: def.GroupWait, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{level!=\"critical\",owner=\"team-A\"}"},
		},
		{
			input: label.LabelSet{"owner": "team-A", "env": "unset"},
			result: []*RouteOpts{
				{Receiver: "notify-A", GroupBy: def.GroupBy, GroupWait: def.GroupWait, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{level!=\"critical\",owner=\"team-A\"}"},
		},
		{
			input: label.LabelSet{"owner": "team-C"},
			result: []*RouteOpts{
				{Receiver: "notify-BC", GroupBy: lset("foo", "bar"), GroupWait: 2 * time.Minute, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{owner=~\"team-(B|C)\"}"},
		},
		{
			input: label.LabelSet{"owner": "team-A", "env": "testing"},
			result: []*RouteOpts{
				{Receiver: "notify-testing", GroupBy: lset(), GroupByAll: true, GroupWait: def.GroupWait, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{level!=\"critical\",owner=\"team-A\"}/{baz!~\".*quux\",env=\"testing\"}"},
		},
		{
			input: label.LabelSet{"owner": "team-A", "env": "production"},
			result: []*RouteOpts{
				{Receiver: "notify-productionA", GroupBy: def.GroupBy, GroupWait: time.Minute, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
				{Receiver: "notify-productionB", GroupBy: lset("job"), GroupWait: 30 * time.Second, GroupInterval: 5 * time.Minute, RepeatInterval: time.Hour},
			},
			keys: []string{
				"{}/{level!=\"critical\",owner=\"team-A\"}/{env=\"production\"}",
				"{}/{level!=\"critical\",owner=\"team-A\"}/{env=~\"produ.*\",job=~\".*\"}",
			},
		},
		{
			input: label.LabelSet{"group_by": "role"},
			result: []*RouteOpts{
				{Receiver: "notify-def", GroupBy: lset("role"), GroupWait: def.GroupWait, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{group_by=\"role\"}"},
		},
		{
			input: label.LabelSet{"env": "testing", "group_by": "role"},
			result: []*RouteOpts{
				{Receiver: "notify-testing", GroupBy: lset("role"), GroupWait: def.GroupWait, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{group_by=\"role\"}/{env=\"testing\"}"},
		},
		{
			input: label.LabelSet{"env": "testing", "group_by": "role", "wait": "long"},
			result: []*RouteOpts{
				{Receiver: "notify-testing", GroupBy: lset("role"), GroupWait: 2 * time.Minute, GroupInterval: def.GroupInterval, RepeatInterval: def.RepeatInterval},
			},
			keys: []string{"{}/{group_by=\"role\"}/{env=\"testing\"}/{wait=\"long\"}"},
		},
	}

	for _, test := range tests {
		var matches []*RouteOpts
		var keys []string

		for _, r := range tree.Match(test.input) {
			matches = append(matches, &r.RouteOpts)
			keys = append(keys, r.Key())
		}

		if !reflect.DeepEqual(matches, test.result) {
			t.Errorf("input %v\nexpected:\n%v\ngot:\n%v", test.input, test.result, matches)
		}
		if !reflect.DeepEqual(keys, test.keys) {
			t.Errorf("input %v\nexpected keys:\n%v\ngot:\n%v", test.input, test.keys, keys)
		}
	}
}

func TestRouteID(t *testing.T) {
	in := `
route:
  receiver: 'default'
  routes:
    - continue: true
      matchers: ['foo="bar"']
      receiver: 'test1'
      routes:
        - matchers: ['bar="baz"']
    - continue: true
      matchers: ['foo="bar"']
      receiver: 'test1'
      routes:
        - matchers: ['bar="baz"']
    - continue: true
      matchers: ['foo="bar"']
      receiver: 'test2'
      routes:
        - matchers: ['bar="baz"']
    - continue: true
      matchers: ['bar="baz"']
      receiver: 'test3'
      routes:
        - matchers: ['baz="qux"']
        - matchers: ['qux="corge"']
    - continue: true
      matchers: ['qux=~"[a-zA-Z0-9]+"']
    - continue: true
      matchers: ['corge!~"[0-9]+"']
receivers:
  - name: 'default'
  - name: 'test1'
  - name: 'test2'
  - name: 'test3'
`
	cfg, err := profile.Load([]byte(in))
	require.NoError(t, err)
	tree := NewRoute(cfg.Route, nil)

	expected := []string{
		"{}",
		"{}/{foo=\"bar\"}/0",
		"{}/{foo=\"bar\"}/0/{bar=\"baz\"}/0",
		"{}/{foo=\"bar\"}/1",
		"{}/{foo=\"bar\"}/1/{bar=\"baz\"}/0",
		"{}/{foo=\"bar\"}/2",
		"{}/{foo=\"bar\"}/2/{bar=\"baz\"}/0",
		"{}/{bar=\"baz\"}/3",
		"{}/{bar=\"baz\"}/3/{baz=\"qux\"}/0",
		"{}/{bar=\"baz\"}/3/{qux=\"corge\"}/1",
		"{}/{qux=~\"[a-zA-Z0-9]+\"}/4",
		"{}/{corge!~\"[0-9]+\"}/5",
	}

	var actual []string
	tree.Walk(func(r *Route) {
		actual = append(actual, r.ID())
	})
	require.ElementsMatch(t, actual, expected)
}

func TestRouteIndices(t *testing.T) {
	tree := loadTestRoute(t)

	// Collect all indices.
	var indices []int
	var totalNodes int
	tree.Walk(func(r *Route) {
		indices = append(indices, r.Idx)
		totalNodes++
	})

	// All indices are unique.
	seenIndices := make(map[int]bool)
	for _, idx := range indices {
		require.False(t, seenIndices[idx], "Index %d appears more than once", idx)
		seenIndices[idx] = true
	}

	// Root index equals total nodes - 1.
	require.Equal(t, totalNodes-1, tree.Idx, "Root index should equal total nodes - 1")

	// All indices are in range [0, totalNodes).
	for _, idx := range indices {
		require.GreaterOrEqual(t, idx, 0, "Index should be >= 0")
		require.Less(t, idx, totalNodes, "Index should be < total nodes")
	}
}
