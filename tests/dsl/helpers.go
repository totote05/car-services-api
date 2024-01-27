package dsl

import (
	"reflect"

	"github.com/stretchr/testify/mock"
)

func GetTypeAsString(i any) string {
	typeOf := reflect.TypeOf(i)
	return typeOf.String()
}

func AnythingOfType(i any) mock.AnythingOfTypeArgument {
	return mock.AnythingOfType(
		GetTypeAsString(i),
	)
}
