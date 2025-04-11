package piscine

func FindNextPrime(nb int) int {
	if nb < 1 {
		return 2
	}
	is_prime := false
	for !is_prime {
		is_prime = IsPrime(nb)
		if is_prime {
			return nb
		}
		nb += 1
	}
	return 1
}
