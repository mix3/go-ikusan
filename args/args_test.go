package args

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	defer func() {
		if err := recover(); err != nil {
			if !reflect.DeepEqual(a, b) {
				t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
			}
		}
	}()
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func Test_DefaultArgs(t *testing.T) {
	Parse([]string{})
	r := GetConfig()
	expect(t, r.HttpHost(), "127.0.0.1")
	expect(t, r.HttpPort(), uint(19300))
	expect(t, r.ReverseProxy(), []string{})
	expect(t, r.Mount(), "")
	expect(t, r.IrcServer(), "")
	expect(t, r.IrcPort(), uint(6667))
	expect(t, r.IrcKeyword(), "")
	expect(t, r.IrcNickname(), "ikusan")
	expect(t, r.IrcUser(), "ikusan")
	expect(t, r.IrcPostInterval(), uint(2))
	expect(t, r.EnableSsl(), false)
	expect(t, r.InsecureSkipVerify(), false)
	expect(t, r.IrcReconnectInterval(), uint(3))
	expect(t, r.NoPostWithJoin(), false)
	expect(t, r.MaxLength(), uint(0))
	expect(t, r.Help(), false)
	expect(t, r.IsValid(), false)
}

func Test_SetShortArgs(t *testing.T) {
	Parse([]string{
		"-o", "127.0.0.2",
		"-p", "19301",
		"-r", "192.168.0.0.2", "-r", "192.168.0.0.3",
		"-m", "ikusan",
		"-S", "192.168.0.0.1",
		"-P", "6668",
		"-K", "fever",
		"-N", "iku",
		"-U", "iku",
		"-i", "1",
		//		"--enable-ssl",
		//		"--insecure-skip-verify",
		"-R", "1",
		"-j",
		"-l", "100",
		"-h",
	})
	r := GetConfig()
	expect(t, r.HttpHost(), "127.0.0.2")
	expect(t, r.HttpPort(), uint(19301))
	expect(t, r.ReverseProxy(), []string{"192.168.0.0.2", "192.168.0.0.3"})
	expect(t, r.Mount(), "ikusan")
	expect(t, r.IrcServer(), "192.168.0.0.1")
	expect(t, r.IrcPort(), uint(6668))
	expect(t, r.IrcKeyword(), "fever")
	expect(t, r.IrcNickname(), "iku")
	expect(t, r.IrcUser(), "iku")
	expect(t, r.IrcPostInterval(), uint(1))
	expect(t, r.EnableSsl(), false)
	expect(t, r.InsecureSkipVerify(), false)
	expect(t, r.IrcReconnectInterval(), uint(1))
	expect(t, r.NoPostWithJoin(), true)
	expect(t, r.MaxLength(), uint(100))
	expect(t, r.Help(), true)
	expect(t, r.IsValid(), true)
}

func Test_SetLongArgs(t *testing.T) {
	Parse([]string{
		"--host", "127.0.0.2",
		"--port", "19301",
		"--reverse-proxy", "192.168.0.0.2", "-r", "192.168.0.0.3",
		"--mount", "ikusan",
		"--Server", "192.168.0.0.1",
		"--Port", "6668",
		"--Keyword", "fever",
		"--Nickname", "iku",
		"--User", "iku",
		"--interval", "1",
		"--enable-ssl",
		"--insecure-skip-verify",
		"--reconnect-interval", "1",
		"--no-post-with-join",
		"--max-length", "100",
		"--help",
	})
	r := GetConfig()
	expect(t, r.HttpHost(), "127.0.0.2")
	expect(t, r.HttpPort(), uint(19301))
	expect(t, r.ReverseProxy(), []string{"192.168.0.0.2", "192.168.0.0.3"})
	expect(t, r.Mount(), "ikusan")
	expect(t, r.IrcServer(), "192.168.0.0.1")
	expect(t, r.IrcPort(), uint(6668))
	expect(t, r.IrcKeyword(), "fever")
	expect(t, r.IrcNickname(), "iku")
	expect(t, r.IrcUser(), "iku")
	expect(t, r.IrcPostInterval(), uint(1))
	expect(t, r.EnableSsl(), true)
	expect(t, r.InsecureSkipVerify(), true)
	expect(t, r.IrcReconnectInterval(), uint(1))
	expect(t, r.NoPostWithJoin(), true)
	expect(t, r.MaxLength(), uint(100))
	expect(t, r.Help(), true)
	expect(t, r.IsValid(), true)
}
