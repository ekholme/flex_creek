package frstr

import flexcreek "github.com/ekholme/flex_creek"

//TODO
//implement WodService methods

type frstrDB struct{}

// hmm -- i might want frstrDB to implement WodService and UserService
// so i might want to wrap those in another interface?
func NewFrstrDB() flexcreek.WodService {
	return &frstrDB{}
}
