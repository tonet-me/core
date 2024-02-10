package richerror

type RichErrorOptions struct {
	Op         OP
	Kind       Kind
	Message    string
	Meta       interface{}
	InnerError error
}

type RichErrorOption func(options *RichErrorOptions)

func WithOp(op OP) RichErrorOption {
	return func(options *RichErrorOptions) {
		options.Op = op
	}
}

func WithMessage(msg string) RichErrorOption {
	return func(options *RichErrorOptions) {
		options.Message = msg
	}
}

func WithKind(kind Kind) RichErrorOption {
	return func(options *RichErrorOptions) {
		options.Kind = kind
	}
}

func WithInnerError(err error) RichErrorOption {
	return func(options *RichErrorOptions) {
		options.InnerError = err
	}
}

func WithMeta(meta interface{}) RichErrorOption {
	return func(options *RichErrorOptions) {
		options.Meta = meta
	}
}
