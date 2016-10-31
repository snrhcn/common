package datautil

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"common/fileutil"
)

const testdbdir = "test"

const invalidFileName = "**" + string(0x0)

func TestMain(m *testing.M) {
	flag.Parse()

	// Setup
	if res, _ := fileutil.PathExists(testdbdir); res {
		os.RemoveAll(testdbdir)
	}

	err := os.Mkdir(testdbdir, 0770)
	if err != nil {
		fmt.Print("Could not create test directory:", err.Error())
		os.Exit(1)
	}

	// Run the tests
	res := m.Run()

	// Teardown
	err = os.RemoveAll(testdbdir)
	if err != nil {
		fmt.Print("Could not remove test directory:", err.Error())
	}

	os.Exit(res)

}

func TestPersistentMap(t *testing.T) {

	// Test main scenario

	pm, err := NewPersistentMap(testdbdir + "/testmap.map")
	if err != nil {
		t.Error(nil)
		return
	}

	pm.Data["test1"] = "test1data"
	pm.Data["test2"] = "test2data"

	pm.Flush()

	pm2, err := LoadPersistentMap(testdbdir + "/testmap.map")

	if len(pm2.Data) != 2 {
		t.Error("Unexpected size of map")
		return
	}

	if pm.Data["test1"] != "test1data" || pm.Data["test2"] != "test2data" {
		t.Error("Unexpected data in map:", pm.Data)
		return
	}

	// Test error cases

	pm, err = NewPersistentMap(invalidFileName)
	if err == nil {
		t.Error("Unexpected result of new map")
		return
	}

	pm, err = LoadPersistentMap(invalidFileName)
	if err == nil {
		t.Error("Unexpected result of new map")
		return
	}

	pm = &PersistentMap{invalidFileName, make(map[string]string)}
	if err := pm.Flush(); err == nil {
		t.Error("Unexpected result of new map")
		return
	}
}
