package templates

const ImageListPlain = `{{ $showRepo := .IncludeRepository }}{{ $showTag := .IncludeTag }}{{ range $image := $.Images }}{{ if $showRepo }}{{ $image.Repository }}/{{ end }}{{ $image.Name }}{{ if $showTag }}:{{ $image.Tag }}{{ end }}
{{ end }}`
