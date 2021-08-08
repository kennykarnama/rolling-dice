package player

type Options struct {
	InitialScore float64
}

type Option func(*Options)

func SetInitialScore(initialScore float64) Option {
	return func(args *Options) {
		args.InitialScore = initialScore
	}
}
