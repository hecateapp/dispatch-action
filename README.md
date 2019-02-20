<p align="center"> 
<img src="https://hecate.co/hecate-logo-168.png">
</p>

# Hecate Dispatch via GitHub Action

GitHub action to send emails notifying of merged pull requests. Super handy for keeping your stakeholders who aren't on GitHub informed of releases.

Free version limited to 5 emails per user/organization per day.

Pro version is a GitHub app with no notification limits, allows org wide (rather than per repo) configuration, can do daily rollup newsletters, notify via slack, and more.

Learn more on the [Hecate Dispatch product page](https://hecate.co/products/dispatch)

## Setup

Example workflow:

```hcl
workflow "Email stakeholders on release" {
  on = "pull_request"
  resolves = ["hecateapp/dispatch-action"]
}

action "hecateapp/dispatch-action" {
  uses = "hecateapp/dispatch-action@v1.0.0"
  secrets = ["GITHUB_TOKEN"]
  env = {
    EMAILS = "your.email@some.domain, another.email@some.domain"
  }
}
```
