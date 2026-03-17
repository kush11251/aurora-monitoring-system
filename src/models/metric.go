package models

import (
    "encoding/json"
)

type Metric struct {
    Name  string `json:"name"`
    Value int    `json:"value"`
}
