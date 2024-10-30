package interfaces

import "slurpy/models"

type DeploymentRepository interface {
	GetAll() ([]models.Deployment, error)
	GetDeploymentByKey(key string) ([]models.Deployment, error)
	SaveDeployment(address *string, name *string, key *string) (int, error)
	SaveParameters(params *[]interface{}, id int) error
}
