package db

func initialTable() []string {
	queries := []string{
		exchangeTable,
		exchangeRateTable,
	}
	return queries
}

const (
	exchangeTable     string = "CREATE TABLE IF NOT EXISTS exchange (id INTEGER PRIMARY KEY AUTO_INCREMENT,`from` VARCHAR(255), `to` VARCHAR(255))"
	exchangeRateTable string = `CREATE TABLE IF NOT EXISTS exchange_rate (id INTEGER PRIMARY KEY AUTO_INCREMENT, exchange_id INTEGER, date DATE, rate FLOAT)`
)
