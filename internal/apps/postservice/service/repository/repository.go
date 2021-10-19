/*
Package repository package contains the repository layer of the post service.
*/
package repository

import "github.com/aidarkhanov/nanoid/v2"

func GenerateID() (string, error) {
	return nanoid.New()
}
