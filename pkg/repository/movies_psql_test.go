package repository

import (
	"database/sql"
	"movieapi"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMovieListPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewMoviePostgres(db)

	type args struct {
		movie    movieapi.MovieList
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
				mock.ExpectQuery("INSERT INTO movielist").
					WithArgs("title", 10, "1999-9-9", "description", `{"a","b"}`).WillReturnRows(rows)

				mock.ExpectCommit()
			},
			input: args{
				movie: movieapi.MovieList{
					Title:       "title",
					Description: "description",
					Rating:      10,
					Date:        "1999-9-9",
					ActorName:   []string{"a", "b"},
				},
			},
			want: 1,
		},
		{
			name: "Empty input fields",
			mock: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO movielist").
					WithArgs("", 10, "1999-9-9", "description", `{"a","b"}`).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			input: args{
				userRole: "1",
				movie: movieapi.MovieList{
					Title:       "",
					Description: "description",
					Rating:      10,
					Date:        "1999-9-9",
					ActorName:   []string{"a", "b"},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.Create("1", tt.input.movie)
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

func TestMovieListPostgres_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewMoviePostgres(db)

	type args struct {
		order string
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    []movieapi.MovieList
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "rating", "date", "title", "description", "actorname"}).
					AddRow(1, "title1", 1, "1999-9-1", "description1", `{"a","b"}`).
					AddRow(2, "title2", 2, "1999-9-2", "description2", `{"a","b"}`).
					AddRow(3, "title3", 3, "1999-9-3", "description3", `{"a","b"}`)

				mock.ExpectQuery("SELECT (.+) FROM movielist ORDER BY (.+) DESC").WillReturnRows(rows)
			},
			input: args{
				order: "rating",
			},
			want: []movieapi.MovieList{
				{Id: 1, Title: "title1", Rating: 1, Date: "1999-9-1", Description: "description1", ActorName: []string{"a", "b"}},
				{Id: 2, Title: "title2", Rating: 2, Date: "1999-9-2", Description: "description2", ActorName: []string{"a", "b"}},
				{Id: 3, Title: "title3", Rating: 3, Date: "1999-9-3", Description: "description3", ActorName: []string{"a", "b"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.ListMovies(tt.input.order)
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

func TestMovieListPostgres_GetByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewMoviePostgres(db)

	type args struct {
		movieName string
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    []movieapi.MovieList
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "rating", "date", "description", "actorname"}).
					AddRow(1, "title1", 1, "1999-9-1", "description1", `{"a","b"}`)
				mock.ExpectQuery("SELECT (.+) FROM movielist WHERE title LIKE  (.+) ").
					WithArgs("title1").WillReturnRows(rows)
			},
			input: args{
				movieName: "title1",
			},
			want: []movieapi.MovieList{
				{Id: 1, Title: "title1", Rating: 1, Date: "1999-9-1", Description: "description1", ActorName: []string{"a", "b"}},
			},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "rating", "description", "actorname"})

				mock.ExpectQuery("SELECT (.+) FROM movielist WHERE title LIKE (.+) ").
					WithArgs("film").WillReturnRows(rows)
			},
			input: args{
				movieName: "film",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetByName(tt.input.movieName)
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

func TestMovieListPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewMoviePostgres(db)

	type args struct {
		userRole string
		movId    int
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
				mock.ExpectExec("DELETE FROM movielist WHERE (.+)").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				movId: 1,
			},
		},
		{
			name: "Not Found",
			mock: func() {
				mock.ExpectExec("DELETE FROM movielist WHERE (.+)").
					WithArgs(100).WillReturnError(sql.ErrNoRows)
			},
			input: args{
				movId: 100,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.Delete("1", tt.input.movId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestMovieListPostgres_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewMoviePostgres(db)

	type args struct {
		userRole string
		input    movieapi.UpdateMovieListInput
		movId    int
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
				mock.ExpectExec("UPDATE movielist SET (.+) WHERE (.+)").
					WithArgs("new title", 10, "new date", "new description", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				movId: 1,
				input: movieapi.UpdateMovieListInput{
					Title:       stringPointer("new title"),
					Rating:      intPointer(10),
					Date:        stringPointer("new date"),
					Description: stringPointer("new description"),
				},
			},
		},
		{
			name: "OK_OnlyTitle",
			mock: func() {
				mock.ExpectExec("UPDATE movielist SET (.+) WHERE (.+)").
					WithArgs("new title", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				movId: 1,
				input: movieapi.UpdateMovieListInput{
					Title: stringPointer("new title"),
				},
			},
		},
		{
			name: "OK_OnlyDate",
			mock: func() {
				mock.ExpectExec("UPDATE movielist SET (.+) WHERE (.+)").
					WithArgs("new date", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				movId: 1,
				input: movieapi.UpdateMovieListInput{
					Date: stringPointer("new date"),
				},
			},
		},
		{
			name: "OK_NoInputFields",
			mock: func() {
				mock.ExpectExec("UPDATE movielist SET  WHERE (.+)").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				movId: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.Update("1", tt.input.movId, tt.input.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func stringPointer(s string) *string {
	return &s
}

func intPointer(s int) *int {
	return &s
}
