package withApis

import (
    dba "app/database"
)

type Foo struct {
    ID int `json:"id"`
    Name string `json:"name"`
}

func GetFoo() (Foo, error) {

    foo := Foo{}

    db, err := dba.DbOpen()
    if err != nil {
        return foo, err
    }
    defer db.Close()

    rows, err := db.Query("select 10, 'foo' from dual")
	if err != nil {
		return foo, err
	}
	defer rows.Close()

    for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)

        foo.ID = id
        foo.Name = name
	}

    return foo, nil
}