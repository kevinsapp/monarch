package fileutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

const (
	tmpDir     string = "tmp"
	tmpTestDir string = tmpDir + "/test/"
)

func TestMain(m *testing.M) {
	// Setup
	MkdirP(tmpTestDir)

	// Execute tests.
	i := m.Run()

	// Teardown
	os.RemoveAll(tmpDir) // Do cleanup

	// Exit
	os.Exit(i)
}

// Unit test CreateAndWriteString
func TestCreateAndWriteString(t *testing.T) {
	ts := time.Now().UnixNano()
	n := fmt.Sprintf("%d_test.txt", ts)
	path := tmpTestDir + n

	// Write string to the file (any string would do)
	str := `Some text to write to a file.`
	err := CreateAndWriteString(path, str)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	// Read in content from migration file to a buffer.
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	// Verify string
	exp := str
	act := string(buf)

	if exp != act {
		t.Errorf("want %q\n; got %q\n", exp, act)
	}
}

// Unit test Mkdir
func TestMkdir(t *testing.T) {
	// Call Mkdir to create a directory.
	dn := `test_directory`
	path := tmpTestDir + dn
	err := MkdirP(path)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(path)

	// Open the directory and get its description.
	dir, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	info, err := dir.Stat()
	if err != nil {
		t.Fatal(err)
	}

	// Verify that it is a directory
	expDir := true
	actDir := info.IsDir()
	if expDir != actDir {
		t.Errorf("want %t; got %t", expDir, actDir)
	}

	// Verify directory has the correct name.
	expName := dn
	actName := info.Name()
	if expName != actName {
		t.Errorf("want %q; got %q", expName, actName)
	}
}

// Unit test ReadFileAsString
func TestReadFileAsString(t *testing.T) {
	ts := time.Now().UnixNano()
	n := fmt.Sprintf("%d_test.txt", ts)
	path := tmpTestDir + n

	// Write string to the file (any string would do)
	str := `Some text to write to a file.`
	err := CreateAndWriteString(path, str)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	// Run ReadFileAsString.
	rstr, err := ReadFileAsString(path)
	if err != nil {
		t.Fatal(err)
	}

	// Verify string
	exp := str
	act := rstr

	if exp != act {
		t.Errorf("want %q\n; got %q\n", exp, act)
	}
}
