/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */

package janusgraph

import (
	"context"
	"fmt"

	"github.com/semi-technologies/weaviate/adapters/connectors/janusgraph/fetch"
	"github.com/semi-technologies/weaviate/adapters/connectors/janusgraph/fetchfuzzy"
	"github.com/semi-technologies/weaviate/adapters/connectors/janusgraph/gremlin"
	"github.com/semi-technologies/weaviate/usecases/traverser"
)

// LocalFetchKindClass based on GraphQL Query params
func (j *Janusgraph) LocalFetchKindClass(ctx context.Context, params *traverser.FetchParams) (interface{}, error) {
	q, err := fetch.NewQuery(*params, &j.state, &j.schema).String()
	if err != nil {
		return nil, fmt.Errorf("could not build query: %s", err)
	}

	res, err := fetch.NewProcessor(j.client, params.Kind, "localhost", &j.state).
		Process(ctx, gremlin.New().Raw(q))
	return res, err
}

// LocalFetchFuzzy based on GraphQL Query params
func (j *Janusgraph) LocalFetchFuzzy(ctx context.Context, words []string) (interface{}, error) {
	q, err := fetchfuzzy.NewQuery(words, &j.state, &j.schema).String()
	if err != nil {
		return nil, fmt.Errorf("could not build query: %s", err)
	}

	res, err := fetchfuzzy.NewProcessor(j.client, "localhost", &j.state).
		Process(ctx, gremlin.New().Raw(q))
	return res, err
}
