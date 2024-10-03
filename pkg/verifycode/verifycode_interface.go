package verifycode

type Strore interface {
	Get(key string, clear bool) string
	Set(key string, answer string) bool
	Verify(id, answer string, clear bool) bool
}
