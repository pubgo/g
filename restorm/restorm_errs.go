package restorm

var Errs = struct {
	DbCreateError    string
	DbDeleteError    string
	DbUpdateError    string
	DbCountError     string
	DbGetError       string
	DbDeleteOneError string
	DbUpdateOneError string
	DbGetOneError    string
}{
	DbCreateError:    "db_create_error",
	DbDeleteError:    "db_delete_error",
	DbUpdateError:    "db_update_error",
	DbCountError:     "db_count_error",
	DbGetError:       "db_get_error",
	DbDeleteOneError: "db_delete_one_error",
	DbUpdateOneError: "db_update_one_error",
	DbGetOneError:    "db_get_one_error",
}
