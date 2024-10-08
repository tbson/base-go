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
  default = getenv("DB_URL")
}

variable "db_url_atlas" {
  type = string
  default = getenv("DB_URL_ATLAS")
}

env "gorm" {
  src = data.external_schema.gorm.url
  url = var.db_url
  dev = var.db_url_atlas
  migration {
    dir = "file://dbversioning/migration"
    revisions_schema = "public"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
