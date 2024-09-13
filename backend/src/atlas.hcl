data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./dbversioning/loader",
  ]
}

variable "db_url" {
  type = string
  default = "postgres://postgres:postgres@basecode_db:5432/basecode_dev?sslmode=disable"
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = var.db_url
  url = var.db_url
  migration {
    dir = "file://dbversioning/migration"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
