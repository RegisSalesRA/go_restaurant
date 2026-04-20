package models

import (
	"time"
)

type Users struct {
	ID        int       
	FirstName string    
	LastName  string    
	Email     string    
	Password  string    
	CreatedAt time.Time 
}