package dice

type RandomValueGen func(min int, max int, num int) []int

type Options struct {
	RandomValueGen RandomValueGen
}

type Option func(*Options)

func SetRandomValueGen(f func(min int, max int, num int) []int) Option {
	return func(args *Options) {
		args.RandomValueGen = f
	}
}
