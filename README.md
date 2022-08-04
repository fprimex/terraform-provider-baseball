How do

```
go build

# edit dev.tfrc to suit
export TF_CLI_CONFIG_FILE="$(pwd)/dev.tfrc

terraform apply
```

Output

```
$ terraform apply
╷
│ Warning: Provider development overrides are in effect
│
│ The following provider development overrides are set in the CLI configuration:
│  - brenthc/baseball in /Users/brent/code/terraform-provider-baseball
│
│ The behavior may therefore not match any released version of the provider and applying changes may cause the state to
│ become incompatible with published releases.
╵

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the
following symbols:
  + create

Terraform will perform the following actions:

  # baseball_bat.smack will be created
  + resource "baseball_bat" "smack" {
      + id      = (known after apply)
      + results = (known after apply)
      + targets = [
          + "clowns",
          + "goobers",
        ]
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  + smacked = (known after apply)

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

baseball_bat.smack: Creating...
baseball_bat.smack: Creation complete after 0s [id=2899211730713887727]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.

Outputs:

smacked = tolist([
  "clowns",
  "goobers",
])
```
