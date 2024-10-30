package repositories

import (
	"fmt"
	"time"

	"slurpy/interfaces"
	"slurpy/models"
)

type DeploymentRepository struct {
	Storage interfaces.Storage
}

func (d *DeploymentRepository) GetAll() ([]models.Deployment, error) {
	d.Storage.Open()
	defer d.Storage.Close()

	sql := `
		SELECT d.id, d.name, d.contract, d.created_at, d.group_name
		FROM deployments AS d
	`
	rows, err := d.Storage.Query(&sql, &[]interface{}{})
	if err != nil {
		fmt.Println("Failed to fetch deployments")
		return nil, err
	}
	defer rows.Close()

	var data []models.Deployment

	for rows.Next() {
		var id int
		var name, contract, createdAt, group string
		err := rows.Scan(&id, &name, &contract, &createdAt, &group)
		if err != nil {
			return nil, err
		}
		date, _ := time.Parse("2006-01-02 15:04:05", createdAt)

		deployment := &models.Deployment{
			Id:       id,
			Name:     name,
			Contract: contract,
			Date:     date,
			Group:    group,
			Options:  []string{},
		}

		sqlParameters := `
			SELECT parameter FROM deployment_parameters
			WHERE deploymentId = $1
		`

		rows, err := d.Storage.Query(&sqlParameters, &[]interface{}{
			&id,
		})

		if err != nil {
			continue
		}
		var parameters []string
		for rows.Next() {
			var param string
			err := rows.Scan(&param)

			if err != nil {
				fmt.Println("Can't retrive parameter for deloyment with contract ", contract)
			}

			parameters = append(parameters, param)
		}
		deployment.Options = parameters
		data = append(data, *deployment)
	}

	return data, nil
}

func (d *DeploymentRepository) GetDeploymentByKey(key string) ([]models.Deployment, error) {
	d.Storage.Open()
	defer d.Storage.Close()

	sql := `
		SELECT d.id, d.name, d.contract, d.created_at, d.group_name
		FROM deployments AS d
		WHERE d.group_name = $1
	`
	rows, err := d.Storage.Query(&sql, &[]interface{}{
		&key,
	})

	if err != nil {
		fmt.Println("Failed to fetch deployments")
		return nil, err
	}
	defer rows.Close()

	var data []models.Deployment
	for rows.Next() {
		var id int
		var name, contract, createdAt, group string
		err := rows.Scan(&id, &name, &contract, &createdAt, &group)
		if err != nil {
			return nil, err
		}

		date, _ := time.Parse("2006-01-02 15:04:05", createdAt)
		deployment := &models.Deployment{
			Id:       id,
			Name:     name,
			Contract: contract,
			Date:     date,
			Group:    group,
			Options:  []string{},
		}

		sqlParameters := `
			SELECT parameter FROM deployment_parameters
			WHERE deploymentId = $1
		`

		rows, err := d.Storage.Query(&sqlParameters, &[]interface{}{
			&id,
		})

		if err != nil {
			continue
		}
		var parameters []string
		for rows.Next() {
			var param string
			err := rows.Scan(&param)

			if err != nil {
				fmt.Println("Can't retrive parameter for deloyment with contract ", contract)
			}

			parameters = append(parameters, param)
		}
		deployment.Options = parameters
		data = append(data, *deployment)
	}

	return data, nil
}

func (d *DeploymentRepository) SaveDeployment(address *string, name *string, key *string) (int, error) {
	d.Storage.Open()
	defer d.Storage.Close()

	insertDeploymentSQL := `
		INSERT INTO deployments (contract, name, created_at, group_name)
		VALUES ($1, $2, datetime('now', 'localtime'), $3)
		RETURNING id
	`
	row := d.Storage.QuerySingle(&insertDeploymentSQL, &[]interface{}{
		&address,
		&name,
		&key,
	})

	var id int
	err := row.Scan(&id)
	return id, err
}

func (d *DeploymentRepository) SaveParameters(params *[]interface{}, id int) error {
	d.Storage.Open()
	defer d.Storage.Close()

	insertParamsSQL := `
		INSERT INTO deployment_parameters (parameter, deploymentId)
		VALUES ($1, $2)
	`

	for _, param := range *params {
		fmt.Println("saving ", param)
		str, ok := param.(string)
		if !ok {
			fmt.Print("Failed to convert parameter")
			continue
		}

		err := d.Storage.Exec(&insertParamsSQL, &[]interface{}{
			str,
			&id,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
