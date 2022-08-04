terraform {
  required_providers {
    baseball = {
      source = "brenthc/baseball"
    }
  }
}

resource "baseball_bat" "smack" {
  targets = [
    "clowns",
    "goobers",
  ]
}

output "smacked" {
  value = baseball_bat.smack.results
}
