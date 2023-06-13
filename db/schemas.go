package db

type DiscountCode struct {
	ID    uint64 `gorm:"primaryKey"`
	Code  uint64
	Valid bool
}
