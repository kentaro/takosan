release:
	gox -output "pkg/{{.Dir}}_{{.OS}}_{{.Arch}}"; ghr -u kentaro -r takosan --delete $(VERSION) pkg/
