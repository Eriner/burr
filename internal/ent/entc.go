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
			// TODO: determine how to embed or handle using atlas to provide
			// versioned migrations between versions. Sequential updates aren't
			// guaranteed and recording migrations on new versions is probably
			// what we want long term. The auto-migration implementation is great
			// for development, it just isn't fit for released software.
			// gen.FeatureVersionedMigration,
		},
	}, entc.Extensions(entviz.Extension{})); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
