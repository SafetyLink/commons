package snowflake

import (
	"fmt"
	"testing"
)

func TestGenerateSnowflake(t *testing.T) {
	fmt.Println(GenerateSnowflakeID())
}
