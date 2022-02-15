package eleventh

import (
	"context"
	"testing"
)

func TestMockShard_Query(t *testing.T) {
	shard := newMockShard()
	_, err := shard.Query(context.Background(), "SELECT * FROM `my_app`.`user` WHERE `id`=?", 1)
	if err != nil {
		t.Fatal(err)
	}
}
