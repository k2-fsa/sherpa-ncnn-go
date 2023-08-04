//go:build (darwin && amd64 && !ios) || (darwin && arm64 && !ios)

package sherpa_ncnn

import (
	sherpa "github.com/k2-fsa/sherpa-ncnn-go-macos"
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
