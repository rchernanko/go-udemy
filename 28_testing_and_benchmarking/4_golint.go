package main

/*
- This is a great little tool!
- Is a linter for go source code
- https://github.com/golang/lint#purpose

"

Golint differs from gofmt. Gofmt reformats Go source code, whereas golint prints out style mistakes.

Golint differs from govet. Govet is concerned with correctness, whereas golint is concerned with coding style.
Golint is in use at Google, and it seeks to match the accepted style of the open source Go project.

The suggestions made by golint are exactly that: suggestions. Golint is not perfect, and has both false positives and
false negatives. Do not treat its output as a gold standard. We will not be adding pragmas or other knobs to suppress
specific warnings, so do not expect or require code to be completely "lint-free". In short, this tool is not, and will
never be, trustworthy enough for its suggestions to be enforced automatically, for example as part of a build process.
Golint makes suggestions for many of the mechanically checkable items listed in Effective Go and the CodeReviewComments
wiki page.

"

- So if I am at `/Users/richardchernanko/Development/go-udemy` and then run a golint ./..., I get stuff like:

```
18_application/2_json_marshal.go:15:6: exported type Club should have comment or be unexported
18_application/2_json_marshal.go:19:2: struct field IgnoreMeInJson should be IgnoreMeInJSON
18_application/3_json_unmarshal.go:12:2: struct field IgnoreMeInJson should be IgnoreMeInJSON
18_application/6_sort_custom.go:44:6: exported type ByAge should have comment or be unexported
18_application/6_sort_custom.go:58:6: exported type ByName should have comment or be unexported
19_exercises_ninja_level_8/5_exercises.go:87:6: exported type ByAge should have comment or be unexported
19_exercises_ninja_level_8/5_exercises.go:93:6: exported type ByLast should have comment or be unexported
1_introduction/1_valuable_resources.go:1:1: don't use an underscore in package name
1_introduction/2_why_go.go:1:1: don't use an underscore in package name
1_introduction/3_how_to_succeed.go:1:1: don't use an underscore in package name
22_channels/6_comma_ok_idiom_with_selects.go:45:11: if block ends with a return statement, so drop this else and outdent its block
2_course_overview/1_course_resources.go:1:1: don't use an underscore in package name
2_course_overview/2_documentation.go:1:1: don't use an underscore in package name
2_course_overview/3_accelerate_learning.go:1:1: don't use an underscore in package name
3_development_environment/come_back_to_this.go:1:1: don't use an underscore in package name
4_variables_values_and_types/1_playground/playground.go:1:1: don't use an underscore in package name
4_variables_values_and_types/7_zero_value/zero_value.go:19:14: should omit type string from declaration of var number2; it will be inferred from the right-hand side
4_variables_values_and_types/7_zero_value/zero_value.go:54:6: exported type Person should have comment or be unexported
6_programming_fundamentals/1_bool_type.go:1:1: don't use an underscore in package name
7_exercises_ninja_level_2/1_exercises.go:1:1: don't use an underscore in package name
```

Same with go vet ./...

Brilliant!

*/
