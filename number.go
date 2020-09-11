package urandom

// Integer returns a random integer value between lb(lowerbound) and ub(upperbound)
func Integer(lb, ub int) int {
	return seed.Intn(ub-lb) + lb
}
