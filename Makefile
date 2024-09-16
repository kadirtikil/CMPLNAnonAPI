serve:
	@templ generate
	@go run main.go

git:
	@git add .
	@git commit -m "automated commit"
	@git push origin HEAD:main