workflow "Merged Pull Request" {
    on = "pull_request"
    resolves = ["dispatch"]
}
action "dispatch" {
    uses = "./"
}