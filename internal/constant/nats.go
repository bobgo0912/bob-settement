package constant

const (
	SettleStream   = "Settle"
	SettleSubject  = "Settle.>"
	SettleConsumer = "Settle_"
	SettleEvent    = "Settle.event.*"

	SettleBeginEvent  = "Settle.event.begin"
	SettleGoingEvent  = "Settle.event.going"
	SettleFinishEvent = "Settle.event.finish"
	SettleCancelEvent = "Settle.event.cancel"
)
