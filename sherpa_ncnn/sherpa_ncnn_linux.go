//go:build (!android && linux && arm64) || (!android && linux && amd64 && !musl) || (!android && linux && arm && !arm7) || (!android && arm7) || (!android && linux && 386 && !musl) || (!android && musl) || (!android && linux && mips) || (!android && linux && mips64) || (!android && linux && mips64le) || (!android && linux && mipsle)

package sherpa_ncnn

import (
	sherpa "github.com/k2-fsa/sherpa-ncnn-go-linux"
)

type ModelConfig = sherpa.ModelConfig
type FeatureConfig = sherpa.FeatureConfig
type DecoderConfig = sherpa.DecoderConfig
type RecognizerConfig = sherpa.RecognizerConfig
type RecognizerResult = sherpa.RecognizerResult

type Recognizer = sherpa.Recognizer
type Stream = sherpa.Stream

var NewRecognizer = sherpa.NewRecognizer
var DeleteRecognizer = sherpa.DeleteRecognizer

var NewStream = sherpa.NewStream
var DeleteStream = sherpa.DeleteStream
