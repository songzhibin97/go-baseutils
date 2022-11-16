package options

// package options: 选项模式汇总

// Option 选项模式
// Option(*Config{})
type Option[V any] func(o V)
