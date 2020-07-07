package main

import (
	"flag"
	"fmt"
	"github.com/couchbase/indexing/secondary/tests/framework/kvutility"
	"os"
)

var options struct {
	bucket        string // bucket to connect
	username      string
	password      string
	kvaddress     string
	prefix        string
	numColls      int
	collsPerScope int
}

func argParse() {

	flag.StringVar(&options.bucket, "bucket", "default",
		"buckets to connect")
	flag.StringVar(&options.username, "username", "Administrator",
		"Cluster username")
	flag.StringVar(&options.password, "password", "asdasd",
		"Cluster password")
	flag.StringVar(&options.kvaddress, "kvaddress", "127.0.0.1:8091",
		"KV address")
	flag.StringVar(&options.prefix, "prefix", "test",
		"Prefix for scope and collection names")
	flag.IntVar(&options.numColls, "numCollections", 1,
		"Number of collections that has to be created")
	flag.IntVar(&options.collsPerScope, "collectionsPerScope", 1,
		"Number of collections per scpoe")
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage : %s [OPTIONS] <addr> \n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	argParse()
	fmt.Printf("Username: %v, password: %v, prefix: %v, numColls: %v, collectionsPerScope: %v\n", options.username, options.password, options.prefix, options.numColls, options.collsPerScope)

	var scope, collection string
	collCounter := 1
	scopeCounter := 1
	for i := 0; i < options.numColls; i++ {
		if i > options.collsPerScope {
			collCounter = 1
			scope = options.prefix + fmt.Sprintf("_scope_%v", scopeCounter)
		}
		collection = options.prefix + fmt.Sprintf("_collection_%v", collCounter)
		collCounter++
		kvutility.CreateCollection(options.bucket, scope, collection, options.username, options.password, options.kvaddress)
	}

}
