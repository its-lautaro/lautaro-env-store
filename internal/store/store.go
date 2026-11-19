// Package store es el archivo ~/.lautaro-env: una línea `export NAME='valor'  # saved FECHA` por variable,
// con permisos 600, pensado para que el shell lo cargue con source.
package store

type Store struct {
	Path string
}
