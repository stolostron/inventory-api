package eventing

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/project-kessel/inventory-api/internal/eventing/api"
	"github.com/project-kessel/inventory-api/internal/eventing/kafka"
	"github.com/project-kessel/inventory-api/internal/eventing/stdout"
)

func New(c CompletedConfig, logger log.Logger) (api.Manager, error) {

	switch c.Eventer {
	case "stdout":
		return stdout.New()
	case "kafka":
		return kafka.New(c.Kafka, logger)
	}
	return nil, fmt.Errorf("unrecognized eventer type: %s", c.Eventer)

}