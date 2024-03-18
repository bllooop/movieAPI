package repository

import (
	"database/sql"
	"movieapi"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestActorListPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewActorPostgres(db)

	type args struct {
		actor    movieapi.ActorList
		userRole string
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO actorlist").
					WithArgs("title", 10, "1999-9-9", "description", `{"a", "b"}`).WillReturnRows(rows)

				mock.ExpectCommit()
			},
			input: args{
				actor: movieapi.ActorList{
					Name:      "name",
					Gender:    "gender",
					Birthdate: "1999",
				},
			},
			want: 1,
		},
		{
			name: "Empty input fields",
			mock: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO actorlist").
					WithArgs("", 10, "1999-9-9", "description", `{"a", "b"}`).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			input: args{
				userRole: "1",
				actor: movieapi.ActorList{
					Name:      "",
					Gender:    "gender",
					Birthdate: "1999",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.CreateActor(tt.input.userRole, tt.input.actor)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestActorListPostgres_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewActorPostgres(db)

	tests := []struct {
		name    string
		mock    func()
		want    []movieapi.ActorList
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "rating", "date", "title", "description", "actorname"}).
					AddRow(1, "name1", 1, "1999-9-1", "description1", `{"a", "b"}`).
					AddRow(2, "name2", 2, "1999-9-2", "description2", `{"a", "b"}`).
					AddRow(3, "name3", 3, "1999-9-3", "description3", `{"a", "b"}`)

				mock.ExpectQuery("SELECT at.id,at.name,at.gender,at.date,array_agg(mt.title) FROM actorlist at LEFT JOIN movielist mt ON at.name = ANY(mt.actorname) GROUP BY (at.id, at.name, at.gender,at.date)").WillReturnRows(rows)
			},
			want: []movieapi.ActorList{
				{Id: 1, Name: "name1", Gender: "male1", Birthdate: "1999-9-1"},
				{Id: 2, Name: "name2", Gender: "male2", Birthdate: "1999-9-2"},
				{Id: 3, Name: "name3", Gender: "male3", Birthdate: "1999-9-3"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.ListActors()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

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

func TestActorListPostgres_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewActorPostgres(db)

	type args struct {
		userRole string
		input    movieapi.UpdateActorListInput
		actorId  int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "OK_AllFields",
			mock: func() {
				mock.ExpectExec("UPDATE actorlist SET (.+) FROM actorlist WHERE (.+)").
					WithArgs("new title", "new description", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				actorId: 1,
				input: movieapi.UpdateActorListInput{
					Name:      stringPointer("new name"),
					Gender:    stringPointer("new gender"),
					Birthdate: stringPointer("1900"),
					//ActorName:   slicePointer([]string{"a,b"}),
				},
			},
		},
		{
			name: "OK_WithoutDescription",
			mock: func() {
				mock.ExpectExec("UPDATE actorlist SET (.+) FROM actorlist WHERE (.+)").
					WithArgs("new title", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				actorId: 1,
				input: movieapi.UpdateActorListInput{
					Name: stringPointer("new name"),
				},
			},
		},
		{
			name: "OK_WithoutTitle",
			mock: func() {
				mock.ExpectExec("UPDATE actorlist SET (.+) FROM actorlist WHERE (.+)").
					WithArgs("new description", 1, 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				actorId: 1,
				input: movieapi.UpdateActorListInput{
					Gender: stringPointer("new gender"),
				},
			},
		},
		{
			name: "OK_NoInputFields",
			mock: func() {
				mock.ExpectExec("UPDATE actorlist SET (.+) FROM actorlist WHERE (.+)").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				actorId: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.Update(tt.input.userRole, tt.input.actorId, tt.input.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
