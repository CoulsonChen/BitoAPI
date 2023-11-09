package service

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/CoulsonChen/BitoAPI/internal/constant"
	"github.com/CoulsonChen/BitoAPI/internal/model"
	nutsutil "github.com/CoulsonChen/BitoAPI/third-party/nutsdb"
	"github.com/nutsdb/nutsdb"
)

var bucket = "Match"

type MatchService struct {
	nuts *nutsdb.DB
}

var instance *MatchService

func InitMatchService() {
	instance = &MatchService{
		nuts: nutsutil.GetInstance(),
	}
}

func GetMatchServiceInstance() *MatchService {
	return instance
}

func (svc *MatchService) AddPerson(person model.Person) int {
	person.Id = rand.Int()
	key := []byte("Male")
	if person.Gender == constant.Female {
		key = []byte("Female")
	}
	val := []byte(strconv.Itoa(person.Id))

	if err := svc.nuts.Update(
		func(tx *nutsdb.Tx) error {
			return tx.ZAdd(bucket, key, person.Height, val)
		}); err != nil {
		log.Fatal(err)
		return 0
	}

	return person.Id
}

func (svc *MatchService) RemovePerson(id int64) {
	// ok := svc.maleSet.Delete(id)
	// if !ok {
	// 	svc.femaleSet.Delete(id)
	// }
}

func (svc *MatchService) PersonMatch(id int64) {
	// ok := svc.maleSet.(id)
	// if !ok {
	// 	svc.femaleSet.Delete(id)
	// }
}
