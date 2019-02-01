package templates

const TagListPlain = `{{ range $tag := $.Tags }}{{ $tag }}
{{ end }}`
