schema_version = 1

project {
  license        = "Apache-2.0"
  copyright_year = 2024
  copyright_holder = "RoseSecurity"

  header_ignore = [
    # tests used within documentation (prose)
    "test/**",

    # GitHub issue template configuration
    ".github/ISSUE_TEMPLATE/*.yml",

    # golangci-lint tooling configuration
    ".golangci.yml",

    # GoReleaser tooling configuration
    ".goreleaser.yml",
  ]
}
