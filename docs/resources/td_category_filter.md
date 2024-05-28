---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_td_category_filter Resource - terraform-provider-bloxone"
subcategory: "Threat Defense"
description: |-
  Creates and Manages Category Filters.
---

# bloxone_td_category_filter (Resource)

Creates and Manages Category Filters.

## Example Usage

```terraform
resource "bloxone_td_category_filter" "example" {
  name       = "example_category_filter"
  categories = ["Tutoring", "College"]

  # Other optional fields
  description = "Example of a Category Filter"
  tags = {
    site = "Site A"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `categories` (List of String) The list of content category names that falls into this category filter.
- `name` (String) The name of the category filter.

### Optional

- `description` (String) The brief description for the category filter.
- `tags` (Map of String) Enables tag support for resource where tags attribute contains user-defined key value pairs

### Read-Only

- `created_time` (String) The time when this Category Filter object was created.
- `id` (Number) The Category Filter object identifier.
- `policies` (List of String) The list of security policy names with which the category filter is associated.
- `updated_time` (String) The time when this Category Filter object was last updated.