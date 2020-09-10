package test

/**
测试文件名必须以"_test.go"结尾；
测试函数名必须以“TestXxx”开始；
命令行下使用"go test"即可启动测试；
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
