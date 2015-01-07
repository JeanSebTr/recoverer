# [recoverer](http://en.wiktionary.org/wiki/recoverer) – Noun

[Agent noun](http://en.wiktionary.org/wiki/agent_noun) of [recover](http://golang.org/pkg/builtin/#recover); one who recovers.

## What?

Very small golang lib to catch panic.

## Should I use it?

Probably not.

## Why?

It's not idiomatic go to catch errors.
Every function which can possily result in an error should explicitly return an error.

This lib is only a simple proof of how well engineered golang is.

## How?

```go
func FuncWhichMustReturnError() (err error) {
	defer recoverer.Catch(&err)

	FuncWhichCanPanic()

	return err
}
```

## How??

A @$#^%#$ 5 liner with a link to documentation in the README…
