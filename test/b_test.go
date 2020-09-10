package test

/**
 1.文件名 xxx_test.go
 2.函数名 Testxxx 开头，即可测试
 */
import (
	"testing"
)

func TestHello(t *testing.T) {
	say("hello world")
}

func say(b string) string {
	return b
}
