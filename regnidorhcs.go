package regnidorhcs

import (
	log "github.com/sirupsen/logrus"
)

// the status of things
const (
	NULL        = ""
	ALIVE       = "alive"
	DEAD        = "dead"
	SCHRODINGER = "schrodinger"
)

type Regnidorhcs struct {
	status  string
    // unknown indicate the status of the program in Schrodinger
	unknown interface{}
}

type regnidorhcs interface {
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

func (r *Regnidorhcs) Init() {
	r.status = NULL
	r.unknown = nil
}

func (r *Regnidorhcs) GetStatus() string {
	return r.status
}

func (r *Regnidorhcs) SetStatus(status string) {
	r.status = status
}

func (r *Regnidorhcs) UpdateStatus(status string) {
	if status != NULL && status != SCHRODINGER || !r.IsSchrodinger() {
		r.SetStatus(status)
	}
}

func (r *Regnidorhcs) IsSchrodinger() (ok bool) {
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

func (r *Regnidorhcs) wantDead(unknown ...interface{}) bool {
    if len(unknown) > 0 {
        r.unknown = unknown[0]
        if r.unknown == nil {
            return true
        }
        return false
    }
    return true
}

func (r *Regnidorhcs) Takedown(value bool, status string) bool {
	if !value {
		log.Fatalf("cannot takedown the program!")
	}
	r.takedown(status)
    return r.IsDead()
}

func (r *Regnidorhcs) takedown(status string) {
    r.SetStatus(status)
}

func (r *Regnidorhcs) Turnup(value bool, status string) bool {
    if value {
        log.Debugf("preparing to turn program up...")
    }
    r.turnup(status)
    return r.IsAlive()
}

func (r *Regnidorhcs) turnup(status string) {
    r.SetStatus(status)
}

func (r *Regnidorhcs) IsRegnidorhcs() bool {
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

func (r *Regnidorhcs) IsAlive() bool {
	if r.status == ALIVE {
		return true
	}
	return false
}

func (r *Regnidorhcs) IsDead() bool {
	if r.status == DEAD {
		return true
	}
	return false
}
