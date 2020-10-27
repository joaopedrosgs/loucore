package loucore

import (
	"context"
	"errors"
	"fmt"
	"github.com/joaopedrosgs/loucore/ent"
	"github.com/joaopedrosgs/loucore/ent/city"
	"github.com/joaopedrosgs/loucore/ent/construction"
	"github.com/joaopedrosgs/loucore/ent/queueitem"
	"github.com/joaopedrosgs/loucore/ent/user"
	"time"
)

func CreateQueueItem(ownerId, cityId, constructionId, action int) (*ent.QueueItem, error) {
	c, err := client.Construction.Get(context.Background(), constructionId)
	if err != nil {
		return nil, fmt.Errorf("(CreateQueueItem) Failed to find construction: %v", err)
	}
	if c.Level >= 10 {
		return nil, errors.New("(CreateQueueItem) Construction already max level!")
	}

	return client.QueueItem.
		Create().
		SetOwnerID(ownerId).
		SetCityID(cityId).
		SetConstructionID(constructionId).
		SetAction(action).
		Save(context.Background())

}
func CompleteQueueItem(queueItemId int) error {

	_, err := client.Construction.Update().Where(
		construction.HasQueueWith(
			queueitem.IDEQ(queueItemId),
			queueitem.CompletionLTE(
				time.Now()))).
		AddLevel(1).
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("(CompleteQueueItem) Failed to add level: %v", err)
	}
	return client.QueueItem.DeleteOneID(queueItemId).Exec(context.Background())

}

func GetUserQueue(ownerId int) ([]*ent.QueueItem, error) {
	return client.QueueItem.
		Query().Where(queueitem.HasOwnerWith(user.IDEQ(ownerId))).All(context.Background())
}

func GetCityQueue(cityId int) ([]*ent.QueueItem, error) {
	return client.QueueItem.
		Query().Where(queueitem.HasCityWith(city.IDEQ(cityId))).All(context.Background())
}

func GetStructureQueue(structureId int) ([]*ent.QueueItem, error) {
	return client.QueueItem.
		Query().Where(queueitem.HasConstructionWith(construction.IDEQ(structureId))).All(context.Background())
}
func UpdateQueue(cityId int) error {
	city, err := client.City.Query().WithQueue(func(query *ent.QueueItemQuery) {
		query.Order(ent.Asc(queueitem.FieldOrder)).Limit(10)

	}).Where(city.ID(cityId)).First(context.Background())
	if err != nil {
		return err
	}
	startTime := MinTime(city.Edges.Queue[0].StartAt, time.Now())
	lastDuration := 0
	var lastCompletion = MinTime(city.Edges.Queue[0].StartAt, time.Now())
	for _, qi := range city.Edges.Queue {
		startTime = lastCompletion.Add(time.Duration(lastDuration) * time.Second)
		qi.Update().SetStartAt(startTime)
		lastDuration = qi.Duration
		lastCompletion = lastCompletion.Add(time.Duration(lastDuration) * time.Second)
	}
	return nil
}
func MinTime(t1, t2 time.Time) time.Time {
	if t1.After(t2) {
		return t2
	}
	return t1
}
