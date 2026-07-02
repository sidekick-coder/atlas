package database 

func Create(filename string) (*Database, error) {
	connection, err := Connect(filename)

	if err != nil {
		panic(err)
	}

	return New(connection), err
}
