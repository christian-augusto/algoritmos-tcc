package constants

import "algoritmos-tcc-golang/utils"

const DEBUG = true
const LOGS = true
const TOTAL_LOGS = DEBUG

var ARRAY_SIZES = utils.IfThenElse(
	DEBUG, []int{8, 16}, []int{8, 16, 100, 1_000, 10_000, 25_000, 50_000, 75_000, 100_000, 1_000_000},
).([]int)

const THREADS_NUMBER = 4

const RANDOM_INT_MIN = 0

var RANDOM_INT_MAX = utils.IfThenElse(
	DEBUG, 100, 20_000_000_00,
).(int)
