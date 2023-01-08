// Package question
package question

import (
	"github.com/go-zookeeper/zk"
	"testing"
	"time"
)

//git.code.oa.com,git.woa.com
func TestZkTemp(t *testing.T) {
	zkConn, _, err := zk.Connect([]string{"30.43.42.120:2181"}, time.Second*15)
	if err != nil {
		t.Fatal(err)
	}
	path := "/for_teset/a"
	if p, err := zkConn.Create(path, []byte{1, 2, 3, 4}, zk.FlagEphemeral, WorldACL(zk.PermAll)); err != nil {
		t.Fatalf("Create returned error: %+v", err)
	} else if p != path {
		t.Fatalf("Create returned different path '%s' != '%s'", p, path)
	}

	time.Sleep(10 * time.Second)

	defer zkConn.Close()
}

func WorldACL(perms int32) []zk.ACL {
	return []zk.ACL{{perms, "world", "anyone"}}
}
