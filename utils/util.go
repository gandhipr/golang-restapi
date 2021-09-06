package utils

import "github.com/gin-gonic/gin"

type Metadata struct {
	Title       string       `yaml:"title,omitempty" validate:"required,title"`
	Version     string       `yaml:"version,omitempty" validate:"required,version"`
	Maintainers []Maintainer `yaml:"maintainers,omitempty" validate:"required,dive"`
	Company     string       `yaml:"company" validate:"required"`
	Website     string       `yaml:"website,omitempty" validate:"required"`
	Source      string       `yaml:"source,omitempty" validate:"required"`
	License     string       `yaml:"license,omitempty" validate:"required"`
	Description string       `yaml:"description,omitempty" validate:"required"`
}

type Maintainer struct {
	Name  string `yaml:"name,omitempty" validate:"required"`
	Email string `yaml:"email,omitempty" validate:"required,email"`
}

type GinContext struct {
	C *gin.Context
}
