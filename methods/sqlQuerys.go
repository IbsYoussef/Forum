package methods

var (
	INSERT string = "INSERT INTO USERS (username, password, email) VALUES (?,?,?)"
	FIND   string = "SELECT COUNT(*) FROM users WHERE email = ?"
	SELECT string = "SELECT USERS WHERE email, password = (?,?);"
)
