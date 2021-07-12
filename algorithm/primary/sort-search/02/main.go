package _02

import "math"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
var bad int

func firstBadVersion(n int) int {
	if isBadVersion(n) {
		return findMid(1, n)
	}
	return findMid(n, math.MaxInt32)
}

func findMid(lft, rgt int) int {
	mid := (lft + rgt) / 2
	if mid == rgt || mid == lft {
		if isBadVersion(lft) && isBadVersion(rgt) {
			return lft
		}
		return rgt
	}
	if isBadVersion(mid) {
		return findMid(lft, mid)
	} else {
		return findMid(mid, rgt)
	}
}

func isBadVersion(version int) bool {
	return version >= bad
}
