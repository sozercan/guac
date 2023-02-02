// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Package struct {
		Namespaces func(childComplexity int) int
		Type       func(childComplexity int) int
	}

	PackageName struct {
		Name     func(childComplexity int) int
		Versions func(childComplexity int) int
	}

	PackageNamespace struct {
		Names     func(childComplexity int) int
		Namespace func(childComplexity int) int
	}

	PackageQualifier struct {
		Key   func(childComplexity int) int
		Value func(childComplexity int) int
	}

	PackageVersion struct {
		Qualifiers func(childComplexity int) int
		Subpath    func(childComplexity int) int
		Version    func(childComplexity int) int
	}

	Query struct {
		Packages func(childComplexity int, pkgSpec *model.PkgSpec) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Package.namespaces":
		if e.complexity.Package.Namespaces == nil {
			break
		}

		return e.complexity.Package.Namespaces(childComplexity), true

	case "Package.type":
		if e.complexity.Package.Type == nil {
			break
		}

		return e.complexity.Package.Type(childComplexity), true

	case "PackageName.name":
		if e.complexity.PackageName.Name == nil {
			break
		}

		return e.complexity.PackageName.Name(childComplexity), true

	case "PackageName.versions":
		if e.complexity.PackageName.Versions == nil {
			break
		}

		return e.complexity.PackageName.Versions(childComplexity), true

	case "PackageNamespace.names":
		if e.complexity.PackageNamespace.Names == nil {
			break
		}

		return e.complexity.PackageNamespace.Names(childComplexity), true

	case "PackageNamespace.namespace":
		if e.complexity.PackageNamespace.Namespace == nil {
			break
		}

		return e.complexity.PackageNamespace.Namespace(childComplexity), true

	case "PackageQualifier.key":
		if e.complexity.PackageQualifier.Key == nil {
			break
		}

		return e.complexity.PackageQualifier.Key(childComplexity), true

	case "PackageQualifier.value":
		if e.complexity.PackageQualifier.Value == nil {
			break
		}

		return e.complexity.PackageQualifier.Value(childComplexity), true

	case "PackageVersion.qualifiers":
		if e.complexity.PackageVersion.Qualifiers == nil {
			break
		}

		return e.complexity.PackageVersion.Qualifiers(childComplexity), true

	case "PackageVersion.subpath":
		if e.complexity.PackageVersion.Subpath == nil {
			break
		}

		return e.complexity.PackageVersion.Subpath(childComplexity), true

	case "PackageVersion.version":
		if e.complexity.PackageVersion.Version == nil {
			break
		}

		return e.complexity.PackageVersion.Version(childComplexity), true

	case "Query.packages":
		if e.complexity.Query.Packages == nil {
			break
		}

		args, err := ec.field_Query_packages_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Packages(childComplexity, args["pkgSpec"].(*model.PkgSpec)), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputPackageQualifierInput,
		ec.unmarshalInputPkgSpec,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../schema/package.graphql", Input: `#
# Copyright 2023 The GUAC Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# NOTE: This is experimental and might change in the future!

# Defines a GraphQL schema for the package trie/tree. This tree closely matches
# the pURL specification (https://github.com/package-url/purl-spec/blob/0dd92f26f8bb11956ffdf5e8acfcee71e8560407/README.rst)
# but deviates from it where GUAC rules state otherwise. In principle, we want
# this to represent a trie for packages, so information that represents a
# smaller collection of packages is being pushed downwards in the trie.

"""
Package represents a package.

In the pURL representation, each Package matches a ` + "`" + `pkg:<type>` + "`" + ` partial pURL.
The ` + "`" + `type` + "`" + ` field matches the pURL types but we might also use ` + "`" + `"guac"` + "`" + ` for the
cases where the pURL representation is not complete or when we have custom
rules.

This node is a singleton: backends guarantee that there is exactly one node with
the same ` + "`" + `type` + "`" + ` value.

Also note that this is named ` + "`" + `Package` + "`" + `, not ` + "`" + `PackageType` + "`" + `. This is only to make
queries more readable.
"""
type Package {
  type: String!
  namespaces: [PackageNamespace!]!
}

"""
PackageNamespace is a namespace for packages.

In the pURL representation, each PackageNamespace matches the
` + "`" + `pkg:<type>/<namespace>/` + "`" + ` partial pURL.

Namespaces are optional and type specific. Because they are optional, we use
empty string to denote missing namespaces.
"""
type PackageNamespace {
  namespace: String!
  names: [PackageName!]!
}

"""
PackageName is a name for packages.

In the pURL representation, each PackageName matches the
` + "`" + `pkg:<type>/<namespace>/<name>` + "`" + ` pURL.

Names are always mandatory.

This is the first node in the trie that can be referred to by other parts of
GUAC.
"""
type PackageName {
  name: String!
  versions: [PackageVersion!]!
}

"""
PackageVersion is a package version.

In the pURL representation, each PackageName matches the
` + "`" + `pkg:<type>/<namespace>/<name>@<version>` + "`" + ` pURL.

Versions are optional and each Package type defines own rules for handling them.
For this level of GUAC, these are just opaque strings.

This node can be referred to by other parts of GUAC.

Subpath and qualifiers are optional. Lack of qualifiers is represented by an
empty list and lack of subpath by empty string (to be consistent with
optionality of namespace and version). Two nodes that have different qualifiers
and/or subpath but the same version mean two different packages in the trie
(they are different). Two nodes that have same version but qualifiers of one are
a subset of the qualifier of the other also mean two different packages in the
trie.
"""
type PackageVersion {
  version: String!
  qualifiers: [PackageQualifier!]!
  subpath: String!
}

"""
PackageQualifier is a qualifier for a package, a key-value pair.

In the pURL representation, it is a part of the ` + "`" + `<qualifiers>` + "`" + ` part of the
` + "`" + `pkg:<type>/<namespace>/<name>@<version>?<qualifiers>` + "`" + ` pURL.

Qualifiers are optional, each Package type defines own rules for handling them,
and multiple qualifiers could be attached to the same package.

This node cannot be directly referred by other parts of GUAC.
"""
type PackageQualifier {
  key: String!
  value: String!
}

"""
PkgSpec allows filtering the list of packages to return.

Each field matches a qualifier from pURL. Use ` + "`" + `null` + "`" + ` to match on all values at
that level. For example, to get all packages in GUAC backend, use a PkgSpec
where every field is ` + "`" + `null` + "`" + `.

Empty string at a field means matching with the empty string. If passing in
qualifiers, all of the values in the list must match.
"""
input PkgSpec {
  type: String
  namespace: String
  name: String
  version: String
  qualifiers: [PackageQualifierInput!]
  subpath: String
}

"""
PackageQualifierInput is the same as PackageQualifier, but usable as query
input.

GraphQL does not allow input types to contain composite types and does not allow
composite types to contain input types. So, although in this case these two
types are semantically the same, we have to duplicate the definition.

Keys are mandatory, but values could also be ` + "`" + `null` + "`" + ` if we want to match all
values for a specific key.

TODO(mihaimaruseac): Formalize empty vs null when the schema is fully done
"""
input PackageQualifierInput {
  key: String!
  value: String
}

extend type Query {
  "Returns all packages"
  packages(pkgSpec: PkgSpec): [Package!]!
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)