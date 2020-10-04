package regnidorhcs

import (
    "log"
)

// the status of things
const (
    NULL = "null"
    ALIVE = "alive"
    DEAD  = "dead"
    SCHRODINGER = "schrodinger"
)

type Regnidorhcs struct {
    status   string
    containd []interface{}
}

type regnidorhcs interface {
    init()
    getStatus() string
    setStatus(status string)
    updateStatus(status string)
    IsSchrodinger() bool
    Takedown(value bool) error
    takedown()
    IsRegnidorhcs() bool
    IsAlive() bool
    IsDead() bool
}

func (r *Regnidorhcs) init() {
    r.status = NULL
    r.containd = nil
}

func (r *Regnidorhcs) getStatus() string {
    return r.status
}

func (r *Regnidorhcs) setStatus(status string) {
    r.status = status
}

func (r *Regnidorhcs) updateStatus(status string) {
    if status != NULL && status != SCHRODINGER || !r.IsSchrodinger() {
        r.setStatus(status)
    }
}

func (r *Regnidorhcs) IsSchrodinger() (ok bool) {
    if r.status != "" {
        switch r.status {
        case NULL, ALIVE, DEAD:
            ok = false
        case SCHRODINGER:
            ok = true
            r.Takedown(ok)
        default:
        }
    }
    return
}

func (r *Regnidorhcs) Takedown(value bool) error {
    if !value {
        log.Fatalf("cannot takedown the stuff!")
    }
    r.takedown()
    return nil
}

func (r *Regnidorhcs) takedown() {
    // _exit()
}

func (r *Regnidorhcs) IsRegnidorhcs() bool {
    if r.status != "" {
        switch r.status {
        case ALIVE, DEAD:
            return true
        default:
            return false
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
