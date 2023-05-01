package snowflake

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var counter int64 = 0
var workerID int64 = 0
var systemID int64 = 0

func GenerateSnowflakeID() int64 {
	if counter == 4095 {
		counter = 0
	}

	currentUnixTime := time.Now().UnixNano() / int64(time.Millisecond)
	binaryUnixTime := fmt.Sprintf("%048b", currentUnixTime)
	workerIDBinaryValue := strconv.FormatInt(workerID, 2)
	systemIDBinaryValue := strconv.FormatInt(systemID, 2)
	counterBinaryValue := fmt.Sprintf("%014b", counter)

	var binarySnowflakeID strings.Builder

	binarySnowflakeID.WriteString(binaryUnixTime)
	binarySnowflakeID.WriteString(workerIDBinaryValue)
	binarySnowflakeID.WriteString(systemIDBinaryValue)
	binarySnowflakeID.WriteString(counterBinaryValue)

	counter++

	snowflake, _ := strconv.ParseInt(binarySnowflakeID.String(), 2, 64)

	return snowflake
}
