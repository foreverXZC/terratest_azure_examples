package test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformHttpExample(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../../clouddrive/database",

		Vars: map[string]interface{}{},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	server := terraform.Output(t, terraformOptions, "sql_server_fqdn")
	port := "1433"
	user := "azureuser"
	password := "P@ssw0rd12345!"
	database := terraform.Output(t, terraformOptions, "database_name")

	maxRetries := 15
	timeBetweenRetries := 5 * time.Second
	description := fmt.Sprintf("Executing commands on database %s", server)

	retry.DoWithRetry(t, description, maxRetries, timeBetweenRetries, func() (string, error) {
		config := fmt.Sprintf("server = %s; port = %s; user id = %s; password = %s; database = %s", server, port, user, password, database)
		db, err := sql.Open("mssql", config)
		if err != nil {
			return "", err
		}

		_, err0 := db.Exec("create table person (id integer, name varchar(30), primary key (id))")
		if err0 != nil {
			return "", err0
		}

		expectedID := 12345
		expectedName := "azure"
		insertion := fmt.Sprintf("insert into person values (%d, '%s')", expectedID, expectedName)
		_, err1 := db.Exec(insertion)
		if err1 != nil {
			return "", err1
		}

		rows, err2 := db.Query("select * from person")
		if err2 != nil {
			return "", err2
		}

		var id int
		var name string
		for rows.Next() {
			err3 := rows.Scan(&id, &name)
			if err3 != nil {
				return "", err3
			}
			fmt.Println(id, name)
			assert.Equal(t, expectedID, id)
			assert.Equal(t, expectedName, name)
		}

		_, err4 := db.Exec("drop table person")
		if err4 != nil {
			return "", err4
		}
		fmt.Println("Executed SQL commands correctly")

		defer rows.Close()
		defer db.Close()

		return "", nil
	})
}
