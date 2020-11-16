package pkg

import (
	"github.com/joaopedrosgs/loucore/ent"
	"reflect"
	"testing"
	"time"
)

func TestCompleteQueueItem(t *testing.T) {
	type args struct {
		queueItemID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Simple case 1", args{testEnv.queue[0].ID}, false},
		{"CompleteQueueItem should fail because the queueItemID doesn't exists", args{-1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CompleteQueueItem(tt.args.queueItemID); (err != nil) != tt.wantErr {
				t.Errorf("CompleteQueueItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateQueueItem(t *testing.T) {
	type args struct {
		ownerID        int
		cityID         int
		constructionID int
		action         int
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.QueueItem
		wantErr bool
	}{
		{"Simple case 1", args{testEnv.accounts[0].ID, testEnv.cities[0].ID, testEnv.constructions[0].ID, 1}, nil, false},
		{"Create queue should fail because the city doesn't exists", args{-1, -1, -1, -1}, nil, true},
		{"Create queue should fail because the city already have the queue full", args{testEnv.accounts[1].ID, testEnv.cities[3].ID, testEnv.constructions[7].ID, 1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qi, err := CreateQueueItem(tt.args.ownerID, tt.args.cityID, tt.args.constructionID, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateQueueItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if qi != nil {
				_ = DeleteQueueItem(qi.ID)
			}
		})
	}
}

func TestGetCityQueue(t *testing.T) {
	type args struct {
		cityID int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantErr bool
	}{
		{"Simple case 1", args{testEnv.cities[0].ID}, 2, false},
		{"Simple case 2", args{testEnv.cities[1].ID}, 3, false},
		{"Simple case 3	", args{testEnv.cities[3].ID}, 10, false},
		{"Get city queue should return 0 since the city doesn't exists", args{-1}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCityQueue(tt.args.cityID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCityQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("GetCityQueue() got = %v, want %v", len(got), tt.wantLen)
			}
		})
	}
}

func TestGetStructureQueue(t *testing.T) {
	type args struct {
		structureID int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantErr bool
	}{
		{"Simple case 1", args{testEnv.constructions[0].ID}, 1, false},
		{"Simple case 1", args{testEnv.constructions[1].ID}, 1, false},
		{"Get structure queue should return 0 since the city doesn't exists", args{-1}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStructureQueue(tt.args.structureID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStructureQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantLen != len(got) {
				t.Errorf("GetStructureQueue() got = %v, want %v", len(got), tt.wantLen)
			}
		})
	}
}

func TestGetUserQueue(t *testing.T) {
	type args struct {
		ownerID int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantErr bool
	}{
		{"Simple case 1", args{testEnv.accounts[0].ID}, 5, false},
		{"Get user queue should return 0 because user doesn't exists", args{-1}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserQueue(tt.args.ownerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("GetUserQueue() got = %v, want %v", len(got), tt.wantLen)
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
		{
			"t1<t2",
			args{time.Now(), time.Now().Add(time.Duration(10) * time.Second)},
			time.Now(),
		}, {
			"t1>t2",
			args{time.Now().Add(time.Duration(10) * time.Second), time.Now()},
			time.Now(),
		}, {
			"t1==t2",
			args{time.Now(), time.Now()},
			time.Now(),
		},
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
		cityID int
		until  time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Simple case 1", args{testEnv.cities[0].ID, time.Now()}, false},
		{"Simple case 2", args{testEnv.cities[2].ID, time.Now().AddDate(0, 0, 1)}, false},
		{"Update queue should fail since the city doesn't exists", args{-1, time.Now()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateQueue(tt.args.cityID, tt.args.until); (err != nil) != tt.wantErr {
				t.Errorf("UpdateQueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteQueueItem(t *testing.T) {
	qi1, _ := CreateQueueItem(testEnv.accounts[0].ID, testEnv.cities[0].ID, testEnv.constructions[0].ID, 1)
	qi2, _ := CreateQueueItem(testEnv.accounts[0].ID, testEnv.cities[0].ID, testEnv.constructions[0].ID, 1)

	type args struct {
		queueID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Simple case 1", args{qi1.ID}, false},
		{"Simple case 2", args{qi2.ID}, false},
		{"DeleteQueueItem should fail because the QueueItem ID doesn't exists", args{-1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteQueueItem(tt.args.queueID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteQueueItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
