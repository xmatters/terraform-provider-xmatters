// Example of an xMatters group roster resource with advanced member configuration
resource "xmatters_group_roster" "advanced_roster" {
  id = "481086d8-357a-4279-b7d5-d7dce48fcd12"
  members = [
    {
      id          = "12345678-1234-1234-1234-123456789012"
      member_type = "PERSON"
    },
    {
      id          = "87654321-4321-4321-4321-210987654321"
      member_type = "PERSON"
    },
    {
      id          = "abcdefab-cdef-abcd-efab-cdefabcdefab"
      member_type = "GROUP"
    }
  ]
}
