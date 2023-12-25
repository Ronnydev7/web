atlas migrate lint \
  --dev-url "docker://postgres/15/test?search_path=public" \
  --dir="file://ent/migrations" \
  --latest=1