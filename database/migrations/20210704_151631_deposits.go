package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Deposits_20210704_151631 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Deposits_20210704_151631{}
	m.Created = "20210704_151631"

	migration.Register("Deposits_20210704_151631", m)
}

// Run the migrations
func (m *Deposits_20210704_151631) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE IF NOT EXISTS "deposit" (
        "id" serial NOT NULL PRIMARY KEY,
        "email" text NOT NULL DEFAULT '' ,
        "txid" text NOT NULL DEFAULT '' ,
        "currency" text NOT NULL DEFAULT '' ,
        "amount" numeric(12, 2) NOT NULL DEFAULT 0 ,
        "status" text NOT NULL DEFAULT '' ,
        "created_at" timestamp with time zone NOT NULL,
        "updated_at" timestamp with time zone NOT NULL
    );
`)
}

// Reverse the migrations
func (m *Deposits_20210704_151631) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE deposit")
}
