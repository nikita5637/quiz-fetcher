package model

import (
	"time"
)

// SyncStatus ...
type SyncStatus int

const (
	// SyncStatusNotSynced ...
	SyncStatusNotSynced SyncStatus = iota
	// SyncStatusInProgress ...
	SyncStatusInProgress
	// SyncStatusOK ...
	SyncStatusOK
	// SyncStatusFailed ...
	SyncStatusFailed
)

func (s SyncStatus) String() string {
	switch s {
	case SyncStatusNotSynced:
		return "SYNC_STATUS_NOT_SYNCED"
	case SyncStatusInProgress:
		return "SYNC_STATUS_IN_PROGRESS"
	case SyncStatusOK:
		return "SYNC_STATUS_OK"
	case SyncStatusFailed:
		return "SYNC_STATUS_FAILED"
	}

	return "SYNC_STATUS_UNKNOWN"
}

// SyncLog ...
type SyncLog struct {
	ID         int
	Name       string
	LastSyncAt time.Time
	Status     SyncStatus
}
