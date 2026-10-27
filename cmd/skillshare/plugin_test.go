package main

import "testing"

func TestPluginOptions(t *testing.T) {
	o, err := parsePluginOptions([]string{"add", "owner/repo", "--plugin", "demo", "--target", "claude", "--target", "codex", "--json"})
	if err != nil || o.request.Source != "owner/repo" || len(o.request.Targets) != 2 || !o.json {
		t.Fatalf("%+v %v", o, err)
	}
	for _, args := range [][]string{{"unknown"}, {"add", "source", "--target"}, {"remove", "a", "b"}, {"list", "--oops"}, {"update", "demo", "--all"}, {"update", "--name", "demo", "-a"}, {"update", "--all", "--source-ref", "main"}, {"sync", "--all"}} {
		if _, err := parsePluginOptions(args); err == nil {
			t.Fatalf("accepted %v", args)
		}
	}
	for _, flag := range []string{"--all", "-a"} {
		u, err := parsePluginOptions([]string{"update", flag, "--target", "claude"})
		if err != nil || !u.request.All || u.request.Name != "" || len(u.request.Targets) != 1 {
			t.Fatalf("update %s: %+v %v", flag, u, err)
		}
	}
}

func TestPluginAntigravityAlias(t *testing.T) {
	o, err := parsePluginOptions([]string{"add", "./demo", "--target", "agy", "--no-tui"})
	if err != nil || o.request.Targets[0] != "antigravity" {
		t.Fatalf("alias: %+v %v", o, err)
	}
}
