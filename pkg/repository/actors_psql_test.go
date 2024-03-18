package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestActorListPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewActorPostgres(db)

	type args struct {
		userRole string
		actorId  int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectExec("DELETE FROM actorlist WHERE (.+)").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				actorId: 1,
			},
		},
		{
			name: "Not Found",
			mock: func() {
				mock.ExpectExec("DELETE FROM actorlist WHERE (.+)").
					WithArgs(100).WillReturnError(sql.ErrNoRows)
			},
			input: args{
				actorId: 100,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.Delete(tt.input.userRole, tt.input.actorId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
