package utils

import (
	"database/sql"
	"reflect"
)

func ScanStructFromRows(rows *sql.Rows, ptr interface{}) error {
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	columnPointers := make([]interface{}, len(columns))
	for i := range columns {
		columnPointers[i] = new(interface{})
	}

	err = rows.Scan(columnPointers...)
	if err != nil {
		return err
	}

	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < len(columns); i++ {
		columnValue := reflect.ValueOf(columnPointers[i].(*interface{})).Elem().Interface()
		field := v.FieldByName(columns[i])
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(columnValue))
		}
	}

	return nil
}
