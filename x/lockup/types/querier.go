package types

// query endpoints supported by the lockup Querier
const (
	QueryModuleBalance                        = "module_balance"
	QueryModuleLockedAmount                   = "module_locked_amount"
	QueryAccountUnlockableCoins               = "account_unlockable_coins"
	QueryAccountLockedCoins                   = "account_locked_coins"
	QueryAccountLockedPastTime                = "account_locked_pasttime"
	QueryAccountUnlockedBeforeTime            = "account_unlocked_beforetime"
	QueryAccountLockedPastTimeDenom           = "account_locked_denom_pasttime"
	QueryLockedByID                           = "locked_by_id"
	QueryAccountLockedLongerThanDuration      = "account_locked_longer_than_duration"
	QueryAccountLockedLongerThanDurationDenom = "account_locked_longer_than_duration_denom"
)
