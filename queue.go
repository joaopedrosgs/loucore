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
	c, err := client.City.Query().WithQueue().WithConstructions(func(query *ent.ConstructionQuery) {
		query.Where(construction.ID(constructionId))
	}).Where(city.IDEQ(cityId)).First(context.Background())

	if err != nil {
		return nil, fmt.Errorf("(CreateQueueItem) Failed to find construction: %v", err)
	}
	if c.Edges.Constructions[0].Level >= 10 {
		return nil, errors.New("(CreateQueueItem) Construction already max level!")
	}
	if len(c.Edges.Queue)>=10 {
		return nil, errors.New("(CreateQueueItem) Queue already full")
	}


	duration:=10
	endsAt := c.QueueEndsAt.Add(time.Duration(duration)*time.Second)
	startsAt:= c.QueueEndsAt
	c.Update().SetQueueEndsAt(endsAt).Save(context.Background())

	return client.QueueItem.
		Create().
		SetOwnerID(ownerId).
		SetCityID(cityId).
		SetConstructionID(constructionId).
		SetAction(action).
		SetStartAt(startsAt).
		SetCompletion(endsAt).
		SetDuration(duration).

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
func UpdateQueue(cityId int) (*ent.City, error) {
	city, err := client.City.Query().WithQueue(func(query *ent.QueueItemQuery) {
		query.Order(ent.Asc(queueitem.FieldCompletion)).Limit(10)

	}).Where(city.ID(cityId)).First(context.Background())
	if err != nil {
		return nil, fmt.Errorf("(UpdateQueue) Failed to get the city and the queue: %v", err)
	}
	lastCompletion := MinTime(city.Edges.Queue[0].StartAt, time.Now())
	for _, qi := range city.Edges.Queue {
		startTime := lastCompletion
		_, err :=qi.Update().SetStartAt(startTime).Save(context.Background())
		if err != nil {
			return nil, fmt.Errorf("(UpdateQueue) Failed to set startsAt: %v", err)
		}
		lastCompletion = lastCompletion.Add(time.Duration(qi.Duration) * time.Second)
	}
	return city.Update().SetQueueEndsAt(lastCompletion).Save(context.Background())
}
func MinTime(t1, t2 time.Time) time.Time {
	if t1.After(t2) {
		return t2
	}
	return t1
}
