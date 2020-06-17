package dmodels

import (
	"time"
)

const (
	ValidatorsTable    = "validators_list_view"
	ValidatorStatsView = "validator_day_stats_view"
)

type Validator struct {
	EntityID          string    `db:"reg_entity_id"`
	ConsensusAddress  string    `db:"reg_consensus_address"`
	ValidateSince     time.Time `db:"created_time"`
	StartBlockLevel   uint64    `db:"start_blk_lvl"`
	LastBlockTime     time.Time `db:"last_block_time"`
	BlocksCount       uint64    `db:"blocks"`
	LastSignatureTime uint64    `db:"last_signature_time"`
	SignaturesCount   uint64    `db:"signatures"`
	EscrowBalance     uint64    `db:"acb_escrow_balance_active"`
	DepositorsNum     uint64    `db:"depositors_num"`
	IsActive          bool      `db:"is_active"`
	ValidatorName     string    `db:"pvl_name"`
	ValidatorFee      uint64    `db:"pvl_fee"`
	WebAddress        string    `db:"pvl_address"`
	AvailabilityScore uint64    `db:"-"`
	Status            string    `db:"-"`
}

type ValidatorStats struct {
	BeginOfPeriod     time.Time
	LastBlock         uint64
	AvailabilityScore uint64
	BlocksCount       uint64
	SignaturesCount   uint64
}
