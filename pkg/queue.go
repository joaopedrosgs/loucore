package pkg

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

func CreateQueueItem(ownerID, cityID, constructionID, action int) (*ent.QueueItem, error) {
	c, err := client.City.Query().WithQueue(func(query *ent.QueueItemQuery) {
		query.Order(ent.Asc(queueitem.FieldPosition))
	}).WithConstructions(func(query *ent.ConstructionQuery) {
		query.Where(construction.ID(constructionID))
	}).Where(city.IDEQ(cityID)).First(context.Background())

	if err != nil {
		return nil, fmt.Errorf("(CreateQueueItem) Failed to find construction: %w", err)
	}
	if c.Edges.Constructions[0].Level >= 10 {
		return nil, errors.New("(CreateQueueItem) Construction already max level!")
	}
	if len(c.Edges.Queue) >= 10 {
		return nil, errors.New("(CreateQueueItem) Queue already full")
	} else if len(c.Edges.Queue) == 0 {
		_, _ = c.Update().SetQueueStartedAt(time.Now()).Save(context.Background())
	}
	nextPosition := len(c.Edges.Queue)
	duration := 10

	return client.QueueItem.
		Create().
		SetOwnerID(ownerID).
		SetCityID(cityID).
		SetConstructionID(constructionID).
		SetPosition(nextPosition).
		SetAction(action).
		SetDuration(duration).
		Save(context.Background())
}
func CompleteQueueItem(queueItemID int) error {
	q, err := client.QueueItem.Query().WithCity().WithConstruction().Where(queueitem.ID(queueItemID)).Only(context.Background())
	if err != nil {
		return fmt.Errorf("(CompleteQueueItem) Failed to query queue item: %w", err)
	}
	_, err = q.Edges.Construction.Update().AddLevel(1).Save(context.Background())
	if err != nil {
		return fmt.Errorf("(CompleteQueueItem) Failed to add level: %w", err)
	}
	_, err = q.Edges.City.Update().SetQueueStartedAt(time.Now()).Save(context.Background())
	return err
}

func GetUserQueue(ownerID int) ([]*ent.QueueItem, error) {
	return client.QueueItem.
		Query().Where(queueitem.HasOwnerWith(user.IDEQ(ownerID))).All(context.Background())
}
func DeleteQueueItem(queueID int) error {
	return client.QueueItem.DeleteOneID(queueID).Exec(context.Background())
}

func GetCityQueue(cityID int) ([]*ent.QueueItem, error) {
	return client.QueueItem.
		Query().Where(queueitem.HasCityWith(city.IDEQ(cityID))).All(context.Background())
}

func GetStructureQueue(structureID int) ([]*ent.QueueItem, error) {
	return client.QueueItem.
		Query().Where(queueitem.HasConstructionWith(construction.IDEQ(structureID))).All(context.Background())
}
func UpdateQueue(cityID int, until time.Time) error {
	c, err := client.City.Query().WithQueue(func(query *ent.QueueItemQuery) {
		query.Order(ent.Asc(queueitem.FieldPosition))
	}).Where(city.ID(cityID)).First(context.Background())
	if err != nil {
		return fmt.Errorf("(UpdateQueue) Failed to get the city and the queue: %w", err)
	}
	secondsToUse := int(until.Sub(c.QueueStartedAt).Seconds())
	for _, qi := range c.Edges.Queue {
		secondsLeft := secondsToUse - qi.Duration

		if secondsLeft < 0 {
			break
		}
		err := CompleteQueueItem(qi.ID)
		fmt.Println("a")
		if err != nil {
			return fmt.Errorf("(UpdateQueue) Failed to CompleteQueueItem: %w", err)
		}
		secondsToUse = secondsLeft
	}
	_, err = c.Update().SetQueueStartedAt(time.Now()).Save(context.Background())
	return err
}
func MinTime(t1, t2 time.Time) time.Time {
	if t1.After(t2) {
		return t2
	}
	return t1
}
