---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_dhcp_option_group Resource - terraform-provider-bloxone"
subcategory: ""
description: |-
  
---

# bloxone_dhcp_option_group (Resource)



## Example Usage

```terraform
resource "bloxone_dhcp_option_group" "example" {
  name     = "example_dhcp_option_group"
  protocol = "ip4"
}

resource "bloxone_dhcp_option_space" "option_space" {
  name     = "option_space"
  protocol = "ip4"
}

resource "bloxone_dhcp_option_code" "option_code" {
  code         = 234
  name         = "option_code"
  option_space = bloxone_dhcp_option_space.option_space.id
  type         = "boolean"
}

resource "bloxone_dhcp_option_group" "example_with_options" {
  name     = "example_dhcp_option_group_with_options"
  protocol = "ip6"

  # Other Optional Fields
  dhcp_options = [
    {
      type         = "option"
      option_code  = bloxone_dhcp_option_code.option_code.id
      option_value = "true"
    }
  ]
  comment = "dhcp option group"
  tags = {
    site = "Test Site"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the option group. Must contain 1 to 256 characters. Can include UTF-8.
- `protocol` (String) The type of protocol (_ip4_ or _ip6_).

### Optional

- `comment` (String) The description for the option group. May contain 0 to 1024 characters. Can include UTF-8.
- `dhcp_options` (Attributes List) The list of DHCP options for the option group. May be either a specific option or a group of options. (see [below for nested schema](#nestedatt--dhcp_options))
- `tags` (Map of String) The tags for the option group in JSON format.

### Read-Only

- `created_at` (String) Time when the object has been created.
- `id` (String) The resource identifier.
- `updated_at` (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.

<a id="nestedatt--dhcp_options"></a>
### Nested Schema for `dhcp_options`

Optional:

- `group` (String) The resource identifier.
- `option_code` (String) The resource identifier.
- `option_value` (String) The option value.
- `type` (String) The type of item.  Valid values are: * _group_ * _option_