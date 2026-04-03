package output

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSON 将数据以缩进 JSON 格式写入 stdout
func JSON(data interface{}) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

// Errorf 将错误信息写入 stderr
func Errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
}
