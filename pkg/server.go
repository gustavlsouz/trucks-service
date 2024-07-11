package pkg

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/common/persistence"
	"github.com/gustavlsouz/trucks-service/internal/driver/controllers"
	"github.com/gustavlsouz/trucks-service/internal/driver/operations"
	truckControllers "github.com/gustavlsouz/trucks-service/internal/truck/controllers"
	truckOperations "github.com/gustavlsouz/trucks-service/internal/truck/operations"
	"github.com/joho/godotenv"
)

func Start(startedSuccessfully chan<- bool, envPath string, migrationPath string) {
	log.Println("Iniciando serviço")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Panicf("Erro ao carregar configuração de variáveis de ambiente: %v\n", err)
	}

	err = persistence.GetPersistenceInstance().Start(migrationPath)

	if err != nil {
		log.Panic("Não foi possível iniciar conexão com o banco de dados: ", err)
	}

	truckController := truckControllers.NewTruckController(
		truckOperations.NewTruckReaderCreator(),
		truckOperations.NewTruckInserterCreator(),
		truckOperations.NewTruckRemoverCreator(),
		truckOperations.NewTruckUpdaterCreator(),
	)

	truckHandler := common.NewHttpHandlerBuilder().
		Post(truckController.Create).
		Get(truckController.Read).
		Put(truckController.Update).
		Delete(truckController.Delete).
		Build()

	http.HandleFunc("/api/truck", truckHandler)

	driverController := controllers.NewDriverController(
		operations.NewDriverReaderCreator(),
		operations.NewDriverInserterCreator(),
		operations.NewDriverRemoverCreator(),
		operations.NewDriverUpdaterCreator(),
	)

	driverHandler := common.NewHttpHandlerBuilder().
		Post(driverController.Create).
		Get(driverController.Read).
		Put(driverController.Update).
		Delete(driverController.Delete).
		Build()

	http.HandleFunc("/api/driver", driverHandler)

	truckDriverController := truckControllers.NewTruckDriverController(
		truckOperations.NewTruckDriverReaderCreator(),
		truckOperations.NewTruckDriverInserterCreator(),
		truckOperations.NewTruckDriverRemoverCreator(),
		truckOperations.NewTruckDriverUpdaterCreator(),
	)

	truckDriverHandler := common.NewHttpHandlerBuilder().
		Post(truckDriverController.Create).
		Get(truckDriverController.Read).
		Put(truckDriverController.Update).
		Delete(truckDriverController.Delete).
		Build()

	http.HandleFunc("/api/truck/relation", truckDriverHandler)

	go func() {
		if err != nil {
			startedSuccessfully <- false
			return
		}
		startedSuccessfully <- true
	}()

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
