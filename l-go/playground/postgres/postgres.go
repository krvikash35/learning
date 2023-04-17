package postgres

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgres struct {
	db *sqlx.DB
}

type Tile struct {
	name   string
	active bool
}

type TileLogEntity struct {
	TileID    string  `db:"tile_id" json:"tile_id"`
	EventName string  `db:"event_name"`
	Changes   Changes `db:"changes"`
}

type Change struct {
	Field    string      `json:"field"`
	OldValue interface{} `json:"old_value"`
	NewValue interface{} `json:"new_value"`
}

type Changes []Change

func (c *Changes) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &c)
}

func NewClient(uname, pwd, host, port, dbName string) (*postgres, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&timezone=UTC", uname,
		pwd, host, port, dbName)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &postgres{
		db: db,
	}, nil
}

func (p *postgres) GetSingleTile() (Tile, error) {
	var name string
	var active bool
	err := p.db.QueryRow("select name, active from tiles limit 1").Scan(&name, &active)
	if err != nil {
		return Tile{}, err
	}

	return Tile{
		name,
		active,
	}, nil
}

func (p *postgres) GetTileLogs() {
	// tileLogsQuery := `select tile_id, event_name, changes from tile_logs`
	tileLogsQuery := `select tile_id, event_name, changes from tile_logs where tile_id ='80331a50-0859-400c-aec1-7431abb18c19'`
	var tileLogs []TileLogEntity

	err := p.db.Select(&tileLogs, tileLogsQuery)
	if err != nil {
		panic(fmt.Errorf("GetTileLogs error %w", err))
	}

	log.Printf("tileLogs %+v", tileLogs)

	t, err := json.Marshal(tileLogs)
	if err != nil {
		panic(fmt.Errorf("err1 %w", err))
	}
	log.Printf("json %+v", string(t))
}

func (p *postgres) InsertTileLog() {
	query := "insert into tile_logs(tile_id, event_name, changed_by, changes, payload) values($1, $2, $3, $4, $5)"

	var diffs Changes

	diffs = append(diffs, Change{Field: "name", OldValue: "vikash", NewValue: "vikash1"})
	diffs = append(diffs, Change{Field: "age", OldValue: 12, NewValue: 13})

	d, err := json.Marshal(diffs)
	if err != nil {
		panic("err2")
	}

	_, err = p.db.Exec(query, "80331a50-0859-400c-aec1-7431abb18c19", "update", "vikash", d, []byte("{}"))
	if err != nil {
		return
	}
}

func (p *postgres) RunQuery() {
	p.db.Exec("")
}

// insert into tile_logs(tile_id, event_name, changed_by, changes, payload) values('80331a50-0859-400c-aec1-7431abb18c19', 'update', 'vikash.kumar', '[{"field": "description.en", "old_value": "en", "new_value": "en1"}, {"field": "active", "old_value": false, "new_value": true}]', '{}');
