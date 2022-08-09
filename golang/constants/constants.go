package constants

import "algoritmos-tcc-golang/utils"

const DEBUG = true
const LOGS = true
const TOTAL_LOGS = DEBUG

var ARRAY_SIZES = utils.IfThenElse(
	DEBUG, []int{8, 16}, []int{
		8, 16, 100, 1_000, 10_000, 25_000, 50_000, 75_000, 100_000, 250_000, 500_000, 750_000, 1_000_000, 5_000_000,
		10_000_000, 25_000_000, 50_000_000, 100_000_000, 250_000_000, 500_000_000,
	},
).([]int)

const THREADS_NUMBER = 4

const RANDOM_INT_MIN = 0

var RANDOM_INT_MAX = utils.IfThenElse(
	DEBUG, 100, 2_000_000_000,
).(int)
