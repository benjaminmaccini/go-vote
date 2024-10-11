package protocol

import (
	"context"
	"database/sql"

	"git.sr.ht/~bmaccini/go-vote/pkg/db"
)

func ValidateVoterExists(q *db.Queries, voter Voter) (bool, error) {
	_, err := q.UpsertVoter(context.Background(), db.UpsertVoterParams{
		ID:  voter.ID,
		Zip: sql.NullString{String: voter.Zip, Valid: true},
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
