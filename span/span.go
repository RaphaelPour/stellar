package span

type Span struct {
	from, to int
}

func (s1 Span) Contains(s2 Span) bool {
	if s1.from <= s2.from && s1.to >= s2.to {
		return true
	}

	if s2.from <= s1.from && s2.to >= s1.to {
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
	if s1.from <= s2.from && s1.to >= s2.to {
		return true
	}

	/*
	 *  x
	 * xxx
	 */
	if s1.from <= s2.from && s1.to >= s2.from {
		return true
	}

	/*
	 *  xx
	 * xx
	 */
	if s2.from <= s1.from && s2.to >= s1.from {
		return true
	}

	/*
	 * xx
	 *  xx
	 */
	if s2.from <= s1.from && s2.to >= s1.to {
		return true
	}

	/*
	 * xx
	 *  xx
	 */
	if s1.to == s2.from || s2.to == s1.from {
		return true
	}

	return false
}
