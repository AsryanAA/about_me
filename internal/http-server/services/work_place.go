package services

import (
	"about_me/internal/models"
	"about_me/internal/storage/sqlite"
)

func CreateWorkPlace(workPlace models.WorkPlace, storage *sqlite.Storage) error {
	_, err := storage.AddWorkPlace(workPlace.WorkPlace, workPlace.BeginDate, workPlace.EndDate, workPlace.WhatDoing)
	if err != nil {
		return err
	}
	return nil
}

func ReadWorkPlaces(storage *sqlite.Storage) ([]models.WorkPlace, error) {
	workPlace, err := storage.GetWorkPlaces()
	if err != nil {
		return workPlace, err
	}
	return workPlace, nil
}

func UpdateWorkPlace(workPlace models.WorkPlace, storage *sqlite.Storage) error {
	err := storage.UpdateWorkPlace(workPlace.Id, workPlace.WorkPlace, workPlace.BeginDate, workPlace.EndDate, workPlace.WhatDoing)
	if err != nil {
		return err
	}
	return nil
}
