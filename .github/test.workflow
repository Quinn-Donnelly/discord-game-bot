workflow "New workflow" {
  on = "push"
  resolves = ["go"]
}

action "go" {
  uses = "docker://golang"
  runs = "go test"
}
