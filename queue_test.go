package loucore

import (
	"fmt"
	"github.com/joaopedrosgs/loucore/ent"
	"reflect"
	"testing"
	"time"
)

func TestCompleteQueueItem(t *testing.T) {
	fmt.Println(testEnv.queue)
	type args struct {
		queueItemId int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"CompleteQueue", args{testEnv.queue[0].ID}, false},
		{"CompleteQueue", args{-1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CompleteQueueItem(tt.args.queueItemId); (err != nil) != tt.wantErr {
				t.Errorf("CompleteQueueItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateQueueItem(t *testing.T) {
	type args struct {
		ownerId        int
		cityId         int
		constructionId int
		action         int
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.QueueItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateQueueItem(tt.args.ownerId, tt.args.cityId, tt.args.constructionId, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateQueueItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateQueueItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCityQueue(t *testing.T) {
	type args struct {
		cityId int
	}
	tests := []struct {
		name    string
		args    args
		want    []*ent.QueueItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCityQueue(tt.args.cityId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCityQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCityQueue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStructureQueue(t *testing.T) {
	type args struct {
		structureId int
	}
	tests := []struct {
		name    string
		args    args
		want    []*ent.QueueItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStructureQueue(tt.args.structureId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStructureQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStructureQueue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserQueue(t *testing.T) {
	type args struct {
		ownerId int
	}
	tests := []struct {
		name    string
		args    args
		want    []*ent.QueueItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserQueue(tt.args.ownerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserQueue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinTime(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinTime(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQueue(t *testing.T) {
	type args struct {
		cityId int
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.City
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateQueue(tt.args.cityId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQueue() got = %v, want %v", got, tt.want)
			}
		})
	}
}
