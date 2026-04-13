package models

import (
	"time"
)

type DailyCashDrawer struct {
	ID            int       
	OpenedAt      time.Time 
	ClosedAt      *time.Time
	InitialValue  int       
	TotalRecorded int       
	FinalCounted  int       
	Status        string    
}