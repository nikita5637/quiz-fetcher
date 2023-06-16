package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
)

// Place represents a row from 'place'.
type Place struct {
	ID         int    `json:"id"`          // id
	ExternalID int    `json:"external_id"` // external_id
	Name       string `json:"name"`        // name
	Address    string `json:"address"`     // address
}

// PlaceStorage is Place service implementation
type PlaceStorage struct {
	db *tx.Manager
}

// NewPlaceStorage creates new instance of PlaceStorage
func NewPlaceStorage(txManager *tx.Manager) *PlaceStorage {
	return &PlaceStorage{
		db: txManager,
	}
}

// GetAll returns all records
func (s *PlaceStorage) GetAll(ctx context.Context) ([]Place, error) {
	return s.Find(ctx, nil, "")
}

// Find perform find request by params
func (s *PlaceStorage) Find(ctx context.Context, q builder.Cond, sort string) ([]Place, error) {
	query := `SELECT id, external_id, name, address FROM place`

	var args []interface{}

	if q != nil {
		var where string
		var err error
		where, args, err = builder.ToSQL(q)
		if err != nil {
			return nil, err
		}
		query += ` WHERE ` + where
	}

	if sort != "" {
		query += ` ` + getOrderStmt(sort)
	}

	rows, err := s.db.Sync(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Place

	for rows.Next() {
		var item Place
		if err := rows.Scan(
			&item.ID,
			&item.ExternalID,
			&item.Name,
			&item.Address,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// FindWithLimit perform find request by params, offset and limit
func (s *PlaceStorage) FindWithLimit(ctx context.Context, q builder.Cond, sort string, offset, limit uint64) ([]Place, error) {
	query := `SELECT id, external_id, name, address FROM place`

	var args []interface{}

	if q != nil {
		var where string
		var err error
		where, args, err = builder.ToSQL(q)
		if err != nil {
			return nil, err
		}
		query += ` WHERE ` + where
	}

	if sort != "" {
		query += ` ` + getOrderStmt(sort)
	}

	if limit != 0 {
		query += ` OFFSET ? LIMIT ?`
		args = append(args, offset)
		args = append(args, limit)
	}

	rows, err := s.db.Sync(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Place

	for rows.Next() {
		var item Place
		if err := rows.Scan(
			&item.ID,
			&item.ExternalID,
			&item.Name,
			&item.Address,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// Total return count(*) by params
func (s *PlaceStorage) Total(ctx context.Context, q builder.Cond) (uint64, error) {
	query := `SELECT count(*) FROM place`

	var args []interface{}

	if q != nil {
		var where string
		var err error
		where, args, err = builder.ToSQL(q)
		if err != nil {
			return 0, err
		}
		query += ` WHERE ` + where
	}

	rows, err := s.db.Sync(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count uint64

	for rows.Next() {
		if err := rows.Scan(
			&count,
		); err != nil {
			return 0, err
		}
	}

	return count, nil
}

// Insert inserts the Place to the database.
func (s *PlaceStorage) Insert(ctx context.Context, item Place) (int, error) {
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO place (` +
		`external_id, name, address` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`
	// run
	logger.Debugf(ctx, sqlstr, item.ExternalID, item.Name, item.Address)

	res, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.ExternalID, item.Name, item.Address)
	if err != nil {
		return 0, err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Update updates a Place in the database.
func (s *PlaceStorage) Update(ctx context.Context, item Place) error {
	// update with primary key
	const sqlstr = `UPDATE place SET ` +
		`external_id = ?, name = ?, address = ? ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, item.ExternalID, item.Name, item.Address, item.ID)
	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.ExternalID, item.Name, item.Address, item.ID); err != nil {
		return err
	}

	return nil
}

// Upsert performs an upsert for Place.
func (s *PlaceStorage) Upsert(ctx context.Context, item Place) error {
	// upsert
	const sqlstr = `INSERT INTO place (` +
		`id, external_id, name, address` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`external_id = VALUES(external_id), name = VALUES(name), address = VALUES(address)`
	// run
	logger.Debugf(ctx, sqlstr, item.ID, item.ExternalID, item.Name, item.Address)
	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.ID, item.ExternalID, item.Name, item.Address); err != nil {
		return err
	}

	return nil
}

// Delete deletes the Place from the database.
func (s *PlaceStorage) Delete(ctx context.Context, id int) error {
	// delete with single primary key
	const sqlstr = `DELETE FROM place ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)

	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, id); err != nil {
		return err
	}

	return nil
}

// GetPlaceByID retrieves a row from 'place' as a Place.
//
// Generated from index 'place_id_pkey'.
func (s *PlaceStorage) GetPlaceByID(ctx context.Context, id int) (*Place, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, external_id, name, address ` +
		`FROM place ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)
	p := Place{}
	if err := s.db.Sync(ctx).QueryRowContext(ctx, sqlstr, id).Scan(&p.ID, &p.ExternalID, &p.Name, &p.Address); err != nil {
		return nil, err
	}
	return &p, nil
}
