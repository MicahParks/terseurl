package storage

import (
	"context"
	"encoding/json"

	"go.etcd.io/bbolt"

	"github.com/MicahParks/terse-URL/models"
)

// BboltVisits if a VisitsStore implementation that relies on a bbolt file for the backend storage.
type BboltVisits struct {
	db           *bbolt.DB
	visitsBucket []byte
}

// NewBboltVisits creates a new NewBboltVisits given the required assets.
func NewBboltVisits(db *bbolt.DB, visitsBucket []byte) (visitsStore VisitsStore) {
	return &BboltVisits{
		db:           db,
		visitsBucket: visitsBucket,
	}
}

// AddVisit inserts the visit into the VisitsStore. This implementation has no network activity and ignores the given
// context.
func (b *BboltVisits) AddVisit(_ context.Context, shortened string, visit *models.Visit) (err error) {

	// Get the existing visits.
	var visits []*models.Visit
	if visits, err = b.readVisits(shortened); err != nil {
		return err
	}

	// Add the visits to the existing visits.
	visits = append(visits, visit)

	// Turn the visits into JSON data.
	var data []byte
	if data, err = json.Marshal(visits); err != nil {
		return err
	}

	// Open the bbolt database for writing, batch if possible.
	if err = b.db.Batch(func(tx *bbolt.Tx) error {

		// Put the updated JSON data into the bucket.
		if err = tx.Bucket(b.visitsBucket).Put([]byte(shortened), data); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// Close closes the bbolt database file. This implementation has no network activity and ignores the given context.
func (b *BboltVisits) Close(_ context.Context) (err error) {
	return b.db.Close()
}

// DeleteVisits deletes all visits associated with the given shortened URL. This implementation has no network activity
// and ignores the given context.
func (b *BboltVisits) DeleteVisits(_ context.Context, shortened string) (err error) {

	// Open the bbolt database for writing, batch if possible.
	if err = b.db.Batch(func(tx *bbolt.Tx) error {

		// Delete all of this shortened URL's visits from the bucket.
		if err = tx.Bucket(b.visitsBucket).Delete([]byte(shortened)); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// ReadVisits gets all visits for the given shortened URL. This implementation has no network activity and ignores the
// given context.
func (b *BboltVisits) ReadVisits(_ context.Context, shortened string) (visits []*models.Visit, err error) {
	return b.readVisits(shortened)
}

// readVisits gets all visits for the given shortened URL.
func (b *BboltVisits) readVisits(shortened string) (visits []*models.Visit, err error) {

	// Open the bbolt database for reading.
	var data []byte
	if err = b.db.View(func(tx *bbolt.Tx) error {

		// Get the Visits from the bucket.
		data = tx.Bucket(b.visitsBucket).Get([]byte(shortened))

		return nil
	}); err != nil {
		return nil, err
	}

	// Only unmarshal data if there was any.
	if data != nil {

		// Turn the JSON data into the Go structure.
		if err = json.Unmarshal(data, &visits); err != nil {
			return nil, err
		}
	}

	return visits, nil
}
