package span

type Span struct {
	From, To int
}

func (s1 Span) Contains(s2 Span) bool {
	if s1.From <= s2.From && s1.To >= s2.To {
		return true
	}

	if s2.From <= s1.From && s2.To >= s1.To {
		return true
	}
	return false
}

func (s1 Span) Overlaps(s2 Span) bool {
	if s1.Contains(s2) {
		return true
	}

	/*
	 * xxx
	 *  x
	 */
	if s1.From <= s2.From && s1.To >= s2.To {
		return true
	}

	/*
	 *  x
	 * xxx
	 */
	if s1.From <= s2.From && s1.To >= s2.From {
		return true
	}

	/*
	 *  xx
	 * xx
	 */
	if s2.From <= s1.From && s2.To >= s1.From {
		return true
	}

	/*
	 * xx
	 *  xx
	 */
	if s2.From <= s1.From && s2.To >= s1.To {
		return true
	}

	/*
	 * xx
	 *  xx
	 */
	if s1.To == s2.From || s2.To == s1.From {
		return true
	}

	return false
}
