atlas migrate diff \
  --dir "file://ent/migrations" \
  --to "ent://ent/schema?globalid=1" \
  --dev-url "docker://postgres/15/test?search_path=public"