package bucket

import (
	"github.com/belak/seabird/bot"
	"github.com/jmoiron/sqlx"
)

func init() {
	bot.RegisterPlugin("bucket", NewBucketPlugin)
}

type BucketPlugin struct {
	db *sqlx.DB
}

func NewBucketPlugin(b *bot.Bot) (bot.Plugin, error) {
	b.LoadPlugin("db")
	p := &BucketPlugin{b.Plugins["db"].(*sqlx.DB)}

	return p, nil
}
