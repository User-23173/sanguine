// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

// AddressRanking gives the amount of transactions that occurred for a specific address across all chains.
type AddressRanking struct {
	Address *string `json:"address"`
	Count   *int    `json:"count"`
}

// BridgeTransaction represents an entire bridge transaction, including both
// to and from transactions. If a `from` transaction does not have a corresponding
// `to` transaction, `pending` will be true.
type BridgeTransaction struct {
	FromInfo    *PartialInfo `json:"fromInfo"`
	ToInfo      *PartialInfo `json:"toInfo"`
	Kappa       *string      `json:"kappa"`
	Pending     *bool        `json:"pending"`
	SwapSuccess *bool        `json:"swapSuccess"`
}

// DateResult is a given statistic for a given date.
type DateResult struct {
	Date  *string  `json:"date"`
	Total *float64 `json:"total"`
}

// HistoricalResult is a given statistic for dates.
type HistoricalResult struct {
	Total       *float64              `json:"total"`
	DateResults []*DateResult         `json:"dateResults"`
	Type        *HistoricalResultType `json:"type"`
}

// PartialInfo is a transaction that occurred on one chain.
type PartialInfo struct {
	ChainID        *int     `json:"chainId"`
	Address        *string  `json:"address"`
	TxnHash        *string  `json:"txnHash"`
	Value          *string  `json:"value"`
	FormattedValue *float64 `json:"formattedValue"`
	USDValue       *float64 `json:"USDValue"`
	TokenAddress   *string  `json:"tokenAddress"`
	TokenSymbol    *string  `json:"tokenSymbol"`
	BlockNumber    *int     `json:"blockNumber"`
	Time           *int     `json:"time"`
}

// TokenCountResult gives the amount of transactions that occurred for a specific token, separated by chain ID.
type TokenCountResult struct {
	ChainID      *int    `json:"chainId"`
	TokenAddress *string `json:"tokenAddress"`
	Count        *int    `json:"count"`
}

// TransactionCountResult gives the amount of transactions that occurred for a specific chain ID.
type TransactionCountResult struct {
	ChainID *int `json:"chainId"`
	Count   *int `json:"count"`
}

// ValueResult is a value result of either USD or numeric value.
type ValueResult struct {
	Value *string `json:"value"`
}

type Direction string

const (
	DirectionIn  Direction = "IN"
	DirectionOut Direction = "OUT"
)

var AllDirection = []Direction{
	DirectionIn,
	DirectionOut,
}

func (e Direction) IsValid() bool {
	switch e {
	case DirectionIn, DirectionOut:
		return true
	}
	return false
}

func (e Direction) String() string {
	return string(e)
}

func (e *Direction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Direction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Direction", str)
	}
	return nil
}

func (e Direction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Duration string

const (
	DurationPastDay   Duration = "PAST_DAY"
	DurationPastMonth Duration = "PAST_MONTH"
	DurationAllTime   Duration = "ALL_TIME"
)

var AllDuration = []Duration{
	DurationPastDay,
	DurationPastMonth,
	DurationAllTime,
}

func (e Duration) IsValid() bool {
	switch e {
	case DurationPastDay, DurationPastMonth, DurationAllTime:
		return true
	}
	return false
}

func (e Duration) String() string {
	return string(e)
}

func (e *Duration) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Duration(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Duration", str)
	}
	return nil
}

func (e Duration) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type HistoricalResultType string

const (
	HistoricalResultTypeBridgevolume HistoricalResultType = "BRIDGEVOLUME"
	HistoricalResultTypeTransactions HistoricalResultType = "TRANSACTIONS"
	HistoricalResultTypeAddresses    HistoricalResultType = "ADDRESSES"
)

var AllHistoricalResultType = []HistoricalResultType{
	HistoricalResultTypeBridgevolume,
	HistoricalResultTypeTransactions,
	HistoricalResultTypeAddresses,
}

func (e HistoricalResultType) IsValid() bool {
	switch e {
	case HistoricalResultTypeBridgevolume, HistoricalResultTypeTransactions, HistoricalResultTypeAddresses:
		return true
	}
	return false
}

func (e HistoricalResultType) String() string {
	return string(e)
}

func (e *HistoricalResultType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HistoricalResultType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HistoricalResultType", str)
	}
	return nil
}

func (e HistoricalResultType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StatisticType string

const (
	StatisticTypeMeanVolumeUsd     StatisticType = "MEAN_VOLUME_USD"
	StatisticTypeMedianVolumeUsd   StatisticType = "MEDIAN_VOLUME_USD"
	StatisticTypeTotalVolumeUsd    StatisticType = "TOTAL_VOLUME_USD"
	StatisticTypeCountTransactions StatisticType = "COUNT_TRANSACTIONS"
	StatisticTypeCountAddresses    StatisticType = "COUNT_ADDRESSES"
)

var AllStatisticType = []StatisticType{
	StatisticTypeMeanVolumeUsd,
	StatisticTypeMedianVolumeUsd,
	StatisticTypeTotalVolumeUsd,
	StatisticTypeCountTransactions,
	StatisticTypeCountAddresses,
}

func (e StatisticType) IsValid() bool {
	switch e {
	case StatisticTypeMeanVolumeUsd, StatisticTypeMedianVolumeUsd, StatisticTypeTotalVolumeUsd, StatisticTypeCountTransactions, StatisticTypeCountAddresses:
		return true
	}
	return false
}

func (e StatisticType) String() string {
	return string(e)
}

func (e *StatisticType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StatisticType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StatisticType", str)
	}
	return nil
}

func (e StatisticType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}