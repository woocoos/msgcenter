package nflogpb

// IsFiringSubset returns whether the given subset is a subset of the alerts
// that were firing at the time of the last notification.
func (m *Entry) IsFiringSubset(subset map[uint64]struct{}) bool {
	set := map[uint64]struct{}{}
	for i := range m.FiringAlerts {
		set[m.FiringAlerts[i]] = struct{}{}
	}

	return isSubset(set, subset)
}

// IsResolvedSubset returns whether the given subset is a subset of the alerts
// that were resolved at the time of the last notification.
func (m *Entry) IsResolvedSubset(subset map[uint64]struct{}) bool {
	set := map[uint64]struct{}{}
	for i := range m.ResolvedAlerts {
		set[m.ResolvedAlerts[i]] = struct{}{}
	}

	return isSubset(set, subset)
}

func isSubset(set, subset map[uint64]struct{}) bool {
	for k := range subset {
		_, exists := set[k]
		if !exists {
			return false
		}
	}

	return true
}
