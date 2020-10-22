package regnidorhcs

import (
	log "github.com/sirupsen/logrus"
)

// the status of program
const (
	NULL        = ""
	ALIVE       = "alive"
	DEAD        = "dead"
	SCHRODINGER = "schrodinger"
)

type regnidorhcs struct {
	status string
	// unknown indicates the status of the program in Schrodinger
	unknown interface{}
}

// R defines the regnidorhcs interface
type R interface {
	Init()
	GetStatus() string
	SetStatus(status string)
	UpdateStatus(status string)
	IsSchrodinger() bool
	wantDead(unknown ...interface{}) bool
	Takedown(value bool, status string) bool
	takedown(status string)
	Turnup(value bool, status string) bool
	turnup(stauts string)
	IsRegnidorhcs() bool
	IsAlive() bool
	IsDead() bool
}

// Init initials regnidorhcs fileds
func (r *regnidorhcs) Init() {
	r.status = NULL
	r.unknown = nil
}

// GetStatus gets the status of program
func (r *regnidorhcs) GetStatus() string {
	return r.status
}

// SetStatus sets the status of program
func (r *regnidorhcs) SetStatus(status string) {
	r.status = status
}

// UpdateStatus updates the status of program
func (r *regnidorhcs) UpdateStatus(status string) {
	if status != NULL && status != SCHRODINGER || !r.IsSchrodinger() {
		r.SetStatus(status)
	}
}

// IsSchrodinger checks if program's status under the schrodinger
func (r *regnidorhcs) IsSchrodinger() (ok bool) {
	if r.status != NULL {
		switch r.status {
		case ALIVE, DEAD:
			ok = false
		case SCHRODINGER:
			ok = true
			if r.wantDead(r.unknown) {
				r.Takedown(ok, DEAD)
			}
			r.Turnup(ok, ALIVE)
		default:
		}
	}
	return
}

func (r *regnidorhcs) wantDead(unknown ...interface{}) bool {
	if len(unknown) > 0 {
		r.unknown = unknown[0]
		if r.unknown == nil {
			return true
		}
		return false
	}
	return true
}

// Takedown takes the program's status down
func (r *regnidorhcs) Takedown(value bool, status string) bool {
	if !value {
		log.Fatalf("cannot takedown the program!")
	}
	r.takedown(status)
	return r.IsDead()
}

func (r *regnidorhcs) takedown(status string) {
	r.SetStatus(status)
}

// Turnup turns the program's status up
func (r *regnidorhcs) Turnup(value bool, status string) bool {
	if value {
		log.Debugf("preparing to turn program up...")
	}
	r.turnup(status)
	return r.IsAlive()
}

func (r *regnidorhcs) turnup(status string) {
	r.SetStatus(status)
}

// IsRegnidorhcs checks if program's status under the regnidorhcs
func (r *regnidorhcs) IsRegnidorhcs() bool {
	if r.status != NULL {
		switch r.status {
		case ALIVE, DEAD:
			return true
		case SCHRODINGER:
			return false
		default:
		}
	}
	return false
}

// IsAlive checks if program's status is alive
func (r *regnidorhcs) IsAlive() bool {
	if r.status == ALIVE {
		return true
	}
	return false
}

// IsDead checks if program's status is dead
func (r *regnidorhcs) IsDead() bool {
	if r.status == DEAD {
		return true
	}
	return false
}
