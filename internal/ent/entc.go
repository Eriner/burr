//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
)

func main() {
	if err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeaturePrivacy,
			gen.FeatureSnapshot,
			gen.FeatureLock,
		},
	}, entc.Extensions(entviz.Extension{})); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
