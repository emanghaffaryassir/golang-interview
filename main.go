package main

import (
	"fmt"
	"time"
)

type Subscription struct {
	UserId           string
	MonthlyFee       float64
	SubscriptionDate time.Time
}
type Bill struct {
	UserId  string
	Status  int64 // 1 pending 2 paid
	Amount  float64
	DueDate time.Time
}

func main() {

	Subscriptions := []Subscription{
		{
			UserId:           "userA",
			MonthlyFee:       50,
			SubscriptionDate: time.Date(2021, 7, 25, 0, 0, 0, 0, time.Local),
		},
		{
			UserId:           "userB",
			MonthlyFee:       50,
			SubscriptionDate: time.Date(2022, 6, 2, 0, 0, 0, 0, time.Local),
		},
		{
			UserId:           "userC",
			MonthlyFee:       50,
			SubscriptionDate: time.Date(2023, 8, 1, 0, 0, 0, 0, time.Local),
		},
	}

	Bills := []Bill{} // Only current month bills

	fmt.Println(Subscriptions, Bills)
}
