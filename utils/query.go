package utils

import (
	"context"
	"database/sql"
	"reflect"
	"task_mission/queries"
)

const (
	selectQuery = `SELECT`
	insertQuery = `INSERT`
	updateQuery = `UPDATE`
	deleteQuery = `DELETE`
)

func Save[T any](ctx context.Context, db *sql.DB, query queries.GoQuery, requests ...any) (*uint64, error) {
	err := QueryValidation(query, insertQuery)
	if err != nil {
		return nil, err
	}
	var id = new(uint64)
	err = db.QueryRowContext(ctx, query.ToString(), requests...).Scan(&id)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func FindAll[T any](ctx context.Context, db *sql.DB, query queries.GoQuery, requests *T) (result []*T, err error) {
	err = QueryValidation(query, selectQuery)
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, query.ToString())
	if err != nil {
		return nil, err
	}

	typ := reflect.TypeOf(requests).Elem()
	var ptrs = make([]interface{}, typ.NumField())
	for i := range ptrs {
		ptrs[i] = reflect.New(typ.Field(i).Type).Interface()
	}

	for rows.Next() {
		request := reflect.New(typ).Elem()
		err := rows.Scan(ptrs...)
		if err != nil {
			return nil, err
		}
		for i, ptr := range ptrs {
			field := request.Field(i)
			field.Set(reflect.ValueOf(ptr).Elem())
		}
		result = append(result, request.Addr().Interface().(*T))
	}
	return result, nil
}

func Delete[T any](ctx context.Context, db *sql.DB, query queries.GoQuery, id uint64) (err error) {
	err = QueryValidation(query, deleteQuery)
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, query.ToString(), id)
	if err != nil {
		return err
	}
	return nil
}

func FindByID[T any](ctx context.Context, db *sql.DB, query queries.GoQuery, id uint64, result *T) error {
	err := QueryValidation(query, selectQuery)
	if err != nil {
		return err
	}
	stmt, err := db.PrepareContext(ctx, query.ToString())
	if err != nil {
		return err
	}
	defer stmt.Close()

	var ptrs []interface{}

	typ := reflect.TypeOf(result).Elem()

	for i := 0; i < typ.NumField(); i++ {
		field := reflect.ValueOf(result).Elem().Field(i)
		ptrs = append(ptrs, field.Addr().Interface())
	}

	err = stmt.QueryRowContext(ctx, id).Scan(ptrs...)
	if err != nil {
		return err
	}
	return nil
}

func Update[T any](ctx context.Context, db *sql.DB, query queries.GoQuery, toupdate *T) (result *T, err error) {
	err = QueryValidation(query, updateQuery)
	if err != nil {
		return nil, err
	}
	return
}
