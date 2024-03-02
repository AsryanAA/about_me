package models

type WorkPlace struct {
	Id        int64  `json:"id"`
	WorkPlace string `json:"work_place"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
	WhatDoing string `json:"what_doing"`
}
