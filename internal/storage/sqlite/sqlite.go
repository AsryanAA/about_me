package sqlite

import (
	"about_me/internal/models"
	"about_me/internal/storage"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS work_places (
		    id INTEGER PRIMARY KEY,
		    work_place TEXT NOT NULL,
		    begin_date TEXT NOT NULL,
		    end_date TEXT,
		    what_doing TEXT);
		CREATE INDEX IF NOT EXISTS idx_alias ON work_places(work_place);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db}, nil
}

func (s *Storage) AddWorkPlace(workPlace, beginDate, endDate, whatDoing string) (int64, error) {
	const op = "storage.sqlite.AddWorkPlace"

	stmt, err := s.db.Prepare(`INSERT INTO work_places(work_place, begin_date, end_date, what_doing) 
									 VALUES (?, ?, ?, ?)`)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(workPlace, beginDate, endDate, whatDoing)
	if err != nil {
		// TODO: refactoring this
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", op, err)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetWorkPlaces() ([]models.WorkPlace, error) {
	const op = "storage.sqlite.GetWorkPlaces"

	stmt, err := s.db.Prepare(`SELECT * FROM work_places`)
	if err != nil {
		return nil, fmt.Errorf("%s: prepare statement %w", op, err)
	}

	var workPlaces []models.WorkPlace
	var workPlace models.WorkPlace
	rows, err := stmt.Query()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrWorkPlaceNotFound
		}

		return nil, fmt.Errorf("%s: execute statement %w", op, err)
	}

	for rows.Next() {
		err = rows.Scan(&workPlace.Id, &workPlace.WorkPlace, &workPlace.BeginDate, &workPlace.EndDate, &workPlace.WhatDoing)
		if err != nil {
			fmt.Printf("%s: error receive record %w", op, err)
		}
		workPlaces = append(workPlaces, workPlace)
	}

	return workPlaces, nil
}

func (s *Storage) UpdateWorkPlace(id int64, workPlace, beginDate, endDate, whatDoing string) error {
	const op = "storage.sqlite.UpdateWorkPlace"

	stmt, err := s.db.Prepare(`UPDATE work_places SET work_place = ?, begin_date = ?, end_date = ?, what_doing = ?
									  WHERE id = ?`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = stmt.Exec(workPlace, beginDate, endDate, whatDoing, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
