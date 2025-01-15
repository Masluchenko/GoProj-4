package files

import (
	"fmt"
	"os"

	"example.com/m/v2/output"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}

}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintErorr(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		output.PrintErorr("Неверный формат URL или Логин")
		return
	}
	fmt.Println("Запись успешна")
}
