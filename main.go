package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nbedregal/gambituser/awsgo"
	"github.com/nbedregal/gambituser/bd"
	"github.com/nbedregal/gambituser/models"
)

func main() {
	lambda.Start(EjecutoLambda)
}

/**
* Función principal de ejecución lambda, que recive el contexto
* y el evento de confirmación luego de un signup en la UI.
 */
func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parámetros, debe enviar 'SecretManager'")
		err := errors.New("Error en los parámetros, debe enviar SecretName")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el secret " + err.Error())
		return event, err
	}

	err = bd.SignUp(datos)
	return event, err

}

/**
* Función que valida la variable de sistema definido en
* la función lambda exista.
 */
func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
