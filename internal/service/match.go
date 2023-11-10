package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"log"

	"github.com/CoulsonChen/BitoAPI/internal/constant"
	"github.com/CoulsonChen/BitoAPI/internal/model"
	redisutil "github.com/CoulsonChen/BitoAPI/third-party/redis"
	"github.com/redis/go-redis/v9"
)

var rkey_user_map = "UserMap"
var rkey_male_pool = "MalePool"
var rkey_female_pool = "FemalePool"

type MatchService struct {
	rdb *redis.Client
}

var instance *MatchService

func InitMatchService() {
	instance = &MatchService{
		rdb: redisutil.GetInstance(),
	}
}

func GetMatchServiceInstance() *MatchService {
	return instance
}

func (svc *MatchService) AddPerson(person model.Person) ([]model.Person, error) {
	person.Id = rand.Int()
	rkey := getGenderKey(person.Gender)

	// add to sorted set
	ctx := context.TODO()
	zset := redis.Z{Score: person.Height, Member: person.Id}
	_, err := svc.rdb.ZAdd(ctx, rkey, zset).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// add to hash map
	err = svc.setPerson(&person)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return svc.QueryMatches(person.Id, 0)
}

func (svc *MatchService) RemovePerson(id int) {
	ctx := context.TODO()
	person, err := svc.getPerson(id)
	if err != nil {
		log.Println(err)
		return
	}

	rkey := getGenderKey(person.Gender)
	_, err = svc.rdb.ZRem(ctx, rkey, id).Result()
	if err != nil {
		log.Println(err)
	}
	_, err = svc.rdb.HDel(ctx, rkey_user_map, strconv.Itoa(id)).Result()
	if err != nil {
		log.Println(err)
	}
}

func (svc *MatchService) QueryMatches(id int, topN int) ([]model.Person, error) {
	result := make([]model.Person, 0, topN)

	ctx := context.TODO()
	// get person info
	person, err := svc.getPerson(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// get top n matched person ids
	rkey := getMatchKey(person.Gender)
	high := "+inf"
	low := "-inf"
	not_equal := "("
	if person.Gender == constant.Male {
		high = not_equal + strconv.Itoa(int(person.Height))
	} else {
		low = not_equal + strconv.Itoa(int(person.Height))
	}
	ids, err := svc.rdb.ZRevRangeByScore(ctx, rkey, &redis.ZRangeBy{
		Max: high, Min: low, Count: int64(topN),
	}).Result()
	if err != nil {
		return nil, err
	}
	// get person info
	datas, err := svc.rdb.HMGet(ctx, rkey_user_map, ids...).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, data := range datas {
		dataStr := fmt.Sprintf("%v", data)
		var p model.Person
		err = json.Unmarshal([]byte(dataStr), &p)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, p)
	}

	return result, nil
}

func (svc *MatchService) PersonMatch(id int, id_to_match int) {
	// get perosn info
	person, err := svc.getPerson(id)
	if err != nil {
		log.Println(err)
		return
	}
	person_to_match, err := svc.getPerson(id_to_match)
	if err != nil {
		log.Println(err)
		return
	}
	// verify
	if person.Dates <= 0 {
		log.Println("user has no quota")
		return
	}
	if person.Gender == constant.Male && person.Height < person_to_match.Height {
		log.Println("it's ok")
	}
	if person_to_match.Dates <= 0 {
		log.Println("person wanna match has no quota")
		return
	}
	if person.Gender == constant.Female && person.Height > person_to_match.Height {
		log.Println("it's ok")
	}
	if person.Gender == person_to_match.Gender {
		log.Println("Cool but not this time")
		return
	}
	// update user map
	log.Println(person)
	log.Println(person_to_match)
	person.Dates -= 1
	err = svc.setPerson(person)
	if err != nil {
		log.Println(err)
		return
	}
	person_to_match.Dates -= 1
	err = svc.setPerson(person_to_match)
	if err != nil {
		log.Println(err)
		return
	}
	// sync matching pool (sorted set)
	ctx := context.TODO()
	var key string
	if person.Dates == 0 {
		key = getGenderKey(person.Gender)
		svc.rdb.ZRem(ctx, key, person.Id)
	}
	if person_to_match.Dates == 0 {
		key = getGenderKey(person_to_match.Gender)
		svc.rdb.ZRem(ctx, key, person_to_match.Id)
	}

}

func getGenderKey(gender constant.Gender) string {
	if gender == constant.Male {
		return rkey_male_pool
	} else {
		return rkey_female_pool
	}
}

func getMatchKey(gender constant.Gender) string {
	if gender == constant.Male {
		return rkey_female_pool
	} else {
		return rkey_male_pool
	}
}

func (svc *MatchService) getPerson(id int) (*model.Person, error) {
	ctx := context.TODO()
	json_p, err := svc.rdb.HGet(ctx, rkey_user_map, strconv.Itoa(id)).Result()
	if err != nil {
		return nil, err
	}
	var person model.Person
	err = json.Unmarshal([]byte(json_p), &person)
	if err != nil {
		return nil, err
	}

	person.Id = id
	return &person, nil
}

func (svc *MatchService) setPerson(person *model.Person) error {
	ctx := context.TODO()
	val, err := json.Marshal(person)
	if err != nil {
		return err
	}
	_, err = svc.rdb.HSet(ctx, rkey_user_map, person.Id, val).Result()
	if err != nil {
		return err
	}

	return nil
}
