[[_TOC_]]

## What is Viper?

It supports:

- setting defaults.
- reading from JSON, TOML, YAML, HCL, envfile and Java properties config files.
- live watching and re-reading of config files (optional).
- reading from environment variables.
- reading from remote config systems (etcd or Consul), and watching changes.
- reading from command line flags.
- reading from buffer.
- setting explicit values.

## Why Viper?

Viper uses the following precedence order. Each item takes precedence over the item below it:

- explicit call to set
- flag
- env
- config 
- key/value store
- default

## Putting Values into Viper

**Establish Defaults**

```go
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
```

## Working with Environment Variables

Viper has full support for environment variables. This enables 12 factor applications out of the box. There are five methods that exist to aid working with ENV:

- AutomaticEnv()
- BindEnv(string...): error
- SetEnvPrefix(string)
- SetEnvReplacer(string...) *string.Replacer
- AllowEmptyEnv(bool)

When working with ENV variables, it's important to recognize that Viper treats ENV variables as case sensitive.

Viper provides a mechanism to try to ensure that ENV variables are unique. By using `SetEnvPrefix`, you can tell Viper to use a prefix while reading from the environment variables. Both `BindEnv` and `AutomaticEnv` will use this prefix.

`BindEnv` takes one or more parameters
