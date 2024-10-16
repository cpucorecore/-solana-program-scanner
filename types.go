package main

import "time"

/*
tx_hash        varchar(64)                 not null,
event          smallint default 0          not null,
token0_amount  varchar(70)                 not null,
token1_amount  varchar(70)                 not null,
maker          varchar(64)                 not null,
token0_address varchar(64)                 not null,
token1_address varchar(64)                 not null,
amount_usd     numeric(70, 18)             not null,
price_usd      numeric(70, 18)             not null,
block          bigint                      not null,
block_at       timestamp(6) with time zone not null,
created_at     timestamp(6) with time zone not null,
index          integer                     not null
*/
type Tx struct {
	TxHash        string
	Event         int8
	Token0Amount  string
	Token1Amount  string
	Maker         string
	Token0Address string
	Token1Address string
	AmountUsd     float64
	PriceUsd      float64
	Block         int64
	BlockAt       time.Time
	CreatedAt     time.Time `xorm:"created`
	Index         int
}
