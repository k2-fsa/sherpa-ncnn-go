//go:build (windows && amd64) || (windows && 386)

package sherpa_ncnn

import (
	sherpa "github.com/k2-fsa/sherpa-ncnn-go-windows"
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
