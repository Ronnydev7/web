atlas migrate apply \
  --dir "file://ent/migrations" \
  --url "$DB_URL?search_path=public&sslmode=disable"