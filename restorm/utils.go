package restorm

import (
	"bufio"
	"bytes"
	"database/sql"
	"github.com/pubgo/g/xerror"
	"io"
	"strings"
)

//Import SQL DDL from sql file
func (t *RestOrm) Import(name string, f io.Reader) (res []sql.Result, err error) {
	defer xerror.RespErr(&err)

	scanner := bufio.NewScanner(f)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.IndexByte(data, ';'); i >= 0 {
			return i + 1, data[0:i], nil
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), data, nil
		}

		// Request more data.
		return 0, nil, nil
	})

	var results []sql.Result
	for scanner.Scan() {
		query := strings.Trim(scanner.Text(), " \t\n\r")
		if len(query) > 0 {
			result, err := t.cfg[name].db.Exec(query)
			xerror.PanicM(err, "Import Exec error, db(%s)", name)
			results = append(results, result)
		}
	}

	res = results
	return
}
