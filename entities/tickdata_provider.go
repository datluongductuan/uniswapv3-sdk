package entities

import "math/big"

type Tick struct {
	LiquidityNet *big.Int
}

// Provides information about ticks
type TickDataProvider interface {
	/**
	 * Return information corresponding to a specific tick
	 * @param tick the tick to load
	 */
	GetTick(tick int64) Tick

	/**
	 * Return the next tick that is initialized within a single word
	 * @param tick The current tick
	 * @param lte Whether the next tick should be lte the current tick
	 * @param tickSpacing The tick spacing of the pool
	 */
	NextInitializedTickWithinOneWord(tick int64, lte bool, tickSpacing int64) (int64, bool)
}