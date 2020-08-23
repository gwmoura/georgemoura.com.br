deploy-images:
	gsutil cp -n -r ./images/* gs://george-moura-site.appspot.com/
deploy: deploy-images
	./scripts/deploy
tests:
	./scripts/run_tests.sh
serve:
	./scripts/serve.sh
start: serve
