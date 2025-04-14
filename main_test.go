package main

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

var testFiles = getTestFiles("./testdata/")

func getTestFiles(testFolder string) []string {
	var res []string

	files, err := os.ReadDir(testFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		res = append(res, testFolder+file.Name())
	}
	return res
}

func TestHeadCommand(t *testing.T) {
	for _, testFile := range testFiles {
		cmd := exec.Command("go", "run", ".", testFile)
		got, err := cmd.Output()
		if err != nil {
			t.Fatalf("Command failed with error: %v", err)
		}

		unixCmd := exec.Command("head", testFile)
		want, err := unixCmd.Output()
		if err != nil {
			t.Fatalf("Command failed with error: %v", err)
		}

		gotString := string(got)
		wantString := string(want)

		if gotString != wantString {
			t.Errorf("\tEXPECTED: %q\n\tGOT: %q\n", wantString, gotString)
		}
	}
}

func TestHeadCommandBytes(t *testing.T) {
	for _, testFile := range testFiles {
		cmd := exec.Command("go", "run", ".", "-c", "30", testFile)
		got, err := cmd.Output()
		if err != nil {
			t.Fatalf("Command failed with error: %v", err)
		}

		unixCmd := exec.Command("head", "-c", "30", testFile)
		want, err := unixCmd.Output()
		if err != nil {
			t.Fatalf("Command failed with error: %v", err)
		}

		if string(got) != string(want) {
			t.Errorf("\tEXPECTED: %q\n\tGOT: %q\n", string(want), string(got))
		}
	}
}

func TestHeadCommandFiveLines(t *testing.T) {
	for _, testFile := range testFiles {
		cmd := exec.Command("go", "run", ".", "-n", "5", testFile)
		got, err := cmd.Output()
		if err != nil {
			t.Fatalf("Command failed with error: %v", err)
		}

		unixCmd := exec.Command("head", "-n", "5", testFile)
		want, err := unixCmd.Output()
		if err != nil {
			t.Fatalf("Command failed with error: %v", err)
		}

		if string(got) != string(want) {
			t.Errorf("\tEXPECTED: %q\n\tGOT: %q\n", string(want), string(got))
		}
	}
}

func TestHeadMultipleFiles(t *testing.T) {

	cmd := exec.Command("./cchead", testFiles...)
	got, err := cmd.Output()
	if err != nil {
		t.Fatalf("Command %s failed with error: %v, got: %s", cmd.String(), err, string(got))
	}

	unixCmd := exec.Command("head", testFiles...)
	want, err := unixCmd.Output()
	if err != nil {
		t.Fatalf("Command failed with error: %v", err)
	}

	if string(got) != string(want) {
		t.Errorf("\tEXPECTED: %q\n\tGOT: %q\n", string(want), string(got))
	}
}
